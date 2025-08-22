package novel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AIRouter struct{}

// InitAIRouter 初始化AI路由信息
func (s *AIRouter) InitAIRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	aiApi := v1.ApiGroupApp.NovelApiGroup.AIApi
	{
		aiRouter := Router.Group("ai").Use(middleware.OperationRecord())
		aiRouter.POST("chat", aiApi.ChatWithAI)                     // AI聊天对话
		aiRouter.POST("generate", aiApi.GenerateContent)             // AI内容生成
		aiRouter.GET("history", aiApi.GetChatHistory)               // 获取聊天历史
		aiRouter.POST("feedback", aiApi.SubmitFeedback)             // 提交反馈
		aiRouter.POST("stats", aiApi.GetUsageStats)                 // 获取使用统计
		aiRouter.POST("writing-assist", aiApi.WritingAssist)        // 写作助手
		aiRouter.POST("plot-generate", aiApi.PlotGenerate)          // 情节生成
		aiRouter.POST("character-generate", aiApi.CharacterGenerate) // 角色生成
		aiRouter.POST("dialogue-generate", aiApi.DialogueGenerate)  // 对话生成
		aiRouter.POST("continue", aiApi.ContinueWriting)            // 续写
	}
	{
		// 公开路由，不需要鉴权
		aiPublicRouter := PublicRouter.Group("ai")
		_ = aiPublicRouter // 预留公开AI接口
	}
}