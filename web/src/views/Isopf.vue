<template>
  <n-h2>International sentiment of public figures</n-h2>
  <n-p>
    The direct dependency ancestor or descended token of a public figure
    in news articles have been collected, while a token must be assigned to
    either of these part of speeches: Adjective, Noun, Verb. After that, the
    top ten most common permutations have been aggregated.
  </n-p>

  <n-tabs
    :justify-content="justifyContentTabs"
    placement="top"
    type="line"
    default-value="biden"
  >
    <n-tab-pane
      name="biden"
      tab="Joe Biden"
    >
      <n-p depth="3">
        Joseph Robinette Biden Jr. is an American politician who served as the
        46th president of the United States from 2021 to 2025.
      </n-p>
      <div
        v-if="joeBiden"
        style="display: block; margin-left: auto; margin-right: auto;"
      >
        <XChart :code="joeBiden" />
      </div>
    </n-tab-pane>
    <n-tab-pane
      name="trump"
      tab="Donald Trump"
    >
      <n-p depth="3">
        Donald John Trump is an American politician, media personality, and
        businessman who is the 47th president of the United States. A member of
        the Republican Party, he served as the 45th president from 2017 to 2021.
      </n-p>
      <div
        v-if="donaldTrump"
        style="display: block; margin-left: auto; margin-right: auto;"
      >
        <XChart :code="donaldTrump" />
      </div>
    </n-tab-pane>
    <n-tab-pane
      name="musk"
      tab="Elon Musk"
    >
      <n-p depth="3">
        Elon Reeve Musk is an international businessman and entrepreneur known
        for his leadership of Tesla, SpaceX, X, and the Department of Government
        Efficiency. Musk has been the wealthiest person in the world since 2021.
      </n-p>
      <div
        v-if="elonMusk"
        style="display: block; margin-left: auto; margin-right: auto;"
      >
        <XChart :code="elonMusk" />
      </div>
    </n-tab-pane>
    <n-tab-pane
      name="putin"
      tab="Vladimir Putin"
    >
      <n-p depth="3">
        Vladimir Vladimirovich Putin is a Russian politician and former
        intelligence officer who has served as President of Russia since 2012,
        having previously served from 2000 to 2008.
      </n-p>
      <div
        v-if="vladimirPutin"
        style="display: block; margin-left: auto; margin-right: auto;"
      >
        <XChart :code="vladimirPutin" />
      </div>
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import XChart from '@/components/XChart.vue'
import { onMounted, ref } from 'vue'

const joeBiden = ref('')
const donaldTrump = ref('')
const elonMusk = ref('')
const vladimirPutin = ref('')

const justifyContentTabs = ref(undefined)

onMounted(async () => {
  if (document.body.clientWidth < 768) {
    justifyContentTabs.value = undefined
  } else {
    justifyContentTabs.value = 'space-evenly'
  }
  window.addEventListener('resize', async () => {
    if (document.body.clientWidth < 768) {
      justifyContentTabs.value = undefined
    } else {
      justifyContentTabs.value = 'space-evenly'
    }
  })

  {
    const response = await fetch('/isopf/biden.json')
    const file = await response.json()
    joeBiden.value = toChart(file, 'Joe Biden')
  }
  {
    const response = await fetch('/isopf/trump.json')
    const file = await response.json()
    donaldTrump.value = toChart(file, 'Donald Trump')
  }
  {
    const response = await fetch('/isopf/musk.json')
    const file = await response.json()
    elonMusk.value = toChart(file, 'Elon Musk')
  }
  {
    const response = await fetch('/isopf/putin.json')
    const file = await response.json()
    vladimirPutin.value = toChart(file, 'Vladimir Putin')
  }
})

const code = `---
config:
  sankey:
---
sankey-beta

%% source,target,value
`

const toChart = (report, publicFigure) => {
  let chart = code
  for (const row of report.ancestors) {
    chart += `${row.heads[0]} (a),${publicFigure},${row.n}\n`
  }
  for (const row of report.descendants) {
    chart += `${publicFigure},${row.heads[0]} (d),${row.n}\n`
  }
  return chart
}
</script>
