<template>
  <div class="title">License Manager Access</div>
  <Form @submit="onSubmit" v-slot="{ errors }" v-if="!sendMessage">
    <FormInput name="Email" v-model="data.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" ico="at-symbol" />
    <div class="form-control mt-6">
      <button type="submit" class="btn " :disabled="loadingStatus">
        <div v-if="loadingStatus">
          <span>Loading...</span>
        </div>
        <span v-else>Request Access Link</span>
      </button>
    </div>
  </Form>
  <div v-else>We sent you a link to access the License Manager. The link expires after 24 hours.</div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { apiPost } from "@/utils/api";
import { FormInput } from "@/components";
import { Form } from "vee-validate";

const loadingStatus = ref<boolean>(false);
const data = ref({ email: "" });
const sendMessage = ref<boolean>(false);

async function signIn(token = {}) {
  try {
    const response = await apiPost(`/api/sign/in`, token, data.value);
    if (response.code === 200) {
      return true;
    } else {
      console.error("Connection error:", response.result);
      return false;
    }
  } catch (error) {
    console.error("API call failed:", error);
    return false;
  }
}

onMounted(async () => {
  const route = useRoute();
  const router = useRouter();
  const token = route.query.token;

  if (token) {
    const success = await signIn({ "token": token });
    if (success) {
      router.push({ name: 'manager-license' });
    }
  }
});

const onSubmit = async () => {
  loadingStatus.value = true;
  const success = await signIn();
  if (success) {
    sendMessage.value = true;
  }
  loadingStatus.value = false;
};

</script>