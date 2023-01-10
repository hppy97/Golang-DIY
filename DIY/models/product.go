package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model for Product - file
type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"Name,omitempty"`
	Description string             `json:"Description,omitempty"`
	Price       float64            `json:"Price,omitempty"`
	Quantity    int64              `json:"Quantity,omitempty"`
}
