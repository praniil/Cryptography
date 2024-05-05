package salting

import (
	"crypto/sha256"
	"encoding/base64"
	// "encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func generateSalt() []byte{
	return []byte("234ksdjf34234")	//hardcoded salt
}

func hashPassword (password string, salt []byte) string{
	saltedPassword := append([]byte(password), salt...)
	//second argument i.e salt... is variadic argument, youre passing each elemnt of the salt slice as separate argument to append
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func hashPass (pass string) string {
	hashed := sha256.Sum256([]byte(pass))
	// return hex.EncodeToString(hashed[:])
	return base64.StdEncoding.EncodeToString(hashed[:])
}

func HashingAndSalting () {
	password := "helloworld123"

	salt := generateSalt()
	hashedAndSaltedPassword := hashPassword(password, salt)
	hashedPassword := hashPass(password)
	fmt.Println("hashed password: ", hashedPassword)
	fmt.Println("hashed and salted password: ", hashedAndSaltedPassword)
}