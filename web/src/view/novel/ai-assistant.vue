<template>
  <div class="ai-assistant-container">
    <!-- AI助手头部 -->
    <div class="ai-header">
      <div class="header-left">
        <el-avatar :size="40" src="/ai-avatar.png">
          <el-icon><Robot /></el-icon>
        </el-avatar>
        <div class="header-info">
          <h3>AI写作助手</h3>
          <span class="status-text">在线 · 随时为您服务</span>
        </div>
      </div>
      <div class="header-right">
        <el-button-group>
          <el-button :type="activeTab === 'chat' ? 'primary' : ''" @click="activeTab = 'chat'">智能对话</el-button>
          <el-button :type="activeTab === 'generate' ? 'primary' : ''" @click="activeTab = 'generate'">内容生成</el-button>
          <el-button :type="activeTab === 'tools' ? 'primary' : ''" @click="activeTab = 'tools'">写作工具</el-button>
        </el-button-group>
      </div>
    </div>

    <!-- 智能对话模块 -->
    <div v-show="activeTab === 'chat'" class="chat-module">
      <div class="chat-container">
        <!-- 聊天历史 -->
        <div class="chat-history" ref="chatHistoryRef">
          <div v-for="message in chatHistory" :key="message.id" class="message-item" :class="message.role">
            <div class="message-avatar">
              <el-avatar v-if="message.role === 'user'" :size="32">
                <el-icon><User /></el-icon>
              </el-avatar>
              <el-avatar v-else :size="32">
                <el-icon><Robot /></el-icon>
              </el-avatar>
            </div>
            <div class="message-content">
              <div class="message-text" v-html="formatMessage(message.content)"></div>
              <div class="message-time">{{ formatTime(message.timestamp) }}</div>
              <div v-if="message.role === 'assistant'" class="message-actions">
                <el-button link size="small" @click="copyMessage(message.content)">
                  <el-icon><CopyDocument /></el-icon>
                </el-button>
                <el-button link size="small" @click="rateMessage(message)">
                  <el-icon><Star /></el-icon>
                </el-button>
                <el-button link size="small" @click="insertToEditor(message.content)">
                  <el-icon><DocumentAdd /></el-icon>
                </el-button>
              </div>
            </div>
          </div>
          <div v-if="isTyping" class="typing-indicator">
            <div class="message-avatar">
              <el-avatar :size="32">
                <el-icon><Robot /></el-icon>
              </el-avatar>
            </div>
            <div class="typing-dots">
              <span></span><span></span><span></span>
            </div>
          </div>
        </div>

        <!-- 输入区域 -->
        <div class="chat-input">
          <div class="input-toolbar">
            <el-button-group size="small">
              <el-button @click="insertQuickPrompt('写作建议')">写作建议</el-button>
              <el-button @click="insertQuickPrompt('情节发展')">情节发展</el-button>
              <el-button @click="insertQuickPrompt('角色分析')">角色分析</el-button>
              <el-button @click="insertQuickPrompt('对话优化')">对话优化</el-button>
            </el-button-group>
          </div>
          <div class="input-area">
            <el-input
              v-model="userInput"
              type="textarea"
              :autosize="{ minRows: 2, maxRows: 6 }"
              placeholder="请输入您的问题或需求..."
              @keydown.ctrl.enter="sendMessage"
            />
            <div class="input-actions">
              <el-button type="primary" :loading="isTyping" @click="sendMessage">
                <el-icon><Promotion /></el-icon>
                发送 (Ctrl+Enter)
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 内容生成模块 -->
    <div v-show="activeTab === 'generate'" class="generate-module">
      <div class="generate-categories">
        <el-row :gutter="20">
          <el-col :span="8" v-for="category in generateCategories" :key="category.type">
            <el-card class="category-card" shadow="hover" @click="selectCategory(category)">
              <div class="category-icon">
                <el-icon :size="32"><component :is="category.icon" /></el-icon>
              </div>
              <h4>{{ category.title }}</h4>
              <p>{{ category.description }}</p>
            </el-card>
          </el-col>
        </el-row>
      </div>

      <!-- 生成表单 -->
      <el-dialog v-model="generateDialogVisible" :title="selectedCategory?.title" width="60%" destroy-on-close>
        <el-form :model="generateForm" label-width="100px">
          <el-form-item v-for="field in selectedCategory?.fields" :key="field.key" :label="field.label">
            <el-input
              v-if="field.type === 'input'"
              v-model="generateForm[field.key]"
              :placeholder="field.placeholder"
            />
            <el-input
              v-else-if="field.type === 'textarea'"
              v-model="generateForm[field.key]"
              type="textarea"
              :autosize="{ minRows: 3, maxRows: 6 }"
              :placeholder="field.placeholder"
            />
            <el-select
              v-else-if="field.type === 'select'"
              v-model="generateForm[field.key]"
              :placeholder="field.placeholder"
              style="width: 100%"
            >
              <el-option
                v-for="option in field.options"
                :key="option.value"
                :label="option.label"
                :value="option.value"
              />
            </el-select>
            <el-input-number
              v-else-if="field.type === 'number'"
              v-model="generateForm[field.key]"
              :min="field.min || 0"
              :max="field.max || 10000"
              style="width: 100%"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="generateDialogVisible = false">取消</el-button>
            <el-button type="primary" :loading="isGenerating" @click="generateContent">生成内容</el-button>
          </span>
        </template>
      </el-dialog>

      <!-- 生成结果 -->
      <div v-if="generateResult" class="generate-result">
        <el-card>
          <template #header>
            <div class="result-header">
              <span>生成结果</span>
              <div>
                <el-button size="small" @click="copyMessage(generateResult)">复制</el-button>
                <el-button size="small" type="primary" @click="insertToEditor(generateResult)">插入编辑器</el-button>
              </div>
            </div>
          </template>
          <div class="result-content" v-html="formatMessage(generateResult)"></div>
        </el-card>
      </div>
    </div>

    <!-- 写作工具模块 -->
    <div v-show="activeTab === 'tools'" class="tools-module">
      <div class="tools-grid">
        <el-row :gutter="20">
          <el-col :span="6" v-for="tool in writingTools" :key="tool.type">
            <el-card class="tool-card" shadow="hover" @click="useTool(tool)">
              <div class="tool-icon">
                <el-icon :size="24"><component :is="tool.icon" /></el-icon>
              </div>
              <h5>{{ tool.title }}</h5>
              <p>{{ tool.description }}</p>
            </el-card>
          </el-col>
        </el-row>
      </div>

      <!-- 工具使用对话框 -->
      <el-dialog v-model="toolDialogVisible" :title="selectedTool?.title" width="50%" destroy-on-close>
        <el-form :model="toolForm" label-width="80px">
          <el-form-item label="内容:" v-if="selectedTool?.needContent">
            <el-input
              v-model="toolForm.content"
              type="textarea"
              :autosize="{ minRows: 4, maxRows: 8 }"
              placeholder="请输入要处理的内容"
            />
          </el-form-item>
          <el-form-item label="要求:">
            <el-input
              v-model="toolForm.requirement"
              type="textarea"
              :autosize="{ minRows: 2, maxRows: 4 }"
              placeholder="请输入具体要求"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="toolDialogVisible = false">取消</el-button>
            <el-button type="primary" :loading="isProcessing" @click="processTool">处理</el-button>
          </span>
        </template>
      </el-dialog>
    </div>

    <!-- 评分对话框 -->
    <el-dialog v-model="ratingDialogVisible" title="评价AI回复" width="400px">
      <div class="rating-content">
        <div class="rating-stars">
          <el-rate v-model="rating" show-text :texts="['很差', '较差', '一般', '较好', '很好']" />
        </div>
        <el-input
          v-model="feedback"
          type="textarea"
          :autosize="{ minRows: 3, maxRows: 5 }"
          placeholder="请输入您的反馈意见（可选）"
        />
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="ratingDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitRating">提交评价</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  chatWithAI,
  getChatHistory,
  submitFeedback,
  writingAssist,
  plotGenerate,
  characterGenerate,
  dialogueGenerate,
  continueWriting
} from '@/api/novel'
import { ElMessage } from 'element-plus'
import { ref, reactive, onMounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import {
  Cpu,
  User,
  CopyDocument,
  Star,
  DocumentAdd,
  Promotion,
  Edit,
  UserFilled,
  ChatDotRound,
  Document,
  Reading,
  MagicStick,
  Tools,
  Search,
  Refresh
} from '@element-plus/icons-vue'

defineOptions({
  name: 'AIAssistant'
})

const route = useRoute()
const novelId = route.params.novelId

const activeTab = ref('chat')
const userInput = ref('')
const chatHistory = ref([])
const isTyping = ref(false)
const isGenerating = ref(false)
const isProcessing = ref(false)
const chatHistoryRef = ref()
const sessionId = ref('novel_' + novelId + '_' + Date.now())

// 生成相关
const generateDialogVisible = ref(false)
const selectedCategory = ref(null)
const generateForm = ref({})
const generateResult = ref('')

// 工具相关
const toolDialogVisible = ref(false)
const selectedTool = ref(null)
const toolForm = ref({
  content: '',
  requirement: ''
})

// 评分相关
const ratingDialogVisible = ref(false)
const currentMessage = ref(null)
const rating = ref(5)
const feedback = ref('')

// 生成分类
const generateCategories = ref([
  {
    type: 'plot',
    title: '情节生成',
    description: '生成精彩的故事情节和转折',
    icon: 'Document',
    fields: [
      { key: 'plotType', label: '情节类型', type: 'select', options: [
        { label: '开场情节', value: '开场' },
        { label: '冲突情节', value: '冲突' },
        { label: '高潮情节', value: '高潮' },
        { label: '结局情节', value: '结局' }
      ]},
      { key: 'characters', label: '相关角色', type: 'input', placeholder: '请输入相关角色' },
      { key: 'setting', label: '场景设定', type: 'textarea', placeholder: '请描述场景设定' },
      { key: 'conflict', label: '冲突设定', type: 'textarea', placeholder: '请描述冲突设定' },
      { key: 'tone', label: '情感基调', type: 'input', placeholder: '如：紧张、温馨、悲伤' },
      { key: 'length', label: '生成长度', type: 'number', min: 100, max: 2000 }
    ]
  },
  {
    type: 'character',
    title: '角色生成',
    description: '创建丰富立体的小说角色',
    icon: 'User',
    fields: [
      { key: 'characterType', label: '角色类型', type: 'select', options: [
        { label: '主角', value: '主角' },
        { label: '配角', value: '配角' },
        { label: '反派', value: '反派' },
        { label: '路人', value: '路人' }
      ]},
      { key: 'role', label: '角色定位', type: 'input', placeholder: '请输入角色定位' },
      { key: 'personality', label: '性格特点', type: 'textarea', placeholder: '请描述性格特点' },
      { key: 'background', label: '背景设定', type: 'textarea', placeholder: '请描述背景设定' },
      { key: 'relationships', label: '关系网络', type: 'textarea', placeholder: '请描述与其他角色的关系' }
    ]
  },
  {
    type: 'dialogue',
    title: '对话生成',
    description: '生成自然流畅的角色对话',
    icon: 'ChatDotRound',
    fields: [
      { key: 'characters', label: '对话角色', type: 'input', placeholder: '请输入对话角色，用逗号分隔' },
      { key: 'situation', label: '对话情境', type: 'textarea', placeholder: '请描述对话发生的情境' },
      { key: 'purpose', label: '对话目的', type: 'input', placeholder: '请输入对话目的' },
      { key: 'tone', label: '对话基调', type: 'input', placeholder: '如：轻松、严肃、紧张' },
      { key: 'length', label: '对话长度', type: 'number', min: 50, max: 1000 }
    ]
  },
  {
    type: 'continue',
    title: '智能续写',
    description: '根据现有内容智能续写',
    icon: 'Edit',
    fields: [
      { key: 'content', label: '当前内容', type: 'textarea', placeholder: '请输入当前内容' },
      { key: 'direction', label: '续写方向', type: 'textarea', placeholder: '请描述希望的续写方向' },
      { key: 'style', label: '写作风格', type: 'input', placeholder: '请输入写作风格' },
      { key: 'length', label: '续写长度', type: 'number', min: 100, max: 2000 }
    ]
  }
])

// 写作工具
const writingTools = ref([
  {
    type: 'optimize',
    title: '内容优化',
    description: '优化文本表达和语言风格',
    icon: 'MagicStick',
    needContent: true
  },
  {
    type: 'review',
    title: '内容审查',
    description: '检查语法、逻辑和连贯性',
    icon: 'Search',
    needContent: true
  },
  {
    type: 'style',
    title: '风格分析',
    description: '分析文本风格和特点',
    icon: 'Document',
    needContent: true
  },
  {
    type: 'expand',
    title: '内容扩展',
    description: '扩展和丰富现有内容',
    icon: 'DocumentAdd',
    needContent: true
  }
])

// 发送消息
const sendMessage = async () => {
  if (!userInput.value.trim()) {
    ElMessage.warning('请输入消息内容')
    return
  }

  const message = userInput.value
  userInput.value = ''

  // 添加用户消息到历史
  chatHistory.value.push({
    id: Date.now(),
    role: 'user',
    content: message,
    timestamp: new Date()
  })

  scrollToBottom()
  isTyping.value = true

  try {
    const res = await chatWithAI({
      sessionId: sessionId.value,
      novelId: novelId,
      userMessage: message,
      messageType: 'chat',
      functionType: 'writing'
    })

    if (res.code === 0) {
      // 添加AI回复到历史
      chatHistory.value.push({
        id: res.data.assistant.id,
        role: 'assistant',
        content: res.data.assistant.aiResponse,
        timestamp: new Date(res.data.assistant.createdAt)
      })
      scrollToBottom()
    }
  } catch (error) {
    ElMessage.error('发送消息失败')
  } finally {
    isTyping.value = false
  }
}

// 插入快速提示
const insertQuickPrompt = (prompt) => {
  userInput.value = `请帮我${prompt}：`
}

// 选择生成分类
const selectCategory = (category) => {
  selectedCategory.value = category
  generateForm.value = {}
  generateDialogVisible.value = true
}

// 生成内容
const generateContent = async () => {
  isGenerating.value = true
  
  try {
    let res
    const type = selectedCategory.value.type
    
    switch (type) {
      case 'plot':
        res = await plotGenerate({
          novelId: novelId,
          ...generateForm.value
        })
        break
      case 'character':
        res = await characterGenerate({
          novelId: novelId,
          ...generateForm.value
        })
        break
      case 'dialogue':
        res = await dialogueGenerate({
          novelId: novelId,
          ...generateForm.value
        })
        break
      case 'continue':
        res = await continueWriting({
          novelId: novelId,
          ...generateForm.value
        })
        break
    }
    
    if (res.code === 0) {
      generateResult.value = res.data.assistant.aiResponse
      generateDialogVisible.value = false
      ElMessage.success('内容生成成功')
    }
  } catch (error) {
    ElMessage.error('内容生成失败')
  } finally {
    isGenerating.value = false
  }
}

// 使用工具
const useTool = (tool) => {
  selectedTool.value = tool
  toolForm.value = {
    content: '',
    requirement: ''
  }
  toolDialogVisible.value = true
}

// 处理工具
const processTool = async () => {
  isProcessing.value = true
  
  try {
    const res = await writingAssist({
      novelId: novelId,
      content: toolForm.value.content,
      assistType: selectedTool.value.type,
      requirement: toolForm.value.requirement
    })
    
    if (res.code === 0) {
      generateResult.value = res.data.assistant.aiResponse
      toolDialogVisible.value = false
      ElMessage.success('处理完成')
    }
  } catch (error) {
    ElMessage.error('处理失败')
  } finally {
    isProcessing.value = false
  }
}

// 复制消息
const copyMessage = async (content) => {
  try {
    await navigator.clipboard.writeText(content)
    ElMessage.success('已复制到剪贴板')
  } catch (error) {
    ElMessage.error('复制失败')
  }
}

// 评价消息
const rateMessage = (message) => {
  currentMessage.value = message
  rating.value = 5
  feedback.value = ''
  ratingDialogVisible.value = true
}

// 提交评价
const submitRating = async () => {
  try {
    await submitFeedback({
      assistantId: currentMessage.value.id,
      userRating: rating.value,
      feedback: feedback.value
    })
    
    ElMessage.success('评价提交成功')
    ratingDialogVisible.value = false
  } catch (error) {
    ElMessage.error('评价提交失败')
  }
}

// 插入到编辑器
const insertToEditor = (content) => {
  // 这里可以实现插入到章节编辑器的逻辑
  ElMessage.success('内容已插入编辑器')
}

// 格式化消息
const formatMessage = (content) => {
  return content.replace(/\n/g, '<br>')
}

// 格式化时间
const formatTime = (time) => {
  return new Date(time).toLocaleTimeString()
}

// 滚动到底部
const scrollToBottom = () => {
  nextTick(() => {
    if (chatHistoryRef.value) {
      chatHistoryRef.value.scrollTop = chatHistoryRef.value.scrollHeight
    }
  })
}

// 加载聊天历史
const loadChatHistory = async () => {
  try {
    const res = await getChatHistory({
      sessionId: sessionId.value,
      page: 1,
      pageSize: 50
    })
    
    if (res.code === 0) {
      chatHistory.value = res.data.list.reverse().map(item => ({
        id: item.id,
        role: 'user',
        content: item.userMessage,
        timestamp: new Date(item.createdAt)
      })).concat(res.data.list.reverse().map(item => ({
        id: item.id,
        role: 'assistant',
        content: item.aiResponse,
        timestamp: new Date(item.createdAt)
      })))
      
      scrollToBottom()
    }
  } catch (error) {
    console.error('加载聊天历史失败:', error)
  }
}

onMounted(() => {
  loadChatHistory()
})
</script>

<style scoped>
.ai-assistant-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: #f5f7fa;
}

