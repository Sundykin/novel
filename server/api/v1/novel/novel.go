package novel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NovelApi struct{}

var novelService = service.ServiceGroupApp.NovelServiceGroup.NovelService

// CreateNovel 创建小说
// @Tags Novel
// @Summary 创建小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body novel.Novel true "小说信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /novel/createNovel [post]
func (n *NovelApi) CreateNovel(c *gin.Context) {
	var novelInfo novel.Novel
	err := c.ShouldBindJSON(&novelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 获取当前用户ID
	userID := utils.GetUserID(c)
	novelInfo.CreatedBy = userID
	
	err = novelService.CreateNovel(&novelInfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteNovel 删除小说
// @Tags Novel
// @Summary 删除小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body novel.Novel true "小说ID"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /novel/deleteNovel [delete]
func (n *NovelApi) DeleteNovel(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := novelService.DeleteNovel(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteNovelByIds 批量删除小说
// @Tags Novel
// @Summary 批量删除小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /novel/deleteNovelByIds [delete]
func (n *NovelApi) DeleteNovelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := novelService.DeleteNovelByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateNovel 更新小说
// @Tags Novel
// @Summary 更新小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body novel.Novel true "小说信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /novel/updateNovel [put]
func (n *NovelApi) UpdateNovel(c *gin.Context) {
	var novelInfo novel.Novel
	err := c.ShouldBindJSON(&novelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	userID := utils.GetUserID(c)
	err = novelService.UpdateNovel(novelInfo, userID)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindNovel 用id查询小说
// @Tags Novel
// @Summary 用id查询小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query novel.Novel true "用id查询小说"
// @Success 200 {object} response.Response{data=novel.Novel,msg=string} "查询成功"
// @Router /novel/findNovel [get]
func (n *NovelApi) FindNovel(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	reNovel, err := novelService.GetNovel(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(gin.H{"novel": reNovel}, c)
}

// GetNovelList 分页获取小说列表
// @Tags Novel
// @Summary 分页获取小说列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.NovelSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /novel/getNovelList [get]
func (n *NovelApi) GetNovelList(c *gin.Context) {
	var pageInfo request.NovelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	userID := utils.GetUserID(c)
	pageInfo.CreatedBy = userID
	
	list, total, err := novelService.GetNovelInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetNovelPublic 获取公开小说列表
// @Tags Novel
// @Summary 获取公开小说列表
// @accept application/json
// @Produce application/json
// @Param data query request.NovelSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /novel/getNovelPublic [get]
func (n *NovelApi) GetNovelPublic(c *gin.Context) {
	var pageInfo request.NovelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	list, total, err := novelService.GetNovelPublicList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// UpdateNovelStatus 更新小说状态
// @Tags Novel
// @Summary 更新小说状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /novel/updateNovelStatus [put]
func (n *NovelApi) UpdateNovelStatus(c *gin.Context) {
	ID := c.Query("ID")
	status := c.Query("status")
	userID := utils.GetUserID(c)
	
	err := novelService.UpdateNovelStatus(ID, status, userID)
	if err != nil {
		global.GVA_LOG.Error("更新状态失败!", zap.Error(err))
		response.FailWithMessage("更新状态失败", c)
		return
	}
	response.OkWithMessage("更新状态成功", c)
}

// GetNovelStatistics 获取小说统计信息
// @Tags Novel
// @Summary 获取小说统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /novel/getNovelStatistics [get]
func (n *NovelApi) GetNovelStatistics(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	
	statistics, err := novelService.GetNovelStatistics(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("获取统计信息失败!", zap.Error(err))
		response.FailWithMessage("获取统计信息失败", c)
		return
	}
	response.OkWithData(statistics, c)
}