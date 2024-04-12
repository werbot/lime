<template>
  <Skeleton class="text-gray-200" v-if="!data" />
  <div class="artboard" v-else>
    <header>
      <h1>Settings</h1>
    </header>
    <Tabs :tabs="secondMenu('setting')" />

    <div class="artboard-content p-5">
      <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
        <FormInput name="Name" v-model="data.name" :error="errors.name" id="name" type="text" rules="required|min:5|max:128" class="max-w-md" />
        <FormInput name="Domain" v-model="data.domain" :error="errors.domain" id="domain" type="text" rules="required|url" class="max-w-md" />
        <FormInput name="Signature" v-model="data.signature" :error="errors.sign" id="sign" type="text" rules="required|min:5|max:128" class="max-w-md" />
        <FormInput name="Support email" v-model="data.email_support" :error="errors.email" id="email" type="email" rules="required|email" class="max-w-md" />
        <div>
          <button type="submit" class="btn mt-4 max-w-md" :disabled="loadingStatus">
            <div v-if="loadingStatus">
              <span>Loading...</span>
            </div>
            <span v-else>Save</span>
          </button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Tabs, FormInput, Skeleton } from "@/components";
import { secondMenu } from "@/utils/menu";
import { Form } from "vee-validate";
import { apiGet, apiUpdate } from "@/utils/api";

const loadingStatus = ref(false);
const data = ref();

onMounted(async () => {
  try {
    const res = await apiGet(`/_/api/setting/site`, {});
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error('Error fetching setting data:', error);
  }
});

const onSubmit = async () => {
  loadingStatus.value = true;
  try {
    const res = await apiUpdate(`/_/api/setting/site`, {}, data.value);
    if (res.code === 200) {
      //getPayments(route.query);
      //closeDrawer();
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }
  loadingStatus.value = false;
};
</script>