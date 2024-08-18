package utils

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func GenerateUid() string {
	return uuid.New().String()
}

func UploadToCloudinary(file *multipart.FileHeader) (string, string, error) {
	defer func() {
		os.Remove("assets/uploads/" + file.Filename)
	}()

	var cloudinary_url = os.Getenv("CLOUDINARY_URL")

	ctx := context.Background()
	cld, err := cloudinary.NewFromURL(cloudinary_url)
	if err != nil {
		fmt.Println(err)
		return "", "", errors.New("error with cloudinary url")
	}

	resp, err := cld.Upload.Upload(ctx, "assets/uploads/"+file.Filename, uploader.UploadParams{PublicID: "image" + "-" + file.Filename + "-" + GenerateUid()})
	if err != nil {
		fmt.Println(err)
		return "", "", errors.New("can't upload images")
	}

	return resp.PublicID, resp.SecureURL, nil

}

func DeleteFromCloudinary(imageId string) error {
	var cloudinary_url = os.Getenv("CLOUDINARY_URL")

	cld, err := cloudinary.NewFromURL(cloudinary_url)
	if err != nil {
		return errors.New("error with cloudinary url")
	}

	resp, err := cld.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: imageId})
	if err != nil || resp.Result != "ok" {
		return errors.New("error deleting deleting images from cloudinary")
	}

	return nil
}
