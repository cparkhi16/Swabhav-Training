package controller

import (
	"fmt"
	"net/http"

	"userPassport/customlogger"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var secretKey = []byte("yogesh")

// Middleware function, which will be called for each request
func CheckAuthentication(next http.Handler) http.Handler {
	var l = customlogger.GetLoggerInstance()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ((r.URL.Path == "/courses/") && (r.Method == "GET")) || ((r.URL.Path == "/login" || r.URL.Path == "/checkToken") && r.Method == "POST") || (r.URL.Path == "/users/" && r.Method == "POST") || (r.URL.Path == "/files/" && r.Method == "POST") {
			l.Debug().Str("route-", r.URL.Path).Msg("Visited Unguarded Route")
			next.ServeHTTP(w, r)
			return
		}
		l.Debug().Str("route-", r.URL.Path).Msg("Visited Guarded Route")
		tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
		if err != nil {
			//http.Error(w, "Forbidden", http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access Denied; Please check the access token"))
			l.Error().Err(err).Msg("Access Denied; Please check the access token")
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secretKey, nil
		})

		if err != nil {
			//http.Error(w, "Forbidden", http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, err.Error())
			// w.Write([]byte("Access Denied; Please check the access token"))
			// l.Error().Err(err).Msg("Access Denied; Please check the access token")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// If token is valid
			// We found the token in our map
			//log.Printf("Authenticated user %s\n", claims)
			l.Error().Interface("claims", claims).Msg("Authenticated user")
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			//http.Error(w, "Forbidden", http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access Denied; Please check the access token"))
			l.Error().Err(err).Msg("Access Denied; Please check the access token")
			return
		}
	})
}
