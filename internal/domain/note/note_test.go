package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	title  = "Title X"
	topics = []string{"dev", "golang"}
)

func Test_NewNote_CreateNote(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	note, _ := NewNote(title, topics)

	assert.Equal(note.Title, title)
	assert.False(note.Finished)
	assert.Equal(len(note.Topics), len(topics))
	assert.Greater(note.CreatedOn, now)
}

func Test_NewNote_IDSNotNill(t *testing.T) {
	assert := assert.New(t)

	note, _ := NewNote(title, topics)

	assert.NotNil(note.ID)
}

func Test_NewNote_MustValidateTitleMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewNote("", topics)

	assert.Equal("title is required with min 5", err.Error())
}
