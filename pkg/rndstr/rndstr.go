// Generate Random string
package rndstr

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(seed int, length int) string {

	rand.Seed(time.Now().Unix() + int64(seed))
	charSet := "abcdedfghijklmnopqrst"
	var output strings.Builder

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}
