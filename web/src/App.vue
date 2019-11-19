<template>
  <div id="app">
    <header class="bg-primary text-white p-5 mb-5">
      <b-container class="text-center">
        <p>
          The bar charts represent the distance between word types in news articles
          and persons. For  example, given the sentence: "I'm Donald Trump, the American president.", the
          distance between the adjective "American" and "Donald Trump" would equal three. Because "the" and ","
          are in between. The cumulated word occurrences are visualized through a color scale. Bars are
          ordered by occurrences as well.
        </p>
        <b-form-select v-model="person" size="lg" :options="persons"></b-form-select>
      </b-container>
    </header>

    <b-container class="mb-4">
      <div v-if="person">
        <h2>{{ person }}</h2>
        <entity-get :entity="person" />
      </div>
    </b-container>

    <footer>
      <b-container>
        <a href="https://github.com/ndabAP/entityscrape">Source code</a>
      </b-container>
    </footer>
  </div>
</template>

<script>
import axios from 'axios'
import head from 'lodash/head'

import EntityGet from './components/EntityGet'

export default {
  components: {
    EntityGet
  },

  data: () => ({
    persons: [],
    person: ''
  }),

  async mounted () {
    const { data: persons } = await axios.get('/api/list')
    this.persons = persons

    this.person = head(persons)
  },

  watch: {
    person () {

    }
  }
}
</script>
