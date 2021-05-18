package security

import (
	"RDMS_server/database"
	"RDMS_server/structs"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

type Token struct {
	Ws_name string `json:"name"`
	jwt.StandardClaims
	Age time.Time `json:"created"`
	Token string `json:"token"`
}

func JwtAuth (r *http.Request) (bool, error) {
	tk := Token{}
	tk.Token = r.Header.Get("Authorization")
	tk.Ws_name = r.Header.Get("Workstation_name")

	if !database.WorkstationAvailability(tk.Ws_name) {
		return false, errors.New("Workstation not found")
	}

	if tk.Token == "" {
		return false, errors.New("Missing Token")
	} else if tk.Ws_name == "" {
		return false, errors.New("Empty Workstation_name header")
	}

	token, err := jwt.ParseWithClaims(tk.Token, &tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	if err != nil {
		return false, errors.New("Invalid token")
	}

	if !token.Valid {
		return false, errors.New("Token is not valid")
	}

	return true, nil

}

func CreateToken(ws structs.Workstation) (Token, error) {
	tk := Token{Ws_name:ws.Name}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	newToken, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return tk, err
	}
	tk.Token = newToken
	tk.Age = time.Now()
	return tk, nil
}