package fs

import (
	"fmt"
	"os"
)

/**
check a path is exist or not
*/
func PathExists(path string) (isExist bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func BasePath() (path string) {
	path, _ = os.Getwd()
	fmt.Println(path)
	return
}

func StoragePath() (path string) {
	path = BasePath() +
		string(os.PathSeparator) +
		"internal" +
		string(os.PathSeparator) +
		"app" +
		string(os.PathSeparator) +
		"storages" +
		string(os.PathSeparator)
	return
}
