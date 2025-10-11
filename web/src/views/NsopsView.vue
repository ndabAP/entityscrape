<template>
  <n-h2>National sentiment of political speeches</n-h2>
  <n-p>
    For each mentioning of a country during a political speech in a governmental
    environment, the countries dependent token has been collected, while a token
    must be assigned to either of these part of speeches: Adjective, Noun, Verb.
    After that, the top ten most common permutations have been aggregated.
  </n-p>

  <n-page-header style="margin-bottom: 12px">
    <n-grid cols="2 400:4 600:6">
      <n-gi>
        <n-statistic
          label="Corpus"
          value="9"
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
          value="26,000"
        />
      </n-gi>
      <n-gi>
        <n-statistic
          label="Range"
          value="2002-2021"
        />
      </n-gi>
    </n-grid>
  </n-page-header>

  <n-tabs
    placement="top"
    default-value="germany"
    style="margin-bottom: 12px"
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
      name="usa"
      tab="United States"
    >
      <XNsops
        label="United States"
        identifier="usa"
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

  <n-text strong>
    Sources
  </n-text>
  <n-ul>
    <n-li>
      <!-- eslint-disable-next-line -->
      <n-a href="https://data.stanford.edu/congress_text" target="_blank">Congressional Record for the 43rd-114th Congresses: Parsed Speeches and Phrase Counts</n-a>,
      Gentzkow, Matthew, Jesse M. Shapiro, and Matt Taddy. Congressional Record for the 43rd-114th Congresses: Parsed
      Speeches and Phrase Counts. Palo Alto, CA: Stanford Libraries [distributor], 2018-01-16.
    </n-li>
    <n-li>
      <!-- eslint-disable-next-line -->
      <n-a href="https://politische-reden.eu" target="_blank">German Political Speeches Corpus and Visualization</n-a>,
      Barbaresi, Adrien (2018). "A corpus of German political speeches from the 21st century", Proceedings of the
      Eleventh International Conference on Language Resources and Evaluation (LREC 2018), European Language Resources
      Association (ELRA), pp. 792–797.<br>
      Adrien Barbaresi. (2019). German Political Speeches Corpus (Version v4.2019) [Data set]. Zenodo
    </n-li>
    <n-li>
      <!-- eslint-disable-next-line -->
      <n-a href="https://discuss-data.net/dataset/fb52dac2-66e3-47a3-86c5-b2a3dadf41bf/files/" target="_blank">Duma Speeches: A Term Frequency Analysis</n-a>,
      Daniel Marcus (Oct. 4, 2021). Made available under the Open Data Commons Attribution License (ODC-BY) v1.0.
    </n-li>
  </n-ul>
</template>

<script setup>
import XNsops from '@/components/XNsops.vue'
</script>
