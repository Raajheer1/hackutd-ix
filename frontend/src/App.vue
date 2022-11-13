<template>
  <router-view />
  <NavBar v-if="store.state.token !== ''"></NavBar>
</template>

<script setup>
import NavBar from "./components/NavBar.vue";

import {useStore} from 'vuex';
const store = useStore();

</script>
<style>
@import url("https://fonts.googleapis.com/css2?family=Montserrat&display=swap");
* {
  font-family: "Montserrat", sans-serif;
  color: #313131;
}
#app{
  height:100%;
  overflow-y: scroll;
  -webkit-overflow-scrolling: touch;
  background-color: #FAF9F6;
}
</style>
