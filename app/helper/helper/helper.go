package helper

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"ransmart_auth/app/models/userModel"
	"strconv"
	"strings"
	"time"
)

func CheckEnv(err error) {
	if err != nil {
		log.Fatal("Failed to load environment")
	}
}

const alphabet = "./ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var bcEncoding = base64.NewEncoding(alphabet)

func Encode(src []byte) []byte {
	n := bcEncoding.EncodedLen(len(src))
	dst := make([]byte, n)
	bcEncoding.Encode(dst, src)
	for dst[n-1] == '=' {
		n--
	}
	return dst[:n]
}

func Decode(src []byte) ([]byte, error) {
	numOfEquals := 4 - (len(src) % 4)
	for i := 0; i < numOfEquals; i++ {
		src = append(src, '=')
	}

	dst := make([]byte, bcEncoding.DecodedLen(len(src)))
	n, err := bcEncoding.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

func UnMarshall(from string, to interface{}) {
	if err := json.Unmarshal([]byte(from), &to); err != nil {
		log.Println(err.Error())
		return
	}
}

func SearchOneUser(models []userModel.User, id int) (interface{}, error) {
	var oneData userModel.User
	for i := 0; i < len(models); i++ {
		if models[i].Id == id {
			oneData = models[i]
			return oneData, nil
		}
	}
	return nil, errors.New("data not found")
}

func SearchUser(models []userModel.User, id int) (interface{}, bool) {
	var oneData userModel.User
	for i := 0; i < len(models); i++ {
		if models[i].Id == id {
			oneData = models[i]
			return oneData, true
		}
	}
	return nil, false
}

func ExtractTokens(token string) string {
	/// Bearer asdhaskdkasdhsagdhasdgjasgdhasdghadgjadsaj
	strarr := strings.Split(token, " ")
	if len(strarr) == 2 {
		return strarr[1]
	}
	return ""
}

func ExpiredTime(menit int) time.Time {
	return time.Now().Add(time.Duration(menit) * time.Minute) // expired date
}

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func StringToint(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("error =>", err)
	}
	return i
}
