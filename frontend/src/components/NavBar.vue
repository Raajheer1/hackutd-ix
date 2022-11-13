<template>
  <div class="navbar">
    <div class="nav-item side" @click="slidersClick">
      <SlidersIcon
        :size="42"
        :color="slidersActive ? activeColor : inactiveColor"
      ></SlidersIcon>
    </div>
    <div class="nav-item" @click="dnutClick">
      <DnutIcon
        :size="60"
        :color="dnutActive ? activeColor : inactiveColor"
      ></DnutIcon>
    </div>
    <div class="nav-item side" @click="piggyClick">
      <PiggyBank
        :size="45"
        :color="piggyActive ? activeColor : inactiveColor"
      ></PiggyBank>
    </div>
  </div>
</template>

<script setup>
import DnutIcon from "./icons/DnutIcon.vue";
import PiggyBank from "./icons/PiggyBank.vue";
import SlidersIcon from "./icons/SlidersIcon.vue";
import { ref } from "vue";

const slidersActive = ref(false);
const dnutActive = ref(true);
const piggyActive = ref(false);

const activeColor = "#313131";
const inactiveColor = "lightGray";

function slidersClick() {
  piggyActive.value = false;
  dnutActive.value = false;
  slidersActive.value = true;
}
function dnutClick() {
  piggyActive.value = false;
  slidersActive.value = false;
  dnutActive.value = true;
}
function piggyClick() {
  slidersActive.value = false;
  dnutActive.value = false;
  piggyActive.value = true;
}
</script>
<style scoped>
.navbar {
  width: 100vw;
  position: fixed;
  bottom: 0;
  left: 0;
  display: flex;
  flex-direction: row;
  padding-bottom: 10px;
}
.nav-item {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.side {
  padding-bottom: 10px;
}
</style>
