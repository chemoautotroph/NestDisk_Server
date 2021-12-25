package commands

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"myServer/config"
	"strings"
)

func Upload(addr, username string) (error, string) {
	err := createBucket(username)
	if err != nil{
		fmt.Println("createBucket err",err)
		return err, ""
	}
	objectName, _ := expressAddr(addr)
	fmt.Printf("object name: %v, addr: %v\n", objectName, addr)
	minioClient := config.GetClient()
	n, err := minioClient.FPutObject(username, objectName, addr, minio.PutObjectOptions{})
	if err != nil{
		fmt.Println("FPutObject error", err)
		return err, ""
	}
	fmt.Printf("Successfully uploaded %s of size %d\n\n", objectName, n)
	message := fmt.Sprintf("Successfully uploaded %s of size %d\n", objectName, n)
	return nil, message
}

func createBucket(username string) error{
	fmt.Println(username)
	bucketName := username
	minioClient := config.GetClient()
	fmt.Println(minioClient)
	exists, err := minioClient.BucketExists(bucketName)
	if err != nil{
		fmt.Println("bucketExists ",err)
		return err
	} else if exists{
		fmt.Printf("%v Bucket already exists\n", bucketName)
		return nil
	} else {
		err := minioClient.MakeBucket(bucketName, "")
		if err != nil{
			fmt.Println("makeBucket ",err)
			return err
		}
		fmt.Printf("Successfully create bucket %v\n", bucketName)
		return nil
	}
}

func expressAddr(addr string) (string,string){
	splitAddr := strings.Split(addr, "/")
	objectName := splitAddr[len(splitAddr)-1]
	size := len(objectName)+1
	filePath := addr[:size-1]
	return objectName, filePath
}