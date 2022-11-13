import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "@/views/Dashboard.vue";
import RiskAdjustment from "@/views/RiskAdjustment.vue";
import Portfolio from "@/views/Portfolio.vue"
import LoginScreen from "@/views/LoginScreen.vue";
const routes = [
  {
    path: "/",
    component: Dashboard,
  },
  {
    path: "/risk-adjust",
    component: RiskAdjustment,
  },
  {
    path: "/portfolio",
    component: Portfolio,
  },
  {
    path: "/login",
    component: LoginScreen,
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
