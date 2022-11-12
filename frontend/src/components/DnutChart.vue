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

const dataValues = ref([30, 40, 60]);
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
          borderWidth: 5,
        },
      ],
      borderWidth: 10
    }));

const options =
  computed <
  ChartOptions <
  "doughnut" >>
    (() => ({
      plugins: {
        legend: {
          position: "bottom",
        },
        title: {
          display: false,
          text: "Chart.js Doughnut Chart",
        },
      },
    }));

const { doughnutChartProps, doughnutChartRef } =
  useDoughnutChart({
    chartData: testData,
    options,
  });
</script>
<style scoped></style>
