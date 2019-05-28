package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"testing"
)

//jwt编码测试
func TestJwtTokenEncode(t *testing.T) {
	var my_data []string
	my_data = append(my_data, "梁宇")
	my_data = append(my_data, "胡益铭")
	my_data = append(my_data, "刘印午")
	my_data = append(my_data, "小明")
	my_data = append(my_data, "阿花")
	toek_map := jwt.MapClaims{}
	toek_map["name"] = "liangyu"
	toek_map["my_data"] = my_data
	token, err := JwtTokenEncode(toek_map, 60)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("token结果是:%s", token)
}

//jwt解码测试
func TestJwtTokenDecode(t *testing.T) {

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTkwMjcyMDMsIm15X2RhdGEiOlsi5qKB5a6HIiwi6IOh55uK6ZOtIiwi5YiY5Y2w5Y2IIiwi5bCP5piOIiwi6Zi_6IqxIl0sIm5hbWUiOiJsaWFuZ3l1In0.pMUtC6FeY6wMYSnLPdHPykEUaTKLiAJMd4rG8WQi2Qg"

	jwt_map, err := JwtTokenDecode(token)
	if err != nil {
		log.Fatalf("错误%v", err)
	}

	log.Println(jwt_map["my_data"])
	//这里会输出变量的类型，也可以用golang的反射，得出变量的类型reflect.TypeOf
	log.Printf("%T", jwt_map["my_data"])
	//动态打印一下类型
	switch my_data_type := jwt_map["my_data"].(type) {
	case []interface{}:
		log.Println("切片")
	default:
		log.Println("什么类型？")
		log.Println(my_data_type)
	}

}
