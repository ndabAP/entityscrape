<template>
    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
    consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
    cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat

    <div ref="tester"></div>
</template>

<script setup>
import Plotly from 'plotly.js-dist-min';
import { onMounted, ref } from 'vue';

const tester = ref();

onMounted(() => {
    var data = {
        type: "sankey",
        orientation: "h",
        node: {
            pad: 15,
            thickness: 30,
            line: {
                color: "black",
                width: 0.5
            },
            label: ["A1", "A2", "B1", "B2", "C1", "C2"],
            color: ["blue", "blue", "blue", "blue", "blue", "blue"]
        },

        link: {
            source: [0, 1, 0, 2, 3, 3],
            target: [2, 3, 3, 4, 4, 5],
            value: [8, 4, 2, 8, 4, 2]
        }
    }

    var data = [data]

    var layout = {
        paper_bgcolor: "rgba(0,0,0,0", //background color of the chart container space
        plot_bgcolor: "rgba(0,0,0,0)", //background color of plot area
    }
    Plotly.react(tester.value, data, layout, { displayModeBar: false })
})
</script>