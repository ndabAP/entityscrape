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
        <n-p>
          For each subject, the heads and dependents in international news have
          been collected, while the token must be assigned to either of these
          parts of speeches: Adjective, Noun, Verb. After that, the top ten most
          common occurrences have been aggregated.
        </n-p>
        <template #footer>
          <n-text strong>
            Subjects
          </n-text>
          <n-p depth="2">
            Donald Trump, Elon Musk, Joe Biden, Vladimir Putin
          </n-p>
          <n-flex
            justify="space-between"
          >
            <n-statistic label="Corpus">
              39,000
            </n-statistic>
            <n-statistic label="Subjects">
              4
            </n-statistic>
            <n-statistic label="Population">
              9,300
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
        <n-p>
          For each mentioning of a country during a political speech in a
          governmental environment, the countries dependent token has been
          collected, while a token must be assigned to either of these part of
          speeches: Adjective, Noun, Verb. After that, the top ten most common
          occurrences have been aggregated.
        </n-p>
        <template #footer>
          <n-text strong>
            Subjects
          </n-text>
          <n-p depth="2">
            Germany, Russia, United States
          </n-p>
          <n-flex
            justify="space-between"
          >
            <n-statistic label="Corpus">
              9
            </n-statistic>
            <n-statistic label="Subjects">
              3
            </n-statistic>
            <n-statistic label="Population">
              26,087
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
        <n-p>
          For each of the three music genres, twenty song lyrics have been used
          to collect the dependency root verb for every verse. After that, the
          top ten most common occurrences have been aggregated.
        </n-p>
        <template #footer>
          <n-text strong>
            Subjects
          </n-text>
          <n-p depth="2">
            Hip-hop, Rock and Roll, Pop
          </n-p>
          <n-flex
            justify="space-between"
          >
            <n-statistic label="Corpus">
              30
            </n-statistic>
            <n-statistic label="Subjects">
              3
            </n-statistic>
            <n-statistic label="Population">
              >100
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
</style>
