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
    v-if="option.series.data.length > 0"
    style="display: block; margin-left: auto; margin-right: auto;"
  >
    <XChart :option="option" />
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
const loading = ref(true)
const error = ref(false)

onMounted(async () => {
  try {
    const report = await fetch(`/entityscrape/isopf/${props.identifier}.json`).then(response => response.json())

    option.series.data.push({ name: props.label })
    for (const heads of report.heads) {
      const word = `${heads.word[0]} (h)`
      option.series.data.push({ name: word })
      option.series.links.push({
        source: word,
        target: props.label,
        value: heads.n
      })
    }
    for (const dependent of report.dependents) {
      const word = `${dependent.word[0]} (d)`
      option.series.data.push({ name: word })
      option.series.links.push({
        source: props.label,
        target: word,
        value: dependent.n
      })
    }
  } catch {
    error.value = true
  } finally {
    loading.value = false
  }
})
</script>
