<template>
  <v-col class="login-screen">
    <!-- Prompt for the user -->
    <p class="login-title">Enter your information below.</p>

    <!-- Username text -->
    <div class="username-title">
      <p> Email: </p>
    </div>
    <!-- Input box for username -->
    <div class="input-box">
      <v-text-field
          label="Enter your Email"
          type="email"
          v-model="email"
      ></v-text-field>
    </div>

    <div class="username-title">
      <p> Age: </p>
    </div>
    <!-- Input box for username -->
    <div class="input-box">
      <v-text-field
          label="Enter your Age"
          type="input"
          v-model="age"
      ></v-text-field>
    </div>
    <div class="username-title">
      <p> Income: </p>
    </div>
    <!-- Input box for username -->
    <div class="input-box">
      <v-text-field
          label="Enter your Income"
          type="Income"
          placeholder="10000"
          v-model="income"
      ></v-text-field>
    </div>
    <!-- Password text -->
    <div class="password-title">
      <p>Password: </p>
    </div>
    <!-- Input box for password -->
    <div class="input-box">
      <v-text-field
          label="Enter your password"
          type="password"
          v-model="password"
      ></v-text-field>
    </div>
    <!-- Button input -->
    <div class="login-button">
      <v-btn class="button-properties"
             v-on:click = clickButton()
             size="x-large"
             variant="flat"
             color="#9bd5f2"
      >
        Login
      </v-btn>
    </div>
  </v-col>
</template>
<script setup>
import {useRouter} from 'vue-router'
import {ref} from 'vue'
const email = ref("")
const password = ref("")
const age = ref(18)
const income = ref(10000)
import {useStore} from 'vuex';
const store = useStore();
import axios from 'axios';
const router = useRouter()
function clickButton(){
  axios.post('https://hackutd.raajpatel.dev/auth/register', {
    email: email.value,
    password: password.value,
    age: age.value,
    income: income.value
  })
      .then((res) => {
        store.state.token = res.data.token;
        router.push({path: '/'})
      })
      .catch((err) => {
        console.log("Invalid credentials")
      })
}
</script>

<style>
.login-screen {
  width: 100vw;
  height: 100vh;
  align-items: center;
  padding-top: 100px;
}

.login-title {
  margin-bottom: 45px;
  width: 100%;
  justify-content: center;
  font-size: 25px;
  font-weight: bold;
  text-align: center;
}

.username-title {
  font-size: 25px;
}

.password-title {
  font-size: 25px;
}

.input-box {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 10px;
  margin-top: 10px;
}

.login-button{
  display: flex;
  justify-content: center;
  margin-top: 30px;
}

</style>