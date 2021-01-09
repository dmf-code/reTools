package fs

import (
	"os"
	"path"
)

func IsDir(dir string) (err error) {
	parent := path.Dir(dir)

	if _, err = os.Stat(parent); os.IsNotExist(err) {
		if err = IsDir(parent); err != nil {
			return
		}
	}

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
	}

	return
}
