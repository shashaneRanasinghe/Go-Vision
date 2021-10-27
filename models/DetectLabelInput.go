package Models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shashaneRanasinghe/Go-Vision/responses"
	"github.com/shashaneRanasinghe/Go-Vision/services"
)

type Image struct {
	Image []byte
}

//the image label creates the parameters needed to call the detectLabel method
// in the aws service
func (i *Image) ImageLabel() responses.ClassifyResponse {
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

	detectLabelInput := services.DetectLabelInput{
		Image:         i.Image,
		MaxLabels:     maxLabels,
		MinConfidence: minConfidence,
	}
	fmt.Println("Testing.......")

	results, err := detectLabelInput.DetectLabels()


	if err != nil {
		response.Error = "Unable to process the request"
		return response
	}

	for _, item := range results.Labels {
		labelList = append(labelList, *item.Name)
	}
	response.Labels = labelList

	return response
}
