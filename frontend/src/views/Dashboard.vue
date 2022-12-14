<template>
  <v-col class="dashboard">
    <!-- TITLE -->
    <p class="dash-title">Dashboard</p>

    <!-- RISK RETURN MENU -->
    <div class="risk-return">
      <div class="risk-card">
        <p class="risk-title">Risk Factor</p>
        <div class="risk-metrics">
          <p class="risk-factor" :style="{ color: riskColor }">
            {{ store.state.riskLevel }}
          </p>
          <p class="risk-label" :style="{ color: riskColor }">
            {{ riskLabel }}
          </p>
        </div>
      </div>
      <div class="risk-card">
        <p class="risk-title">Expected Return</p>
        <p class="return-metric">99%</p>
      </div>
    </div>
    <!-- PORTFOLIO PIE CHART -->
    <div class="pie-chart">
      <DnutChart
        :values="[stockPercent, savingsPercent, bondsPercent]"
      ></DnutChart>
    </div>
    <div class="money-label">
      <p class="money-label-title">Account value</p>
      <p class="money-label-metric">
        {{
          parseInt(accountValue).toLocaleString("en-US", {
            style: "currency",
            currency: "USD",
            maximumFractionDigits: 0,
          })
        }}
      </p>
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
        <p class="investment-percent">
          {{ Math.floor(stockPercent) }}%
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
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { computed } from "@vue/reactivity";
import getStockPrice from "@/util/getStockPrice.js";
import axios from "axios";
const store = useStore();
const router = useRouter();

if (store.state.token === "") {
  router.push({ path: "/login" });
}

fetchFolio();
async function fetchFolio() {
  const bonds = await axios.get(
    `${import.meta.env.VITE_API_URL}/api/bond`,
    {
      headers: {
        "Content-Type": "text/plain",
        Authorization: `Bearer ${store.state.token}`,
      },
    }
  );
  const stocks = await axios.get(
    `${import.meta.env.VITE_API_URL}/api/stock`,
    {
      headers: {
        "Content-Type": "text/plain",
        Authorization: `Bearer ${store.state.token}`,
      },
    }
  );
  const savings = await axios.get(
    `${import.meta.env.VITE_API_URL}/api/saving`,
    {
      headers: {
        "Content-Type": "text/plain",
        Authorization: `Bearer ${store.state.token}`,
      },
    }
  );

  const portfolio = {
    stocks: stocks.data.stocks,
    bonds: bonds.data.bonds,
    savings: savings.data.savings,
  };
  getStockPrice(portfolio.stocks);
  for (var i = 0; i < portfolio.savings.length; i++) {
    portfolio.savings[i].price = portfolio.savings[i].amount;
  }
  console.log(portfolio);

  const stocksValue = portfolio.stocks.reduce(
    (acc, stock) => acc + stock.price * stock.shares ,
    0
  );
  const savingsValue = portfolio.savings.reduce(
    (acc, account) => acc + parseInt(account.amount),
    0
  );
  console.log(stocksValue);
  console.log(savingsValue);
  store.state.accountValue = stocksValue + savingsValue;
  let stockP = (stocksValue / store.state.accountValue) * 100;
  let savingsP =
    (savingsValue / store.state.accountValue) * 100;
  let bondP = Math.floor(Math.random() * 10 + 1);
  stockP -= Math.floor(bondP / 2);
  savingsP -= Math.ceil(bondP / 2);

  store.state.savingsPercent = savingsP;
  store.state.stockPercent = stockP;
  store.state.bondsPercent = bondP;
  store.state.portfolio = portfolio;
}

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

const stockPercent = computed(() => store.state.stockPercent);
const savingsPercent = computed(
  () => store.state.savingsPercent
);
const bondsPercent = computed(() => store.state.bondsPercent);
const accountValue = computed(() => store.state.accountValue);
const stocksValue =
  accountValue.value * (stockPercent.value / 100);
const savingsValue =
  accountValue.value * (savingsPercent.value / 100);
const bondsValue =
  accountValue.value * (bondsPercent.value / 100);

const riskColor = ref("black");
const riskLabel = ref("");
switch (store.state.riskLevel) {
  case 0:
    riskColor.value = "#33CC00";
    riskLabel.value = "Low";
    break;
  case 1:
    riskColor.value = "#55AA00";
    riskLabel.value = "Medium";
    break;
  case 2:
    riskColor.value = "#778800";
    riskLabel.value = "Balanced";
    break;
  case 3:
    riskColor.value = "#BB4400";
    riskLabel.value = "High";
    break;
  case 4:
    riskColor.value = "#DD2200";
    riskLabel.value = "Extreme";
    break;
}
const x = computed(() => {
  console.log(store.state.riskLevel);
  switch (store.state.riskLevel) {
    case 0:
      riskColor.value = "#33CC00";
      riskLabel.value = "Low";
      break;
    case 1:
      riskColor.value = "#55AA00";
      riskLabel.value = "Medium";
      break;
    case 2:
      riskColor.value = "#778800";
      riskLabel.value = "Balanced";
      break;
    case 3:
      riskColor.value = "#BB4400";
      riskLabel.value = "High";
      break;
    case 5:
      riskColor.value = "#DD2200";
      riskLabel.value = "Extreme";
      break;
  }
  return 1;
});
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
  width: 100px;
}
.risk-metrics {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.return-metric {
  font-size: 28px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.risk-factor {
  font-size: 30px;
  text-align: center;
  font-weight: bold;
}
.risk-label {
  font-size: 15px;
  font-weight: bold;
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
</style>
