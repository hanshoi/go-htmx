package db

import (
	"database/sql"
	"log"
)

type List struct {
	Id       uint32
	Name     string
	Filename string
}

type ListRepo struct {
	db *sql.DB
}

func (repo *ListRepo) GetLists() []List {
	var lists []List
	rows, err := repo.db.Query("SELECT * FROM lists")
	if err != nil {
		log.Fatalln(err)
		return lists
	}
	defer rows.Close()

	for rows.Next() {
		var list List
		if err := rows.Scan(&list.Id, &list.Filename, &list.Name); err != nil {
			log.Fatalln(err)
			return lists
		}
		lists = append(lists, list)
	}
	return lists
}
