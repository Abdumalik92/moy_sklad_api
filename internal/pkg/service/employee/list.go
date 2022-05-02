package employee

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/utils"
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
	strHash := utils.AppSettings.UserParams.Login + ":" + utils.AppSettings.UserParams.Password
	bs64 := base64.StdEncoding.EncodeToString([]byte(strHash))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+bs64)

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

	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("GetEmployeeList unmarshal body err ", err.Error())
		return err
	}
	return nil
}
