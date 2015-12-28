package utils

import (
	"os"
	"path"
	"regexp"
)

func GetProjectName() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("[^a-z0-9]")
	return re.ReplaceAllString(path.Base(wd), ""), nil
}
