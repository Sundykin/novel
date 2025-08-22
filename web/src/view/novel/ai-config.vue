<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="配置名称" prop="configName">
          <el-input v-model="searchInfo.configName" placeholder="搜索配置名称" />
        </el-form-item>
        <el-form-item label="服务提供商" prop="provider">
          <el-select v-model="searchInfo.provider" clearable placeholder="请选择">
            <el-option label="OpenAI" value="openai" />
            <el-option label="Claude" value="claude" />
            <el-option label="Gemini" value="gemini" />
            <el-option label="千问" value="qianwen" />
            <el-option label="文心" value="baidu" />
            <el-option label="混元" value="tencent" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="激活" value="active" />
            <el-option label="停用" value="inactive" />
            <el-option label="测试中" value="testing" />
            <el-option label="错误" value="error" />
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
        <el-button type="primary" icon="plus" @click="openDialog">新增配置</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <el-button icon="refresh" style="margin-left: 10px;" @click="testAllConfigs">批量测试</el-button>
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
        <el-table-column align="left" label="配置名称" prop="configName" width="150" />
        <el-table-column align="left" label="服务提供商" prop="provider" width="120">
          <template #default="scope">
            <el-tag :type="getProviderTagType(scope.row.provider)">{{ getProviderLabel(scope.row.provider) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="AI模型" prop="model" width="150" />
        <el-table-column align="left" label="默认配置" prop="isDefault" width="100">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isDefault"
              @change="setDefault(scope.row)"
              :disabled="scope.row.isDefault"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="每日限制" prop="dailyLimit" width="100" />
        <el-table-column align="left" label="每月限制" prop="monthlyLimit" width="100" />
        <el-table-column align="left" label="创建时间" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="testConfig(scope.row)">测试</el-button>
            <el-button type="primary" link class="table-button" @click="updateConfigFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link class="table-button" @click="copyConfig(scope.row)">复制</el-button>
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

    <!-- AI配置编辑对话框 -->
    <el-drawer destroy-on-close size="600" v-model="dialogFormVisible" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加AI配置' : '编辑AI配置' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="120px">
        <el-form-item label="配置名称:" prop="configName">
          <el-input v-model="formData.configName" clearable placeholder="请输入配置名称" />
        </el-form-item>
        
        <el-form-item label="服务提供商:" prop="provider">
          <el-select v-model="formData.provider" placeholder="请选择服务提供商" style="width:100%" @change="onProviderChange">
            <el-option label="OpenAI" value="openai" />
            <el-option label="Claude" value="claude" />
            <el-option label="Gemini" value="gemini" />
            <el-option label="千问" value="qianwen" />
            <el-option label="文心" value="baidu" />
            <el-option label="混元" value="tencent" />
            <el-option label="本地模型" value="local" />
            <el-option label="自定义" value="custom" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="AI模型:" prop="model">
          <el-select v-model="formData.model" placeholder="请选择AI模型" style="width:100%" filterable allow-create>
            <el-option
              v-for="model in availableModels"
              :key="model.value"
              :label="model.label"
              :value="model.value"
            />
          </el-select>
        </el-form-item>
        
        <el-form-item label="API密钥:" prop="apiKey">
          <el-input
            v-model="formData.apiKey"
            type="password"
            show-password
            clearable
            placeholder="请输入API密钥"
          />
        </el-form-item>
        
        <el-form-item label="API端点:" prop="apiEndpoint">
          <el-input v-model="formData.apiEndpoint" clearable placeholder="请输入API端点URL" />
        </el-form-item>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="最大Token数:" prop="maxTokens">
              <el-input-number v-model="formData.maxTokens" :min="1" :max="32000" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="创造性参数:" prop="temperature">
              <el-slider v-model="formData.temperature" :min="0" :max="2" :step="0.1" show-input />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="核采样参数:" prop="topP">
              <el-slider v-model="formData.topP" :min="0" :max="1" :step="0.1" show-input />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="频率惩罚:" prop="frequencyPenalty">
              <el-slider v-model="formData.frequencyPenalty" :min="-2" :max="2" :step="0.1" show-input />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-form-item label="存在惩罚:" prop="presencePenalty">
          <el-slider v-model="formData.presencePenalty" :min="-2" :max="2" :step="0.1" show-input style="width: 48%" />
        </el-form-item>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="每日限制:" prop="dailyLimit">
              <el-input-number v-model="formData.dailyLimit" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="每月限制:" prop="monthlyLimit">
              <el-input-number v-model="formData.monthlyLimit" :min="1" style="width:100%" />
            </el-form-item>
          </el-col>
        </el-row>
        
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="状态:" prop="status">
              <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%">
                <el-option label="激活" value="active" />
                <el-option label="停用" value="inactive" />
                <el-option label="测试中" value="testing" />
                <el-option label="错误" value="error" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="设为默认:">
              <el-switch v-model="formData.isDefault" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-drawer>

    <!-- 测试结果对话框 -->
    <el-dialog v-model="testDialogVisible" title="配置测试结果" width="60%">
      <div class="test-result">
        <div class="test-info">
          <el-descriptions :column="2" border>
            <el-descriptions-item label="配置名称">{{ testResult.configName }}</el-descriptions-item>
            <el-descriptions-item label="服务提供商">{{ getProviderLabel(testResult.provider) }}</el-descriptions-item>
            <el-descriptions-item label="AI模型">{{ testResult.model }}</el-descriptions-item>
            <el-descriptions-item label="测试状态">
              <el-tag :type="testResult.success ? 'success' : 'danger'">
                {{ testResult.success ? '成功' : '失败' }}
              </el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="响应时间">{{ testResult.responseTime }}ms</el-descriptions-item>
            <el-descriptions-item label="Token使用">{{ testResult.tokenUsed }}</el-descriptions-item>
          </el-descriptions>
        </div>
        
        <div class="test-content" v-if="testResult.response">
          <h4>AI回复内容：</h4>
          <div class="response-content">{{ testResult.response }}</div>
        </div>
        
        <div class="test-error" v-if="testResult.error">
          <h4>错误信息：</h4>
          <el-alert :title="testResult.error" type="error" show-icon :closable="false" />
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createAIConfig,
  deleteAIConfig,
  updateAIConfig,
  findAIConfig,
  getAIConfigList,
  testAIConfig
} from '@/api/novel'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted, computed } from 'vue'

