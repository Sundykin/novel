package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelScene 场景模型
type NovelScene struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	Name        string `json:"name" gorm:"size:200;not null;comment:场景名称"`
	Location    string `json:"location" gorm:"size:200;not null;comment:场景位置"`
	Description string `json:"description" gorm:"type:text;not null;comment:场景描述"`
	
	// 详细信息
	Atmosphere      string `json:"atmosphere" gorm:"type:text;comment:氛围描述"`
	KeyFeatures     string `json:"keyFeatures" gorm:"type:text;comment:关键特征"`
	PhysicalLayout  string `json:"physicalLayout" gorm:"type:text;comment:物理布局"`
	Environment     string `json:"environment" gorm:"type:text;comment:环境描述"`
	
	// 关联信息
	RelatedCharacters string `json:"relatedCharacters" gorm:"type:text;comment:相关角色"`
	RelatedEvents     string `json:"relatedEvents" gorm:"type:text;comment:相关事件"`
	RelatedItems      string `json:"relatedItems" gorm:"type:text;comment:相关物品"`
	
	// 时间信息
	TimeOfDay       string `json:"timeOfDay" gorm:"size:50;comment:时间段"`
	Season          string `json:"season" gorm:"size:20;comment:季节"`
	Weather         string `json:"weather" gorm:"size:50;comment:天气"`
	
	// 出现信息
	FirstAppearance int    `json:"firstAppearance" gorm:"comment:首次出现章节"`
	LastAppearance  int    `json:"lastAppearance" gorm:"comment:最后出现章节"`
	AppearanceCount int    `json:"appearanceCount" gorm:"default:0;comment:出现次数"`
	
	// 分类信息
	Category        string `json:"category" gorm:"size:50;comment:场景分类"`
	SubCategory     string `json:"subCategory" gorm:"size:50;comment:子分类"`
	Tags            string `json:"tags" gorm:"type:text;comment:标签"`
	
	// 状态信息
	Status          string `json:"status" gorm:"size:20;default:active;comment:场景状态"`
	Importance      int    `json:"importance" gorm:"default:1;comment:重要程度1-10"`
	
	// 地理信息
	Region          string `json:"region" gorm:"size:100;comment:所属区域"`
	Country         string `json:"country" gorm:"size:100;comment:所属国家"`
	City            string `json:"city" gorm:"size:100;comment:所属城市"`
	
	// 关联信息
	NovelID         uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy       uint   `json:"createdBy" gorm:"comment:创建者ID"`
}

// TableName 设置表名
func (NovelScene) TableName() string {
	return "novel_scenes"
}

// SceneCategory 场景分类常量
const (
	SceneCategoryIndoor    = "indoor"    // 室内
	SceneCategoryOutdoor   = "outdoor"   // 户外
	SceneCategoryNatural   = "natural"   // 自然
	SceneCategoryUrban     = "urban"     // 城市
	SceneCategoryRural     = "rural"     // 乡村
	SceneCategoryMystic    = "mystic"    // 神秘
	SceneCategoryBattle    = "battle"    // 战斗
	SceneCategorySocial    = "social"    // 社交
	SceneCategoryPrivate   = "private"   // 私人
	SceneCategoryPublic    = "public"    // 公共
)

// SceneStatus 场景状态常量
const (
	SceneStatusActive    = "active"    // 活跃
	SceneStatusInactive  = "inactive"  // 非活跃
	SceneStatusDestroyed = "destroyed" // 已毁坏
	SceneStatusAbandoned = "abandoned" // 已废弃
	SceneStatusHidden    = "hidden"    // 隐藏
	SceneStatusArchived  = "archived"  // 已归档
)

// TimeOfDay 时间段常量
const (
	TimeOfDayDawn     = "dawn"     // 黎明
	TimeOfDayMorning  = "morning"  // 上午
	TimeOfDayNoon     = "noon"     // 中午
	TimeOfDayAfternoon = "afternoon" // 下午
	TimeOfDayEvening  = "evening"  // 傍晚
	TimeOfDayNight    = "night"    // 夜晚
	TimeOfDayMidnight = "midnight" // 午夜
)

// Season 季节常量
const (
	SeasonSpring = "spring" // 春季
	SeasonSummer = "summer" // 夏季
	SeasonAutumn = "autumn" // 秋季
	SeasonWinter = "winter" // 冬季
)