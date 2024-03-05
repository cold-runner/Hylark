package pkg

import (
	"bytes"
	"encoding/base64"
	"io"
	"mime/multipart"
	"os"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gabriel-vasile/mimetype"
)

// FileTypeMap TODO 新建类型处理
var FileTypeMap = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
}

func MustFileType(data []byte) (fileType string) {
	mime, _ := mimetype.DetectReader(bytes.NewReader(data))

	return FileTypeMap[mime.String()]
}

func FileTypeFromBinary(data []byte) (fileType string, err error) {
	mime, err := mimetype.DetectReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	return FileTypeMap[mime.String()], nil
}

func FileType(file io.Reader) (fileType string, err error) {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}

	return FileTypeMap[mime.String()], nil
}

func FileTypeFromBs64(bs64 string) (fileType string, err error) {
	data, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		return "", err
	}
	mime := mimetype.Detect(data)
	return FileTypeMap[mime.String()], nil
}

func FormFileType(fileHeader *multipart.FileHeader) (fileType string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	mime, err := mimetype.DetectReader(file)

	return FileTypeMap[mime.String()], nil
}

func IsFileImage(file *os.File) (bool, error) {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false, err
	}
	if !mime.Is("image/jpeg") && !mime.Is("image/png") {
		return false, nil
	}
	return true, nil
}

func IsFormFileImage(fileHeader *multipart.FileHeader) bool {
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false
	}
	if !mime.Is(consts.MIMEImageJPEG) && !mime.Is(consts.MIMEImagePNG) {
		return false
	}
	return true
}
