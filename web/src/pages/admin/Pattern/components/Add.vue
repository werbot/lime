<template>
  <header>
    <h1>Add new pattern</h1>
  </header>

  <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
    <FormInput name="Name" v-model="patternData.name" :error="errors.name" id="name" type="text" rules="required" title="Name" />

    <hr />
    <div>
      <label class="label">
        <span>Limit</span>
      </label>
      <div class="flex text-gray-400 text-sm">
        <span class="grow">Name</span>
        <span class="grow -ml-5">Value</span>
      </div>
      <div class="flex" v-for="(limit, index) in data.limit" :key="index">
        <div class="grow pr-3">
          <FormInput title="Key" v-model="limit.key" :error="errors[`limit-key-${index}`]" :id="`limit-key-${index}`" rules="required|alpha_num" type="text" />
        </div>
        <div class="grow">
          <FormInput title="Value" v-model="limit.value" :error="errors[`limit-value-${index}`]" :id="`limit-value-${index}`" rules="required|numeric" type="text" />
        </div>
        <div class="flex-none cursor-pointer pl-3 pt-4" @click="deleteLimitRecord(limit.key)">
          <SvgIcon name="trash" class="h-5 w-5 text-red-500" stroke="currentColor" />
        </div>
      </div>
      <div class="flex py-4">
        <div class="grow"></div>
        <div class="mt-2 flex-none">
          <a href="#" class="shrink-0 rounded-lg bg-gray-200 p-2 text-sm font-medium text-gray-700" @click="addLimitRecord()">
            Add limit record
          </a>
        </div>
      </div>
    </div>

    <hr />
    <div class="mt-5 flex flex-row">
      <FormSelect name="Term" v-model="patternData.term" :options="data.lists.terms" id="term" :error="errors.term" class="mr-5 w-10 flex-grow" />
      <div class="grow"></div>
      <FormInput name="Price" v-model="data.price" :error="errors.price" id="price" type="text" rules="required" title="Price" class="mr-5 w-10 flex-grow" />
      <FormSelect name="Currency" v-model="patternData.currency" :options="data.lists.currency" id="currency" :error="errors.currency" class="w-6 flex-grow" />
    </div>

    <hr />
    <div>
      <label class="label">
        <span>Check strategies</span>
      </label>

      <div class="space-y-4" v-for="(check, index) in data.check" :key="index">
        <FormToggle :name="`Check the client's ${check.key}`" v-model="check.value" :id="check.key" />
      </div>
    </div>

    <hr />
    <div class="mt-5 flex flex-row">
      <FormToggle name="Status" v-model="patternData.status" class="mr-5 flex-grow" id="status" />
      <FormToggle name="Private use" v-model="patternData.private" class="flex-grow" id="private" />
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
import { ref, inject, onMounted } from "vue";
import { useRoute } from "vue-router";
import { SvgIcon, FormInput, FormSelect, FormToggle } from "@/components";
import { Form } from "vee-validate";
import { termObj, currencyObj, arrayToObject, reduceToObject, costFormat, costStripe } from "@/utils";
import { apiPost } from "@/utils/api";

const route = useRoute();
const getPatterns = inject("getPatterns") as Function;
const closeDrawer = inject("closeDrawer") as Function;

const loadingStatus = ref({
  save: false,
});

const data = ref({
  lists: {
    terms: {},
    currency: {},
  },
  limit: [],
  check: [],
  price: null,
});

const patternData = ref({
  name: null,
  limit: [],
  term: null,
  price: null,
  currency: null,
  check: { country: false, ip: false, mac: false },
  status: false,
  private: false,
});

onMounted(async () => {
  loadingStatus.value.save = true;

  data.value = {
    ...data.value,
    lists: {
      currency: arrayToObject(currencyObj, currency => currency.name),
      terms: arrayToObject(termObj, terms => terms.name),
    },
    check: Object.entries(patternData.value.check ?? []).map(([key, value]) => ({ key, value })),
    price: costFormat(patternData.value.price),
  }

  loadingStatus.value.save = false;
});

const onSubmit = async () => {
  loadingStatus.value.save = true;

  try {
    const addData = { ...patternData.value };
    addData.limit = reduceToObject(data.value.limit, Number);
    addData.check = reduceToObject(data.value.check, Boolean);
    addData.term = Number(addData.term);
    addData.price = costStripe(data.value.price);
    addData.currency = Number(addData.currency);

    const res = await apiPost(`/_/api/pattern`, {}, addData);
    if (res.code === 200) {
      getPatterns(route.query);
      closeDrawer();
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }

  loadingStatus.value.save = false;
};

const addLimitRecord = () => {
  data.value.limit.push({ key: null, value: null });
};

const deleteLimitRecord = (key: string) => {
  const index = data.value.limit.findIndex(item => item.key === key);
  if (index !== -1) {
    data.value.limit.splice(index, 1);
  }
};
</script>