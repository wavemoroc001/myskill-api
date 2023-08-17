package main

import "myskill-api/db"

func main() {
	db, cleanupDBConnFunc := db.NewDBConn()
	defer cleanupDBConnFunc()
	if db != nil {
		println("connect success")
	}
}
