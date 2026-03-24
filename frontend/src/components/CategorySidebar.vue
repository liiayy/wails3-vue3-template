<script setup lang="ts">
interface Category {
  id: string
  label: string
  icon: any // 此处接收图标组件对象，以支持离线按需加载
  path: string
}

defineProps<{
  categories: Category[]
  activeValue: string
}>()

const emit = defineEmits<{
  (e: 'change', path: string): void
}>()
</script>

<template>
  <aside
    class="w-52 border-r border-[var(--td-border-level-1-color)] bg-[var(--td-bg-color-container)] shrink-0 flex flex-col"
  >
    <div class="p-1">
      <nav class="space-y-1">
        <div
          v-for="cat in categories"
          :key="cat.id"
          @click="emit('change', cat.path)"
          class="flex items-center gap-2.5 px-3 py-2 rounded-lg cursor-pointer transition-all duration-200"
          :class="
            activeValue === cat.id
              ? 'bg-[var(--td-brand-color-light)] text-[var(--td-brand-color)] font-medium shadow-sm'
              : 'text-[var(--td-text-color-secondary)] hover:bg-[var(--td-bg-color-secondarycontainer)]'
          "
        >
          <!-- 核心改动：使用动态组件渲染导入的对象 -->
          <component :is="cat.icon" size="16" />
          <span class="text-sm">
            {{ $t(cat.label) }}
          </span>
        </div>
      </nav>
    </div>
  </aside>
</template>
