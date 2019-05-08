package v1

import (
	"blog-api-by-gin/models"
	"blog-api-by-gin/packages/e"
	"blog-api-by-gin/packages/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Articles(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	title := c.Query("title")
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	if param := c.Query("status"); param != "" {
		status := com.StrTo(param).MustInt()
		maps["status"] = status
	}
	if title != "" {
		maps["title"] = title
	}
	if tagId > 0 {
		maps["tag_id"] = tagId
	}

	data["list"] = models.GetArticles(util.GetPageOffset(c), util.GetPageSize(c), maps)
	data["total"] = models.GetArticlesTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	var code int
	var msg string

	title := c.PostForm("title")
	status := com.StrTo(c.PostForm("status")).MustInt()
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 50, "title").Message("标题最长为50个字符")
	valid.Range(status, 0, 1, "status").Message("状态值错误")
	valid.Required(desc, "desc").Message("描述字段不能为空")
	valid.MaxSize(desc, 50, "desc").Message("描述字段最多50个字符")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Min(tagId, 1, "tag_id").Message("标签id必须大于0")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		msg = com.ToStr(valid.Errors)
	}else {
		models.AddArticle(title,status,desc,content,tagId,createdBy)
		code = e.SUCCESS
		msg = e.GetMsg(code)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":make(map[string]string),
	})

	// todo 验证是否已经存在






}
