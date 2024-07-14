package common

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

func GenerateUID(username string) string {
	nameParts := strings.Fields(username)

	randomString := generateRandomString()

	uniqueString := fmt.Sprintf("%s.%s", strings.ToLower(getInitials(nameParts)), randomString)

	return uniqueString
}

func generateRandomString() string {
	length := 8
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func getInitials(nameParts []string) string {
	var initials []byte
	for _, part := range nameParts {
		initials = append(initials, part[0])
	}
	return string(initials)
}

func GenerateFeedID(userID string) string {
	randomString := generateRandomString()
	return fmt.Sprintf("%s-%s", userID+"feed", randomString)
}
