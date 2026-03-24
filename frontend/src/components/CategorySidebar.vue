<script setup lang="ts">
import { Icon } from 'tdesign-icons-vue-next'

interface Category {
  id: string
  label: string
  icon: string
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
    class="w-44 border-r border-[var(--td-border-level-1-color)] bg-[var(--td-bg-color-container)] shrink-0 flex flex-col"
  >
    <div class="p-2">
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
          <icon :name="cat.icon" size="16" />
          <span class="text-sm">
            {{ $t(cat.label) }}
          </span>
        </div>
      </nav>
    </div>
  </aside>
</template>
