package params_service

import (
	"gfast/app/model/admin/sys_config"
)

//保存参数
func AddSave(req *sys_config.AddReq, userId int) (id int64, err error) {
	return sys_config.AddSave(req, userId)
}

//修改保存
func EditSave(req *sys_config.EditReq, userId int) (int64, error) {
	return sys_config.EditSave(req, userId)
}

//验证参数键名是否存在
func CheckConfigKeyUniqueAll(configKey string) error {
	return sys_config.CheckConfigKeyUniqueAll(configKey)
}

//列表分页搜索查询
func SelectListByPage(req *sys_config.SelectPageReq) (total, page int, list []*sys_config.Entity, err error) {
	return sys_config.SelectListByPage(req)
}

//通过id查询参数数据
func GetParamsById(id int) (entity *sys_config.Entity, err error) {
	return sys_config.GetParamsById(id)
}

//检查键是否已经存在
func CheckConfigKeyUnique(configKey string, configId int64) error {
	return sys_config.CheckConfigKeyUnique(configKey, configId)
}

//删除参数
func DeleteByIds(ids []int) error {
	return sys_config.DeleteByIds(ids)
}