package wordcounter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// WordCounter merely counts the words written to it
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		return 0, err
	}
	*c = WordCounter(count)
	return count, nil
}
