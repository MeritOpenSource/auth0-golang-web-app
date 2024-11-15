package merit

import (
	"github.com/google/uuid"
	"time"
)

type Container struct {
	Id                   uuid.UUID           `json:"id"`
	AcceptedAt           *time.Time          `json:"acceptedAt"`
	Active               bool                `json:"active"`
	ActivenessFailures   []*ValidationResult `json:"activenessFailures"`
	AuthorizedAt         *time.Time          `json:"authorizedAt"`
	Completed            bool                `json:"completed"`
	CompletenessFailures []*ValidationResult `json:"completenessFailures"`
	CreatedAt            time.Time           `json:"createdAt"`
	Description          string              `json:"description"`
	Fields               []ContainerField    `json:"fields"`
	Issuer               Entity              `json:"issuer"`
	Name                 string              `json:"name"`
	Recipient            *Entity             `json:"recipient"`
	RejectedAt           *time.Time          `json:"rejectedAt"`
	RevokedAt            *time.Time          `json:"revokedAt"`
	State                ContainerState      `json:"state"`
	TemplateId           uuid.UUID           `json:"templateId"`
	UpdatedAt            time.Time           `json:"updatedAt"`
}

type ContainerState struct {
	Name       string    `json:"name"`
	OccurredAt time.Time `json:"occurredAt"`
}

type ContainerField struct {
	CanShare             bool        `json:"canShare"`
	Description          string      `json:"description"`
	FieldKind            FieldKind   `json:"fieldKind"`
	Metadata             *string     `json:"metadata"`
	Name                 string      `json:"name"`
	TemplateFieldID      uuid.UUID   `json:"templateFieldID"`
	TemplateFieldLineage []uuid.UUID `json:"templateFieldLineage"`
	UpdatedAt            time.Time   `json:"updatedAt"`
	ValidationErrors     any         `json:"validationErrors"`
	Value                string      `json:"value,omitempty"`
}
