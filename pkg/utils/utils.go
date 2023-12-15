package utils

import (
	"os"

	"github.com/artyomtugaryov/vpnbot/pkg/utils/errors"
)

func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errors.Wrap("Cannot check if directory exists", err)
}

func MkDir(path string) error {
	exists, err := IsExist(path)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return errors.Wrap("Cannot create a directory", err)
	}
	return nil
}
