package service

import (
	"github.com/stretchr/testify/assert"
	"kirmac-site-backend/domain"
	"kirmac-site-backend/services"
	"kirmac-site-backend/services/model"
	"testing"
)

var propertyService services.IPropertyService

func TestMain(m *testing.M) {
	initialProperties := []domain.Property{
		{
			ID:          3,
			Location:    "Antalya, Turkey",
			Price:       1800000,
			Title:       "Seaside Penthouse in Antalya",
			Description: "Stunning penthouse apartment with panoramic sea views in the beautiful coastal city of Antalya.",
			Bedrooms:    3,
			Bathrooms:   2,
			SquareFeet:  2200,
			AgentName:   "Ayse Kaya",
			AgentTitle:  "Luxury Property Specialist",
			ImageURLs:   []string{"https://example.com/antalya_penthouse1.jpg", "https://example.com/antalya_penthouse2.jpg", "https://example.com/antalya_penthouse3.jpg"},
		},
		{
			ID:          4,
			Location:    "Bodrum, Turkey",
			Price:       3500000,
			Title:       "Luxury Beach Villa in Bodrum",
			Description: "Stunning beachfront villa with private pool and direct access to the Aegean Sea.",
			Bedrooms:    6,
			Bathrooms:   5,
			SquareFeet:  5000,
			AgentName:   "Mehmet Yilmaz",
			AgentTitle:  "Luxury Property Consultant",
			ImageURLs:   []string{"https://example.com/bodrum_villa1.jpg", "https://example.com/bodrum_villa2.jpg"},
		},
		{
			ID:          5,
			Location:    "Ankara, Turkey",
			Price:       800000,
			Title:       "Modern City Apartment",
			Description: "Centrally located modern apartment with panoramic city views in Ankara.",
			Bedrooms:    3,
			Bathrooms:   2,
			SquareFeet:  1500,
			AgentName:   "Zeynep Kaya",
			AgentTitle:  "City Center Specialist",
			ImageURLs:   []string{"https://example.com/ankara_apt1.jpg", "https://example.com/ankara_apt2.jpg"},
		},
		{
			ID:          6,
			Location:    "Izmir, Turkey",
			Price:       1200000,
			Title:       "Seaside Condo in Izmir",
			Description: "Beautiful condo with sea view, located in the vibrant Alsancak district of Izmir.",
			Bedrooms:    4,
			Bathrooms:   3,
			SquareFeet:  2000,
			AgentName:   "Can Demir",
			AgentTitle:  "Izmir Coast Expert",
			ImageURLs:   []string{"https://example.com/izmir_condo1.jpg", "https://example.com/izmir_condo2.jpg"},
		},
		{
			ID:          7,
			Location:    "Cappadocia, Turkey",
			Price:       950000,
			Title:       "Unique Cave House in Cappadocia",
			Description: "One-of-a-kind cave house with modern amenities in the heart of Cappadocia.",
			Bedrooms:    2,
			Bathrooms:   2,
			SquareFeet:  1200,
			AgentName:   "Ayse Yildiz",
			AgentTitle:  "Cappadocia Property Specialist",
			ImageURLs:   []string{"https://example.com/cappadocia_cave1.jpg", "https://example.com/cappadocia_cave2.jpg"},
		},
		{
			ID:          10,
			Location:    "Trabzon, Turkey",
			Price:       750000,
			Title:       "Black Sea View Apartment",
			Description: "Modern apartment with stunning Black Sea views in Trabzon.",
			Bedrooms:    3,
			Bathrooms:   2,
			SquareFeet:  1600,
			AgentName:   "Emre Sahin",
			AgentTitle:  "Black Sea Region Specialist",
			ImageURLs:   []string{"https://example.com/trabzon_apt1.jpg", "https://example.com/trabzon_apt2.jpg"},
		},
		{
			ID:          11,
			Location:    "Alanya, Turkey",
			Price:       450000,
			Title:       "Beachfront Studio in Alanya",
			Description: "Cozy beachfront studio apartment in the popular tourist destination of Alanya.",
			Bedrooms:    1,
			Bathrooms:   1,
			SquareFeet:  600,
			AgentName:   "Selin Aydin",
			AgentTitle:  "Alanya Beach Property Expert",
			ImageURLs:   []string{"https://example.com/alanya_studio1.jpg", "https://example.com/alanya_studio2.jpg"},
		},
		{
			ID:          12,
			Location:    "Eskisehir, Turkey",
			Price:       350000,
			Title:       "Student-Friendly Apartment",
			Description: "Modern apartment ideal for students, close to university campuses in Eskisehir.",
			Bedrooms:    2,
			Bathrooms:   1,
			SquareFeet:  900,
			AgentName:   "Ahmet Celik",
			AgentTitle:  "Student Housing Specialist",
			ImageURLs:   []string{"https://example.com/eskisehir_apt1.jpg", "https://example.com/eskisehir_apt2.jpg"},
		},
		{
			ID:          13,
			Location:    "Cesme, Turkey",
			Price:       2200000,
			Title:       "Luxury Beach House in Cesme",
			Description: "Elegant beach house with private garden and pool in the exclusive Cesme Peninsula.",
			Bedrooms:    4,
			Bathrooms:   3,
			SquareFeet:  2800,
			AgentName:   "Deniz Korkmaz",
			AgentTitle:  "Cesme Luxury Property Advisor",
			ImageURLs:   []string{"https://example.com/cesme_house1.jpg", "https://example.com/cesme_house2.jpg"},
		},
		{
			ID:          14,
			Location:    "Cesme, Turkey",
			Price:       2200000,
			Title:       "Luxury Beach House in Cesme",
			Description: "Elegant beach house with private garden and pool in the exclusive Cesme Peninsula.",
			Bedrooms:    4,
			Bathrooms:   3,
			SquareFeet:  2800,
			AgentName:   "Deniz Korkmaz",
			AgentTitle:  "Cesme Luxury Property Advisor",
			ImageURLs:   []string{"https://example.com/cesme_house1.jpg", "https://example.com/cesme_house2.jpg"},
		},
		{
			ID:          8,
			Location:    "Bursa, Turkey",
			Price:       600000,
			Title:       "Traditional Ottoman House",
			Description: "Beautifully restored Ottoman-era house in the historic district of Bursa.",
			Bedrooms:    10,
			Bathrooms:   10,
			SquareFeet:  1800,
			AgentName:   "Leyla Ozturk",
			AgentTitle:  "Historical Property Consultant",
			ImageURLs:   []string{"https://example.com/bursa_ottoman1.jpg", "https://example.com/bursa_ottoman2.jpg"},
		},
	}
	fakePropertyRepository := NewFakePropertyRepository(initialProperties)
	propertyService = services.NewPropertyService(fakePropertyRepository)
}

