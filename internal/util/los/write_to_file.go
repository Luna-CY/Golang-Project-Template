package los

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"os"
)

func WriteToFile(path string, content string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return errors.New("创建文件失败: %s, err: %s", path, err)
	}

	defer func() {
		_ = f.Close()
	}()

	if _, err := f.WriteString(content); nil != err {
		return errors.New("写入文件失败: %s, err: %s", path, err)
	}

	return nil
}
