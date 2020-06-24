package util

import (
	"crypto/sha256"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"translation_web/entry"
)


func BandingParams(transParams *entry.TranslationParams) {
	u1 := uuid.NewV4()
	stamp := time.Now().Unix()
	transParams.Salt = u1.String()
	transParams.CurTime = strconv.FormatInt(stamp, 10)
	transParams.SignType = "v3"
	input := truncate(transParams.Q)
	//生成签名
	instr := transParams.AppKey + input + u1.String() + transParams.CurTime + transParams.AppPwd
	sig := sha256.Sum256([]byte(instr))
	transParams.Sign = HexBuffToString(sig[:])
}
//加密
func HexBuffToString(buff []byte) string {
	var ret string
	for _, value := range buff {
		str := strconv.FormatUint(uint64(value), 16)
		if len([]rune(str)) == 1 {
			ret = ret + "0" + str
		} else {
			ret = ret + str
		}
	}
	return ret
}
//input处理
func truncate(q string) string {
	res := make([]byte, 10)
	qlen := len([]rune(q))
	if qlen <= 20 {
		return q
	} else {
		temp := []byte(q)
		copy(res, temp[:10])
		lenstr := strconv.Itoa(qlen)
		res = append(res, lenstr...)
		res = append(res, temp[qlen-10:qlen]...)
		return string(res)
	}
}
//发送请求
func DoPost(transParams entry.TranslationParams) []byte {
	data := make(url.Values, 0)
	data["q"] = []string{transParams.Q}
	data["from"] = []string{transParams.From}
	data["to"] = []string{transParams.To}
	data["appKey"] = []string{transParams.AppKey}
	data["salt"] = []string{transParams.Salt}
	data["sign"] = []string{transParams.Sign}
	data["signType"] = []string{"v3"}
	data["curtime"] = []string{transParams.CurTime}

	var resp *http.Response
	var err error
	resp, err = http.PostForm("https://openapi.youdao.com/api",data)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("json解析结果: ", string(body))
	return body
}
