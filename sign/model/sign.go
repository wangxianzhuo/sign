package sign

import "time"

type Sign struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	ReferenceID string    `json:"reference_id"`
	Tag         string    `json:"tag"`
	CreatedTime time.Time `json:"created_time"`
}
