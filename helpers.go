package auth

// File provides global helpers for salt and hashing.

import (
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSALT Generates salt based on settings value.
func GenerateSALT(length int) string {
	return randomString(length)
}

// GenerateSessionID Generates session ID based on settings value and UNIX time.
func GenerateSessionID(length int) string {
	return UXSecs() + randomString(length)
}

// HASHPassword Uses bcrypt to generate password hash.
// Suggested hash cost is 14.
func HASHPassword(password, salt string, hashCost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password+salt), hashCost)
}

// UXSecs Returns UNIX time in seconds.
func UXSecs() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// UXNano Returns UNIX time in nanoseconds.
func UXNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func CheckPasswordHash(password, salt, hashedPassword string) bool {
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
	return errCompare == nil
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := make([]byte, length)
	for i := range result {
		result[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(result)
}
