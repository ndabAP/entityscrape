<template>
  <n-h2>Root verbs of music genres</n-h2>
  <n-p>
    For each of the three music genres, twenty song lyrics have been used to
    collect the dependency root verb for every song. After that,
    the top ten most common occurences have been aggregated.
  </n-p>

  <n-page-header style="margin-bottom: 12px">
    <n-grid cols="2 400:4 600:6">
      <n-gi>
        <n-statistic
          label="Corpus"
          value="60"
        />
      </n-gi>
      <n-gi>
        <n-statistic
          label="Subjects"
          value="3"
        />
      </n-gi>
      <n-gi>
        <n-statistic
          label="Population"
          value="60"
        />
      </n-gi>
      <n-gi>
        <n-statistic
          label="Range"
          value="ca. 1950-ca. 2020"
        />
      </n-gi>
    </n-grid>
  </n-page-header>

  <n-tabs
    placement="top"
    type="bar"
    size="large"
    default-value="hip_hop"
  >
    <n-tab-pane
      name="hip_hop"
      tab="Hip-hop"
    >
      <XRvomg
        label="Hip-hop"
        identifier="hip_hop"
        text="Hip-hop is a popular music genre that emerged in the early 1970s alongside a hip-hop subculture built by the African-American and Latino communities of New York City."
      />
    </n-tab-pane>
    <n-tab-pane
      name="rock_and_roll"
      tab="Rock and Roll"
    >
      <XRvomg
        label="Rock and Roll"
        identifier="rock_and_roll"
        text="Rock and roll is a genre of popular music that evolved in the United States during the late 1940s and early 1950s."
      />
    </n-tab-pane>
    <n-tab-pane
      name="pop"
      tab="Pop"
    >
      <XRvomg
        label="Pop"
        identifier="pop"
        text="Pop music is a genre of popular music that originated in its modern form during the mid-1950s in the United States and the United Kingdom."
      />
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import XRvomg from '@/components/XRvomg.vue'
</script>
