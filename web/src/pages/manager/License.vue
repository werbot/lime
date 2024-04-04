<template>
  <div class="artboard">
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
          <th>Pattern</th>
          <th class="w-24">Term</th>
          <th class="w-24">Price</th>
          <th class="w-48">Created</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.licenses" :key="index" class="cursor" :class="{ 'bg-red-50': !item.status }" @click="openDrawerDesc(item.id)">
          <td>
            <div class="flex items-center">
              <span class="dot" :class="item.status ? 'bg-green-500' : 'bg-red-500'"></span>
            </div>
          </td>
          <td>{{ item.pattern.name }}</td>
          <td>
            <Badge :name="termFormat[item.pattern.term].name" :color="termFormat[item.pattern.term].color" />
          </td>
          <td>{{ priceFormat(item.pattern.price) }} {{ currency[item.pattern.currency] }}</td>
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
          <td>Created</td>
          <td>{{ formatDate(dataFull.created) }}</td>
        </tr>
        <tr v-if="dataFull.updated">
          <td>Updated</td>
          <td>{{ formatDate(dataFull.updated) }}</td>
        </tr>
        <tr>
          <td>Pattern</td>
          <td>{{ dataFull.pattern.name }} </td>
        </tr>
        <tr>
          <td>Limits</td>
          <td class="!p-0">
            <table class="mini">
              <tr v-for="(value, key, index) in dataFull.pattern.limit" :key="index">
                <td>{{ key }}</td>
                <td>{{ value }}</td>
              </tr>
            </table>
          </td>
        </tr>
        <tr>
          <td>Price</td>
          <td>{{ priceFormat(dataFull.pattern.price) }} {{ currency[dataFull.pattern.currency-1] }}</td>
        </tr>
        <tr>
          <td>Term</td>
          <td>
            <Badge :name="termFormat[dataFull.pattern.term].name" :color="termFormat[dataFull.pattern.term].color" />
          </td>
        </tr>
        <tr>
          <td>Hash</td>
          <td>{{ dataFull.hash }}</td>
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
import { SvgIcon, Badge, Pagination, Drawer } from "@/components";
import { termFormat, priceFormat, currency, formatDate } from "@/utils";
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
  if (route.params.license_slug) {
    openDrawerDesc(<string>route.params.license_slug)
  }
});

const getData = async (routeQuery: any) => {
  try {
    const res = await apiGet(`/api/license`, routeQuery);
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error('Error fetching license data:', error);
  }
};

const onSelectPage = (e: any) => {
  getData(e);
};

const openDrawerDesc = async (id: string) => {
  try {
    const res = await apiGet(`/api/license/lic_${id}`, {});
    if (res.code === 200) {
      dataFull.value = res.result;
      isDrawer.value.open = true;
      isDrawer.value.action = "desc";
    }
  } catch (error) {
    console.error('Error fetching license data:', error);
  }
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