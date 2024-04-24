<template>
  <header>
    <h1>Add new payment</h1>
  </header>

  <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
    <FormSelect name="Customer" v-model="paymentData.customer.id" :options="data.lists.customers" id="customer" :error="errors.customer" :required="true" />
    <FormSelect name="Pattern" v-model="paymentData.pattern.id" :options="data.lists.patterns" id="pattern" :error="errors.pattern" :required="true" />

    <hr />
    <div class="mt-5 flex flex-row">
      <FormSelect name="Provider" v-model="paymentData.transaction.provider" :options="data.lists.providers" id="provider" :error="errors.provider" :required="true"
        class="mr-5 w-10 flex-grow" />
      <FormSelect name="Status" v-model="paymentData.transaction.status" :options="data.lists.status" id="status" :error="errors.status" :required="true" class="w-10 flex-grow" />
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
          <FormInput v-model="meta.key" :error="errors[`meta-key-${index}`]" :id="`meta-key-${index}`" rules="required" type="text" />
        </div>
        <div class="grow">
          <FormInput v-model="meta.value" :error="errors[`meta-value-${index}`]" :id="`meta-value-${index}`" rules="required" type="text" />
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
            <span v-else>Create</span>
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
import { SvgIcon, FormSelect, FormInput } from "@/components";
import { Form } from "vee-validate";
import { arrayToObject, reduceToObject, paymentProvidersObj, paymentStatusObj } from "@/utils";
import { apiGet, apiPost } from "@/utils/api";

const route = useRoute();
const getPayments = inject("getPayments") as Function;
const closeDrawer = inject('closeDrawer') as Function;

const paymentData = ref({
  customer: {
    id: null,
  },
  pattern: {
    id: null,
  },
  transaction: {
    provider: null,
    status: null,
    meta: {},
  }
});

const loadingStatus = ref({
  save: false,
});

const data = ref({
  lists: {
    customers: {},
    patterns: {},
    providers: {},
    status: {},
  },
  meta: [],
});

onMounted(async () => {
  loadingStatus.value.save = true;

  data.value = {
    ...data.value,
    lists: {
      customers: {},
      patterns: {},
      providers: arrayToObject(paymentProvidersObj, provider => provider.name),
      status: arrayToObject(paymentStatusObj, status => status.name),
    }
  }

  await fetchListData('customers', 'email');
  await fetchListData('patterns', 'name');

  loadingStatus.value.save = false;
});

const onSubmit = async () => {
  loadingStatus.value.save = true;
  try {
    const addData = { ...paymentData.value };
    addData.transaction.provider = Number(addData.transaction.provider);
    addData.transaction.status = Number(addData.transaction.status);
    addData.transaction.meta = reduceToObject(data.value.meta, String);

    const res = await apiPost(`/_/api/payment`, {}, addData);
    if (res.code === 200) {
      getPayments(route.query);
      closeDrawer();
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }
  loadingStatus.value.save = false;
};

async function fetchListData(listType, propertyKey) {
  try {
    const res = await apiGet(`/_/api/list/${listType}`, {});
    if (res.code === 200) {
      data.value.lists[listType] = res.result[listType].reduce((accumulator, item) => {
        accumulator[item.id] = item[propertyKey];
        return accumulator;
      }, {});
    }
  } catch (error) {
    console.error(`Error fetching ${listType}:`, error);
  }
}

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