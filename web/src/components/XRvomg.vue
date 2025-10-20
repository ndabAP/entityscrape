<template>
  <n-p depth="3">
    {{ text }}
  </n-p>
  <n-skeleton
    v-if="loading"
    round
    text
    :repeat="2"
  />
  <n-skeleton
    v-if="loading"
    round
    text
    style="width: 60%"
  />
  <div
    v-if="option.series?.[0].data.length > 0"
    style="display: block; margin-left: auto; margin-right: auto;"
  >
    <XChart
      :option="option"
    />
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
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

const loading = ref(true)
const error = ref(false)

onMounted(async () => {
  try {
    const report = await fetch(`/entityscrape/rvomg/${props.identifier}.json`).then(response => response.json())
    for (const { word, n } of report) {
      option.xAxis.data.push(word[0])
      option.series[0].data.push(n)
    }
  } catch {
    error.value = true
  } finally {
    loading.value = false
  }
})
</script>
