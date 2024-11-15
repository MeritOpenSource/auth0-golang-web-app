package user

import (
	"01-Login/platform/merit"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our logged-in user page.
func Handler(meritClient merit.PlatformClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		profile := session.Get("profile")

		rawUserID := profile.(map[string]any)["entityID"]
		userID, err := uuid.Parse(rawUserID.(string))
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		payload := new(struct {
			Nickname   string    `json:"nickname"`
			Picture    string    `json:"picture"`
			EntityID   uuid.UUID `json:"entityID"`
			Containers string    `json:"containers"`
		})
		payload.Nickname = profile.(map[string]any)["nickname"].(string)
		payload.Picture = profile.(map[string]any)["picture"].(string)
		payload.EntityID = userID
		at := session.Get("access_token").(string)

		containers, err := meritClient.GetContainers(ctx, userID, at)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		output, err := json.MarshalIndent(containers, "", "  ")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		payload.Containers = string(output)

		ctx.HTML(http.StatusOK, "user.html", payload)
	}
}
