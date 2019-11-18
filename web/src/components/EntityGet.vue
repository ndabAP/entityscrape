<template>
    <div>
      <p v-if="news > 0">Based on {{ news }} news and {{ associations }} associations.</p>

      <b-form-select
        class="mb-4"
        v-model="partOfSpeech"
        :options="partOfSpeeches">
      </b-form-select>

      <div :id="`${dash}-chart`" />
    </div>
</template>

<script>
import axios from 'axios'
import Plotly from 'plotly.js/lib/core'
import kebabCase from 'lodash/kebabCase'

Plotly.register([
  require('plotly.js/lib/bar')
])

const DEFAULT_CHART = {
  type: 'bar',
  x: [],
  y: [],
  text: [],
  textposition: 'auto',
  orientation: 'h',
  marker: {
    color: []
  }
}

export default {
  props: {
    entity: {
      type: String,
      required: true
    }
  },

  async mounted () {
    await this.getEntity()
    Plotly.newPlot(`${this.dash}-chart`, [this.chart], this.layout, this.options)

    await this.getNews()
    await this.getAssociations()
  },

  data: () => ({
    partOfSpeech: 'adj',
    partOfSpeeches: [
      { value: 'adj', text: 'Adjective' },
      { value: 'adp', text: 'Adposition' },
      { value: 'adv', text: 'Adverb' },
      { value: 'conj', text: 'Conjunction' },
      { value: 'det', text: 'Determiner' },
      { value: 'noun', text: 'Noun' },
      { value: 'num', text: 'Cardinal number' },
      { value: 'pron', text: 'Pronoun' },
      { value: 'prt', text: 'Particle' },
      { value: 'punct', text: 'Punctuation' },
      { value: 'verb', text: 'Verb' },
      { value: 'x', text: 'Other' }
    ],

    news: 0,
    associations: 0,
    chart: DEFAULT_CHART,

    layout: {
      autosize: true,

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
  }),

  methods: {
    async getEntity () {
      return new Promise(async resolve => {
        let { data: chart } = await axios.get(`/api/entities?entity=${this.uri}&part-of-speech=${this.partOfSpeech}`)
        this.chart.x = []
        this.chart.y = []
        this.chart.text = []
        this.chart.marker.color = []

        chart.map(({ count, distance, word }) => {
          this.chart.x.push(Math.round(distance))
          this.chart.y.push(word)
          this.chart.text.push(Math.round(distance))
          this.chart.marker.color.push(count)
        })

        resolve()
      })
    },

    async getAssociations () {
      return new Promise(async resolve => {
        const { data: associations } = await axios.get(`/api/associations?entity=${this.uri}`)
        this.associations = associations

        resolve()
      })
    },

    async getNews () {
      return new Promise(async resolve => {
        let { data: count } = await axios.get(`/api/news?entity=${this.uri}`)
        this.news = count

        resolve()
      })
    }
  },

  computed: {
    dash: {
      get () {
        return kebabCase(this.entity)
      }
    },

    uri: {
      get () {
        return encodeURIComponent(this.entity)
      }
    }
  },

  watch: {
    async partOfSpeech () {
      await this.getEntity()
      Plotly.update(`${this.dash}-chart`, {}, this.layout, this.options)
    },

    async entity () {
      await this.getEntity()
      Plotly.newPlot(`${this.dash}-chart`, [this.chart], this.layout, this.options)

      await this.getNews()
      await this.getAssociations()
    }
  }
}
</script>
