package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelOutline 小说大纲模型
type NovelOutline struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	VolumeNumber int    `json:"volumeNumber" gorm:"not null;comment:卷号"`
	ChapterStart int    `json:"chapterStart" gorm:"not null;comment:起始章节"`
	ChapterEnd   int    `json:"chapterEnd" gorm:"not null;comment:结束章节"`
	Title        string `json:"title" gorm:"size:200;not null;comment:大纲标题"`
	Summary      string `json:"summary" gorm:"type:text;not null;comment:大纲摘要"`
	
	// 详细内容
	MainPlot              string `json:"mainPlot" gorm:"type:text;comment:主线剧情"`
	SubPlots              string `json:"subPlots" gorm:"type:text;comment:支线剧情"`
	CharacterDevelopment  string `json:"characterDevelopment" gorm:"type:text;comment:角色发展"`
	KeyEvents             string `json:"keyEvents" gorm:"type:text;comment:关键事件"`
	
	// 伏笔相关
	ForeshadowingPlanted  string `json:"foreshadowingPlanted" gorm:"type:text;comment:埋下的伏笔"`
	ForeshadowingRevealed string `json:"foreshadowingRevealed" gorm:"type:text;comment:揭示的伏笔"`
	
	// 状态信息
	Status                string `json:"status" gorm:"size:20;default:draft;comment:大纲状态"`
	Progress              int    `json:"progress" gorm:"default:0;comment:完成进度0-100"`
	
	// 目标信息
	WordCountTarget       int    `json:"wordCountTarget" gorm:"default:0;comment:目标字数"`
	EstimatedDays         int    `json:"estimatedDays" gorm:"default:0;comment:预计天数"`
	
	// 关联信息
	NovelID               uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy             uint   `json:"createdBy" gorm:"comment:创建者ID"`
	
	// 关联关系
	Chapters              []NovelChapter `json:"chapters" gorm:"foreignKey:OutlineID"`
}

// TableName 设置表名
func (NovelOutline) TableName() string {
	return "novel_outlines"
}

// OutlineStatus 大纲状态常量
const (
	OutlineStatusDraft     = "draft"     // 草稿
	OutlineStatusActive    = "active"    // 活跃
	OutlineStatusCompleted = "completed" // 已完成
	OutlineStatusArchived  = "archived"  // 已归档
	OutlineStatusCancelled = "cancelled" // 已取消
)