package errcode

var (
	ErrorGetTagListFail = NewError(2001_0001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(2001_0002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(2001_0003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(2001_0004, "删除标签失败")
	ErrorCountTagFail   = NewError(2001_0005, "统计标签失败")

	ErrorGetArticleFail    = NewError(2002_0001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(2002_0002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(2002_0003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(2002_0004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(2002_0005, "删除文章失败")
)
