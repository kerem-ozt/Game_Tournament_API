package services

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"

	"github.com/kamva/mgm/v3"
	db "github.com/kerem-ozt/GoodBlast_API/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateTournament create new tournament record
// func CreateTournament(participants []primitive.ObjectID) (*db.Tournament, error) {
func CreateTournament(participants ...primitive.ObjectID) (*db.Tournament, error) {
	tournament := db.NewTournament(participants)
	err := mgm.Coll(tournament).Create(tournament)
	if err != nil {
		return nil, errors.New("cannot create new tournament")
	}

	return tournament, nil
}

// GetTournaments get paginated tournaments list
// func GetTournaments(userId primitive.ObjectID, page int, limit int) ([]db.Tournament, error) {
func GetTournaments(page int, limit int) ([]db.Tournament, error) {
	var tournaments []db.Tournament

	findOptions := options.Find().
		SetSkip(int64(page * limit)).
		SetLimit(int64(limit + 1))

	err := mgm.Coll(&db.Tournament{}).SimpleFind(
		&tournaments,
		bson.M{},
		findOptions,
	)

	if err != nil {
		return nil, errors.New("cannot find tournaments")
	}

	return tournaments, nil
}

func GetTournamentById(tournamentId primitive.ObjectID) (*db.Tournament, error) {
	tournament := &db.Tournament{}
	err := mgm.Coll(tournament).FindByID(tournamentId, tournament)
	if err != nil {
		return nil, errors.New("cannot find tournament")
	}

	return tournament, nil
}

func ProgressTournament(tournamentID primitive.ObjectID) ([]primitive.ObjectID, error) {
	tournament := &db.Tournament{}

	type Participant struct {
		ID   primitive.ObjectID `bson:"id"`
		Rank int
	}

	err := mgm.Coll(tournament).FindByID(tournamentID, tournament)
	if err != nil {
		return nil, errors.New("cannot find tournament")
	}

	var participants []Participant
	for _, objID := range tournament.Participants {
		id, err := primitive.ObjectIDFromHex(objID.Hex())
		if err != nil {
			return nil, errors.New("invalid participant ID")
		}
		participants = append(participants, Participant{ID: id, Rank: 0})
	}

	for round := 1; len(tournament.Participants) > 1; round++ {
		for i := len(tournament.Participants) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			tournament.Participants[i], tournament.Participants[j] = tournament.Participants[j], tournament.Participants[i]
		}

		winnerCount := len(tournament.Participants) / 2

		winnersSlice := tournament.Participants[:winnerCount]

		tournament.Participants = tournament.Participants[:winnerCount]

		var winners []Participant
		for _, winner := range winnersSlice {
			for j := range participants {
				if participants[j].ID == winner {
					participants[j].Rank = round
					winners = append(winners, participants[j])
					break
				}
			}
		}

		fmt.Println("Round", round, "Winners:", winners)
	}

	for _, participant := range participants {
		err := UpdateUserStat(participant.ID, participant.Rank*100, 0)
		if err != nil {
			return nil, errors.New("cannot update user progress")
		}
	}

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Rank > participants[j].Rank
	})

	top3winnerIDs := []primitive.ObjectID{participants[0].ID, participants[1].ID, participants[2].ID}

	for i, reward := range []int{5000, 3000, 2000} {
		err := UpdateUserStat(top3winnerIDs[i], 0, reward)
		if err != nil {
			return nil, errors.New("cannot update user progress")
		}
	}

	return top3winnerIDs, nil
}
