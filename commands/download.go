package commands

import (
	"github.com/minio/minio-go/v6"
	"myServer/config"
)

func Download(bucketName, objectName, filePath string) error {
	minioClient := config.GetClient()
	err := minioClient.FGetObject(bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
