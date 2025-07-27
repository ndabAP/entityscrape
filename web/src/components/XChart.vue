<template>
    <pre class="mermaid" ref="element" v-html="code"></pre>
</template>

<script setup>
import mermaid from 'mermaid';
import { onMounted, ref } from 'vue';

defineProps({
    code: String,
})

const toggleSankey = () => {
    if (document.body.clientWidth < 768) {
        config.sankey.height = 800
    } else {
        config.sankey.height = 500
    }
    mermaid.initialize(config)
}

const fontFamily = 'v-sans, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol"'
const config = {
    theme: "dark",
    darkMode: true,
    useMaxWidth: true,
    themeVariables: {
        fontFamily: fontFamily,
        fontSize: "16px",
    },
    startOnLoad: false,
    securityLevel: "loose",
    sankey: {
        width: 1200,
        height: 500
    }
};

const element = ref(null)

onMounted(async () => {
    toggleSankey()
    const html = element.value.innerHTML
    window.addEventListener("resize", async () => {
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