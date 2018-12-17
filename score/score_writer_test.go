package scoreboard

import (
	"bytes"
	"strings"
	"testing"
)

func TestWritingOfScoresToSpecifiedOut(t *testing.T) {

	buff := &bytes.Buffer{}
	sb := Monitor{IoWriter: buff}
	aggregator := NewAggregate(sb)
	writeMessage := "Hello"
	aggregator.Save(writeMessage)
	aggregator.Dump()

	readfromOut := strings.TrimSuffix(buff.String(), "\n")
	if readfromOut != writeMessage {
		t.Errorf("Message are wrtten to specified out sent:  `%s` and wrote `%s`",
			writeMessage, readfromOut)
	}

}
