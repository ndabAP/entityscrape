<template>
  <n-h2>Root verbs of music genres</n-h2>
  <n-p>
    For each of the three music genres, twenty song lyrics have been used to
    collect the dependency root verb for every song. After that,
    the top ten most common occurences have been aggregated.
  </n-p>

  <n-tabs
    :justify-content="justifyContentTabs"
    placement="top"
    type="line"
    default-value="hip_hop"
  >
    <n-tab-pane
      name="hip_hop"
      tab="Hip-hop"
    >
      <n-p depth="3">
        Hip-hop or hip hop is a popular music genre that emerged in the early
        1970s alongside a hip-hop subculture built by the African-American and
        Latino communities of New York City.
      </n-p>
      <n-table
        v-if="hipHop.length > 0"
        :bordered="false"
        :single-line="false"
      >
        <thead>
          <tr>
            <th>Word</th>
            <th>Occurence</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="column of hipHop"
            :key="column"
          >
            <td
              v-for="row of column"
              :key="row"
            >
              {{ row }}
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-tab-pane>
    <n-tab-pane
      name="rock_and_roll"
      tab="Rock and Roll"
    >
      <n-p depth="3">
        Rock and roll is a genre of popular music that evolved in the United
        States during the late 1940s and early 1950s.
      </n-p>
      <n-table
        v-if="hipHop.length > 0"
        :bordered="false"
        :single-line="false"
      >
        <thead>
          <tr>
            <th>Word</th>
            <th>Occurence</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="column of rock"
            :key="column"
          >
            <td
              v-for="row of column"
              :key="row"
            >
              {{ row }}
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-tab-pane>
    <n-tab-pane
      name="pop"
      tab="Pop"
    >
      <n-p depth="3">
        Pop music, or simply pop, is a genre of popular music that originated in
        its modern form during the mid-1950s in the United States and the
        United Kingdom.
      </n-p>
      <n-table
        v-if="hipHop.length > 0"
        :bordered="false"
        :single-line="false"
      >
        <thead>
          <tr>
            <th>Word</th>
            <th>Occurence</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="column of pop"
            :key="column"
          >
            <td
              v-for="row of column"
              :key="row"
            >
              {{ row }}
            </td>
          </tr>
        </tbody>
      </n-table>
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import { onMounted, ref } from 'vue'

const hipHop = ref([])
const rock = ref([])
const pop = ref([])

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
    const response = await fetch('/rvomg/rap.json')
    const file = await response.json()
    hipHop.value = toTable(file)
  }
  {
    const response = await fetch('/rvomg/rock.json')
    const file = await response.json()
    rock.value = toTable(file)
  }
  {
    const response = await fetch('/rvomg/pop.json')
    const file = await response.json()
    pop.value = toTable(file)
  }
})

const toTable = report => {
  const rows = []
  for (const row of report) {
    rows.push([row.word[0], row.n])
  }
  return rows
}
</script>
