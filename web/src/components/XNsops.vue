<template>
  <n-p depth="3">
    {{ text }}
  </n-p>
  <div
    v-if="option.series.data.length > 0"
    style="display: block; margin-left: auto; margin-right: auto;"
  >
    <XChart
      :option="option"
    />
  </div>
</template>

<script setup>
import { onMounted, reactive } from 'vue'
import XChart from './XChart.vue'

const props = defineProps({
  label: String,
  identifier: String,
  text: String
})

const option = reactive({
  series: {
    type: 'sunburst',
    data: [],
    nodeClick: false,
    radius: ['20%', '100%']
  }
})
onMounted(async () => {
  const report = await fetch(`/nsops/${props.identifier}.json`).then(response => response.json())

  for (const [pos, words] of Object.entries(Object.groupBy(report, ({ pos }) => pos))) {
    option.series.data.push({
      name: pos,
      children: words.map(({ word, n }) => ({ name: `${word[0]}\n(${word[1]})`, value: n }))
    })
  }
})
</script>
