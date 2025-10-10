<template>
  <n-h2>International sentiment of public figures</n-h2>
  <n-p>
    For each subject, the heads and dependents in international news have
    been collected, while the token must be assigned to either of these
    parts of speeches: Adjective, Noun, Verb. After that, the top ten most
    common occurrences have been aggregated.
  </n-p>

  <n-text strong>
    Sources
  </n-text>
  <n-ul>
    <n-li>
      <!-- eslint-disable-next-line -->
      <n-a href="https://github.com/Webhose/free-news-datasets" target="_blank">Webz.io News Dataset Repository</n-a>, Webz.io LTD B.S.R Tower 4, 7 Mesada st.
      Bnei Brak, Tel Aviv District, POB 105 ISRAEL 5126112 (January 1st, 2024)
    </n-li>
  </n-ul>

  <n-tabs
    placement="top"
    default-value="musk"
  >
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
      name="biden"
      tab="Joe Biden"
    >
      <XIsopf
        label="Joe Biden"
        text="Joseph Robinette Biden Jr. is an American politician who served as the 46th president of the United States from 2021 to 2025."
        identifier="biden"
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
