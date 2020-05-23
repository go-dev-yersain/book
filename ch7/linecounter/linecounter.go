package linecounter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// LineCounter counts the lines written to it
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	io.LimitedReader
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		return 0, err
	}
	*c = LineCounter(count)
	return count, nil
}
