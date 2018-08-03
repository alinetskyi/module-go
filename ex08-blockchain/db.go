package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

func checkFirstBlock(database *sql.DB) bool {
	rows, _ := database.Query("SELECT data FROM block WHERE id = 1")
	var data string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data)
		return true
	}
	return false
}

func createDatabase(database *sql.DB) *sql.Stmt {
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS block (id INTEGER PRIMARY      KEY, timestamp INTEGER, data BLOB, hash TEXT    , prevHash TEXT)")
	statement.Exec()
	if !checkFirstBlock(database) {
		block := NewGenesisBlock()
		createBlock(block, database, statement)
	}
	return statement
}

func createBlock(block *Block, database *sql.DB, statement *sql.Stmt) *sql.Stmt {
	statement, _ = database.Prepare("INSERT INTO block (timestamp, data, hash, prevHash)        VALUES (?, ?, ?, ?)")
	statement.Exec(block.Timestamp, block.Data, block.Hash, block.PrevBlockHash)
	return statement
}

func printList(database *sql.DB) {
	rows, _ := database.Query("SELECT id, data, hash, prevHash FROM block")
	var id int
	var data string
	var hash []byte
	var prevHash []byte
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &data, &hash, &prevHash)
		fmt.Println("Block's id is: ", strconv.Itoa(id))
		fmt.Println("Block's data is: ", data)
		fmt.Printf("hash: %x\n", hash)
		fmt.Println()
	}
}
