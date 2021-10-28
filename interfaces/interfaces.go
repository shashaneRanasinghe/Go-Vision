package interfaces

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/shashaneRanasinghe/Go-Vision/entities"
	"github.com/shashaneRanasinghe/Go-Vision/responses"
)

type AWSService interface {
	DetectLabels(dlInput entities.DetectLabelInput) (*rekognition.DetectLabelsOutput, error)
}

type ModelService interface {
	ImageLabel(image []byte) responses.ClassifyResponse
}
