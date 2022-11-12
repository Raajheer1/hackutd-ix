<template>
  <v-col class="dashboard">
    <!-- TITLE -->
    <p class="dash-title">Dashboard</p>

    <!-- RISK RETURN MENU -->
    <div class="risk-return">
      <div class="risk-card">
        <p class="risk-title">Current Risk</p>
        <p class="risk-metric">99%</p>
      </div>
      <div class="risk-card">
        <p class="risk-title">Projected Return</p>
        <p class="risk-metric">99%</p>
      </div>
    </div>
    <!-- PORTFOLIO PIE CHART -->
    <div class="pie-chart">
      <DnutChart></DnutChart>
    </div>
    <!-- TIME SELECTOR -->
    <div class="time-selector"></div>
    <!-- PORTFOLIO BREAKDOWN -->
    <div class="breakdown"></div>
  </v-col>
</template>

<script setup>
import DnutChart from "@/components/DnutChart.vue"
</script>
<style scoped>
.dashboard {
  width: 100vw;
  height: 100vh;
  align-items: center;
}
.dash-title {
  margin-bottom: 10px;
  width: 100%;
  justify-content: center;
  font-size: 30px;
  font-weight: bold;
  text-align: center;
}
.risk-return {
  display: flex;
  flex-wrap: nowrap;
  justify-content: center;
}
.risk-card {
  width: 160px;
  margin: auto;
  display: flex;
  border-color: black;
  border-style: solid;
  border-width: 1px;
  border-radius: 15px;
  padding: 10px;
  justify-content: center;
}
.risk-title {
  font-weight: bold;
  font-size: 18px;
  width: fit-content;
}
.risk-metric {
  font-size: 28px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
