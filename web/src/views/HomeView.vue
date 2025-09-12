<template>
  <n-grid
    :x-gap="16"
    :y-gap="16"
    cols="1 s:1 m:2 l:3"
    responsive="screen"
  >
    <n-grid-item>
      <n-card
        :segmented="{
          content: false
        }"
        size="large"
        :bordered="false"
        title="International sentiment of public figures"
      >
        <template #header-extra>
          isopf
        </template>
        <n-p>
          The direct dependency ancestor or descended token of a public figure
          in news articles have been collected, while a token must be assigned to
          either of these part of speeches: Adjective, Noun, Verb. After that, the
          top ten most common permutations have been aggregated.
        </n-p>
        <template #footer>
          <n-p depth="3">
            Donald Trump, Elon Musk, Joe Biden, Vladimir Putin
          </n-p>
          <n-flex
            justify="space-between"
          >
            <n-statistic label="Corpus">
              33,000
            </n-statistic>
            <n-statistic label="Figures">
              4
            </n-statistic>
            <n-statistic label="Samples">
              >10,000
            </n-statistic>
            <n-statistic label="Sample">
              100 %
            </n-statistic>
          </n-flex>
        </template>
        <template #action>
          <n-button
            text
            block
            round
            size="large"
            @click="() => router.push('isopf')"
          >
            View case study
          </n-button>
        </template>
      </n-card>
    </n-grid-item>

    <n-grid-item>
      <n-card
        :segmented="{
          content: false
        }"
        size="large"
        :bordered="false"
        title="National sentiment of political speeches"
      >
        <template #header-extra>
          nsops
        </template>
        <n-p>
          For each mentioning of a country during a political speech in a
          governmental environment, the countries dependent token has been
          collected, while a token must be assigned to either of these part of
          speeches: Adjective, Noun, Verb. After that, the top ten most common
          permutations have been aggregated.
        </n-p>
        <template #footer>
          <n-p depth="3">
            Germany, Russia, United States
          </n-p>
          <n-flex
            justify="space-between"
          >
            <n-statistic label="Corpus">
              9
            </n-statistic>
            <n-statistic label="Countries">
              3
            </n-statistic>
            <n-statistic label="Samples">
              >10,000
            </n-statistic>
            <n-statistic label="Sample">
              90 %
            </n-statistic>
          </n-flex>
        </template>
        <template #action>
          <n-button
            block
            round
            text
            size="large"
            @click="router.push('nsops')"
          >
            View case study
          </n-button>
        </template>
      </n-card>
    </n-grid-item>

    <n-grid-item>
      <n-card
        :segmented="{
          content: false
        }"
        size="large"
        :bordered="false"
        title="Root verbs of music genres"
      >
        <template #header-extra>
          rvomg
        </template>
        <n-p>
          For each of the three music genres, twenty song lyrics have been used to
          collect the dependency root verb for every song. After that,
          the top ten most common occurences have been aggregated.
        </n-p>
        <template #footer>
          <n-p depth="3">
            Hip-hop, Rock and Roll, Pop
          </n-p>
          <n-flex
            justify="space-between"
          >
            <n-statistic label="Corpus">
              30
            </n-statistic>
            <n-statistic label="Genres">
              3
            </n-statistic>
            <n-statistic label="Samples">
              >100
            </n-statistic>
            <n-statistic label="Sample">
              100 %
            </n-statistic>
          </n-flex>
        </template>
        <template #action>
          <n-button
            block
            round
            text
            size="large"
            @click="router.push('rvomg')"
          >
            View case study
          </n-button>
        </template>
      </n-card>
    </n-grid-item>
  </n-grid>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<style scoped>
.n-grid > div {
  display: flex;
}

.n-card {
  height: 100%;
}

p {
  text-align: justify;
}
</style>

<style>
.n-statistic .n-statistic-value .n-statistic-value__content {
  font-size: unset !important;
}
</style>
