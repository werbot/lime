<template>
  <Form @submit="onSubmit" v-slot="{ errors }" class="mx-auto mb-0 max-w-md space-y-4">
    <FormInput name="Email" v-model="data.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" ico="at-symbol" />
    <div class="form-control mt-6">
      <button type="submit" class="btn" :disabled="loading">
        <div v-if="loading">
          <span>Loading...</span>
        </div>
        <span v-else>Login</span>
      </button>
    </div>
  </Form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { apiPost } from "@/utils/api";
import { FormInput } from "@/components";
import { Form } from "vee-validate";

const loading = ref(false);
const data = ref({ email: "" });

const onSubmit = async () => {
  loading.value = true;

  apiPost(`/api/sign/in`, data.value).then(res => {
    if (res.success) {
      window.location.href = "/";
    } else {
      //showMessage(res.result, "connextError");
      console.log("connextError")
    }
  });

  loading.value = false;
};

</script>