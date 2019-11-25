package repository

import (
	"log"
	"time"

	model "github.com/mg/microgardener/model"
	persistence "github.com/mg/microgardener/persistence"
)

type EventRepository interface {
	Save(model.Event)
	GetLastEventByType(eventType string) model.Event
}

type SqlEventRepository struct{}

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

func (repository SqlEventRepository) GetLastEventByType(eventType string) model.Event {
	db := persistence.Init()
	defer db.Close()

	sqlStmt := `
		SELECT * FROM public.events WHERE event_type = $1 ORDER BY created_at DESC LIMIT 1
	`

	rows := db.QueryRow(sqlStmt, eventType)

	type event struct {
		id         string
		eventType  string
		controller string
		payload    string
		createdAt  time.Time
	}

	var e event
	err := rows.Scan(&e.id, &e.eventType, &e.controller, &e.payload, &e.createdAt)
	if err != nil {
		log.Print(err)
		return nil
	}

	return model.NewEvent(e.id, e.controller, e.eventType, e.payload, e.createdAt)
}
