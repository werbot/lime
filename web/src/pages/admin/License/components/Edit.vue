<template>
  <header>
    <h1>Edit license</h1>
  </header>

  {{ drawer.data }}

  <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
    <FormInput name="Email" v-model="data.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" />
    <FormInput name="Password" v-model="data.password" :error="errors.password" id="password" type="password" rules="required" title="Password" />
  </Form>

  <div class="pt-8">
    <div class="flex">
      <div class="flex-none">
        <button type="submit" class="btn btn-green" :disabled="loadingStatus">
          <div v-if="loadingStatus">
            <span>Loading...</span>
          </div>
          <span v-else>Save</span>
        </button>
        <button class="btn ml-5" @click="closeDrawer()">Close</button>
      </div>
      <div class="grow"></div>
      <div class="flex-none">
        <button class="btn btn-red" @click="closeDrawer()">Delete pattern</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, inject } from 'vue';
import { FormInput } from "@/components";
import { Form } from "vee-validate";

const closeDrawer = inject('closeDrawer') as Function;
const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});

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