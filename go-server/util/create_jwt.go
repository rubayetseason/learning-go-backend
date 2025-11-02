package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"isShopOwner"`
	Exp         int64  `json:"exp"`
	Iat         int64  `json:"iat"`
}

func CreateJwt(secret string, data Payload) (string, error) {

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrayHeader, err := json.Marshal(header) // creates the byte array of the header
	if err != nil {
		return "", err
	}

	header64 := base64UrlEncode(byteArrayHeader) // encodes the header to base64url

	byteArrayPayload, err := json.Marshal(data) // creates the byte array of the payload
	if err != nil {
		return "", err
	}

	payload64 := base64UrlEncode(byteArrayPayload) // encodes the payload to base64url

	message := header64 + "." + payload64
	byte64Message := []byte(message) // creates the byte array of the message

	byteArraySecret := []byte(secret) // creates the byte array of the secret

	h := hmac.New(sha256.New, byteArraySecret)

	h.Write(byte64Message)

	signature := h.Sum(nil) // creates the signature

	signature64 := base64UrlEncode(signature) // encodes the signature to base64url

	jwt := header64 + "." + payload64 + "." + signature64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
