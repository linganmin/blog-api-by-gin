package models

type Articles struct {
	Model
	Title string `json:"title"`
	TagId int `json:"tag_id"`
	Status int `json:"status"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

func GetArticles(pageOffset int,pageSize int,maps interface{})(articles []Articles)  {
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&articles)
	return
}

func GetArticlesTotal(maps interface{})(count int)  {
	db.Model(&Articles{}).Where(maps).Count(&count)
	return
}

func AddArticle(title string, status int, desc string, content string, tagId int, createdBy string) bool {
	db.Model(Articles{}).Create(&Articles{
		Title:title,
		Status:status,
		Content:content,
		TagId:tagId,
		Desc:desc,
		CreatedBy:createdBy,
	})
	return  true
}