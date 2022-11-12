<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',

  components: {
    HelloWorld,
  },

  data: () => ({
    //
  }),
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Outfit&display=swap');
*{
  font-family: 'Outfit', sans-serif;
}
</style>
