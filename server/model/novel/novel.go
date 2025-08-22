package novel

import (
	"time"
	"gorm.io/gorm"
)

// Novel 小说主表模型
type Novel struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 基本信息
	Title       string `json:"title" gorm:"size:200;not null;comment:小说标题"`
	Subtitle    string `json:"subtitle" gorm:"size:200;comment:副标题"`
	Description string `json:"description" gorm:"type:text;comment:小说简介"`
	Synopsis    string `json:"synopsis" gorm:"type:text;comment:故事梗概"`
	
	// 分类信息
	Genre       string `json:"genre" gorm:"size:50;not null;comment:小说类型"`
	SubGenre    string `json:"subGenre" gorm:"size:50;comment:子类型"`
	Tags        string `json:"tags" gorm:"type:text;comment:标签"`
	Keywords    string `json:"keywords" gorm:"type:text;comment:关键词"`
	
	// 状态信息
	Status      string `json:"status" gorm:"size:20;default:draft;comment:小说状态"`
	Progress    int    `json:"progress" gorm:"default:0;comment:完成进度0-100"`
	
	// 统计信息
	TotalChapters    int `json:"totalChapters" gorm:"default:0;comment:总章节数"`
	CurrentChapter   int `json:"currentChapter" gorm:"default:0;comment:当前章节"`
	TotalWords       int `json:"totalWords" gorm:"default:0;comment:总字数"`
	TargetWords      int `json:"targetWords" gorm:"default:0;comment:目标字数"`
	
	// 时间信息
	StartDate        *time.Time `json:"startDate" gorm:"comment:开始写作时间"`
	TargetEndDate    *time.Time `json:"targetEndDate" gorm:"comment:目标完成时间"`
	ActualEndDate    *time.Time `json:"actualEndDate" gorm:"comment:实际完成时间"`
	LastUpdateDate   *time.Time `json:"lastUpdateDate" gorm:"comment:最后更新时间"`
	
	// 设置信息
	Language         string `json:"language" gorm:"size:10;default:zh;comment:语言"`
	TargetAudience   string `json:"targetAudience" gorm:"size:50;comment:目标读者"`
	ContentRating    string `json:"contentRating" gorm:"size:20;comment:内容分级"`
	
	// 创作信息
	AuthorNotes      string `json:"authorNotes" gorm:"type:text;comment:作者备注"`
	WritingStyle     string `json:"writingStyle" gorm:"type:text;comment:写作风格"`
	Themes           string `json:"themes" gorm:"type:text;comment:主题"`
	
	// 发布信息
	IsPublic         bool   `json:"isPublic" gorm:"default:false;comment:是否公开"`
	PublishPlatform  string `json:"publishPlatform" gorm:"type:text;comment:发布平台"`
	
	// 关联信息
	CreatedBy        uint   `json:"createdBy" gorm:"not null;comment:创建者ID"`
	Collaborators    string `json:"collaborators" gorm:"type:text;comment:协作者"`
	
	// 关联关系
	Characters       []NovelCharacter     `json:"characters" gorm:"foreignKey:NovelID"`
	Outlines         []NovelOutline       `json:"outlines" gorm:"foreignKey:NovelID"`
	Chapters         []NovelChapter       `json:"chapters" gorm:"foreignKey:NovelID"`
	Foreshadowings   []NovelForeshadowing `json:"foreshadowings" gorm:"foreignKey:NovelID"`
	Worldbuildings   []NovelWorldbuilding `json:"worldbuildings" gorm:"foreignKey:NovelID"`
	Items            []NovelItem          `json:"items" gorm:"foreignKey:NovelID"`
	Scenes           []NovelScene         `json:"scenes" gorm:"foreignKey:NovelID"`
}

// TableName 设置表名
func (Novel) TableName() string {
	return "novels"
}

// NovelGenre 小说类型常量
const (
	NovelGenreFantasy     = "fantasy"     // 奇幻
	NovelGenreSciFi       = "sci_fi"      // 科幻
	NovelGenreRomance     = "romance"     // 言情
	NovelGenreMystery     = "mystery"     // 悬疑
	NovelGenreThriller    = "thriller"    // 惊悚
	NovelGenreHistorical  = "historical"  // 历史
	NovelGenreContemporary = "contemporary" // 现代
	NovelGenreUrban       = "urban"       // 都市
	NovelGenreWuxia       = "wuxia"       // 武侠
	NovelGenreXianxia     = "xianxia"     // 仙侠
	NovelGenreGameWorld   = "game_world"  // 游戏
	NovelGenreSliceOfLife = "slice_of_life" // 日常
)

// NovelStatus 小说状态常量
const (
	NovelStatusDraft      = "draft"      // 草稿
	NovelStatusPlanning   = "planning"   // 规划中
	NovelStatusWriting    = "writing"    // 写作中
	NovelStatusRevision   = "revision"   // 修订中
	NovelStatusCompleted  = "completed"  // 已完成
	NovelStatusPublished  = "published"  // 已发布
	NovelStatusOnHold     = "on_hold"    // 暂停
	NovelStatusCancelled  = "cancelled"  // 已取消
	NovelStatusArchived   = "archived"   // 已归档
)

// ContentRating 内容分级常量
const (
	ContentRatingG  = "G"  // 全年龄
	ContentRatingPG = "PG" // 建议家长指导
	ContentRatingPG13 = "PG13" // 13岁以上
	ContentRatingR  = "R"  // 限制级
	ContentRatingNC17 = "NC17" // 17岁以下禁止
)