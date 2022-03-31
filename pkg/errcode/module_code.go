package errcode

var (
	ErrorGetTagListFail = NewError(2001_0001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(2001_0002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(2001_0003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(2001_0004, "删除标签失败")
	ErrorCountTagFail   = NewError(2001_005, "统计标签失败")
)
