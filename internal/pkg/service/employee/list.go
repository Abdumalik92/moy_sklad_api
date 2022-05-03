package employee

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Abdumalik92/moy_sklad_api/internal/middlware"
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetEmployeeList(request models.Request, response *models.Response) error {
	url := "https://online.moysklad.ru/api/remap/1.2/entity/employee?limit=" + fmt.Sprint(request.Limit) + "&offset=" + fmt.Sprint(request.Offset)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println("GetEmployeeList req err ", err.Error())
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+middlware.HashAuth())

	res, err := client.Do(req)
	if err != nil {
		log.Println("GetEmployeeList res err ", err.Error())
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("GetEmployeeList body err ", err.Error())
		return err
	}
	if res.StatusCode != 200 {
		log.Println("GetEmployeeList body res ", string(body))
		return errors.New(string(body))
	}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("GetEmployeeList unmarshal body err ", err.Error())
		return err
	}
	return nil
}
