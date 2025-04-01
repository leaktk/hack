package main

import (
	"io"
	"mime/quotedprintable"
	"os"
)

func main() {
  if len(os.Args) == 2 && os.Args[1] == "-d" {
    // Decode quote-printable
	  io.Copy(os.Stdout, quotedprintable.NewReader(os.Stdin))
  } else {
    // Encode quote-printable
	  w := quotedprintable.NewWriter(os.Stdout)
	  defer w.Close()
	  io.Copy(w, os.Stdin)
  }
}
