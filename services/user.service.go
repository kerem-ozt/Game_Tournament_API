package services

import (
	"errors"

	"github.com/kamva/mgm/v3"
	db "github.com/kerem-ozt/GoodBlast_API/models/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser create a user record
func CreateUser(name string, email string, plainPassword string) (*db.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("cannot generate hashed password")
	}

	user := db.NewUser(email, string(password), name, db.RoleUser, db.InitialLevel, db.InitialCoin, db.InitialProgress)
	err = mgm.Coll(user).Create(user)
	if err != nil {
		return nil, errors.New("cannot create new user")
	}

	return user, nil
}

// FindUserById find user by id
func FindUserById(userId primitive.ObjectID) (*db.User, error) {
	user := &db.User{}
	err := mgm.Coll(user).FindByID(userId, user)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return user, nil
}

// FindUserByEmail find user by email
func FindUserByEmail(email string) (*db.User, error) {
	user := &db.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return user, nil
}

// CheckUserMail search user by email, return error if someone uses
func CheckUserMail(email string) error {
	user := &db.User{}
	userCollection := mgm.Coll(user)
	err := userCollection.First(bson.M{"email": email}, user)
	if err == nil {
		return errors.New("email is already in use")
	}

	return nil
}

// Progress update users score
func Progress(userId primitive.ObjectID, score int) error {
	user := &db.User{}
	err := mgm.Coll(user).FindByID(userId, user)
	if err != nil {
		return errors.New("cannot find user")
	}

	user.Progress += score
	err = mgm.Coll(user).Update(user)
	if err != nil {
		return errors.New("cannot update user")
	}

	return nil
}

// Attend to tournament
func Attend(userId primitive.ObjectID, tournamentId primitive.ObjectID) error {
	user := &db.User{}
	err := mgm.Coll(user).FindByID(userId, user)
	if err != nil {
		return errors.New("cannot find user")
	}

	tournament := &db.Tournament{}
	err = mgm.Coll(tournament).FindByID(tournamentId, tournament)
	if err != nil {
		return errors.New("cannot find tournament")
	}

	tournament.Participants = append(tournament.Participants, userId)
	err = mgm.Coll(tournament).Update(tournament)
	if err != nil {
		return errors.New("cannot update tournament")
	}

	return nil
}
