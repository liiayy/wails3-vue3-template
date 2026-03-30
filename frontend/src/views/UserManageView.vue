<script setup lang="ts">
import { ref, reactive, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { MessagePlugin, DialogPlugin } from 'tdesign-vue-next'
import {
  listUsers,
  registerUser,
  updateUser,
  deleteUser,
  exportUsers,
  type User,
} from '../api/user'
import { useAsyncAction, useDebounce } from '@/composables'
import { SearchIcon, AddIcon, EditIcon, DeleteIcon, DownloadIcon } from 'tdesign-icons-vue-next'

const { t } = useI18n()

// ========== 搜索与分页 ==========
const keyword = ref('')
const debouncedKeyword = useDebounce(keyword, 300)
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
})

// ========== 表格数据 ==========
const tableData = ref<User[]>([])
const tableLoading = ref(false)

async function fetchList() {
  tableLoading.value = true
  try {
    const res = await listUsers(debouncedKeyword.value, pagination.current, pagination.pageSize)
    tableData.value = res.items
    pagination.total = res.total
  } catch (err) {
    // handleResult 已自动弹出报错，此处仅处理本地状态（如关闭加载）
    console.error(err)
  } finally {
    tableLoading.value = false
  }
}

// 搜索关键词变化时重置到第一页
watch(debouncedKeyword, () => {
  pagination.current = 1
  fetchList()
})

// 分页变化
function onPageChange(pageInfo: any) {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
  fetchList()
}

// 初始化加载
fetchList()

// ========== 表格列定义 ==========
const columns = computed(() => [
  { colKey: 'id', title: t('users.colId'), width: 80, align: 'center' as const },
  { colKey: 'name', title: t('users.colName'), ellipsis: true },
  { colKey: 'email', title: t('users.colEmail'), ellipsis: true },
  { colKey: 'operation', title: t('users.colAction'), width: 180, align: 'center' as const },
])

// ========== 新增/编辑弹窗 ==========
const dialogVisible = ref(false)
const isEdit = ref(false)
const formData = reactive({ id: 0, name: '', email: '' })

function openCreateDialog() {
  isEdit.value = false
  formData.id = 0
  formData.name = ''
  formData.email = ''
  dialogVisible.value = true
}

function openEditDialog(row: User) {
  isEdit.value = true
  formData.id = row.id
  formData.name = row.name
  formData.email = row.email
  dialogVisible.value = true
}

const { execute: doSubmit, loading: submitLoading } = useAsyncAction(async () => {
  if (!formData.name) throw new Error(t('users.nameRequired'))
  if (!formData.email) throw new Error(t('users.emailRequired'))

  if (isEdit.value) {
    await updateUser(formData.id, formData.name, formData.email)
    MessagePlugin.success(t('users.updateSuccess'))
  } else {
    await registerUser(formData.name, formData.email)
    MessagePlugin.success(t('users.createSuccess'))
  }
  dialogVisible.value = false
  fetchList()
})

// ========== 删除 ==========
function handleDelete(row: User) {
  const confirmDialog = DialogPlugin.confirm({
    header: t('users.deleteConfirmTitle'),
    body: t('users.deleteConfirmBody').replace('{name}', row.name),
    theme: 'danger',
    onConfirm: async () => {
      try {
        await deleteUser(row.id)
        MessagePlugin.success(t('users.deleteSuccess'))
        fetchList()
      } catch (err) {
        // handleResult 已处理报错
      }
      confirmDialog.destroy()
    },
  })
}

// ========== 导出 ==========
const { execute: handleExport, loading: exportLoading } = useAsyncAction(async () => {
  await exportUsers()
  MessagePlugin.success(t('users.exportSuccess'))
})
</script>

