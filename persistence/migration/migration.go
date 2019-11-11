package migration

import (
	"database/sql"
	"fmt"

	db "github.com/mg/microgardener/persistence"
)

var migrations = map[int]func(*sql.DB){
	201911062209: migration201911062209,
	201911091708: migration201911091708,
}

func MigrateDb(version int) {
	fmt.Println("Run migration")
	if version == 0 {
		panic("Please provide a migration version in command argument")
	}

	if nil == migrations[version] {
		panic("The migration does not exist")
	}

	if fn := migrations[version]; fn != nil {
		connection := db.Init()
		fn(connection)
	}
}
