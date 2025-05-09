package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
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

		buf := make([]byte, 0, 512)
		_, err = file.Read(buf)
		if err != nil {
			log.Printf("[ERR] skipping file: error=%q, path=%s", err, path)
			return nil
		}

		if filetype.IsArchive(buf) {
			fmt.Println(path)
		}

		return nil
	})
}
