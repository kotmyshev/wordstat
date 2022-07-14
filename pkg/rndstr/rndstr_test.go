// Generate Random string
package rndstr

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {

	var rnds string
	lena := 42

	rnds = GenerateRandomString(42, lena)
	fmt.Println("Random String: ", rnds)
	if len(rnds) != lena {
		t.Errorf("invalid length")
	}

}
