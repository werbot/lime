<template>
  <header>
    <h1>Add new license</h1>
  </header>

  <div class="pt-4 flex flex-row justify-end">
    <button class="btn btn-red" @click="closeDrawer()">Delete customer</button>
    <button class="btn ml-5" @click="closeDrawer()">Close</button>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue';

const closeDrawer = inject('closeDrawer') as Function;

const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
</script>