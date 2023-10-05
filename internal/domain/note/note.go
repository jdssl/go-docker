package note

import (
	"time"

	"github.com/rs/xid"
)

type Topic struct {
	Name string
}

type Note struct {
	ID        string
	Title     string
	Finished  bool
	Topics    []Topic
	CreatedOn time.Time
}

func NewNote(title string, finished bool, names []string) (*Note, error) {
	topics := make([]Topic, len(names))
	for i, name := range names {
		topics[i].Name = name
	}

	note := &Note{
		ID:        xid.New().String(),
		Title:     title,
		Finished:  finished,
		Topics:    topics,
		CreatedOn: time.Now(),
	}

	return note, nil
}
