<template>
  <n-p depth="3">
    {{ text }}
  </n-p>
  <div
    v-if="option.series[0].data.length > 0"
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
  xAxis: {
    type: 'category',
    data: []
  },
  yAxis: {
    type: 'value'
  },
  colorBy: 'data',
  series: [
    {
      type: 'bar',
      data: []
    }
  ]
})
onMounted(async () => {
  const report = await fetch(`/rvomg/${props.identifier}.json`).then(response => response.json())
  console.debug(report)
  for (const { word, n } of report) {
    option.xAxis.data.push(word[0])
    option.series[0].data.push(n)
  }
})
</script>
