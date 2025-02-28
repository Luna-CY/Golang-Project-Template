package los

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"os"
)

func WriteToFile(path string, content string) errors.Error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if nil != err {
		return errors.New(errors.ErrorTypeServerInternalError, "IUI_OS.WTF_LE.11", "创建文件失败: %s, err: %s", path, err)
	}

	defer func() {
		_ = f.Close()
	}()

	if _, err := f.WriteString(content); nil != err {
		return errors.New(errors.ErrorTypeServerInternalError, "IUI_OS.WTF_LE.19", "写入文件失败: %s, err: %s", path, err)
	}

	return nil
}
