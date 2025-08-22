package novel

import (
	"time"
	"gorm.io/gorm"
)

// AIAssistant AI助手模型
type AIAssistant struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	SessionID    string `json:"sessionId" gorm:"size:100;not null;comment:会话ID"`
	UserID       uint   `json:"userId" gorm:"not null;comment:用户ID"`
	NovelID      uint   `json:"novelId" gorm:"comment:关联小说ID"`
	
	// 对话信息
	UserMessage  string `json:"userMessage" gorm:"type:text;not null;comment:用户消息"`
	AIResponse   string `json:"aiResponse" gorm:"type:text;comment:AI回复"`
	MessageType  string `json:"messageType" gorm:"size:50;default:chat;comment:消息类型"`
	
	// 功能信息
	FunctionType string `json:"functionType" gorm:"size:50;comment:功能类型"`
	Context      string `json:"context" gorm:"type:text;comment:上下文信息"`
	Parameters   string `json:"parameters" gorm:"type:text;comment:参数信息"`
	
	// 状态信息
	Status       string `json:"status" gorm:"size:20;default:completed;comment:处理状态"`
	ProcessTime  int    `json:"processTime" gorm:"comment:处理时间(毫秒)"`
	TokenUsed    int    `json:"tokenUsed" gorm:"comment:使用的Token数量"`
	
	// 质量评估
	UserRating   *int   `json:"userRating" gorm:"comment:用户评分1-5"`
	Feedback     string `json:"feedback" gorm:"type:text;comment:用户反馈"`
	
	// 关联信息
	CreatedBy    uint   `json:"createdBy" gorm:"comment:创建者ID"`
}

// TableName 设置表名
func (AIAssistant) TableName() string {
	return "ai_assistants"
}

// 消息类型常量
const (
	MessageTypeChat         = "chat"          // 普通聊天
	MessageTypeGenerate     = "generate"      // 内容生成
	MessageTypeAnalyze      = "analyze"       // 内容分析
	MessageTypeOptimize     = "optimize"      // 内容优化
	MessageTypeSuggestion   = "suggestion"    // 建议
	MessageTypeCorrection   = "correction"    // 纠错
)

// 功能类型常量
const (
	FunctionTypeWriting     = "writing"       // 写作助手
	FunctionTypePlot        = "plot"          // 情节生成
	FunctionTypeCharacter   = "character"     // 角色生成
	FunctionTypeDialogue    = "dialogue"      // 对话生成
	FunctionTypeDescription = "description"   // 描述生成
	FunctionTypeOutline     = "outline"       // 大纲生成
	FunctionTypeReview      = "review"        // 内容审查
	FunctionTypeTranslate   = "translate"     // 翻译
	FunctionTypeStyle       = "style"         // 风格分析
	FunctionTypeContinue    = "continue"      // 续写
)

// 处理状态常量
const (
	StatusPending    = "pending"     // 处理中
	StatusCompleted  = "completed"   // 已完成
	StatusFailed     = "failed"      // 失败
	StatusCancelled  = "cancelled"   // 已取消
)

// AITemplate AI模板模型
type AITemplate struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	Name         string `json:"name" gorm:"size:200;not null;comment:模板名称"`
	Description  string `json:"description" gorm:"type:text;comment:模板描述"`
	Category     string `json:"category" gorm:"size:50;not null;comment:模板分类"`
	
	// 模板内容
	Prompt       string `json:"prompt" gorm:"type:text;not null;comment:提示词模板"`
	SystemPrompt string `json:"systemPrompt" gorm:"type:text;comment:系统提示词"`
	Examples     string `json:"examples" gorm:"type:text;comment:示例内容"`
	
	// 配置信息
	Parameters   string `json:"parameters" gorm:"type:text;comment:参数配置"`
	MaxTokens    int    `json:"maxTokens" gorm:"default:2000;comment:最大Token数"`
	Temperature  float32 `json:"temperature" gorm:"default:0.7;comment:创造性参数"`
	
	// 使用统计
	UsageCount   int    `json:"usageCount" gorm:"default:0;comment:使用次数"`
	SuccessRate  float32 `json:"successRate" gorm:"default:0;comment:成功率"`
	AvgRating    float32 `json:"avgRating" gorm:"default:0;comment:平均评分"`
	
	// 状态信息
	Status       string `json:"status" gorm:"size:20;default:active;comment:模板状态"`
	IsPublic     bool   `json:"isPublic" gorm:"default:false;comment:是否公开"`
	
	// 关联信息
	CreatedBy    uint   `json:"createdBy" gorm:"comment:创建者ID"`
}

