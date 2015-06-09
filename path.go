package spruce

import (
	"fmt"
	"os"
	"path/filepath"
)

func AbsPathify(src string) string {

	if filepath.IsAbs(src) && src != "/" {
		return filepath.Clean(src)
	}

	cwd, err := os.Getwd()
	if err != nil {
		// TODO Handle error
		fmt.Errorf("ERROR: %v", err)
		cwd = "./"
	}

	absPath := filepath.Join(cwd, src)

	return filepath.Clean(absPath)
}
