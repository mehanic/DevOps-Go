package factorvar

import (
	"fmt"
	"io"
)

func Writeto(w io.Writer, lines ...string) (n int64, err error) {
	var nw int
	for _, line := range lines {
		nw, err = fmt.Fprint(w, line)
		n += int64(nw)
		if err != nil {
			return
		}
	}
	return
}
