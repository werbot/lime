<template>
  <div class="artboard">
    <header>
      <h1>Audit</h1>
    </header>

    <table v-if="data.total > 0">
      <thead>
        <tr>
          <th>Section</th>
          <th class="w-36">Action</th>
          <th class="w-48">Created</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in data.audits" :key="index" class="cursor" @click="openDrawerDesc(item.id)">
          <td>{{ sections[item.section] }}</td>
          <td>
            <Badge :name="actionFormat[item.action].name" :color="actionFormat[item.action].color" />
          </td>
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
    <div class="rounded border border-solid border-gray-300">
      <table class="mini" v-if="dataFull.id">
        <tr>
          <td class="w-32">ID</td>
          <td>
            {{ dataFull.id }}
          </td>
        </tr>
      </table>
    </div>
  </Drawer>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { SvgIcon, Badge, Pagination, Drawer } from "@/components";
import { sections, actionFormat, formatDate } from "@/utils";
import { apiGet } from "@/utils/api";

const isDrawer = ref({
  open: false,
})
const data: any = ref({});
const dataFull: any = ref({});
const route = useRoute();

onMounted(() => {
  getData(route.query);
  if (route.params.audit_slug) {
    openDrawerDesc(<string>route.params.audit_slug)
  }
});

const getData = async (routeQuery: any) => {
  apiGet(`/_/api/audit`, routeQuery).then(res => {
    if (res.code === 200) {
      data.value = res.result;
    }
  });
};

const onSelectPage = (e: any) => {
  getData(e)
};

const openDrawerDesc = async (id: string) => {
  apiGet(`/_/api/audit/${id}`, {}).then(res => {
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