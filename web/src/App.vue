<template>
  <n-config-provider
    :theme="darkTheme"
    :theme-overrides="themeOverrides"
  >
    <n-layout
      content-style="max-width: 1024px; margin-left: auto; margin-right: auto;"
    >
      <n-menu
        style="padding-top: 8px"
        mode="horizontal"
        responsive
        :options="menuOptions"
      />

      <n-layout-content
        content-style="padding: 6px 21px;"
      >
        <RouterView />
      </n-layout-content>
    </n-layout>
  </n-config-provider>
</template>

<script setup>
import {
  OpenOutline
} from '@vicons/ionicons5'
import { darkTheme, NIcon } from 'naive-ui'
import { h } from 'vue'
import { RouterLink } from 'vue-router'

function renderIcon (icon) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

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
  },
  {
    label: () =>
      h(
        'a',
        {
          href: 'https://github.com/ndabAP/entityscrape',
          target: '_blank'
        },
        'GitHub'
      ),
    key: 'github',
    icon: renderIcon(OpenOutline)
  }
]
const themeOverrides = {
  Alert: {
    borderRadius: '16px'
  },
  Card: {
    borderRadius: '16px'
  }
}
</script>

<style>
html {
  background-color: rgb(16, 16, 20);
}

.n-statistic .n-statistic-value .n-statistic-value__content {
  font-size: unset !important;
}
</style>
