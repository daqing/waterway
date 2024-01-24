package ext_setup_api

import (
	"github.com/daqing/waterway/lib/api_resp"
	"github.com/daqing/waterway/lib/repo"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type CreateParams struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

func CreateAction(c *gin.Context) {
	var p CreateParams

	if err := c.BindJSON(&p); err != nil {
		api_resp.LogError(c, err)
		return
	}

	if len(p.Type) == 0 {
		api_resp.ErrorCodeMsg(c, 400, "Type missing")
		return
	}

	if len(p.URL) == 0 {
		api_resp.ErrorCodeMsg(c, 400, "Url missing")
		return
	}

	viper.SetConfigFile("./config/db.yaml")
	viper.Set("db_type", p.Type)
	viper.Set("db_url", p.URL)
	if err := viper.WriteConfig(); err != nil {
		api_resp.LogError(c, err)
		return
	}

	if !repo.Setup() {
		api_resp.ErrorCodeMsg(c, 500, "repo setup failed")
		return
	}

	api_resp.OK(c, "ok")
}
