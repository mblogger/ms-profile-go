package handleProfile

import (
	"github.com/gin-gonic/gin"

	utils "biodata/lib"
	templates "biodata/lib/html"
	model "biodata/model"
	homeTemplate "biodata/templates"
)

func Home(c *gin.Context) {
	var data model.HomeData
	data.Title = "MS - Profile"
	data.Message = "Welcome, you can see my profile by visiting /admin/profile link!"
	templates.Read("home", homeTemplate.TPL, data, c)
}

func GetProfileById(id string) string {
	return utils.GetLoginData(id)
}
