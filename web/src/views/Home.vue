<template>
  <n-flex
    size="large"
    :wrap="wrap"
  >
    <n-card
      :bordered="true"
      :segmented="{
        content: false,
        footer: false,
      }"
      size="large"
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
      <n-p>
        Test subjects are: Donald Trump, Elon Musk, Joe Biden and Vladimir Putin
      </n-p>
      <template #footer>
        <n-flex
          justify="space-between"
          style="margin-bottom: 24px"
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
          block
          round
          size="large"
          @click="() => push('isopf')"
        >
          View case study
        </n-button>
      </template>
    </n-card>

    <n-card
      :bordered="true"
      :segmented="{
        content: false,
        footer: false,
      }"
      size="large"
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
      <n-p>
        Test subjects are: Germany, Russia and United States
      </n-p>
      <template #footer>
        <n-flex
          justify="space-between"
          style="margin-bottom: 24px"
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
          size="large"
          @click="push('nsops')"
        >
          View case study
        </n-button>
      </template>
    </n-card>

    <n-card
      :bordered="true"
      :segmented="{
        content: false,
        footer: false,
      }"
      size="large"
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
      <n-p>
        Test subjects are: Hip-hop, Rock and Roll, Pop
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
          size="large"
          @click="push('rvomg')"
        >
          View case study
        </n-button>
      </template>
    </n-card>
  </n-flex>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const push = name => router.push({ name })

const wrap = ref(false)
onMounted(() => {
  wrap.value = window.innerWidth < 1024
  window.addEventListener('resize', () => {
    wrap.value = window.innerWidth < 1024
  })
})
</script>

<style>
.n-statistic .n-statistic-value .n-statistic-value__content {
    font-size: unset;
}
</style>
