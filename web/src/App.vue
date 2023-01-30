<template>
  <n-space vertical>
    <n-select v-model="entity" :options="entities" />
  </n-space>
</template>

<script>
import { defineComponent, onMounted, ref, reactive, watch } from "vue";
import { Bar } from "vue-chartjs";
import entities from "../../source/entities.csv?raw";

export default defineComponent({
  components: { Bar },

  setup() {
    const entity = ref(null);

    watch(entity, () => {});

    return {
      entity,
      entities: entities.split(/\r?\n/).map((entities) => {
        return {
          label: entities.split(",").at(0),
          value: entities.split(",").at(0),
        };
      }),
    };
  },
});
</script>

<style>
body {
  padding: 24px;
}
</style>
