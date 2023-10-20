package note

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	title  = "Title X"
	topics = []string{"dev", "golang"}
	fake   = faker.New()
)

func Test_NewNote_CreateNote(t *testing.T) {
	now := time.Now().Add(-time.Minute)

	note, _ := NewNote(title, topics)

	assert.Equal(t, note.Title, title)
	assert.False(t, note.Finished)
	assert.Equal(t, len(note.Topics), len(topics))
	assert.Greater(t, note.CreatedOn, now)
}

func Test_NewNote_IDSNotNill(t *testing.T) {
	note, _ := NewNote(title, topics)

	assert.NotNil(t, note.ID)
}

func Test_NewNote_MustValidateTitleMin(t *testing.T) {
	_, err := NewNote("", topics)

	assert.Equal(t, "title is required with min 5", err.Error())
}

func Test_NewNote_MustValidateTitleMax(t *testing.T) {
	_, err := NewNote(fake.Lorem().Text(181), topics)

	assert.Equal(t, "title is required with max 180", err.Error())
}

func Test_NewNote_MustValidateTopicsMin(t *testing.T) {
	_, err := NewNote(title, nil)

	assert.Equal(t, "topics is required with min 1", err.Error())
}
