<template>
    <n-flex size="large" :wrap="wrap">
        <n-card :bordered="false" title="International sentiment of brands">
            <template #header-extra>
                <n-icon size="21">
                    <Business />
                </n-icon>
            </template>

            <n-p>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore
                et dolore magna aliqua.
            </n-p>
            <template #footer>
                <n-button size="small" circle block>View</n-button>
            </template>
        </n-card>

        <n-card :bordered="false" title="National sentiment of political speeches">
            <template #header-extra>
                <n-icon size="21">
                    <Earth />
                </n-icon>
            </template>

            <n-p>
                Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
                consequat.
            </n-p>
            <template #footer>
                <n-button size="small" circle block>View</n-button>
            </template>

        </n-card>

        <n-card :bordered="false" title="Root verbs of music genres">
            <template #header-extra>
                <n-icon size="21">
                    <MusicalNotes />
                </n-icon>
            </template>

            <n-p>
                Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
                pariatur.
                Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
                consequat.
            </n-p>
            <template #footer>
                <n-button size="small" circle block>View</n-button>
            </template>
        </n-card>
    </n-flex>
</template>

<script setup>
import { Business, Earth, MusicalNotes } from '@vicons/ionicons5';
import { onMounted, ref } from 'vue';

const wrap = ref(false)

onMounted(() => {
    wrap.value = window.innerWidth < 768
    window.addEventListener("resize", () => {
        wrap.value = window.innerWidth < 768
    })
})
</script>

<style></style>