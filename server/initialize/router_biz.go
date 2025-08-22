package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}

func initBizRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	novelRouter := router.RouterGroupApp.Novel
	{
		// 代码生成器历史 开始
		novelRouter.NovelRouter.InitNovelRouter(Router, PublicRouter)
		novelRouter.AIRouter.InitAIRouter(Router, PublicRouter)
		// 代码生成器历史 结束
	}
}
