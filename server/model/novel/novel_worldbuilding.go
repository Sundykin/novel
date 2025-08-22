package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelWorldbuilding 世界观设定模型
type NovelWorldbuilding struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	Category    string `json:"category" gorm:"size:50;not null;comment:设定分类"`
	Name        string `json:"name" gorm:"size:200;not null;comment:设定名称"`
	Description string `json:"description" gorm:"type:text;not null;comment:设定描述"`
	
	// 详细信息
	Details         string `json:"details" gorm:"type:text;comment:详细信息"`
	RelatedElements string `json:"relatedElements" gorm:"type:text;comment:相关元素"`
	Rules           string `json:"rules" gorm:"type:text;comment:规则说明"`
	History         string `json:"history" gorm:"type:text;comment:历史背景"`
	
	// 状态信息
	Status          string `json:"status" gorm:"size:20;default:draft;comment:设定状态"`
	Importance      int    `json:"importance" gorm:"default:1;comment:重要程度1-10"`
	
	// 标签和分类
	Tags            string `json:"tags" gorm:"type:text;comment:标签"`
	SubCategory     string `json:"subCategory" gorm:"size:50;comment:子分类"`
	
	// 关联信息
	NovelID         uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy       uint   `json:"createdBy" gorm:"comment:创建者ID"`
}

// TableName 设置表名
func (NovelWorldbuilding) TableName() string {
	return "novel_worldbuildings"
}

// WorldbuildingCategory 世界观分类常量
const (
	WorldbuildingCategoryGeography = "geography" // 地理
	WorldbuildingCategoryHistory   = "history"   // 历史
	WorldbuildingCategoryCulture   = "culture"   // 文化
	WorldbuildingCategoryPolitics  = "politics"  // 政治
	WorldbuildingCategoryReligion  = "religion"  // 宗教
	WorldbuildingCategoryMagic     = "magic"     // 魔法
	WorldbuildingCategoryTechnology = "technology" // 科技
	WorldbuildingCategoryEconomy   = "economy"   // 经济
	WorldbuildingCategorySociety   = "society"   // 社会
	WorldbuildingCategoryLanguage  = "language"  // 语言
	WorldbuildingCategoryOrganization = "organization" // 组织
)

// WorldbuildingStatus 世界观状态常量
const (
	WorldbuildingStatusDraft    = "draft"    // 草稿
	WorldbuildingStatusActive   = "active"   // 活跃
	WorldbuildingStatusArchived = "archived" // 已归档
	WorldbuildingStatusObsolete = "obsolete" // 已废弃
)