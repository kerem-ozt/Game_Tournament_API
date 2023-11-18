package services

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"

	"github.com/kamva/mgm/v3"
	db "github.com/kerem-ozt/GoodBlast_API/models/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTournament create new tournament record
func CreateTournament(participants []primitive.ObjectID) (*db.Tournament, error) {
	tournament := db.NewTournament(participants)
	err := mgm.Coll(tournament).Create(tournament)
	if err != nil {
		return nil, errors.New("cannot create new tournament")
	}

	return tournament, nil
}

// GetTournaments get paginated tournaments list
func GetTournaments(page int, limit int) ([]db.Tournament, error) {
	var tournaments []db.Tournament

	err := mgm.Coll(&db.Tournament{}).SimpleFind(&tournaments, nil, nil)

	fmt.Println("tournaments: ", tournaments)

	if err != nil {
		fmt.Println("Error finding tournaments:", err)
		return nil, errors.New("cannot find tournaments")
	}

	fmt.Println("Number of Tournaments:", len(tournaments))
	for _, t := range tournaments {
		fmt.Println("Tournament:", t)
	}

	return tournaments, nil
}

func ProgressTournament(tournamentID primitive.ObjectID) error {
	tournament := &db.Tournament{}

	type Participant struct {
		ID   primitive.ObjectID `bson:"id"`
		Rank int
	}

	err := mgm.Coll(tournament).FindByID(tournamentID, tournament)
	if err != nil {
		return errors.New("cannot find tournament")
	}

	var participants []Participant
	for _, objID := range tournament.Participants {
		id, err := primitive.ObjectIDFromHex(objID.Hex())
		if err != nil {
			return errors.New("invalid participant ID")
		}
		participants = append(participants, Participant{ID: id, Rank: 0})
	}

	fmt.Println(participants)

	for round := 1; len(tournament.Participants) > 1; round++ {
		for i := len(tournament.Participants) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			tournament.Participants[i], tournament.Participants[j] = tournament.Participants[j], tournament.Participants[i]
		}

		winnerCount := len(tournament.Participants) / 2

		winnersSlice := tournament.Participants[:winnerCount]

		fmt.Println("Round", round, "Winners Slice:", winnersSlice)

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

	fmt.Println("Final Winner:", tournament.Participants[0])

	fmt.Println("Final Participants:", participants)

	for _, participant := range participants {
		err = Progress(participant.ID, participant.Rank*100)
		if err != nil {
			return errors.New("cannot update user progress")
		}
	}

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Rank > participants[j].Rank
	})

	fmt.Println("After Sorting:", participants)

	top3winnerIDs := []primitive.ObjectID{participants[0].ID, participants[1].ID, participants[2].ID}

	fmt.Println("After Sorting:", top3winnerIDs)

	return nil
}