defineOptions({
  name: 'AIConfig'
})

const formData = ref({
  configName: '',
  provider: 'openai',
  model: 'gpt-3.5-turbo',
  apiKey: '',
  apiEndpoint: '',
  maxTokens: 2000,
  temperature: 0.7,
  topP: 1.0,
  frequencyPenalty: 0,
  presencePenalty: 0,
  dailyLimit: 1000,
  monthlyLimit: 30000,
  isDefault: false,
  status: 'active'
})

const rule = reactive({
  configName: [{
    required: true,
    message: '请输入配置名称',
    trigger: 'blur'
  }],
  provider: [{
    required: true,
    message: '请选择服务提供商',
    trigger: 'change'
  }],
  model: [{
    required: true,
    message: '请选择AI模型',
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
const testDialogVisible = ref(false)
const testResult = ref({})
const tableData = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)

// 可用模型列表
const modelOptions = {
  openai: [
    { label: 'GPT-3.5 Turbo', value: 'gpt-3.5-turbo' },
    { label: 'GPT-4', value: 'gpt-4' },
    { label: 'GPT-4 Turbo', value: 'gpt-4-turbo' },
    { label: 'GPT-4o', value: 'gpt-4o' }
  ],
  claude: [
    { label: 'Claude 3 Haiku', value: 'claude-3-haiku-20240307' },
    { label: 'Claude 3 Sonnet', value: 'claude-3-sonnet-20240229' },
    { label: 'Claude 3 Opus', value: 'claude-3-opus-20240229' }
  ],
  gemini: [
    { label: 'Gemini Pro', value: 'gemini-pro' },
    { label: 'Gemini Pro Vision', value: 'gemini-pro-vision' }
  ],
  qianwen: [
    { label: '通义千问-Turbo', value: 'qwen-turbo' },
    { label: '通义千问-Plus', value: 'qwen-plus' },
    { label: '通义千问-Max', value: 'qwen-max' }
  ],
  baidu: [
    { label: '文心一言', value: 'ernie-bot' },
    { label: '文心一言-Turbo', value: 'ernie-bot-turbo' },
    { label: '文心一言4.0', value: 'ernie-bot-4' }
  ],
  tencent: [
    { label: '混元-Standard', value: 'hunyuan-standard' },
    { label: '混元-Pro', value: 'hunyuan-pro' }
  ]
}

const availableModels = computed(() => {
  return modelOptions[formData.value.provider] || []
})

// 服务提供商变更
const onProviderChange = () => {
  const models = availableModels.value
  if (models.length > 0) {
    formData.value.model = models[0].value
  }
  
  // 设置默认端点
  const endpoints = {
    openai: 'https://api.openai.com/v1/chat/completions',
    claude: 'https://api.anthropic.com/v1/messages',
    gemini: 'https://generativelanguage.googleapis.com/v1/models',
    qianwen: 'https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation',
    baidu: 'https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat',
    tencent: 'https://hunyuan.tencentcloudapi.com'
  }
  
  formData.value.apiEndpoint = endpoints[formData.value.provider] || ''
}

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

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAIConfigList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 字典方法
const getProviderLabel = (value) => {
  const providerMap = {
    openai: 'OpenAI',
    claude: 'Claude',
    gemini: 'Gemini',
    qianwen: '千问',
    baidu: '文心',
    tencent: '混元',
    local: '本地模型',
    custom: '自定义'
  }
  return providerMap[value] || value
}

const getStatusLabel = (value) => {
  const statusMap = {
    active: '激活',
    inactive: '停用',
    testing: '测试中',
    error: '错误'
  }
  return statusMap[value] || value
}

const getProviderTagType = (value) => {
  const typeMap = {
    openai: 'primary',
    claude: 'success',
    gemini: 'warning',
    qianwen: 'info',
    baidu: 'danger',
    tencent: 'primary'
  }
  return typeMap[value] || ''
}

const getStatusTagType = (value) => {
  const typeMap = {
    active: 'success',
    inactive: 'info',
    testing: 'warning',
    error: 'danger'
  }
  return typeMap[value] || ''
}

// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 设置默认配置
const setDefault = async (row) => {
  try {
    const res = await updateAIConfig({ ...row, isDefault: true })
    if (res.code === 0) {
      ElMessage.success('设置默认配置成功')
      getTableData()
    }
  } catch (error) {
    ElMessage.error('设置默认配置失败')
    row.isDefault = false
  }
}

// 测试配置
const testConfig = async (row) => {
  try {
    ElMessage.info('正在测试配置...')
    const res = await testAIConfig({ configId: row.ID })
    
    testResult.value = {
      configName: row.configName,
      provider: row.provider,
      model: row.model,
      success: res.code === 0,
      responseTime: res.data?.responseTime || 0,
      tokenUsed: res.data?.tokenUsed || 0,
      response: res.data?.response || '',
      error: res.code !== 0 ? res.msg : ''
    }
    
    testDialogVisible.value = true
    
    if (res.code === 0) {
      ElMessage.success('配置测试成功')
    } else {
      ElMessage.error('配置测试失败')
    }
  } catch (error) {
    ElMessage.error('配置测试失败')
  }
}

// 批量测试
const testAllConfigs = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要测试的配置')
    return
  }
  
  ElMessage.info('正在批量测试配置...')
  
  for (const config of multipleSelection.value) {
    await testConfig(config)
  }
  
  ElMessage.success('批量测试完成')
}

// 复制配置
const copyConfig = (row) => {
  formData.value = {
    ...row,
    ID: undefined,
    configName: row.configName + ' (副本)',
    isDefault: false,
    status: 'inactive'
  }
  type.value = 'create'
  dialogFormVisible.value = true
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除这个配置吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteConfigFunc(row)
  })
}

