package go_file_util

import (
	"crypto/md5"
	"fmt"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

// sign 签名
// 传入密码,加密
func SignJWT(secret string, uname string, upass string) (jwtStr string) {
	key := []byte(secret)
	fmt.Println(secret)
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key},
		(&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}

	cl := jwt.Claims{
		// Registered claims : 这里有一组预定义的声明，它们不是强制的，但是推荐
		// 比如：iss (issuer), exp (expiration time), sub (subject), aud (audience)等。
		Issuer:    uname,
		Subject:   upass,
		NotBefore: jwt.NewNumericDate(time.Now()),
		Audience:  jwt.Audience{"name", "admin"},
	}
	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		panic(err)
	}

	// fmt.Println(raw)
	return raw
}

// 解析jwt
// 传入key(之前加密的密码),raw(jwt令牌)
func ParseJWT(key string, raw string) {
	var sharedKey = []byte(key)
	tok, err := jwt.ParseSigned(raw)
	if err != nil {
		panic(err)
	}
	out := jwt.Claims{}
	// 解析出issuer(uname)和subject(upass),校验
	if err := tok.Claims(sharedKey, &out); err != nil {
		panic(err)
	}
	fmt.Printf("iss: %s, sub: %s\n", out.Issuer, out.Subject)
}

// DM5加密
func MD5_malred(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// 销毁TokenMap的方法
// 定时销毁token(默认2小时)
func DestoryTokenMap(tokenMap map[string]string) {
	for k := range tokenMap {
		delete(tokenMap, k)
	}
}
