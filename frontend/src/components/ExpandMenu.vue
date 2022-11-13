<template>
  <div
    v-if="expand"
    @click="emit('collapse')"
    class="expand-background"
  ></div>
  <Transition :name="props.menuType">
    <div v-if="expand" class="expander">
      <v-table class="table" :class="[props.menuType]">
        <thead>
          <tr>
            <th class="table-head text-left">
              {{ props.menuType }}
            </th>
            <th class="text-left table-head">{{column2Header}}</th>
            <th class="text-right table-head">{{column3Header}}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in items" :key="item.name">
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
        </tbody>
      </v-table>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed } from "vue";
const props = defineProps({
  items: Array,
  menuType: String,
  expand: Boolean,
});
const emit = defineEmits(["collapse"]);

const items = computed(() => props.items);
const expand = computed(() => props.expand);
const column2Header = computed(() => {
  if (props.menuType == "Stock") return "Shares";
  if (props.menuType == "Savings") return "";
  if (props.menuType == "Bonds") return "Count";
});
const column3Header = computed(() => {
  if (props.menuType == "Stock") return "Price";
  if (props.menuType == "Savings") return "Ammount";
  if (props.menuType == "Bonds") return "Value";
});
</script>
<style scoped>
.expander {
  position: absolute;
  bottom: 81px;
  left: 16px;
  width: 91%;
  border-radius: 15px;
}
.expand-background {
  width: 100vw;
  height: 100vh;
  top: 0;
  left: 0;
  position: fixed;
}
.table {
  min-height: 450px;
}
.Stock {
  background-color: #7fcef5;
}
.Savings {
  background-color: #f57fce;
}
.Bond {
  background-color: #f5ed7f;
}
.table-head {
  font-size: 25px !important;
  font-weight: bold;
  color: white;
}
.table-entry {
  font-size: 20px !important;
  color: white;
}
.table-right {
  text-align: right;
}

.Stock-enter-active,
.Stock-leave-active {
  animation: expand-left 0.3s;
}

.Stock-enter-from,
.Stock-leave-to {
  animation: expand-left 0.3s reverse;
}
@keyframes expand-left {
  0% {
    transform-origin: bottom left;
    transform: scale(0);
  }
  100% {
    transform-origin: bottom left;
    transform: scale(1);
  }
}
.Savings-enter-active,
.Savings-leave-active {
  animation: expand-center 0.3s;
}

.Savings-enter-from,
.Savings-leave-to {
  animation: expand-center 0.3s reverse;
}
@keyframes expand-center {
  0% {
    transform-origin: bottom center;
    transform: scale(0);
  }
  100% {
    transform-origin: bottom center;
    transform: scale(1);
  }
}
.Bond-enter-active,
.Bond-leave-active {
  animation: expand-right 0.3s;
}

.Bond-enter-from,
.Bond-leave-to {
  animation: expand-right 0.3s reverse;
}
@keyframes expand-right {
  0% {
    transform-origin: bottom right;
    transform: scale(0);
  }
  100% {
    transform-origin: bottom right;
    transform: scale(1);
  }
}
</style>