// TableName 设置表名
func (AITemplate) TableName() string {
	return "ai_templates"
}

// 模板分类常量
const (
	TemplateCategoryWriting     = "writing"      // 写作类
	TemplateCategoryPlot        = "plot"         // 情节类
	TemplateCategoryCharacter   = "character"    // 角色类
	TemplateCategoryDialogue    = "dialogue"     // 对话类
	TemplateCategoryDescription = "description"  // 描述类
	TemplateCategoryOutline     = "outline"      // 大纲类
	TemplateCategoryStyle       = "style"        // 风格类
	TemplateCategoryOther       = "other"        // 其他类
)

// 模板状态常量
const (
	TemplateStatusActive   = "active"    // 激活
	TemplateStatusInactive = "inactive"  // 停用
	TemplateStatusDraft    = "draft"     // 草稿
	TemplateStatusArchived = "archived"  // 归档
)

// AIConfig AI配置模型
type AIConfig struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	UserID       uint   `json:"userId" gorm:"not null;comment:用户ID"`
	ConfigName   string `json:"configName" gorm:"size:100;not null;comment:配置名称"`
	
	// AI服务配置
	Provider     string `json:"provider" gorm:"size:50;not null;comment:AI服务提供商"`
	Model        string `json:"model" gorm:"size:100;not null;comment:AI模型"`
	APIKey       string `json:"apiKey" gorm:"size:500;comment:API密钥"`
	APIEndpoint  string `json:"apiEndpoint" gorm:"size:500;comment:API端点"`
	
	// 默认参数
	MaxTokens    int     `json:"maxTokens" gorm:"default:2000;comment:最大Token数"`
	Temperature  float32 `json:"temperature" gorm:"default:0.7;comment:创造性参数"`
	TopP         float32 `json:"topP" gorm:"default:1.0;comment:核采样参数"`
	FrequencyPenalty float32 `json:"frequencyPenalty" gorm:"default:0;comment:频率惩罚"`
	PresencePenalty  float32 `json:"presencePenalty" gorm:"default:0;comment:存在惩罚"`
	
	// 使用限制
	DailyLimit   int    `json:"dailyLimit" gorm:"default:1000;comment:每日使用限制"`
	MonthlyLimit int    `json:"monthlyLimit" gorm:"default:30000;comment:每月使用限制"`
	
	// 状态信息
	IsDefault    bool   `json:"isDefault" gorm:"default:false;comment:是否默认配置"`
	Status       string `json:"status" gorm:"size:20;default:active;comment:配置状态"`
	
	// 关联信息
	CreatedBy    uint   `json:"createdBy" gorm:"comment:创建者ID"`
}

// TableName 设置表名
func (AIConfig) TableName() string {
	return "ai_configs"
}

// AI服务提供商常量
const (
	ProviderOpenAI    = "openai"     // OpenAI
	ProviderClaude    = "claude"     // Anthropic Claude
	ProviderGemini    = "gemini"     // Google Gemini
	ProviderQianwen   = "qianwen"    // 阿里千问
	ProviderBaidu     = "baidu"      // 百度文心
	ProviderTencent   = "tencent"    // 腾讯混元
	ProviderLocal     = "local"      // 本地模型
	ProviderCustom    = "custom"     // 自定义
)

// 配置状态常量
const (
	ConfigStatusActive   = "active"    // 激活
	ConfigStatusInactive = "inactive"  // 停用
	ConfigStatusTesting  = "testing"   // 测试中
	ConfigStatusError    = "error"     // 错误
)