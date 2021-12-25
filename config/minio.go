package config

import (
	"fmt"
	"github.com/minio/minio-go/v6"
)

var minioClient *minio.Client

func MinioInit() {
	//endpoint := Config.GetString("minio_endpoint")
	//accessKeyID := Config.GetString("minio_accessKeyID")
	//secretAccessKey := Config.GetString("minio_secretAccessKey")
	//useSSL := Config.GetBool("minio_useSSL")
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true
	var err error
	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", minioClient)

	//// 创建一个叫mymusic的存储桶。
	//bucketName := "nkytest"
	//location := "us-east-1"
	//
	//err = minioClient.MakeBucket(bucketName, location)
	//if err != nil {
	//	// 检查存储桶是否已经存在。
	//	exists, err := minioClient.BucketExists(bucketName)
	//	if err == nil && exists {
	//		log.Printf("We already own %s\n", bucketName)
	//	} else {
	//		log.Fatalln("bucket exists",err)
	//	}
	//}
	//log.Printf("Successfully created %s\n", bucketName)
	//
	//// 上传一个zip文件。
	//objectName := "1.jpg"
	//filePath := "D:/1.jpg"
	//// contentType := "application/jpg"
	//
	//// 使用FPutObject上传一个zip文件。
	//n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{})
	//if err != nil {
	//	log.Fatalln("fputobject err ",err)
	//}
	//
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

func GetClient() *minio.Client{
	return minioClient
}
