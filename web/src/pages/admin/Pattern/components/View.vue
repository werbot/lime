<template>
  <header>
    <h1>Pattern description</h1>
  </header>

  <div class="rounded border border-solid border-gray-300">
    <table class="mini">
      <tr>
        <td class="w-32">ID</td>
        <td>
          <span class="dot mr-2" :class="drawer.data.status ? 'bg-green-500' : 'bg-red-500'"></span>
          {{ drawer.data.id }}
        </td>
      </tr>
      <tr>
        <td>Name</td>
        <td>{{ drawer.data.name }}</td>
      </tr>
      <tr>
        <td>Created</td>
        <td>{{ formatDate(drawer.data.created) }}</td>
      </tr>
      <tr v-if="drawer.data.created!==drawer.data.updated">
        <td>Updated</td>
        <td>{{ formatDate(drawer.data.updated) }}</td>
      </tr>
      <tr>
        <td>Limits</td>
        <td class="!p-0">
          <table class="mini">
            <tr v-for="(value, key, index) in drawer.data.limit" :key="index">
              <td>{{ key }}</td>
              <td>
                <Badge :name="String(value)" />
              </td>
            </tr>
          </table>
        </td>
      </tr>
      <tr>
        <td>Price</td>
        <td>
          {{ priceFormat(drawer.data.price) }}
          {{ currency[drawer.data.currency-1] }}
        </td>
      </tr>
      <tr>
        <td>Licenses</td>
        <td>
          <Badge :name="String(drawer.data.licenses.total)" />
        </td>
      </tr>
      <tr>
        <td>Term</td>
        <td>
          <Badge :name="termFormat[drawer.data.term].name" :color="termFormat[drawer.data.term].color" />
        </td>
      </tr>
      <tr>
        <td>Check</td>
        <td class="!p-0">
          <table class="mini">
            <tr v-for="(value, key, index) in drawer.data.check" :key="index">
              <td>{{ key }}</td>
              <td>
                <SvgIcon name="check" class="text-green-500" v-if="value === true" />
                <SvgIcon name="x-mark" class="text-red-500" v-else />
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </div>

  <div class="pt-4">
    <button class="btn" @click="closeDrawer()">Close</button>
  </div>
</template>

<script setup lang="ts">
import { inject } from "vue";
import { SvgIcon, Badge } from "@/components";
import { termFormat, priceFormat, currency, formatDate } from "@/utils";

const closeDrawer = inject("closeDrawer") as Function;
const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
</script>
