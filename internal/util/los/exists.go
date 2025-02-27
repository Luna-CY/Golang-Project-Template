package los

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"os"
)

func CheckPathExists(path string) (bool, errors.Error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, errors.New(errors.ErrorTypeServerInternalError, "IU.I_OS.CPE_TS.18", err)
}
