<template>
  <Skeleton class="text-gray-200" v-if="!drawer" />
  <div v-else>
    <header>
      <h1>Customer description</h1>
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
          <td class="w-32">Email</td>
          <td>{{ drawer.data.email }}</td>
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
import { Skeleton } from "@/components";
import { formatDate } from "@/utils";

const openDrawerEdit = inject("openDrawerEdit") as Function;
const closeDrawer = inject('closeDrawer') as Function;

defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
</script>