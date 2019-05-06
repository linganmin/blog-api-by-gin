package v1

import (
	"ginhello/models"
	"ginhello/packages/e"
	"ginhello/packages/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tags(c *gin.Context) {

	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), util.GetPageSize(c), maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	var code int
	var msg string


	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state","0")).MustInt()
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name,100,"name").Message("名称最长为100字符")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy,100,"created_by").Message("创建人最长为100个字符")
	valid.Range(state,0,1,"state").Message("状态值错误")


	if !valid.HasErrors(){
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			res  := models.AddTag(name,state,createdBy)
			if !res {
				code = e.ERROR
			}
		}else {
			code = e.EXIST
		}

		msg = e.GetMsg(code)
	}else {
		msg = com.ToStr(valid.Errors)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":make(map[string]string),
	})

}

func UpdateTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
