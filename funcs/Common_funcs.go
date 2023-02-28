package funcs

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"path"
	"rj97807_work_serve/utils"
	"time"
)

type UserClaim struct {
	Uid  string
	Name string
	Role int
	jwt.StandardClaims
}

var Ak = utils.QiniuAK
var Sk = utils.QiniuSK
var Bucket = utils.QiniuBucket

var putPolicy = storage.PutPolicy{Scope: Bucket}
var mac = qbox.NewMac(Ak, Sk)
var UpToken = putPolicy.UploadToken(mac)

// Md5 加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// YieldToken 生成Token
func YieldToken(exTime int, uid, name string) (string, error) {
	uc := UserClaim{
		Uid:  uid,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(exTime)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(utils.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyzeToken token解析
func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(t *jwt.Token) (interface{}, error) {
		return []byte(utils.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token已过期")
	}
	return uc, err
}

// GetUUID 生成唯一ID
func GetUUID() string {
	return uuid.New().String()
}

// FileUploadQiniuSdk SDK方式上传文件
func FileUploadQiniuSdk(r *http.Request) (string, string, error) {
	file, fileHeader, err := r.FormFile("file")
	//os.Create(file)
	byteFile := make([]byte, fileHeader.Size)
	_, err = file.Read(byteFile)
	if err != nil {
		logx.Error("读取出错:" + err.Error())
		return "", "", err
	}
	pwd, _ := os.Getwd()
	fileBuff, _ := os.OpenFile(pwd+"/files_template/"+fileHeader.Filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer fileBuff.Close()
	defer os.Remove(pwd + "/files_template/" + fileHeader.Filename)

	_, err = fileBuff.Write(byteFile)
	if err != nil {
		//fmt.Println("存入文件错误" + err.Error())
		logx.Error("保存出错:"+err.Error()+",pwd:"+pwd, ",打开文件:"+pwd+"/files_template/"+fileHeader.Filename)
		return "", "", err
	}
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		logx.Error("NewFileRecorder出错:" + err.Error())
		return "", "", err
	}
	putExtra := storage.RputV2Extra{Recorder: recorder}
	key := utils.CosKey + GetUUID() + path.Ext(fileHeader.Filename)
	err = resumeUploader.PutFile(context.Background(), &ret, UpToken, key, "./files_template/"+fileHeader.Filename, &putExtra)
	if err != nil {
		logx.Error("PutFile出错:" + err.Error())
		return "", "", err
	}
	return ret.Key, ret.Hash, nil
}

// RandCode 验证码随机数生成
func RandCode() string {
	s := "0123456789abcdefghijklmnopqrstuvwxyz"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < utils.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// MailSendCode 验证码发送到邮箱
func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "Grj <531505871@qq.com>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("验证码:<h1>" + code + "</h1>")
	// 密码是邮箱生成的授权码
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "531505871@qq.com", utils.EmailPwd, "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}
	return nil
}
