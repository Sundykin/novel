package novel

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel/request"
	"gorm.io/gorm"
)

type AIService struct{}

// OpenAI API 请求结构
type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float32   `json:"temperature,omitempty"`
	TopP        float32   `json:"top_p,omitempty"`
	Stream      bool      `json:"stream,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
	Error   *APIError `json:"error,omitempty"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}

// AI助手对话
func (aiService *AIService) ChatWithAI(req *request.AIChatRequest) (*novel.AIAssistant, error) {
	startTime := time.Now()
	
	// 获取用户AI配置
	config, err := aiService.GetUserAIConfig(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("获取AI配置失败: %v", err)
	}
	
	// 构建消息历史
	messages, err := aiService.BuildMessageHistory(req.SessionID, req.UserMessage, req.Context)
	if err != nil {
		return nil, fmt.Errorf("构建消息历史失败: %v", err)
	}
	
	// 调用AI API
	aiResponse, tokenUsed, err := aiService.CallAIAPI(config, messages)
	if err != nil {
		return nil, fmt.Errorf("AI API调用失败: %v", err)
	}
	
	processTime := int(time.Since(startTime).Milliseconds())
	
	// 保存对话记录
	assistant := &novel.AIAssistant{
		SessionID:    req.SessionID,
		UserID:       req.UserID,
		NovelID:      req.NovelID,
		UserMessage:  req.UserMessage,
		AIResponse:   aiResponse,
		MessageType:  req.MessageType,
		FunctionType: req.FunctionType,
		Context:      req.Context,
		Parameters:   req.Parameters,
		Status:       novel.StatusCompleted,
		ProcessTime:  processTime,
		TokenUsed:    tokenUsed,
		CreatedBy:    req.UserID,
	}
	
	err = global.GVA_DB.Create(assistant).Error
	if err != nil {
		return nil, fmt.Errorf("保存对话记录失败: %v", err)
	}
	
	return assistant, nil
}

// 内容生成
func (aiService *AIService) GenerateContent(req *request.AIGenerateRequest) (*novel.AIAssistant, error) {
	startTime := time.Now()
	
	// 获取模板
	template, err := aiService.GetTemplate(req.TemplateID)
	if err != nil {
		return nil, fmt.Errorf("获取模板失败: %v", err)
	}
	
	// 构建提示词
	prompt := aiService.BuildPrompt(template, req.Parameters)
	
	// 获取用户AI配置
	config, err := aiService.GetUserAIConfig(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("获取AI配置失败: %v", err)
	}
	
	// 构建消息
	messages := []Message{
		{Role: "system", Content: template.SystemPrompt},
		{Role: "user", Content: prompt},
	}
	
	// 调用AI API
	aiResponse, tokenUsed, err := aiService.CallAIAPI(config, messages)
	if err != nil {
		return nil, fmt.Errorf("AI API调用失败: %v", err)
	}
	
	processTime := int(time.Since(startTime).Milliseconds())
	
	// 保存生成记录
	assistant := &novel.AIAssistant{
		SessionID:    req.SessionID,
		UserID:       req.UserID,
		NovelID:      req.NovelID,
		UserMessage:  prompt,
		AIResponse:   aiResponse,
		MessageType:  novel.MessageTypeGenerate,
		FunctionType: req.FunctionType,
		Context:      req.Context,
		Parameters:   req.Parameters,
		Status:       novel.StatusCompleted,
		ProcessTime:  processTime,
		TokenUsed:    tokenUsed,
		CreatedBy:    req.UserID,
	}
	
	err = global.GVA_DB.Create(assistant).Error
	if err != nil {
		return nil, fmt.Errorf("保存生成记录失败: %v", err)
	}
	
	return assistant, nil
}

