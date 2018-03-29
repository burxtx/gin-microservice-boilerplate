package endpoints

import (
	"fmt"
	"net/http"

	"github.com/burxtx/gin-microservice-boilerplate/app/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type AuthEndpoint struct {
}

func (a *AuthEndpoint) Logout(c *gin.Context) {
	cfg := config.GetConfig()
	secret := cfg.GetString("http.secret")
	session_domain := cfg.GetString("http.session_domain")
	store := sessions.NewCookieStore([]byte(secret))
	callback := fmt.Sprintf("%s://%s%s", "http", c.Request.Host, "/")
	cas := fmt.Sprintf("%s/logout?service=%s", cfg.GetString("auth.cas"), callback)
	session, err := store.Get(c.Request, session_domain)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
	var msg []byte
	if session != nil {
		session.Options.MaxAge = -1
		saveErr := session.Save(c.Request, c.Writer)
		if saveErr != nil {
			http.Error(c.Writer, saveErr.Error(), http.StatusInternalServerError)
		}
		http.Redirect(c.Writer, c.Request, cas, 302)
	} else {
		msg = []byte("already logout!")
	}
	c.JSON(200, gin.H{"msg": msg})
	return
}

func (a *AuthEndpoint) GetCurrentUser(c *gin.Context) {
	cfg := config.GetConfig()
	secret := cfg.GetString("http.secret")
	session_domain := cfg.GetString("http.session_domain")
	store := sessions.NewCookieStore([]byte(secret))
	session, err := store.Get(c.Request, session_domain)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	user := session.Values["user"]
	c.JSON(200, gin.H{"user": user})
}
