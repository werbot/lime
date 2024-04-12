<template>
  <Skeleton class="text-gray-200" v-if="!drawer.data" />
  <div v-else>
    <header>
      <h1>Edit <span class="text-green-600">{{ drawer.data.email }}</span> customer</h1>
    </header>

    <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
      <FormInput name="Email" v-model="drawer.data.email" :error="errors.email" id="email" type="email" rules="required|email" />

      <div class="mt-5 flex flex-row">
        <FormToggle name="Status" v-model="drawer.data.status" class="mr-5 flex-grow" id="status" />
      </div>

      <div class="pt-8">
        <div class="flex">
          <div class="flex-none">
            <button type="submit" class="btn btn-green" :disabled="loadingStatus.save">
              <div v-if="loadingStatus.save">
                <span>Loading...</span>
              </div>
              <span v-else>Save</span>
            </button>
            <div class="btn ml-5 cursor-pointer" @click="closeDrawer()">Close</div>
          </div>
          <div class="grow"></div>
          <div class="flex-none">
            <div class="btn btn-red cursor-pointer disabled" v-if="drawer.data.payments.total > 0">
              <span>Delete</span>
            </div>
            <div class="btn btn-red cursor-pointer" :class="{ 'disabled': loadingStatus.delete }" @click="onDelete()" v-else>
              <span v-if="loadingStatus.delete">Loading...</span>
              <span v-else>Delete</span>
            </div>
          </div>
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
import { apiUpdate, apiDelete } from "@/utils/api";

const route = useRoute();
const getCustomers = inject("getCustomers") as Function;
const closeDrawer = inject('closeDrawer') as Function;

const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});

const loadingStatus = ref({
  save: false,
  delete: false,
});

const onSubmit = async () => {
  loadingStatus.value.save = true;
  try {
    const updatedData = { ...props.drawer.data };

    const res = await apiUpdate(`/_/api/customer/${updatedData.id}`, {}, updatedData);
    if (res.code === 200) {
      getCustomers(route.query);
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }
  loadingStatus.value.save = false;
};

const onDelete = async () => {
  loadingStatus.value.delete = true;
  try {
    const res = await apiDelete(`/_/api/customer/${props.drawer.data.id}`, {});
    if (res.code === 200) {
      console.log(res);
      getCustomers(route.query);
      closeDrawer();
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }
  loadingStatus.value.delete = false;
};
</script>