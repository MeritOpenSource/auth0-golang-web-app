package merit

import "github.com/google/uuid"

type Entity struct {
	Id    uuid.UUID `json:"id"`
	State string    `json:"state"`
	Type  string    `json:"type"`
}
