package employee

import (
	"encoding/json"
	"errors"
	"github.com/Abdumalik92/moy_sklad_api/internal/middlware"
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func UpdateEmployee(request models.CreateEmployee, response *models.Row) error {
	url := "https://online.moysklad.ru/api/remap/1.2/entity/employee/" + request.ID
	method := "PUT"
	b, err := json.Marshal(request)
	if err != nil {
		log.Println("UpdateEmployee marshal json err ", err.Error())
		return err
	}
	client := &http.Client{}
	payload := strings.NewReader(string(b))

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println("UpdateEmployee req err ", err.Error())
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+middlware.HashAuth())

	res, err := client.Do(req)
	if err != nil {
		log.Println("UpdateEmployee res err ", err.Error())
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("UpdateEmployee body err ", err.Error())
		return err
	}
	if res.StatusCode != 200 {
		log.Println("UpdateEmployee body res ", string(body))
		return errors.New(string(body))
	}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("UpdateEmployee unmarshal body err ", err.Error())
		return err
	}
	return nil
}
