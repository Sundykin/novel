<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="章节标题" prop="chapterTitle">
          <el-input v-model="searchInfo.chapterTitle" placeholder="搜索章节标题" />
        </el-form-item>
        <el-form-item label="卷号" prop="volume">
          <el-input-number v-model="searchInfo.volume" :min="1" placeholder="卷号" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="草稿" value="draft" />
            <el-option label="写作中" value="writing" />
            <el-option label="已完成" value="completed" />
            <el-option label="已发布" value="published" />
            <el-option label="需修订" value="revision" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增章节</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <el-button icon="sort" style="margin-left: 10px;" @click="openSortDialog">章节排序</el-button>
        <el-button icon="document" style="margin-left: 10px;" @click="batchExport">批量导出</el-button>
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
        <el-table-column align="left" label="章节号" prop="chapterNumber" width="80" sortable />
        <el-table-column align="left" label="卷号" prop="volume" width="80" sortable />
        <el-table-column align="left" label="章节标题" prop="chapterTitle" min-width="200" show-overflow-tooltip />
        <el-table-column align="left" label="字数" prop="wordCount" width="100" sortable>
          <template #default="scope">
            <span>{{ scope.row.wordCount || 0 }}</span>
            <span v-if="scope.row.targetWordCount" class="text-gray-400"> / {{ scope.row.targetWordCount }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="进度" prop="progress" width="120">
          <template #default="scope">
            <el-progress :percentage="getProgress(scope.row)" :stroke-width="6" :show-text="false" />
            <span class="text-xs text-gray-500">{{ getProgress(scope.row) }}%</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="情感基调" prop="emotionalTone" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.emotionalTone" size="small" :type="getToneTagType(scope.row.emotionalTone)">
              {{ scope.row.emotionalTone }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="更新时间" prop="updatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="280">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="viewContent(scope.row)">查看</el-button>
            <el-button type="primary" link class="table-button" @click="editContent(scope.row)">编辑</el-button>
            <el-button type="primary" link class="table-button" @click="updateChapterFunc(scope.row)">设置</el-button>
            <el-button type="primary" link class="table-button" @click="copyChapter(scope.row)">复制</el-button>
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

    <!-- 章节编辑对话框 -->
    <el-drawer destroy-on-close size="900" v-model="dialogFormVisible" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加章节' : '编辑章节' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="章节号:" prop="chapterNumber">
              <el-input-number v-model="formData.chapterNumber" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="卷号:" prop="volume">
              <el-input-number v-model="formData.volume" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="状态:" prop="status">
              <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%">
                <el-option label="草稿" value="draft" />
                <el-option label="写作中" value="writing" />
                <el-option label="已完成" value="completed" />
                <el-option label="已发布" value="published" />
                <el-option label="需修订" value="revision" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="章节标题:" prop="chapterTitle">
          <el-input v-model="formData.chapterTitle" clearable placeholder="请输入章节标题" />
        </el-form-item>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="目标字数:" prop="targetWordCount">
              <el-input-number v-model="formData.targetWordCount" :min="0" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="情感基调:" prop="emotionalTone">
              <el-input v-model="formData.emotionalTone" clearable placeholder="如：紧张、温馨、悲伤" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="主要事件:" prop="mainEvents">
          <el-input v-model="formData.mainEvents" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入本章主要事件" />
        </el-form-item>
        
        <el-form-item label="主线剧情:" prop="mainPlot">
          <el-input v-model="formData.mainPlot" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入主线剧情发展" />
        </el-form-item>
        
        <el-form-item label="支线剧情:" prop="subPlots">
          <el-input v-model="formData.subPlots" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入支线剧情" />
        </el-form-item>
        
        <el-form-item label="角色发展:" prop="characterDevelopment">
          <el-input v-model="formData.characterDevelopment" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入角色发展情况" />
        </el-form-item>
        
        <el-form-item label="关键事件:" prop="keyEvents">
          <el-input v-model="formData.keyEvents" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入关键事件" />
        </el-form-item>
        
        <el-form-item label="埋下的伏笔:" prop="foreshadowingPlanted">
          <el-input v-model="formData.foreshadowingPlanted" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入本章埋下的伏笔" />
        </el-form-item>
        
        <el-form-item label="揭示的伏笔:" prop="foreshadowingRevealed">
          <el-input v-model="formData.foreshadowingRevealed" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入本章揭示的伏笔" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 章节内容编辑对话框 -->
    <el-dialog v-model="contentDialogVisible" title="章节内容" width="90%" destroy-on-close>
      <div class="content-editor">
        <div class="editor-toolbar">
          <el-button-group>
            <el-button size="small" @click="insertTemplate('dialogue')">对话模板</el-button>
            <el-button size="small" @click="insertTemplate('scene')">场景描述</el-button>
            <el-button size="small" @click="insertTemplate('action')">动作描述</el-button>
          </el-button-group>
          <div class="word-count">
            字数：{{ contentWordCount }} / {{ currentChapter.targetWordCount || 0 }}
          </div>
        </div>
        <el-input
          v-model="chapterContent"
          type="textarea"
          :autosize="{ minRows: 20, maxRows: 30 }"
          placeholder="请输入章节内容..."
          @input="updateWordCount"
        />
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="contentDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveContent">保存内容</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 章节内容查看对话框 -->
    <el-dialog v-model="viewDialogVisible" title="章节内容" width="80%" destroy-on-close>
      <div class="content-viewer">
        <div class="chapter-header">
          <h2>{{ currentChapter.chapterTitle }}</h2>
          <div class="chapter-meta">
            <el-tag>第{{ currentChapter.volume }}卷</el-tag>
            <el-tag type="primary">第{{ currentChapter.chapterNumber }}章</el-tag>
            <el-tag :type="getStatusTagType(currentChapter.status)">{{ getStatusLabel(currentChapter.status) }}</el-tag>
          </div>
        </div>
        <div class="chapter-content">
          <pre>{{ currentChapter.content || '暂无内容' }}</pre>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createChapter,
  deleteChapter,
  deleteChapterByIds,
  updateChapter,
  findChapter,
  getChapterList
} from '@/api/novel'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'NovelChapters'
})

