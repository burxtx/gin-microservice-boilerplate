package transport

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/burxtx/gin-microservice-boilerplate/app/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func CasAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		secret := cfg.GetString("http.secret")
		session_domain := cfg.GetString("http.session_domain")
		store := sessions.NewCookieStore([]byte(secret))
		session, getErr := store.Get(c.Request, session_domain)
		if getErr != nil {
			http.Error(c.Writer, getErr.Error(), http.StatusInternalServerError)
		}
		var ticket string
		callback := fmt.Sprintf("%s://%s%s", "http", c.Request.Host, c.Request.URL.Path)
		cas := fmt.Sprintf("%s/login?service=%s", cfg.GetString("auth.cas"), callback)
		v := session.Values["user"]
		if v != nil {
			// http.Redirect(c.Writer, c.Request, cas, 302)
			// } else {
			c.Next()
			return
		}

		ticket = getTicketParam(c)
		if len(ticket) == 0 {
			http.Redirect(c.Writer, c.Request, cas, 302)
			return
		}

		validateUrl := fmt.Sprintf("%s/validate?service=%s&ticket=%s", cfg.GetString("auth.cas"), callback, ticket)
		client := http.Client{}
		request, requestErr := http.NewRequest("GET", validateUrl, nil)
		if requestErr != nil {
			http.Error(c.Writer, requestErr.Error(), http.StatusNotAcceptable)
		}
		resp, validateErr := client.Do(request)
		if validateErr != nil {
			http.Error(c.Writer, validateErr.Error(), http.StatusNotAcceptable)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
		content := string(body)
		lines := strings.Split(content, "\n")
		if lines[0] != "yes" {
			http.Error(c.Writer, "Get username authorization info failed!", http.StatusInternalServerError)
		}
		session.Values["user"] = lines[1]
		session.Save(c.Request, c.Writer)
		c.Next()
	}
}

func getTicketParam(c *gin.Context) string {
	requestErr := c.Request.ParseForm()
	if requestErr != nil {
		http.Error(c.Writer, requestErr.Error(), http.StatusBadRequest)
	}
	return c.Request.Form.Get("ticket")
}
