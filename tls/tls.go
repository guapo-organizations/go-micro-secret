package tls

import (
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	//获取当前文件，也就是tls文件夹的绝对路径
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}


//文件的绝对路径
//
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}
