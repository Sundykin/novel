import service from '@/utils/request'

// @Tags Novel
// @Summary 创建小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Novel true "小说信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /novel/createNovel [post]
export const createNovel = (data) => {
  return service({
    url: '/novel/createNovel',
    method: 'post',
    data
  })
}

// @Tags Novel
// @Summary 删除小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Novel true "小说信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /novel/deleteNovel [delete]
export const deleteNovel = (params) => {
  return service({
    url: '/novel/deleteNovel',
    method: 'delete',
    params
  })
}

// @Tags Novel
// @Summary 批量删除小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除小说"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /novel/deleteNovelByIds [delete]
export const deleteNovelByIds = (params) => {
  return service({
    url: '/novel/deleteNovelByIds',
    method: 'delete',
    params
  })
}

// @Tags Novel
// @Summary 更新小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Novel true "小说信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /novel/updateNovel [put]
export const updateNovel = (data) => {
  return service({
    url: '/novel/updateNovel',
    method: 'put',
    data
  })
}

// @Tags Novel
// @Summary 用id查询小说
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Novel true "用id查询小说"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /novel/findNovel [get]
export const findNovel = (params) => {
  return service({
    url: '/novel/findNovel',
    method: 'get',
    params
  })
}

// @Tags Novel
// @Summary 分页获取小说列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取小说列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /novel/getNovelList [get]
export const getNovelList = (params) => {
  return service({
    url: '/novel/getNovelList',
    method: 'get',
    params
  })
}

// @Tags Novel
// @Summary 获取公开小说列表
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公开小说列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /novel/getNovelPublic [get]
export const getNovelPublic = (params) => {
  return service({
    url: '/novel/getNovelPublic',
    method: 'get',
    params
  })
}

// @Tags Novel
// @Summary 更新小说状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /novel/updateNovelStatus [put]
export const updateNovelStatus = (params) => {
  return service({
    url: '/novel/updateNovelStatus',
    method: 'put',
    params
  })
}

// @Tags Novel
// @Summary 获取小说统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /novel/getNovelStatistics [get]
export const getNovelStatistics = (params) => {
  return service({
    url: '/novel/getNovelStatistics',
    method: 'get',
    params
  })
}

// ==================== 角色管理 API ====================

// @Tags Character
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelCharacter true "角色信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /character/createCharacter [post]
export const createCharacter = (data) => {
  return service({
    url: '/novel/createCharacter',
    method: 'post',
    data
  })
}

// @Tags Character
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelCharacter true "角色信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /character/deleteCharacter [delete]
export const deleteCharacter = (params) => {
  return service({
    url: '/novel/deleteCharacter',
    method: 'delete',
    params
  })
}

// @Tags Character
// @Summary 批量删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /character/deleteCharacterByIds [delete]
export const deleteCharacterByIds = (params) => {
  return service({
    url: '/novel/deleteCharacterByIds',
    method: 'delete',
    params
  })
}

// @Tags Character
// @Summary 更新角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelCharacter true "角色信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /character/updateCharacter [put]
export const updateCharacter = (data) => {
  return service({
    url: '/novel/updateCharacter',
    method: 'put',
    data
  })
}

// @Tags Character
// @Summary 用id查询角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NovelCharacter true "用id查询角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /character/findCharacter [get]
export const findCharacter = (params) => {
  return service({
    url: '/novel/findCharacter',
    method: 'get',
    params
  })
}

// @Tags Character
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取角色列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /character/getCharacterList [get]
export const getCharacterList = (params) => {
  return service({
    url: '/novel/getCharacterList',
    method: 'get',
    params
  })
}

// ==================== 伏笔管理 API ====================

// @Tags Foreshadowing
// @Summary 创建伏笔
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelForeshadowing true "伏笔信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /foreshadowing/createForeshadowing [post]
export const createForeshadowing = (data) => {
  return service({
    url: '/novel/createForeshadowing',
    method: 'post',
    data
  })
}

