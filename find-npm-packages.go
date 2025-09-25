package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type PackageInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

func main() {
	var packages []PackageInfo

	if len(os.Args) != 2 {
		log.Fatal("ERROR invalid arguments provided\n\nUSAGE\n\tfind-npm-packages path")
	}

	err := filepath.WalkDir(os.Args[1], func(path string, d fs.DirEntry, err error) error {
		var packageInfo PackageInfo

		if d.IsDir() || d.Name() != "package.json" {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			log.Printf("skipping path: %v path=%q", err, path)
			return nil
		}

		if mode := info.Mode(); mode&fs.ModeType != 0 {
			log.Printf("skipping non-regular file: type=%q path=%q", mode, path)
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			log.Printf("could not open file: %v path=%q", err, path)
			return nil
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&packageInfo); err != nil {
			log.Printf("could not decode file: %v path=%q", err, path)
			return nil
		}
		packageInfo.Path = path
		packages = append(packages, packageInfo)
		return nil
	})

	if err != nil {
		log.Printf("error walking path: %v", err)
	}

	data, err := json.Marshal(packages)
	if err != nil {
		log.Fatalf("could not marshal data: %v", err)
	}

	fmt.Println(string(data))
}
