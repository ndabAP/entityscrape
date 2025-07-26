<template>
  <n-config-provider :theme="darkTheme">
    <n-layout position="absolute">
      <n-layout-header style="padding: 16px;">

        <n-flex justify="space-between">
          <span>assocentity showcase</span>
          <n-a href="http://localhost">Github</n-a>
        </n-flex>
      </n-layout-header>

      <n-layout content-style="padding: 16px;">
        <n-flex align="center" justify="center">
          <n-button>International sentiment of brands</n-button>
          <n-button>National sentiment of political speeches</n-button>
          <n-button>Root verbs of music genres</n-button>
        </n-flex>

        <n-layout content-style="padding: 16px;">
          <router-view />
        </n-layout>
      </n-layout>

      <n-layout-footer position="absolute" style="padding: 16px">
        Copyright 2025 Julian Claus
      </n-layout-footer>
    </n-layout>
  </n-config-provider>
</template>

<script setup>
import { darkTheme } from 'naive-ui';
</script>

<style>
.n-layout-header,
.n-layout-footer {
  background-color: transparent;
}

.n-layout-header a:hover {
  background-color: transparent;
}
</style>