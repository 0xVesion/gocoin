package core

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateNonce() string {
	length := 64
	b := make([]byte, length)
	rand.Read(b)

	return fmt.Sprintf("%x", b)[:length]
}

func CurrentTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func ParseTimestamp(timestamp string) time.Time {
	t, _ := time.Parse(time.RFC3339, timestamp)

	return t
}
