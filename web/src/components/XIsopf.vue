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
  const report = await fetch(`/isopf/${props.identifier}.json`).then(response => response.json())

  option.series.data.push({ name: props.label })
  for (const ancestor of report.ancestors) {
    const value = `${ancestor.heads[0]} (a)`
    option.series.data.push({ name: value })
    option.series.links.push({
      source: value,
      target: props.label,
      value: ancestor.n
    })
  }
  for (const descendant of report.descendants) {
    const value = `${descendant.heads[0]} (d)`
    option.series.data.push({ name: value })
    option.series.links.push({
      source: props.label,
      target: value,
      value: descendant.n
    })
  }
})
</script>
