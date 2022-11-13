import { createStore } from "vuex";

export default createStore({
    state:{
        token:"",
        accountValue: 10000,
        stockPercent: 55,
        bondsPercent: 15,
        savingsPercent: 30,
        riskFactor: 13,
        riskLevel: 3
    }
})