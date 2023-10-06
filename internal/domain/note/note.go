package note

import (
	internalerrors "go-docker/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Topic struct {
	Name string
}

type Note struct {
	ID        string    `validate:"required"`
	Title     string    `validate:"min=5,max=180"`
	Finished  bool      `validate:"boolean,isdefault=false"`
	Topics    []Topic   `validate:"min=1,dive"`
	CreatedOn time.Time `validate:"required"`
}

func NewNote(title string, names []string) (*Note, error) {
	topics := make([]Topic, len(names))
	for i, name := range names {
		topics[i].Name = name
	}

	note := &Note{
		ID:        xid.New().String(),
		Title:     title,
		Finished:  false,
		Topics:    topics,
		CreatedOn: time.Now(),
	}

	err := internalerrors.ValidateStruct(note)
	if err == nil {
		return note, nil
	}

	return nil, err
}
