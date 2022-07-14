package filestat

import (
	"strings"
	"testing"
)

func TestSplitLineToWords(t *testing.T) {
	sb := "xxxxx!yyyyy:zzz.aaa.bbb.cc:dd:ee-ff"
	sa := "xxxxx yyyyy zzz aaa bbb cc dd ee ff"
	words := splitLineToWords(sb, " .:-!")
	sr := strings.Join(words, " ")

	if sa != sr {
		t.Errorf("split test error")
	} else {
		t.Logf("split test pass")
	}
}