const route = useRoute()
const novelId = route.params.novelId

const formData = ref({
  chapterNumber: 1,
  volume: 1,
  chapterTitle: '',
  content: '',
  mainEvents: '',
  emotionalTone: '',
  mainPlot: '',
  subPlots: '',
  characterDevelopment: '',
  keyEvents: '',
  foreshadowingPlanted: '',
  foreshadowingRevealed: '',
  wordCount: 0,
  targetWordCount: 3000,
  status: 'draft',
  novelId: novelId
})

const rule = reactive({
  chapterNumber: [{
    required: true,
    message: '请输入章节号',
    trigger: 'blur'
  }],
  volume: [{
    required: true,
    message: '请输入卷号',
    trigger: 'blur'
  }],
  chapterTitle: [{
    required: true,
    message: '请输入章节标题',
    trigger: 'blur'
  }]
})

const searchInfo = ref({
  novelId: novelId
})

const elSearchFormRef = ref()
const elFormRef = ref()
const multipleTable = ref()
const multipleSelection = ref([])
const type = ref('')
const dialogFormVisible = ref(false)
const contentDialogVisible = ref(false)
const viewDialogVisible = ref(false)
const currentChapter = ref({})
const chapterContent = ref('')
const contentWordCount = ref(0)
const tableData = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)

// 计算进度
const getProgress = (row) => {
  if (!row.targetWordCount || row.targetWordCount === 0) return 0
  return Math.min(Math.round((row.wordCount || 0) / row.targetWordCount * 100), 100)
}

// 更新字数统计
const updateWordCount = () => {
  contentWordCount.value = chapterContent.value.length
}

