package merit

import (
	"context"
	"github.com/google/uuid"
)

// PlatformClient is based on the [Merit] platform API.
// For this sample, we only are requesting [Containers] belonging to the
// user's entity ID.
//
// [Merit]: https://merit-developer-portal.readme.io/
// [Containers]: https://merit-developer-portal.readme.io/reference/getcontainers
type PlatformClient interface {
	GetContainers(ctx context.Context, entityID uuid.UUID, token string) (*ContainersResponse, error)
}

type ContainersResponse struct {
	Cursor     PaginatedCursor `json:"cursor"`
	HasMore    bool            `json:"hasMore"`
	Containers []*Container    `json:"containers"`
}

type FieldKind struct {
	FieldType string    `json:"fieldType"`
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
}

type PaginatedCursor struct {
	Current Cursor `json:"current"`
}

type Cursor struct {
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	StartingAfter string `json:"starting_after"`
}
