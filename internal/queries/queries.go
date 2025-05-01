package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronGulman/chat-gin/internal/chat"
	"log"
	"time"
)

type Queries struct {
	DB *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{DB: db}
}

const queryTimeout = 5 * time.Second

func (q *Queries) GetUser(userID int) (*chat.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	query := `SELECT id, name, password, email FROM users WHERE id = ?`

	var user chat.User
	err := q.DB.QueryRowContext(ctx, query, userID).Scan(
		&user.Id,
		&user.Name,
		&user.Password,
		&user.Email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user found with id %d", userID)
		}
		return nil, fmt.Errorf("query failed: %w", err)
	}

	if user.Email == "" {
		user.Email = "none"
	}

	log.Printf("[INFO] Retrieved user: %+v\n", user)
	return &user, nil
}

func (q *Queries) AddUser(name, password, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	query := `INSERT INTO users (name, password, email) VALUES (?, ?, ?)`

	result, err := q.DB.ExecContext(ctx, query, name, password, email)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("[INFO] Inserted user '%s' (affected rows: %d)", name, rowsAffected)
	return nil
}

func (q *Queries) SaveMsg(content string) error {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	query := `INSERT INTO messages (content) VALUES (?)`

	_, err := q.DB.ExecContext(ctx, query, content)
	if err != nil {
		return fmt.Errorf("failed to save message: %w", err)
	}

	log.Printf("[INFO] Saved message: %s", content)
	return nil
}

func (q *Queries) AllMsg() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	query := `SELECT content FROM messages`

	rows, err := q.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", err)
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var msg string
		if err := rows.Scan(&msg); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	log.Printf("[INFO] Retrieved %d messages", len(messages))
	return messages, nil
}
