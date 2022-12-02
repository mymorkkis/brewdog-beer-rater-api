package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Could install the migrate CLI and use instead but wanted to attempt a script
// https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

func main() {
	migrationName := getMigrationName()
	wd, err := os.Getwd()
	checkErr(err)

	dirPath := filepath.Join(wd, "db", "migrations")
	nextVersion := getNextVersionNumber(dirPath)

	downFile := fmt.Sprintf("%v_%v.down.sql", nextVersion, migrationName)
	createFile(dirPath, downFile)

	upFile := fmt.Sprintf("%v_%v.up.sql", nextVersion, migrationName)
	createFile(dirPath, upFile)
}

func getMigrationName() string {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		log.Fatal(
			"Invalid number of args, 1 expected. Usage: go run PATH_TO_SCRIPT MIGRATION_NAME",
		)
	}
	return os.Args[1]
}

func getNextVersionNumber(dirPath string) int {
	files, err := ioutil.ReadDir(dirPath)
	checkErr(err)

	currentVersion := int(len(files) / 2)

	if currentVersion%2 != 0 {
		log.Fatal("migration files are out of sync so cannot create new ones")
	}
	return currentVersion + 1
}

func createFile(dirPath, fileName string) {
	filePath := filepath.Join(dirPath, fileName)
	f, err := os.Create(filePath)
	checkErr(err)
	f.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
