<template>
  <v-chart
    class="chart"
    :option="props.option"
    autoresize
  />
</template>

<script setup>
import { SankeyChart } from 'echarts/charts'
import { use } from 'echarts/core'
import { SVGRenderer } from 'echarts/renderers'
import VChart from 'vue-echarts'

use([
  SVGRenderer,
  SankeyChart
])

const props = defineProps({
  option: Object,
  type: String
})
</script>

<style scoped>
.chart {
  height: 550px;
  width: 100%;
}
</style>
