package file

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestReadField(t *testing.T) {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}
