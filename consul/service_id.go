package consul

import (
	"crypto/md5"
	"fmt"
)

//创建服务的唯一id
func createAgentServiceUniqueID(name string, address string, port int) string {
	ID := fmt.Sprintf("service-name:%s,service-address:%s,service-port:%d", name, address, port)
	unique_id := md5.Sum([]byte(ID))
	return fmt.Sprintf("%x", unique_id)
}
