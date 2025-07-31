<template>
  <n-config-provider :theme="darkTheme">
    <n-layout position="absolute">
      <n-menu
        mode="horizontal"
        :options="menuOptions"
        responsive
      />

      <div style="padding: 0 20px;">
      <RouterView  />
    </div>
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
]
</script>

<style>
.n-menu .n-menu-item-content .n-menu-item-content-header {
  white-space: break-spaces;
}
</style>