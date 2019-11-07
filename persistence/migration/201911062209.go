package migration

import (
	"database/sql"
	"fmt"
)

func migration201911062209(connection *sql.DB) {
	fmt.Println("We are starting the migration 201911062209")
	defer connection.Close()

	var stmt *sql.Stmt
	var err error

	stmt, err = connection.Prepare(`CREATE TABLE IF NOT EXISTS public.events (
		id uuid NOT NULL,
		controller VARCHAR(100) NOT NULL,
		event_type VARCHAR(100) NOT NULL,
		payload json NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		PRIMARY KEY(id)
	);`)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	stmt, err = connection.Prepare("CREATE INDEX IF NOT EXISTS controller ON public.events (controller);")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	stmt, err = connection.Prepare("CREATE INDEX IF NOT EXISTS event_type ON public.events (event_type);")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	stmt, err = connection.Prepare("CREATE INDEX IF NOT EXISTS created_at ON public.events (created_at);")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}
