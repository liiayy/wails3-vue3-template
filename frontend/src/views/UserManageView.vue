<script setup lang="ts">
import { ref, reactive, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { MessagePlugin, DialogPlugin } from 'tdesign-vue-next'
import { listUsers, registerUser, updateUser, deleteUser, type User } from '../api/user'
import { useAsyncAction, useDebounce } from '../composables'
import { SearchIcon, AddIcon, EditIcon, DeleteIcon } from 'tdesign-icons-vue-next'

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
    tableData.value = res?.items || []
    pagination.total = res?.total || 0
  } catch (err) {
    MessagePlugin.error(t('users.loadFailed') + ': ' + String(err))
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
        MessagePlugin.error(t('users.deleteFailed') + ': ' + String(err))
      }
      confirmDialog.destroy()
    },
  })
}
</script>

<template>
  <div class="space-y-5">
    <!-- 页面标题 + 操作栏 -->
    <div class="flex items-center justify-between">
      <h2 class="text-xl font-bold text-[var(--td-text-color-primary)]">{{ $t('users.title') }}</h2>

      <div class="flex items-center gap-3">
        <t-input
          v-model="keyword"
          :placeholder="$t('users.searchPlaceholder')"
          clearable
          style="width: 260px"
        >
          <template #prefixIcon><SearchIcon /></template>
        </t-input>

        <t-button theme="primary" @click="openCreateDialog">
          <template #icon><AddIcon /></template>
          {{ $t('users.addUser') }}
        </t-button>
      </div>
    </div>

    <!-- 数据表格 -->
    <t-card :bordered="false" class="shadow-sm">
      <t-table
        :data="tableData"
        :columns="columns"
        :loading="tableLoading"
        row-key="id"
        stripe
        hover
        :pagination="pagination"
        @page-change="onPageChange"
      >
        <!-- 操作列 -->
        <template #operation="{ row }">
          <div class="flex items-center justify-center gap-2">
            <t-button variant="text" theme="primary" size="small" @click="openEditDialog(row)">
              <template #icon><EditIcon /></template>
              {{ $t('common.edit') }}
            </t-button>
            <t-button variant="text" theme="danger" size="small" @click="handleDelete(row)">
              <template #icon><DeleteIcon /></template>
              {{ $t('common.delete') }}
            </t-button>
          </div>
        </template>
      </t-table>
    </t-card>

    <!-- 新增/编辑弹窗 -->
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
      width="480px"
    >
      <t-form :data="formData" label-align="top" class="space-y-4">
        <t-form-item :label="$t('users.colName')" name="name">
          <t-input v-model="formData.name" :placeholder="$t('users.namePlaceholder')" clearable />
        </t-form-item>
        <t-form-item :label="$t('users.colEmail')" name="email">
          <t-input v-model="formData.email" :placeholder="$t('users.emailPlaceholder')" clearable />
        </t-form-item>
      </t-form>
    </t-dialog>
  </div>
</template>
