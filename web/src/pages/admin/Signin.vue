<template>
  <div class="title">Admin SingIn</div>
  <Form @submit="onSubmit" v-slot="{ errors }" class="mx-auto mb-0 max-w-md space-y-4">
    <FormInput name="Email" v-model="data.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" ico="at-symbol" />
    <FormInput name="Password" v-model="data.password" :error="errors.password" id="password" type="password" rules="required" title="Password" ico="at-symbol" />
    <div class="form-control">
      <button type="submit" class="btn mt-4" :disabled="loadingStatus">
        <div v-if="loadingStatus">
          <span>Loading...</span>
        </div>
        <span v-else>Login</span>
      </button>
    </div>
  </Form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { apiPost } from "@/utils/api";
import { FormInput } from "@/components";
import { Form } from "vee-validate";

const loadingStatus = ref(false);
const data = ref({ email: "", password: "" });
const router = useRouter();

const onSubmit = async () => {
  loadingStatus.value = true;

  apiPost(`/_/api/sign/in`, data.value).then(res => {
    if (res.code === 200) {
      router.push({ name: 'admin-license' })
    } else {
      //showMessage(res.result, "connextError");
      console.log("connextError")
    }
  });

  loadingStatus.value = false;
};

</script>