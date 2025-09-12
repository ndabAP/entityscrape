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
      type="sankey"
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
    type: 'sankey',
    data: [],
    links: [],
    label: {
      color: 'white',
      textBorderColor: 'transparent'
    }
  }
})
onMounted(async () => {
  const response = await fetch(`/isopf/${props.identifier}.json`)
  const file = await response.json()
  option.series.data.push({ name: props.label })
  for (const row of file.ancestors) {
    const value = `${row.heads[0]} (a)`
    option.series.data.push({ name: value })
    option.series.links.push({
      source: value,
      target: props.label,
      value: row.n
    })
  }
  for (const row of file.descendants) {
    const value = `${row.heads[0]} (d)`
    option.series.data.push({ name: value })
    option.series.links.push({
      source: props.label,
      target: value,
      value: row.n
    })
  }
})
</script>
