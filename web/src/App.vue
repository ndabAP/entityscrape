<template>
  <div id="app">
    <veui-grid-container>
      <h1>Entity scrape</h1>
        <veui-grid-row>
          <veui-grid-column :span="12">
            <h2>Donald Trump</h2>
            <v-chart :options="trump"/>
          </veui-grid-column>
          <veui-grid-column :span="12">
            <h2>Angela Merkel</h2>
            <v-chart :options="merkel"/>
          </veui-grid-column>
        </veui-grid-row>
      </veui-grid-container>
  </div>
</template>

<script>
import { GridContainer, GridRow, GridColumn } from 'veui'
import axios from 'axios'

export default {
  components: {
    'veui-grid-container': GridContainer,
    'veui-grid-row': GridRow,
    'veui-grid-column': GridColumn
  },

  async mounted () {
    let { data: trump } = await axios.get('/api/entities?entity=Donald%20Trump')

    this.trump.visualMap.min = Math.min(...trump.map(({ count }) => count))
    this.trump.visualMap.max = Math.max(...trump.map(({ count }) => count))
    trump = trump.map(({ count, word, distance }) => ([count, distance, word])).reverse()

    this.trump.dataset.source = [...this.trump.dataset.source, ...trump]

    let { data: merkel } = await axios.get('/api/entities?entity=Angela%20Merkel')

    this.merkel.visualMap.min = Math.min(...merkel.map(({ count }) => count))
    this.merkel.visualMap.max = Math.max(...merkel.map(({ count }) => count))
    merkel = merkel.map(({ count, word, distance }) => ([count, distance, word])).reverse()

    this.merkel.dataset.source = [...this.merkel.dataset.source, ...merkel]
  },

  data: () => ({
    trump: {
      dataset: {
        source: [
          ['count', 'distance', 'word']
        ]
      },

      xAxis: { name: 'Distance' },
      yAxis: { type: 'category' },

      visualMap: {
        orient: 'horizontal',
        left: 'center',
        dimension: 0,
        min: 10,
        max: 100,
        text: ['High count', 'Low count'],
        inRange: {
          color: ['#D7DA8B', '#E15457']
        }
      },

      series: [{
        type: 'bar',
        encode: {
          x: 1,
          y: 2
        }
      }]
    },

    merkel: {
      dataset: {
        source: [
          ['count', 'distance', 'word']
        ]
      },

      xAxis: { name: 'Distance' },
      yAxis: { type: 'category' },

      visualMap: {
        orient: 'horizontal',
        left: 'center',
        dimension: 0,
        min: 10,
        max: 100,
        text: ['High count', 'Low count'],
        inRange: {
          color: ['#D7DA8B', '#E15457']
        }
      },

      series: [{
        type: 'bar',
        encode: {
          x: 1,
          y: 2
        }
      }]
    }
  })
}
</script>

<style lang="less">
@import "~veui-theme-one/common.less";
</style>
