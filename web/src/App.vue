<template>
  <n-space vertical>
    <n-select v-model="selectedEntity" :options="availableEntities" />
  </n-space>
</template>

<script>
import { defineComponent, onMounted, ref, reactive, watch } from "vue";
import { Bar } from "vue-chartjs";
import entities from "../../source/entities.csv?raw";

export default defineComponent({
  components: { Bar },

  setup() {
    const availableEntities = entities.split(/\r?\n/).map((entities) => {
      return {
        label: entities.split(",").at(0),
        value: entities.split(",").at(0),
      };
    });
    const selectedEntity = ref(null);

    watch(selectedEntity, () => { });

    return {
      availableEntities,
      selectedEntity,
    };
  },
});
</script>

<style>
body {
  padding: 24px;
}
</style>
