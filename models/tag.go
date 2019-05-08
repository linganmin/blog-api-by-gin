package models

type Tags struct {
	Model
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	Status    int    `json:"status"`
}

// 获取tag列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tags) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// 获取tag 总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tags{}).Where(maps).Count(&count)
	return
}

// 增加tag
func AddTag(name string, status int, createdBy string) bool {
	db.Create(&Tags{
		Name:      name,
		Status:     status,
		CreatedBy: createdBy,
	})

	return true
}

// 通过标签名字判断是否已存在
func ExistTagByName(name string) bool {
	var tag Tags

	db.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

// 通过id判断是否已存在
func ExistTagById(id int) bool {
	var tag Tags
	db.Select("id").Where("id = ? ", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// 更新标签
func UpdateTag(id int, params interface{}) bool {
	db.Model(&Tags{}).Where("id = ?", id).Updates(params)
	return true
}

// 删除标签
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tags{})
	return true
}
