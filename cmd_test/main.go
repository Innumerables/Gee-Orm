package main

import (
	"fmt"
	geeorm "gee-orm"

	_ "github.com/go-sql-driver/mysql"
)

// func main() {
// 	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/demo?charset=utf8mb4&parseTime=True&loc=Local")
// 	defer func() { _ = db.Close() }()
// 	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
// 	_, _ = db.Exec("CREATE TABLE User(Name text);")
// 	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
// 	if err == nil {
// 		affected, _ := result.RowsAffected()
// 		log.Println(affected)
// 	}
// 	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
// 	var name string
// 	if err := row.Scan(&name); err == nil {
// 		log.Println(name)
// 	}
// }

func main() {
	engine, _ := geeorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:13306)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
