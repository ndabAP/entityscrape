<template>
  <n-config-provider
    :theme="darkTheme"
    :theme-overrides="themeOverrides"
  >
    <n-layout
      position="absolute"
      content-style="max-width: 1280px; margin-left: auto; margin-right: auto;"
    >
      <n-menu
        style="padding-top: 8px"
        mode="horizontal"
        :options="menuOptions"
      />

      <n-layout-content
        content-style="padding: 6px 24px;"
      >
        <RouterView />
      </n-layout-content>
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

<style>
p {
  text-align: justify;
}
</style>
