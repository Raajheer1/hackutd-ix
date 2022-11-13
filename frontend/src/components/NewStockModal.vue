<template>
  <Modal @close="$emit('close')">
    <h2 class="modal-title">Input Stock Information</h2>
    <v-text-field
      label="Ticker"
      variant="outlined"
      class="input"
      v-model="securityName"
    ></v-text-field>
    <v-text-field
      label="Number of shares"
      variant="outlined"
      class="input"
      v-model="securityCount"
    ></v-text-field>

    <v-btn
      v-if="!showConfirmed"
      variant="tonal"
      color="#7fcef5"
      @click="submitSecurity"
    >
      Add stock
    </v-btn>
    <ConfirmedIcon :size="40" v-else></ConfirmedIcon>
  </Modal>
</template>

<script setup>
import Modal from "./Modal.vue";
import axios from "axios";
import ConfirmedIcon from "./icons/ConfirmedIcon.vue";
import { ref } from "vue";
const emit = defineEmits(["close"]);
const showConfirmed = ref(false);
const securityCount = ref();
const securityName = ref("");
async function submitSecurity() {
  console.log(securityCount.value, securityName.value);
  const userdata = await axios.get(
    `${import.meta.env.VITE_API_URL}/api/user`,
    {
      headers: {
        "Content-Type": "application/json",
        Authorization:
          "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Njg0MzAyMTUsInVzZXJfaWQiOjF9.YRAOnrg_X3Bf6ypAMue1_1DdAfA2jzxkoqwvku4sNxk",
      },
    }
  );
  console.log(userdata);
  let user_id = userdata.data.data.id;
  console.log(user_id);
  console.log(parseInt(securityCount.value));
  console.log("Ticker: " + securityName.value);
  const response = await axios.post(
    `${import.meta.env.VITE_API_URL}/api/stock/`,
    {
      ticker: securityName.value,
      shares: parseInt(securityCount.value),
      user_id: user_id,
    },
    {
      headers: {
        "Content-Type": "application/json",
        Authorization:
          "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Njg0MzAyMTUsInVzZXJfaWQiOjF9.YRAOnrg_X3Bf6ypAMue1_1DdAfA2jzxkoqwvku4sNxk",
      },
    }
  );

  console.log(response);
  showConfirmed.value = true;
  setTimeout(() => emit("close"), 1000);
}
</script>
<style scoped>
.modal-title {
  margin: 5px;
  margin-bottom: 20px;
}
.input {
  border-radius: 10px !important;
}
</style>
