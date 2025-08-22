package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// AIChatRequest AI聊天请求
type AIChatRequest struct {
	SessionID    string `json:"sessionId" binding:"required"`    // 会话ID
	UserID       uint   `json:"userId" binding:"required"`       // 用户ID
	NovelID      uint   `json:"novelId"`                        // 小说ID
	UserMessage  string `json:"userMessage" binding:"required"`  // 用户消息
	MessageType  string `json:"messageType"`                    // 消息类型
	FunctionType string `json:"functionType"`                   // 功能类型
	Context      string `json:"context"`                        // 上下文信息
	Parameters   string `json:"parameters"`                     // 参数信息
}

// AIGenerateRequest AI内容生成请求
type AIGenerateRequest struct {
	SessionID    string `json:"sessionId" binding:"required"`   // 会话ID
	UserID       uint   `json:"userId" binding:"required"`      // 用户ID
	NovelID      uint   `json:"novelId"`                       // 小说ID
	TemplateID   uint   `json:"templateId" binding:"required"` // 模板ID
	FunctionType string `json:"functionType"`                  // 功能类型
	Context      string `json:"context"`                       // 上下文信息
	Parameters   string `json:"parameters"`                    // 参数信息
}

// AITemplateSearch AI模板搜索
type AITemplateSearch struct {
	request.PageInfo
	Name        string `json:"name" form:"name"`               // 模板名称
	Category    string `json:"category" form:"category"`       // 模板分类
	Status      string `json:"status" form:"status"`           // 模板状态
	IsPublic    *bool  `json:"isPublic" form:"isPublic"`       // 是否公开
	CreatedBy   uint   `json:"createdBy" form:"createdBy"`     // 创建者ID
	Keyword     string `json:"keyword" form:"keyword"`         // 关键词搜索
}

// AIConfigSearch AI配置搜索
type AIConfigSearch struct {
	request.PageInfo
	UserID      uint   `json:"userId" form:"userId"`           // 用户ID
	ConfigName  string `json:"configName" form:"configName"`   // 配置名称
	Provider    string `json:"provider" form:"provider"`       // 服务提供商
	Status      string `json:"status" form:"status"`           // 配置状态
	IsDefault   *bool  `json:"isDefault" form:"isDefault"`     // 是否默认
}

// AIAssistantSearch AI助手记录搜索
type AIAssistantSearch struct {
	request.PageInfo
	SessionID    string `json:"sessionId" form:"sessionId"`       // 会话ID
	UserID       uint   `json:"userId" form:"userId"`             // 用户ID
	NovelID      uint   `json:"novelId" form:"novelId"`           // 小说ID
	MessageType  string `json:"messageType" form:"messageType"`   // 消息类型
	FunctionType string `json:"functionType" form:"functionType"` // 功能类型
	Status       string `json:"status" form:"status"`             // 处理状态
	StartDate    string `json:"startDate" form:"startDate"`       // 开始日期
	EndDate      string `json:"endDate" form:"endDate"`           // 结束日期
	Keyword      string `json:"keyword" form:"keyword"`           // 关键词搜索
}

// AIFeedbackRequest AI反馈请求
type AIFeedbackRequest struct {
	AssistantID uint   `json:"assistantId" binding:"required"` // 助手记录ID
	UserRating  int    `json:"userRating" binding:"min=1,max=5"` // 用户评分
	Feedback    string `json:"feedback"`                       // 反馈内容
}

// AIUsageStatsRequest AI使用统计请求
type AIUsageStatsRequest struct {
	UserID    uint   `json:"userId" binding:"required"`    // 用户ID
	StartDate string `json:"startDate" binding:"required"` // 开始日期
	EndDate   string `json:"endDate" binding:"required"`   // 结束日期
	GroupBy   string `json:"groupBy"`                      // 分组方式: day, week, month
}

// AIWritingAssistRequest 写作助手请求
type AIWritingAssistRequest struct {
	UserID      uint   `json:"userId" binding:"required"`     // 用户ID
	NovelID     uint   `json:"novelId" binding:"required"`    // 小说ID
	ChapterID   uint   `json:"chapterId"`                     // 章节ID
	Content     string `json:"content" binding:"required"`    // 当前内容
	AssistType  string `json:"assistType" binding:"required"` // 辅助类型
	Requirement string `json:"requirement"`                   // 具体要求
	Length      int    `json:"length"`                        // 生成长度
}

// AIPlotGenerateRequest 情节生成请求
type AIPlotGenerateRequest struct {
	UserID       uint   `json:"userId" binding:"required"`      // 用户ID
	NovelID      uint   `json:"novelId" binding:"required"`     // 小说ID
	PlotType     string `json:"plotType" binding:"required"`    // 情节类型
	Characters   string `json:"characters"`                     // 相关角色
	Setting      string `json:"setting"`                       // 场景设定
	Conflict     string `json:"conflict"`                      // 冲突设定
	Tone         string `json:"tone"`                          // 情感基调
	Length       int    `json:"length"`                        // 生成长度
	Requirement  string `json:"requirement"`                   // 具体要求
}

