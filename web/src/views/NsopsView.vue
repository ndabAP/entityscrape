<template>
  <n-h2>National sentiment of political speeches</n-h2>
  <n-p>
    For each mentioning of a country during a political speech in a governmental
    environment, the countries dependent token has been collected, while a token
    must be assigned to either of these part of speeches: Adjective, Noun, Verb.
    After that, the top ten most common permutations have been aggregated.
  </n-p>

  <n-tabs
    placement="top"
    type="bar"
    size="large"
    default-value="germany"
  >
    <n-tab-pane
      name="germany"
      tab="Germany"
    >
      <XNsops
        label="Germany"
        identifier="germany"
        text="Germany is a Western European country with a landscape of forests, rivers, mountain ranges and North Sea beaches. It has over 2 millennia of history."
      />
    </n-tab-pane>

    <n-tab-pane
      name="united_states"
      tab="United States"
    >
      <XNsops
        label="United States"
        identifier="united_states"
        text="The U.S. is a country of 50 states covering a vast swath of North America, with Alaska in the northwest and Hawaii extending the nation's presence into the Pacific Ocean."
      />
    </n-tab-pane>

    <n-tab-pane
      name="russia"
      tab="Russia"
    >
      <XNsops
        label="Russia"
        identifier="russia"
        text="Russia is a country spanning Eastern Europe and North Asia. It is the largest country in the world, and extends across eleven time zones, sharing land borders with fourteen countries."
      />
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import XNsops from '@/components/XNsops.vue'

</script>
