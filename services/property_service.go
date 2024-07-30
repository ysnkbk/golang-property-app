package services

import (
	"errors"
	"kirmac-site-backend/domain"
	"kirmac-site-backend/persistence"
	"kirmac-site-backend/services/model"
)

// IPropertyService defines the service interface for property operations
type IPropertyService interface {
	GetAllProperties() ([]model.PropertyCreate, error)
	GetPropertyById(id int64) (model.PropertyCreate, error)
	AddProperty(property model.PropertyCreate) error
	UpdateProperty(id int64, property model.PropertyCreate) error
	DeleteById(id int64) (bool, error)
}

// PropertyService implements IPropertyService and provides business logic for property operations
type PropertyService struct {
	repository persistence.IPropertyRepository
}

// NewPropertyService creates a new instance of PropertyService
func NewPropertyService(repository persistence.IPropertyRepository) *PropertyService {
	return &PropertyService{
		repository: repository,
	}
}

// GetAllProperties retrieves all properties
func (service *PropertyService) GetAllProperties() ([]model.PropertyCreate, error) {
	properties, err := service.repository.GetAllProperties()
	if err != nil {
		return nil, err
	}
	var propertyModels []model.PropertyCreate
	for _, property := range properties {
		propertyModels = append(propertyModels, model.PropertyCreate{
			Location:    property.Location,
			Price:       property.Price,
			Title:       property.Title,
			Description: property.Description,
			Bedrooms:    property.Bedrooms,
			Bathrooms:   property.Bathrooms,
			SquareFeet:  property.SquareFeet,
			AgentName:   property.AgentName,
			AgentTitle:  property.AgentTitle,
			ImageURLs:   property.ImageURLs,
		})
	}
	return propertyModels, nil
}

// GetPropertyById retrieves a property by id
func (service *PropertyService) GetPropertyById(id int64) (model.PropertyCreate, error) {
	property, err := service.repository.GetPropertyById(id)
	if err != nil {
		return model.PropertyCreate{}, err
	}
	return model.PropertyCreate{
		Location:    property.Location,
		Price:       property.Price,
		Title:       property.Title,
		Description: property.Description,
		Bedrooms:    property.Bedrooms,
		Bathrooms:   property.Bathrooms,
		SquareFeet:  property.SquareFeet,
		AgentName:   property.AgentName,
		AgentTitle:  property.AgentTitle,
		ImageURLs:   property.ImageURLs,
	}, nil
}

// AddProperty adds a new property
func (service *PropertyService) AddProperty(property model.PropertyCreate) error {
	err := validateProperty(property)
	if err != nil {
		return err
	}
	service.repository.AddProperty(domain.Property{
		Location:    property.Location,
		Price:       property.Price,
		Title:       property.Title,
		Description: property.Description,
		Bedrooms:    property.Bedrooms,
		Bathrooms:   property.Bathrooms,
		SquareFeet:  property.SquareFeet,
		AgentName:   property.AgentName,
		AgentTitle:  property.AgentTitle,
		ImageURLs:   property.ImageURLs,
	})
	return nil
}

// UpdateProperty updates a property
func (service *PropertyService) UpdateProperty(id int64, property model.PropertyCreate) error {
	err := validateProperty(property)
	if err != nil {
		return err
	}

	return service.repository.UpdateProperty(id, domain.Property{
		Location:    property.Location,
		Price:       property.Price,
		Title:       property.Title,
		Description: property.Description,
		Bedrooms:    property.Bedrooms,
		Bathrooms:   property.Bathrooms,
		SquareFeet:  property.SquareFeet,
		AgentName:   property.AgentName,
		AgentTitle:  property.AgentTitle,
		ImageURLs:   property.ImageURLs,
	})
}

// DeleteById deletes a property by id
func (service *PropertyService) DeleteById(id int64) (bool, error) {
	return service.repository.DeleteById(id)
}

func validateProperty(property model.PropertyCreate) error {
	if property.Price <= 0 {
		return errors.New("Price must be greater than zero")
	}
	return nil
}
