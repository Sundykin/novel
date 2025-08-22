package novel

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AIApi struct{}

var aiService = service.ServiceGroupApp.NovelServiceGroup.AIService

// ChatWithAI AI聊天对话
// @Tags AI
// @Summary AI聊天对话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIChatRequest true "AI聊天请求"
// @Success 200 {object} response.Response{data=novel.AIAssistant,msg=string} "聊天成功"
// @Router /ai/chat [post]
func (a *AIApi) ChatWithAI(c *gin.Context) {
	var req request.AIChatRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	assistant, err := aiService.ChatWithAI(&req)
	if err != nil {
		global.GVA_LOG.Error("AI聊天失败!", zap.Error(err))
		response.FailWithMessage("AI聊天失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(gin.H{"assistant": assistant}, "聊天成功", c)
}

// GenerateContent AI内容生成
// @Tags AI
// @Summary AI内容生成
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIGenerateRequest true "AI生成请求"
// @Success 200 {object} response.Response{data=novel.AIAssistant,msg=string} "生成成功"
// @Router /ai/generate [post]
func (a *AIApi) GenerateContent(c *gin.Context) {
	var req request.AIGenerateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	assistant, err := aiService.GenerateContent(&req)
	if err != nil {
		global.GVA_LOG.Error("AI内容生成失败!", zap.Error(err))
		response.FailWithMessage("AI内容生成失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(gin.H{"assistant": assistant}, "生成成功", c)
}

// GetChatHistory 获取聊天历史
// @Tags AI
// @Summary 获取聊天历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param sessionId query string true "会话ID"
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ai/history [get]
func (a *AIApi) GetChatHistory(c *gin.Context) {
	sessionID := c.Query("sessionId")
	if sessionID == "" {
		response.FailWithMessage("会话ID不能为空", c)
		return
	}
	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	
	history, total, err := aiService.GetChatHistory(sessionID, page, pageSize)
	if err != nil {
		global.GVA_LOG.Error("获取聊天历史失败!", zap.Error(err))
		response.FailWithMessage("获取聊天历史失败", c)
		return
	}
	
	response.OkWithDetailed(response.PageResult{
		List:     history,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, "获取成功", c)
}

// SubmitFeedback 提交反馈
// @Tags AI
// @Summary 提交反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIFeedbackRequest true "反馈请求"
// @Success 200 {object} response.Response{msg=string} "提交成功"
// @Router /ai/feedback [post]
func (a *AIApi) SubmitFeedback(c *gin.Context) {
	var req request.AIFeedbackRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	err = aiService.SubmitFeedback(req.AssistantID, req.UserRating, req.Feedback)
	if err != nil {
		global.GVA_LOG.Error("提交反馈失败!", zap.Error(err))
		response.FailWithMessage("提交反馈失败", c)
		return
	}
	
	response.OkWithMessage("提交成功", c)
}

// GetUsageStats 获取使用统计
// @Tags AI
// @Summary 获取使用统计
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIUsageStatsRequest true "统计请求"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取成功"
// @Router /ai/stats [post]
func (a *AIApi) GetUsageStats(c *gin.Context) {
	var req request.AIUsageStatsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	// 解析日期
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		response.FailWithMessage("开始日期格式错误", c)
		return
	}
	
	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		response.FailWithMessage("结束日期格式错误", c)
		return
	}
	
	stats, err := aiService.GetUsageStats(req.UserID, startDate, endDate)
	if err != nil {
		global.GVA_LOG.Error("获取使用统计失败!", zap.Error(err))
		response.FailWithMessage("获取使用统计失败", c)
		return
	}
	
	response.OkWithDetailed(stats, "获取成功", c)
}

// WritingAssist 写作助手
// @Tags AI
// @Summary 写作助手
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIWritingAssistRequest true "写作助手请求"
// @Success 200 {object} response.Response{data=novel.AIAssistant,msg=string} "辅助成功"
// @Router /ai/writing-assist [post]
func (a *AIApi) WritingAssist(c *gin.Context) {
	var req request.AIWritingAssistRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	// 构建AI请求
	chatReq := &request.AIChatRequest{
		SessionID:    "writing_" + strconv.Itoa(int(req.NovelID)),
		UserID:       req.UserID,
		NovelID:      req.NovelID,
		UserMessage:  buildWritingPrompt(req),
		MessageType:  novel.MessageTypeGenerate,
		FunctionType: novel.FunctionTypeWriting,
		Context:      req.Content,
	}
	
	assistant, err := aiService.ChatWithAI(chatReq)
	if err != nil {
		global.GVA_LOG.Error("写作助手失败!", zap.Error(err))
		response.FailWithMessage("写作助手失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(gin.H{"assistant": assistant}, "辅助成功", c)
}

// PlotGenerate 情节生成
// @Tags AI
// @Summary 情节生成
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIPlotGenerateRequest true "情节生成请求"
// @Success 200 {object} response.Response{data=novel.AIAssistant,msg=string} "生成成功"
// @Router /ai/plot-generate [post]
func (a *AIApi) PlotGenerate(c *gin.Context) {
	var req request.AIPlotGenerateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	// 构建AI请求
	chatReq := &request.AIChatRequest{
		SessionID:    "plot_" + strconv.Itoa(int(req.NovelID)),
		UserID:       req.UserID,
		NovelID:      req.NovelID,
		UserMessage:  buildPlotPrompt(req),
		MessageType:  novel.MessageTypeGenerate,
		FunctionType: novel.FunctionTypePlot,
	}
	
	assistant, err := aiService.ChatWithAI(chatReq)
	if err != nil {
		global.GVA_LOG.Error("情节生成失败!", zap.Error(err))
		response.FailWithMessage("情节生成失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(gin.H{"assistant": assistant}, "生成成功", c)
}

// CharacterGenerate 角色生成
// @Tags AI
// @Summary 角色生成
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AICharacterGenerateRequest true "角色生成请求"
// @Success 200 {object} response.Response{data=novel.AIAssistant,msg=string} "生成成功"
// @Router /ai/character-generate [post]
func (a *AIApi) CharacterGenerate(c *gin.Context) {
	var req request.AICharacterGenerateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	// 构建AI请求
	chatReq := &request.AIChatRequest{
		SessionID:    "character_" + strconv.Itoa(int(req.NovelID)),
		UserID:       req.UserID,
		NovelID:      req.NovelID,
		UserMessage:  buildCharacterPrompt(req),
		MessageType:  novel.MessageTypeGenerate,
		FunctionType: novel.FunctionTypeCharacter,
	}
	
	assistant, err := aiService.ChatWithAI(chatReq)
	if err != nil {
		global.GVA_LOG.Error("角色生成失败!", zap.Error(err))
		response.FailWithMessage("角色生成失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(gin.H{"assistant": assistant}, "生成成功", c)
}

// DialogueGenerate 对话生成
// @Tags AI
// @Summary 对话生成
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIDialogueGenerateRequest true "对话生成请求"
// @Success 200 {object} response.Response{data=novel.AIAssistant,msg=string} "生成成功"
// @Router /ai/dialogue-generate [post]
func (a *AIApi) DialogueGenerate(c *gin.Context) {
	var req request.AIDialogueGenerateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	// 构建AI请求
	chatReq := &request.AIChatRequest{
		SessionID:    "dialogue_" + strconv.Itoa(int(req.NovelID)),
		UserID:       req.UserID,
		NovelID:      req.NovelID,
		UserMessage:  buildDialoguePrompt(req),
		MessageType:  novel.MessageTypeGenerate,
		FunctionType: novel.FunctionTypeDialogue,
	}
	
	assistant, err := aiService.ChatWithAI(chatReq)
	if err != nil {
		global.GVA_LOG.Error("对话生成失败!", zap.Error(err))
		response.FailWithMessage("对话生成失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(gin.H{"assistant": assistant}, "生成成功", c)
}

// ContinueWriting 续写
// @Tags AI
// @Summary 续写
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIContinueRequest true "续写请求"
// @Success 200 {object} response.Response{data=novel.AIAssistant,msg=string} "续写成功"
// @Router /ai/continue [post]
func (a *AIApi) ContinueWriting(c *gin.Context) {
	var req request.AIContinueRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 设置用户ID
	req.UserID = utils.GetUserID(c)
	
	// 构建AI请求
	chatReq := &request.AIChatRequest{
		SessionID:    "continue_" + strconv.Itoa(int(req.NovelID)),
		UserID:       req.UserID,
		NovelID:      req.NovelID,
		UserMessage:  buildContinuePrompt(req),
		MessageType:  novel.MessageTypeGenerate,
		FunctionType: novel.FunctionTypeContinue,
		Context:      req.Content,
	}
	
	assistant, err := aiService.ChatWithAI(chatReq)
	if err != nil {
		global.GVA_LOG.Error("续写失败!", zap.Error(err))
		response.FailWithMessage("续写失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(gin.H{"assistant": assistant}, "续写成功", c)
}

// 构建写作助手提示词
func buildWritingPrompt(req request.AIWritingAssistRequest) string {
	prompt := "请作为专业的小说写作助手，根据以下要求提供帮助：\n\n"
	prompt += "辅助类型：" + req.AssistType + "\n"
	if req.Requirement != "" {
		prompt += "具体要求：" + req.Requirement + "\n"
	}
	if req.Length > 0 {
		prompt += "生成长度：" + strconv.Itoa(req.Length) + "字\n"
	}
	prompt += "\n当前内容：\n" + req.Content
	return prompt
}

// 构建情节生成提示词
func buildPlotPrompt(req request.AIPlotGenerateRequest) string {
	prompt := "请生成一个" + req.PlotType + "情节，要求如下：\n\n"
	if req.Characters != "" {
		prompt += "相关角色：" + req.Characters + "\n"
	}
	if req.Setting != "" {
		prompt += "场景设定：" + req.Setting + "\n"
	}
	if req.Conflict != "" {
		prompt += "冲突设定：" + req.Conflict + "\n"
	}
	if req.Tone != "" {
		prompt += "情感基调：" + req.Tone + "\n"
	}
	if req.Length > 0 {
		prompt += "生成长度：" + strconv.Itoa(req.Length) + "字\n"
	}
	if req.Requirement != "" {
		prompt += "具体要求：" + req.Requirement + "\n"
	}
	return prompt
}

// 构建角色生成提示词
func buildCharacterPrompt(req request.AICharacterGenerateRequest) string {
	prompt := "请创建一个" + req.CharacterType + "角色，要求如下：\n\n"
	if req.Role != "" {
		prompt += "角色定位：" + req.Role + "\n"
	}
	if req.Personality != "" {
		prompt += "性格特点：" + req.Personality + "\n"
	}
	if req.Background != "" {
		prompt += "背景设定：" + req.Background + "\n"
	}
	if req.Relationships != "" {
		prompt += "关系网络：" + req.Relationships + "\n"
	}
	if req.Requirement != "" {
		prompt += "具体要求：" + req.Requirement + "\n"
	}
	return prompt
}

// 构建对话生成提示词
func buildDialoguePrompt(req request.AIDialogueGenerateRequest) string {
	prompt := "请生成一段对话，要求如下：\n\n"
	prompt += "对话角色：" + req.Characters + "\n"
	prompt += "对话情境：" + req.Situation + "\n"
	if req.Purpose != "" {
		prompt += "对话目的：" + req.Purpose + "\n"
	}
	if req.Tone != "" {
		prompt += "对话基调：" + req.Tone + "\n"
	}
	if req.Length > 0 {
		prompt += "对话长度：" + strconv.Itoa(req.Length) + "字\n"
	}
	if req.Requirement != "" {
		prompt += "具体要求：" + req.Requirement + "\n"
	}
	return prompt
}

// 构建续写提示词
func buildContinuePrompt(req request.AIContinueRequest) string {
	prompt := "请根据以下内容进行续写：\n\n"
	if req.Direction != "" {
		prompt += "续写方向：" + req.Direction + "\n"
	}
	if req.Style != "" {
		prompt += "写作风格：" + req.Style + "\n"
	}
	if req.Length > 0 {
		prompt += "续写长度：" + strconv.Itoa(req.Length) + "字\n"
	}
	if req.Requirement != "" {
		prompt += "具体要求：" + req.Requirement + "\n"
	}
	prompt += "\n当前内容：\n" + req.Content
	return prompt
}

// CreateAIConfig 创建AI配置
// @Tags AI
// @Summary 创建AI配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body novel.AIConfig true "AI配置信息"
// @Success 200 {object} response.Response{data=novel.AIConfig,msg=string} "创建成功"
// @Router /novel/createAIConfig [post]
func (a *AIApi) CreateAIConfig(c *gin.Context) {
	var aiConfig novel.AIConfig
	err := c.ShouldBindJSON(&aiConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	err = aiService.CreateAIConfig(&aiConfig)
	if err != nil {
		global.GVA_LOG.Error("创建AI配置失败!", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(aiConfig, "创建成功", c)
}

// DeleteAIConfig 删除AI配置
// @Tags AI
// @Summary 删除AI配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body novel.AIConfig true "AI配置信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /novel/deleteAIConfig [delete]
func (a *AIApi) DeleteAIConfig(c *gin.Context) {
	var aiConfig novel.AIConfig
	err := c.ShouldBindJSON(&aiConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	err = aiService.DeleteAIConfig(&aiConfig)
	if err != nil {
		global.GVA_LOG.Error("删除AI配置失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}
	
	response.OkWithMessage("删除成功", c)
}

// DeleteAIConfigByIds 批量删除AI配置
// @Tags AI
// @Summary 批量删除AI配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AI配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /novel/deleteAIConfigByIds [delete]
func (a *AIApi) DeleteAIConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	err = aiService.DeleteAIConfigByIds(&IDS)
	if err != nil {
		global.GVA_LOG.Error("批量删除AI配置失败!", zap.Error(err))
		response.FailWithMessage("删除失败: "+err.Error(), c)
		return
	}
	
	response.OkWithMessage("删除成功", c)
}

// UpdateAIConfig 更新AI配置
// @Tags AI
// @Summary 更新AI配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body novel.AIConfig true "AI配置信息"
// @Success 200 {object} response.Response{data=novel.AIConfig,msg=string} "更新成功"
// @Router /novel/updateAIConfig [put]
func (a *AIApi) UpdateAIConfig(c *gin.Context) {
	var aiConfig novel.AIConfig
	err := c.ShouldBindJSON(&aiConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	err = aiService.UpdateAIConfig(&aiConfig)
	if err != nil {
		global.GVA_LOG.Error("更新AI配置失败!", zap.Error(err))
		response.FailWithMessage("更新失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(aiConfig, "更新成功", c)
}

// FindAIConfig 根据ID获取AI配置
// @Tags AI
// @Summary 根据ID获取AI配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param ID query string true "AI配置ID"
// @Success 200 {object} response.Response{data=novel.AIConfig,msg=string} "查询成功"
// @Router /novel/findAIConfig [get]
func (a *AIApi) FindAIConfig(c *gin.Context) {
	ID := c.Query("ID")
	aiConfig, err := aiService.GetAIConfig(ID)
	if err != nil {
		global.GVA_LOG.Error("查询AI配置失败!", zap.Error(err))
		response.FailWithMessage("查询失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(aiConfig, "查询成功", c)
}

// GetAIConfigList 分页获取AI配置列表
// @Tags AI
// @Summary 分页获取AI配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.AIConfigSearch true "分页获取AI配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /novel/getAIConfigList [get]
func (a *AIApi) GetAIConfigList(c *gin.Context) {
	var pageInfo request.AIConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	list, total, err := aiService.GetAIConfigInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取AI配置列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// TestAIConfig 测试AI配置
// @Tags AI
// @Summary 测试AI配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AIConfigTestRequest true "AI配置测试请求"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "测试成功"
// @Router /novel/testAIConfig [post]
func (a *AIApi) TestAIConfig(c *gin.Context) {
	var req request.AIConfigTestRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	result, err := aiService.TestAIConfig(&req)
	if err != nil {
		global.GVA_LOG.Error("测试AI配置失败!", zap.Error(err))
		response.FailWithMessage("测试失败: "+err.Error(), c)
		return
	}
	
	response.OkWithDetailed(result, "测试成功", c)
}