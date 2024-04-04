<template>
  <div class="artboard">
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
        <tr v-for="(item, index) in data.audits" :key="index" class="cursor" @click="openDrawerDesc(item.id)">
          <td v-if="item.customer.email === 'admin'">
            <Badge name="admin" color="indigo" />
          </td>
          <td :class="{ 'text-red-500': !item.customer.status }" v-else>
            <router-link active-class="current" :to="{ name: 'admin-customer-description', params: { customer_slug: item.customer.id } }">
              {{ item.customer.email }}
            </router-link>
          </td>
          <td>{{ sections[item.section] }}</td>
          <td>
            <Badge :name="actionFormat[item.action].name" :color="actionFormat[item.action].color" />
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

  <Drawer :is-open="isDrawer.open" @close="closeDrawer" maxWidth="600px">
    <header>
      <h1>Audit description</h1>
    </header>
    <div class="rounded border border-solid border-gray-300">
      <table class="mini" v-if="dataFull.id">
        <tr>
          <td class="w-32">ID</td>
          <td>{{ dataFull.id }}</td>
        </tr>
        <tr>
          <td class="w-32">Section</td>
          <td>{{ sections[dataFull.section] }}</td>
        </tr>
        <tr>
          <td class="w-32">Customer</td>
          <td v-if="dataFull.customer.email === 'admin'">
            <Badge name="admin" color="indigo" />
          </td>
          <td :class="{ 'text-red-500': !dataFull.customer.status }" v-else>
            <router-link active-class="current" :to="{ name: 'admin-customer-description', params: { customer_slug: dataFull.customer.id } }">
              {{ dataFull.customer.email }}
            </router-link>
          </td>
        </tr>
        <tr>
          <td class="w-32">Action</td>
          <td>
            <Badge :name="actionFormat[dataFull.action].name" :color="actionFormat[dataFull.action].color" />
          </td>
        </tr>
        <tr>
          <td class="w-32">User Agent</td>
          <td>{{ dataFull.metadata.request.user_agent }}</td>
        </tr>
        <tr>
          <td class="w-32">User IP</td>
          <td>{{ dataFull.metadata.request.user_ip }}</td>
        </tr>
        <tr>
          <td class="w-32">User Country</td>
          <td>{{ dataFull.metadata.request.user_country }}</td>
        </tr>
        <tr v-if="dataFull.metadata.data">
          <td class="w-32">Data</td>
          <td class="!p-0">
            <table class="mini">
              <tr v-for="(value, key, index) in dataFull.metadata.data" :key="index">
                <td class="!break-normal !pl-2">{{ key }}</td>
                <td>{{ value.value }}</td>
              </tr>
            </table>
          </td>
        </tr>
        <tr>
          <td>Created</td>
          <td>{{ formatDate(dataFull.created) }}</td>
        </tr>
      </table>
    </div>

    <div class="pt-4">
      <button class="btn" @click="closeDrawer">Close</button>
    </div>
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { Badge, Pagination, Drawer } from "@/components";
import { sections, actionFormat, formatDate } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  open: false,
});
const data: any = ref({});
const dataFull: any = ref({});
const route = useRoute();

onMounted(() => {
  getData(route.query);
  if (route.params.audit_slug) {
    openDrawerDesc(<string>route.params.audit_slug);
  }
});

const getData = async (routeQuery: any) => {
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
  getData(e);
};

const openDrawerDesc = async (id: string) => {
  try {
    const res = await apiGet(`/_/api/audit/${id}`, {});
    if (res.code === 200) {
      const { metadata } = res.result;
      if (metadata.data) {
        metadata.data = Object.entries(metadata.data).reduce(
          (acc, [key, value]) => ({
            ...acc,
            [key]: {
              value,
              isLong: String(value).length > 32,
            },
          }),
          {},
        );
      }
      dataFull.value = res.result;
    }
    isDrawer.value.open = true;
  } catch (error) {
    console.error("Error fetching audit data:", error);
  }
};

const closeDrawer = async () => {
  dataFull.value = {};
  isDrawer.value.open = false;
};

/*
const makeVisible = async (key) => {
  const item = dataFull.value.metadata.data[key];
  if (item) {
    item.isLong = !item.isLong;
  }
};
*/
</script>
