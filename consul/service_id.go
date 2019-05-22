package consul

import (
	"crypto/md5"
	"fmt"
)

//创建服务的唯一id
func CreateAgentServiceUniqueID(name string, address string, port string) string {
	ID := fmt.Sprintf("service-name:%s,service-address:%s,service-port:%s", name, address, port)
	unique_id := md5.Sum([]byte(ID))
	return fmt.Sprintf("%x", unique_id)
}
