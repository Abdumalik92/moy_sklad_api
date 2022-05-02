package employee

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/utils"

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
	strHash := utils.AppSettings.UserParams.Login + ":" + utils.AppSettings.UserParams.Password
	bs64 := base64.StdEncoding.EncodeToString([]byte(strHash))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+bs64)

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
	}
	str = auth.AccessToken
	return str, nil
}