.ai-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background: white;
  border-bottom: 1px solid #e4e7ed;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-info h3 {
  margin: 0;
  color: #303133;
}

.status-text {
  font-size: 12px;
  color: #909399;
}

.chat-module {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.chat-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}

.chat-history {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  background: white;
  margin: 20px;
  border-radius: 8px;
}

.message-item {
  display: flex;
  margin-bottom: 20px;
  align-items: flex-start;
}

.message-item.user {
  flex-direction: row-reverse;
}

.message-avatar {
  margin: 0 10px;
}

.message-content {
  max-width: 70%;
  background: #f0f2f5;
  padding: 12px 16px;
  border-radius: 12px;
  position: relative;
}

.message-item.user .message-content {
  background: #409eff;
  color: white;
}

.message-text {
  line-height: 1.6;
  word-wrap: break-word;
}

.message-time {
  font-size: 11px;
  color: #909399;
  margin-top: 5px;
}

.message-item.user .message-time {
  color: rgba(255, 255, 255, 0.8);
}

.message-actions {
  display: flex;
  gap: 5px;
  margin-top: 8px;
}

.typing-indicator {
  display: flex;
  align-items: center;
}

.typing-dots {
  display: flex;
  gap: 4px;
  padding: 12px 16px;
  background: #f0f2f5;
  border-radius: 12px;
}

