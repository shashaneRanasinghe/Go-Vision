package models

import (
	"log"
	"os"
	"strconv"

	"github.com/shashaneRanasinghe/Go-Vision/entities"
	"github.com/shashaneRanasinghe/Go-Vision/interfaces"
	"github.com/shashaneRanasinghe/Go-Vision/responses"
)

type Model struct {
}

var aws interfaces.AWSService

func NewModelService(service interfaces.AWSService) interfaces.ModelService {
	aws = service
	return &Model{}
}

//the image label creates the parameters needed to call the detectLabel method
// in the aws service
func (m *Model) ImageLabel(image []byte) responses.ClassifyResponse {
	var labelList []string
	response := responses.ClassifyResponse{}

	maxLabels, err1 := strconv.ParseInt(os.Getenv("MAX_LABELS"), 10, 64)
	if err1 != nil {
		log.Printf("%v", err1)
		response.Error = "Unable to process the request"
		return response

	}
	minConfidence, err2 := strconv.ParseFloat(os.Getenv("MIN_CONFIDENCE"), 64)
	if err2 != nil {
		log.Printf("%v", err2)
		response.Error = "Unable to process the request"
		return response

	}

	detectLabelInput := entities.DetectLabelInput{
		Image:         image,
		MaxLabels:     maxLabels,
		MinConfidence: minConfidence,
	}

	results, err := aws.DetectLabels(detectLabelInput)

	if err != nil {
		response.Error = "Unable to process the request"
		return response
	}

	for _, item := range results.Labels {
		labelList = append(labelList, *item.Name)
	}
	response.Labels = labelList
	log.Printf("%v", response)
	return response
}
