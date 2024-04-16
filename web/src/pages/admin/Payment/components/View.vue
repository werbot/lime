<template>
  <Skeleton class="text-gray-200" v-if="!drawer.data" />
  <div v-else>
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
            <router-link @click="closeDrawer()" :to="{ name: 'admin-customer-description', params: { customer_slug: drawer.data.customer.id } }">
              {{ drawer.data.customer.email }}
            </router-link>
          </td>
        </tr>
        <tr>
          <td>License</td>
          <td>
            <div v-if="drawer.data.transaction.status === 1 && drawer.data.pattern.licenses.total === 0" class="relative inline-flex">
              <SvgIcon name="ticket" class="text-red-500 mr-2" /> Issue a license
            </div>
            <div v-else-if="drawer.data.transaction.status === 1 && drawer.data.pattern.licenses.total > 0" class="flex flex-col">
              <div v-for="license in drawer.data.pattern.licenses.licenses" :key="license.id" class="inline-flex">
                <SvgIcon name="ticket" class="mr-2" :class="license.status ? 'text-green-500' : 'text-red-500'" />
                <router-link @click="closeDrawer()" :to="{ name: 'admin-license-description', params: { license_slug: license.id } }">
                  {{ license.id }}
                </router-link>
              </div>
            </div>
            <div v-else class="relative inline-flex">
              <SvgIcon name="ticket" class="text-gray-200 mr-2" /> It's impossible to issue a license.
            </div>
          </td>
        </tr>
        <tr>
          <td>Pattern</td>
          <td class="inline-flex">
            <SvgIcon name="pattern" class="mr-2" />
            <router-link @click="closeDrawer()" :to="{ name: 'admin-pattern-description', params: { pattern_slug: drawer.data.pattern.id } }">
              {{ drawer.data.pattern.name }}
            </router-link>
          </td>
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
          <td>
            <Badge :name="paymentProvidersObj[drawer.data.transaction.provider - 1].name" :color="paymentProvidersObj[drawer.data.transaction.provider - 1].color" />
          </td>
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

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <div class="btn cursor-pointer" @click="closeDrawer()">Close</div>
        </div>
        <div class="grow"></div>
        <div class="flex-none">
          <div class="btn btn-green cursor-pointer" @click="openDrawerEdit(drawer.data.id)">Edit</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue';
import { termObj, currencyObj, paymentProvidersObj, formatDate, priceFormat, paymentStatusObj } from "@/utils";
import { Skeleton, Badge, SvgIcon } from "@/components";

const openDrawerEdit = inject("openDrawerEdit") as Function;
const closeDrawer = inject('closeDrawer') as Function;

defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
</script>