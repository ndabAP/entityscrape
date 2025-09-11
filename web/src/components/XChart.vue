<template>
  <pre
    ref="element"
    class="mermaid"
    v-html="code"
  />
</template>

<script setup>
import mermaid from 'mermaid'
import { onMounted, ref } from 'vue'

defineProps({
  code: String
})

const toggleSankey = () => {
  if (document.body.clientWidth < 768) {
    config.sankey.height = 1200
  } else if (document.body.clientWidth < 470) {
    config.sankey.height = 1600
  } else {
    config.sankey.height = 850
  }
  mermaid.initialize(config)
}

const fontFamily = 'v-sans, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol"'
const config = {
  theme: 'dark',
  darkMode: true,
  useMaxWidth: true,
  themeVariables: {
    fontFamily,
    fontSize: '24px'
  },
  startOnLoad: false,
  securityLevel: 'loose',
  sankey: {
    width: 1200,
    height: 800
  }
}

const element = ref(null)

onMounted(async () => {
  toggleSankey()
  const html = element.value.innerHTML
  window.addEventListener('resize', async () => {
    toggleSankey()
    mermaid.initialize(config)

    element.value.removeAttribute('data-processed')
    element.value.innerHTML = html
    await mermaid.run({
      nodes: [element.value]
    })
  })
  await mermaid.run()
})
</script>

<style>
.mermaid { display: flex !important; justify-content: center }
</style>
