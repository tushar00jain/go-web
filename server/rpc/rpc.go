package rpc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	// "github.com/tushar00jain/go-web/server/schemas"

	"golang.org/x/net/context"

	pb "github.com/tushar00jain/service/protos"
)

type AddressBookServer struct {
	Db *sql.DB
}

func (ab *AddressBookServer) GetPersons (ctx context.Context, e *pb.Empty) (*pb.Persons, error) {
	query := "SELECT p.\"Id\", p.\"Name\", p.\"Email\"," +
	"json_agg(json_build_object('phoneNumber', ph.\"Number\", 'phoneType', ph.\"Type\")) AS \"PhoneNumbers\"" +
	"FROM Person p INNER JOIN PhoneNumber ph ON p.\"Id\"=ph.\"PersonId\" GROUP BY p.\"Id\";"
	rows, err := ab.Db.Query(query)

	if err != nil {
		fmt.Println("selct error")
		return nil, err
	}

	var persons pb.Persons
	persons.People = make([]*pb.Person, 0)
	for rows.Next() {
		person := &pb.Person{}
		var temp string
		err = rows.Scan(&person.Id, &person.Name, &person.Email, &temp)

		if err != nil {
			fmt.Println("query error")
			return nil, err
		}

		tbs := []byte(temp)
		if err := json.Unmarshal(tbs, &person.PhoneNumbers); err != nil {
			fmt.Println("unmarshal error")
			return nil, err
		}

		persons.People = append(persons.People, person)
	}
	return &persons, nil
}

func (ab *AddressBookServer) GetAddressBook (ctx context.Context, i *pb.Id) (*pb.Book, error) {
	var query = "SELECT d.*, json_agg(json_build_object('phoneNumber', ph.\"Number\", 'phoneType', ph.\"Type\")) AS \"PhoneNumbers\" FROM (" +
	"SELECT p.*" +
	"FROM Person p INNER JOIN AddressBook a ON p.\"Id\"=a.\"People\" AND a.\"Self\"=" + "1" +
	") d INNER JOIN PhoneNumber ph ON d.\"Id\"=ph.\"PersonId\" GROUP BY d.\"Id\", d.\"Name\", d.\"Email\";"

	rows, err := ab.Db.Query(query)
	if err != nil {
		fmt.Println("selct error")
		return nil, err
	}

	var persons pb.Book
	persons.People = make([]*pb.Person, 0)
	for rows.Next() {
		person := &pb.Person{}
		var temp string

		err = rows.Scan(&person.Id, &person.Name, &person.Email, &temp)
		if err != nil {
			fmt.Println("query error")
			return nil, err
		}

		tbs := []byte(temp)
		if err := json.Unmarshal(tbs, &person.PhoneNumbers); err != nil {
			fmt.Println("unmarshal error")
			return nil, err
		}

		persons.People = append(persons.People, person)
	}

	return &persons, nil
}