// TestGetAllProperties tests the GetAllProperties method of the PropertyService
func TestGetAllProperties(t *testing.T) {
	actualProperties, err := propertyService.GetAllProperties()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Run("TestGetAllProperties", func(t *testing.T) {
		assert.Equal(t, 11, len(actualProperties))
	})
}

// TestGetPropertyById tests the GetPropertyById method of the PropertyService
func TestGetPropertyById(t *testing.T) {
	actualProperty, err := propertyService.GetPropertyById(3)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Run("TestGetPropertyById", func(t *testing.T) {
		expectedProperty := domain.Property{
			ID:          3,
			Location:    "Antalya, Turkey",
			Price:       1800000,
			Title:       "Seaside Penthouse in Antalya",
			Description: "Stunning penthouse apartment with panoramic sea views in the beautiful coastal city of Antalya.",
			Bedrooms:    3,
			Bathrooms:   2,
			SquareFeet:  2200,
			AgentName:   "Ayse Kaya",
			AgentTitle:  "Luxury Property Specialist",
			ImageURLs:   []string{"https://example.com/antalya_penthouse1.jpg", "https://example.com/antalya_penthouse2.jpg", "https://example.com/antalya_penthouse3.jpg"},
		}
		assert.Equal(t, expectedProperty, actualProperty)
	})
}

// TestAddProperty tests the AddProperty method of the PropertyService
func TestAddProperty(t *testing.T) {
	property := model.PropertyCreate{
		Location:    "Istanbul, Turkey",
		Price:       2500000,
		Title:       "Luxury Bosphorus Mansion",
		Description: "Historic mansion with panoramic Bosphorus views in the heart of Istanbul.",
		Bedrooms:    8,
		Bathrooms:   6,
		SquareFeet:  8000,
		AgentName:   "Ali Yilmaz",
		AgentTitle:  "Bosphorus Property Specialist",
		ImageURLs:   []string{"https://example.com/istanbul_mansion1.jpg", "https://example.com/istanbul_mansion2.jpg"},
	}
	err := propertyService.AddProperty(property)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Run("TestAddProperty", func(t *testing.T) {
		actualProperties, err := propertyService.GetAllProperties()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		assert.Equal(t, 12, len(actualProperties))
	})
}
