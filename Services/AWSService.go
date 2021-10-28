package services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/shashaneRanasinghe/Go-Vision/entities"
	"github.com/shashaneRanasinghe/Go-Vision/interfaces"
)

type AWSservice struct{}

func NewAWSService() interfaces.AWSService {
	return &AWSservice{}
}

// the DetectLabels function calls the AWS API and gets the labels corresponding to the
// image given
func (a *AWSservice) DetectLabels(dlInput entities.DetectLabelInput) (*rekognition.DetectLabelsOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewSharedCredentials(os.Getenv("AWS_CREDENTIALS_FILEPATH"),
			os.Getenv("AWS_PROFILE")),
	})

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	svc := rekognition.New(sess)

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: dlInput.Image,
		},
		MaxLabels:     aws.Int64(dlInput.MaxLabels),
		MinConfidence: aws.Float64(dlInput.MinConfidence),
	}

	results, err := svc.DetectLabels(input)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return results, nil
}
