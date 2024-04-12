<template>
  <Skeleton class="text-gray-200" v-if="!data" />
  <div class="artboard" v-else>
    <header>
      <h1>Settings</h1>
    </header>
    <Tabs :tabs="secondMenu('setting')" />

    <div class="artboard-content p-5">
      <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
        <div class="flex flex-row">
          <FormInput name="Sender Name" v-model="data.sender_name" :error="errors.name" id="name" type="text" rules="required|min:5|max:128" class="mr-5" />
          <FormInput name="Sender Email" v-model="data.sender_email" :error="errors.email" id="email" type="email" rules="required|email" />
        </div>

        <hr />
        <h2>SMTP settings</h2>
        <div class="flex flex-row">
          <FormInput name="Host" v-model="data.smtp.host" :error="errors.host" id="host" type="text" rules="required" class="mr-5" />
          <FormInput name="Port" v-model="data.smtp.port" :error="errors.port" id="port" type="text" rules="required|numeric" class="w-20 mr-5" />
          <FormSelect name="Encryption" v-model="data.smtp.encryption" :zeroOption="false" :options="['None', 'SSL/TLS', 'STARTTLS']" :error="errors.encryption" rules="required" id="encryption" />
        </div>

        <div class="flex flex-row">
          <FormInput name="Username" v-model="data.smtp.username" :error="errors.username" id="username" type="text" rules="required|min:3|max:20" class="max-w-md mr-5" />
          <FormInput name="Password" v-model="data.smtp.password" :error="errors.password" id="password" type="text" rules="required|min:3|max:20" class="max-w-md" />
        </div>

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
import { Tabs, FormInput, FormSelect, Skeleton } from "@/components";
import { secondMenu } from "@/utils/menu";
import { Form } from "vee-validate";
import { apiGet, apiUpdate } from "@/utils/api";

const loadingStatus = ref(false);
const data = ref();

onMounted(async () => {
  try {
    const res = await apiGet(`/_/api/setting/mail`, {});
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error('Error fetching mail data:', error);
  }
});

const onSubmit = async () => {
  loadingStatus.value = true;
  try {
    const updatedData = { ...data.value };
    updatedData.smtp.port = Number(updatedData.smtp.port);
    updatedData.smtp.encryption = Number(updatedData.smtp.encryption);

    const res = await apiUpdate(`/_/api/setting/mail`, {}, data.value);
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