.typing-dots span {
  width: 6px;
  height: 6px;
  background: #909399;
  border-radius: 50%;
  animation: typing 1.4s infinite;
}

.typing-dots span:nth-child(2) {
  animation-delay: 0.2s;
}

.typing-dots span:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes typing {
  0%, 60%, 100% {
    transform: translateY(0);
  }
  30% {
    transform: translateY(-10px);
  }
}

.chat-input {
  background: white;
  padding: 20px;
  margin: 0 20px 20px;
  border-radius: 8px;
}

.input-toolbar {
  margin-bottom: 10px;
}

.input-area {
  display: flex;
  gap: 10px;
  align-items: flex-end;
}

.input-area .el-textarea {
  flex: 1;
}

.generate-module {
  padding: 20px;
}

.generate-categories {
  margin-bottom: 20px;
}

.category-card {
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
  height: 150px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.category-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
}

.category-icon {
  color: #409eff;
  margin-bottom: 10px;
}

.category-card h4 {
  margin: 0 0 8px 0;
  color: #303133;
}

.category-card p {
  margin: 0;
  color: #606266;
  font-size: 12px;
}

.generate-result {
  margin-top: 20px;
}

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.result-content {
  line-height: 1.8;
  white-space: pre-wrap;
}

.tools-module {
  padding: 20px;
}

.tools-grid {
  margin-bottom: 20px;
}

.tool-card {
  text-align: center;
  cursor: pointer;
  transition: all 0.3s;
  height: 120px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.tool-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.1);
}

.tool-icon {
  color: #67c23a;
  margin-bottom: 8px;
}

.tool-card h5 {
  margin: 0 0 5px 0;
  color: #303133;
}

.tool-card p {
  margin: 0;
  color: #606266;
  font-size: 11px;
}

.rating-content {
  text-align: center;
}

.rating-stars {
  margin-bottom: 20px;
}
</style>