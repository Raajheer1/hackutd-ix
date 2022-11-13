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
        <p class="risk-title">Expected Return</p>
        <p class="risk-metric">99%</p>
      </div>
    </div>
    <!-- PORTFOLIO PIE CHART -->
    <div class="pie-chart">
      <DnutChart></DnutChart>
    </div>
    <div class="money-label">
      <p class="money-label-title">Account value</p>
      <p class="money-label-metric">$10,000</p>
    </div>
    <!-- BREAKDOWN SUBTITLE -->
    <div class="breakdown-title">Breakdown</div>
    <!-- PORTFOLIO BREAKDOWN -->
    <div class="breakdown">
      <div
        class="investment-card stocks"
        @click="openStockMenu = true"
      >
        <p class="investment-label stocks-solid">Stocks</p>
        <p class="investment-percent">33%</p>
        <p class="investment-ammount">$3,333</p>
      </div>
      <ExpandMenu
        :expand="openStockMenu"
        @collapse="openStockMenu = false"
        :items="stocks"
        menu-type="Stock"
      />
      <div
        class="investment-card savings"
        @click="openSavingsMenu = true"
      >
        <p class="investment-label savings-solid">Savings</p>
        <p class="investment-percent">33%</p>
        <p class="investment-ammount">$3,333</p>
      </div>
      <ExpandMenu
        :expand="openSavingsMenu"
        @collapse="openSavingsMenu = false"
        :items="savings"
        menu-type="Savings"
      />
      <div
        class="investment-card bonds"
        @click="openBondMenu = true"
      >
        <p class="investment-label bonds-solid">Bonds</p>
        <p class="investment-percent">33%</p>
        <p class="investment-ammount">$3,333</p>
      </div>
      <ExpandMenu
        :expand="openBondMenu"
        @collapse="openBondMenu = false"
        :items="bonds"
        menu-type="Bond"
      />
    </div>
  </v-col>
</template>

<script setup>
import DnutChart from "@/components/DnutChart.vue";
import ExpandMenu from "../components/ExpandMenu.vue";
import { ref } from "vue";
const stocks = [
  { name: "Apple", shares: 3, price: 98.88 },
  { name: "Google", shares: 4, price: 3143.62 },
  { name: "Amazon", shares: 4, price: 132.44 },
  { name: "Tesla", shares: 5, price: 324.94 },
  { name: "GME", shares: 1124, price: 26.44 },
];
const savings = [
  { name: "Bank of America", price: 15234 },
  { name: "Fidelity 401k", price: 24912 },
];
const bonds = [
  { name: "Treasury", shares: 1, price: 543 },
  { name: "Treasury", shares: 2, price: 442 },
  { name: "Treasury", shares: 1, price: 1392 },
];
const openStockMenu = ref(false);
const openSavingsMenu = ref(false);
const openBondMenu = ref(false);
</script>
<style scoped>
.dashboard {
  width: 100vw;
  height: 100vh;
  align-items: center;
  padding-top: 60px;
}
.dash-title {
  margin-bottom: 20px;
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
  /* margin-bottom: 5px; */
}
.risk-card {
  width: 175px;
  margin: auto;
  display: flex;
  border-radius: 15px;
  padding: 10px;
  justify-content: center;
  background-color: #d8d8d8bc;
}
.risk-title {
  font-weight: bold;
  font-size: 21px;
  width: fit-content;
}
.risk-metric {
  font-size: 28px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.money-label {
  position: absolute;
  top: 355px;
  left: 50%;
  transform: translate(-50%);
}
.money-label-title {
  font-size: 18px;
  color: rgba(164, 164, 164, 0.733);
  text-align: center;
  margin-bottom: -6px;
}
.money-label-metric {
  font-size: 35px;
  text-align: center;
}
.breakdown-title {
  font-size: 27px;
  font-weight: bold;
  margin-bottom: 10px;
  margin-left: 10px;
  margin-top: -10px;
}
.breakdown {
  display: flex;
  flex-direction: row;
  width: 100%;
}
.investment-card {
  border-radius: 15px;
  margin: auto;
  padding: 10px;
  justify-content: center;
  align-items: center;
  width: 115px;
}
.investment-percent {
  font-size: 30px;
  text-align: center;
}
.investment-ammount {
  font-size: 20px;
  text-align: center;
}
.investment-label {
  font-size: 22px;
  font-weight: bold;
  text-align: center;
  border-radius: 10px;
}
.stocks {
  background-color: #7fcef572;
}
.savings {
  background-color: #f57fce72;
}
.bonds {
  background-color: #f5ed7f72;
}
.stocks-solid {
  background-color: #7fcef5;
}
.savings-solid {
  background-color: #f57fce;
}
.bonds-solid {
  background-color: #f5ed7f;
}
</style>
