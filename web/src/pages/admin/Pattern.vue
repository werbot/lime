<template>
  <div class="artboard">
    <header>
      <h1>Pattern</h1>
      <label class="plus" @click="openDrawerAdd()">
        <SvgIcon name="plus_square" />
        add new
      </label>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th class="w-12"></th>
          <th>Name</th>
          <th class="w-24">Licenses</th>
          <th class="w-24">Term</th>
          <th class="w-24">Price</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.patterns" :key="index" :class="{ 'opacity-30': item.private }" class="cursor" @click="openDrawerDesc(item.id)">
          <td>
            <div class="flex items-center">
              <span class="dot" :class="item.status ? 'bg-green-500' : 'bg-red-500'"></span>
            </div>
          </td>
          <td>{{ item.name }}</td>
          <td><Badge :name="String(item.licenses.total)" /></td>
          <td>
            <Badge :name="termFormat[item.term].name" :color="termFormat[item.term].color" />
          </td>
          <td>{{ priceFormat(item.price) }} {{ currency[item.currency] }}</td>
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
          <td>Name</td>
          <td>
            {{ dataFull.name }}
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
          <td>Limits</td>
          <td class="!p-0">
            <table class="mini">
              <tr v-for="(value, key, index) in dataFull.limit" :key="index">
                <td>{{ key }}</td>
                <td>{{ value }}</td>
              </tr>
            </table>
          </td>
        </tr>
        <tr>
          <td>Price</td>
          <td>{{ priceFormat(dataFull.price) }} {{ currency[dataFull.currency] }}</td>
        </tr>
        <tr>
          <td>Licenses</td>
          <td>
            <Badge :name="dataFull.licenses.total" />
          </td>
        </tr>
        <tr>
          <td>Term</td>
          <td>
            <Badge :name="termFormat[dataFull.term].name" :color="termFormat[dataFull.term].color" />
          </td>
        </tr>
        <tr>
          <td>Check</td>
          <td class="!p-0">
            <table class="mini">
              <tr v-for="(value, key, index) in dataFull.check" :key="index">
                <td>{{ key }}</td>
                <td>
                  <SvgIcon name="check" class="text-green-500" v-if="value === 1" />
                  <SvgIcon name="x-mark" class="text-red-500" v-else />
                </td>
              </tr>
            </table>
          </td>
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
  if (route.params.pattern_slug) {
    openDrawerDesc(<string>route.params.pattern_slug)
  }
});

const getData = async (routeQuery: any) => {
  apiGet(`/_/api/pattern`, routeQuery).then(res => {
    if (res.code === 200) {
      data.value = res.result;
    }
  })
};

const onSelectPage = (e: any) => {
  getData(e);
};

const openDrawerDesc = async (id: string) => {
  apiGet(`/_/api/pattern/${id}`, {}).then(res => {
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