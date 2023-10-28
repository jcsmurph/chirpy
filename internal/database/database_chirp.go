package database

import "errors"

type Chirp struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
}

func (db *DB) CreateChirp(body string, authorID int) (Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	id := len(dbStructure.Chirps) + 1
	chirp := Chirp{
		ID:       id,
		Body:     body,
		AuthorID: authorID,
	}
	dbStructure.Chirps[id] = chirp

	err = db.writeDB(dbStructure)
	if err != nil {
		return Chirp{}, err
	}

	return chirp, nil
}

func (db *DB) GetChirps() ([]Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	chirps := make([]Chirp, 0, len(dbStructure.Chirps))
	for _, chirp := range dbStructure.Chirps {
		chirps = append(chirps, chirp)
	}

	return chirps, nil
}

func (db *DB) GetChirpID(id int) (Chirp, error) {
	dbStructure, dbErr := db.loadDB()
	if dbErr != nil {
		return Chirp{}, nil
	}

	chirp, err := dbStructure.Chirps[id]

	if err == false {
		return Chirp{}, errors.New("Chirp does not exist")
	}

	return chirp, nil

}

func (db *DB) DeleteChirp(chirpID, authorID int) error {

	dbStructure, dbErr := db.loadDB()
	if dbErr != nil {
		return dbErr
	}

	for id, chirp := range dbStructure.Chirps {
		if chirpID == chirp.ID {
			if authorID == chirp.AuthorID {
                delete(dbStructure.Chirps, id)
                return nil
			} else {
                return errors.New("User is not the Author of the chirp")
			}
		}
	}
	return errors.New("Chirp does not exist")
}
