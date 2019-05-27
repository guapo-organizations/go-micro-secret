package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey []byte

func init() {

	if mySigningKey == nil {
		mySigningKey = []byte("zldzJwtSigninKey")
	}
}

//jwt 编码,
//jwt.MapClaims是编码的东西
//minute 是token最多可以维持多久有效
func JwtTokenEncode(jwt_map_claims jwt.MapClaims, minute time.Duration) (string, error) {
	if minute != 0 {
		//过期时间是一个时间戳，先获取当前时间，加上什么时候结束，转化为int
		jwt_map_claims["exp"] = time.Now().Add(minute * time.Minute).Unix()
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt_map_claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("jwt加密失败：%s", err)
	}
	return tokenString, nil
}

//jwt 解码
func JwtTokenDecode(jwt_token string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(jwt_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil
	}

	//过期了
	return nil, fmt.Errorf("无效的令牌")
}
