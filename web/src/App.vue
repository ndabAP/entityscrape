<template>
  <n-config-provider :theme="darkTheme">
    <n-layout position="absolute">
      <n-layout>
        <n-menu :mode="menuMode" responsive :options="menuOptions" />
      </n-layout>

      <n-layout content-style="padding: 16px; max-width: 1200px">
        <router-view />
      </n-layout>
    </n-layout>
  </n-config-provider>
</template>

<script setup>
import { darkTheme } from 'naive-ui';
import { h, onMounted, ref } from 'vue';
import { RouterLink } from 'vue-router';

const menuMode = ref('horizontal');
const toggleMenuMode = () => {
  if (document.body.clientWidth < 768) {
    menuMode.value = 'vertical';
  } else {
    menuMode.value = 'horizontal';
  }
}
onMounted(() => {
  toggleMenuMode()
  window.addEventListener("resize", toggleMenuMode)
})

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
        { default: () => 'Home' }
      ),
    key: 'home',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'isob'
          }
        },
        { default: () => 'International sentiment of brands' }
      ),
    key: 'isob',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'nsops'
          }
        },
        { default: () => 'National sentiment of political speeches' }
      ),
    key: 'nsops',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'rvomg'
          }
        },
        { default: () => 'Root verbs of music genres' }
      ),
    key: 'rvomg',
  },
]
</script>
