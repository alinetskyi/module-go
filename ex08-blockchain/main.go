package main

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./blockchain.db")
	statement := createDatabase(database)
	//addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	//listBlockCmd := flag.NewFlagSet("list", flag.ExitOnError)
	chain := NewBlockchain()
	cmd := HandleArgs(chain)
	switch cmd[0] {
	case "add":
		//chain.StageBlock(cmd[1])

		if cmd[2] == "" {
			panic("[ERROR] You can't mine with difficulty 0")
		}
		i, err := strconv.Atoi(cmd[2])
		if err == nil && i < 10 {
			//_ = addBlockCmd.Parse(cmd)
			block := chain.AddBlock(cmd[1], i)
			statement = createBlock(block, database, statement)
		} else {
			panic("[ERROR] You need to specify difficulty and it has to be more than 24")
		}
	case "list":
		//err := listBlockCmd.Parse(cmd[1:])
		//if err != nil {
		//	log.Panic(err)
		//}
		printList(database)
		//b.ListBlocks()
		//case "mine":
	}

}
func HandleArgs(b *Blockchain) []string {
	if len(os.Args) < 2 {
		panic("[ERROR]: Sorry, nothing to do...!")
	}
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 4 {
		panic("[EROR]: Sorry, way tooo many arguments")
	}
	return argsWithoutProg
}

/*
func RunCommand(b *Blockchain, cmd []string, database *sql.DB) {
	switch cmd[0] {
	case "add":
		b.StageBlock(cmd[1])
	case "list":
		err := listBlockCmd.Parse(cmd[1:])
		if err != nil {
			log.Panic(err)
		}
		printList(database)
		//b.ListBlocks()
	case "mine":
		if cmd[1] == "" {
			panic("[ERROR] You can't mine with difficulty 0")
		}
		i, err := strconv.Atoi(cmd[1])
		if err == nil && i < 10 {
			err := addBlockCmd.Parse(i)
			b.AddBlock(i)
			statement = createBlock(block, database, statement)
		} else {
			panic("[ERROR] You need to specify difficulty and it has to be more than 24")
		}
	}
}*/
