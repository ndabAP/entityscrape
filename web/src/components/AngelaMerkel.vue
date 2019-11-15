<template>
    <div>
        <p v-if="news > 0">Based on {{ news }} news.</p>
        <div id="angela-merkel-chart" />
    </div>
</template>

<script>
import axios from 'axios'
import Plotly from 'plotly.js-dist'

export default {
  async mounted () {
    let { data: chart } = await axios.get('/api/entities?entity=Angela%20Merkel')
    chart.map(({ count, distance, word }) => {
      this.chart.x.push(Math.round(distance))
      this.chart.text.push(Math.round(distance))
      this.chart.y.push(word)
      this.chart.marker.color.push(count)
    })

    Plotly.newPlot('angela-merkel-chart', [this.chart], this.layout, this.options)

    let { data: count } = await axios.get('/api/news?entity=Angela%20Merkel')

    this.news = count
  },

  data: () => ({
    news: 0,
    chart: {
      type: 'bar',
      x: [],
      y: [],
      text: [],
      textposition: 'auto',
      orientation: 'h',
      marker: {
        color: []
      }
    },

    layout: {
      width: 500,
      height: 500,

      font: {
        family: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"',
        size: 14,
        color: '#000000'
      },

      margin: {
        b: 0,
        t: 0,
        pad: 5
      },

      xaxis: {
        fixedrange: true,
        zeroline: false,
        ticks: '',
        showgrid: false
      },

      yaxis: {
        fixedrange: true,
        zeroline: false,
        ticks: '',
        showgrid: false
      },

      showlegend: false,
      hovermode: false
    },

    options: {
      displayModeBar: false,
      responsive: true
    }
  })
}
</script>
