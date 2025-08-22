package novel

import (
	"errors"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel/request"
	"gorm.io/gorm"
)

type NovelService struct{}

// CreateNovel 创建小说
func (novelService *NovelService) CreateNovel(novelInfo *novel.Novel) (err error) {
	err = global.GVA_DB.Create(novelInfo).Error
	return err
}

// DeleteNovel 删除小说
func (novelService *NovelService) DeleteNovel(ID string, userID uint) (err error) {
	id, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}
	
	// 检查权限
	var novelInfo novel.Novel
	err = global.GVA_DB.Where("id = ? AND created_by = ?", id, userID).First(&novelInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("小说不存在或无权限删除")
		}
		return err
	}
	
	// 软删除
	err = global.GVA_DB.Delete(&novelInfo).Error
	return err
}

// DeleteNovelByIds 批量删除小说
func (novelService *NovelService) DeleteNovelByIds(IDs []string, userID uint) (err error) {
	var ids []int
	for _, id := range IDs {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		ids = append(ids, intID)
	}
	
	// 检查权限
	var count int64
	err = global.GVA_DB.Model(&novel.Novel{}).Where("id IN ? AND created_by = ?", ids, userID).Count(&count).Error
	if err != nil {
		return err
	}
	if int(count) != len(ids) {
		return errors.New("部分小说不存在或无权限删除")
	}
	
	// 批量软删除
	err = global.GVA_DB.Where("id IN ? AND created_by = ?", ids, userID).Delete(&novel.Novel{}).Error
	return err
}

// UpdateNovel 更新小说
func (novelService *NovelService) UpdateNovel(novelInfo novel.Novel, userID uint) (err error) {
	// 检查权限
	var existingNovel novel.Novel
	err = global.GVA_DB.Where("id = ? AND created_by = ?", novelInfo.ID, userID).First(&existingNovel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("小说不存在或无权限修改")
		}
		return err
	}
	
	// 更新
	err = global.GVA_DB.Model(&existingNovel).Updates(novelInfo).Error
	return err
}

// GetNovel 根据ID获取小说
func (novelService *NovelService) GetNovel(ID string, userID uint) (novelInfo novel.Novel, err error) {
	id, err := strconv.Atoi(ID)
	if err != nil {
		return novelInfo, err
	}
	
	err = global.GVA_DB.Where("id = ? AND created_by = ?", id, userID).Preload("Characters").Preload("Outlines").Preload("Chapters").Preload("Foreshadowings").Preload("Worldbuildings").Preload("Items").Preload("Scenes").First(&novelInfo).Error
	return
}

// GetNovelInfoList 分页获取小说列表
func (novelService *NovelService) GetNovelInfoList(info request.NovelSearch) (list []novel.Novel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	
	// 构建查询条件
	db := global.GVA_DB.Model(&novel.Novel{})
	
	// 用户权限过滤
	if info.CreatedBy != 0 {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	
	// 关键词搜索
	if info.Keyword != "" {
		db = db.Where("title LIKE ? OR description LIKE ? OR synopsis LIKE ?", "%"+info.Keyword+"%", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
	}
	
	// 类型筛选
	if info.Genre != "" {
		db = db.Where("genre = ?", info.Genre)
	}
	
	// 状态筛选
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	
	// 时间范围筛选
	if info.StartCreatedAt != nil && *info.StartCreatedAt != "" {
		db = db.Where("created_at >= ?", *info.StartCreatedAt)
	}
	if info.EndCreatedAt != nil && *info.EndCreatedAt != "" {
		db = db.Where("created_at <= ?", *info.EndCreatedAt)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Order("updated_at DESC").Find(&list).Error
	return list, total, err
}

// GetNovelPublicList 获取公开小说列表
func (novelService *NovelService) GetNovelPublicList(info request.NovelSearch) (list []novel.Novel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	
	// 构建查询条件
	db := global.GVA_DB.Model(&novel.Novel{}).Where("is_public = ?", true)
	
	// 关键词搜索
	if info.Keyword != "" {
		db = db.Where("title LIKE ? OR description LIKE ?", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
	}
	
	// 类型筛选
	if info.Genre != "" {
		db = db.Where("genre = ?", info.Genre)
	}
	
	// 状态筛选
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	
	err = db.Limit(limit).Offset(offset).Order("updated_at DESC").Find(&list).Error
	return list, total, err
}

// UpdateNovelStatus 更新小说状态
func (novelService *NovelService) UpdateNovelStatus(ID string, status string, userID uint) (err error) {
	id, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}
	
	// 检查权限
	var novelInfo novel.Novel
	err = global.GVA_DB.Where("id = ? AND created_by = ?", id, userID).First(&novelInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("小说不存在或无权限修改")
		}
		return err
	}
	
	// 更新状态
	err = global.GVA_DB.Model(&novelInfo).Update("status", status).Error
	return err
}

// GetNovelStatistics 获取小说统计信息
func (novelService *NovelService) GetNovelStatistics(ID string, userID uint) (statistics map[string]interface{}, err error) {
	id, err := strconv.Atoi(ID)
	if err != nil {
		return nil, err
	}
	
	// 检查权限
	var novelInfo novel.Novel
	err = global.GVA_DB.Where("id = ? AND created_by = ?", id, userID).First(&novelInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("小说不存在或无权限查看")
		}
		return nil, err
	}
	
	statistics = make(map[string]interface{})
	
	// 基本统计
	statistics["novel"] = novelInfo
	
	// 角色统计
	var characterCount int64
	global.GVA_DB.Model(&novel.NovelCharacter{}).Where("novel_id = ?", id).Count(&characterCount)
	statistics["characterCount"] = characterCount
	
	// 章节统计
	var chapterCount int64
	var totalWords int64
	global.GVA_DB.Model(&novel.NovelChapter{}).Where("novel_id = ?", id).Count(&chapterCount)
	global.GVA_DB.Model(&novel.NovelChapter{}).Where("novel_id = ?", id).Select("COALESCE(SUM(word_count), 0)").Scan(&totalWords)
	statistics["chapterCount"] = chapterCount
	statistics["totalWords"] = totalWords
	
	// 伏笔统计
	var foreshadowingCount int64
	var plantedCount, revealedCount int64
	global.GVA_DB.Model(&novel.NovelForeshadowing{}).Where("novel_id = ?", id).Count(&foreshadowingCount)
	global.GVA_DB.Model(&novel.NovelForeshadowing{}).Where("novel_id = ? AND status = ?", id, "planted").Count(&plantedCount)
	global.GVA_DB.Model(&novel.NovelForeshadowing{}).Where("novel_id = ? AND status = ?", id, "revealed").Count(&revealedCount)
	statistics["foreshadowingCount"] = foreshadowingCount
	statistics["plantedCount"] = plantedCount
	statistics["revealedCount"] = revealedCount
	
	// 世界观统计
	var worldbuildingCount int64
	global.GVA_DB.Model(&novel.NovelWorldbuilding{}).Where("novel_id = ?", id).Count(&worldbuildingCount)
	statistics["worldbuildingCount"] = worldbuildingCount
	
	// 物品统计
	var itemCount int64
	global.GVA_DB.Model(&novel.NovelItem{}).Where("novel_id = ?", id).Count(&itemCount)
	statistics["itemCount"] = itemCount
	
	// 场景统计
	var sceneCount int64
	global.GVA_DB.Model(&novel.NovelScene{}).Where("novel_id = ?", id).Count(&sceneCount)
	statistics["sceneCount"] = sceneCount
	
	return statistics, nil
}