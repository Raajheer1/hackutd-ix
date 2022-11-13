<template>
  <Modal @close="$emit('close')">
    <h2 class="modal-title">Input Bond Information</h2>
    <v-text-field
      label="Purchase date (yyyy-mm)"
      variant="outlined"
      class="input"
      v-model="purchaseDate"
    ></v-text-field>
    <v-select
      :items="['5y', '10y', '30y']"
      label="Maturity"
      variant="outlined"
      v-model="maturity"
    ></v-select>

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
import {useStore} from 'vuex'
const store = useStore()
const emit = defineEmits(["close"]);
const showConfirmed = ref(false);
const purchaseDate = ref();
const maturity = ref();
async function submitSecurity() {
  console.log(purchaseDate.value, maturity.value);
  const userdata = await axios.get(
    `${import.meta.env.VITE_API_URL}/api/user`,
    {
      headers: {
        "Content-Type": "application/json",
        Authorization:
        `Bearer ${store.state.token}`,
      },
    }
  );
  console.log(userdata);
  let user_id = userdata.data.data.id;
  console.log(user_id);
  const response = await axios.post(
    `${import.meta.env.VITE_API_URL}/api/bond/`,
    {
      purchase_date: purchaseDate.value + "-01",
      maturity: maturity.value,
      user_id: user_id,
    },
    {
      headers: {
        "Content-Type": "application/json",
        Authorization:
        `Bearer ${store.state.token}`,
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
