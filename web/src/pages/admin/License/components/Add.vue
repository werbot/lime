<template>
  <Skeleton class="text-gray-200" v-if="!data.lists.payments" />
  <div v-else>
    <header>
      <h1>Add new license</h1>
    </header>

    <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
      <FormSelect name="Payment" v-model="licenseData.payment.id" :options="data.lists.payments" id="payment" :error="errors.payment" :required="true"
        @change="onSelectPayment()" />
    </Form>

    <div v-if="licenseData.payment.id" class="pt-4">
      <Skeleton class="text-gray-200" v-if="!licenseData.pattern" />
      <div v-else>
        <div class="rounded border border-solid border-gray-300">
          <table class="mini">
            <tr>
              <td class="w-32">Name</td>
              <td>{{ licenseData.pattern.name }}</td>
            </tr>
            <tr>
              <td>Limits</td>
              <td class="!p-0">
                <table class="mini">
                  <tr v-for="(value, key, index) in licenseData.pattern.limit" :key="index">
                    <td>{{ key }}</td>
                    <td>
                      <Badge :name="String(value)" />
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
          </table>
        </div>

        <div v-if="licenseData.pattern.check.country" class="pt-8">
          <div class="flex">
            <div><label class="label"><span>Country</span></label></div>
            <div class="grow"></div>
            <div class="flex-none">
              <FormToggle name="Black-list" l-color="green" r-color="red" v-model="data.check.country.black" />
            </div>
          </div>
          <div class="flex">
            <div class="grow">
              <FormInput v-model="tmp.country.req" :id="`country-value`" type="text" @keyup="searchCountries()" />
            </div>
          </div>
          <div class="flex-col">
            <div v-if="tmp.country.search?.length > 0" class="mt-3">
              <Badge v-for="(item, index) in tmp.country.search" :key="index" :name="item.value" class="mr-1 cursor-pointer" @click="addCountry(index)" />
            </div>
            <div class="pt-4 pb-2">
              <span class="mr-3 mb-3 inline-flex items-center rounded border px-2 py-1" v-for="(item, index) in data.check.country.list"
                :class="data.check.country.black ? 'bg-red-100 border-red-300' : 'bg-green-100  border-green-300'">
                <span>{{ item.value }}</span>
                <SvgIcon name="trash" class="ml-1 h-4 w-4 opacity-75 cursor-pointer" @click="removeCountry(index)" />
              </span>
            </div>
          </div>
          <hr />
        </div>

        <div v-if="licenseData.pattern.check.ip" class="pt-8">
          <div class="flex">
            <div><label class="label"><span>IP address</span></label></div>
            <div class="grow"></div>
            <div class="flex-none">
              <FormToggle name="Black-list" l-color="green" r-color="red" v-model="data.check.ip.black" />
            </div>
          </div>
          <div class="flex">
            <div class="grow">
              <FormInput v-model="tmp.ip.req" id="ip-value" type="text" @keyup="searchIP()" />
            </div>
          </div>
          <div class="flex-col">
            <div v-if="tmp.ip.search" class="mt-3">
              <Badge :name="(tmp.ip.search.start === tmp.ip.search.end ? tmp.ip.search.start : `${tmp.ip.search.start}-${tmp.ip.search.end}`)" class="mr-1 cursor-pointer"
                @click="addIP()" />
            </div>
            <div class="pt-4 pb-2">
              <span class="mr-3 mb-3 inline-flex items-center rounded border px-2 py-1" v-for="(item, index) in data.check.ip.list"
                :class="data.check.ip.black ? 'bg-red-100 border-red-300' : 'bg-green-100  border-green-300'">
                <span>{{ (item.start === item.end ? item.start : `${item.start}-${item.end}`) }}</span>
                <SvgIcon name="trash" class="ml-1 h-4 w-4 opacity-75 cursor-pointer" @click="removeIP(index)" />
              </span>
            </div>
          </div>
          <hr />
        </div>

        <div v-if="licenseData.pattern.check.mac" class="pt-8">
          <div class="flex">
            <div><label class="label"><span>MAC address</span></label></div>
            <div class="grow"></div>
            <div class="flex-none">
              <FormToggle name="Black-list" l-color="green" r-color="red" v-model="data.check.mac.black" />
            </div>
          </div>
          <div class="flex">
            <div class="grow">
              <FormInput v-model="tmp.mac.req" id="mac-value" type="text" @keyup="searchMAC()" />
            </div>
          </div>
          <div class="flex-col">
            <div v-if="tmp.mac.search" class="mt-3">
              <Badge :name="tmp.mac.search" class="mr-1 cursor-pointer" @click="addMAC()" />
            </div>
            <div class="pt-4 pb-2">
              <span class="mr-3 mb-3 inline-flex items-center rounded border px-2 py-1" v-for="(item, index) in data.check.mac.list"
                :class="data.check.mac.black ? 'bg-red-100 border-red-300' : 'bg-green-100  border-green-300'">
                <span>{{ item }}</span>
                <SvgIcon name="trash" class="ml-1 h-4 w-4 opacity-75 cursor-pointer" @click="removeMAC(index)" />
              </span>
            </div>
          </div>
          <hr />
        </div>

      </div>
    </div>

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <div class="btn btn-green cursor-pointer disabled" v-if="!licenseData.payment.id">
            <span>Create</span>
          </div>
          <div class="btn btn-green cursor-pointer" :class="{ 'disabled': loadingStatus }" @click="onSubmit()" v-else>
            <span v-if="loadingStatus">Loading...</span>
            <span v-else>Create</span>
          </div>
          <div class="btn ml-5 cursor-pointer" @click="closeDrawer()">Close</div>
        </div>
        <div class="grow"></div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, inject, onMounted } from 'vue';
