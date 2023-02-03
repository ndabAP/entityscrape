<template>
  <n-space vertical>
    <h1>entityscrape</h1>
    <p>This a social experiment which shows the mean distance between part of speeches
(e. g. adjectives or nouns) in news articles (like from NBC or CNN) and randomly
selected entities (like Xi Jingping or ISIS).</p>

<p>The Go package <a href="https://github.com/ndabAP/assocentity">assocentity</a> was
used for creating this experiment. You can create new ones with updating the
<code>source/entities.txt</code> file and run the CLI with the provided Visual Studio Code
debug configuration.</p>

    <n-form-item size="small" label="Entity">
      <n-select v-model:value="selectedEntity" :options="availableEntities" />
    </n-form-item>
    <n-form-item  size="small" label="Part of speech">
      <n-select v-model:value="selectedPos" :options="availablePos" />
    </n-form-item>

    <Bar :options="{
      indexAxis: 'y',
      responsive: true,
    }" :data="chartData" />

   <small><b>Data source</b>:
      <a target="_blank" href="https://dataverse.harvard.edu/dataset.xhtml?persistentId=doi:10.7910/DVN/GMFCTR">dai,
        tianru, 2017, "News Articles", Harvard Dataverse, V1</a></small>
  </n-space>
</template>

<script>
import {
  defineComponent,
  onMounted,
  ref,
  computed,
  reactive,
  watch,
} from "vue";
import { Bar } from "vue-chartjs";
import entities from "../../source/entities.txt?raw";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
} from "chart.js";

ChartJS.register(
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale
);

const pos = {
  ADJ: "Adjective",
  ADV: "Adverb",
  CONJ: "Conjunction",
  DET: "Determiner",
  NOUN: "Noun",
  NUM: "Number",
  PRON: "Pronoun",
  PRT: "Preposition",
  VERB: "Verb",
};

export default defineComponent({
  components: { Bar },

  setup() {
    const selectedEntity = ref(entities.split(/\r?\n/).at(2).split(",").at(0));
    const selectedPos = ref(null);

    const availableEntities = entities.split(/\r?\n/).map((entities) => {
      return {
        label: entities.split(",").at(0),
        value: entities.split(",").at(0),
      };
    });
    const availablePos = [];
    for (const [identifier, name] of Object.entries(pos)) {
      availablePos.push({
        label: name,
        value: identifier,
      });
    }

    let meanN = [];
    const fetchEntity = async (entity) => {
      entity = entity.toLowerCase().replace(/ /g, "+");
      const response = await fetch(`${import.meta.env.BASE_URL}${entity}.json`);
      meanN = await response.json();
    };

    const filterPosMeanN = () => {
      posMeanN.splice(0);
      meanN.forEach((meanN) => {
        if (selectedPos.value === meanN.pos) {
          posMeanN.push(meanN);
        }
      });
    };

    let posMeanN = reactive([]);
    const chartData = computed(() => {
      if (posMeanN.length === 0) {
        return {
          labels: [],
          datasets: [],
        };
      }

      return {
        labels: posMeanN.map((meanN) => {
          return meanN.text;
        }),

        datasets: [
          {
            label: "Mean distances",
            data: posMeanN.map((meanN) => {
              return meanN.distance;
            }),
          },
        ],
      };
    });

    onMounted(async () => {
      await fetchEntity(selectedEntity.value);

      watch(selectedPos, () => {
        filterPosMeanN();
      });
      selectedPos.value = "ADJ";

      watch(selectedEntity, async (entity) => {
        await fetchEntity(entity);
        filterPosMeanN();
      });
    });

    return {
      chartData,

      availableEntities,
      availablePos,
      selectedEntity,
      selectedPos,
    };
  },
});
</script>

<style>
body {
  padding: 24px;
  max-width: 720px;
  margin: 0 auto;
}
</style>
