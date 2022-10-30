package lib

import (
	"encoding/json"
	"burlyeducation/log"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Authorization string
	AppId         string
}

type Error struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

type AuthClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Mobile string `json:"verified_mobile"`
}

var clients map[string]map[string]string = make(map[string]map[string]string)

func init() {
	secrets, _ := config.String("jwt::clients")
	json.Unmarshal([]byte(secrets), &clients)
}

func (auth Auth) IsAuthorized() (bool, string) {

	tokenParams := strings.Split(auth.Authorization, " ")
	tokenType := tokenParams[0]
	strToken := tokenParams[1]
	appId := auth.AppId
	secret := clients[appId]["secret"]

	secretManager, _ := config.String("SECRET_MANAGER")
	if secretManager == "aws" {
		secret, err := AWSSecretManager{}.GetApiSecret()
		if err != nil {
			log.Warning(1104, map[string]interface{}{"error_details": err})
			return false, "Token, Invalid auth"
		}
	}

	if len(secret) == 0 || tokenType != "Bearer" {
		return false, "Token, Invalid format"
	}

	//Validate token
	var res string
	var resBool bool
	token, err := Tokener{}.Validate(strToken, secret, &AuthClaims{})

	if err == nil {
		myClaim := token.Claims.(*AuthClaims)
		res = fmt.Sprintf("%s %d %s, %v", "Valid", myClaim.UserId, myClaim.Issuer, myClaim.ExpiresAt)
		resBool = true
	} else {
		res = fmt.Sprintf("%s", err)
		log.Warning(1102, map[string]interface{}{"error_details": res})
		resBool = false
	}

	return resBool, res
}

func ApplyAuth(ctx *context.Context) {

	isAuthEnable, _ := config.Bool("jwt::enable_authentication")
	if isAuthEnable {
		var authorization string = ctx.Input.Header("Authorization")
		var appId string = ctx.Input.Header("X-TNL-TOKEN-TYPE")

		if authorization != "" {
			isAuthorized, _ := Auth{Authorization: authorization, AppId: appId}.IsAuthorized()
			if !isAuthorized {
				ctx.ResponseWriter.WriteHeader(401)
				ctx.Output.JSON(Error{0, 11103, "Invalid Token"}, true, true)
			}
		} else {
			log.Warning(1503, map[string]interface{}{"error_details": "Missing Auth header"})
			ctx.ResponseWriter.WriteHeader(401)
			ctx.Output.JSON(Error{0, 11103, "Please pass Auth header"}, true, true)
		}
	}
}
