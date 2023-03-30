package repository

import (
	"database/sql"

	"formative-15/structs"
)

func GetAllPerson(db *sql.DB) (error, []structs.Person) {
	sql := "SELECT * FROM person"

	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	results := []structs.Person{}

	for rows.Next() {
		var person = structs.Person{}

		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return err, nil
		}

		results = append(results, person)
	}

	return nil, results
}

func InsertPerson(db *sql.DB, person structs.Person) error {
	sql := "INSERT INTO person (id, first_name, last_name) VALUES ($1, $2, $3)"

	_, err := db.Exec(sql, person.ID, person.FirstName, person.LastName)

	return err
}

func UpdatePerson(db *sql.DB, person structs.Person) error {
	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"

	_, err := db.Exec(sql, person.FirstName, person.LastName, person.ID)

	return err
}

func DeletePerson(db *sql.DB, person structs.Person) error {
	sql := "DELETE FROM person WHERE ID = $1"

	_, err := db.Exec(sql, person.ID)

	return err
}
