package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shashaneRanasinghe/Go-Vision/interfaces"
	"github.com/shashaneRanasinghe/Go-Vision/models"
	"github.com/shashaneRanasinghe/Go-Vision/services"
)

var awsService interfaces.AWSService = services.NewAWSService()
var model interfaces.ModelService = models.NewModelService(awsService)

//the classify method gets an image and returns the labels for
//the given image
func Classify(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v", err)
		_, _ = fmt.Fprint(w, err)
		return
	}
	response := model.ImageLabel(data)

	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v", err)
		_, _ = fmt.Fprint(w, err)
		return
	}
	w.Write(res)

}
