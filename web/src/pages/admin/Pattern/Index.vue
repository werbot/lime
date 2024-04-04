<template>
  <div class="artboard">
    <header>
      <h1>Patterns</h1>
      <label class="plus" @click="openDrawerAdd()">
        <SvgIcon name="plus-square" />
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
          <th class="w-20"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.patterns" :key="index" :class="{ 'opacity-30': item.private }" class="cursor">
          <td @click="openDrawerView(item.id)">
            <div class="flex items-center">
              <span class="dot" :class="item.status ? 'bg-green-500' : 'bg-red-500'"></span>
            </div>
          </td>
          <td @click="openDrawerView(item.id)">{{ item.name }}</td>
          <td @click="openDrawerView(item.id)">
            <Badge :name="String(item.licenses.total)" />
          </td>
          <td @click="openDrawerView(item.id)">
            <Badge :name="termFormat[item.term].name" :color="termFormat[item.term].color" />
          </td>
          <td @click="openDrawerView(item.id)">
            {{ priceFormat(item.price) }} {{ currency[item.currency-1] }}
          </td>
          <td>
            <div class="flex">
              <div class="pr-3">
                <SvgIcon name="pencil-square" @click="openDrawerEdit(item.id)" />
              </div>
              <div>
                <SvgIcon name="document-duplicate" @click="openDrawerClone(item.id, item.name)" />
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
    <Add v-if="isDrawer.action === 'add'" />
    <Clone :drawer="isDrawer" v-if="isDrawer.action === 'clone'" />
    <Edit :drawer="isDrawer" v-if="isDrawer.action === 'edit'" />
    <View :drawer="isDrawer" v-if="isDrawer.action === 'view'" />
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref, provide } from "vue";
import { useRoute } from "vue-router";
import { View, Edit, Add, Clone } from "./components";
import { SvgIcon, Badge, Pagination, Drawer } from "@/components";
import { termFormat, priceFormat, currency } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref<{
  data: Object;
  open: boolean;
  action: string;
}>({
  data: {},
  open: false,
  action: null,
});
const data: any = ref({});
const route = useRoute();

onMounted(() => {
  getPatterns(route.query);
  if (route.params.pattern_slug) {
    openDrawerView(<string>route.params.pattern_slug);
  }
});

const getPatterns = async (routeQuery: any) => {
  try {
    const res = await apiGet(`/_/api/pattern`, routeQuery);
    if (res.code === 200) {
      data.value = res.result;
    }
  } catch (error) {
    console.error("Error fetching audit data:", error);
  }
};

const onSelectPage = (e: any) => {
  getPatterns(e);
};

const getPattern = async (id: string, action: string) => {
  try {
    const res = await apiGet(`/_/api/pattern/${id}`, {});
    if (res.code === 200) {
      isDrawer.value.data = res.result;
      isDrawer.value.open = true;
      isDrawer.value.action = action;
    }
  } catch (error) {
    console.error("Error fetching audit data:", error);
  }
};

const openDrawerView = async (id: string) => {
  closeDrawer();
  getPattern(id, "view");
};

const openDrawerEdit = async (id: string) => {
  closeDrawer();
  getPattern(id, "edit");
};

const openDrawerAdd = async () => {
  closeDrawer();
  isDrawer.value.open = true;
  isDrawer.value.action = "add";
};

const openDrawerClone = async (id: string, name: string) => {
  closeDrawer();
  isDrawer.value.data = {
    id: id,
    name: name,
  };
  isDrawer.value.open = true;
  isDrawer.value.action = "clone";
};

const closeDrawer = async () => {
  isDrawer.value.data = {};
  isDrawer.value.open = false;
  isDrawer.value.action = null;
};

provide("getPatterns", getPatterns);
provide("openDrawerEdit", openDrawerEdit);
provide("closeDrawer", closeDrawer);
</script>
