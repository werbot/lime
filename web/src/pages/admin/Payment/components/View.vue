<template>
  <header>
    <h1>Payment description</h1>
  </header>

  <div class="rounded border border-solid border-gray-300">
    <table class="mini" v-if="drawer.data.id">
      <tr>
        <td class="w-32">ID</td>
        <td>{{ drawer.data.id }}</td>
      </tr>
      <tr>
        <td>Customer</td>
        <td>
          <span class="dot mr-2" :class="drawer.data.customer.status ? 'bg-green-500' : 'bg-red-500'"></span>
          {{ drawer.data.customer.email }}
        </td>
      </tr>
      <tr>
        <td>Pattern</td>
        <td>{{ drawer.data.pattern.name }}</td>
      </tr>
      <tr>
        <td>Term</td>
        <td>
          <Badge :name="termObj[drawer.data.pattern.term - 1].name" :color="termObj[drawer.data.pattern.term - 1].color" />
        </td>
      </tr>
      <tr>
        <td>Price</td>
        <td>{{ priceFormat(drawer.data.pattern.price) }} {{ currencyObj[drawer.data.pattern.currency - 1].name }}</td>
      </tr>
      <tr>
        <td>Provider</td>
        <td>{{ drawer.data.transaction.provider }}</td>
      </tr>
      <tr>
        <td>Status</td>
        <td>
          <Badge :name="paymentStatusObj[drawer.data.transaction.status - 1].name" :color="paymentStatusObj[drawer.data.transaction.status - 1].color" />
        </td>
      </tr>
      <tr>
        <td>Created</td>
        <td>{{ formatDate(drawer.data.created) }}</td>
      </tr>
      <tr v-if="drawer.data.created !== drawer.data.updated">
        <td>Updated</td>
        <td>{{ formatDate(drawer.data.updated) }}</td>
      </tr>
    </table>
  </div>

  <div class="pt-4">
    <button class="btn" @click="closeDrawer()">Close</button>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue';
import { termObj, currencyObj, formatDate, priceFormat, paymentStatusObj } from "@/utils";
import { Badge } from "@/components";

const closeDrawer = inject('closeDrawer') as Function;
const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
</script>