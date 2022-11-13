<template>
  <router-view />
  <NavBar></NavBar>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import NavBar from "./components/NavBar.vue";

export default {
  name: "App",

  components: {
    HelloWorld,
    NavBar,
  },

  data: () => ({
    //
  }),
};
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
