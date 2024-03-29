<template>
  <div class="artboard">
    <header>
      <h1>Customers</h1>
      <label class="plus" @click="openDrawerAdd()">
        <SvgIcon name="plus_square" />
        add new
      </label>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th class="w-12"></th>
          <th>Email</th>
          <th class="w-48">Created</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.customers" :key="index" class="cursor" :class="{ 'bg-red-50': !item.status }" @click="openDrawerDesc(item.id)">
          <td>
            <div class="flex items-center">
              <span class="dot" :class="item.status ? 'bg-green-500' : 'bg-red-500'"></span>
            </div>
          </td>
          <td :class="{ 'text-red-500': !item.status }">{{ item.email }}</td>
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
    <div class="rounded border border-solid border-gray-300" v-if="isDrawer.action === 'desc'">
      <table class="mini" v-if="dataFull.id">
        <tr>
          <td class="w-32">ID</td>
          <td>
            <span class="dot mr-2" :class="dataFull.status ? 'bg-green-500' : 'bg-red-500'"></span>
            {{ dataFull.id }}
          </td>
        </tr>
        <tr>
          <td class="w-32">Email</td>
          <td>{{ dataFull.email }}</td>
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

    <div v-if="isDrawer.action === 'add'">
      <h2>Add</h2>
    </div>
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { SvgIcon, Pagination, Drawer } from "@/components";
import { formatDate } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  open: false,
  action: null,
})
const data: any = ref({});
const dataFull: any = ref({});
const route = useRoute();

onMounted(() => {
  getData(route.query);
  if (route.params.customer_slug) {
    openDrawerDesc(<string>route.params.customer_slug)
  }
});

const getData = async (routeQuery: any) => {
  apiGet(`/_/api/customer`, routeQuery).then(res => {
    if (res.code === 200) {
      data.value = res.result;
    }
  });
};

const onSelectPage = (e: any) => {
  getData(e)
};

const openDrawerDesc = async (id: string) => {
  apiGet(`/_/api/customer/${id}`, {}).then(res => {
    if (res.code === 200) {
      dataFull.value = res.result;
      isDrawer.value.open = true;
      isDrawer.value.action = "desc";
    }
  });
};

const openDrawerAdd = async () => {
  isDrawer.value.open = true;
  isDrawer.value.action = "add";
};

const closeDrawer = async () => {
  dataFull.value = {};
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};
</script>