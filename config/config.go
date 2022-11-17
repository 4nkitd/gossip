package config

import (
	"fs/api/utils"
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()

var AWS struct {
	AccessKey    string
	SecretKey    string
	Region       string
	S3bucketName string
	S3bucketPath string
}

var TotalInstaces int

func GetConfig() {

	utils.ThorwIfError(err)

	AWS.AccessKey = os.Getenv("AWS_ACCESS_KEY")
	AWS.SecretKey = os.Getenv("AWS_SECRET")
	AWS.Region = os.Getenv("AWS_REGION")
	AWS.S3bucketName = os.Getenv("AWS_S3_BUCKET_NAME")
	AWS.S3bucketPath = os.Getenv("AWS_S3_BUCKET_PATH")

	TotalInstaces = utils.StringToInt(os.Getenv("TOTAL_INSTANCES"))

}

func ENV(key string) string {
	var err = godotenv.Load()
	utils.ThorwIfError(err)
	return os.Getenv(key)
}
