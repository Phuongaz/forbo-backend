package common

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const endpoint = "db:9000"
const accessKey = "minioadmin"
const secretKey = "minioadmin"

func ConnectMinIO() (c *minio.Client, err error) {
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient, err
}

func SetPermission(client *minio.Client, bucketName string) error {
	policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]`
	err := client.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return err
}

func CreateBucket(client *minio.Client, bucketName string, objectName string, data io.Reader) error {
	_, err := client.GetBucketPolicy(context.Background(), bucketName)

	if err != nil {
		log.Fatal(err)
	}

	n, err := client.PutObject(context.Background(), bucketName, objectName, data, -1, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n.Size)

	return err
}

func UploadFile(client *minio.Client, bucketName string, objectName string, file *multipart.FileHeader) error {
	fileData, err := file.Open()
	if err != nil {
		log.Fatalln(err)
	}

	err = CreateBucket(client, bucketName, objectName, fileData)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func DownloadFile(client *minio.Client, bucketName string, objectName string) (string, error) {

	downloadPath := "/tmp/" + objectName
	file, err := os.Create(downloadPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	object, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}
	defer object.Close()

	_, err = io.Copy(file, object)
	if err != nil {
		return "", err
	}

	return downloadPath, nil
}

func GetObject(client *minio.Client, bucketName string, objectName string) (io.Reader, error) {
	object, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	return object, err
}
