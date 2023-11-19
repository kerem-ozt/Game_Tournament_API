// models/leaderboard.go

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Leaderboard represents the leaderboard data stored in MongoDB
type Leaderboard struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Type  string             `bson:"type"` // global, country, or tournament
	Users []LeaderboardUser  `bson:"users"`
}

// GetID implements mgm.Model.
func (l *Leaderboard) GetID() interface{} {
	return l.ID
}

// PrepareID implements mgm.Model.
func (l *Leaderboard) PrepareID(id interface{}) (interface{}, error) {
	return primitive.ObjectIDFromHex(id.(string))
}

// SetID implements mgm.Model.
func (l *Leaderboard) SetID(id interface{}) {
	l.ID = id.(primitive.ObjectID)
}

// LeaderboardUser represents a user in the leaderboard
type LeaderboardUser struct {
	UserID   primitive.ObjectID `bson:"userId"`
	Progress int                `bson:"progress"`
	Country  string             `bson:"country"`
}
