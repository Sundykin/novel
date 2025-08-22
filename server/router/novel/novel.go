package novel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NovelRouter struct{}

// InitNovelRouter 初始化小说路由信息
func (s *NovelRouter) InitNovelRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	novelRouter := Router.Group("novel").Use(middleware.OperationRecord())
	novelRouterWithoutRecord := Router.Group("novel")
	novelRouterWithoutAuth := RouterPub.Group("novel")
	
	var novelApi = v1.ApiGroupApp.NovelApiGroup.NovelApi
	{
		novelRouter.POST("createNovel", novelApi.CreateNovel)             // 新建小说
		novelRouter.DELETE("deleteNovel", novelApi.DeleteNovel)           // 删除小说
		novelRouter.DELETE("deleteNovelByIds", novelApi.DeleteNovelByIds) // 批量删除小说
		novelRouter.PUT("updateNovel", novelApi.UpdateNovel)              // 更新小说
		novelRouter.PUT("updateNovelStatus", novelApi.UpdateNovelStatus)  // 更新小说状态
	}
	{
		novelRouterWithoutRecord.GET("findNovel", novelApi.FindNovel)                   // 根据ID获取小说
		novelRouterWithoutRecord.GET("getNovelList", novelApi.GetNovelList)             // 获取小说列表
		novelRouterWithoutRecord.GET("getNovelStatistics", novelApi.GetNovelStatistics) // 获取小说统计信息
	}
	{
		novelRouterWithoutAuth.GET("getNovelPublic", novelApi.GetNovelPublic) // 获取公开小说列表
	}
}

	// 角色路由
	var characterApi = v1.ApiGroupApp.NovelApiGroup.CharacterApi
	{
		novelRouter.POST("createCharacter", characterApi.CreateCharacter)
		novelRouter.DELETE("deleteCharacter", characterApi.DeleteCharacter)
		novelRouter.DELETE("deleteCharacterByIds", characterApi.DeleteCharacterByIds)
		novelRouter.PUT("updateCharacter", characterApi.UpdateCharacter)
	}
	{
		novelRouterWithoutRecord.GET("findCharacter", characterApi.FindCharacter)
		novelRouterWithoutRecord.GET("getCharacterList", characterApi.GetCharacterList)
	}

	// 章节路由
	var chapterApi = v1.ApiGroupApp.NovelApiGroup.ChapterApi
	{
		novelRouter.POST("createChapter", chapterApi.CreateChapter)
		novelRouter.DELETE("deleteChapter", chapterApi.DeleteChapter)
		novelRouter.DELETE("deleteChapterByIds", chapterApi.DeleteChapterByIds)
		novelRouter.PUT("updateChapter", chapterApi.UpdateChapter)
	}
	{
		novelRouterWithoutRecord.GET("findChapter", chapterApi.FindChapter)
		novelRouterWithoutRecord.GET("getChapterList", chapterApi.GetChapterList)
	}

	// 伏笔路由
	var foreshadowingApi = v1.ApiGroupApp.NovelApiGroup.ForeshadowingApi
	{
		novelRouter.POST("createForeshadowing", foreshadowingApi.CreateForeshadowing)
		novelRouter.DELETE("deleteForeshadowing", foreshadowingApi.DeleteForeshadowing)
		novelRouter.DELETE("deleteForeshadowingByIds", foreshadowingApi.DeleteForeshadowingByIds)
		novelRouter.PUT("updateForeshadowing", foreshadowingApi.UpdateForeshadowing)
	}
	{
		novelRouterWithoutRecord.GET("findForeshadowing", foreshadowingApi.FindForeshadowing)
		novelRouterWithoutRecord.GET("getForeshadowingList", foreshadowingApi.GetForeshadowingList)
	}

	// 大纲路由
	var outlineApi = v1.ApiGroupApp.NovelApiGroup.OutlineApi
	{
		novelRouter.POST("createOutline", outlineApi.CreateOutline)
		novelRouter.DELETE("deleteOutline", outlineApi.DeleteOutline)
		novelRouter.DELETE("deleteOutlineByIds", outlineApi.DeleteOutlineByIds)
		novelRouter.PUT("updateOutline", outlineApi.UpdateOutline)
	}
	{
		novelRouterWithoutRecord.GET("findOutline", outlineApi.FindOutline)
		novelRouterWithoutRecord.GET("getOutlineList", outlineApi.GetOutlineList)
	}

	// AI路由
	var aiApi = v1.ApiGroupApp.NovelApiGroup.AIApi
	{
		novelRouter.POST("chat", aiApi.ChatWithAI)
		novelRouter.POST("generate", aiApi.GenerateContent)
		novelRouter.POST("feedback", aiApi.SubmitFeedback)
		novelRouter.POST("stats", aiApi.GetUsageStats)
		novelRouter.POST("writing-assist", aiApi.WritingAssist)
		novelRouter.POST("plot-generate", aiApi.PlotGenerate)
		novelRouter.POST("character-generate", aiApi.CharacterGenerate)
		novelRouter.POST("dialogue-generate", aiApi.DialogueGenerate)
		novelRouter.POST("continue", aiApi.ContinueWriting)
		novelRouter.POST("createAIConfig", aiApi.CreateAIConfig)
		novelRouter.DELETE("deleteAIConfig", aiApi.DeleteAIConfig)
		novelRouter.DELETE("deleteAIConfigByIds", aiApi.DeleteAIConfigByIds)
		novelRouter.PUT("updateAIConfig", aiApi.UpdateAIConfig)
		novelRouter.POST("testAIConfig", aiApi.TestAIConfig)
	}
	{
		novelRouterWithoutRecord.GET("history", aiApi.GetChatHistory)
		novelRouterWithoutRecord.GET("findAIConfig", aiApi.FindAIConfig)
		novelRouterWithoutRecord.GET("getAIConfigList", aiApi.GetAIConfigList)
	}