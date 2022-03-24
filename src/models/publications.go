package models

import (
	"errors"
	"strings"
	"time"
)

// Posts has all the fields on a post post.
type Posts struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick uint64    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedIn  time.Time `json:"createdIn,omitempty"`
}

func (post *Posts) Prepare() error {
	if erro := post.validate(); erro != nil {
		return erro
	}

	post.format()
	return nil
}

func (post *Posts) validate() error {
	if post.Title == "" {
		return errors.New("Field 'Title' is required.")
	}

	if post.Content == "" {
		return errors.New("Field 'Content' is required.")
	}

	return nil
}

func (post *Posts) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
