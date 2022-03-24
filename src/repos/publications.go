package repos

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func NewPublicationsRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(post models.Publications) (uint64, error) {
	statement, erro := repository.db.Prepare("insert into posts (title, content, author_id) value (?, ?, ?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(post.Title, post.Content, post.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastInsertID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertID), nil
}
