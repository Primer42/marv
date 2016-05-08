package marv

import (
	"bufio"
	"bytes"
)

func Preprocess(dirty []byte) (clean []Line, err error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(dirty))

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			clean = append(clean, Line{scanner.Text()})
		}
	}

	return
}
