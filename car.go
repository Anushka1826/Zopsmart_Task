package models

// Car represents the structure of a car entity
type Car struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Make     string `json:"make" bson:"make"`
	Model    string `json:"model" bson:"model"`
	Year     int    `json:"year" bson:"year"`
	Owner    string `json:"owner" bson:"owner"`
	InGarage bool   `json:"in_garage" bson:"in_garage"`
}