// @Tags Foreshadowing
// @Summary 删除伏笔
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelForeshadowing true "伏笔信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /foreshadowing/deleteForeshadowing [delete]
export const deleteForeshadowing = (params) => {
  return service({
    url: '/novel/deleteForeshadowing',
    method: 'delete',
    params
  })
}

// @Tags Foreshadowing
// @Summary 批量删除伏笔
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除伏笔"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /foreshadowing/deleteForeshadowingByIds [delete]
export const deleteForeshadowingByIds = (params) => {
  return service({
    url: '/novel/deleteForeshadowingByIds',
    method: 'delete',
    params
  })
}

// @Tags Foreshadowing
// @Summary 更新伏笔
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelForeshadowing true "伏笔信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /foreshadowing/updateForeshadowing [put]
export const updateForeshadowing = (data) => {
  return service({
    url: '/novel/updateForeshadowing',
    method: 'put',
    data
  })
}

// @Tags Foreshadowing
// @Summary 用id查询伏笔
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NovelForeshadowing true "用id查询伏笔"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /foreshadowing/findForeshadowing [get]
export const findForeshadowing = (params) => {
  return service({
    url: '/novel/findForeshadowing',
    method: 'get',
    params
  })
}

// @Tags Foreshadowing
// @Summary 分页获取伏笔列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取伏笔列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /foreshadowing/getForeshadowingList [get]
export const getForeshadowingList = (params) => {
  return service({
    url: '/novel/getForeshadowingList',
    method: 'get',
    params
  })
}

// @Tags Foreshadowing
// @Summary 更新伏笔状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /foreshadowing/updateForeshadowingStatus [put]
export const updateForeshadowingStatus = (params) => {
  return service({
    url: '/novel/updateForeshadowingStatus',
    method: 'put',
    params
  })
}

// ==================== 大纲管理 API ====================

// @Tags Outline
// @Summary 创建大纲
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelOutline true "大纲信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /outline/createOutline [post]
export const createOutline = (data) => {
  return service({
    url: '/novel/createOutline',
    method: 'post',
    data
  })
}

// @Tags Outline
// @Summary 删除大纲
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelOutline true "大纲信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outline/deleteOutline [delete]
export const deleteOutline = (params) => {
  return service({
    url: '/novel/deleteOutline',
    method: 'delete',
    params
  })
}

// @Tags Outline
// @Summary 批量删除大纲
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除大纲"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outline/deleteOutlineByIds [delete]
export const deleteOutlineByIds = (params) => {
  return service({
    url: '/novel/deleteOutlineByIds',
    method: 'delete',
    params
  })
}

// @Tags Outline
// @Summary 更新大纲
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelOutline true "大纲信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /outline/updateOutline [put]
export const updateOutline = (data) => {
  return service({
    url: '/novel/updateOutline',
    method: 'put',
    data
  })
}

// @Tags Outline
// @Summary 用id查询大纲
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NovelOutline true "用id查询大纲"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /outline/findOutline [get]
export const findOutline = (params) => {
  return service({
    url: '/novel/findOutline',
    method: 'get',
    params
  })
}

// @Tags Outline
// @Summary 分页获取大纲列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取大纲列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outline/getOutlineList [get]
export const getOutlineList = (params) => {
  return service({
    url: '/novel/getOutlineList',
    method: 'get',
    params
  })
}

// ==================== 章节管理 API ====================

// @Tags Chapter
// @Summary 创建章节
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelChapter true "章节信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chapter/createChapter [post]
export const createChapter = (data) => {
  return service({
    url: '/novel/createChapter',
    method: 'post',
    data
  })
}

// @Tags Chapter
// @Summary 删除章节
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelChapter true "章节信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chapter/deleteChapter [delete]
export const deleteChapter = (params) => {
  return service({
    url: '/novel/deleteChapter',
    method: 'delete',
    params
  })
}

// @Tags Chapter
// @Summary 批量删除章节
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除章节"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chapter/deleteChapterByIds [delete]
export const deleteChapterByIds = (params) => {
  return service({
    url: '/novel/deleteChapterByIds',
    method: 'delete',
    params
  })
}

