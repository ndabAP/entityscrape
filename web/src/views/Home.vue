<template>
  <n-flex
    size="large"
    :wrap="wrap"
  >
    <n-card
      :bordered="false"
      title="International sentiment of brands"
    >
      <n-p>
        At max three direct dependency ancestor token of each brand have
        been collected. After that, the top ten most common permutations
        have been aggregated.
      </n-p>
      <n-p>
        Possesive noun suffixes, non-ASCII characters and multi-token
        brands have been filtered out.
      </n-p>
      <template #footer>
        <n-flex
          justify="space-between"
          style="margin-bottom: 24px"
        >
          <n-statistic label="Corpus">
            8,000
          </n-statistic>
          <n-statistic label="Brands">
            3
          </n-statistic>
          <n-statistic label="Samples">
            4,000
          </n-statistic>
        </n-flex>
        <n-button block>
          View case study
        </n-button>
      </template>
    </n-card>

    <n-card
      :bordered="false"
      title="National sentiment of political speeches"
    >
      <n-p>
        The direct dependency ancestor or descended token of a country
        have been collected. After that, the top ten most common
        permutations have been aggregated.
      </n-p>
      <n-p>
        As for ancestors and descendends, only verbs, nouns and
        adjectives are included.
      </n-p>
      <template #footer>
        <n-flex
          justify="space-between"
          style="margin-bottom: 24px"
        >
          <n-statistic label="Corpus">
            8,000
          </n-statistic>
          <n-statistic label="Countries">
            3
          </n-statistic>
          <n-statistic label="Samples">
            4,000
          </n-statistic>
        </n-flex>
        <n-button block>
          View case study
        </n-button>
      </template>
    </n-card>

    <n-card
      :bordered="false"
      title="Root verbs of music genres"
    >
      <n-p>
        For each of the three music genres, there are ten songs used to
        collect the dependency root verb for every song. After that,
        the top ten most common occurences have been aggregated.
      </n-p>
      <n-p>
        A verse has been considered as a new sentence and a terminal
        punctuation has been added. Due to its nature of song lyrics,
        some dependency trees might have been falsely parsed.
      </n-p>
      <template #footer>
        <n-flex
          justify="space-between"
          style="margin-bottom: 24px"
        >
          <n-statistic label="Corpus">
            30
          </n-statistic>
          <n-statistic label="Genres">
            3
          </n-statistic>
          <n-statistic label="Samples">
            4,000
          </n-statistic>
        </n-flex>
        <n-button block>
          View case study
        </n-button>
      </template>
    </n-card>
  </n-flex>
</template>

<script setup>
import { onMounted, ref } from 'vue'

const wrap = ref(false)

onMounted(() => {
  wrap.value = window.innerWidth < 768
  window.addEventListener('resize', () => {
    wrap.value = window.innerWidth < 768
  })
})
</script>

<style>
.n-statistic .n-statistic-value .n-statistic-value__content {
    font-size: unset;
}
</style>
