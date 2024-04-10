<template>
  <div class="artboard">
    <header>
      <h1>Payments</h1>
      <label class="plus" @click="openDrawerAdd()">
        <SvgIcon name="plus-square" />
        add new
      </label>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th class="w-52">Customer</th>
          <th>Pattern</th>
          <th class="w-24">Term</th>
          <th class="w-24">Price</th>
          <th class="w-24">Provider</th>
          <th class="w-28">Status</th>
          <th class="w-8"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.payments" :key="index" class="cursor">
          <td :class="{ 'text-red-500': !item.customer.status }">{{ item.customer.email }}</td>
          <td @click="openDrawerView(item.id)">{{ item.pattern.name }}</td>
          <td @click="openDrawerView(item.id)">
            <Badge :name="termObj[item.pattern.term - 1].name" :color="termObj[item.pattern.term - 1].color" />
          </td>
          <td @click="openDrawerView(item.id)">{{ priceFormat(item.pattern.price) }} {{ currencyObj[item.pattern.currency - 1].name }}</td>
          <td @click="openDrawerView(item.id)">
            <Badge :name="paymentProvidersObj[item.transaction.provider - 1].name" :color="paymentProvidersObj[item.transaction.provider - 1].color" />
          </td>
          <td @click="openDrawerView(item.id)">
            <Badge :name="paymentStatusObj[item.transaction.status - 1].name" :color="paymentStatusObj[item.transaction.status - 1].color" />
          </td>
          <td>
            <div class="flex">
              <div>
                <SvgIcon name="pencil-square" @click="openDrawerEdit(item.id)" />
              </div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="desc">Empty</div>

    <div class="artboard-content">
      <Pagination :total="data.total" @selectPage="onSelectPage" />
    </div>
  </div>

  <Drawer :is-open="isDrawer.open" @close="closeDrawer()" maxWidth="600px">
    <View :drawer="isDrawer" v-if="isDrawer.action === 'view'" />
    <Edit :drawer="isDrawer" v-if="isDrawer.action === 'edit'" />
    <Add v-if="isDrawer.action === 'add'" />
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref, provide } from "vue";
import { useRoute } from "vue-router";
import { View, Edit, Add } from "./components";
import { SvgIcon, Badge, Pagination, Drawer } from "@/components";
import { termObj, paymentStatusObj, paymentProvidersObj, priceFormat, currencyObj } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  data: {},
  open: false,
  action: null,
})
const data: any = ref({});
const route = useRoute();

onMounted(() => {
  getPayments(route.query);
  if (route.params.customer_slug) {
    openDrawerView(<string>route.params.customer_slug)
  }
});

const getPayments = async (routeQuery: any) => {
  try {
    const res = await apiGet(`/_/api/payment`, routeQuery);
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error('Error fetching payment data:', error);
  }
};

const onSelectPage = (e: any) => {
  getPayments(e)
};

const getPayment = async (id: string, action: string) => {
  try {
    const res = await apiGet(`/_/api/payment/${id}`, {});
    if (res.code === 200) {
      isDrawer.value.data = res.result;
      isDrawer.value.open = true;
      isDrawer.value.action = action;
    }
  } catch (error) {
    console.error('Error fetching payment data:', error);
  }
};

const openDrawerView = async (id: string) => {
  getPayment(id, "view")
};

const openDrawerEdit = async (id: string) => {
  getPayment(id, "edit")
};

const openDrawerAdd = async () => {
  isDrawer.value.open = true;
  isDrawer.value.action = "add";
};

const closeDrawer = async () => {
  isDrawer.value.data = {};
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};


provide("getPayments", getPayments);
provide('closeDrawer', closeDrawer);
</script>