package repository

import (
	"log"

	model "github.com/mg/microgardener/model"
	persistence "github.com/mg/microgardener/persistence"
)

type EventRepository interface {
	Save(model.Event)
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
