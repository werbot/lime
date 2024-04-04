<template>
  <div class="artboard">
    <header>
      <h1>Settings</h1>
    </header>
    <Tabs :tabs="secondMenu('setting')" />

    <div class="artboard-content p-5">
      <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
        <FormInput name="Email" v-model="data.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" class="max-w-md" />
        <FormInput name="Password" v-model="data.password" :error="errors.password" id="password" type="password" rules="required" title="Password" class="max-w-md" />
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
import { ref } from "vue";
import { Tabs, FormInput } from "@/components";
import { secondMenu } from "@/utils/menu";
import { Form } from "vee-validate";

const loadingStatus = ref(false);
const data = ref({
  email: null,
  password: null,
});

const onSubmit = async () => {
  loadingStatus.value = true;

  loadingStatus.value = false;
};
</script>