<template>
  <div class="artboard">
    <header>
      <h1>Customers</h1>
      <label class="plus" @click="openDrawerAdd()">
        <SvgIcon name="plus-square" />
        add new
      </label>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th class="w-12"></th>
          <th>Email</th>
          <th class="w-24">Payments</th>
          <th class="w-48">Created</th>
          <th class="w-8"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.customers" :key="index" class="cursor" :class="{ 'bg-red-50': !item.status }">
          <td @click="openDrawerView(item.id)">
            <div class="flex items-center">
              <span class="dot" :class="item.status ? 'bg-green-500' : 'bg-red-500'"></span>
            </div>
          </td>
          <td @click="openDrawerView(item.id)" :class="{ 'text-red-500': !item.status }">{{ item.email }}</td>
          <td @click="openDrawerView(item.id)">
            <Badge :name="String(item.payments.total)" />
          </td>
          <td @click="openDrawerView(item.id)">{{ formatDate(item.created) }}</td>
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
import { SvgIcon, Pagination, Drawer, Badge } from "@/components";
import { formatDate } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  data: {},
  open: false,
  action: null,
})
const data: any = ref({});
const route = useRoute();

onMounted(() => {
  getCustomers(route.query);
  if (route.params.customer_slug) {
    openDrawerView(<string>route.params.customer_slug)
  }
});

const getCustomers = async (routeQuery: any) => {
  try {
    const res = await apiGet(`/_/api/customer`, routeQuery);
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error('Error fetching customer data:', error);
  }
};

const onSelectPage = (e: any) => {
  getCustomers(e)
};

const getCustomer = async (id: string, action: string) => {
  try {
    const res = await apiGet(`/_/api/customer/${id}`, {});
    if (res.code === 200) {
      isDrawer.value.data = res.result;
      isDrawer.value.open = true;
      isDrawer.value.action = action;
    }
  } catch (error) {
    console.error('Error fetching customer data:', error);
  }
};

const openDrawerView = async (id: string) => {
  getCustomer(id, "view")
};

const openDrawerEdit = async (id: string) => {
  getCustomer(id, "edit")
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

provide("getCustomers", getCustomers);
provide('closeDrawer', closeDrawer);
</script>