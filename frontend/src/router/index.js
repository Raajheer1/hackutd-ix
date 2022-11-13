import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "@/views/Dashboard.vue";
import RiskAdjustment from "@/views/RiskAdjustment.vue";
import Portfolio from "@/views/Portfolio.vue"
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
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
