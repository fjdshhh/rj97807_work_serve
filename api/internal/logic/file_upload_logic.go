package logic

import (
	"context"
	"rj97807_work_serve/api/models"
	"rj97807_work_serve/funcs"
	"rj97807_work_serve/utils"

	"rj97807_work_serve/api/internal/svc"
	"rj97807_work_serve/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest, userUid string) (resp *types.FileUploadResponse, err error) {
	response := &models.RepositoryModel{
		Uid:       funcs.GetUUID(),
		BelongUid: userUid,
		Md5:       req.Md5,
		Name:      req.Name,
		Ext:       req.Ext,
		Size:      req.Size,
		Path:      utils.CosBucket + req.Path,
	}
	err = l.svcCtx.EngineWeb.Table("repository").Create(&response).Error
	if err != nil {
		return
	}
	resp = new(types.FileUploadResponse)
	resp.Uid = response.Uid
	resp.Name = response.Name
	resp.Ext = response.Ext
	return
}
