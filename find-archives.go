package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("USAGE:\n\tfind-archives <search-root>")
	}

	searchRoot := os.Args[1]
	filepath.WalkDir(searchRoot, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if d.Type() & os.ModeSymlink != 0 {
			log.Printf("[WRN] skipping symlink: path=%q", path)
		}

		file, err := os.Open(path)
		if err != nil {
			log.Printf("[ERR] skipping file: error=%q, path=%s", err, path)
			return nil
		}
		defer file.Close()

		buf := make([]byte, 512)
		_, err = file.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("[ERR] skipping file: error=%q, path=%s", err, path)
				return nil
			}
		}

		kind, _ := filetype.Archive(buf)
		if kind == types.Unknown {
			return nil
		}

		fmt.Println(kind.MIME.Value, path)

		return nil
	})
}
