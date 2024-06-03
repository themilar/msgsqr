package models

import (
	"database/sql"
	"errors"
	"time"
)

type Message struct {
	ID        string
	Title     string
	Content   string
	Created   time.Time
	CreatedAt time.Time
}

type MessageModel struct {
	DB *sql.DB
}

var ErrNoRecord error = errors.New("models: no matching record found")

func (m *MessageModel) Insert(title, content string) (int, error) {
	statement := `INSERT INTO message (title, content) VALUES ($1, $2) returning id;`
	_, err := m.DB.Exec(statement, title, content)
	if err != nil {
		return 0, err
	}
	var id int
	err = m.DB.QueryRow(statement, title, content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MessageModel) Get(id int) (*Message, error) {
	statement := `SELECT id,title,content,created FROM message WHERE id=$1`
	mess := &Message{}
	err := m.DB.QueryRow(statement, id).Scan(&mess.ID, &mess.Title, &mess.Content, &mess.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}

	}
	return mess, nil
}
func (m *MessageModel) Latest() ([]*Message, error) {
	statement := `SELECT id,title,content,created,created_at FROM message ORDER BY id DESC LIMIT 10`
	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	messages := []*Message{}
	for rows.Next() {
		mess := &Message{}
		err := rows.Scan(&mess.ID, &mess.Title, &mess.Content, &mess.Created, &mess.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, mess)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
