package database

import (
	"errors"
	"time"
)

func (db *DB) RevokeToken(token string) error {
	dbStructure, err := db.loadDB()

	if err != nil {
		return errors.New("Unable to load the database")
	}

	dbStructure.Token[token] = time.Now()

	err = db.writeDB(dbStructure)
	if err != nil {
		return err
	}
	return nil

}

func (db *DB) CheckRevokeToken(token string) error {
	dbStructure, err := db.loadDB()

	if err != nil {
		return errors.New("Unable to load the database")
	}

    _, ok := dbStructure.Token[token]

    if ok {
        return errors.New("Token has been revoked")
    }

    return nil
}
