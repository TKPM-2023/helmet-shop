package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/config"
	"log"
	"net/http"
)

type S3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(
	config *config.Config,
) *S3Provider {
	provider := &S3Provider{
		bucketName: config.S3BucketName,
		region:     config.S3Region,
		apiKey:     config.S3ApiKey,
		secret:     config.S3SecretKey,
		domain:     config.S3Domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secret, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider
}

func (provider *S3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	fmt.Println("fileType====>", fileType)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, err
}
