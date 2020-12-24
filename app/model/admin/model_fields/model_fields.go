// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package model_fields

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// AddReq 用于存储新增请求的请求参数
type AddReq struct {
	ModelId      uint     `p:"modelId" `
	FieldName    string   `p:"fieldName" v:"required#字段名称不能为空"`
	FieldTitle   string   `p:"fieldTitle" `
	FieldType    string   `p:"fieldType" `
	FieldLength  string   `p:"fieldLength" `
	FieldDefault string   `p:"fieldDefault" `
	FieldData    string   `p:"fieldData" `
	FieldDesc    string   `p:"fieldDesc" `
	FieldRules   []string `p:"fieldRules" `
	CreateBy     uint64
	FieldWidth   string `p:"fieldWidth"`
	FieldAlign   string `p:"fieldAlign"`
}

// EditReq 用于存储修改请求参数
type EditReq struct {
	FieldId      int64    `p:"fieldId" v:"required#主键ID不能为空"`
	ModelId      uint     `p:"modelId" `
	FieldName    string   `p:"fieldName" v:"required#字段名称不能为空"`
	FieldTitle   string   `p:"fieldTitle" `
	FieldType    string   `p:"fieldType" `
	FieldLength  string   `p:"fieldLength" `
	FieldDefault string   `p:"fieldDefault" `
	FieldData    string   `p:"fieldData" `
	FieldDesc    string   `p:"fieldDesc" `
	FieldRules   []string `p:"fieldRules" `
	UpdateBy     uint64
	FieldWidth   string `p:"fieldWidth"`
	FieldAlign   string `p:"fieldAlign"`
}

type RemoveReq struct {
	Ids []int `p:"ids"` //删除id
}

type FieldInfo struct {
	FieldId      uint64 `p:"field_id,primary" json:"field_id"`      // 模型字段ID
	ModelId      uint   `p:"model_id"         json:"model_id"`      // 模型ID
	FieldName    string `p:"field_name"       json:"field_name"`    // 字段名称
	FieldTitle   string `p:"field_title"      json:"field_title"`   // 字段标题
	FieldType    string `p:"field_type"       json:"field_type"`    // 字段类型
	FieldLength  string `p:"field_length"     json:"field_length"`  // 字段长度
	FieldDefault string `p:"field_default"    json:"field_default"` // 字段默认值
	FieldData    string `p:"field_data"       json:"field_data"`    // 字段数据
	FieldDesc    string `p:"field_desc"       json:"field_desc"`    // 字段描述
	FieldRules   string `p:"field_rules"      json:"field_rules"`   // 字段规则
	FieldSort    int64  `p:"field_sort"       json:"field_sort"`    // 字段排序
	CreateBy     uint64 `p:"create_by"        json:"create_by"`     // 创建人
	UpdateBy     uint64 `p:"update_by"        json:"update_by"`     // 修改人
	CreateTime   uint64 `p:"create_time"      json:"create_time"`   // 创建时间
	UpdateTime   uint64 `p:"update_time"      json:"update_time"`   // 修改时间
	FieldWidth   string `p:"field_width"      json:"field_width"`   //字段列表宽度
	FieldAlign   string `p:"field_align"     json:"field_align"`    //字段对齐方式
	ModelEdit    string `p:"model_edit" json:"model_edit"`
	ModelIndexes string `p:"model_indexes" json:"model_indexes"`
	ModelList    string `p:"model_list" json:"model_list"`
	ModelOrder   string `p:"model_order" json:"model_order"`
	ModelPk      string `p:"model_pk" json:"model_pk"`
	ModelSort    string `p:"model_sort" json:"model_sort"`
	SearchList   string `p:"search_list" json:"search_list"`
}

type SetFieldsAttrReq struct {
	ModelId    int    `p:"modelId" json:"model_id"`
	PkId       uint64 `p:"pkId" json:"pk_id"`
	FieldsList []FieldInfo
}

