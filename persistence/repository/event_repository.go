package repository

import (
	"fmt"
	"log"

	model "github.com/mg/microgardener/model"
	persistence "github.com/mg/microgardener/persistence"
)

type EventRepository interface {
	Save(model.Event)
	GetLastEventByType(eventType string)
}

type SqlEventRepository struct {
}

func (repository SqlEventRepository) Save(event model.Event) {
	db := persistence.Init()
	defer db.Close()

	sqlStmt := `
		INSERT INTO public.events (id, controller, event_type, payload, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	result, err := db.Exec(
		sqlStmt,
		event.GetID(),
		event.GetController(),
		event.GetEventType(),
		event.GetPayload(),
		event.GetCreatedAt(),
	)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	log.Printf("New event created %d", rowsAffected)
}

func (repository SqlEventRepository) GetLastEventByType(eventType string) {
	db := persistence.Init()
	defer db.Close()

	sqlStmt := `
		SELECT payload, created_at FROM public.events WHERE event_type = $1 ORDER BY created_at DESC LIMIT 1
	`

	rows, err := db.Query(sqlStmt, eventType)
	if err != nil {
		panic(err)
	}

	var payload string
	var createdAt string

	rows.Next()
	err = rows.Scan(&payload, &createdAt)
	if err != nil {
		panic(err)
	}
	fmt.Println(payload, createdAt)
}
