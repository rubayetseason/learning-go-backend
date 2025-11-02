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

// CreateJwt generates a JWT token using HMAC-SHA256
func CreateJwt(secret string, data Payload) (string, error) {

	// 1. Create the header with algorithm and type
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// 2. Convert header struct to JSON
	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	// 3. Encode the JSON header to Base64 URL format
	header64 := base64UrlEncode(byteArrayHeader)

	// 4. Convert payload struct to JSON
	byteArrayPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 5. Encode the JSON payload to Base64 URL format
	payload64 := base64UrlEncode(byteArrayPayload)

	// 6. Combine header + payload (without signature) → message to be signed
	message := header64 + "." + payload64

	// 7. Convert message and secret to byte arrays
	byte64Message := []byte(message)
	byteArraySecret := []byte(secret)

	// 8. Create a new HMAC using SHA-256 and the secret key
	h := hmac.New(sha256.New, byteArraySecret)

	// 9. Add the message (header.payload) to HMAC for hashing
	h.Write(byte64Message)

	// 10. Generate the final HMAC signature (in bytes)
	signature := h.Sum(nil)

	// 11. Encode the signature to Base64 URL format
	signature64 := base64UrlEncode(signature)

	// 12. Combine header, payload, and signature to get the final JWT token
	jwt := header64 + "." + payload64 + "." + signature64

	// 13. Return the token
	return jwt, nil
}

// Helper function to Base64 URL-encode data without padding (=)
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
