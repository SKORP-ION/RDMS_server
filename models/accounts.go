package models

import (
	u "Rostelecom_Device_Management_System/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"strings"
)

type Token struct {
	WsId uint
	jwt.StandardClaims
}

type Workstation struct {
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
	Token string `json:"token" ;sql:"-"`
}

func (ws *Workstation) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(ws.Name, "@") {
		return u.Message(false, "Email address is requires"), false
	}

	if len(ws.Password) < 64 {
		return u.Message(false, "Wrong key"), false
	}

	temp := &Workstation{}

	err := GetDB().Table("workstations").Where("ws_name = ?", ws.Name).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Already in use"), false
	}

	return u.Message(false, "Requirement passed"), true
}
