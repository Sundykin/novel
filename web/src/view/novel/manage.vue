<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="小说ID" prop="novelId">
          <el-input v-model="searchInfo.novelId" placeholder="请输入小说ID" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <div class="novel-manage-container">
      <!-- 小说基本信息卡片 -->
      <el-card class="novel-info-card" shadow="hover">
        <template #header>
          <div class="card-header">
            <span class="card-title">小说信息</span>
            <el-button type="primary" size="small" @click="editNovel">编辑</el-button>
          </div>
        </template>
        <div v-if="novelInfo.id" class="novel-info">
          <div class="info-row">
            <span class="label">标题：</span>
            <span class="value">{{ novelInfo.title }}</span>
          </div>
          <div class="info-row">
            <span class="label">类型：</span>
            <el-tag :type="getGenreTagType(novelInfo.genre)">{{ getGenreLabel(novelInfo.genre) }}</el-tag>
          </div>
          <div class="info-row">
            <span class="label">状态：</span>
            <el-tag :type="getStatusTagType(novelInfo.status)">{{ getStatusLabel(novelInfo.status) }}</el-tag>
          </div>
          <div class="info-row">
            <span class="label">进度：</span>
            <el-progress :percentage="novelInfo.progress" :stroke-width="8" />
          </div>
          <div class="info-row">
            <span class="label">字数：</span>
            <span class="value">{{ novelInfo.totalWords }} / {{ novelInfo.targetWords }}</span>
          </div>
          <div class="info-row">
            <span class="label">简介：</span>
            <span class="value description">{{ novelInfo.description }}</span>
          </div>
        </div>
        <div v-else class="no-data">
          <el-empty description="请选择一个小说进行管理" />
        </div>
      </el-card>

      <!-- 功能模块导航 -->
      <div class="module-nav">
        <el-card shadow="hover">
          <template #header>
            <span class="card-title">管理模块</span>
          </template>
          <div class="nav-grid">
            <div class="nav-item" @click="goToModule('characters')">
              <el-icon class="nav-icon"><User /></el-icon>
              <span class="nav-text">角色管理</span>
              <span class="nav-count">{{ statistics.characterCount || 0 }}</span>
            </div>
            <div class="nav-item" @click="goToModule('chapters')">
              <el-icon class="nav-icon"><Document /></el-icon>
              <span class="nav-text">章节管理</span>
              <span class="nav-count">{{ statistics.chapterCount || 0 }}</span>
            </div>
            <div class="nav-item" @click="goToModule('outlines')">
              <el-icon class="nav-icon"><List /></el-icon>
              <span class="nav-text">大纲规划</span>
              <span class="nav-count">{{ statistics.outlineCount || 0 }}</span>
            </div>
            <div class="nav-item" @click="goToModule('foreshadowing')">
              <el-icon class="nav-icon"><View /></el-icon>
              <span class="nav-text">伏笔管理</span>
              <span class="nav-count">{{ statistics.foreshadowingCount || 0 }}</span>
            </div>
            <div class="nav-item" @click="goToModule('worldbuilding')">
              <el-icon class="nav-icon"><Monitor /></el-icon>
              <span class="nav-text">世界观</span>
              <span class="nav-count">{{ statistics.worldbuildingCount || 0 }}</span>
            </div>
            <div class="nav-item" @click="goToModule('items')">
              <el-icon class="nav-icon"><Box /></el-icon>
              <span class="nav-text">物品道具</span>
              <span class="nav-count">{{ statistics.itemCount || 0 }}</span>
            </div>
            <div class="nav-item" @click="goToModule('scenes')">
              <el-icon class="nav-icon"><Location /></el-icon>
              <span class="nav-text">场景管理</span>
              <span class="nav-count">{{ statistics.sceneCount || 0 }}</span>
            </div>
            <div class="nav-item" @click="goToModule('relationships')">
              <el-icon class="nav-icon"><Connection /></el-icon>
              <span class="nav-text">关系网络</span>
              <span class="nav-count">{{ statistics.relationshipCount || 0 }}</span>
            </div>
          </div>
        </el-card>
      </div>

      <!-- 统计信息 -->
      <el-card class="statistics-card" shadow="hover">
        <template #header>
          <span class="card-title">统计信息</span>
        </template>
        <div class="statistics-grid">
          <div class="stat-item">
            <div class="stat-number">{{ statistics.totalWords || 0 }}</div>
            <div class="stat-label">总字数</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ statistics.chapterCount || 0 }}</div>
            <div class="stat-label">章节数</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ statistics.characterCount || 0 }}</div>
            <div class="stat-label">角色数</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ statistics.plantedCount || 0 }}</div>
            <div class="stat-label">已埋伏笔</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ statistics.revealedCount || 0 }}</div>
            <div class="stat-label">已揭示伏笔</div>
          </div>
          <div class="stat-item">
            <div class="stat-number">{{ novelInfo.progress || 0 }}%</div>
            <div class="stat-label">完成进度</div>
          </div>
        </div>
      </el-card>

      <!-- 最近活动 -->
      <el-card class="activity-card" shadow="hover">
        <template #header>
          <span class="card-title">最近活动</span>
        </template>
        <el-timeline>
          <el-timeline-item
            v-for="activity in recentActivities"
            :key="activity.id"
            :timestamp="activity.timestamp"
            :type="activity.type"
          >
            {{ activity.description }}
          </el-timeline-item>
        </el-timeline>
      </el-card>
    </div>
  </div>
