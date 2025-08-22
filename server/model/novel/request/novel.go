package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel"
)

// NovelSearch 小说搜索结构体
type NovelSearch struct {
	novel.Novel
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword        string  `json:"keyword" form:"keyword"`        // 关键词搜索
	Genre          string  `json:"genre" form:"genre"`            // 类型筛选
	Status         string  `json:"status" form:"status"`          // 状态筛选
	CreatedBy      uint    `json:"createdBy" form:"createdBy"`    // 创建者筛选
}

// CharacterSearch 角色搜索结构体
type CharacterSearch struct {
	novel.NovelCharacter
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword        string  `json:"keyword" form:"keyword"`        // 关键词搜索
	Role           string  `json:"role" form:"role"`              // 角色类型筛选
	Status         string  `json:"status" form:"status"`          // 状态筛选
	NovelID        uint    `json:"novelId" form:"novelId"`        // 小说ID筛选
}

// ForeshadowingSearch 伏笔搜索结构体
type ForeshadowingSearch struct {
	novel.NovelForeshadowing
	request.PageInfo
	StartCreatedAt  *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt    *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword         string  `json:"keyword" form:"keyword"`                // 关键词搜索
	ImportanceLevel string  `json:"importanceLevel" form:"importanceLevel"` // 重要程度筛选
	Status          string  `json:"status" form:"status"`                  // 状态筛选
	Type            string  `json:"type" form:"type"`                      // 类型筛选
	NovelID         uint    `json:"novelId" form:"novelId"`                // 小说ID筛选
}

// OutlineSearch 大纲搜索结构体
type OutlineSearch struct {
	novel.NovelOutline
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword        string  `json:"keyword" form:"keyword"`        // 关键词搜索
	Status         string  `json:"status" form:"status"`          // 状态筛选
	VolumeNumber   int     `json:"volumeNumber" form:"volumeNumber"` // 卷号筛选
	NovelID        uint    `json:"novelId" form:"novelId"`        // 小说ID筛选
}

// ChapterSearch 章节搜索结构体
type ChapterSearch struct {
	novel.NovelChapter
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword        string  `json:"keyword" form:"keyword"`        // 关键词搜索
	Status         string  `json:"status" form:"status"`          // 状态筛选
	Volume         int     `json:"volume" form:"volume"`          // 卷号筛选
	OutlineID      uint    `json:"outlineId" form:"outlineId"`    // 大纲ID筛选
	NovelID        uint    `json:"novelId" form:"novelId"`        // 小说ID筛选
}

// WorldbuildingSearch 世界观搜索结构体
type WorldbuildingSearch struct {
	novel.NovelWorldbuilding
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword        string  `json:"keyword" form:"keyword"`        // 关键词搜索
	Category       string  `json:"category" form:"category"`      // 分类筛选
	Status         string  `json:"status" form:"status"`          // 状态筛选
	NovelID        uint    `json:"novelId" form:"novelId"`        // 小说ID筛选
}

// ItemSearch 物品搜索结构体
type ItemSearch struct {
	novel.NovelItem
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword        string  `json:"keyword" form:"keyword"`        // 关键词搜索
	Category       string  `json:"category" form:"category"`      // 分类筛选
	Status         string  `json:"status" form:"status"`          // 状态筛选
	Owner          string  `json:"owner" form:"owner"`            // 拥有者筛选
	NovelID        uint    `json:"novelId" form:"novelId"`        // 小说ID筛选
}

// SceneSearch 场景搜索结构体
type SceneSearch struct {
	novel.NovelScene
	request.PageInfo
	StartCreatedAt *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword        string  `json:"keyword" form:"keyword"`        // 关键词搜索
	Category       string  `json:"category" form:"category"`      // 分类筛选
	Status         string  `json:"status" form:"status"`          // 状态筛选
	Location       string  `json:"location" form:"location"`      // 位置筛选
	NovelID        uint    `json:"novelId" form:"novelId"`        // 小说ID筛选
}

// RelationshipSearch 关系搜索结构体
type RelationshipSearch struct {
	novel.NovelRelationship
	request.PageInfo
	StartCreatedAt   *string `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt     *string `json:"endCreatedAt" form:"endCreatedAt"`
	Keyword          string  `json:"keyword" form:"keyword"`                  // 关键词搜索
	RelationshipType string  `json:"relationshipType" form:"relationshipType"` // 关系类型筛选
	Status           string  `json:"status" form:"status"`                    // 状态筛选
	CharacterAID     uint    `json:"characterAId" form:"characterAId"`        // 角色A筛选
	CharacterBID     uint    `json:"characterBId" form:"characterBId"`        // 角色B筛选
	NovelID          uint    `json:"novelId" form:"novelId"`                  // 小说ID筛选
}

// NovelStatisticsRequest 小说统计请求
type NovelStatisticsRequest struct {
	NovelID   uint   `json:"novelId" form:"novelId"`     // 小说ID
	StartDate string `json:"startDate" form:"startDate"` // 开始日期
	EndDate   string `json:"endDate" form:"endDate"`     // 结束日期
	Type      string `json:"type" form:"type"`           // 统计类型
}

// BatchOperationRequest 批量操作请求
type BatchOperationRequest struct {
	IDs       []uint `json:"ids" form:"ids"`             // ID列表
	Operation string `json:"operation" form:"operation"` // 操作类型
	Value     string `json:"value" form:"value"`         // 操作值
}