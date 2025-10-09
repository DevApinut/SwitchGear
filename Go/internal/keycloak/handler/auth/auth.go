package authhandler

import (
	"SwitchGear/internal/keycloak/auth"
	"SwitchGear/internal/keycloak/config"
	store "SwitchGear/internal/keycloak/store/redis"
	"SwitchGear/internal/keycloak/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	cfg          *config.Config
	serverAddr   string
	authClient   *auth.Client
	authStore    store.AuthStore
	sessionStore store.SessionStore
}

func New(
	cfg *config.Config,
	serverAddr string,
	authClient *auth.Client,
	authStore store.AuthStore,
	sessionStore store.SessionStore,
) *AuthHandler {
	return &AuthHandler{
		cfg:          cfg,
		serverAddr:   serverAddr,
		authClient:   authClient,
		authStore:    authStore,
		sessionStore: sessionStore,
	}
}

func (a *AuthHandler) RedirectToKeycloak(c *gin.Context) {
	stateID, err := utils.GenerateRandomBase64Str()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err = a.authStore.SetState(c, stateID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Redirect(http.StatusFound, a.authClient.Oauth.AuthCodeURL(stateID))
}

func (a *AuthHandler) RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login.tmpl", gin.H{
		"Title":       "Welcome",
		"Message":     "Please log in to continue",
		"KeycloakURL": fmt.Sprintf("http://%s/login-keycloak", a.serverAddr),
	})
}

func (a *AuthHandler) Logout(c *gin.Context) {
	// Get id_token from cookie for id_token_hint
	idToken, err := c.Cookie("id_token")
	if err != nil || idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id_token for logout"})
		return
	}
	logoutURL := fmt.Sprintf(
		"http://localhost:8080/realms/Devpool_project/protocol/openid-connect/logout?id_token_hint=%s&post_logout_redirect_uri=%s",
		idToken,
		a.cfg.Auth.LogoutRedirect,
	)
	// Optionally clear cookies
	c.SetCookie("id_token", "", -1, "/", "", true, true)
	c.SetCookie("session_id", "", -1, "/", "", true, true)
	c.SetCookie("user_email", "", -1, "/", "", true, true)
	c.Redirect(http.StatusFound, logoutURL)
}
func (a *AuthHandler) CallbackLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Logout Success"})
}
