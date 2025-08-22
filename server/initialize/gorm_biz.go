package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/novel"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		// 小说系统模型
		&novel.Novel{},
		&novel.NovelCharacter{},
		&novel.NovelRelationship{},
		&novel.NovelForeshadowing{},
		&novel.NovelOutline{},
		&novel.NovelChapter{},
		&novel.NovelWorldbuilding{},
		&novel.NovelItem{},
		&novel.NovelScene{},
		&novel.AIAssistant{},
		&novel.AITemplate{},
		&novel.AIConfig{},
	)
	if err != nil {
		return err
	}
	return nil
}
