package langgo

import (
	"fmt"
	"io"
	"log"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}
}
