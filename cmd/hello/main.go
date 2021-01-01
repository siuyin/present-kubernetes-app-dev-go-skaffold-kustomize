package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/siuyin/dflt"
)

func main() {
	fmt.Println("hello service starting on port 8080.")

	msg := dflt.EnvString("MESSAGE", "world")
	dispFile := dflt.EnvString("DISPLAY_FILE", "testdata/myfile.txt")
	passwd := dflt.EnvString("PASSWORD", "Ag00d.Password!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s. My password is %s.\n", msg, passwd) // changed from "Hello..."
		displayFileContents(w, dispFile)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func displayFileContents(w io.Writer, fn string) {
	f := openFileForRead(fn)
	defer f.Close()

	copyToWriter(f, w)
}
func openFileForRead(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
func copyToWriter(r io.Reader, w io.Writer) {
	_, err := io.Copy(w, r)
	if err != nil {
		log.Fatalf("could not copy to writer: %v", err)
	}
}
