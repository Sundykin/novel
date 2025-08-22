<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="小说标题" prop="title">
          <el-input v-model="searchInfo.title" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="类型" prop="genre">
          <el-select v-model="searchInfo.genre" clearable placeholder="请选择">
            <el-option label="奇幻" value="fantasy" />
            <el-option label="科幻" value="sci_fi" />
            <el-option label="言情" value="romance" />
            <el-option label="悬疑" value="mystery" />
            <el-option label="武侠" value="wuxia" />
            <el-option label="仙侠" value="xianxia" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="草稿" value="draft" />
            <el-option label="规划中" value="planning" />
            <el-option label="写作中" value="writing" />
            <el-option label="已完成" value="completed" />
            <el-option label="已发布" value="published" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
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
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="小说标题" prop="title" width="200" />
        <el-table-column align="left" label="类型" prop="genre" width="120">
          <template #default="scope">
            <el-tag :type="getGenreTagType(scope.row.genre)">{{ getGenreLabel(scope.row.genre) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="进度" prop="progress" width="120">
          <template #default="scope">
            <el-progress :percentage="scope.row.progress" :stroke-width="8" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="总字数" prop="totalWords" width="120" />
        <el-table-column align="left" label="目标字数" prop="targetWords" width="120" />
        <el-table-column align="left" label="简介" prop="description" show-overflow-tooltip />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">查看</el-button>
            <el-button type="primary" link class="table-button" @click="updateNovelFunc(scope.row)">变更</el-button>
            <el-button type="primary" link class="table-button" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link class="table-button" @click="manageNovel(scope.row)">管理</el-button>
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="小说标题:" prop="title">
          <el-input v-model="formData.title" clearable placeholder="请输入小说标题" />
        </el-form-item>
        <el-form-item label="副标题:" prop="subtitle">
          <el-input v-model="formData.subtitle" clearable placeholder="请输入副标题" />
        </el-form-item>
        <el-form-item label="类型:" prop="genre">
          <el-select v-model="formData.genre" placeholder="请选择类型" style="width:100%">
            <el-option label="奇幻" value="fantasy" />
            <el-option label="科幻" value="sci_fi" />
            <el-option label="言情" value="romance" />
            <el-option label="悬疑" value="mystery" />
            <el-option label="武侠" value="wuxia" />
            <el-option label="仙侠" value="xianxia" />
            <el-option label="都市" value="urban" />
            <el-option label="历史" value="historical" />
            <el-option label="游戏" value="game_world" />
          </el-select>
        </el-form-item>
        <el-form-item label="子类型:" prop="subGenre">
          <el-input v-model="formData.subGenre" clearable placeholder="请输入子类型" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%">
            <el-option label="草稿" value="draft" />
            <el-option label="规划中" value="planning" />
            <el-option label="写作中" value="writing" />
            <el-option label="修订中" value="revision" />
            <el-option label="已完成" value="completed" />
            <el-option label="已发布" value="published" />
            <el-option label="暂停" value="on_hold" />
          </el-select>
        </el-form-item>
        <el-form-item label="目标字数:" prop="targetWords">
          <el-input-number v-model="formData.targetWords" :min="0" style="width:100%" />
        </el-form-item>
        <el-form-item label="目标读者:" prop="targetAudience">
          <el-input v-model="formData.targetAudience" clearable placeholder="请输入目标读者" />
        </el-form-item>
        <el-form-item label="内容分级:" prop="contentRating">
          <el-select v-model="formData.contentRating" placeholder="请选择内容分级" style="width:100%">
            <el-option label="全年龄" value="G" />
            <el-option label="建议家长指导" value="PG" />
            <el-option label="13岁以上" value="PG13" />
            <el-option label="限制级" value="R" />
          </el-select>
        </el-form-item>
        <el-form-item label="简介:" prop="description">
          <el-input v-model="formData.description" :autosize="{ minRows: 4, maxRows: 8}" type="textarea" placeholder="请输入小说简介" />
        </el-form-item>
        <el-form-item label="故事梗概:" prop="synopsis">
          <el-input v-model="formData.synopsis" :autosize="{ minRows: 4, maxRows: 8}" type="textarea" placeholder="请输入故事梗概" />
        </el-form-item>
        <el-form-item label="标签:" prop="tags">
          <el-input v-model="formData.tags" clearable placeholder="请输入标签，用逗号分隔" />
        </el-form-item>
        <el-form-item label="关键词:" prop="keywords">
          <el-input v-model="formData.keywords" clearable placeholder="请输入关键词，用逗号分隔" />
        </el-form-item>
        <el-form-item label="作者备注:" prop="authorNotes">
          <el-input v-model="formData.authorNotes" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入作者备注" />
        </el-form-item>
        <el-form-item label="写作风格:" prop="writingStyle">
          <el-input v-model="formData.writingStyle" :autosize="{ minRows: 3, maxRows: 6}" type="textarea" placeholder="请输入写作风格" />
        </el-form-item>
        <el-form-item label="主题:" prop="themes">
          <el-input v-model="formData.themes" clearable placeholder="请输入主题" />
        </el-form-item>
        <el-form-item label="是否公开:">
          <el-switch v-model="formData.isPublic" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createNovel,
  deleteNovel,
  deleteNovelByIds,
  updateNovel,
  findNovel,
  getNovelList
} from '@/api/novel'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

defineOptions({
  name: 'Novel'
})

const router = useRouter()

// 自动化生成的字典
const formData = ref({
  title: '',
  subtitle: '',
  description: '',
  synopsis: '',
  genre: '',
  subGenre: '',
  tags: '',
  keywords: '',
  status: 'draft',
  progress: 0,
  totalChapters: 0,
  currentChapter: 0,
  totalWords: 0,
  targetWords: 0,
  targetAudience: '',
  contentRating: 'G',
  authorNotes: '',
  writingStyle: '',
  themes: '',
  isPublic: false,
  publishPlatform: '',
  collaborators: ''
})

// 验证规则
const rule = reactive({
  title: [{
    required: true,
    message: '请输入小说标题',
    trigger: 'blur'
  }],
  genre: [{
    required: true,
    message: '请选择类型',
    trigger: 'change'
  }],
  status: [{
    required: true,
    message: '请选择状态',
    trigger: 'change'
  }]
})

const searchInfo = ref({})

const elSearchFormRef = ref()
const elFormRef = ref()
const multipleTable = ref()
const multipleSelection = ref([])
const type = ref('')
const dialogFormVisible = ref(false)
const tableData = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)

