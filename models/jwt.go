package models

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	db "github.com/mrTavas/fw-backend/dbconn"
)

const tokenExpiredTime = 100000 // 1 - примерно 10-20 секунд, после токен просрочен
//const tokenExpiredTime = 1440
const refreshTokenExpiredMinutes = 201600
const configjwtSec = "mySecret"

// Sessions godoc
type Sessions struct {
	tableName           struct{}  `sql:"cws_sessions"`
	ID                  int       `json:"id" sql:",pk"`
	UserUUID            string    `json:"user_uuid"`
	RefreshToken        string    `json:"refresh_token" description:"Токен для обновления"`
	SessionEnd          time.Time `json:"session_end" description:"Дата когда токен отозван"`
	RefreshTokenUsed    time.Time `json:"refresh_token_used" description:"Дата использования токена"`
	RefreshTokenExpired time.Time `json:"refrash_expired" description:"Дата протухания токена"`
	CreatedAt           time.Time `sql:"default:now()" json:"created_at" description:"Дата создания"`
}

type (
	// LoginRequest requested data when logging in
	LoginRequest struct {
		Phone    int    `json:"phone"`
		Password string `json:"password"`
	}

	// LoginRefreshRequest godoc
	LoginRefreshRequest struct {
		RefreshToken string `json:"refresh"`
	}

	// TokenClaim JWT token structure
	TokenClaim struct {
		//Role   string `json:"role"`
		UserID int `json:"user_id"`
		Phone  int `json:"login"`
		jwt.StandardClaims
	}

	// LoginResponse responsed when requesting token
	LoginResponse struct {
		UserUUID               string    `json:"user_uuid"`
		Token                  string    `json:"token"`
		RefreshToken           string    `json:"refresh_token"`
		RefreshTokenExpiration time.Time `json:"refresh_expiration"`
	}
)

func (logResp *LoginResponse) NewRefreshToken(userID string) error {

	newToken, err := uuid.NewV4()
	if err != nil {
		return err
	}

	logResp.UserUUID = userID
	logResp.RefreshToken = newToken.String()

	dur := time.Minute * time.Duration(refreshTokenExpiredMinutes)
	logResp.RefreshTokenExpiration = time.Now().Add(dur)

	err = logResp.saveTokenData(newToken.String())
	if err != nil {
		return err
	}
	return nil
}

// saveTokenData expired existing and create new token for user
func (logResp *LoginResponse) saveTokenData(uuid string) error {

	var sessNew Sessions

	sessNew.UserUUID = logResp.UserUUID
	sessNew.RefreshToken = uuid
	sessNew.RefreshTokenExpired = logResp.RefreshTokenExpiration

	_, err := db.Conn.Model(&sessNew).Returning("*").Insert()
	if err != nil {
		return errors.New("Ошибка сохранения новой сессии")
	}

	return nil
}

func (logResp *LoginResponse) GenerateJWT(user Managers) error {

	jwtSec := configjwtSec
	mySigningKey := []byte(jwtSec)

	claims := TokenClaim{
		UserID: user.ID,
		//Role:   user.Role,
		Phone: user.Phone,
	}
	claims.IssuedAt = time.Now().Unix()

	dur := time.Minute * time.Duration(tokenExpiredTime)
	claims.ExpiresAt = time.Now().Add(dur).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return err
	}
	logResp.Token = ss

	return nil
}

func ExpireUserTokens(userUUID string) error {

	var sessOld Sessions

	_, err := db.Conn.Model(&sessOld).
		Set("refresh_token_used = ?", time.Now()).
		Where("user_uuid = ?", userUUID).
		Update()

	if err != nil {
		return err
	}
	return nil
}

func RefreshJWTToken(token string) (LoginResponse, error) {

	var newLogin LoginResponse

	User, err := expireToken(token)
	if err != nil {
		return newLogin, err
	}

	err = newLogin.NewRefreshToken(User.UUID)
	if err != nil {
		return newLogin, err
	}

	err = newLogin.GenerateJWT(User)
	if err != nil {
		return newLogin, err
	}

	return newLogin, nil
}

func expireToken(token string) (Managers, error) {

	var oper Managers
	var sessOld Sessions

	_, err := db.Conn.Model(&sessOld).
		Set("refresh_token_used = ?", time.Now()).
		Where("refresh_token = ? AND CURRENT_TIMESTAMP < refresh_token_expired AND refresh_token_used is NULL", token).
		Returning("*").
		Update(&sessOld)

	if err != nil {
		return oper, errors.New("Refresh token not found")
	}

	err = db.Conn.Model(&oper).
		Where("ID = ?", sessOld.UserUUID).
		First()
	if err != nil {
		return oper, err
	}
	return oper, nil
}

func (logResp *LoginResponse) GenerateJWTWorker(user Workers) error {

	jwtSec := configjwtSec
	mySigningKey := []byte(jwtSec)

	claims := TokenClaim{
		UserID: user.ID,
		//Role:   user.Role,
		Phone: user.Phone,
	}
	claims.IssuedAt = time.Now().Unix()

	dur := time.Minute * time.Duration(tokenExpiredTime)
	claims.ExpiresAt = time.Now().Add(dur).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return err
	}
	logResp.Token = ss

	return nil
}

func RefreshJWTTokenWorker(token string) (LoginResponse, error) {

	var newLogin LoginResponse

	User, err := expireTokenWorker(token)
	if err != nil {
		return newLogin, err
	}

	err = newLogin.NewRefreshToken(User.UUID)
	if err != nil {
		return newLogin, err
	}

	err = newLogin.GenerateJWTWorker(User)
	if err != nil {
		return newLogin, err
	}

	return newLogin, nil
}

func expireTokenWorker(token string) (Workers, error) {

	var oper Workers
	var sessOld Sessions

	_, err := db.Conn.Model(&sessOld).
		Set("refresh_token_used = ?", time.Now()).
		Where("refresh_token = ? AND CURRENT_TIMESTAMP < refresh_token_expired AND refresh_token_used is NULL", token).
		Returning("*").
		Update(&sessOld)

	if err != nil {
		return oper, errors.New("Refresh token not found")
	}

	err = db.Conn.Model(&oper).
		Where("ID = ?", sessOld.UserUUID).
		First()
	if err != nil {
		return oper, err
	}
	return oper, nil
}