// 获取用户AI配置
func (aiService *AIService) GetUserAIConfig(userID uint) (*novel.AIConfig, error) {
	var config novel.AIConfig
	err := global.GVA_DB.Where("user_id = ? AND is_default = ? AND status = ?", 
		userID, true, novel.ConfigStatusActive).First(&config).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有默认配置，创建一个基础配置
			return aiService.CreateDefaultConfig(userID)
		}
		return nil, err
	}
	
	return &config, nil
}

// 创建默认配置
func (aiService *AIService) CreateDefaultConfig(userID uint) (*novel.AIConfig, error) {
	config := &novel.AIConfig{
		UserID:           userID,
		ConfigName:       "默认配置",
		Provider:         novel.ProviderOpenAI,
		Model:            "gpt-3.5-turbo",
		MaxTokens:        2000,
		Temperature:      0.7,
		TopP:             1.0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		DailyLimit:       1000,
		MonthlyLimit:     30000,
		IsDefault:        true,
		Status:           novel.ConfigStatusActive,
		CreatedBy:        userID,
	}
	
	err := global.GVA_DB.Create(config).Error
	if err != nil {
		return nil, err
	}
	
	return config, nil
}

// 构建消息历史
func (aiService *AIService) BuildMessageHistory(sessionID, userMessage, context string) ([]Message, error) {
	messages := []Message{}
	
	// 添加系统提示
	systemPrompt := "你是一个专业的小说写作助手，擅长帮助作者进行创作。请根据用户的需求提供有用的建议和内容。"
	messages = append(messages, Message{Role: "system", Content: systemPrompt})
	
	// 如果有上下文，添加上下文信息
	if context != "" {
		messages = append(messages, Message{Role: "system", Content: "上下文信息: " + context})
	}
	
	// 获取最近的对话历史（最多10条）
	var history []novel.AIAssistant
	err := global.GVA_DB.Where("session_id = ?", sessionID).
		Order("created_at DESC").Limit(10).Find(&history).Error
	
	if err != nil {
		return nil, err
	}
	
	// 反转历史记录顺序
	for i := len(history) - 1; i >= 0; i-- {
		h := history[i]
		messages = append(messages, Message{Role: "user", Content: h.UserMessage})
		messages = append(messages, Message{Role: "assistant", Content: h.AIResponse})
	}
	
	// 添加当前用户消息
	messages = append(messages, Message{Role: "user", Content: userMessage})
	
	return messages, nil
}

// 调用AI API
func (aiService *AIService) CallAIAPI(config *novel.AIConfig, messages []Message) (string, int, error) {
	switch config.Provider {
	case novel.ProviderOpenAI:
		return aiService.CallOpenAI(config, messages)
	default:
		return "", 0, fmt.Errorf("不支持的AI服务提供商: %s", config.Provider)
	}
}

// 调用OpenAI API
func (aiService *AIService) CallOpenAI(config *novel.AIConfig, messages []Message) (string, int, error) {
	apiKey := config.APIKey
	if apiKey == "" {
		// 使用系统默认API Key (需要在配置中添加)
		apiKey = "" // 暂时留空，需要用户配置
	}
	
	if apiKey == "" {
		return "", 0, errors.New("未配置OpenAI API Key")
	}
	
	endpoint := config.APIEndpoint
	if endpoint == "" {
		endpoint = "https://api.openai.com/v1/chat/completions"
	}
	
	reqBody := OpenAIRequest{
		Model:       config.Model,
		Messages:    messages,
		MaxTokens:   config.MaxTokens,
		Temperature: config.Temperature,
		TopP:        config.TopP,
	}
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", 0, fmt.Errorf("序列化请求失败: %v", err)
	}
	
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", 0, fmt.Errorf("创建请求失败: %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("API请求失败: %v", err)
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("读取响应失败: %v", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("API请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}
	
	var apiResp OpenAIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return "", 0, fmt.Errorf("解析响应失败: %v", err)
	}
	
	if apiResp.Error != nil {
		return "", 0, fmt.Errorf("API错误: %s", apiResp.Error.Message)
	}
	
	if len(apiResp.Choices) == 0 {
		return "", 0, errors.New("API响应中没有选择项")
	}
	
	return apiResp.Choices[0].Message.Content, apiResp.Usage.TotalTokens, nil
}

