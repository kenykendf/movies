package service

import (
	"context"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type UploadService struct {
	CloudName string
	APIKey    string
	APISecret string
	Directory string
}

func NewUploadService(CloudName, APIKey, APISecret, Directory string) *UploadService {
	return &UploadService{
		CloudName: CloudName,
		APIKey:    APIKey,
		APISecret: APISecret,
		Directory: Directory,
	}
}

func (us *UploadService) UploadImage(input *multipart.FileHeader) (string, error) {

	// validate Image mime and extention
	err := validateImage(input)
	if err != nil {
		return "", err
	}

	cld, _ := cloudinary.NewFromParams(us.CloudName, us.APIKey, us.APISecret)

	ctx := context.Background()
	file, _ := input.Open()

	uploaded, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{
			Folder:   us.CloudName,
			PublicID: input.Filename,
		},
	)
	if err != nil {
		return "", err
	}

	return uploaded.SecureURL, nil
}

const (
	imageMime = "image/jpeg, image/jpg, image/png"
	imageExt  = "jpeg, jpg, png"
	imageSize = 5000000 // size in KiloByte
)

func validateImage(image *multipart.FileHeader) error {
	image.Open()

	mime := image.Header.Get("Content-Type")
	if !checkMime(mime) {
		return errors.New("invalid image mime type")
	}

	filename := image.Filename
	if !checkExt(filename) {
		return errors.New("invalid image extention, only allow : " + imageExt)
	}

	if image.Size > imageSize {
		return errors.New("image's size is too big")
	}

	return nil
}

func checkMime(mimeType string) bool {
	mime := strings.Split(imageMime, ", ")

	for _, allowedMime := range mime {
		if mimeType == allowedMime {
			return true
		}
	}
	return false
}

func checkExt(filename string) bool {
	file := strings.Split(filename, ".")
	fileExt := file[len(file)-1]

	ext := strings.Split(imageExt, ", ")

	for _, allowedExt := range ext {
		if fileExt == allowedExt {
			return true
		}
	}
	return false
}
