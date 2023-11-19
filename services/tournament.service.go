package services

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"

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

func FindTournamentByStartDateToday() (*db.Tournament, error) {
	// Get the current date in UTC
	now := time.Now().UTC()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// Find the tournament with the matching start_date
	tournament := &db.Tournament{}
	err := mgm.Coll(tournament).First(bson.M{"startTime": startOfDay}, tournament)

	if err != nil {
		return nil, errors.New("cannot find tournament")
	}

	return tournament, nil
}

func CreateTournamentGroups() ([]db.TournamentGroup, error) {
	groups := make([]db.TournamentGroup, 0)

	group := db.TournamentGroup{
		GroupID:      primitive.NewObjectID(),
		Participants: []primitive.ObjectID{}, // Empty participants for now
	}

	groups = append(groups, group)

	todayTournament, err := FindTournamentByStartDateToday()

	if err != nil {
		return nil, err
	}

	if todayTournament == nil {
		return nil, errors.New("no tournament found for today")
	}

	// Add the new groups to the existing groups
	todayTournament.Groups = append(todayTournament.Groups, groups...)

	// Save the updated tournament to the database
	err = mgm.Coll(todayTournament).Update(todayTournament)
	if err != nil {
		return nil, errors.New("cannot update tournament with groups: " + err.Error())
	}

	// Check if groups is not empty before returning
	if len(groups) > 0 {
		return groups, nil
	}

	return nil, errors.New("no groups created")
}

func CreateTournamentGroups0(participants []primitive.ObjectID) ([]db.TournamentGroup, error) {
	groups := make([]db.TournamentGroup, 0)
	group := db.TournamentGroup{
		GroupID:      primitive.NewObjectID(),
		Participants: participants,
	}
	groups = append(groups, group)

	fmt.Println("Groups to be added:", groups)

	// Split participants into groups of MaxParticipants
	// for i := 0; i < len(participants); i += db.MaxParticipants {
	// 	end := i + db.MaxParticipants
	// 	if end > len(participants) {
	// 		end = len(participants)
	// 	}

	// 	group := db.TournamentGroup{
	// 		GroupID:      primitive.NewObjectID(),
	// 		Participants: participants[i:end],
	// 	}

	// 	groups = append(groups, group)
	// }

	todayTournament, err := FindTournamentByStartDateToday()

	fmt.Println("todayTournament", todayTournament)

	if err != nil {
		return nil, err
	}

	if todayTournament == nil {
		return nil, errors.New("no tournament found for today")
	}

	fmt.Println("Groups to be added:", groups)

	todayTournament.Groups = groups
	err = mgm.Coll(todayTournament).Update(todayTournament)
	if err != nil {
		return nil, errors.New("cannot update tournament with groups")
	}

	return groups, nil

	// tournament := &db.Tournament{}
	// err := mgm.Coll(tournament).FindByID(tournamentID, tournament)
	// if err != nil {
	// 	return errors.New("cannot find tournament")
	// }

	// // Update the tournament with the new groups
	// tournament.Groups = groups
	// err = mgm.Coll(tournament).Update(tournament)
	// if err != nil {
	// 	return errors.New("cannot update tournament with groups")
	// }

	// return nil

	// return groups, nil
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
		err := UpdateProgress(participant.ID, participant.Rank*100, 0)
		if err != nil {
			return nil, errors.New("cannot update user progress")
		}
	}

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Rank > participants[j].Rank
	})

	top3winnerIDs := []primitive.ObjectID{participants[0].ID, participants[1].ID, participants[2].ID}

	for i, reward := range []int{5000, 3000, 2000, 1000, 1000, 1000, 1000, 1000, 1000, 1000} {
		err := UpdateProgress(participants[i].ID, 0, reward)
		if err != nil {
			return nil, errors.New("cannot update user progress")
		}
	}

	return top3winnerIDs, nil
}
