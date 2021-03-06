// Code generated by sqlc. DO NOT EDIT.

package ondeck

import (
	"database/sql"
	"fmt"
	"time"
)

type Status string

const (
	StatusOpen   Status = "open"
	StatusClosed Status = "closed"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type City struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Venues are places where muisc happens
type Venue struct {
	ID int64 `json:"id"`
	// Venues can be either open or closed
	Status   Status         `json:"status"`
	Statuses sql.NullString `json:"statuses"`
	// This value appears in public URLs
	Slug            string         `json:"slug"`
	Name            string         `json:"name"`
	City            string         `json:"city"`
	SpotifyPlaylist string         `json:"spotify_playlist"`
	SongkickID      sql.NullString `json:"songkick_id"`
	Tags            sql.NullString `json:"tags"`
	CreatedAt       time.Time      `json:"created_at"`
}
