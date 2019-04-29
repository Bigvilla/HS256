package HS256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//Header 定义加密算法
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"JWT"`
}

//PayLoad 定义PayLoad
type PayLoad struct {
	Exp int64  `json:"exp"`
	Iss string `json:"iss"`
}

//Token 输出生成的Token
func Token(key, secret string) {
	token := GetToken(key, secret)
	fmt.Println(token)
}

//GetToken 定义获取token的算法
func GetToken(key, secret string) string {
	dtime, _ := time.ParseDuration("24h") //默认24h
	expirtime := time.Now().Add(dtime).Unix()
	var HeaDetail = Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	var PayData = PayLoad{
		Exp: expirtime,
		Iss: key,
	}
	jsonHead, _ := json.Marshal(HeaDetail)
	jsonData, _ := json.Marshal(PayData)
	Secret := secret
	TokenValue := GenToken(jsonHead, jsonData, PayData.Iss, Secret)
	return TokenValue
}

//Gensign 使用HS256算法生成签名
func Gensign(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	signature := strings.TrimRight(base64.URLEncoding.EncodeToString(expectedMAC), "=")
	return signature
}

//GenToken 生成token
func GenToken(jsonHead []byte, jsonData []byte, key, secret string) string {
	EnHeader := strings.TrimRight(base64.URLEncoding.EncodeToString(jsonHead), "=")
	EnPay := strings.TrimRight(base64.URLEncoding.EncodeToString(jsonData), "=")
	Str := strings.Join([]string{EnHeader, EnPay}, ".")
	TokenStr := Gensign(Str, secret)
	token := Str + "." + TokenStr
	return token
}
