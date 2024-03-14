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

onMounted(() => {
  const token = useRoute().query.token;
  const router = useRouter();
  if (token) {
    apiPost(`/api/sign/in?token=${token}`, data.value).then(res => {
      if (res.code === 200) {
        router.push({ name: 'manager-license' })
      } else {
        //showMessage(res.result, "connextError");
        console.log("connextError")
      }
    });
  }
});

const onSubmit = async () => {
  loadingStatus.value = true;
  apiPost(`/api/sign/in`, data.value).then(res => {
    if (res.code === 200) {
      sendMessage.value = true;
    } else {
      //showMessage(res.result, "connextError");
      console.log("connextError")
    }
  });
  loadingStatus.value = false;
};

</script>