<template>
  <Skeleton class="text-gray-200" v-if="!drawer.data" />
  <div v-else>
    <header>
      <h1>License description</h1>
    </header>

    <div class="rounded border border-solid border-gray-300">
      <table class="mini" v-if="drawer.data.id">
        <tr>
          <td class="w-32">ID</td>
          <td>
            <span class="dot mr-2" :class="drawer.data.status ? 'bg-green-500' : 'bg-red-500'"></span>
            {{ drawer.data.id }}
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
        <tr>
          <td>Payment</td>
          <td class="inline-flex">
            <SvgIcon name="banknotes" class="mr-2" />
            <router-link @click="closeDrawer()" :to="{ name: 'admin-payment-description', params: { payment_slug: drawer.data.payment.id } }">
              {{ drawer.data.payment.id }}
            </router-link>
          </td>
        </tr>
        <tr>
          <td>Customer</td>
          <td class="inline-flex">
            <SvgIcon name="users" :class="drawer.data.payment.customer.status ? 'text-green-500' : 'text-red-500'" class="mr-2" />
            <router-link @click="closeDrawer()" :to="{ name: 'admin-customer-description', params: { customer_slug: drawer.data.payment.customer.id } }">
              {{ drawer.data.payment.customer.email }}
            </router-link>
          </td>
        </tr>
        <tr>
          <td>Pattern</td>
          <td>
            <router-link @click="closeDrawer()" :to="{ name: 'admin-pattern-description', params: { pattern_slug: drawer.data.payment.pattern.id } }">
              {{ drawer.data.payment.pattern.name }}
            </router-link>
          </td>
        </tr>
        <tr>
          <td>Limits</td>
          <td class="!p-0">
            <table class="mini">
              <tr v-for="(value, key, index) in drawer.data.payment.pattern.limit" :key="index">
                <td>{{ key }}</td>
                <td><Badge :name="String(value)" /></td>
              </tr>
            </table>
          </td>
        </tr>
        <tr>
          <td>Price</td>
          <td>{{ priceFormat(drawer.data.payment.pattern.price) }} {{ currencyObj[drawer.data.payment.pattern.currency - 1].name }}</td>
        </tr>
        <tr>
          <td>Term</td>
          <td>
            <Badge :name="termObj[drawer.data.payment.pattern.term - 1].name" :color="termObj[drawer.data.payment.pattern.term - 1].color" />
          </td>
        </tr>
        <tr>
          <td>Hash</td>
          <td>{{ drawer.data.hash }}</td>
        </tr>
        <tr>
          <td>License</td>
          <td>{{ drawer.data.data }}</td>
        </tr>
      </table>
    </div>

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <div class="btn cursor-pointer" @click="closeDrawer()">Close</div>
        </div>
        <div class="grow"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue';
import { SvgIcon, Skeleton, Badge } from "@/components";
import { termObj, currencyObj, priceFormat, formatDate } from "@/utils";

const closeDrawer = inject('closeDrawer') as Function;

defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
</script>