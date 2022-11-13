<template>
  <DoughnutChart v-bind="doughnutChartProps" />
</template>

<script setup lang="ts">
import { computed, defineComponent, ref } from "vue";
import { DoughnutChart, useDoughnutChart } from "vue-chart-3";
import {
  Chart,
  ChartData,
  ChartOptions,
  registerables,
} from "chart.js";
Chart.register(...registerables);
const props = defineProps({
  values: {
    type: Array,
    default: [33,33,33]
  }
})
const dataValues = computed(()=>props.values);
const dataLabels = ref([
  "Stocks",
  "Saving",
  "Bonds",
]);

const testData =
  computed <
  ChartData <
  "doughnut" >>
    (() => ({
      labels: dataLabels.value,
      datasets: [
        {
          data: dataValues.value,
          backgroundColor: [
          "#7fcef5", "#f57fce", "#f5ed7f"
          ],
          borderWidth: 10,
          borderRadius: 15,
        },
      ],
    }));

const options =
  computed <
  ChartOptions <
  "doughnut" >>
    (() => ({
      plugins: {
        legend: {
          display: false,
          position: "left",
          align:"center",
          labels:{
            display: true,
            font:{
                size: 1
            },
            boxHeight: 20,
            boxWidth: 20,
            color: 'black',
            fullSize: false,
          }
        },
        title: {
          display: false,
          text: "Chart.js Doughnut Chart",
        },
      },
      cutout: 120,
      radius: '80%',
      borderColor: "#FAF9F6"
    }));

const { doughnutChartProps, doughnutChartRef } =
  useDoughnutChart({
    chartData: testData,
    options:options,
  });
</script>
<style scoped></style>