// @Tags Chapter
// @Summary 更新章节
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NovelChapter true "章节信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chapter/updateChapter [put]
export const updateChapter = (data) => {
  return service({
    url: '/novel/updateChapter',
    method: 'put',
    data
  })
}

// @Tags Chapter
// @Summary 用id查询章节
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NovelChapter true "用id查询章节"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chapter/findChapter [get]
export const findChapter = (params) => {
  return service({
    url: '/novel/findChapter',
    method: 'get',
    params
  })
}

// @Tags Chapter
// @Summary 分页获取章节列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取章节列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chapter/getChapterList [get]
export const getChapterList = (params) => {
  return service({
    url: '/novel/getChapterList',
    method: 'get',
    params
  })
}

// @Tags Chapter
// @Summary 更新章节状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chapter/updateChapterStatus [put]
export const updateChapterStatus = (params) => {
  return service({
    url: '/novel/updateChapterStatus',
    method: 'put',
    params
  })
}

// AI助手相关API
export const chatWithAI = (data) => {
  return service({
    url: '/novel/chat',
    method: 'post',
    data
  })
}

export const generateContent = (data) => {
  return service({
    url: '/novel/generate',
    method: 'post',
    data
  })
}

export const getChatHistory = (params) => {
  return service({
    url: '/novel/history',
    method: 'get',
    params
  })
}

export const submitFeedback = (data) => {
  return service({
    url: '/novel/feedback',
    method: 'post',
    data
  })
}

export const getUsageStats = (data) => {
  return service({
    url: '/novel/stats',
    method: 'post',
    data
  })
}

export const writingAssist = (data) => {
  return service({
    url: '/novel/writing-assist',
    method: 'post',
    data
  })
}

export const plotGenerate = (data) => {
  return service({
    url: '/novel/plot-generate',
    method: 'post',
    data
  })
}

export const characterGenerate = (data) => {
  return service({
    url: '/novel/character-generate',
    method: 'post',
    data
  })
}

export const dialogueGenerate = (data) => {
  return service({
    url: '/novel/dialogue-generate',
    method: 'post',
    data
  })
}

export const continueWriting = (data) => {
  return service({
    url: '/novel/continue',
    method: 'post',
    data
  })
}

export const outlineGenerate = (data) => {
  return service({
    url: '/ai/outline-generate',
    method: 'post',
    data
  })
}

export const reviewContent = (data) => {
  return service({
    url: '/ai/review',
    method: 'post',
    data
  })
}

export const optimizeContent = (data) => {
  return service({
    url: '/ai/optimize',
    method: 'post',
    data
  })
}

export const translateContent = (data) => {
  return service({
    url: '/ai/translate',
    method: 'post',
    data
  })
}

export const analyzeStyle = (data) => {
  return service({
    url: '/ai/style-analyze',
    method: 'post',
    data
  })
}

// ==================== AI配置管理 API ====================

export const createAIConfig = (data) => {
  return service({
    url: '/novel/createAIConfig',
    method: 'post',
    data
  })
}

export const deleteAIConfig = (params) => {
  return service({
    url: '/novel/deleteAIConfig',
    method: 'delete',
    params
  })
}

export const deleteAIConfigByIds = (params) => {
  return service({
    url: '/novel/deleteAIConfigByIds',
    method: 'delete',
    params
  })
}

export const updateAIConfig = (data) => {
  return service({
    url: '/novel/updateAIConfig',
    method: 'put',
    data
  })
}

export const findAIConfig = (params) => {
  return service({
    url: '/novel/findAIConfig',
    method: 'get',
    params
  })
}

export const getAIConfigList = (params) => {
  return service({
    url: '/novel/getAIConfigList',
    method: 'get',
    params
  })
}

export const testAIConfig = (data) => {
  return service({
    url: '/novel/testAIConfig',
    method: 'post',
    data
  })
}