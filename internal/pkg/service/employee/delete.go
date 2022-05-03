package employee

import (
	"errors"
	"github.com/Abdumalik92/moy_sklad_api/internal/middlware"
	"io/ioutil"
	"log"
	"net/http"
)

func DeleteEmployee(id string) error {
	url := "https://online.moysklad.ru/api/remap/1.2/entity/employee/" + id
	method := "DELETE"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("DeleteEmployee req err ", err.Error())
		return err
	}

	req.Header.Add("Authorization", "Basic "+middlware.HashAuth())

	res, err := client.Do(req)
	if err != nil {
		log.Println("DeleteEmployee res err ", err.Error())
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("DeleteEmployee body err ", err.Error())
		return err
	}
	if res.StatusCode != 200 {
		log.Println("DeleteEmployee body res ", string(body))
		return errors.New(string(body))
	}
	return nil
}
