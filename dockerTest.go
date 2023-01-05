package main

import (
	"fmt"
	"os"
	"transmitter/sql"
)

func main() {
	var sql sql.Database = &sql.Mysql{
		Config: sql.Config{
			Username:  "XXX",
			Password:  "XXX",
			Protocol:  "tcp",
			Host:      "XXX.XXX.XXX.XXX",
			Port:      3306,
			Database:  "XXX",
			Chartset:  "utf8",
			Location:  "Asia%2FTaipei",
			Parsetime: true,
		},
	}
	defer sql.Close()
	err := sql.Conn()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("PASS")
		os.Exit(0)
	}

}
