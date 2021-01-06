// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"io"
	"path"
)

const serviceBasePath = "drive/v1"

type Service struct {
	conf     *config.Config
	basePath string
	Medias   *MediaService
	Files    *FileService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf:     conf,
		basePath: serviceBasePath,
	}
	s.Medias = newMediaService(s)
	s.Files = newFileService(s)
	return s
}

type MediaService struct {
	service *Service
}

func newMediaService(service *Service) *MediaService {
	return &MediaService{
		service: service,
	}
}

type FileService struct {
	service *Service
}

func newFileService(service *Service) *FileService {
	return &FileService{
		service: service,
	}
}

type FileUploadPrepareReqCall struct {
	ctx   *core.Context
	files *FileService
	body  *UploadInfo

	optFns []request.OptFn
}

func (rc *FileUploadPrepareReqCall) Do() (*FileUploadPrepareResult, error) {
	httpPath := path.Join(rc.files.service.basePath, "files/upload_prepare")
	var result = &FileUploadPrepareResult{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) UploadPrepare(ctx *core.Context, body *UploadInfo, optFns ...request.OptFn) *FileUploadPrepareReqCall {
	return &FileUploadPrepareReqCall{
		ctx:    ctx,
		files:  files,
		body:   body,
		optFns: optFns,
	}
}

type MediaUploadPrepareReqCall struct {
	ctx    *core.Context
	medias *MediaService
	body   *UploadInfo

	optFns []request.OptFn
}

func (rc *MediaUploadPrepareReqCall) Do() (*MediaUploadPrepareResult, error) {
	httpPath := path.Join(rc.medias.service.basePath, "medias/upload_prepare")
	var result = &MediaUploadPrepareResult{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.medias.service.conf, req)
	return result, err
}

func (medias *MediaService) UploadPrepare(ctx *core.Context, body *UploadInfo, optFns ...request.OptFn) *MediaUploadPrepareReqCall {
	return &MediaUploadPrepareReqCall{
		ctx:    ctx,
		medias: medias,
		body:   body,
		optFns: optFns,
	}
}

type MediaUploadPartReqCall struct {
	ctx    *core.Context
	medias *MediaService
	body   *request.FormData

	optFns []request.OptFn
}

func (rc *MediaUploadPartReqCall) SetUploadId(uploadId string) {
	rc.body.AddParam("upload_id", uploadId)
}
func (rc *MediaUploadPartReqCall) SetSeq(seq int) {
	rc.body.AddParam("seq", seq)
}
func (rc *MediaUploadPartReqCall) SetSize(size int) {
	rc.body.AddParam("size", size)
}
func (rc *MediaUploadPartReqCall) SetChecksum(checksum string) {
	rc.body.AddParam("checksum", checksum)
}
func (rc *MediaUploadPartReqCall) SetFile(file *request.File) {
	rc.body.AddFile("file", file)
}
func (rc *MediaUploadPartReqCall) Do() (*response.NoData, error) {
	httpPath := path.Join(rc.medias.service.basePath, "medias/upload_part")
	var result = &response.NoData{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.medias.service.conf, req)
	return result, err
}

func (medias *MediaService) UploadPart(ctx *core.Context, optFns ...request.OptFn) *MediaUploadPartReqCall {
	return &MediaUploadPartReqCall{
		ctx:    ctx,
		medias: medias,
		body:   request.NewFormData(),
		optFns: optFns,
	}
}

type FileUploadPartReqCall struct {
	ctx   *core.Context
	files *FileService
	body  *request.FormData

	optFns []request.OptFn
}

func (rc *FileUploadPartReqCall) SetUploadId(uploadId string) {
	rc.body.AddParam("upload_id", uploadId)
}
func (rc *FileUploadPartReqCall) SetSeq(seq int) {
	rc.body.AddParam("seq", seq)
}
func (rc *FileUploadPartReqCall) SetSize(size int) {
	rc.body.AddParam("size", size)
}
func (rc *FileUploadPartReqCall) SetChecksum(checksum string) {
	rc.body.AddParam("checksum", checksum)
}
func (rc *FileUploadPartReqCall) SetFile(file *request.File) {
	rc.body.AddFile("file", file)
}
func (rc *FileUploadPartReqCall) Do() (*response.NoData, error) {
	httpPath := path.Join(rc.files.service.basePath, "files/upload_part")
	var result = &response.NoData{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) UploadPart(ctx *core.Context, optFns ...request.OptFn) *FileUploadPartReqCall {
	return &FileUploadPartReqCall{
		ctx:    ctx,
		files:  files,
		body:   request.NewFormData(),
		optFns: optFns,
	}
}

type MediaUploadFinishReqCall struct {
	ctx    *core.Context
	medias *MediaService
	body   *MediaUploadFinishReqBody

	optFns []request.OptFn
}

func (rc *MediaUploadFinishReqCall) Do() (*MediaUploadFinishResult, error) {
	httpPath := path.Join(rc.medias.service.basePath, "medias/upload_finish")
	var result = &MediaUploadFinishResult{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.medias.service.conf, req)
	return result, err
}

func (medias *MediaService) UploadFinish(ctx *core.Context, body *MediaUploadFinishReqBody, optFns ...request.OptFn) *MediaUploadFinishReqCall {
	return &MediaUploadFinishReqCall{
		ctx:    ctx,
		medias: medias,
		body:   body,
		optFns: optFns,
	}
}

type FileUploadFinishReqCall struct {
	ctx   *core.Context
	files *FileService
	body  *FileUploadFinishReqBody

	optFns []request.OptFn
}

func (rc *FileUploadFinishReqCall) Do() (*FileUploadFinishResult, error) {
	httpPath := path.Join(rc.files.service.basePath, "files/upload_finish")
	var result = &FileUploadFinishResult{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) UploadFinish(ctx *core.Context, body *FileUploadFinishReqBody, optFns ...request.OptFn) *FileUploadFinishReqCall {
	return &FileUploadFinishReqCall{
		ctx:    ctx,
		files:  files,
		body:   body,
		optFns: optFns,
	}
}

type FileUploadAllReqCall struct {
	ctx   *core.Context
	files *FileService
	body  *request.FormData

	optFns []request.OptFn
}

func (rc *FileUploadAllReqCall) SetFileName(fileName string) {
	rc.body.AddParam("file_name", fileName)
}
func (rc *FileUploadAllReqCall) SetParentType(parentType string) {
	rc.body.AddParam("parent_type", parentType)
}
func (rc *FileUploadAllReqCall) SetParentNode(parentNode string) {
	rc.body.AddParam("parent_node", parentNode)
}
func (rc *FileUploadAllReqCall) SetSize(size int) {
	rc.body.AddParam("size", size)
}
func (rc *FileUploadAllReqCall) SetChecksum(checksum string) {
	rc.body.AddParam("checksum", checksum)
}
func (rc *FileUploadAllReqCall) SetFile(file *request.File) {
	rc.body.AddFile("file", file)
}
func (rc *FileUploadAllReqCall) Do() (*FileUploadAllResult, error) {
	httpPath := path.Join(rc.files.service.basePath, "files/upload_all")
	var result = &FileUploadAllResult{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) UploadAll(ctx *core.Context, optFns ...request.OptFn) *FileUploadAllReqCall {
	return &FileUploadAllReqCall{
		ctx:    ctx,
		files:  files,
		body:   request.NewFormData(),
		optFns: optFns,
	}
}

type MediaUploadAllReqCall struct {
	ctx    *core.Context
	medias *MediaService
	body   *request.FormData

	optFns []request.OptFn
}

func (rc *MediaUploadAllReqCall) SetFileName(fileName string) {
	rc.body.AddParam("file_name", fileName)
}
func (rc *MediaUploadAllReqCall) SetParentType(parentType string) {
	rc.body.AddParam("parent_type", parentType)
}
func (rc *MediaUploadAllReqCall) SetParentNode(parentNode string) {
	rc.body.AddParam("parent_node", parentNode)
}
func (rc *MediaUploadAllReqCall) SetSize(size int) {
	rc.body.AddParam("size", size)
}
func (rc *MediaUploadAllReqCall) SetChecksum(checksum string) {
	rc.body.AddParam("checksum", checksum)
}
func (rc *MediaUploadAllReqCall) SetFile(file *request.File) {
	rc.body.AddFile("file", file)
}
func (rc *MediaUploadAllReqCall) Do() (*MediaUploadAllResult, error) {
	httpPath := path.Join(rc.medias.service.basePath, "medias/upload_all")
	var result = &MediaUploadAllResult{}
	req := request.NewRequest(httpPath, "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.medias.service.conf, req)
	return result, err
}

func (medias *MediaService) UploadAll(ctx *core.Context, optFns ...request.OptFn) *MediaUploadAllReqCall {
	return &MediaUploadAllReqCall{
		ctx:    ctx,
		medias: medias,
		body:   request.NewFormData(),
		optFns: optFns,
	}
}

type MediaDownloadReqCall struct {
	ctx        *core.Context
	medias     *MediaService
	pathParams map[string]interface{}

	optFns []request.OptFn
	result io.Writer
}

func (rc *MediaDownloadReqCall) SetFileToken(fileToken string) {
	rc.pathParams["file_token"] = fileToken
}
func (rc *MediaDownloadReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}
func (rc *MediaDownloadReqCall) Do() (io.Writer, error) {
	httpPath := path.Join(rc.medias.service.basePath, "medias/:file_token/download")
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest(httpPath, "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.medias.service.conf, req)
	return rc.result, err
}

func (medias *MediaService) Download(ctx *core.Context, optFns ...request.OptFn) *MediaDownloadReqCall {
	return &MediaDownloadReqCall{
		ctx:        ctx,
		medias:     medias,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type FileDownloadReqCall struct {
	ctx        *core.Context
	files      *FileService
	pathParams map[string]interface{}

	optFns []request.OptFn
	result io.Writer
}

func (rc *FileDownloadReqCall) SetFileToken(fileToken string) {
	rc.pathParams["file_token"] = fileToken
}
func (rc *FileDownloadReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}
func (rc *FileDownloadReqCall) Do() (io.Writer, error) {
	httpPath := path.Join(rc.files.service.basePath, "files/:file_token/download")
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest(httpPath, "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return rc.result, err
}

func (files *FileService) Download(ctx *core.Context, optFns ...request.OptFn) *FileDownloadReqCall {
	return &FileDownloadReqCall{
		ctx:        ctx,
		files:      files,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type MediaBatchGetTmpDownloadUrlReqCall struct {
	ctx    *core.Context
	medias *MediaService

	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *MediaBatchGetTmpDownloadUrlReqCall) SetFileTokens(fileTokens ...string) {
	rc.queryParams["file_tokens"] = fileTokens
}
func (rc *MediaBatchGetTmpDownloadUrlReqCall) Do() (*MediaBatchGetTmpDownloadUrlResult, error) {
	httpPath := path.Join(rc.medias.service.basePath, "medias/batch_get_tmp_download_url")
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &MediaBatchGetTmpDownloadUrlResult{}
	req := request.NewRequest(httpPath, "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.medias.service.conf, req)
	return result, err
}

func (medias *MediaService) BatchGetTmpDownloadUrl(ctx *core.Context, optFns ...request.OptFn) *MediaBatchGetTmpDownloadUrlReqCall {
	return &MediaBatchGetTmpDownloadUrlReqCall{
		ctx:         ctx,
		medias:      medias,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}
