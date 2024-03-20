<template>
  <div class="artboard">
    <header>
      <h1>Payments</h1>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th class="w-52">Customer</th>
          <th>Pattern</th>
          <th class="w-24">Term</th>
          <th class="w-24">Price</th>
          <th class="w-24">Provider</th>
          <th class="w-32">Status</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.payments" :key="index" class="cursor" @click="openDrawerDesc(item.id)">
          <td :class="{ 'text-red-500': !item.customer.status }">{{ item.customer.email }}</td>
          <td>{{ item.pattern.name }}</td>
          <td>
            <Badge :name="termFormat[item.pattern.term].name" :color="termFormat[item.pattern.term].color" />
          </td>
          <td>{{ priceFormat(item.pattern.price) }} {{ currency[item.pattern.currency] }}</td>
          <td>{{ item.transaction.provider }}</td>
          <td>
            <Badge :name="paymentStatusFormat[item.transaction.status].name" :color="paymentStatusFormat[item.transaction.status].color" />
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="desc">Empty</div>

    <div class="artboard-content">
      <Pagination :total="data.total" @selectPage="onSelectPage" />
    </div>
  </div>

  <Drawer :is-open="isDrawer.open" @close="closeDrawer" maxWidth="600px">
    <div class="rounded border border-solid border-gray-300">
      <table class="mini" v-if="dataFull.id">
        <tr>
          <td class="w-32">ID</td>
          <td>{{ dataFull.id }}</td>
        </tr>
        <tr>
          <td>Customer</td>
          <td>
            <span class="dot mr-2" :class="Object(dataFull.customer).status ? 'bg-green-500' : 'bg-red-500'"></span>
            {{ Object(dataFull.customer).email }}
          </td>
        </tr>
        <tr>
          <td>Pattern</td>
          <td>{{ Object(dataFull.pattern).name }}</td>
        </tr>
        <tr>
          <td>Term</td>
          <td><Badge :name="termFormat[Object(dataFull.pattern).term].name" :color="termFormat[Object(dataFull.pattern).term].color" /></td>
        </tr>
        <tr>
          <td>Price</td>
          <td>{{ priceFormat(Object(dataFull.pattern).price) }} {{ currency[Object(dataFull.pattern).currency] }}</td>
        </tr>
        <tr>
          <td>Provider</td>
          <td>{{ Object(dataFull.transaction).provider }}</td>
        </tr>
        <tr>
          <td>Status</td>
          <td>
            <Badge :name="paymentStatusFormat[Object(dataFull.transaction).status].name" :color="paymentStatusFormat[Object(dataFull.transaction).status].color" />
          </td>
        </tr>
        <tr>
          <td>Created</td>
          <td>{{ formatDate(dataFull.created) }}</td>
        </tr>
        <tr v-if="dataFull.updated">
          <td>Updated</td>
          <td>{{ formatDate(dataFull.updated) }}</td>
        </tr>
      </table>
    </div>
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { SvgIcon, Badge, Pagination, Drawer } from "@/components";
import { termFormat, paymentStatusFormat, priceFormat, currency, formatDate } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  open: false,
})
const data: any = ref({});
const dataFull: any = ref({});
const route = useRoute();

onMounted(() => {
  getData(route.query);
  if (route.params.payment_slug) {
    openDrawerDesc(<string>route.params.payment_slug)
  }
});

const getData = async (routeQuery: any) => {
  apiGet(`/_/api/payment`, routeQuery).then(res => {
    if (res.code === 200) {
      data.value = res.result;
    }
  });
};

const onSelectPage = (e: any) => {
  getData(e);
};

const openDrawerDesc = async (id: string) => {
  apiGet(`/_/api/payment/${id}`, {}).then(res => {
    if (res.code === 200) {
      dataFull.value = res.result;
      isDrawer.value.open = true;
    }
  });
};

const closeDrawer = async () => {
  dataFull.value = {};
  isDrawer.value.open = false;
};
</script>