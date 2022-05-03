package employee

import (
	"encoding/json"
	"github.com/Abdumalik92/moy_sklad_api/internal/middlware"
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"io/ioutil"
	"log"
	"net/http"
)

func Auth() (str string, err error) {
	var auth models.Auth

	url := "https://online.moysklad.ru/api/remap/1.2/security/token"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println("Auth req err ", err.Error())
		return str, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+middlware.HashAuth())

	res, err := client.Do(req)
	if err != nil {
		log.Println("Auth res err ", err.Error())
		return str, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Auth body err ", err.Error())
		return str, err
	}

	if err := json.Unmarshal(body, &auth); err != nil {
		log.Println("Auth unmarshal body err ", err.Error())
		return str, err
	}
	str = auth.AccessToken
	return str, nil
}
