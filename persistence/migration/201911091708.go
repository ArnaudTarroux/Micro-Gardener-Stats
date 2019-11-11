package migration

import (
	"database/sql"
	"fmt"
)

func migration201911091708(connection *sql.DB) {
	fmt.Println("We are starting the migration 201911091708")
	defer connection.Close()

	var stmt *sql.Stmt
	var err error

	stmt, err = connection.Prepare(`CREATE TABLE IF NOT EXISTS public.weather (
		id uuid NOT NULL,
		controller VARCHAR(100) NOT NULL,
		humidity NUMERIC (5, 2) NOT NULL,
		temperature NUMERIC (5, 2) NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		PRIMARY KEY(id)
	);`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	stmt, err = connection.Prepare("CREATE INDEX IF NOT EXISTS controller ON public.weather (controller);")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}
