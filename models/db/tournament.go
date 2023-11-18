package models

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var StartTime = time.Now().Truncate(24 * time.Hour)
var EndTime = StartTime.Add(24 * time.Hour)

type Tournament struct {
	mgm.DefaultModel `bson:",inline"`
	StartTime        time.Time `bson:"startTime"`
	EndTime          time.Time `bson:"endTime"`
	MinLevels        int       `bson:"minLevels"`
	EntryFee         int       `bson:"entryFee"`
	MaxParticipants  int       `bson:"maxParticipants"`
	// Participants     []primitive.ObjectID `bson:"participants"`
	Participants []primitive.ObjectID `json:"participants" binding:"required"`
	Scores       []TournamentScore    `bson:"scores"`
}

type TournamentScore struct {
	UserID primitive.ObjectID `bson:"userId"`
	Score  int                `bson:"score"`
}

func NewTournament(participants []primitive.ObjectID) *Tournament {
	return &Tournament{
		StartTime:       StartTime,
		EndTime:         EndTime,
		MinLevels:       10,
		EntryFee:        500,
		MaxParticipants: 35,
		// Participants:    []primitive.ObjectID{},
		Participants: participants,
		Scores:       []TournamentScore{},
	}
}

func (model *Tournament) CollectionName() string {
	return "tournaments"
}
