package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelCharacter 小说角色模型
type NovelCharacter struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	Name        string `json:"name" gorm:"size:100;not null;comment:角色姓名"`
	Role        string `json:"role" gorm:"size:20;not null;comment:角色类型;default:supporting"`
	Description string `json:"description" gorm:"type:text;comment:角色描述"`
	
	// 详细信息
	Appearance   string `json:"appearance" gorm:"type:text;comment:外貌描述"`
	Personality  string `json:"personality" gorm:"type:text;comment:性格描述"`
	Background   string `json:"background" gorm:"type:text;comment:背景故事"`
	Abilities    string `json:"abilities" gorm:"type:text;comment:能力技能"`
	
	// 状态信息
	Status       string `json:"status" gorm:"size:20;default:active;comment:状态"`
	Importance   int    `json:"importance" gorm:"default:1;comment:重要程度1-10"`
	
	// 关联信息
	NovelID      uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy    uint   `json:"createdBy" gorm:"comment:创建者ID"`
	
	// 关联关系
	RelationshipsA []NovelRelationship `json:"relationshipsA" gorm:"foreignKey:CharacterAID"`
	RelationshipsB []NovelRelationship `json:"relationshipsB" gorm:"foreignKey:CharacterBID"`
}

// TableName 设置表名
func (NovelCharacter) TableName() string {
	return "novel_characters"
}

// CharacterRole 角色类型常量
const (
	CharacterRoleProtagonist = "protagonist" // 主角
	CharacterRoleAntagonist  = "antagonist"  // 反派
	CharacterRoleSupporting  = "supporting"  // 配角
	CharacterRoleMinor       = "minor"       // 次要角色
)

// CharacterStatus 角色状态常量
const (
	CharacterStatusActive   = "active"   // 活跃
	CharacterStatusInactive = "inactive" // 非活跃
	CharacterStatusDead     = "dead"     // 死亡
	CharacterStatusMissing  = "missing"  // 失踪
)