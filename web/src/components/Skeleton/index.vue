<template>
  <div class="animate-pulse">
    <div class="flex items-center justify-between space-x-4">
      <div v-if="avatar" class="w-12 h-12 bg-current rounded-full shrink-0"></div>
      <div v-if="title" class="w-full h-8 bg-current rounded-lg"></div>
    </div>

    <div v-if="lines > 0" class="mt-6 space-y-2">
      <div v-for="(setting, index) in lineSettings" class="h-4 bg-current rounded-lg" :class="setting.margin" :key="index"></div>
    </div>

    <div v-if="footer" class="w-1/3 h-10 mt-6 bg-current rounded-lg"></div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from "vue";

const props = defineProps({
  avatar: {
    type: Boolean,
    default: false,
  },
  title: {
    type: Boolean,
    default: true,
  },
  lines: {
    type: Number,
    default: 3,
  },
  footer: {
    type: Boolean,
    default: true,
  },
});

const lineSettings = new Array(props.lines).fill(undefined).map(() => {
  const margins = [4, 5, 6, 8, 10, 16, 24, 32];
  const margin = margins[Math.floor(Math.random() * margins.length)]

  return {
    margin: `mr-${margin}`,
  };
});
</script>