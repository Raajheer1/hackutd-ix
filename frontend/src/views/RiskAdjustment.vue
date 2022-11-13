<template>
  <div class="adjuster">
    <p class="adjust-title">Risk Adjuster</p>
    <!-- PORTFOLIO BREAKDOWN -->
    <div class="breakdown">
      <div
        class="investment-card stocks"
        @click="openStockMenu = true"
      >
        <p class="investment-label stocks-solid">Stocks</p>
        <p class="investment-percent">
          {{ Math.floor(stocksPercent) }}%
        </p>
        <p class="investment-ammount">
          {{
            stocksValue.toLocaleString("en-US", {
              style: "currency",
              currency: "USD",
              maximumFractionDigits: 0,
            })
          }}
        </p>
      </div>
      <div
        class="investment-card savings"
        @click="openSavingsMenu = true"
      >
        <p class="investment-label savings-solid">Savings</p>
        <p class="investment-percent">
          {{ Math.floor(savingsPercent) }}%
        </p>
        <p class="investment-ammount">
          {{
            savingsValue.toLocaleString("en-US", {
              style: "currency",
              currency: "USD",
              maximumFractionDigits: 0,
            })
          }}
        </p>
      </div>
      <div
        class="investment-card bonds"
        @click="openBondMenu = true"
      >
        <p class="investment-label bonds-solid">Bonds</p>
        <p class="investment-percent">
          {{ Math.floor(bondsPercent) }}%
        </p>
        <p class="investment-ammount">
          {{
            bondsValue.toLocaleString("en-US", {
              style: "currency",
              currency: "USD",
              maximumFractionDigits: 0,
            })
          }}
        </p>
      </div>
    </div>
    <!-- PORTFOLIO PIE CHART -->
    <div class="pie-chart">
      <DnutChart
        :values="[stocksPercent, savingsPercent, bondsPercent]"
      ></DnutChart>
    </div>
    <!-- PIE CHART CHANGES LEGEND -->
    <div class="changes-legend">
      <div class="change">
        <div class="change-box stocks-solid"></div>
        <p class="change-label">Stocks:</p>
        <p class="change-metric">-0.23%</p>
      </div>
      <div class="change">
        <div class="change-box savings-solid"></div>
        <p class="change-label">Savings:</p>
        <p class="change-metric">-0.23%</p>
      </div>
      <div class="change">
        <div class="change-box bonds-solid"></div>
        <p class="change-label">Bonds:</p>
        <p class="change-metric">-0.23%</p>
      </div>
    </div>
    <!-- Slider Labels -->
    <div class="slider-labels">
      <p class="slider-label left">
        Risk Factor:
        <span class="slider-value risk">{{ riskValue }}</span>
      </p>
      <p class="slider-label right">
        Return:
        <span class="slider-value">15%</span>
      </p>
    </div>
    <!-- RISK/REWARD SLIDER -->
    <div class="risk-slider">
      <v-slider
        v-model="riskSlider"
        color="#DCCBAB"
        step="1"
        track-size="9"
        thumb-size="33"
        show-ticks="always"
        hide-details
        :min="0"
        :max="4"
      ></v-slider>
      <p class="risk-messages">
        {{ sliderMessages[riskValue] }}
      </p>
    </div>
  </div>
</template>

<script setup>
import DnutChart from "@/components/DnutChart.vue";
import { ref, computed } from "vue";
import { useStore } from "vuex";

const store = useStore();
const riskSlider = ref(store.state.riskLevel);
const riskValue = computed(() => riskSlider.value);

const stocksPercent = ref(store.state.stockPercent);
const savingsPercent = ref(store.state.savingsPercent);
const bondsPercent = ref(store.state.bondsPercent);

const accountValue = store.state.accountValue;
const stocksValue = accountValue * (stocksPercent.value / 100);
const savingsValue = accountValue * (savingsPercent.value / 100);
const bondsValue = accountValue * (bondsPercent.value / 100);


const sliderMessages = {
  0: "Low Risk",
  1: "Moderate Risk",
  2: "Balanced Risk",
  3: "High yield",
  4: "Bet the farm",
};
</script>
<style scoped>
.adjuster {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 60px;
}
.adjust-title {
  margin-bottom: 10px;
  width: 100%;
  justify-content: center;
  font-size: 30px;
  font-weight: bold;
  text-align: center;
}
.breakdown {
  display: flex;
  flex-direction: row;
  width: 100%;
}
.pie-chart {
  width: 366px;
}
.changes-legend {
  position: absolute;
  left: 50%;
  transform: translate(-50%);
  top: 386px;
}
.change {
  display: flex;
  justify-content: center;
}
.change-box {
  width: 20px;
  height: 20px;
  border-radius: 5px;
  margin-bottom: 5px;
  margin: auto;
}
.change-label {
  font-size: 17px;
  margin: auto;
  width: 65px;
  margin-inline: 7px;
  font-weight: bold;
}
.change-metric {
  font-size: 17px;
  margin: auto;
  font-weight: bold;
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
  background-color: #7fcef585;
}
.savings {
  background-color: #f57fce85;
}
.bonds {
  background-color: #f5ed7f85;
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
.risk-slider {
  width: 87%;
  height: 60px;
}
.risk-messages {
  width: 100%;
  text-align: center;
  margin-top: 10px;
  font-size: 20px;
  font-weight: bold;
}
.slider-labels {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
  margin-bottom: 20px;
}
.slider-label {
  font-size: 20px;
  font-weight: bold;
  margin: auto;
  background-color: #d8d8d8bc;
  padding: 10px;
  border-radius: 10px;
  height: 52px;
}
.slider-label.left {
  width: 193px;
}
.slider-value {
  color: #e39400;
  font-weight: bold;
  font-size: 21px;
  margin: auto;
}
.slider-value.risk {
  margin-left: 15px;
  font-size: 21px;
  font-weight: bold;
}
</style>
