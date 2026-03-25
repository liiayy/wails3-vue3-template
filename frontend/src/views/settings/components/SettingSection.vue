<script setup lang="ts">
export interface SettingItem {
  key: string
  label: string
  description?: string
  type?: 'select' | 'switch' | 'slot'
  value?: any
  options?: Array<{ value: string; label: string }>
}

interface Props {
  title: string
  items?: SettingItem[]
}

defineProps<Props>()
</script>

<template>
  <section class="space-y-2">
    <div class="text-[15px] font-bold px-6">
      {{ title }}
    </div>

    <div class="bg-[var(--td-bg-color-container)] rounded-2xl shadow-sm px-6 py-2">
      <slot>
        <!-- 默认渲染：基于 items 配置 -->
        <div
          v-for="item in items"
          :key="item.key"
          class="flex items-center justify-between py-3 border-b border-[var(--td-component-border)] last:border-0"
        >
          <div class="flex flex-col gap-1">
            <span class="text-[14px] text-[var(--td-text-color-primary)]">{{ item.label }}</span>
            <span v-if="item.description" class="text-[12px] text-[var(--td-text-color-placeholder)]">
              {{ item.description }}
            </span>
          </div>

          <!-- Select 类型 -->
          <t-select
            v-if="item.type === 'select'"
            :value="item.value"
            class="!w-[200px]"
            variant="outline"
            size="small"
            @change="$emit('change', item.key, $event)"
          >
            <t-option
              v-for="opt in item.options"
              :key="opt.value"
              :value="opt.value"
              :label="opt.label"
            />
          </t-select>

          <!-- Switch 类型 -->
          <t-switch
            v-else-if="item.type === 'switch'"
            :value="item.value"
            size="large"
            @change="$emit('change', item.key, $event)"
          />

          <!-- 自定义插槽类型 -->
          <slot v-else-if="item.type === 'slot'" :name="item.key" :item="item" />
        </div>
      </slot>
    </div>
  </section>
</template>
