<template>
  <n-config-provider :theme="darkTheme">
    <n-layout position="absolute">
      <n-menu
        mode="horizontal"
        :options="menuOptions"
      />

      <div style="padding: 0 20px;">
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
</script>

<style>

</style>
