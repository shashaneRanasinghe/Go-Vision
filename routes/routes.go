package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shashaneRanasinghe/Go-Vision/models"
)

func Classify(w http.ResponseWriter, r *http.Request) {
	data,err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v",err)
		_, _ = fmt.Fprint(w, err)
		return
	}

	image := Models.Image{
		Image: data,
	}

	response := image.ImageLabel()
	w.WriteHeader(http.StatusOK)
	res,err := json.Marshal(response)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v",err)
		_, _ = fmt.Fprint(w, err)
		return
	}
	w.Write(res)

}