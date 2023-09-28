package gate

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/dgrijalva/jwt-go"
)

func (c *Gate) Protected(next http.HandlerFunc, roles ...string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			StatusCode int
			Data       interface{}
			Message    string
		}{
			http.StatusUnauthorized,
			nil,
			"Operasi tidak diijinkan",
		}

		// Get the authorization header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Check if the authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Extract the token from the authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Here you would verify that the signing method is correct
			// and that the key used to sign the token is valid
			// For simplicity, we'll just return an arbitrary key here
			return []byte(utstring.GetEnv(models.AppApiSecret)), nil
		})

		// Check if there was an error parsing the token
		if err != nil {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Get the userId claim from the token
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Add the userId claim to the request context
		var session models.Session
		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}
		session.UserID = userID
		session.ExpiresAt = claims.ExpiresAt

		if len(roles) > 0 {
			userRole, errx := c.UserUsecase.UserRoleGetByID(context.Background(), userID)
			if errx != nil {
				c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
				return
			}

			roleAllowed := false
			for _, role := range roles {
				if role == userRole.RoleName {
					roleAllowed = true
					break
				}
			}
			if !roleAllowed {
				c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
				return
			}
		}

		ctx := context.WithValue(r.Context(), "session", session)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (c *Gate) XenditProtected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			StatusCode int
			Data       interface{}
			Message    string
		}{
			http.StatusUnauthorized,
			nil,
			"Operasi tidak diijinkan",
		}

		// Get the authorization header from the request
		authHeader := r.Header.Get("x-callback-token")
		if authHeader == "" {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		xenditCallbackToken := utstring.GetEnv(models.XenditCallbackTokenEnv, "just-for-ignoring-not-setup-env")

		if authHeader != xenditCallbackToken {
			resp.Message = "Operasi tidak diijinkan"
			c.ResponseJSON(w, resp.StatusCode, resp.Data, resp.Message)
			return
		}

		// Call the next handler with the new context
		next.ServeHTTP(w, r)
	})
}
