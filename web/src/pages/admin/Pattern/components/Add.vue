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
      <div class="flex" v-for="(data, index) in limitData" :key="index">
        <div class="grow pr-3">
          <FormInput title="Key" v-model="data.key" :error="errors[`limit-key-${index}`]" :id="`limit-key-${index}`" rules="required|alpha_num" type="text" />
        </div>
        <div class="grow">
          <FormInput title="Value" v-model="data.value" :error="errors[`limit-value-${index}`]" :id="`limit-value-${index}`" rules="required|numeric" type="text" />
        </div>
        <div class="flex-none cursor-pointer pl-3 pt-4" @click="deleteLimitRecord(data.key)">
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
      <FormSelect name="Term" v-model="patternData.term" :options="termList" id="term" :error="errors.term" class="mr-5 w-10 flex-grow" />
      <div class="grow"></div>
      <FormInput name="Price" v-model="priceData" :error="errors.price" id="price" type="text" rules="required" title="Price" class="mr-5 w-10 flex-grow" />
      <FormSelect name="Currency" v-model="patternData.currency" :options="currencyList" id="currency" :error="errors.currency" class="w-6 flex-grow" />
    </div>

    <hr />
    <div>
      <label class="label">
        <span>Check strategies</span>
      </label>

      <div class="space-y-4" v-for="(data, index) in checkData" :key="index">
        <FormToggle :name="`Check the client's ${data.key}`" v-model="data.value" :id="data.key" />
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
import { ref, inject } from "vue";
import { useRoute } from "vue-router";
import { SvgIcon, FormInput, FormSelect, FormToggle } from "@/components";
import { Form } from "vee-validate";
import { term, currency, costFormat, costStripe } from "@/utils";
import { apiPost } from "@/utils/api";

const route = useRoute();

const getPatterns = inject("getPatterns") as Function;
const closeDrawer = inject("closeDrawer") as Function;

const loadingStatus = ref({
  save: false,
  delete: false,
});

const patternData = ref({
  name: "",
  limit: [],
  term: 0,
  price: 0,
  currency: 0,
  check: { "country": false, "ip": false, "mac": false },
  status: false,
  private: false,
});

const limitData = ref([]);

const termList = term.reduce((accumulator, currentValue, index) => {
  accumulator[index + 1] = currentValue;
  return accumulator;
}, {});

const currencyList = currency.reduce((accumulator, currentValue, index) => {
  accumulator[index + 1] = currentValue;
  return accumulator;
}, {});

const priceData = ref();
priceData.value = costFormat(patternData.value.price);

const checkData = ref([]);
checkData.value = Object.entries(patternData.value.check).map(([key, value]) => ({ key, value }));

const onSubmit = async () => {
  loadingStatus.value.save = true;

  try {
    patternData.value.limit = limitData.value.reduce((obj, item) => {
      obj[item.key] = Number(item.value);
      return obj;
    }, {});

    patternData.value.term = Number(patternData.value.term);
    patternData.value.price = costStripe(priceData.value);
    patternData.value.currency = Number(patternData.value.currency)

    patternData.value.check = checkData.value.reduce((obj, item) => {
      obj[item.key] = Boolean(item.value);
      return obj;
    }, {});

    const res = await apiPost(`/_/api/pattern`, {}, patternData.value);
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
  limitData.value.push({ key: "", value: 0 });
};

const deleteLimitRecord = (key: string) => {
  const index = limitData.value.findIndex(item => item.key === key);
  if (index !== -1) {
    limitData.value.splice(index, 1);
  }
};
</script>