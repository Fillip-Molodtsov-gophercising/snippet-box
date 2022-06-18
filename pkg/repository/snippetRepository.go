package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Fillip-Molodtsov-gophercising/snippet-box/pkg/model"
)

type SnippetGetter interface {
	Get(id int) (*model.Snippet, error)
	Latest() ([]*model.Snippet, error)
}

type SnippetUpdater interface {
	Insert(title, content string, expires int) (int, error)
}

type SnippetGetterUpdater interface {
	SnippetGetter
	SnippetUpdater
}

type PostgresSnippetRepository struct {
	db *sql.DB
}

func MakePostgresSnippetRepository(db *sql.DB) PostgresSnippetRepository {
	res := PostgresSnippetRepository{}
	res.SetDB(db)
	return res
}

func (m *PostgresSnippetRepository) SetDB(db *sql.DB) {
	m.db = db
}

func (m *PostgresSnippetRepository) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
			VALUES($1, $2, now(), now() + cast($3 as interval)) returning id`
	row := m.db.QueryRow(stmt, title, content, fmt.Sprintf("%d days", expires))
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *PostgresSnippetRepository) Get(id int) (*model.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
								WHERE expires > now() AND id = $1`

	s := &model.Snippet{}
	err := m.db.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *PostgresSnippetRepository) Latest() ([]*model.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
                                            WHERE expires > now() ORDER BY created DESC LIMIT 10`

	rows, err := m.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*model.Snippet{}

	for rows.Next() {

		s := &model.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
