package service

import (
	"kirmac-site-backend/domain"
)

type FakePropertyRepository struct {
	properties []domain.Property
}

func NewFakePropertyRepository(initialProperty []domain.Property) *FakePropertyRepository {
	return &FakePropertyRepository{
		properties: initialProperty,
	}
}

func (repository *FakePropertyRepository) GetAllProperties() ([]domain.Property, error) {
	return repository.properties, nil
}

func (repository *FakePropertyRepository) GetPropertyById(id int64) (domain.Property, error) {
	for _, property := range repository.properties {
		if property.ID == id {
			return property, nil
		}
	}
	return domain.Property{}, nil
}

func (repository *FakePropertyRepository) AddProperty(property domain.Property) domain.Property {
	repository.properties = append(repository.properties, property)
	return property
}

func (repository *FakePropertyRepository) DeleteById(id int64) (bool, error) {
	for i, property := range repository.properties {
		if property.ID == id {
			repository.properties = append(repository.properties[:i], repository.properties[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (repository *FakePropertyRepository) UpdateProperty(id int64, property domain.Property) error {
	for i, p := range repository.properties {
		if p.ID == id {
			repository.properties[i] = property
			return nil
		}
	}
	return nil
}
