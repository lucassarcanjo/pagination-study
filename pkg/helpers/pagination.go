package helpers

import (
	"encoding/base64"
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

type PaginationInfo struct {
	Cursor string `json:"cursor"`
	Total  int64  `json:"total"`
}

type Cursor struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func GeneratePager(nextId uuid.UUID, createdAt time.Time, total int64) PaginationInfo {
	cursor := Cursor{
		ID:        nextId,
		CreatedAt: createdAt,
	}

	encodedCursor := encodeCursor(cursor)

	return PaginationInfo{
		Cursor: encodedCursor,
		Total:  total,
	}
}

func encodeCursor(cursor Cursor) string {
	bytes, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}

	encodedCursor := base64.URLEncoding.EncodeToString(bytes)
	return encodedCursor
}

func DecodeCursor(encodedString string) (*Cursor, error) {
	bytes, err := base64.URLEncoding.DecodeString(encodedString)
	if err != nil {
		return nil, err
	}

	var cursor Cursor
	if err := json.Unmarshal(bytes, &cursor); err != nil {
		return nil, err
	}

	return &cursor, nil
}
