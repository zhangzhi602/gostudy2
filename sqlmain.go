package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func queryInfo( id string)  (string,error){

	db, errOpen := sql.Open("mysql", "name:password@tcp(127.0.0.1:3306)/world?charset=utf8")
	if errOpen != nil {
		return "",errors.Wrap(errOpen," Open  DB is error")
	}
	defer db.Close()
	var im_id int
	var im_address string

	errTables := db.QueryRow("SELECT im_id ,im_address FROM ts_ipmanager WHERE im_id=?", id).Scan(&im_id)
	if errTables != nil {
		des := "not found im_id = " + id
		errors.Wrap(errTables,des)
		return "",errors.Wrap(errTables,des)
	}
	return im_address,nil
}
