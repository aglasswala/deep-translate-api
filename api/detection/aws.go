package detection

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewAWSSession() (*session.Session, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		return awsSession, err
	}

	return awsSession, nil
}

func UploadFileToS3(bucket string, key string, file io.Reader, awsSession *session.Session) error {
	uploader := s3manager.NewUploader(awsSession)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

func RekognizeObjects(awsSession *session.Session, bucket string, key string) (*rekognition.DetectLabelsOutput, error) {
	svc := rekognition.New(awsSession)
	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(key),
			},
		},
		MaxLabels:     aws.Int64(123),
		MinConfidence: aws.Float64(70.000000),
	}

	result, err := svc.DetectLabels(input)
	if err != nil {
		return result, err
	}

	return result, nil
}

func DetectObjects(file io.Reader, filename string) (*rekognition.DetectLabelsOutput, error) {
	// Create a new AWS Session
	awsSession, err := NewAWSSession()
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}

	// Upload to file to S3
	bucket := os.Getenv("AWS_S3_BUCKET")
	err = UploadFileToS3(bucket, filename, file, awsSession)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}

	// Detect the objects
	result, err := RekognizeObjects(awsSession, bucket, filename)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}

	return result, nil
}
