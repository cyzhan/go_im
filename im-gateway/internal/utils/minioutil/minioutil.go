package minioutil

import (
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint        = os.Getenv("MINIO_ENDPOINT")
	accessKeyID     = os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("MINIO_SECRET_ACCESS_KEY")
	secure          bool
	client          *minio.Client
)

func InitMinio() {
	var err error
	secure, err = strconv.ParseBool(os.Getenv("MINIO_SECURE"))
	if err != nil {
		panic(err)
	}

	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: secure,
	})

	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}

func GetClient() *minio.Client {
	return client
}

func Upload(ctx *gin.Context, bucketName string, objectPath string, file *multipart.FileHeader, contentType string) (path string, err error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}

	src, err := file.Open()
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}
	defer src.Close()

	info, err := minioClient.PutObject(ctx, bucketName, objectPath, src, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Printf(err.Error())
		return "", err
	}
	log.Printf("info.Bucket = %s", info.Bucket)

	if secure {
		path = "https://"
	} else {
		path = "http://"
	}
	path = path + endpoint + "/" + bucketName + "/" + objectPath
	return path, nil
}
