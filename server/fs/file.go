package fs

import (
	"io/ioutil"
	"log"
)

// FileJSON wraps the os.FileInfo nicely.
type FileJSON struct {
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	IsDirectory bool   `json:"isDirectory"`
}

// GetDirectoryItems Gets a list of directory items.
func GetDirectoryItems(path string) []FileJSON {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	fileItems := make([]FileJSON, len(files))
	for i, f := range files {
		fileItems[i] = FileJSON{Name: f.Name(), Size: f.Size(), IsDirectory: f.IsDir()}
	}

	return fileItems
}