</template>

<script setup>
import {
  findNovel,
  getNovelStatistics
} from '@/api/novel'
import { ElMessage } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  User,
  Document,
  List,
  View,
  Monitor,
  Box,
  Location,
  Connection
} from '@element-plus/icons-vue'

defineOptions({
  name: 'NovelManage'
})

const route = useRoute()
const router = useRouter()

const searchInfo = ref({
  novelId: route.params.id || ''
})

const novelInfo = ref({})
const statistics = ref({})
const recentActivities = ref([
  {
    id: 1,
    timestamp: '2024-01-15 14:30',
    type: 'primary',
    description: '创建了新角色：李明'
  },
  {
    id: 2,
    timestamp: '2024-01-15 10:20',
    type: 'success',
    description: '完成了第5章的写作'
  },
  {
    id: 3,
    timestamp: '2024-01-14 16:45',
    type: 'warning',
    description: '埋下了新的伏笔：神秘信件'
  }
])

// 重置
const onReset = () => {
  searchInfo.value = {
    novelId: ''
  }
}

// 搜索
const onSubmit = () => {
  if (searchInfo.value.novelId) {
    loadNovelData(searchInfo.value.novelId)
  }
}

// 加载小说数据
const loadNovelData = async (novelId) => {
  try {
    const res = await findNovel({ ID: novelId })
    if (res.code === 0) {
      novelInfo.value = res.data.novel
      await loadStatistics(novelId)
    } else {
      ElMessage.error('获取小说信息失败')
    }
  } catch (error) {
    ElMessage.error('获取小说信息失败')
  }
}

// 加载统计信息
const loadStatistics = async (novelId) => {
  try {
    const res = await getNovelStatistics({ ID: novelId })
    if (res.code === 0) {
      statistics.value = res.data
    }
  } catch (error) {
    console.error('获取统计信息失败:', error)
  }
}

// 编辑小说
const editNovel = () => {
  router.push({
    name: 'Novel',
    query: { edit: novelInfo.value.id }
  })
}

// 跳转到模块
const goToModule = (module) => {
  if (!novelInfo.value.id) {
    ElMessage.warning('请先选择一个小说')
    return
  }
  
  const routes = {
    characters: 'NovelCharacters',
    chapters: 'NovelChapters',
    outlines: 'NovelOutlines',
    foreshadowing: 'NovelForeshadowing',
    worldbuilding: 'NovelWorldbuilding',
    items: 'NovelItems',
    scenes: 'NovelScenes',
    relationships: 'NovelRelationships'
  }
  
  if (routes[module]) {
    router.push({
      name: routes[module],
      params: { novelId: novelInfo.value.id }
    })
  }
}

