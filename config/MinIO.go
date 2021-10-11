package config

import (
	"github.com/minio/minio-go/v6"
)

var MinIO *minio.Client

type MinIOConf struct {
	StorageDB 		string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName  	string
}

func initMinIO() {
	conf := &MinIOConf{
		// StorageDB:       Config.GetString("storageEngine"),
		Endpoint:        Config.GetString("MINIO_END_POINT"),
		AccessKeyID:     Config.GetString("MINIO_ACCESS_KEY"),
		SecretAccessKey: Config.GetString("MINIO_SECRET_KEY"),
		UseSSL:          Config.GetBool("MINIO_USE_SSL"),
		// BucketName:      Config.GetString("BUCKET_NAME"),
	}

	minioClient, err := minio.New(conf.Endpoint, conf.AccessKeyID, conf.SecretAccessKey, conf.UseSSL)
	if err != nil{
	}
	MinIO = minioClient

}