// 批量删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除选中的配置吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteAIConfig({ IDs })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 编辑
const updateConfigFunc = async(row) => {
  const res = await findAIConfig({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.config
    dialogFormVisible.value = true
  }
}

// 删除
const deleteConfigFunc = async (row) => {
  const res = await deleteAIConfig({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
  }
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
    configName: '',
    provider: 'openai',
    model: 'gpt-3.5-turbo',
    apiKey: '',
    apiEndpoint: '',
    maxTokens: 2000,
    temperature: 0.7,
    topP: 1.0,
    frequencyPenalty: 0,
    presencePenalty: 0,
    dailyLimit: 1000,
    monthlyLimit: 30000,
    isDefault: false,
    status: 'active'
  }
}

// 确定
const enterDialog = async() => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createAIConfig(formData.value)
        break
      case 'update':
        res = await updateAIConfig(formData.value)
        break
      default:
        res = await createAIConfig(formData.value)
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
.test-result {
  padding: 20px;
}

.test-info {
  margin-bottom: 20px;
}

.test-content {
  margin-bottom: 20px;
}

.test-content h4 {
  margin: 0 0 10px 0;
  color: #303133;
}

.response-content {
  padding: 15px;
  background: #f5f7fa;
  border-radius: 6px;
  line-height: 1.6;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.test-error h4 {
  margin: 0 0 10px 0;
  color: #303133;
}
</style>