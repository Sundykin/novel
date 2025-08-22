package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelChapter 小说章节模型
type NovelChapter struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	ChapterNumber int    `json:"chapterNumber" gorm:"not null;comment:章节号"`
	Volume        int    `json:"volume" gorm:"not null;comment:卷号"`
	ChapterTitle  string `json:"chapterTitle" gorm:"size:200;not null;comment:章节标题"`
	Content       string `json:"content" gorm:"type:longtext;comment:章节内容"`
	
	// 章节规划
	MainEvents           string `json:"mainEvents" gorm:"type:text;comment:主要事件"`
	EmotionalTone        string `json:"emotionalTone" gorm:"size:100;comment:情感基调"`
	MainPlot             string `json:"mainPlot" gorm:"type:text;comment:主线剧情"`
	SubPlots             string `json:"subPlots" gorm:"type:text;comment:支线剧情"`
	CharacterDevelopment string `json:"characterDevelopment" gorm:"type:text;comment:角色发展"`
	KeyEvents            string `json:"keyEvents" gorm:"type:text;comment:关键事件"`
	
	// 伏笔相关
	ForeshadowingPlanted  string `json:"foreshadowingPlanted" gorm:"type:text;comment:埋下的伏笔"`
	ForeshadowingRevealed string `json:"foreshadowingRevealed" gorm:"type:text;comment:揭示的伏笔"`
	
	// 统计信息
	WordCount       int `json:"wordCount" gorm:"default:0;comment:实际字数"`
	WordCountTarget int `json:"wordCountTarget" gorm:"default:0;comment:目标字数"`
	
	// 状态信息
	Status          string    `json:"status" gorm:"size:20;default:draft;comment:章节状态"`
	PublishedAt     *time.Time `json:"publishedAt" gorm:"comment:发布时间"`
	
	// 质量评估
	QualityScore    *int   `json:"qualityScore" gorm:"comment:质量评分1-10"`
	ReviewNotes     string `json:"reviewNotes" gorm:"type:text;comment:审阅备注"`
	
	// 关联信息
	OutlineID       uint   `json:"outlineId" gorm:"not null;comment:所属大纲ID"`
	NovelID         uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy       uint   `json:"createdBy" gorm:"comment:创建者ID"`
	
	// 关联关系
	Outline         NovelOutline `json:"outline" gorm:"foreignKey:OutlineID"`
}

// TableName 设置表名
func (NovelChapter) TableName() string {
	return "novel_chapters"
}

// ChapterStatus 章节状态常量
const (
	ChapterStatusDraft     = "draft"     // 草稿
	ChapterStatusWriting   = "writing"   // 写作中
	ChapterStatusReview    = "review"    // 审阅中
	ChapterStatusRevision  = "revision"  // 修订中
	ChapterStatusCompleted = "completed" // 已完成
	ChapterStatusPublished = "published" // 已发布
	ChapterStatusArchived  = "archived"  // 已归档
)

// EmotionalTone 情感基调常量
const (
	EmotionalToneHappy     = "happy"     // 欢快
	EmotionalToneSad       = "sad"       // 悲伤
	EmotionalToneExciting  = "exciting"  // 激动
	EmotionalToneTense     = "tense"     // 紧张
	EmotionalToneMysterious = "mysterious" // 神秘
	EmotionalToneRomantic  = "romantic"  // 浪漫
	EmotionalToneAction    = "action"    // 动作
	EmotionalToneDrama     = "drama"     // 戏剧
)