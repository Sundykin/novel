<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="伏笔标题" prop="title">
          <el-input v-model="searchInfo.title" placeholder="搜索伏笔标题" />
        </el-form-item>
        <el-form-item label="重要程度" prop="importanceLevel">
          <el-select v-model="searchInfo.importanceLevel" clearable placeholder="请选择">
            <el-option label="低" value="low" />
            <el-option label="中" value="medium" />
            <el-option label="高" value="high" />
            <el-option label="关键" value="critical" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="已埋下" value="planted" />
            <el-option label="已揭示" value="revealed" />
            <el-option label="已解决" value="resolved" />
            <el-option label="已废弃" value="abandoned" />
          </el-select>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="searchInfo.type" clearable placeholder="请选择">
            <el-option label="情节伏笔" value="plot" />
            <el-option label="角色伏笔" value="character" />
            <el-option label="物品伏笔" value="item" />
            <el-option label="环境伏笔" value="environment" />
            <el-option label="对话伏笔" value="dialogue" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增伏笔</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <el-button icon="view" style="margin-left: 10px;" @click="openTrackingView">伏笔追踪图</el-button>
        <el-button icon="document" style="margin-left: 10px;" @click="exportReport">导出报告</el-button>
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
        <el-table-column align="left" label="伏笔ID" prop="foreshadowingId" width="120" />
        <el-table-column align="left" label="标题" prop="title" min-width="180" show-overflow-tooltip />
        <el-table-column align="left" label="类型" prop="type" width="100">
          <template #default="scope">
            <el-tag :type="getTypeTagType(scope.row.type)" size="small">{{ getTypeLabel(scope.row.type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="重要程度" prop="importanceLevel" width="120">
          <template #default="scope">
            <el-tag :type="getImportanceTagType(scope.row.importanceLevel)">{{ getImportanceLabel(scope.row.importanceLevel) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="埋下章节" prop="plantedChapter" width="100">
          <template #default="scope">
            <span v-if="scope.row.plantedChapter">第{{ scope.row.plantedChapter }}章</span>
            <span v-else class="text-gray-400">未设置</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="揭示章节" prop="revealedChapter" width="100">
          <template #default="scope">
            <span v-if="scope.row.revealedChapter">第{{ scope.row.revealedChapter }}章</span>
            <span v-else class="text-gray-400">未揭示</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="质量评分" prop="qualityScore" width="120">
          <template #default="scope">
            <el-rate v-if="scope.row.qualityScore" v-model="scope.row.qualityScore" disabled show-score text-color="#ff9900" :max="10" />
            <span v-else class="text-gray-400">未评分</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="更新时间" prop="updatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="viewDetails(scope.row)">详情</el-button>
            <el-button type="primary" link class="table-button" @click="updateForeshadowingFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link class="table-button" @click="updateStatus(scope.row)">状态</el-button>
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

    <!-- 伏笔编辑对话框 -->
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加伏笔' : '编辑伏笔' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="伏笔ID:" prop="foreshadowingId">
              <el-input v-model="formData.foreshadowingId" clearable placeholder="请输入伏笔唯一标识" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="伏笔类型:" prop="type">
              <el-select v-model="formData.type" placeholder="请选择伏笔类型" style="width:100%">
                <el-option label="情节伏笔" value="plot" />
                <el-option label="角色伏笔" value="character" />
                <el-option label="物品伏笔" value="item" />
                <el-option label="环境伏笔" value="environment" />
                <el-option label="对话伏笔" value="dialogue" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="伏笔标题:" prop="title">
          <el-input v-model="formData.title" clearable placeholder="请输入伏笔标题" />
        </el-form-item>
        
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="重要程度:" prop="importanceLevel">
              <el-select v-model="formData.importanceLevel" placeholder="请选择" style="width:100%">
                <el-option label="低" value="low" />
                <el-option label="中" value="medium" />
                <el-option label="高" value="high" />
                <el-option label="关键" value="critical" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="状态:" prop="status">
              <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%">
                <el-option label="已埋下" value="planted" />
                <el-option label="已揭示" value="revealed" />
                <el-option label="已解决" value="resolved" />
                <el-option label="已废弃" value="abandoned" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="质量评分:" prop="qualityScore">
              <el-rate v-model="formData.qualityScore" show-text :max="10" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="埋下章节:" prop="plantedChapter">
              <el-input-number v-model="formData.plantedChapter" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="揭示章节:" prop="revealedChapter">
              <el-input-number v-model="formData.revealedChapter" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="解决章节:" prop="resolvedChapter">
              <el-input-number v-model="formData.resolvedChapter" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="伏笔内容:" prop="content">
          <el-input v-model="formData.content" :autosize="{ minRows: 4, maxRows: 8}" type="textarea" placeholder="请详细描述伏笔内容" />
        </el-form-item>
        
        <el-form-item label="伏笔描述:" prop="description">
          <el-input v-model="formData.description" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入伏笔的背景描述" />
        </el-form-item>
        
        <el-form-item label="关联角色:" prop="relatedCharacters">
          <el-input v-model="formData.relatedCharacters" clearable placeholder="请输入相关角色，用逗号分隔" />
        </el-form-item>
        
        <el-form-item label="关联事件:" prop="relatedEvents">
          <el-input v-model="formData.relatedEvents" clearable placeholder="请输入相关事件，用逗号分隔" />
        </el-form-item>
        
        <el-form-item label="预期效果:" prop="expectedImpact">
          <el-input v-model="formData.expectedImpact" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请描述这个伏笔的预期效果" />
        </el-form-item>
        
        <el-form-item label="备注:" prop="notes">
          <el-input v-model="formData.notes" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="其他备注信息" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 伏笔详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="伏笔详情" width="70%" destroy-on-close>
      <div v-if="currentForeshadowing" class="foreshadowing-detail">
        <div class="detail-header">
          <h2>{{ currentForeshadowing.title }}</h2>
          <div class="tags">
            <el-tag :type="getTypeTagType(currentForeshadowing.type)">{{ getTypeLabel(currentForeshadowing.type) }}</el-tag>
            <el-tag :type="getImportanceTagType(currentForeshadowing.importanceLevel)">{{ getImportanceLabel(currentForeshadowing.importanceLevel) }}</el-tag>
            <el-tag :type="getStatusTagType(currentForeshadowing.status)">{{ getStatusLabel(currentForeshadowing.status) }}</el-tag>
          </div>
        </div>
        
        <el-descriptions :column="2" border>
          <el-descriptions-item label="伏笔ID">{{ currentForeshadowing.foreshadowingId }}</el-descriptions-item>
          <el-descriptions-item label="质量评分">
            <el-rate v-if="currentForeshadowing.qualityScore" v-model="currentForeshadowing.qualityScore" disabled :max="10" />
            <span v-else>未评分</span>
          </el-descriptions-item>
          <el-descriptions-item label="埋下章节">
            <span v-if="currentForeshadowing.plantedChapter">第{{ currentForeshadowing.plantedChapter }}章</span>
            <span v-else>未设置</span>
          </el-descriptions-item>
          <el-descriptions-item label="揭示章节">
            <span v-if="currentForeshadowing.revealedChapter">第{{ currentForeshadowing.revealedChapter }}章</span>
            <span v-else>未揭示</span>
          </el-descriptions-item>
          <el-descriptions-item label="解决章节">
            <span v-if="currentForeshadowing.resolvedChapter">第{{ currentForeshadowing.resolvedChapter }}章</span>
            <span v-else>未解决</span>
          </el-descriptions-item>
          <el-descriptions-item label="关联角色">{{ currentForeshadowing.relatedCharacters || '无' }}</el-descriptions-item>
          <el-descriptions-item label="关联事件" :span="2">{{ currentForeshadowing.relatedEvents || '无' }}</el-descriptions-item>
          <el-descriptions-item label="伏笔内容" :span="2">
            <div class="content-text">{{ currentForeshadowing.content }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="伏笔描述" :span="2">
            <div class="content-text">{{ currentForeshadowing.description }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="预期效果" :span="2">
            <div class="content-text">{{ currentForeshadowing.expectedImpact }}</div>
          </el-descriptions-item>
          <el-descriptions-item label="备注" :span="2">
            <div class="content-text">{{ currentForeshadowing.notes || '无' }}</div>
          </el-descriptions-item>
        </el-descriptions>
      </div>
    </el-dialog>

    <!-- 状态更新对话框 -->
    <el-dialog v-model="statusDialogVisible" title="更新伏笔状态" width="500px">
      <el-form :model="statusForm" label-width="100px">
        <el-form-item label="当前状态:">
          <el-tag :type="getStatusTagType(statusForm.currentStatus)">{{ getStatusLabel(statusForm.currentStatus) }}</el-tag>
        </el-form-item>
        <el-form-item label="新状态:" prop="newStatus">
          <el-select v-model="statusForm.newStatus" placeholder="请选择新状态" style="width:100%">
            <el-option label="已埋下" value="planted" />
            <el-option label="已揭示" value="revealed" />
            <el-option label="已解决" value="resolved" />
            <el-option label="已废弃" value="abandoned" />
          </el-select>
        </el-form-item>
        <el-form-item label="更新说明:" prop="updateNote">
          <el-input v-model="statusForm.updateNote" type="textarea" :autosize="{ minRows: 3, maxRows: 5}" placeholder="请输入状态更新说明" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="statusDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmStatusUpdate">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createForeshadowing,
  deleteForeshadowing,
  deleteForeshadowingByIds,
  updateForeshadowing,
  findForeshadowing,
  getForeshadowingList
} from '@/api/novel'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'NovelForeshadowing'
})

const route = useRoute()
const novelId = route.params.novelId

const formData = ref({
  foreshadowingId: '',
  title: '',
  content: '',
  description: '',
  importanceLevel: 'medium',
  type: 'plot',
  category: '',
  status: 'planted',
  qualityScore: null,
  plantedChapter: null,
  revealedChapter: null,
  resolvedChapter: null,
  relatedCharacters: '',
  relatedEvents: '',
  expectedImpact: '',
  notes: '',
  novelId: novelId
})

const rule = reactive({
  foreshadowingId: [{
    required: true,
    message: '请输入伏笔ID',
    trigger: 'blur'
  }],
  title: [{
    required: true,
    message: '请输入伏笔标题',
    trigger: 'blur'
  }],
  content: [{
    required: true,
    message: '请输入伏笔内容',
    trigger: 'blur'
  }]
})

const searchInfo = ref({
  novelId: novelId
})

const statusForm = ref({
  id: null,
  currentStatus: '',
  newStatus: '',
  updateNote: ''
})

const elSearchFormRef = ref()
const elFormRef = ref()
const multipleTable = ref()
const multipleSelection = ref([])
const type = ref('')
const dialogFormVisible = ref(false)
const detailDialogVisible = ref(false)
const statusDialogVisible = ref(false)
const currentForeshadowing = ref(null)
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
  const table = await getForeshadowingList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 字典方法
const getTypeLabel = (value) => {
  const typeMap = {
    plot: '情节',
    character: '角色',
    item: '物品',
    environment: '环境',
    dialogue: '对话'
  }
  return typeMap[value] || value
}

const getImportanceLabel = (value) => {
  const importanceMap = {
    low: '低',
    medium: '中',
    high: '高',
    critical: '关键'
  }
  return importanceMap[value] || value
}

const getStatusLabel = (value) => {
  const statusMap = {
    planted: '已埋下',
    revealed: '已揭示',
    resolved: '已解决',
    abandoned: '已废弃'
  }
  return statusMap[value] || value
}

const getTypeTagType = (value) => {
  const typeMap = {
    plot: 'primary',
    character: 'success',
    item: 'warning',
    environment: 'info',
    dialogue: 'danger'
  }
  return typeMap[value] || ''
}

const getImportanceTagType = (value) => {
  const typeMap = {
    low: 'info',
    medium: 'warning',
    high: 'danger',
    critical: 'danger'
  }
  return typeMap[value] || ''
}

const getStatusTagType = (value) => {
  const typeMap = {
    planted: 'warning',
    revealed: 'primary',
    resolved: 'success',
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
  currentForeshadowing.value = row
  detailDialogVisible.value = true
}

// 更新状态
const updateStatus = (row) => {
  statusForm.value = {
    id: row.ID,
    currentStatus: row.status,
    newStatus: row.status,
    updateNote: ''
  }
  statusDialogVisible.value = true
}

// 确认状态更新
const confirmStatusUpdate = async () => {
  if (!statusForm.value.newStatus) {
    ElMessage.warning('请选择新状态')
    return
  }
  
  try {
    const updateData = {
      ID: statusForm.value.id,
      status: statusForm.value.newStatus,
      updateNote: statusForm.value.updateNote
    }
    
    const res = await updateForeshadowing(updateData)
    if (res.code === 0) {
      ElMessage.success('状态更新成功')
      statusDialogVisible.value = false
      getTableData()
    }
  } catch (error) {
    ElMessage.error('状态更新失败')
  }
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除这个伏笔吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteForeshadowingFunc(row)
  })
}

// 批量删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除选中的伏笔吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteForeshadowingByIds({ IDs })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 编辑
const updateForeshadowingFunc = async(row) => {
  const res = await findForeshadowing({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.foreshadowing
    dialogFormVisible.value = true
  }
}

// 删除
const deleteForeshadowingFunc = async (row) => {
  const res = await deleteForeshadowing({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
  }
}

// 打开追踪视图
const openTrackingView = () => {
  ElMessage.info('伏笔追踪图功能开发中...')
}

// 导出报告
const exportReport = () => {
  ElMessage.info('导出报告功能开发中...')
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
    foreshadowingId: '',
    title: '',
    content: '',
    description: '',
    importanceLevel: 'medium',
    type: 'plot',
    category: '',
    status: 'planted',
    qualityScore: null,
    plantedChapter: null,
    revealedChapter: null,
    resolvedChapter: null,
    relatedCharacters: '',
    relatedEvents: '',
    expectedImpact: '',
    notes: '',
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
        res = await createForeshadowing(formData.value)
        break
      case 'update':
        res = await updateForeshadowing(formData.value)
        break
      default:
        res = await createForeshadowing(formData.value)
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
.foreshadowing-detail {
  padding: 20px;
}

.detail-header {
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e4e7ed;
}

.detail-header h2 {
  margin: 0 0 10px 0;
  color: #303133;
}

.tags .el-tag {
  margin-right: 8px;
}

.content-text {
  line-height: 1.6;
  white-space: pre-wrap;
  word-wrap: break-word;
}
</style>