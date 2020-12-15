package models

import (
	"errors"
	"strings"
	"time"
)

// Publication struct
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Prepare publication request body
func (publication *Publication) Prepare() error {
	if error := publication.validate(); error != nil {
		return error
	}

	publication.format()

	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("Title is required")
	}

	if publication.Content == "" {
		return errors.New("Content is required")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
