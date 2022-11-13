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
import { ref, computed } from "vue";
import {useRouter} from "vue-router"
import {useRoute} from "vue-router"
const router = useRouter();
const route = useRoute();
console.log(route.path)

const slidersActive = computed(() => route.path == '/risk-adjust');
const dnutActive = computed(() => route.path == '/');
const piggyActive = computed(() => route.path == '/portfolio');;

const activeColor = "#313131";
const inactiveColor = "#afafaf";

function slidersClick() {
  router.push({path:"/risk-adjust"})
}
function dnutClick() {
  router.push({path:"/"})
}
function piggyClick() {
  router.push({path:'/portfolio'})
}
</script>
<style scoped>
.navbar {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  display: flex;
  flex-direction: row;
  padding-bottom: 10px;
  background-color: #FAF9F6;
  color: #afafaf;
  z-index: 101;
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
