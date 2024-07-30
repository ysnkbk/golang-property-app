package model

type PropertyCreate struct {
	Location    string   `json:"location"`
	Price       int      `json:"price"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Bedrooms    int      `json:"bedrooms"`
	Bathrooms   int      `json:"bathrooms"`
	SquareFeet  int      `json:"square_feet"`
	AgentName   string   `json:"agent_name"`
	AgentTitle  string   `json:"agent_title"`
	ImageURLs   []string `json:"image_urls"`
}