// AICharacterGenerateRequest 角色生成请求
type AICharacterGenerateRequest struct {
	UserID        uint   `json:"userId" binding:"required"`       // 用户ID
	NovelID       uint   `json:"novelId" binding:"required"`      // 小说ID
	CharacterType string `json:"characterType" binding:"required"` // 角色类型
	Role          string `json:"role"`                            // 角色定位
	Personality   string `json:"personality"`                     // 性格特点
	Background    string `json:"background"`                      // 背景设定
	Relationships string `json:"relationships"`                   // 关系网络
	Requirement   string `json:"requirement"`                     // 具体要求
}

// AIDialogueGenerateRequest 对话生成请求
type AIDialogueGenerateRequest struct {
	UserID      uint   `json:"userId" binding:"required"`     // 用户ID
	NovelID     uint   `json:"novelId" binding:"required"`    // 小说ID
	Characters  string `json:"characters" binding:"required"` // 对话角色
	Situation   string `json:"situation" binding:"required"`  // 对话情境
	Purpose     string `json:"purpose"`                       // 对话目的
	Tone        string `json:"tone"`                          // 对话基调
	Length      int    `json:"length"`                        // 对话长度
	Requirement string `json:"requirement"`                   // 具体要求
}

// AIOutlineGenerateRequest 大纲生成请求
type AIOutlineGenerateRequest struct {
	UserID       uint   `json:"userId" binding:"required"`      // 用户ID
	NovelID      uint   `json:"novelId" binding:"required"`     // 小说ID
	OutlineType  string `json:"outlineType" binding:"required"` // 大纲类型
	Genre        string `json:"genre"`                         // 小说类型
	Theme        string `json:"theme"`                         // 主题
	MainPlot     string `json:"mainPlot"`                      // 主线剧情
	Characters   string `json:"characters"`                    // 主要角色
	ChapterCount int    `json:"chapterCount"`                  // 章节数量
	Requirement  string `json:"requirement"`                   // 具体要求
}

// AIReviewRequest 内容审查请求
type AIReviewRequest struct {
	UserID      uint   `json:"userId" binding:"required"`     // 用户ID
	NovelID     uint   `json:"novelId" binding:"required"`    // 小说ID
	Content     string `json:"content" binding:"required"`    // 待审查内容
	ReviewType  string `json:"reviewType" binding:"required"` // 审查类型
	Aspects     string `json:"aspects"`                       // 审查方面
	Requirement string `json:"requirement"`                   // 具体要求
}

// AIOptimizeRequest 内容优化请求
type AIOptimizeRequest struct {
	UserID       uint   `json:"userId" binding:"required"`      // 用户ID
	NovelID      uint   `json:"novelId" binding:"required"`     // 小说ID
	Content      string `json:"content" binding:"required"`     // 待优化内容
	OptimizeType string `json:"optimizeType" binding:"required"` // 优化类型
	TargetStyle  string `json:"targetStyle"`                    // 目标风格
	Requirement  string `json:"requirement"`                    // 具体要求
}

// AIContinueRequest 续写请求
type AIContinueRequest struct {
	UserID      uint   `json:"userId" binding:"required"`     // 用户ID
	NovelID     uint   `json:"novelId" binding:"required"`    // 小说ID
	ChapterID   uint   `json:"chapterId"`                     // 章节ID
	Content     string `json:"content" binding:"required"`    // 当前内容
	Direction   string `json:"direction"`                     // 续写方向
	Style       string `json:"style"`                        // 写作风格
	Length      int    `json:"length"`                        // 续写长度
	Requirement string `json:"requirement"`                   // 具体要求
}

// AITranslateRequest 翻译请求
type AITranslateRequest struct {
	UserID       uint   `json:"userId" binding:"required"`       // 用户ID
	Content      string `json:"content" binding:"required"`      // 待翻译内容
	SourceLang   string `json:"sourceLang" binding:"required"`   // 源语言
	TargetLang   string `json:"targetLang" binding:"required"`   // 目标语言
	TranslateType string `json:"translateType"`                  // 翻译类型
	Requirement  string `json:"requirement"`                     // 具体要求
}

// AIStyleAnalyzeRequest 风格分析请求
type AIStyleAnalyzeRequest struct {
	UserID      uint   `json:"userId" binding:"required"`     // 用户ID
	NovelID     uint   `json:"novelId" binding:"required"`    // 小说ID
	Content     string `json:"content" binding:"required"`    // 待分析内容
	AnalyzeType string `json:"analyzeType" binding:"required"` // 分析类型
	Aspects     string `json:"aspects"`                       // 分析方面
	Requirement string `json:"requirement"`                   // 具体要求
}