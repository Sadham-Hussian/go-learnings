package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func createDatabase(conn driver.Conn) error {
	// database name : golang
	err := conn.Exec(context.Background(), `
    CREATE DATABASE IF NOT EXISTS golang
	`)
	return err
}

func dropTable(conn driver.Conn) error {
	// table name : test
	err := conn.Exec(context.Background(), `DROP TABLE IF EXISTS test`)
	return err
}

func createTable(conn driver.Conn) error {
	// table name : test
	err := conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS golang.test(
		x String,
		y String
	)ENGINE = Memory`)
	return err
}

func insertData(conn driver.Conn) error {
	// table name : test
	err := conn.Exec(context.Background(), `INSERT INTO golang.test (x,y) VALUES ('sadham', 'hussian')`)
	return err
}

func selectData(conn driver.Conn) error {
	// select Data
	result, err := conn.Query(context.Background(), `SELECT * FROM golang.test`)
	if err != nil {
		return err
	}
	var (
		x string
		y string
	)
	for result.Next() {
		result.Scan(&x, &y)
		log.Println("x = ", x, " y = ", y)
	}

	return nil
}

func main() {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", "127.0.0.1", 9000)},
		Auth: clickhouse.Auth{
			Database: "test",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		log.Println(err)
	}
	v, err := conn.ServerVersion()
	fmt.Println(v)
	if err != nil {
		log.Println(err)
	}

	// *Creating Database
	err = createDatabase(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println("Database golang created!!!")

	err = dropTable(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println("Table test droped!!!")

	// *Creating Table
	err = createTable(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println("Table test created!!!")

	// *Insert data into Table
	err = insertData(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println("Data inserted!!!")

	// *Select data from Table
	err = selectData(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println("Data selected!!!")
}
