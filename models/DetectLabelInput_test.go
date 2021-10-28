package models

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"

	"github.com/shashaneRanasinghe/Go-Vision/entities"
	mock_interfaces "github.com/shashaneRanasinghe/Go-Vision/mocks"
)

func TestImageLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var testFloat float64 = 10

	testImageFile, err := os.Open("../testImages/images.jpg")
	if err != nil {
		log.Printf("Couldnt open image %v", err)
	}

	reader := bufio.NewReader(testImageFile)
	img, err1 := ioutil.ReadAll(reader)
	if err1 != nil {
		log.Printf("%v", err1)
	}

	_ = godotenv.Load("../.env")

	detectLabelsInput := entities.DetectLabelInput{
		Image:         img,
		MaxLabels:     10,
		MinConfidence: 75.0,
	}

	boundingBox := rekognition.BoundingBox{
		Height: &testFloat,
		Left:   &testFloat,
		Top:    &testFloat,
		Width:  &testFloat,
	}
	recognitionInstance := rekognition.Instance{
		BoundingBox: &boundingBox,
		Confidence:  &testFloat,
	}

	labelArray := []*rekognition.Instance{&recognitionInstance}

	rLabel := rekognition.Label{
		Confidence: nil,
		Instances:  labelArray,
		Name:       nil,
		Parents:    nil,
	}

	labels := []*rekognition.Label{&rLabel}

	detectLabelsOutput := rekognition.DetectLabelsOutput{
		LabelModelVersion:     nil,
		Labels:                labels,
		OrientationCorrection: nil,
	}

	mockAWSService := mock_interfaces.NewMockAWSService(ctrl)

	awsService := mockAWSService
	model := NewModelService(awsService)

	actual := model.ImageLabel(img)

	mockAWSService.EXPECT().DetectLabels(detectLabelsInput).Return(&detectLabelsOutput, nil)

	expectedErrors := ""

	if actual.Error != expectedErrors {
		t.Fail()
	}

}
