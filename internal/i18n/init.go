package i18n

import (
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
	"strings"
)

var languages = map[string]map[string]string{}

func Init() error {
	if err := filepath.Walk(filepath.Join("config", "i18n"), func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if nil != err {
			return err
		}

		var messages = map[string]string{}
		if err := toml.NewDecoder(f).Decode(&messages); nil != err {
			return err
		}

		languages[strings.Split(info.Name(), ".")[0]] = messages
		_ = f.Close()

		return nil
	}); nil != err {
		return err
	}

	return nil
}
