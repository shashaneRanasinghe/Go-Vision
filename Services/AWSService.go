package Services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type DetectLabelInput struct{
	Image []byte
	MaxLabels int64
	MinConfidence float64
}

func (d *DetectLabelInput) DetectLabels() (*rekognition.DetectLabelsOutput,error){

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewSharedCredentials(os.Getenv("AWS_CREDENTIALS_FILEPATH"),
		os.Getenv("AWS_PROFILE")),
	})
	
	if err != nil{
		log.Printf("%v",err)
		return nil,err
	}

	svc := rekognition.New(sess)

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
				Bytes: d.Image,
		},
		MaxLabels:     aws.Int64(100),
		MinConfidence: aws.Float64(75.000000),
	}


	results,err := svc.DetectLabels(input)
	if err != nil{
		log.Printf("%v",err)
		return nil,err
	}
	return results,nil
}
