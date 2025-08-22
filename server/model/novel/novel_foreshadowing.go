package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelForeshadowing 伏笔管理模型
type NovelForeshadowing struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	ForeshadowingID  string `json:"foreshadowingId" gorm:"size:50;uniqueIndex;not null;comment:伏笔唯一标识"`
	Title            string `json:"title" gorm:"size:200;not null;comment:伏笔标题"`
	Content          string `json:"content" gorm:"type:text;not null;comment:伏笔内容"`
	Description      string `json:"description" gorm:"type:text;comment:伏笔描述"`
	
	// 重要性和类型
	ImportanceLevel  string `json:"importanceLevel" gorm:"size:20;default:medium;comment:重要程度"`
	Type             string `json:"type" gorm:"size:50;comment:伏笔类型"`
	Category         string `json:"category" gorm:"size:50;comment:伏笔分类"`
	
	// 状态信息
	Status           string `json:"status" gorm:"size:20;default:planted;comment:伏笔状态"`
	QualityScore     *int   `json:"qualityScore" gorm:"comment:质量评分1-10"`
	
	// 章节信息
	PlantedChapter   *int   `json:"plantedChapter" gorm:"comment:埋下伏笔的章节"`
	RevealedChapter  *int   `json:"revealedChapter" gorm:"comment:揭示伏笔的章节"`
	ResolvedChapter  *int   `json:"resolvedChapter" gorm:"comment:解决伏笔的章节"`
	
	// 关联信息
	RelatedCharacters string `json:"relatedCharacters" gorm:"type:text;comment:相关角色"`
	RelatedEvents     string `json:"relatedEvents" gorm:"type:text;comment:相关事件"`
	Tags              string `json:"tags" gorm:"type:text;comment:标签"`
	
	// 所属信息
	NovelID           uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy         uint   `json:"createdBy" gorm:"comment:创建者ID"`
}

// TableName 设置表名
func (NovelForeshadowing) TableName() string {
	return "novel_foreshadowings"
}

// ForeshadowingImportance 重要程度常量
const (
	ForeshadowingImportanceLow      = "low"      // 低
	ForeshadowingImportanceMedium   = "medium"   // 中
	ForeshadowingImportanceHigh     = "high"     // 高
	ForeshadowingImportanceCritical = "critical" // 关键
)

// ForeshadowingStatus 伏笔状态常量
const (
	ForeshadowingStatusPlanted    = "planted"    // 已埋下
	ForeshadowingStatusDeveloping = "developing" // 发展中
	ForeshadowingStatusRevealed   = "revealed"   // 已揭示
	ForeshadowingStatusResolved   = "resolved"   // 已解决
	ForeshadowingStatusAbandoned  = "abandoned"  // 已废弃
)

// ForeshadowingType 伏笔类型常量
const (
	ForeshadowingTypeCharacter = "character" // 角色伏笔
	ForeshadowingTypePlot      = "plot"      // 情节伏笔
	ForeshadowingTypeWorld     = "world"     // 世界观伏笔
	ForeshadowingTypeItem      = "item"      // 物品伏笔
	ForeshadowingTypeEvent     = "event"     // 事件伏笔
	ForeshadowingTypeMystery   = "mystery"   // 悬疑伏笔
)