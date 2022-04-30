package utils

import (
	"errors"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/google/uuid"

	httpErr "github.com/iamaul/go-evonix-backend-api/pkg/errors"
)

var allowedImagesContentType = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
}

func determineFileContentType(fileHeader textproto.MIMEHeader) (string, error) {
	contentTypes := fileHeader["Content-Type"]
	if len(contentTypes) < 1 {
		return "", httpErr.ErrNotAllowedImageHeader
	}
	return contentTypes[0], nil
}

func CheckImageContentType(image *multipart.FileHeader) error {
	// Check content type from header
	if !IsAllowedImageHeader(image) {
		return httpErr.ErrNotAllowedImageHeader
	}

	// Check real content type
	imageFile, err := image.Open()
	if err != nil {
		return httpErr.ErrBadRequest
	}
	defer imageFile.Close()

	fileHeader := make([]byte, 512)
	if _, err = imageFile.Read(fileHeader); err != nil {
		return httpErr.ErrBadRequest
	}

	if !IsAllowedImageContentType(fileHeader) {
		return httpErr.ErrNotAllowedImageHeader
	}
	return nil
}

func IsAllowedImageHeader(image *multipart.FileHeader) bool {
	contentType, err := determineFileContentType(image.Header)
	if err != nil {
		return false
	}
	_, allowed := allowedImagesContentType[contentType]
	return allowed
}

func GetImageExtension(image *multipart.FileHeader) (string, error) {
	contentType, err := determineFileContentType(image.Header)
	if err != nil {
		return "", err
	}

	extension, has := allowedImagesContentType[contentType]
	if !has {
		return "", errors.New("prohibited image extension")
	}
	return extension, nil
}

func GetImageContentType(image []byte) (string, bool) {
	contentType := http.DetectContentType(image)
	extension, allowed := allowedImagesContentType[contentType]
	return extension, allowed
}

func IsAllowedImageContentType(image []byte) bool {
	_, allowed := GetImageContentType(image)
	return allowed
}

func GetUniqFileName(customerId string, fileExtension string) string {
	randString := uuid.New().String()
	return "customerId_" + customerId + "_" + randString + "." + fileExtension
}
