package model

import "time"

type event struct {
	id         string
	controller string
	eventType  string
	payload    string
	createdAt  time.Time
}

type Event interface {
	GetID() string
	GetController() string
	GetEventType() string
	GetPayload() string
	GetCreatedAt() time.Time
}

func (e event) GetID() string {
	return e.id
}

func (e event) GetEventType() string {
	return e.eventType
}

func (e event) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e event) GetPayload() string {
	return e.payload
}

func (e event) GetController() string {
	return e.controller
}

func NewEvent(id string, controller string, eventType string, payload string, createdAt time.Time) Event {
	event := event{
		id:         id,
		controller: controller,
		eventType:  eventType,
		payload:    payload,
		createdAt:  createdAt,
	}

	return event
}
