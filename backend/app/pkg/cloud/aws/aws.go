package aws

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_REGION string = "ap-southeast-1"
)

type AWSService interface {
	UploadFileS3(filePath, fileName string, folder string) (objectURL string, err error)
	UploadExcelToS3(buffer *bytes.Buffer, folder, fileName string) (objectURL string, err error)
	GetFile(key string) ([]byte, int64, error)
	DeleteFileFromS3(key string) error
}

type AWSServiceImpl struct {
	session *session.Session
}

func AWSServiceInit() *AWSServiceImpl {
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")

	awsSession := NewClient(accessKey, secretKey)
	return &AWSServiceImpl{
		session: awsSession,
	}
}

func (s *AWSServiceImpl) UploadFileS3(filePath, fileName string, folder string) (objectURL string, err error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	defer os.Remove(filePath)

	bucket := os.Getenv("AWS_S3_BUCKET")

	_, err = s3.New(s.session).PutObject(
		&s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(folder + "/" + fileName),
			ContentType: aws.String(http.DetectContentType(data)),
			Body:        bytes.NewReader(data),
			ACL:         aws.String("public-read"),
		})
	objectURL = "https://%s.s3.%s.amazonaws.com/%s/%s"
	objectURL = fmt.Sprintf(objectURL, bucket, AWS_REGION, folder, fileName)
	return objectURL, err
}

func (s *AWSServiceImpl) UploadExcelToS3(buffer *bytes.Buffer, folder, fileName string) (objectURL string, err error) {
	uploader := s3manager.NewUploader(s.session)
	bucket := os.Getenv("AWS_S3_BUCKET")
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(folder + "/" + fileName),
		ContentType: aws.String("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"),
		Body:        buffer,
		ACL:         aws.String("public-read"),
	})

	// Return any error from the upload
	objectURL = "https://%s.s3.%s.amazonaws.com/%s/%s"
	objectURL = fmt.Sprintf(objectURL, bucket, AWS_REGION, folder, fileName)
	return objectURL, err
}

func (s *AWSServiceImpl) GetFile(key string) ([]byte, int64, error) {
	bucket := os.Getenv("AWS_S3_BUCKET")
	req, err := s3.New(s.session).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, 0, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Failed to read object body: %s", err)
	}
	fileSize := req.ContentLength
	return body, *fileSize, err

}

func (s *AWSServiceImpl) DeleteFileFromS3(key string) error {
	bucket := os.Getenv("AWS_S3_BUCKET")

	svc := s3.New(s.session)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := svc.DeleteObject(input)
	if err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to wait for object deletion: %v", err)
	}

	return nil
}
