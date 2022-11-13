import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "@/views/Dashboard.vue";
import Splash from "@/views/Splash.vue";
import RiskAdjustment from "@/views/RiskAdjustment.vue";
import Portfolio from "@/views/Portfolio.vue"
import LoginScreen from "@/views/LoginScreen.vue";
import RegisterScreen from "@/views/RegisterScreen.vue";
const routes = [
  {
    path: "/",
    component: Dashboard,
  },
  {
    path: "/splash",
    component: Splash,
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
  {
    path: "/register",
    component: RegisterScreen,
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
