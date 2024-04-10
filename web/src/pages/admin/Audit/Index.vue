<template>
  <Skeleton class="text-gray-200" v-if="!data" />
  <div class="artboard" v-else>
    <header>
      <h1>Audit</h1>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th class="w-52">Customer</th>
          <th class="w-36">Section</th>
          <th class="w-36">Action</th>
          <th></th>
          <th class="w-48">Created</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.audits" :key="index" class="cursor" @click="openDrawerView(item.id)">
          <td v-if="item.customer.email === 'admin'">
            <Badge name="admin" color="indigo" />
          </td>
          <td :class="{ 'text-red-500': !item.customer.status }" v-else>
            <router-link active-class="current" :to="{ name: 'admin-customer-description', params: { customer_slug: item.customer.id } }">
              {{ item.customer.email }}
            </router-link>
          </td>
          <td>{{ sectionsObj[item.section - 1].name }}</td>
          <td>
            <Badge :name="actionObj[item.action - 1].name" :color="actionObj[item.action - 1].color" />
          </td>
          <td></td>
          <td>{{ formatDate(item.created) }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else class="desc">Empty</div>

    <div class="artboard-content">
      <Pagination :total="data.total" @selectPage="onSelectPage" />
    </div>
  </div>

  <Drawer :is-open="isDrawer.open" @close="closeDrawer()" maxWidth="600px">
    <View :drawer="isDrawer" />
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref, provide } from "vue";
import { useRoute } from "vue-router";
import { View } from "./components";
import { Skeleton, Badge, Pagination, Drawer } from "@/components";
import { sectionsObj, actionObj, formatDate } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  data: {},
  open: false,
});
const data: any = ref({});
const route = useRoute();

onMounted(() => {
  getAudits(route.query);
  if (route.params.audit_slug) {
    openDrawerView(<string>route.params.audit_slug);
  }
});

const getAudits = async (routeQuery: any) => {
  try {
    const res = await apiGet(`/_/api/audit`, routeQuery);
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error("Error fetching audit data:", error);
  }
};

const onSelectPage = (e: any) => {
  getAudits(e);
};

const getAudit = async (id: string) => {
  try {
    const res = await apiGet(`/_/api/audit/${id}`, {});
    if (res.code === 200) {
      isDrawer.value.data = res.result;
      isDrawer.value.open = true;
    }
  } catch (error) {
    console.error('Error fetching customer data:', error);
  }
};

const openDrawerView = async (id: string) => {
  getAudit(id)
};

const closeDrawer = async () => {
  isDrawer.value.data = {};
  isDrawer.value.open = false;
};

provide('closeDrawer', closeDrawer);
</script>
