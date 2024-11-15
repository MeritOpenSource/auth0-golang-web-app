package login

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"01-Login/platform/authenticator"
)

// Handler for our login.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the state inside the session.
		session := sessions.Default(ctx)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Note: The "audience" parameter is used to specify the API service that the client is
		// interacting with
		url := auth.AuthCodeURL(state,
			oauth2.SetAuthURLParam("audience", os.Getenv("MERIT_AUDIENCE")),
		)
		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
