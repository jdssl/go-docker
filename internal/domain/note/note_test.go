package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	title    = "Title X"
	finished = false
	topics   = []string{"dev", "golang"}
)

func Test_NewNote_CreateNote(t *testing.T) {
	assert := assert.New(t)

	note, _ := NewNote(title, finished, topics)

	assert.Equal(note.Title, title)
	assert.False(note.Finished)
	assert.Equal(len(note.Topics), len(topics))
}
