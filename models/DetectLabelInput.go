package Models

import (
	"log"
	"os"
	"strconv"

	"github.com/shashaneRanasinghe/Go-Vision/Responses"
	"github.com/shashaneRanasinghe/Go-Vision/Services"
)
type Image struct{
	Image []byte
}


func (i *Image) ImageLabel() (responses.ClassifyResponse) {
	var labelList []string
	response :=  responses.ClassifyResponse{}

	maxLabels,err1 := strconv.ParseInt(os.Getenv("MAX_LABELS"),10,64)
	if err1 != nil{
		log.Printf("%v",err1)
		response.Error = "Unable to process the request"
		return response

	}
	minConfidence,err2 := strconv.ParseFloat(os.Getenv("MIN_CONFIDENCE"),64)
	if err2 != nil{
		log.Printf("%v",err2)
		response.Error = "Unable to process the request"
		return response

	}

	detectLabelInput := Services.DetectLabelInput{
		Image: i.Image,
		MaxLabels: maxLabels,
		MinConfidence: minConfidence,
	}

	results,err := detectLabelInput.DetectLabels()
	
	if err != nil {
		response.Error = "Unable to process the request"
		return response
	}


	for _,item := range results.Labels {
		labelList = append(labelList, *item.Name)
	}
	response.Labels = labelList

	return response
}