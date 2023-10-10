package dao

import (
	"github.com/gin/blog-service/internal/model"
	"github.com/gin/blog-service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:      name,
		State:     state,
		BaseModel: &model.BaseModel{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		BaseModel: &model.BaseModel{Id: id},
	}
	//利用此处真正的传参更新数据库，而不是走shouldbind的传参，直接利用param里面的数据，
	//因为对于里面没有传参的数据，他会置空或者置零，此时依旧传进去会导致数据库的数据不是期望的，
	//gorm也不会对空数据进行修改，在拼装SQL上就拦截了，所以对于为空的数据如果传过来需要验证
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{BaseModel: &model.BaseModel{Id: id}}
	return tag.Delete(d.engine)
}
