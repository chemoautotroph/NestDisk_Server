package commands

import (
	"fmt"
	"myServer/config"
)

func Show(username, myPrefix string) ([]string, error){
	fmt.Println("commands.show")
	doneCh := make(chan struct{})
	defer close(doneCh)
	minioClient := config.GetClient()

	var allObject []string

	for objectInfo := range minioClient.ListObjects(username, myPrefix, true, doneCh) {
		if objectInfo.Err != nil {
			fmt.Println(objectInfo.Err)
			return nil, objectInfo.Err
		}
		objName := objectInfo.Key
		// lastModifiied := objectInfo.LastModified

		allObject = append(allObject, objName)
	}

	return allObject, nil
}
