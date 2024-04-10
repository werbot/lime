<template>
  <header>
    <h1>Audit description</h1>
  </header>

  <div class="rounded border border-solid border-gray-300">
    <table class="mini" v-if="drawer.data.id">
      <tr>
        <td class="w-32">ID</td>
        <td>{{ drawer.data.id }}</td>
      </tr>
      <tr>
        <td class="w-32">Section</td>
        <td>{{ sectionsObj[drawer.data.section - 1].name }}</td>
      </tr>
      <tr>
        <td class="w-32">Customer</td>
        <td v-if="drawer.data.customer.email === 'admin'">
          <Badge name="admin" color="indigo" />
        </td>
        <td :class="{ 'text-red-500': !drawer.data.customer.status }" v-else>
          <router-link active-class="current" :to="{ name: 'admin-customer-description', params: { customer_slug: drawer.data.customer.id } }">
            {{ drawer.data.customer.email }}
          </router-link>
        </td>
      </tr>
      <tr>
        <td class="w-32">Action</td>
        <td>
          <Badge :name="actionObj[drawer.data.action - 1].name" :color="actionObj[drawer.data.action - 1].color" />
        </td>
      </tr>
      <tr>
        <td class="w-32">User Agent</td>
        <td>{{ drawer.data.metadata.request.user_agent }}</td>
      </tr>
      <tr>
        <td class="w-32">User IP</td>
        <td>{{ drawer.data.metadata.request.user_ip }}</td>
      </tr>
      <tr>
        <td class="w-32">User Country</td>
        <td>{{ drawer.data.metadata.request.user_country }}</td>
      </tr>
      <tr v-if="drawer.data.metadata.data">
        <td class="w-32">Data</td>
        <td class="!p-0">
          <table class="mini">
            <tr v-for="(value, key, index) in drawer.data.metadata.data" :key="index">
              <td class="!break-normal !pl-2">{{ key }}</td>
              <td>{{ value }}</td>
            </tr>
          </table>
        </td>
      </tr>
      <tr>
        <td>Created</td>
        <td>{{ formatDate(drawer.data.created) }}</td>
      </tr>
    </table>
  </div>

  <div class="pt-4">
    <button class="btn" @click="closeDrawer()">Close</button>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue';
import { Badge } from "@/components";
import { sectionsObj, actionObj, formatDate } from "@/utils";

const closeDrawer = inject('closeDrawer') as Function;
const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
</script>