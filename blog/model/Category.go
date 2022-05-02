package model

import (
	"blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	// 感觉有点多余字段，mysql默认会创建id字段
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	//Cid  int    `gorm:"type:int;not null" json:"cid"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	gorm.Model
}

// 对数据库的操作DAO

// 查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //用户已存在
	} else {
		return errmsg.SUCCESS // 用户不存在，用户名可用
	}
}

// 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 查询分类列表
// 返回Category类型的切片
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cates []Category
	var total int64
	// 分页
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// TODO:查询该分类下的所有文章

// 删除分类(软删除)
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}