# HS256
## 一、说明
### 1、 遵循JWT（JSON Web Token）开放标准
### 2、 JWT由三部分组成
- Header
- Payload
- Signature

**备注：** HS256.go生成的样式：aaaaaa.bbbbbbb.cccccc
## 二、使用
    go get -u "github/Bigvilla/HS256"       
    import "HS256"
    HS256.Token(key,secret)

