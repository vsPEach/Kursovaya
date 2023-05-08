package storage

import "errors"

var (
	ErrEventDoesNotExist     = errors.New("event doesn't exists")
	ErrEventAlreadyExist     = errors.New("event already exist")
	ErrNothingEventsForDay   = errors.New("no one events for this day")
	ErrNothingEventsForWeek  = errors.New("no one events for this week")
	ErrNothingEventsForMonth = errors.New("no one events for this month")
)
