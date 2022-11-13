import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate";

export default createStore({
  state: {
    token: "",
    accountValue: 10000,
    stockPercent: 55,
    bondsPercent: 15,
    savingsPercent: 30,
    riskFactor: 13,
    riskLevel: 3,
    portfolio: Object
  },
  plugins: [createPersistedState()],
});
