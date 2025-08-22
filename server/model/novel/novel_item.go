package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelItem 物品道具模型
type NovelItem struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	Name        string `json:"name" gorm:"size:200;not null;comment:物品名称"`
	Category    string `json:"category" gorm:"size:50;not null;comment:物品分类"`
	Description string `json:"description" gorm:"type:text;not null;comment:物品描述"`
	
	// 详细信息
	Properties  string `json:"properties" gorm:"type:text;comment:物品属性"`
	Appearance  string `json:"appearance" gorm:"type:text;comment:外观描述"`
	Origin      string `json:"origin" gorm:"type:text;comment:来源背景"`
	History     string `json:"history" gorm:"type:text;comment:历史记录"`
	
	// 位置和所有者
	Owner       string `json:"owner" gorm:"size:200;comment:当前拥有者"`
	Location    string `json:"location" gorm:"size:200;comment:当前位置"`
	PreviousOwners string `json:"previousOwners" gorm:"type:text;comment:历史拥有者"`
	
	// 状态信息
	Status      string `json:"status" gorm:"size:20;default:available;comment:物品状态"`
	Condition   string `json:"condition" gorm:"size:20;default:good;comment:物品状况"`
	Importance  int    `json:"importance" gorm:"default:1;comment:重要程度1-10"`
	
	// 功能信息
	Abilities   string `json:"abilities" gorm:"type:text;comment:特殊能力"`
	Limitations string `json:"limitations" gorm:"type:text;comment:使用限制"`
	SideEffects string `json:"sideEffects" gorm:"type:text;comment:副作用"`
	
	// 出现信息
	FirstAppearance int    `json:"firstAppearance" gorm:"comment:首次出现章节"`
	LastAppearance  int    `json:"lastAppearance" gorm:"comment:最后出现章节"`
	
	// 标签和分类
	Tags        string `json:"tags" gorm:"type:text;comment:标签"`
	SubCategory string `json:"subCategory" gorm:"size:50;comment:子分类"`
	
	// 关联信息
	NovelID     uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy   uint   `json:"createdBy" gorm:"comment:创建者ID"`
}

// TableName 设置表名
func (NovelItem) TableName() string {
	return "novel_items"
}

// ItemCategory 物品分类常量
const (
	ItemCategoryWeapon     = "weapon"     // 武器
	ItemCategoryArmor      = "armor"      // 防具
	ItemCategoryTool       = "tool"       // 工具
	ItemCategoryMagic      = "magic"      // 魔法物品
	ItemCategoryArtifact   = "artifact"   // 神器
	ItemCategoryConsumable = "consumable" // 消耗品
	ItemCategoryJewelry    = "jewelry"    // 首饰
	ItemCategoryDocument   = "document"   // 文档
	ItemCategoryKey        = "key"        // 钥匙
	ItemCategoryTreasure   = "treasure"   // 宝物
	ItemCategoryOrdinary   = "ordinary"   // 普通物品
)

// ItemStatus 物品状态常量
const (
	ItemStatusAvailable = "available" // 可用
	ItemStatusInUse     = "in_use"    // 使用中
	ItemStatusLost      = "lost"      // 丢失
	ItemStatusDestroyed = "destroyed" // 损毁
	ItemStatusStolen    = "stolen"    // 被盗
	ItemStatusHidden    = "hidden"    // 隐藏
	ItemStatusSealed    = "sealed"    // 封印
)

// ItemCondition 物品状况常量
const (
	ItemConditionPerfect  = "perfect"  // 完美
	ItemConditionGood     = "good"     // 良好
	ItemConditionFair     = "fair"     // 一般
	ItemConditionPoor     = "poor"     // 较差
	ItemConditionBroken   = "broken"   // 损坏
	ItemConditionRuined   = "ruined"   // 毁坏
)