<template>
  <header>
    <h1>Edit <span class="text-green-600">{{ drawer.data.id }}</span> payment</h1>
  </header>

  <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
    <div class="mt-5 flex flex-row">
      <FormSelect name="Status" v-model="drawer.data.transaction.status" :options="data.lists.status" id="status" :error="errors.status" :required="true" class="w-10 flex-grow" />
    </div>

    <hr />
    <div>
      <label class="label">
        <span>Metadata</span>
      </label>
      <div class="flex text-gray-400 text-sm">
        <span class="grow">Name</span>
        <span class="grow -ml-5">Value</span>
      </div>
      <div class="flex" v-for="(meta, index) in data.meta" :key="index">
        <div class="grow pr-3">
          <FormInput title="Key" v-model="meta.key" :error="errors[`meta-key-${index}`]" :id="`meta-key-${index}`" rules="required" type="text" />
        </div>
        <div class="grow">
          <FormInput title="Value" v-model="meta.value" :error="errors[`meta-value-${index}`]" :id="`meta-value-${index}`" rules="required" type="text" />
        </div>
        <div class="flex-none cursor-pointer pl-3 pt-4" @click="deleteMetaRecord(meta.key)">
          <SvgIcon name="trash" class="h-5 w-5 text-red-500" stroke="currentColor" />
        </div>
      </div>
      <div class="flex py-4">
        <div class="grow"></div>
        <div class="mt-2 flex-none">
          <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addMetaRecord()">
            Add metadata record
          </a>
        </div>
      </div>
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
      </div>
    </div>
  </Form>
</template>

<script setup lang="ts">
import { ref, inject, onMounted } from 'vue';
import { useRoute } from "vue-router";
import { SvgIcon, FormInput, FormSelect } from "@/components";
import { Form } from "vee-validate";
import { arrayToObject, reduceToObject, paymentStatusObj } from "@/utils";
import { apiUpdate } from "@/utils/api";

const route = useRoute();
const getPayments = inject("getPayments") as Function;
const closeDrawer = inject('closeDrawer') as Function;

const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});

const loadingStatus = ref({
  save: false,
});

const data = ref({
  lists: {
    status: {},
  },
  meta: [],
});

onMounted(async () => {
  loadingStatus.value.save = true;

  data.value = {
    ...data.value,
    lists: {
      status: arrayToObject(paymentStatusObj, status => status.name),
    },
    meta: Object.entries(props.drawer.data.transaction.meta ?? []).map(([key, value]) => ({ key, value })),
  }

  loadingStatus.value.save = false;
});

const onSubmit = async () => {
  loadingStatus.value.save = true;
  try {
    const updatedData = { ...props.drawer.data };
    updatedData.transaction.status = Number(updatedData.transaction.status);
    updatedData.transaction.meta = reduceToObject(data.value.meta, String);

    const res = await apiUpdate(`/_/api/payment/${updatedData.id}`, {}, updatedData);
    if (res.code === 200) {
      getPayments(route.query);
      //closeDrawer();
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }
  loadingStatus.value.save = false;
};

const addMetaRecord = () => {
  data.value.meta.push({ key: null, value: null });
};

const deleteMetaRecord = (key: string) => {
  const index = data.value.meta.findIndex(item => item.key === key);
  if (index !== -1) {
    data.value.meta.splice(index, 1);
  }
};
</script>