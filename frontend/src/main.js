import { createApp } from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import { loadFonts } from "./plugins/webfontloader";
import router from "@/router";
import store from "@/store";
import Vue3TouchEvents from "vue3-touch-events";

loadFonts();

createApp(App)
  .use(vuetify)
  .use(router)
  .use(store)
  .use(Vue3TouchEvents, {
    touchClass: "active",
  })
  .mount("#app");
