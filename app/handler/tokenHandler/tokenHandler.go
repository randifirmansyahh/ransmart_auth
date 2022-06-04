package tokenHandler

import (
	"fmt"
	"log"
	"net/http"
	"ransmart_auth/app/helper/helper"
	"ransmart_auth/app/helper/response"
	"ransmart_auth/app/helper/tokenHelper"
	"reflect"

	"github.com/golang-jwt/jwt"
)

var (
	AUD            = tokenHelper.AUD
	ISS            = tokenHelper.ISS
	LOGIN_SECRET   = tokenHelper.LOGIN_SECRET
	MESSAGE        = "Unathorized!"
	MESSAGE_KOSONG = "Request authorize kosong"
)

func GetToken(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ngambil header authorization request
		authorization := r.Header.Get("authorization")

		if reflect.ValueOf(authorization).IsZero() {
			// Handling pertama Jika Header Authorization nya kosong
			log.Println(MESSAGE_KOSONG)
			response.Response(w, http.StatusUnauthorized, MESSAGE_KOSONG, nil)
			return
		}

		// ngekstrak token
		newToken := helper.ExtractTokens(authorization)

		// nge parse
		_, err := jwt.Parse(newToken, func(token *jwt.Token) (interface{}, error) {

			// cek audience
			CheckAudience := token.Claims.(jwt.MapClaims).VerifyAudience(AUD, false)
			if !CheckAudience {
				return nil, fmt.Errorf("invalid audience")
			}

			// cek iss
			CheckISS := token.Claims.(jwt.MapClaims).VerifyIssuer(ISS, false)
			if !CheckISS {
				return nil, fmt.Errorf("invalid iss")
			}

			return []byte(LOGIN_SECRET), nil
		})

		// unauth
		if err != nil {
			log.Println("401 unauthorized !!")
			response.Response(w, http.StatusUnauthorized, MESSAGE, nil)
			return
		}

		// end
		f.ServeHTTP(w, r)
	})
}
