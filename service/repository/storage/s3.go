package storage

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/zerolog/log"
)

type S3Storage struct {
	Region     string
	PrivateKey string
	PublicKey  string
	Bucket     string
}

func InitS3Storage() IStorage {
	s3Storage := &S3Storage{
		Region:     utstring.GetEnv(models.AWSS3Region, "ap-southeast-1"),
		PrivateKey: utstring.GetEnv(models.AWSS3PrivateKey, "C31SFH80ZuxRdkJTvRERShTsreQPBygjeFH6aitr"),
		PublicKey:  utstring.GetEnv(models.AWSS3PublicKey, "AKIA3I6OUCO7J76KIVYP"),
		Bucket:     utstring.GetEnv(models.AWSS3Bucket, "jobhunapp"),
	}
	return s3Storage
}

func (s *S3Storage) AddCloudFileAndGetURL(localFilePath string) (urlLink string, err error) {
	const ACL = "public-read"
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s.Region),
		Credentials: credentials.NewStaticCredentials(s.PublicKey, s.PrivateKey, "")},
	)

	if err != nil {
		log.Error().Err(err)
		return
	}
	// Open the file for use
	file, err := os.Open(localFilePath)
	if err != nil {
		log.Error().Err(err)
		return
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, err := file.Stat()
	if err != nil {
		log.Error().Err(err)
		return
	}

	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(localFilePath),
		ACL:    aws.String(ACL),
		Body:   bytes.NewReader(buffer),
	})

	if err != nil {
		log.Error().Err(err)
		return
	}

	urlLink = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s",
		s.Bucket,
		s.Region,
		localFilePath,
	)

	return
}
