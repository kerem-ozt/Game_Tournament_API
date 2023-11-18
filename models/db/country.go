// models/country.go

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Country represents the country data stored in MongoDB
type Country struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty"`
	Name  string               `bson:"name"`
	Users []primitive.ObjectID `bson:"users"`
}
