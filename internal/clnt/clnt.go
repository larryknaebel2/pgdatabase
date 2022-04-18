package clnt

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

//github.com/jackc/pgx/v4/pgxpool

type Supplies struct {
	Id           int64
	Name         string
	Description  string
	Manufacturer string
	Color        string
	Inventory    int64
}

var dbPool *pgxpool.Pool

func init() {
	var databaseUrl string = "postgres://postgres:welcome1@localhost:5433/school"
	var errI error
	dbPool, errI = pgxpool.Connect(context.Background(), databaseUrl)
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if errI != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", errI)
		os.Exit(1)
	}
}

func GetSupplyById(id int64) (Supplies, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// postgres://postgres:welcome1@localhost:5433/school
	//set DATABASE_URL=postgres://postgres:welcome1@localhost:5433/school

	//dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	//defer conn.Close(context.Background())
	//defer dbPool.Close()
	var s Supplies
	//var name string
	//var inventory int64
	//err = conn.QueryRow(context.Background(), "select name, inventory from supplies where id=$1", 1).Scan(&name, &inventory)

	err := dbPool.QueryRow(context.Background(),
		`select id, name, description, manufacturer, color, inventory from supplies where id=$1`,
		id).Scan(&s.Id,
		&s.Name,
		&s.Description,
		&s.Manufacturer,
		&s.Color,
		&s.Inventory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(s.Name,
		s.Description,
		s.Inventory)
	return s, nil
}

func GetAllSupplies() ([]Supplies, error) {
	// dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	// //conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer dbPool.Close()
	rows, err := dbPool.Query(context.Background(),
		`select * from supplies`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	var allsupplies []Supplies
	var numitems int64 = 0
	for rows.Next() {
		var s Supplies
		if err2 := rows.Scan(&s.Id,
			&s.Name,
			&s.Description,
			&s.Manufacturer,
			&s.Color,
			&s.Inventory); err2 != nil {
			return nil, err2
		}
		allsupplies = append(allsupplies, s)
		numitems += 1

	}
	//fmt.Printf("%v Count: %d", allsupplies, numitems)
	return allsupplies, nil

}
