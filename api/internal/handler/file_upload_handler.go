package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"path"
	"rj97807_work_serve/api/models"
	"rj97807_work_serve/funcs"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rj97807_work_serve/api/internal/logic"
	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		//文件获取-前端使用form/data的contentType
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		//判断文件是否已经存在
		Tob := make([]byte, fileHeader.Size)
		_, err = file.Read(Tob)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(Tob))
		response := new(models.RepositoryModel)
		svcCtx.EngineWeb.Table("repository").Where("md5=?", hash).Find(&response)
		// 如果文件存在
		if response.Uid != "" {
			response.BelongUid = r.Header.Get("Uid")
			err = svcCtx.EngineWeb.Table("repository").Create(&response).Error
			if err != nil {
				httpx.Error(w, errors.New("文件已存在，但发生错误:"+err.Error()))
				return
			}
			httpx.OkJson(w, &types.FileUploadResponse{
				Uid:  response.Uid,
				Ext:  response.Ext,
				Name: response.Name,
			})
			return
		}
		//如果文件不存在--方式有两种，参考ReadMe
		//cosPath, cosMd5, err := funcs.FileUploadQiniuSdk(r)
		cosPath, _, err := funcs.FileUploadQiniuSdk(r)
		if err != nil {
			logx.Error("initUploadFile:" + err.Error())
			httpx.Error(w, err)
			return
		}
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		//req.Md5 = cosMd5
		//自己的md5与七牛云的md5加密方式不同，导致文件秒传无法生效，所以这个地方还是保存我们自己的md5
		req.Md5 = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req, r.Header.Get("Uid"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
