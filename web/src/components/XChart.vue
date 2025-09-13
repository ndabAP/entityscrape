<template>
  <v-chart
    class="chart"
    :option="props.option"
    autoresize
  />
</template>

<script setup>
import { BarChart, SankeyChart, SunburstChart } from 'echarts/charts'
import { GridComponent } from 'echarts/components'
import { use } from 'echarts/core'
import { SVGRenderer } from 'echarts/renderers'
import VChart from 'vue-echarts'

use([
  SVGRenderer,
  SankeyChart,
  SunburstChart,
  BarChart,
  GridComponent
])

const props = defineProps({
  option: Object
})
</script>

<style scoped>
.chart {
  height: 600px;
  width: 100%;
}
</style>
