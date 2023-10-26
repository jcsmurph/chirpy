package database

import (
	"errors"
)

var ErrAlreadyExists = errors.New("already exists")

type User struct {
	Email          string `json:"email"`
	ID             int    `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (db *DB) CreateUser(email, hashedPassword string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	id := len(dbStructure.Users) + 1
	user := User{
		ID:    id,
		Email: email,
        HashedPassword: hashedPassword,
	}

	dbStructure.Users[id] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) GetUsers() ([]User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(dbStructure.Chirps))
	for _, user := range dbStructure.Users {
		users = append(users, user)
	}

	return users, nil
}

func (db *DB) GetUserID(id int) (User, error) {
	dbStructure, dbErr := db.loadDB()
	if dbErr != nil {
		return User{}, dbErr
	}

	user, err := dbStructure.Users[id]

	if err == false {
		return User{}, errors.New("User does not exist")
	}

	return user, nil

}

func (db *DB) GetUserByEmail(email string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	for _, user := range dbStructure.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return User{}, ErrNotExist
}
