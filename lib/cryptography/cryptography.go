package cryptography

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

var (
	argon2IdHashPasses    uint32 = 3
	argon2IdHashMemory    uint32 = 64 * 1024
	argon2IdHashThreads   uint8  = 4
	argon2IdHashKeyLength uint32 = 128 // 1024 bit key like the js client
)

func Argon2Id(password string, optionalSalt ...[]byte) ([]byte, []byte) {
	passwordBytes := []byte(password)

	salt := []byte{}
	if len(optionalSalt) != 1 {
		salt = RandomBytes(16)
	} else {
		salt = optionalSalt[0]
	}

	if pepper := os.Getenv("CRYPTO_PEPPER"); pepper != "" {
		pepperedMasterHash := append(passwordBytes, []byte(pepper)...)
		return salt, argon2.IDKey(pepperedMasterHash, salt, argon2IdHashPasses, argon2IdHashMemory, argon2IdHashThreads, argon2IdHashKeyLength)
	}

	return salt, argon2.IDKey(passwordBytes, salt, argon2IdHashPasses, argon2IdHashMemory, argon2IdHashThreads, argon2IdHashKeyLength)
}

func CompareTokens(token string, strengthenedToken []byte, salt []byte) bool {
	_, testStrengthenedToken := Argon2Id(token, salt)
	return subtle.ConstantTimeCompare(strengthenedToken, testStrengthenedToken) == 1
}

func RandomBytes(n int) []byte {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return bytes
}

func RandomHex(n int) string {
	bytes := RandomBytes(n / 2)
	return hex.EncodeToString(bytes)
}

func RandomUUID() uuid.UUID {
	random, _ := uuid.NewRandom()
	return random
}
