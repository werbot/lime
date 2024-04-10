<template>
  <Skeleton class="text-gray-200" v-if="!data" />
  <div class="artboard" v-else>
    <header>
      <h1>Licenses</h1>
      <label class="plus" @click="openDrawerAdd()">
        <SvgIcon name="plus-square" />
        add new
      </label>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th class="w-12"></th>
          <th class="w-52">Customer</th>
          <th>Pattern</th>
          <th class="w-24">Term</th>
          <th class="w-24">Price</th>
          <th class="w-48">Created</th>
          <th class="w-8"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.licenses" :key="index" class="cursor" :class="{ 'bg-red-50': !item.status }">
          <td @click="openDrawerView(item.id)">
            <div class="flex items-center">
              <span class="dot" :class="item.status ? 'bg-green-500' : 'bg-red-500'"></span>
            </div>
          </td>
          <td @click="openDrawerView(item.id)" :class="{ 'text-red-500': !item.customer.status }">{{ item.customer.email }}</td>
          <td @click="openDrawerView(item.id)">{{ item.pattern.name }}</td>
          <td @click="openDrawerView(item.id)">
            <Badge :name="termObj[item.pattern.term - 1].name" :color="termObj[item.pattern.term - 1].color" />
          </td>
          <td @click="openDrawerView(item.id)">{{ priceFormat(item.pattern.price) }} {{ currencyObj[item.pattern.currency - 1].name }}</td>
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
    <Add :drawer="isDrawer" v-if="isDrawer.action === 'add'" />
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref, provide } from "vue";
import { useRoute } from "vue-router";
import { View, Edit, Add } from "./components";
import { Skeleton, SvgIcon, Badge, Pagination, Drawer } from "@/components";
import { termObj, currencyObj, priceFormat, formatDate } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  data: {},
  open: false,
  action: null,
})
const data: any = ref({});
const route = useRoute();

onMounted(() => {
  getLicenses(route.query);
  if (route.params.license_slug) {
    openDrawerView(<string>route.params.license_slug)
  }
});

const getLicenses = async (routeQuery: any) => {
  try {
    const res = await apiGet(`/_/api/license`, routeQuery);
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error('Error fetching license data:', error);
  }
};

const onSelectPage = (e: any) => {
  getLicenses(e);
};

const getLicense = async (id: string, action: string) => {
  closeDrawer();
  try {
    const res = await apiGet(`/_/api/license/lic_${id}`, {});
    if (res.code === 200) {
      isDrawer.value.data = res.result;
      isDrawer.value.open = true;
      isDrawer.value.action = action;
    }
  } catch (error) {
    console.error('Error fetching license data:', error);
  }
};

const openDrawerView = async (id: string) => {
  getLicense(id, "view")
};

const openDrawerEdit = async (id: string) => {
  getLicense(id, "edit")
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

provide("openDrawerEdit", openDrawerEdit);
provide('closeDrawer', closeDrawer);
</script>