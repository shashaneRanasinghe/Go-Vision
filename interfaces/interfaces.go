package interfaces

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/shashaneRanasinghe/Go-Vision/responses"
)

type AWSService interface{
	DetectLabels () (*rekognition.DetectLabelsOutput,error)
}

type ModelService interface{
	ImageLabel () (responses.ClassifyResponse)
}