// 插入模板
const insertTemplate = (type) => {
  const templates = {
    dialogue: '\n\n"对话内容，" 角色说道。\n\n',
    scene: '\n\n[场景描述：时间、地点、环境氛围]\n\n',
    action: '\n\n[动作描述：角色的具体行为和反应]\n\n'
  }
  
  if (templates[type]) {
    chapterContent.value += templates[type]
    updateWordCount()
  }
}

// 保存内容
const saveContent = async () => {
  try {
    const updateData = {
      ...currentChapter.value,
      content: chapterContent.value,
      wordCount: contentWordCount.value
    }
    
    const res = await updateChapter(updateData)
    if (res.code === 0) {
      ElMessage.success('保存成功')
      contentDialogVisible.value = false
      getTableData()
    }
  } catch (error) {
    ElMessage.error('保存失败')
  }
}

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
  const table = await getChapterList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    draft: '草稿',
    writing: '写作中',
    completed: '已完成',
    published: '已发布',
    revision: '需修订'
  }
  return statusMap[value] || value
}

const getStatusTagType = (value) => {
  const typeMap = {
    draft: 'info',
    writing: 'warning',
    completed: 'success',
    published: 'primary',
    revision: 'danger'
  }
  return typeMap[value] || ''
}

const getToneTagType = (value) => {
  const toneTypes = {
    '紧张': 'danger',
    '温馨': 'success',
    '悲伤': 'warning',
    '欢乐': 'primary',
    '神秘': 'info'
  }
  return toneTypes[value] || ''
}

// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 查看内容
const viewContent = (row) => {
  currentChapter.value = row
  viewDialogVisible.value = true
}

// 编辑内容
const editContent = async (row) => {
  const res = await findChapter({ ID: row.ID })
  if (res.code === 0) {
    currentChapter.value = res.data.chapter
    chapterContent.value = res.data.chapter.content || ''
    updateWordCount()
    contentDialogVisible.value = true
  }
}

// 复制章节
const copyChapter = (row) => {
  formData.value = {
    ...row,
    ID: undefined,
    chapterNumber: row.chapterNumber + 1,
    chapterTitle: row.chapterTitle + ' (副本)',
    content: '',
    wordCount: 0,
    status: 'draft'
  }
  type.value = 'create'
  dialogFormVisible.value = true
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除这个章节吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteChapterFunc(row)
  })
}

// 批量删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除选中的章节吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteChapterByIds({ IDs })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 编辑
const updateChapterFunc = async(row) => {
  const res = await findChapter({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.chapter
    dialogFormVisible.value = true
  }
}

// 删除
const deleteChapterFunc = async (row) => {
  const res = await deleteChapter({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
  }
}

// 打开排序对话框
const openSortDialog = () => {
  ElMessage.info('章节排序功能开发中...')
}

// 批量导出
const batchExport = () => {
  ElMessage.info('批量导出功能开发中...')
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
    chapterNumber: 1,
    volume: 1,
    chapterTitle: '',
    content: '',
    mainEvents: '',
    emotionalTone: '',
    mainPlot: '',
    subPlots: '',
    characterDevelopment: '',
    keyEvents: '',
    foreshadowingPlanted: '',
    foreshadowingRevealed: '',
    wordCount: 0,
    targetWordCount: 3000,
    status: 'draft',
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
        res = await createChapter(formData.value)
        break
      case 'update':
        res = await updateChapter(formData.value)
        break
      default:
        res = await createChapter(formData.value)
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
.content-editor {
  margin-bottom: 20px;
}

.editor-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
  background: #f5f7fa;
  border-radius: 4px;
}

.word-count {
  font-size: 14px;
  color: #606266;
}

.content-viewer {
  max-height: 70vh;
  overflow-y: auto;
}

.chapter-header {
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #e4e7ed;
}

.chapter-header h2 {
  margin: 0 0 10px 0;
  color: #303133;
}

.chapter-meta .el-tag {
  margin-right: 8px;
}

.chapter-content {
  line-height: 1.8;
  font-size: 16px;
  color: #303133;
}

.chapter-content pre {
  white-space: pre-wrap;
  word-wrap: break-word;
  font-family: inherit;
  margin: 0;
}
</style>