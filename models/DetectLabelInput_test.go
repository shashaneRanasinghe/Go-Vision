package Models

import (
	"bufio"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	// "github.com/shashaneRanasinghe/Go-Vision/responses"
	mock_interfaces "github.com/shashaneRanasinghe/Go-Vision/mocks"
)

// func TestImageLabel1(t *testing.T){
	
// 	testImageFile,err := os.Open("../testImages/images.jpg")
// 	if err != nil{
// 		log.Printf("Couldnt open image %v",err)
// 	}
// 	reader := bufio.NewReader(testImageFile)
// 	content,err1 := ioutil.ReadAll(reader)
// 	if err1 != nil{
// 		log.Printf("%v",err1)
// 	}

// 	labelArray := []string{"tree","beach","sea"}

// 	resp := responses.ClassifyResponse{
// 		Labels: labelArray,
// 		Error: "",
// 	}

// 	img := Image{
// 		Image: content,
// 	}

// 	var tests = []struct{
// 		image Image
// 		expected responses.ClassifyResponse
// 	}{
// 		{img,resp},
// 	}

// 	for _,test := range tests{
// 		testName := fmt.Sprintf("%v",testImageFile)
// 		t.Run(testName,func(t *testing.T) {
// 			err := godotenv.Load("../.env")
// 			if err != nil{
// 				t.Logf("env failed %v",err)
// 				t.Fail()
// 			}
// 			actual := test.image.ImageLabel()
// 			fmt.Print(actual)
// 			if actual.Error != test.expected.Error{
// 				t.Errorf("got %v, want %v",actual.Error,test.expected.Error)
// 			}
// 		})
// 	}
// }

func TestImageLabel (t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var testFloat float64 = 10

	boundingBox := rekognition.BoundingBox{
		Height: &testFloat,
		Left: &testFloat,
		Top: &testFloat,
		Width: &testFloat,
	}
	recognitionInstance := rekognition.Instance{
		BoundingBox:&boundingBox,
		Confidence:&testFloat,
	}

	labelArray := []*rekognition.Instance{&recognitionInstance}

	rLabel := rekognition.Label{
		Confidence: nil,
		Instances: labelArray,
		Name: nil,
		Parents: nil,
	}

	labels := []*rekognition.Label{&rLabel}

	detectLabelsOutput := rekognition.DetectLabelsOutput{
		LabelModelVersion:nil,
		Labels:labels,
		OrientationCorrection:nil,
	}

	mockAWSService := mock_interfaces.NewMockAWSService(ctrl)
	mockAWSService.EXPECT().DetectLabels().Return(&detectLabelsOutput,nil)

	testImageFile,err := os.Open("../testImages/images.jpg")
	if err != nil{
		log.Printf("Couldnt open image %v",err)
	}

	reader := bufio.NewReader(testImageFile)
	content,err1 := ioutil.ReadAll(reader)
	if err1 != nil{
		log.Printf("%v",err1)
	}

	_ = Image{
		Image: content,
	}

	_ = godotenv.Load("../.env")

	// expectedErrors := ""

	// actual := img.ImageLabel()
	// if actual.Error != expectedErrors{
	// 	t.Fail()
	// }

}