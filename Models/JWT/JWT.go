package JWT

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Iss string 	`json:"iss"`
	Exp uint 	`json:"exp"`
	Sub string 	`json:"sub"`
	Aud int 	`json:"aud"` //用户ID
	Ndf uint 	`json:"ndf"`
	Iat uint 	`json:"iat"`
	Jti uint 	`json:"jti"`
}

type Jwt struct{
	Header Header
	Payload Payload
	Secret string
}

/* func (j Jwt)Encoding() string
* 参照当前Header与PayLoad的值刷新Jwt
* 成功返回刷新后的Jwt，失败则令Jwt"",返回""
****************************************************/
func (j Jwt)Encoding() string {
	header,err:=json.Marshal(j.Header)
	if err!=nil {
		return ""
	}
	Header1:=base64.StdEncoding.EncodeToString(header)
	payload,err:=json.Marshal(j.Payload)
	if err!=nil {
		//j.Jwt=""
		return ""
	}
	Payload1:=base64.StdEncoding.EncodeToString(payload)
	hash := hmac.New(sha256.New,[]byte(j.Secret))
	hash.Write(
		[]byte(Header1+ "."+
			Payload1+ "."))

	return Header1+"."+Payload1+"."+base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

/* func (j *Jwt)Decoding(JWT string) error
* 参照当前的JWT刷新Payload的值
 ****************************************/
func (j *Jwt)Decoding(jwt string) error {
	hps:=strings.Split(jwt,".")
	if len(hps)!=3 {
		return fmt.Errorf("error: RefreshHP Signature error")
	}
	p,err:=base64.StdEncoding.DecodeString(hps[1])
	if err!=nil{
		return err
	}
	//fmt.Println(string(p))
	err = json.Unmarshal(p,&j.Payload)
	return err
}