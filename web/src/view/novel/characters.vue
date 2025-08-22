<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="角色名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="搜索角色名称" />
        </el-form-item>
        <el-form-item label="角色类型" prop="role">
          <el-select v-model="searchInfo.role" clearable placeholder="请选择">
            <el-option label="主角" value="protagonist" />
            <el-option label="重要配角" value="major_supporting" />
            <el-option label="次要配角" value="minor_supporting" />
            <el-option label="反派" value="antagonist" />
            <el-option label="路人" value="background" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="活跃" value="active" />
            <el-option label="死亡" value="dead" />
            <el-option label="失踪" value="missing" />
            <el-option label="退场" value="retired" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增角色</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <el-button icon="connection" style="margin-left: 10px;" @click="openRelationshipDialog">管理关系</el-button>
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
        <el-table-column align="left" label="头像" prop="avatar" width="80">
          <template #default="scope">
            <el-avatar :size="50" :src="scope.row.avatar" :alt="scope.row.name">
              {{ scope.row.name?.charAt(0) }}
            </el-avatar>
          </template>
        </el-table-column>
        <el-table-column align="left" label="角色名称" prop="name" width="120" />
        <el-table-column align="left" label="角色类型" prop="role" width="120">
          <template #default="scope">
            <el-tag :type="getRoleTagType(scope.row.role)">{{ getRoleLabel(scope.row.role) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="重要程度" prop="importance" width="120">
          <template #default="scope">
            <el-rate v-model="scope.row.importance" disabled show-score />
          </template>
        </el-table-column>
        <el-table-column align="left" label="年龄" prop="age" width="80" />
        <el-table-column align="left" label="性别" prop="gender" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.gender === 'male' ? 'primary' : 'danger'" size="small">
              {{ scope.row.gender === 'male' ? '男' : scope.row.gender === 'female' ? '女' : '未知' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="首次出场" prop="firstAppearance" width="120" />
        <el-table-column align="left" label="描述" prop="description" show-overflow-tooltip />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="viewDetails(scope.row)">详情</el-button>
            <el-button type="primary" link class="table-button" @click="updateCharacterFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link class="table-button" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link class="table-button" @click="manageRelations(scope.row)">关系</el-button>
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

    <!-- 角色编辑对话框 -->
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加角色' : '编辑角色' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="角色名称:" prop="name">
              <el-input v-model="formData.name" clearable placeholder="请输入角色名称" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="角色类型:" prop="role">
              <el-select v-model="formData.role" placeholder="请选择角色类型" style="width:100%">
                <el-option label="主角" value="protagonist" />
                <el-option label="重要配角" value="major_supporting" />
                <el-option label="次要配角" value="minor_supporting" />
                <el-option label="反派" value="antagonist" />
                <el-option label="路人" value="background" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="年龄:" prop="age">
              <el-input-number v-model="formData.age" :min="0" :max="1000" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="性别:" prop="gender">
              <el-select v-model="formData.gender" placeholder="请选择性别" style="width:100%">
                <el-option label="男" value="male" />
                <el-option label="女" value="female" />
                <el-option label="未知" value="unknown" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="重要程度:" prop="importance">
              <el-rate v-model="formData.importance" show-text />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="状态:" prop="status">
              <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%">
                <el-option label="活跃" value="active" />
                <el-option label="死亡" value="dead" />
                <el-option label="失踪" value="missing" />
                <el-option label="退场" value="retired" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="首次出场:" prop="firstAppearance">
              <el-input v-model="formData.firstAppearance" clearable placeholder="如：第1章" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="头像URL:" prop="avatar">
          <el-input v-model="formData.avatar" clearable placeholder="请输入头像URL" />
        </el-form-item>
        
        <el-form-item label="角色描述:" prop="description">
          <el-input v-model="formData.description" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入角色描述" />
        </el-form-item>
        
        <el-form-item label="外貌特征:" prop="appearance">
          <el-input v-model="formData.appearance" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入外貌特征" />
        </el-form-item>
        
        <el-form-item label="性格特点:" prop="personality">
          <el-input v-model="formData.personality" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入性格特点" />
        </el-form-item>
        
        <el-form-item label="背景故事:" prop="background">
          <el-input v-model="formData.background" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入背景故事" />
        </el-form-item>
        
        <el-form-item label="能力技能:" prop="abilities">
          <el-input v-model="formData.abilities" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入能力技能" />
        </el-form-item>
        
        <el-form-item label="目标动机:" prop="goals">
          <el-input v-model="formData.goals" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="请输入目标动机" />
        </el-form-item>
        
        <el-form-item label="标签:" prop="tags">
          <el-input v-model="formData.tags" clearable placeholder="请输入标签，用逗号分隔" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 角色详情对话框 -->
    <el-dialog v-model="detailDialogVisible" title="角色详情" width="60%" destroy-on-close>
      <div v-if="currentCharacter" class="character-detail">
        <div class="detail-header">
          <el-avatar :size="80" :src="currentCharacter.avatar" :alt="currentCharacter.name">
            {{ currentCharacter.name?.charAt(0) }}
          </el-avatar>
          <div class="header-info">
            <h2>{{ currentCharacter.name }}</h2>
            <div class="tags">
              <el-tag :type="getRoleTagType(currentCharacter.role)">{{ getRoleLabel(currentCharacter.role) }}</el-tag>
              <el-tag :type="getStatusTagType(currentCharacter.status)">{{ getStatusLabel(currentCharacter.status) }}</el-tag>
            </div>
          </div>
        </div>
        
        <el-descriptions :column="2" border>
          <el-descriptions-item label="年龄">{{ currentCharacter.age }}</el-descriptions-item>
          <el-descriptions-item label="性别">{{ currentCharacter.gender === 'male' ? '男' : currentCharacter.gender === 'female' ? '女' : '未知' }}</el-descriptions-item>
          <el-descriptions-item label="重要程度">
            <el-rate v-model="currentCharacter.importance" disabled />
          </el-descriptions-item>
          <el-descriptions-item label="首次出场">{{ currentCharacter.firstAppearance }}</el-descriptions-item>
          <el-descriptions-item label="描述" :span="2">{{ currentCharacter.description }}</el-descriptions-item>
          <el-descriptions-item label="外貌特征" :span="2">{{ currentCharacter.appearance }}</el-descriptions-item>
          <el-descriptions-item label="性格特点" :span="2">{{ currentCharacter.personality }}</el-descriptions-item>
          <el-descriptions-item label="背景故事" :span="2">{{ currentCharacter.background }}</el-descriptions-item>
          <el-descriptions-item label="能力技能" :span="2">{{ currentCharacter.abilities }}</el-descriptions-item>
          <el-descriptions-item label="目标动机" :span="2">{{ currentCharacter.goals }}</el-descriptions-item>
        </el-descriptions>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCharacter,
  deleteCharacter,
  deleteCharacterByIds,
  updateCharacter,
  findCharacter,
  getCharacterList
} from '@/api/novel'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'NovelCharacters'
})

const route = useRoute()
const novelId = route.params.novelId

const formData = ref({
  name: '',
  role: 'minor_supporting',
  age: 0,
  gender: 'unknown',
  importance: 1,
  status: 'active',
  firstAppearance: '',
  avatar: '',
  description: '',
  appearance: '',
  personality: '',
  background: '',
  abilities: '',
  goals: '',
  tags: '',
  novelId: novelId
})

const rule = reactive({
  name: [{
    required: true,
    message: '请输入角色名称',
    trigger: 'blur'
  }],
  role: [{
    required: true,
    message: '请选择角色类型',
    trigger: 'change'
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
const detailDialogVisible = ref(false)
const currentCharacter = ref(null)
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
  const table = await getCharacterList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 字典方法
const getRoleLabel = (value) => {
  const roleMap = {
    protagonist: '主角',
    major_supporting: '重要配角',
    minor_supporting: '次要配角',
    antagonist: '反派',
    background: '路人'
  }
  return roleMap[value] || value
}

const getStatusLabel = (value) => {
  const statusMap = {
    active: '活跃',
    dead: '死亡',
    missing: '失踪',
    retired: '退场'
  }
  return statusMap[value] || value
}

const getRoleTagType = (value) => {
  const typeMap = {
    protagonist: 'danger',
    major_supporting: 'primary',
    minor_supporting: 'success',
    antagonist: 'warning',
    background: 'info'
  }
  return typeMap[value] || ''
}

const getStatusTagType = (value) => {
  const typeMap = {
    active: 'success',
    dead: 'danger',
    missing: 'warning',
    retired: 'info'
  }
  return typeMap[value] || ''
}

// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除这个角色吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCharacterFunc(row)
  })
}

// 批量删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除选中的角色吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteCharacterByIds({ IDs })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 编辑
const updateCharacterFunc = async(row) => {
  const res = await findCharacter({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.character
    dialogFormVisible.value = true
  }
}

// 删除
const deleteCharacterFunc = async (row) => {
  const res = await deleteCharacter({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
  }
}

// 查看详情
const viewDetails = (row) => {
  currentCharacter.value = row
  detailDialogVisible.value = true
}

// 管理关系
const manageRelations = (row) => {
  // TODO: 实现关系管理功能
  ElMessage.info('关系管理功能开发中...')
}

// 打开关系对话框
const openRelationshipDialog = () => {
  // TODO: 实现关系管理功能
  ElMessage.info('关系管理功能开发中...')
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
    name: '',
    role: 'minor_supporting',
    age: 0,
    gender: 'unknown',
    importance: 1,
    status: 'active',
    firstAppearance: '',
    avatar: '',
    description: '',
    appearance: '',
    personality: '',
    background: '',
    abilities: '',
    goals: '',
    tags: '',
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
        res = await createCharacter(formData.value)
        break
      case 'update':
        res = await updateCharacter(formData.value)
        break
      default:
        res = await createCharacter(formData.value)
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
.character-detail {
  padding: 20px;
}

.detail-header {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e4e7ed;
}

.header-info {
  margin-left: 20px;
  flex: 1;
}

.header-info h2 {
  margin: 0 0 10px 0;
  color: #303133;
}

.tags .el-tag {
  margin-right: 8px;
}
</style>