<template>
  <div class="user-manage-container p-2 space-y-6">
    <!-- 1. 顶部标题 & 全局操作 -->
    <section class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-bold text-[var(--td-text-color-primary)] tracking-tight">
          {{ $t('users.title') }}
        </h1>
        <p class="text-[var(--td-text-color-secondary)] text-sm font-normal mt-1">
           管理系统中的所有用户信息，支持搜索、新增、编辑及数据导出
        </p>
      </div>
      <div class="flex items-center gap-2">
        <t-button
          variant="outline"
          theme="default"
          :loading="exportLoading"
          @click="handleExport"
        >
          <template #icon><DownloadIcon /></template>
          {{ $t('users.exportBtn') }}
        </t-button>
        <t-button theme="primary" @click="openCreateDialog">
          <template #icon><AddIcon /></template>
          {{ $t('users.addUser') }}
        </t-button>
      </div>
    </section>

    <!-- 2. 核心内容区 (过滤 + 表格) -->
    <t-card :bordered="false" class="shadow-md rounded-xl overflow-hidden">
      <!-- 过滤栏 -->
      <div class="flex flex-wrap items-center justify-between gap-4 mb-6">
        <div class="flex items-center gap-4">
          <t-input
            v-model="keyword"
            :placeholder="$t('users.searchPlaceholder')"
            clearable
            style="width: 280px"
          >
            <template #prefixIcon><SearchIcon /></template>
          </t-input>
        </div>
        
        <div class="text-[var(--td-text-color-placeholder)] text-sm">
          共 <span class="text-[var(--td-brand-color)] font-medium">{{ pagination.total }}</span> 个用户
        </div>
      </div>

      <!-- 表格 -->
      <t-table
        :data="tableData"
        :columns="columns"
        :loading="tableLoading"
        row-key="id"
        stripe
        hover
        vertical-align="middle"
        :pagination="pagination"
        @page-change="onPageChange"
        class="user-table"
      >
        <!-- 头像/名称列增强 -->
        <template #name="{ row }">
          <div class="flex items-center gap-3">
            <t-avatar size="small" :hide-on-load-failed="false">
              {{ row.name.charAt(0).toUpperCase() }}
            </t-avatar>
            <span class="font-medium text-[var(--td-text-color-primary)]">{{ row.name }}</span>
          </div>
        </template>

        <!-- 邮箱列 -->
        <template #email="{ row }">
          <span class="text-[var(--td-text-color-secondary)]">{{ row.email }}</span>
        </template>

        <!-- 操作列 -->
        <template #operation="{ row }">
          <div class="flex items-center gap-1">
            <t-tooltip content="编辑用户信息">
              <t-button variant="text" theme="primary" shape="square" @click="openEditDialog(row)">
                <EditIcon />
              </t-button>
            </t-tooltip>
            <t-tooltip content="删除该用户">
              <t-button variant="text" theme="danger" shape="square" @click="handleDelete(row)">
                <DeleteIcon />
              </t-button>
            </t-tooltip>
          </div>
        </template>

        <!-- 空状态 -->
        <template #empty>
          <div class="flex flex-col items-center justify-center py-12 opacity-50">
            <div class="text-4xl mb-2">🔍</div>
            <p>暂无符合搜索条件的用户</p>
          </div>
        </template>
      </t-table>
    </t-card>

    <!-- 3. 新增/编辑弹窗 -->
    <t-dialog
      v-model:visible="dialogVisible"
      :header="isEdit ? $t('users.editUser') : $t('users.addUser')"
      :confirm-btn="{
        content: isEdit ? $t('common.save') : $t('common.create'),
        loading: submitLoading,
        theme: 'primary',
      }"
      :cancel-btn="$t('common.cancel')"
      :on-confirm="() => doSubmit()"
      width="520px"
      placement="center"
      destroy-on-close
    >
      <div class="pt-2">
        <t-form :data="formData" label-align="top" colon>
          <t-form-item :label="$t('users.colName')" name="name" help="支持中英文，不可超过32个字符">
            <t-input 
              v-model="formData.name" 
              :placeholder="$t('users.namePlaceholder')" 
              clearable 
              autofocus
            />
          </t-form-item>
          <t-form-item :label="$t('users.colEmail')" name="email" help="用于接收系统通知和找回密码">
            <t-input v-model="formData.email" :placeholder="$t('users.emailPlaceholder')" clearable />
          </t-form-item>
        </t-form>
      </div>
    </t-dialog>
  </div>
</template>

<style scoped>
.user-manage-container {
  max-width: 1400px;
  margin: 0 auto;
}

.user-table :deep(.t-table__header tr) {
  background-color: var(--td-bg-color-secondarycontainer);
}

.user-table :deep(.t-table__content) {
  border-radius: 8px;
}

/* 自定义卡片样式 */
:deep(.t-card) {
  padding: 24px;
}

/* 响应式调整 */
@media (max-width: 640px) {
  .user-manage-container {
    padding: 12px;
  }
}
</style>
