<template>
  <n-config-provider
    :theme="darkTheme"
    :theme-overrides="themeOverrides"
  >
    <n-layout position="absolute">
      <n-menu
        mode="horizontal"
        :options="menuOptions"
      />

      <div style="padding: 6px 20px;">
        <RouterView />
      </div>
    </n-layout>
  </n-config-provider>
</template>

<script setup>
import { darkTheme } from 'naive-ui'
import { h } from 'vue'
import { RouterLink } from 'vue-router'

const menuOptions = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'home'
          }
        },
        { default: () => 'Studies' }
      ),
    key: 'home'
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'about'
          }
        },
        { default: () => 'About' }
      ),
    key: 'about'
  }
]
const themeOverrides = {
  Card: {
    borderRadius: '16px'
  }
}
</script>
