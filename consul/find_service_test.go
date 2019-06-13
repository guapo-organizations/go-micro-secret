package consul

import (
	"testing"
)

//命令行运行 go test -v  ; -v 加上-v(verbose)选项就可以看到完整过程，否则无法打印日志
func TestFindService(t *testing.T) {
	CreateConfig("106.122.76.72", "8500")
	config, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	service, err := FindService(config, "zldz.sms", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", service)
}
