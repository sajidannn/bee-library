package utils

import (
	"bee-library/config"
	"context"
	"log"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func DeleteFromCloudinary(imageURL string) error {
	cld, err := config.SetupCloudinary()
	if err != nil {
		return err
	}

	publicID := ExtractPublicID(imageURL)

	_, err = cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		log.Println("Error deleting file from Cloudinary:", err)
		return err
	}

	return nil
}

func ExtractPublicID(imageURL string) string {
	parts := strings.Split(imageURL, "/")
	filename := parts[len(parts)-1]
	publicID := strings.Split(filename, ".")[0]
	return publicID
}
