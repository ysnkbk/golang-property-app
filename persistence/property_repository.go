package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"kirmac-site-backend/domain"
)

const (
	getAllPropertiesQuery = `SELECT id, location, price, title, description, bedrooms, bathrooms, square_feet, agent_name, agent_title, image_urls FROM properties`
	getPropertyByIdQuery  = `SELECT id, location, price, title, description, bedrooms, bathrooms, square_feet, agent_name, agent_title, image_urls FROM properties WHERE id = $1`
	addPropertyQuery      = `INSERT INTO properties (location, price, title, description, bedrooms, bathrooms, square_feet, agent_name, agent_title, image_urls) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	deletePropertyQuery   = `DELETE FROM properties WHERE id = $1`
	updatePropertyQuery   = `UPDATE properties SET location = $1, price = $2, title = $3, description = $4, bedrooms = $5, bathrooms = $6, square_feet = $7, agent_name = $8, agent_title = $9, image_urls = $10 WHERE id = $11`
)

// IPropertyRepository is an interface for the property repository
type IPropertyRepository interface {
	GetAllProperties() ([]domain.Property, error)
	GetPropertyById(id int64) (domain.Property, error)
	AddProperty(property domain.Property) domain.Property
	DeleteById(id int64) (bool, error)
	UpdateProperty(id int64, property domain.Property) error
}

// PropertyRepository is a struct for the property repository
type PropertyRepository struct {
	dbPool *pgxpool.Pool
}

// NewPropertyRepository creates a new property repository
func NewPropertyRepository(dbPool *pgxpool.Pool) IPropertyRepository {
	return &PropertyRepository{dbPool: dbPool}
}

func (propertyRepository *PropertyRepository) GetAllProperties() ([]domain.Property, error) {
	ctx := context.Background()
	propertiesRows, err := propertyRepository.dbPool.Query(ctx, getAllPropertiesQuery)
	if err != nil {
		log.Errorf("Veritabanı sorgusu hatası: %v", err)
		return nil, err
	}
	defer propertiesRows.Close()

	properties, err := propertyRepository.scanProperties(propertiesRows)
	if err != nil {
		log.Errorf("Satırları tarama hatası: %v", err)
		return nil, err
	}

	if len(properties) == 0 {
		log.Info("Hiç mülk bulunamadı")
	}

	return properties, nil
}

// GetPropertyById gets a property by id
func (propertyRepository *PropertyRepository) GetPropertyById(id int64) (domain.Property, error) {
	var p domain.Property

	ctx := context.Background()
	row := propertyRepository.dbPool.QueryRow(ctx, getPropertyByIdQuery, id)
	err := row.Scan(
		&p.ID,
		&p.Location,
		&p.Price,
		&p.Title,
		&p.Description,
		&p.Bedrooms,
		&p.Bathrooms,
		&p.SquareFeet,
		&p.AgentName,
		&p.AgentTitle,
		pq.Array(&p.ImageURLs),
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.Property{}, fmt.Errorf("property not found: %v", err)
		}
		return domain.Property{}, fmt.Errorf("unable to read row: %v", err)
	}

	return p, nil
}

// AddProperty adds a property
func (propertyRepository *PropertyRepository) AddProperty(property domain.Property) domain.Property {
	ctx := context.Background()
	var id int64
	err := propertyRepository.dbPool.QueryRow(ctx, addPropertyQuery, property.Location, property.Price, property.Title, property.Description, property.Bedrooms, property.Bathrooms, property.SquareFeet, property.AgentName, property.AgentTitle, pq.Array(property.ImageURLs)).Scan(&id)
	if err != nil {
		log.Errorf("Unable to add property: %v\n", err)
		return domain.Property{} // Consider handling the error more gracefully
	}
	property.ID = id
	return property
}

// DeleteById deletes a property by id
func (propertyRepository *PropertyRepository) DeleteById(id int64) (bool, error) {
	ctx := context.Background()
	cmdTag, err := propertyRepository.dbPool.Exec(ctx, deletePropertyQuery, id)
	if err != nil {
		return false, fmt.Errorf("unable to delete property: %v", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return false, nil // No rows affected, but also no error
	}
	return true, nil
}

// UpdateProperty updates a property
func (propertyRepository *PropertyRepository) UpdateProperty(id int64, property domain.Property) error {
	ctx := context.Background()

	_, err := propertyRepository.dbPool.Exec(ctx, updatePropertyQuery,
		property.Location,
		property.Price,
		property.Title,
		property.Description,
		property.Bedrooms,
		property.Bathrooms,
		property.SquareFeet,
		property.AgentName,
		property.AgentTitle,
		pq.Array(property.ImageURLs),
		id)
	if err != nil {
		return fmt.Errorf("unable to update property: %v", err)
	}
	_, err = propertyRepository.DeleteById(id)
	if err != nil {
		return fmt.Errorf("unable to delete property: %v", err)
	}
	propertyRepository.AddProperty(property)
	return nil
}

// scanProperties scans properties
func (propertyRepository *PropertyRepository) scanProperties(rows pgx.Rows) ([]domain.Property, error) {
	var properties []domain.Property
	for rows.Next() {
		var p domain.Property
		err := rows.Scan(&p.ID, &p.Location, &p.Price, &p.Title, &p.Description, &p.Bedrooms, &p.Bathrooms, &p.SquareFeet, &p.AgentName, &p.AgentTitle, pq.Array(&p.ImageURLs))
		if err != nil {
			log.Errorf("Error readinig scaining rows: %v", err)
			return nil, err
		}
		properties = append(properties, p)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("Error scaninng row: %v", err)
		return nil, err
	}
	return properties, nil
}

func (propertyRepository *PropertyRepository) scanProperty(row pgx.Row) (domain.Property, error) {
	var p domain.Property
	var imageURLs string // Assuming imageURLs are stored as a JSON-encoded string in the database
	err := row.Scan(&p.ID, &p.Location, &p.Price, &p.Title, &p.Description, &p.Bedrooms, &p.Bathrooms, &p.SquareFeet, &p.AgentName, &p.AgentTitle, &imageURLs)
	if err != nil {
		log.Errorf("Error scanning row: %v\n", err)
		return domain.Property{}, err
	}
	// Decode JSON-encoded imageURLs string to []string
	if err := json.Unmarshal([]byte(imageURLs), &p.ImageURLs); err != nil {
		log.Errorf("Error decoding image URLs: %v\n", err)
		return domain.Property{}, err
	}
	return p, nil
}
