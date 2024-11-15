package merit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
)

func New() (PlatformClient, error) {
	platformAPIRoot := os.Getenv("MERIT_API_URL")
	return &meritPlatform{platformRoot: platformAPIRoot}, nil
}

var _ PlatformClient = (*meritPlatform)(nil)

type meritPlatform struct {
	platformRoot string
}

func (s *meritPlatform) GetContainers(ctx context.Context, entityID uuid.UUID, token string) (*ContainersResponse, error) {
	url := fmt.Sprintf("%s/containers?recipient_id=%s", s.platformRoot, entityID.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var containers ContainersResponse
	if err := json.NewDecoder(resp.Body).Decode(&containers); err != nil {
		return nil, err
	}
	return &containers, nil
}
