<template>
  <Skeleton class="text-gray-200" v-if="!data" />
  <div v-else>
    <header>
      <h1>Add new customer</h1>
    </header>

    <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
      <FormInput name="Email" v-model="data.email" :error="errors.email" id="email" type="email" rules="required|email" title="Email" />

      <div class="mt-5 flex flex-row">
        <FormToggle name="Status" v-model="data.status" class="mr-5 flex-grow" id="status" />
      </div>

      <div class="pt-8">
        <div class="flex">
          <div class="flex-none">
            <button type="submit" class="btn btn-green" :disabled="loadingStatus.save">
              <div v-if="loadingStatus.save">
                <span>Loading...</span>
              </div>
              <span v-else>Create</span>
            </button>
            <div class="btn ml-5 cursor-pointer" @click="closeDrawer()">Close</div>
          </div>
          <div class="grow"></div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { ref, inject } from 'vue';
import { useRoute } from "vue-router";
import { Skeleton, FormInput, FormToggle } from "@/components";
import { Form } from "vee-validate";
import { apiPost } from "@/utils/api";

const route = useRoute();
const getCustomers = inject("getCustomers") as Function;
const closeDrawer = inject('closeDrawer') as Function;

const loadingStatus = ref({
  save: false,
});

const data = ref({
  email: null,
  status: false,
});

const onSubmit = async () => {
  loadingStatus.value.save = true;
  try {
    const res = await apiPost(`/_/api/customer`, {}, data.value);
    if (res.code === 200) {
      getCustomers(route.query);
      closeDrawer();
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }
  loadingStatus.value.save = false;
};
</script>