package hmac

//hmac: type fo cryptographic hash function that uses a secret key to authenticate a message
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateHMACSHA256(message string, secretKey []byte) string{
	h := hmac.New(sha256.New, secretKey)		//only instance or setup to add messages
	h.Write([]byte(message))
	hmacValue := h.Sum(nil)		//execution
	return hex.EncodeToString(hmacValue)
}

func verifyHMACSHA256(message string, secretKey []byte, expectedHmacCode string) bool {
	expectedHmacBytes, err := hex.DecodeString(expectedHmacCode)
	if err != nil {
		return false
	}
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(message))
	hmacValue := h.Sum(nil)

	result := hmac.Equal(hmacValue, expectedHmacBytes)
	return result

}

func Hmac () {
	fmt.Println("in hmac")
	message:= "hey mate how are you?"
	secretKey := []byte("secret")
	hmacSHA256 := generateHMACSHA256(message, secretKey)
	fmt.Println("hmac SHA256: ", hmacSHA256)

	verifyHMAC := verifyHMACSHA256(message, secretKey, hmacSHA256)
	fmt.Println("verification: ", verifyHMAC)

	verifyHMAC = verifyHMACSHA256("hello Kate u so beautiful", secretKey, hmacSHA256)
	fmt.Println("verification: ", verifyHMAC)

}