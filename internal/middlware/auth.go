package middlware

import (
	"encoding/base64"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/utils"
)

func HashAuth() string {
	strHash := utils.AppSettings.UserParams.Login + ":" + utils.AppSettings.UserParams.Password
	return base64.StdEncoding.EncodeToString([]byte(strHash))
}
