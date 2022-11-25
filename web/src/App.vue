<template>
  <n-space vertical>
    <n-select v-model="value" :options="options" />
  </n-space>
</template>

<script>
import { defineComponent, ref } from "vue";
import { Bar } from "vue-chartjs";

export default defineComponent({
  components: { Bar },
  setup() {
    return {
      value: ref(null),
      options: [
        {
          label: "Everybody's Got Something to Hide Except Me and My Monkey",
          value: "song0",
        },
      ],
    };
  },
});
</script>

<style>
body {
  padding: 24px;
}
</style>