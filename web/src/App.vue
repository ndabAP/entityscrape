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

      <b-row>
        <b-col>
          <h2>Donald Trump</h2>
          <donald-trump />
        </b-col>

        <b-col>
          <h2>Angela Merkel</h2>
          <angela-merkel />
        </b-col>
      </b-row>

      <b-row>
        <b-col>
          <h2>Xi Jinping</h2>
          <xi-jinping />
        </b-col>

        <b-col>
          <h2>Elon Musk</h2>
          <elon-musk />
        </b-col>
      </b-row>

    <a class="mt-4" href="https://github.com/ndabAP/entityscrape">Source code</a>
    </b-container>
  </div>
</template>

<script>
import AngelaMerkel from './components/AngelaMerkel'
import DonaldTrump from './components/DonaldTrump'
import XiJinping from './components/XiJinping'
import ElonMusk from './components/ElonMusk'

export default {
  components: {
    AngelaMerkel,
    DonaldTrump,
    XiJinping,
    ElonMusk
  }
}
</script>
