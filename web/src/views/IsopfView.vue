<template>
  <n-h2>International sentiment of public figures</n-h2>
  <n-p>
    The direct dependency ancestor or descended token of a public figure in news
    articles have been collected, while a token must be assigned to either of
    these part of speeches: Adjective, Noun, Verb. After that, the top ten most
    common permutations have been aggregated.
  </n-p>

  <n-tabs
    placement="top"
    type="bar"
    size="large"
    default-value="biden"
  >
    <n-tab-pane
      name="biden"
      tab="Joe Biden"
    >
      <XIsopf
        label="Joe Biden"
        text=" Joseph Robinette Biden Jr. is an American politician who served as the 46th president of the United States from 2021 to 2025."
        identifier="biden"
      />
    </n-tab-pane>
    <n-tab-pane
      name="trump"
      tab="Donald Trump"
    >
      <XIsopf
        label="Donald Trump"
        text="Donald John Trump is an American politician, media personality, and businessman who is the 47th president of the United States. A member of the Republican Party, he served as the 45th president from 2017 to 2021."
        identifier="trump"
      />
    </n-tab-pane>
    <n-tab-pane
      name="musk"
      tab="Elon Musk"
    >
      <XIsopf
        label="Elon Musk"
        text="Elon Reeve Musk is an international businessman and entrepreneur known for his leadership of Tesla, SpaceX, X, and the Department of Government Efficiency. Musk has been the wealthiest person in the world since 2021."
        identifier="musk"
      />
    </n-tab-pane>
    <n-tab-pane
      name="putin"
      tab="Vladimir Putin"
    >
      <XIsopf
        label="Vladimir Putin"
        text="Vladimir Vladimirovich Putin is a Russian politician and former intelligence officer who has served as President of Russia since 2012, having previously served from 2000 to 2008."
        identifier="putin"
      />
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import XIsopf from '@/components/XIsopf.vue'
</script>
