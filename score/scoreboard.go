package scoreboard

import (
	"fmt"
	"io"
)

type Monitor struct {
	IoWriter io.Writer
}

//Custom logs goes here
func (m Monitor) Write(p []byte) (n int, err error) {
	return fmt.Fprintln(m.IoWriter, string(p))
}
