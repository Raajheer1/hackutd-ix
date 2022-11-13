<template>
  <div class="portfolio">
    <p class="folio-title">Portfolio</p>
    <div class="folio-table">
      <v-table class="table Stock">
        <thead>
          <tr>
            <th class="table-head text-left">Stock</th>
            <th class="text-left table-head">Shares</th>
            <th class="text-right table-head">Price</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in stocks">
            <td class="table-entry">{{ item.name }}</td>
            <td class="table-entry text-center">
              {{ item.shares }}
            </td>
            <td class="table-entry text-right">
              {{
                item.price.toLocaleString("en-US", {
                  style: "currency",
                  currency: "USD",
                })
              }}
            </td>
          </tr>
          <tr>
            <td></td>
            <td class="table-head">Total:</td>
            <td class="table-head">$15,000</td>
          </tr>
        </tbody>
      </v-table>
    </div>
    <div class="folio-table">
      <v-table class="table Savings">
        <thead>
          <tr>
            <th class="table-head text-left">Account</th>
            <th class="text-left table-head"></th>
            <th class="text-right table-head">Price</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in savings">
            <td class="table-entry">{{ item.name }}</td>
            <td class="table-entry text-center"></td>
            <td class="table-entry text-right">
              {{
                item.price.toLocaleString("en-US", {
                  style: "currency",
                  currency: "USD",
                })
              }}
            </td>
          </tr>
        </tbody>
      </v-table>
    </div>
    <div class="folio-table">
      <v-table class="table Bond">
        <thead>
          <tr>
            <th class="table-head text-left smaller">Bond</th>
            <th class="text-left table-head smaller count">Count</th>
            <th class="text-left table-head smaller yield">Yield</th>
            <th class="text-right table-head smaller">Price</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in bonds">
            <td class="table-entry smaller">{{ item.name }}</td>
            <td class="table-entry text-center smaller">
              {{ item.shares }}
            </td>
            <td class="table-entry text-center smaller">
              {{ item.yield }}
            </td>
            <td class="table-entry text-right smaller">
              {{
                item.price.toLocaleString("en-US", {
                  style: "currency",
                  currency: "USD",
                })
              }}
            </td>
          </tr>
        </tbody>
      </v-table>
    </div>
    <!-- Spacer to force scroll -->
    <div class="scroll-spacer"></div>
  </div>
</template>

<script setup>
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
  { name: "Treasury", shares: 1, yield: 3.45, price: 543 },
  { name: "Treasury", shares: 2, yield: 4.01, price: 442 },
  { name: "Treasury", shares: 1, yield: 2.82, price: 1392 },
];
</script>
<style scoped>
.portfolio {
  width: 100%;
  height: 100%;
  height: fit-content;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 50px;
}
.folio-title {
  margin-bottom: 5px;
  width: 100%;
  justify-content: center;
  font-size: 30px;
  font-weight: bold;
  text-align: center;
}
.folio-table {
  width: 90%;
  margin-bottom: 10px;
  border-radius: 15px;
}
.table {
  width: 100% !important;
}
.table-head {
  font-size: 24px !important;
  font-weight: bold;
  color: white;
}
.table-head.smaller {
  font-size: 18px !important;
}
.table-head.smaller.count{
    padding: 0;
}
.table-head.smaller.yield{
    padding-right: 0;
}
.table-entry {
  font-size: 20px !important;
  color: white;
}
.table-entry.smaller {
  font-size: 17px !important;
}

.table-right {
  text-align: right;
}
.Stock {
  background-color: #58cef5;
}
.Savings {
  background-color: #f75ac3;
}
.Bond {
  background-color: #dfd438;
}
.scroll-spacer {
  width: 100%;
  height: 200px;
}
</style>
