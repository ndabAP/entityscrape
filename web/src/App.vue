<template>
  <n-space vertical>
    <n-select v-model:value="selectedEntity" :options="availableEntities" />
    <n-select v-model:value="selectedPos" :options="availablePos" />

    <Bar
      :options="{
        responsive: true,
      }"
      :data="chartData"
    />
  </n-space>
</template>

<script>
import { defineComponent, onMounted, ref, computed, watch } from "vue";
import { Bar } from "vue-chartjs";
import entities from "../../source/entities.csv?raw";
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
  ANY: "Any",
  ADJ: "Adjective",
  ADV: "Adverb",
  AFFIX: "Affix",
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
    const selectedEntity = ref(null);
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

    let meanN = null;
    const fetchEntity = async (entity) => {
      entity = entity.toLowerCase().replace(/ /g, "+");
      const response = await fetch(`/${entity}.json`);
      meanN = await response.json();
    };

    let posMeanN = []
    const chartData = computed(() => {
      if (posMeanN.length === 0) {
        return {
          labels: [],
          datasets: [],
        };
      }

      // Update chart

      return {
        labels: posMeanN.map((meanN) => {
          return meanN.text;
        }),
        datasets: [
          {
            data: posMeanN.map((meanN) => {
              return meanN.distance;
            }),
          },
        ],
      };
    });

    watch(selectedEntity, async () => {
      await fetchEntity(selectedEntity);
    });
    watch(selectedPos, () => {
      posMeanN = meanN.filter((meanN) => {
        return selectedPos.value === meanN.pos;
      });
    });

    onMounted(async () => {
      await fetchEntity(entities.split(/\r?\n/).at(0).split(",").at(0));
      selectedPos.value = "NOUN";
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
