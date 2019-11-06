package persistence

import (
	"fmt"
)

var migrations = map[int]func(){
	201911062209: migration201911062209,
}

func MigrateDb(version int) {
	fmt.Println("Run migration")
	if version == 0 {
		panic("Please provide a migration version in command argument")
	}

	fmt.Println(version)
	if nil == migrations[version] {
		panic("The migration does not exist")
	}

	migrations[version]()
}