// 重置
const onReset = () => {
  searchInfo.value = {}
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

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getNovelList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

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

// 多选数据
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteNovelFunc(row)
  })
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deleteNovelByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const updateNovelFunc = async(row) => {
  const res = await findNovel({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.novel
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteNovelFunc = async (row) => {
  const res = await deleteNovel({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogTitle = ref('新增')

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogTitle.value = '新增'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    title: '',
    subtitle: '',
    description: '',
    synopsis: '',
    genre: '',
    subGenre: '',
    tags: '',
    keywords: '',
    status: 'draft',
    progress: 0,
    totalChapters: 0,
    currentChapter: 0,
    totalWords: 0,
    targetWords: 0,
    targetAudience: '',
    contentRating: 'G',
    authorNotes: '',
    writingStyle: '',
    themes: '',
    isPublic: false,
    publishPlatform: '',
    collaborators: ''
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createNovel(formData.value)
        break
      case 'update':
        res = await updateNovel(formData.value)
        break
      default:
        res = await createNovel(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

// 查看详情
const getDetails = async (row) => {
  // 这里可以打开详情页面或者弹窗显示详情
  const res = await findNovel({ ID: row.ID })
  if (res.code === 0) {
    // 可以在这里处理详情显示逻辑
    console.log('小说详情:', res.data.novel)
  }
}

// 管理小说
const manageNovel = (row) => {
  router.push({
    name: 'NovelManage',
    params: { id: row.ID }
  })
}
</script>

<style>
</style>