// 获取模板
func (aiService *AIService) GetTemplate(templateID uint) (*novel.AITemplate, error) {
	var template novel.AITemplate
	err := global.GVA_DB.Where("id = ? AND status = ?", templateID, novel.TemplateStatusActive).First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// 构建提示词
func (aiService *AIService) BuildPrompt(template *novel.AITemplate, parameters string) string {
	prompt := template.Prompt
	
	// 解析参数并替换占位符
	if parameters != "" {
		var params map[string]interface{}
		err := json.Unmarshal([]byte(parameters), &params)
		if err == nil {
			for key, value := range params {
				placeholder := fmt.Sprintf("{{%s}}", key)
				prompt = strings.ReplaceAll(prompt, placeholder, fmt.Sprintf("%v", value))
			}
		}
	}
	
	return prompt
}

// 获取对话历史
func (aiService *AIService) GetChatHistory(sessionID string, page, pageSize int) ([]novel.AIAssistant, int64, error) {
	var history []novel.AIAssistant
	var total int64
	
	db := global.GVA_DB.Model(&novel.AIAssistant{}).Where("session_id = ?", sessionID)
	
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = db.Order("created_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&history).Error
	if err != nil {
		return nil, 0, err
	}
	
	return history, total, nil
}

// 用户反馈
func (aiService *AIService) SubmitFeedback(assistantID uint, rating int, feedback string) error {
	return global.GVA_DB.Model(&novel.AIAssistant{}).Where("id = ?", assistantID).
		Updates(map[string]interface{}{
			"user_rating": rating,
			"feedback":    feedback,
		}).Error
}

// 获取使用统计
func (aiService *AIService) GetUsageStats(userID uint, startDate, endDate time.Time) (map[string]interface{}, error) {
	var stats struct {
		TotalRequests int64 `json:"totalRequests"`
		TotalTokens   int `json:"totalTokens"`
		AvgProcessTime float64 `json:"avgProcessTime"`
		SuccessRate   float64 `json:"successRate"`
	}
	
	db := global.GVA_DB.Model(&novel.AIAssistant{}).Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, startDate, endDate)
	
	// 总请求数
	db.Count(&stats.TotalRequests)
	
	// 总Token使用量
	db.Select("COALESCE(SUM(token_used), 0)").Scan(&stats.TotalTokens)
	
	// 平均处理时间
	db.Select("COALESCE(AVG(process_time), 0)").Scan(&stats.AvgProcessTime)
	
	// 成功率
	var successCount int64
	db.Where("status = ?", novel.StatusCompleted).Count(&successCount)
	if stats.TotalRequests > 0 {
		stats.SuccessRate = float64(successCount) / float64(stats.TotalRequests) * 100
	}
	
	return map[string]interface{}{
		"totalRequests":  stats.TotalRequests,
		"totalTokens":    stats.TotalTokens,
		"avgProcessTime": stats.AvgProcessTime,
		"successRate":    stats.SuccessRate,
	}, nil
}

// AI配置管理相关方法

// CreateAIConfig 创建AI配置
func (aiService *AIService) CreateAIConfig(aiConfig *novel.AIConfig) error {
	return global.GVA_DB.Create(aiConfig).Error
}

// DeleteAIConfig 删除AI配置
func (aiService *AIService) DeleteAIConfig(aiConfig *novel.AIConfig) error {
	return global.GVA_DB.Delete(aiConfig).Error
}

// DeleteAIConfigByIds 批量删除AI配置
func (aiService *AIService) DeleteAIConfigByIds(ids *request.IdsReq) error {
	return global.GVA_DB.Delete(&[]novel.AIConfig{}, "id in ?", ids.Ids).Error
}

// UpdateAIConfig 更新AI配置
func (aiService *AIService) UpdateAIConfig(aiConfig *novel.AIConfig) error {
	return global.GVA_DB.Save(aiConfig).Error
}

// GetAIConfig 根据ID获取AI配置
func (aiService *AIService) GetAIConfig(id string) (*novel.AIConfig, error) {
	var aiConfig novel.AIConfig
	err := global.GVA_DB.Where("id = ?", id).First(&aiConfig).Error
	return &aiConfig, err
}

// GetAIConfigInfoList 分页获取AI配置列表
func (aiService *AIService) GetAIConfigInfoList(info request.AIConfigSearch) (list []novel.AIConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	
	// 构建查询条件
	db := global.GVA_DB.Model(&novel.AIConfig{})
	
	// 关键词搜索
	if info.Keyword != "" {
		db = db.Where("name LIKE ? OR description LIKE ?", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
	}
	
	// 提供商筛选
	if info.Provider != "" {
		db = db.Where("provider = ?", info.Provider)
	}
	
	// 状态筛选
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Order("updated_at DESC").Find(&list).Error
	return list, total, err
}

// TestAIConfig 测试AI配置
func (aiService *AIService) TestAIConfig(req *request.AIConfigTestRequest) (map[string]interface{}, error) {
	// 构建测试消息
	messages := []Message{
		{
			Role:    "user",
			Content: req.TestMessage,
		},
	}
	
	// 创建临时配置用于测试
	testConfig := &novel.AIConfig{
		Provider: req.Provider,
		APIKey:   req.APIKey,
		BaseURL:  req.BaseURL,
		Model:    req.Model,
	}
	
	// 调用AI API
	response, tokens, err := aiService.CallAIAPI(testConfig, messages)
	if err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"success":  true,
		"response": response,
		"tokens":   tokens,
		"message":  "AI配置测试成功",
	}, nil
}

// WritingAssist 写作助手
func (aiService *AIService) WritingAssist(req *request.AIWritingAssistRequest) (*novel.AIAssistant, error) {
	// 获取用户AI配置
	config, err := aiService.GetUserAIConfig(req.UserID)
	if err != nil {
		return nil, err
	}
	
	// 构建写作助手提示
	prompt := fmt.Sprintf("请为以下内容提供写作建议：\n\n内容：%s\n\n写作类型：%s\n目标：%s", req.Content, req.WritingType, req.Goal)
	
	messages := []Message{
		{Role: "user", Content: prompt},
	}
	
	// 调用AI API
	response, tokens, err := aiService.CallAIAPI(config, messages)
	if err != nil {
		return nil, err
	}
	
	// 保存助手记录
	assistant := &novel.AIAssistant{
		UserID:      req.UserID,
		SessionID:   req.SessionID,
		RequestType: "writing_assist",
		UserMessage: req.Content,
		AIResponse:  response,
		TokenUsed:   tokens,
		Status:      novel.StatusCompleted,
	}
	
	err = global.GVA_DB.Create(assistant).Error
	return assistant, err
}

// PlotGenerate 情节生成
func (aiService *AIService) PlotGenerate(req *request.AIPlotGenerateRequest) (*novel.AIAssistant, error) {
	// 获取用户AI配置
	config, err := aiService.GetUserAIConfig(req.UserID)
	if err != nil {
		return nil, err
	}
	
	// 构建情节生成提示
	prompt := fmt.Sprintf("请生成一个%s类型的情节：\n\n背景：%s\n角色：%s\n冲突：%s\n长度要求：%s", req.PlotType, req.Background, req.Characters, req.Conflict, req.Length)
	
	messages := []Message{
		{Role: "user", Content: prompt},
	}
	
	// 调用AI API
	response, tokens, err := aiService.CallAIAPI(config, messages)
	if err != nil {
		return nil, err
	}
	
	// 保存助手记录
	assistant := &novel.AIAssistant{
		UserID:      req.UserID,
		SessionID:   req.SessionID,
		RequestType: "plot_generate",
		UserMessage: prompt,
		AIResponse:  response,
		TokenUsed:   tokens,
		Status:      novel.StatusCompleted,
	}
	
	err = global.GVA_DB.Create(assistant).Error
	return assistant, err
}

// CharacterGenerate 角色生成
func (aiService *AIService) CharacterGenerate(req *request.AICharacterGenerateRequest) (*novel.AIAssistant, error) {
	// 获取用户AI配置
	config, err := aiService.GetUserAIConfig(req.UserID)
	if err != nil {
		return nil, err
	}
	
	// 构建角色生成提示
	prompt := fmt.Sprintf("请创建一个%s角色：\n\n角色类型：%s\n性格特点：%s\n背景故事：%s\n外貌描述：%s", req.CharacterType, req.Role, req.Personality, req.Background, req.Appearance)
	
	messages := []Message{
		{Role: "user", Content: prompt},
	}
	
	// 调用AI API
	response, tokens, err := aiService.CallAIAPI(config, messages)
	if err != nil {
		return nil, err
	}
	
	// 保存助手记录
	assistant := &novel.AIAssistant{
		UserID:      req.UserID,
		SessionID:   req.SessionID,
		RequestType: "character_generate",
		UserMessage: prompt,
		AIResponse:  response,
		TokenUsed:   tokens,
		Status:      novel.StatusCompleted,
	}
	
	err = global.GVA_DB.Create(assistant).Error
	return assistant, err
}

// DialogueGenerate 对话生成
func (aiService *AIService) DialogueGenerate(req *request.AIDialogueGenerateRequest) (*novel.AIAssistant, error) {
	// 获取用户AI配置
	config, err := aiService.GetUserAIConfig(req.UserID)
	if err != nil {
		return nil, err
	}
	
	// 构建对话生成提示
	prompt := fmt.Sprintf("请生成一段对话：\n\n场景：%s\n角色：%s\n情绪：%s\n目的：%s", req.Scene, req.Characters, req.Emotion, req.Purpose)
	
	messages := []Message{
		{Role: "user", Content: prompt},
	}
	
	// 调用AI API
	response, tokens, err := aiService.CallAIAPI(config, messages)
	if err != nil {
		return nil, err
	}
	
	// 保存助手记录
	assistant := &novel.AIAssistant{
		UserID:      req.UserID,
		SessionID:   req.SessionID,
		RequestType: "dialogue_generate",
		UserMessage: prompt,
		AIResponse:  response,
		TokenUsed:   tokens,
		Status:      novel.StatusCompleted,
	}
	
	err = global.GVA_DB.Create(assistant).Error
	return assistant, err
}

// ContinueWriting 续写
func (aiService *AIService) ContinueWriting(req *request.AIContinueRequest) (*novel.AIAssistant, error) {
	// 获取用户AI配置
	config, err := aiService.GetUserAIConfig(req.UserID)
	if err != nil {
		return nil, err
	}
	
	// 构建续写提示
	prompt := fmt.Sprintf("请根据以下内容进行续写：\n\n当前内容：%s\n\n续写方向：%s\n写作风格：%s\n续写长度：约%d字", req.Content, req.Direction, req.Style, req.Length)
	
	messages := []Message{
		{Role: "user", Content: prompt},
	}
	
	// 调用AI API
	response, tokens, err := aiService.CallAIAPI(config, messages)
	if err != nil {
		return nil, err
	}
	
	// 保存助手记录
	assistant := &novel.AIAssistant{
		UserID:      req.UserID,
		SessionID:   req.SessionID,
		RequestType: "continue_writing",
		UserMessage: req.Content,
		AIResponse:  response,
		TokenUsed:   tokens,
		Status:      novel.StatusCompleted,
	}
	
	err = global.GVA_DB.Create(assistant).Error
	return assistant, err
}