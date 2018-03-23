package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	//dataSourceName = "root:RussellCloud2017@test.dl.russellcloud.com:3306/russell"

	// OjbK ^_^
	dataSourceName = "root:RussellCloud2017@tcp(139.224.114.10:3306)/Russell"
	database *sql.DB
)

func init()  {
	err := error(nil)
	database, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	defer database.Close()
	getTaskSQL := "show tables;"
	rows, err := database.Query(getTaskSQL)
	if (err != nil) {
		fmt.Println(err)
	} else {
		//fmt.Println(rows)
	}
	defer rows.Close()
	for rows.Next(){
		tablename := ""
		err := rows.Scan(&tablename)
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println(tablename)
		}
	}

	var tablename string;
	stmt, err := database.Prepare("SELECT id from user where username=?")
	err = stmt.QueryRow("Danceiny").Scan(&tablename)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(tablename)
	}
}
