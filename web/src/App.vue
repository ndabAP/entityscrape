<template>
  <n-space vertical>
    <h1>entityscrape</h1>
    <p>
      This a social experiment which shows the mean distance between part of
      speeches (e. g. adjectives or nouns) in news articles (like from NBC or
      CNN) and randomly selected entities (like Xi Jingping or ISIS).
    </p>

    <p>
      The Go package
      <a href="https://github.com/ndabAP/assocentity">assocentity</a> was used
      for creating this experiment. You can create new ones with updating the
      <code>source/entities.txt</code> file and run the CLI with the provided
      Visual Studio Code debug configuration. The experiments source code can be
      found at
      <a target="_blank" href="https://github.com/ndabAP/entityscrape">Github</a
      >.
    </p>

    <n-form-item size="small" label="Entity">
      <n-select v-model:value="selectedEntity" :options="availableEntities" />
    </n-form-item>
    <n-form-item size="small" label="Part of speech">
      <n-select v-model:value="selectedPos" :options="availablePos" />
    </n-form-item>

    <div style="height: 475px">
      <Bar
        :data="chartData"
        :options="{
          indexAxis: 'y',
          maintainAspectRatio: false,
          responsive: true,

          scales: {
            x: {
              grid: {
                drawBorder: false,
                display: false,
              },
            },

            y: {
              ticks: {
                font: {
                  size: 13,
                  family:
                    'v-sans, system-ui, -apple-system, BlinkMacSystemFont',
                },
              },
              grid: {
                display: false,
              },
            },
          },
        }"
      />
    </div>

    <small>
      <b>Data source</b>:
      <a
        target="_blank"
        href="https://dataverse.harvard.edu/dataset.xhtml?persistentId=doi:10.7910/DVN/GMFCTR"
        >dai, tianru, 2017, "News Articles", Harvard Dataverse, V1</a
      >
    </small>
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
import { darkTheme } from "naive-ui";

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
  ADP: "Adposition",
  ADV: "Adverb",
  CONJ: "Conjunction",
  DET: "Determiner",
  NOUN: "Noun",
  NUM: "Number",
  PRON: "Pronoun",
  PRT: "Particle",
  VERB: "Verb",
};

export default defineComponent({
  components: { Bar },

  setup() {
    const selectedEntity = ref(entities.split(/\r?\n/).at(0).split(",").at(0));
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

    let mean = [];
    const fetchEntity = async (entity) => {
      entity = entity.toLowerCase().replace(/ /g, "+");
      const response = await fetch(`${import.meta.env.BASE_URL}${entity}.json`);
      mean = await response.json();
    };

    const filterPosMean = () => {
      posMean.splice(0);
      mean.forEach((meanN) => {
        if (selectedPos.value === meanN.pos) {
          posMean.push(meanN);
        }
      });
    };

    let posMean = reactive([]);
    const chartData = computed(() => {
      if (posMean.length === 0) {
        return {
          labels: [],
          datasets: [],
        };
      }

      return {
        labels: posMean.map((mean) => {
          return mean.text;
        }),

        datasets: [
          {
            label: "Mean distances",
            data: posMean.map((mean) => {
              return mean.distance;
            }),
          },
        ],
      };
    });

    onMounted(async () => {
      await fetchEntity(selectedEntity.value);

      watch(selectedPos, () => {
        filterPosMean();
      });
      selectedPos.value = "ADJ";

      watch(selectedEntity, async (entity) => {
        await fetchEntity(entity);
        filterPosMean();
      });
    });

    return {
      darkTheme,

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
