package novel

import (
	"time"
	"gorm.io/gorm"
)

// NovelRelationship 角色关系模型
type NovelRelationship struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关系信息
	CharacterAID     uint   `json:"characterAId" gorm:"not null;comment:角色A的ID"`
	CharacterBID     uint   `json:"characterBId" gorm:"not null;comment:角色B的ID"`
	RelationshipType string `json:"relationshipType" gorm:"size:50;not null;comment:关系类型"`
	Description      string `json:"description" gorm:"type:text;comment:关系描述"`
	
	// 关系强度和状态
	Strength         int    `json:"strength" gorm:"default:5;comment:关系强度1-10"`
	Status           string `json:"status" gorm:"size:20;default:active;comment:关系状态"`
	
	// 时间线信息
	StartChapter     *int   `json:"startChapter" gorm:"comment:关系开始章节"`
	EndChapter       *int   `json:"endChapter" gorm:"comment:关系结束章节"`
	
	// 关联信息
	NovelID          uint   `json:"novelId" gorm:"not null;comment:所属小说ID"`
	CreatedBy        uint   `json:"createdBy" gorm:"comment:创建者ID"`
	
	// 关联关系
	CharacterA       NovelCharacter `json:"characterA" gorm:"foreignKey:CharacterAID"`
	CharacterB       NovelCharacter `json:"characterB" gorm:"foreignKey:CharacterBID"`
}

// TableName 设置表名
func (NovelRelationship) TableName() string {
	return "novel_relationships"
}

// RelationshipType 关系类型常量
const (
	RelationshipTypeFamily   = "family"   // 家庭关系
	RelationshipTypeFriend   = "friend"   // 朋友关系
	RelationshipTypeEnemy    = "enemy"    // 敌对关系
	RelationshipTypeLover    = "lover"    // 恋人关系
	RelationshipTypeMentor   = "mentor"   // 师徒关系
	RelationshipTypeRival    = "rival"    // 竞争关系
	RelationshipTypeAlly     = "ally"     // 盟友关系
	RelationshipTypeColleague = "colleague" // 同事关系
)

// RelationshipStatus 关系状态常量
const (
	RelationshipStatusActive    = "active"    // 活跃
	RelationshipStatusInactive  = "inactive"  // 非活跃
	RelationshipStatusBroken    = "broken"    // 破裂
	RelationshipStatusDeveloping = "developing" // 发展中
)