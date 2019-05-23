package consul

import (
	"crypto/md5"
	"fmt"
)

//创建服务的唯一id
func CreateAgentServiceUniqueID(name string) string {
	ID := fmt.Sprintf("zldz:%s:zldz", name)
	unique_id := md5.Sum([]byte(ID))
	return fmt.Sprintf("%x", unique_id)
}
