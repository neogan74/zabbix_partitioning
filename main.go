/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"neogan74/zabbix_partitioning/cmd"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	host     string
	user     string
	password string
	db       string
}

type City struct {
	Id         int
	Name       string
	Population int
}

func main() {
	test()
}
func test() {

	cmd.Execute()
	db, err := sql.Open("mysql", "root@tcp(zbx:3306)/testdb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)

	res, err := db.Query("SELECT * FROM cities")

	defer res.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for res.Next() {
		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%v\n", city)
	}

	sql := "INSERT INTO cities(name, population) VALUES ('Moscow',12306000)"
	res2, err := db.Exec(sql)
	if err != nil {
		log.Fatalln(err)
	}

	lastId, err := res2.LastInsertId()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("The last inserted row id: %d\n", lastId)

	var someID int = 12

	res3, err := db.Query("SELECT * FROM cities WHERE Id = ?", someID)
	defer res3.Close()

	if err != nil {
		log.Fatalln(err)
	}

	if res3.Next() {
		var city City
		err := res3.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%v\n", city)
	} else {
		fmt.Println("No city found.")
	}

}
