<template>
  <div id="app">
    <b-container>
      <h1>Entity scrape</h1>

      <p>
        The bar charts represent the distance between word types (default: adjectives) in news articles
        and persons. For  example, given the sentence: "I'm Donald Trump, the American president.", the
        distance between the adjective "American" and "Donald Trump" would equal three. Because "the" and ","
        are in between. The cumulated word occurrences are visualized through a color scale. Bars are
        ordered by occurrences as well.
      </p>

      <b-row class="mb-4">
        <b-col>
          <h2>Donald Trump</h2>
          <entity-get entity="Donald Trump" />
        </b-col>

        <b-col>
          <h2>Angela Merkel</h2>
           <entity-get entity="Angela Merkel" />
        </b-col>
      </b-row>

      <b-row class="mb-4">
        <b-col>
          <h2>Xi Jinping</h2>
           <entity-get entity="Xi Jinping" />
        </b-col>

        <b-col>
          <h2>Elon Musk</h2>
           <entity-get entity="Elon Musk" />
        </b-col>
      </b-row>

    <a class="mt-4" href="https://github.com/ndabAP/entityscrape">Source code</a>
    </b-container>
  </div>
</template>

<script>
import EntityGet from './components/EntityGet'

export default {
  components: {
    EntityGet
  }
}
</script>