import { useRoute } from "vue-router";
import { Skeleton, SvgIcon, FormToggle, FormInput, FormSelect, Badge } from "@/components";
import { Form } from "vee-validate";
import { apiGet, apiPost } from "@/utils/api";
import { Address4, Address6 } from "ip-address";

const route = useRoute();
const getLicenses = inject("getLicenses") as Function;
const closeDrawer = inject('closeDrawer') as Function;

const licenseData = ref({
  payment: {
    id: null,
  },
  pattern: null,
});

const loadingStatus = ref(false);

const data = ref({
  lists: {
    payments: {},
  },

  check: {
    country: {
      black: false,
      list: [],
    },
    ip: {
      black: false,
      list: [],
    },
    mac: {
      black: false,
      list: [],
    }
  },
});

const tmp = ref({
  payments: [],
  country: {
    req: null,
    search: null,
  },
  ip: {
    req: null,
    search: null,
  },
  mac: {
    req: null,
    search: null,
  },
})

onMounted(async () => {
  loadingStatus.value = true;
  try {
    const res = await apiGet(`/_/api/payment/no_lic`, {});
    if (res.code === 200) {
      tmp.value.payments = res.result.payments;

      data.value.lists.payments = res.result.payments.reduce((accumulator, item) => {
        accumulator[item.id] = item.customer.email;
        return accumulator;
      }, {});
    }
  } catch (error) {
    console.error('Error fetching payment data:', error);
  }
  loadingStatus.value = false;
});

const onSubmit = async () => {
  loadingStatus.value = true;
  try {
    const dataCheck = { ...data.value.check };
    const addData = {
      id: licenseData.value.payment.id,
      pattern: {
        check: {
          country: {
            black: dataCheck.country.black,
            list: dataCheck.country.list.map(country => country.key)
          },
          ip: {
            black: dataCheck.ip.black,
            list: dataCheck.ip.list,
          },
          mac: {
            black: dataCheck.mac.black,
            list: dataCheck.mac.list,
          },
        },
      },
    }

    const res = await apiPost(`/_/api/license`, {}, addData);
    if (res.code === 200) {
      getLicenses(route.query);
      closeDrawer();
    }
  } catch (error) {
    console.error("Error fetching license data:", error);
  }
  loadingStatus.value = false;
};

const onSelectPayment = async () => {
  licenseData.value.pattern = null
  const patternID = tmp.value.payments.find(payment => payment.id === licenseData.value.payment.id)?.pattern.id;

  if (patternID) {
    try {
      const res = await apiGet(`/_/api/pattern/${patternID}`, {});
      if (res.code === 200) {
        licenseData.value.pattern = res.result
      }
    } catch (error) {
      console.error('Error fetching pattern data:', error);
    }
  }
}

const searchCountries = async () => {
  const country = tmp.value.country.req
  if (country.length > 2) {
    try {
      const res = await apiGet(`/_/api/list/countries/${country}`, {});
      if (res.code === 200) {
        tmp.value.country.search = res.result;
        if (tmp.value.country.search.length > 0) {
          Object.entries(data.value.check.country.list).forEach((n) => {
            Object.entries(tmp.value.country.search).forEach((s) => {
              if (s[1]["key"] === n[1]["key"]) {
                tmp.value.country.search.splice(s[0], 1);
              }
            });
          });
        }
      }
    } catch (error) {
      console.error('Error fetching country data:', error);
    }
  } else {
    tmp.value.country.search = null;
  }
};

const addCountry = async (index) => {
  data.value.check.country.list.push(tmp.value.country.search[index]);
  tmp.value.country.search.splice(index, 1);
};

const removeCountry = async (index) => {
  data.value.check.country.list.splice(index, 1);
  searchCountries();
};

const searchIP = async () => {
  const ipData = tmp.value.ip;
  const ip = ipData.req;
  ipData.search = null;

  if (ip.length > 2) {
    let address: Address4 | Address6
    if (Address4.isValid(ip)) {
      address = new Address4(ip);
    } else if (Address6.isValid(ip)) {
      address = new Address6(ip);
    }

    if (address) {
      const start = address.startAddress().addressMinusSuffix
      const end = address.endAddress().addressMinusSuffix
      const isRangeInList = Object.values(data.value.check.ip.list).some(range =>
        range.start === start && range.end === end
      );

      if (!isRangeInList) {
        ipData.search = { start, end };
      }
    }
  }
}

const addIP = async () => {
  data.value.check.ip.list.push(tmp.value.ip.search);
  tmp.value.ip.search = null;
};

const removeIP = async (index) => {
  data.value.check.ip.list.splice(index, 1);
  searchIP();
};

const searchMAC = async () => {
  const regex = /^([0-9A-F]{2}[:-]){5}([0-9A-F]{2})$/;
  const macData = tmp.value.mac;
  const mac = macData.req;
  macData.search = null;

  if (mac.length > 11) {
    if (regex.test(mac)) {
      const isRangeInList = Object.values(data.value.check.mac.list).some(range =>
        range === mac
      );

      if (!isRangeInList) {
        macData.search = mac;
      }
    }
  }
}

const addMAC = async () => {
  data.value.check.mac.list.push(tmp.value.mac.search);
  tmp.value.mac.search = null;
};

const removeMAC = async (index) => {
  data.value.check.mac.list.splice(index, 1);
  searchMAC();
};
</script>