// 生成字典方法
const getGenreLabel = (value) => {
  const genreMap = {
    fantasy: '奇幻',
    sci_fi: '科幻',
    romance: '言情',
    mystery: '悬疑',
    wuxia: '武侠',
    xianxia: '仙侠',
    urban: '都市',
    historical: '历史',
    game_world: '游戏'
  }
  return genreMap[value] || value
}

const getStatusLabel = (value) => {
  const statusMap = {
    draft: '草稿',
    planning: '规划中',
    writing: '写作中',
    revision: '修订中',
    completed: '已完成',
    published: '已发布',
    on_hold: '暂停',
    cancelled: '已取消',
    archived: '已归档'
  }
  return statusMap[value] || value
}

const getGenreTagType = (value) => {
  const typeMap = {
    fantasy: 'primary',
    sci_fi: 'success',
    romance: 'danger',
    mystery: 'warning',
    wuxia: 'info',
    xianxia: 'primary'
  }
  return typeMap[value] || ''
}

const getStatusTagType = (value) => {
  const typeMap = {
    draft: 'info',
    planning: 'warning',
    writing: 'primary',
    revision: 'warning',
    completed: 'success',
    published: 'success',
    on_hold: 'warning',
    cancelled: 'danger',
    archived: 'info'
  }
  return typeMap[value] || ''
}

// 初始化
onMounted(() => {
  if (route.params.id) {
    searchInfo.value.novelId = route.params.id
    loadNovelData(route.params.id)
  }
})
</script>

<style scoped>
.novel-manage-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: auto auto auto;
  gap: 20px;
  margin-top: 20px;
}

.novel-info-card {
  grid-column: 1 / 2;
  grid-row: 1 / 2;
}

.module-nav {
  grid-column: 2 / 3;
  grid-row: 1 / 3;
}

.statistics-card {
  grid-column: 1 / 2;
  grid-row: 2 / 3;
}

.activity-card {
  grid-column: 1 / 3;
  grid-row: 3 / 4;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.novel-info .info-row {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.novel-info .label {
  width: 80px;
  font-weight: 500;
  color: #606266;
}

.novel-info .value {
  flex: 1;
  color: #303133;
}

.novel-info .description {
  line-height: 1.5;
  max-height: 60px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}

.nav-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  background: #fafafa;
}

.nav-item:hover {
  border-color: #409eff;
  background: #ecf5ff;
  transform: translateY(-2px);
}

.nav-icon {
  font-size: 24px;
  color: #409eff;
  margin-bottom: 8px;
}

.nav-text {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 4px;
}

.nav-count {
  font-size: 12px;
  color: #909399;
  background: #f0f2f5;
  padding: 2px 8px;
  border-radius: 12px;
}

.statistics-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.stat-item {
  text-align: center;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 8px;
}

.stat-number {
  font-size: 24px;
  font-weight: 600;
  color: #409eff;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 12px;
  color: #909399;
}

.no-data {
  text-align: center;
  padding: 40px;
}

@media (max-width: 768px) {
  .novel-manage-container {
    grid-template-columns: 1fr;
    grid-template-rows: auto auto auto auto;
  }
  
  .novel-info-card {
    grid-column: 1 / 2;
    grid-row: 1 / 2;
  }
  
  .statistics-card {
    grid-column: 1 / 2;
    grid-row: 2 / 3;
  }
  
  .module-nav {
    grid-column: 1 / 2;
    grid-row: 3 / 4;
  }
  
  .activity-card {
    grid-column: 1 / 2;
    grid-row: 4 / 5;
  }
  
  .nav-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .statistics-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>