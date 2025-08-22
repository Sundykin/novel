<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="大纲标题" prop="title">
          <el-input v-model="searchInfo.title" placeholder="搜索大纲标题" />
        </el-form-item>
        <el-form-item label="卷号" prop="volume">
          <el-input-number v-model="searchInfo.volume" :min="1" placeholder="卷号" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="规划中" value="planning" />
            <el-option label="进行中" value="in_progress" />
            <el-option label="已完成" value="completed" />
            <el-option label="需修订" value="revision" />
            <el-option label="已废弃" value="abandoned" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增大纲</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <el-button icon="sort" style="margin-left: 10px;" @click="openSortDialog">大纲排序</el-button>
        <el-button icon="document" style="margin-left: 10px;" @click="exportOutlines">导出大纲</el-button>
      </div>
      
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="卷号" prop="volume" width="80" sortable />
        <el-table-column align="left" label="起始章节" prop="startChapter" width="100" sortable />
        <el-table-column align="left" label="结束章节" prop="endChapter" width="100" sortable />
        <el-table-column align="left" label="大纲标题" prop="title" min-width="200" show-overflow-tooltip />
        <el-table-column align="left" label="进度" prop="progress" width="120">
          <template #default="scope">
            <el-progress :percentage="scope.row.progress || 0" :stroke-width="6" :show-text="false" />
            <span class="text-xs text-gray-500">{{ scope.row.progress || 0 }}%</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="目标字数" prop="targetWordCount" width="120">
          <template #default="scope">
            <span>{{ scope.row.targetWordCount || 0 }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="预计天数" prop="estimatedDays" width="100">
          <template #default="scope">
            <span>{{ scope.row.estimatedDays || 0 }}天</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="更新时间" prop="updatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="viewDetails(scope.row)">详情</el-button>
            <el-button type="primary" link class="table-button" @click="updateOutlineFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link class="table-button" @click="generateChapters(scope.row)">生成章节</el-button>
            <el-button type="primary" link class="table-button" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 大纲编辑对话框 -->
    <el-drawer destroy-on-close size="900" v-model="dialogFormVisible" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加大纲' : '编辑大纲' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="卷号:" prop="volume">
              <el-input-number v-model="formData.volume" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="起始章节:" prop="startChapter">
              <el-input-number v-model="formData.startChapter" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="结束章节:" prop="endChapter">
              <el-input-number v-model="formData.endChapter" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="大纲标题:" prop="title">
          <el-input v-model="formData.title" clearable placeholder="请输入大纲标题" />
        </el-form-item>
        
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="状态:" prop="status">
              <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%">
                <el-option label="规划中" value="planning" />
                <el-option label="进行中" value="in_progress" />
                <el-option label="已完成" value="completed" />
                <el-option label="需修订" value="revision" />
                <el-option label="已废弃" value="abandoned" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="进度:" prop="progress">
              <el-slider v-model="formData.progress" :min="0" :max="100" show-input />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="目标字数:" prop="targetWordCount">
              <el-input-number v-model="formData.targetWordCount" :min="0" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="预计天数:" prop="estimatedDays">
              <el-input-number v-model="formData.estimatedDays" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="优先级:" prop="priority">
              <el-rate v-model="formData.priority" show-text :max="5" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="大纲摘要:" prop="summary">
          <el-input v-model="formData.summary" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入大纲摘要" />
        </el-form-item>
        
        <el-form-item label="主线剧情:" prop="mainPlot">
          <el-input v-model="formData.mainPlot" :autosize="{ minRows: 4, maxRows: 8}" type="textarea" placeholder="请详细描述主线剧情发展" />
        </el-form-item>
        
        <el-form-item label="支线剧情:" prop="subPlots">
          <el-input v-model="formData.subPlots" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请描述支线剧情" />
        </el-form-item>
        
        <el-form-item label="角色发展:" prop="characterDevelopment">
          <el-input v-model="formData.characterDevelopment" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请描述角色在此阶段的发展" />
        </el-form-item>
        
        <el-form-item label="关键事件:" prop="keyEvents">
          <el-input v-model="formData.keyEvents" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请列出关键事件" />
        </el-form-item>
        
        <el-form-item label="伏笔安排:" prop="foreshadowingPlans">
          <el-input v-model="formData.foreshadowingPlans" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请描述伏笔的埋设和揭示计划" />
        </el-form-item>
        
        <el-form-item label="冲突设置:" prop="conflicts">
          <el-input v-model="formData.conflicts" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请描述主要冲突和矛盾" />
        </el-form-item>
        
        <el-form-item label="情感基调:" prop="emotionalTone">
          <el-input v-model="formData.emotionalTone" clearable placeholder="如：紧张刺激、温馨感人、悲伤沉重" />
        </el-form-item>
        
        <el-form-item label="写作要点:" prop="writingNotes">
          <el-input v-model="formData.writingNotes" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="写作时需要注意的要点" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 大纲详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="大纲详情" width="80%" destroy-on-close>
      <div v-if="currentOutline" class="outline-detail">
        <div class="detail-header">
          <h2>{{ currentOutline.title }}</h2>
          <div class="meta-info">
            <el-tag>第{{ currentOutline.volume }}卷</el-tag>
            <el-tag type="primary">第{{ currentOutline.startChapter }}-{{ currentOutline.endChapter }}章</el-tag>
            <el-tag :type="getStatusTagType(currentOutline.status)">{{ getStatusLabel(currentOutline.status) }}</el-tag>
            <el-rate v-model="currentOutline.priority" disabled :max="5" show-text />
          </div>
        </div>
        
        <div class="progress-section">
          <div class="progress-item">
            <span class="label">完成进度：</span>
            <el-progress :percentage="currentOutline.progress || 0" :stroke-width="8" />
          </div>
          <div class="stats-grid">
            <div class="stat-item">
              <div class="stat-value">{{ currentOutline.targetWordCount || 0 }}</div>
              <div class="stat-label">目标字数</div>
            </div>
            <div class="stat-item">
              <div class="stat-value">{{ currentOutline.estimatedDays || 0 }}</div>
              <div class="stat-label">预计天数</div>
            </div>
            <div class="stat-item">
              <div class="stat-value">{{ currentOutline.endChapter - currentOutline.startChapter + 1 }}</div>
              <div class="stat-label">章节数量</div>
            </div>
          </div>
        </div>
        
        <el-descriptions :column="1" border>
          <el-descriptions-item label="大纲摘要">
            <div class="content-text">{{ currentOutline.summary }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="主线剧情">
            <div class="content-text">{{ currentOutline.mainPlot }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="支线剧情">
            <div class="content-text">{{ currentOutline.subPlots }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="角色发展">
            <div class="content-text">{{ currentOutline.characterDevelopment }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="关键事件">
            <div class="content-text">{{ currentOutline.keyEvents }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="伏笔安排">
            <div class="content-text">{{ currentOutline.foreshadowingPlans }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="冲突设置">
            <div class="content-text">{{ currentOutline.conflicts }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="情感基调">
            <div class="content-text">{{ currentOutline.emotionalTone }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="写作要点">
            <div class="content-text">{{ currentOutline.writingNotes }}</div>
          </el-descriptions-item>
        </el-descriptions>
      </div>
    </el-dialog>

    <!-- 生成章节对话框 -->
    <el-dialog v-model="generateDialogVisible" title="生成章节" width="600px">
      <div class="generate-form">
        <el-alert
          title="将根据大纲自动生成章节框架"
          type="info"
          :closable="false"
          style="margin-bottom: 20px;"
        />
        <el-form :model="generateForm" label-width="120px">
          <el-form-item label="生成范围:">
            <span>第{{ generateForm.startChapter }}章 - 第{{ generateForm.endChapter }}章</span>
          </el-form-item>
          <el-form-item label="章节标题前缀:">
            <el-input v-model="generateForm.titlePrefix" placeholder="如：第一卷" />
          </el-form-item>
          <el-form-item label="每章目标字数:">
            <el-input-number v-model="generateForm.wordsPerChapter" :min="1000" :max="10000" />
          </el-form-item>
          <el-form-item label="生成选项:">
            <el-checkbox-group v-model="generateForm.options">
              <el-checkbox label="mainEvents">主要事件</el-checkbox>
              <el-checkbox label="characterDev">角色发展</el-checkbox>
              <el-checkbox label="foreshadowing">伏笔安排</el-checkbox>
              <el-checkbox label="conflicts">冲突设置</el-checkbox>
            </el-checkbox-group>
          </el-form-item>
        </el-form>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="generateDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmGenerate">生成章节</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createOutline,
  deleteOutline,
  deleteOutlineByIds,
  updateOutline,
  findOutline,
  getOutlineList
} from '@/api/novel'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'NovelOutlines'
})

const route = useRoute()
const novelId = route.params.novelId

const formData = ref({
  volume: 1,
  startChapter: 1,
  endChapter: 10,
  title: '',
  summary: '',
  mainPlot: '',
  subPlots: '',
  characterDevelopment: '',
  keyEvents: '',
  foreshadowingPlans: '',
  conflicts: '',
  emotionalTone: '',
  writingNotes: '',
  status: 'planning',
  progress: 0,
  targetWordCount: 30000,
  estimatedDays: 30,
  priority: 3,
  novelId: novelId
})

const rule = reactive({
  volume: [{
    required: true,
    message: '请输入卷号',
    trigger: 'blur'
  }],
  startChapter: [{
    required: true,
    message: '请输入起始章节',
    trigger: 'blur'
  }],
  endChapter: [{
    required: true,
    message: '请输入结束章节',
    trigger: 'blur'
  }],
  title: [{
    required: true,
    message: '请输入大纲标题',
    trigger: 'blur'
  }]
})

const searchInfo = ref({
  novelId: novelId
})

const generateForm = ref({
  startChapter: 1,
  endChapter: 10,
  titlePrefix: '',
  wordsPerChapter: 3000,
  options: ['mainEvents', 'characterDev']
})

const elSearchFormRef = ref()
const elFormRef = ref()
const multipleTable = ref()
const multipleSelection = ref([])
const type = ref('')
const dialogFormVisible = ref(false)
const detailDialogVisible = ref(false)
const generateDialogVisible = ref(false)
const currentOutline = ref(null)
const tableData = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)

// 重置
const onReset = () => {
  searchInfo.value = {
    novelId: novelId
  }
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getOutlineList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 字典方法
const getStatusLabel = (value) => {
  const statusMap = {
    planning: '规划中',
    in_progress: '进行中',
    completed: '已完成',
    revision: '需修订',
    abandoned: '已废弃'
  }
  return statusMap[value] || value
}

const getStatusTagType = (value) => {
  const typeMap = {
    planning: 'warning',
    in_progress: 'primary',
    completed: 'success',
    revision: 'danger',
    abandoned: 'info'
  }
  return typeMap[value] || ''
}

// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 查看详情
const viewDetails = (row) => {
  currentOutline.value = row
  detailDialogVisible.value = true
}

// 生成章节
const generateChapters = (row) => {
  generateForm.value = {
    startChapter: row.startChapter,
    endChapter: row.endChapter,
    titlePrefix: row.title,
    wordsPerChapter: Math.round((row.targetWordCount || 30000) / (row.endChapter - row.startChapter + 1)),
    options: ['mainEvents', 'characterDev']
  }
  generateDialogVisible.value = true
}

// 确认生成章节
const confirmGenerate = () => {
  ElMessage.info('章节生成功能开发中...')
  generateDialogVisible.value = false
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除这个大纲吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteOutlineFunc(row)
  })
}

// 批量删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除选中的大纲吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteOutlineByIds({ IDs })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 编辑
const updateOutlineFunc = async(row) => {
  const res = await findOutline({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.outline
    dialogFormVisible.value = true
  }
}

// 删除
const deleteOutlineFunc = async (row) => {
  const res = await deleteOutline({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
  }
}

// 打开排序对话框
const openSortDialog = () => {
  ElMessage.info('大纲排序功能开发中...')
}

// 导出大纲
const exportOutlines = () => {
  ElMessage.info('导出大纲功能开发中...')
}

// 打开对话框
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭对话框
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    volume: 1,
    startChapter: 1,
    endChapter: 10,
    title: '',
    summary: '',
    mainPlot: '',
    subPlots: '',
    characterDevelopment: '',
    keyEvents: '',
    foreshadowingPlans: '',
    conflicts: '',
    emotionalTone: '',
    writingNotes: '',
    status: 'planning',
    progress: 0,
    targetWordCount: 30000,
    estimatedDays: 30,
    priority: 3,
    novelId: novelId
  }
}

// 确定
const enterDialog = async() => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createOutline(formData.value)
        break
      case 'update':
        res = await updateOutline(formData.value)
        break
      default:
        res = await createOutline(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage.success('操作成功')
      closeDialog()
      getTableData()
    }
  })
}

onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.outline-detail {
  padding: 20px;
}

.detail-header {
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e4e7ed;
}

.detail-header h2 {
  margin: 0 0 15px 0;
  color: #303133;
}

.meta-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.progress-section {
  margin-bottom: 20px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
}

.progress-item {
  margin-bottom: 15px;
}

.progress-item .label {
  font-weight: 500;
  color: #606266;
  margin-bottom: 8px;
  display: block;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.stat-item {
  text-align: center;
  padding: 15px;
  background: white;
  border-radius: 6px;
  border: 1px solid #e4e7ed;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  color: #409eff;
  margin-bottom: 5px;
}

.stat-label {
  font-size: 12px;
  color: #909399;
}

.content-text {
  line-height: 1.6;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.generate-form {
  padding: 20px;
}
</style>