// SelectPageReq 用于存储分页查询的请求参数
type SelectPageReq struct {
	ModelId   int64  `p:"modelId"`   //模型ID
	BeginTime string `p:"beginTime"` //开始时间
	EndTime   string `p:"endTime"`   //结束时间
	PageNum   int64  `p:"pageNum"`   //当前页码
	PageSize  int    `p:"pageSize"`  //每页数
}

// 根据ID查询记录
func GetByID(id int64) (*Entity, error) {
	entity, err := Model.FindOne(id)
	if err != nil {
		g.Log().Error(err)
		return nil, gerror.New("根据ID查询记录出错")
	}
	if entity == nil {
		return nil, gerror.New("没有查询到对应模型的字段")
	}
	return entity, nil
}

// AddSave 添加
func AddSave(req *AddReq) error {
	if err := Exists(req.FieldName, req.ModelId, 0); err != nil {
		return err
	}
	entity := new(Entity)
	entity.ModelId = req.ModelId
	entity.FieldName = req.FieldName
	entity.FieldTitle = req.FieldTitle
	entity.FieldType = req.FieldType
	entity.FieldLength = req.FieldLength
	entity.FieldDefault = req.FieldDefault
	entity.FieldData = req.FieldData
	entity.FieldDesc = req.FieldDesc
	entity.FieldRules = gstr.Join(req.FieldRules, ",")
	entity.CreateBy = req.CreateBy
	entity.FieldSort = 1000
	entity.FieldWidth = req.FieldWidth
	entity.FieldAlign = req.FieldAlign
	time := gconv.Uint64(gtime.Timestamp())
	entity.CreateTime = time
	entity.UpdateTime = time
	result, err := entity.Insert()
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

//判断字段是否存在
func Exists(name string, modelId, fieldsId uint) error {
	model := Model.Where(Columns.FieldName, name).And(Columns.ModelId, modelId)
	if fieldsId != 0 {
		model = model.And(Columns.FieldId, fieldsId)
	}
	entity, err := model.FindOne()
	if err != nil {
		g.Log().Error(err)
		return gerror.New("判断数据重复时出错")
	}
	if entity != nil {
		return gerror.New("已存在相同名称的字段")
	}
	return nil
}

// 删除
func DeleteByIds(Ids []int) error {
	_, err := Model.Delete("field_id in(?)", Ids)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除失败")
	}
	return nil
}

// 根据ID来修改信息
func EditSave(req *EditReq) error {
	// 先根据ID来查询要修改的记录
	entity, err := GetByID(req.FieldId)
	if err != nil {
		return err
	}

	// 修改实体
	entity.ModelId = req.ModelId
	entity.FieldName = req.FieldName
	entity.FieldTitle = req.FieldTitle
	entity.FieldType = req.FieldType
	entity.FieldLength = req.FieldLength
	entity.FieldDefault = req.FieldDefault
	entity.FieldData = req.FieldData
	entity.FieldDesc = req.FieldDesc
	entity.FieldRules = gstr.Join(req.FieldRules, ",")
	entity.UpdateBy = req.UpdateBy
	entity.UpdateTime = gconv.Uint64(gtime.Timestamp())
	entity.FieldWidth = req.FieldWidth
	entity.FieldAlign = req.FieldAlign
	_, err = Model.Save(entity)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("修改失败")
	}
	return nil
}

// 获取所有数据
func SelectListAll(req *SelectPageReq) (list []*Entity, err error) {
	model := Model
	if req != nil {
		if req.ModelId != 0 {
			model = model.Where(Columns.ModelId, req.ModelId)
		}
	}
	// 查询
	list, err = model.Order("field_sort asc,field_id asc").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("查询失败")
		return
	}
	return
}

//通过模型ID删除对应模型字段
func DeleteByModelIds(modelIds []int, tx *gdb.TX) error {
	_, err := Model.TX(tx).Delete(Columns.ModelId+" in(?)", modelIds)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除字段信息失败")
	}
	return nil
}
