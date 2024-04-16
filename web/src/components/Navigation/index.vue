<template>
  <nav class="sidebar">
    <ul>
      <li v-for="(item) in mainMenu()">
        <router-link :to="item.link" :class="{ current: (route.name as string).startsWith(item.link.name) }">
          <SvgIcon :name="item.icon" />
          <span>{{ item.name }}</span>
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import { mainMenu } from "@/utils/menu";
import { SvgIcon } from "@/components";

const route = useRoute();
</script>

<style lang="scss">
.sidebar {
  @apply mb-6 mt-1;

  li {
    @apply mb-4 w-3 sm:w-36;

    a {
      @apply m-2 leading-5 text-gray-400;

      &:hover {
        @apply m-0 rounded bg-gray-100 p-2;
      }

      span {
        @apply hidden sm:inline-block;
      }

      svg {
        @apply inline-block -mt-1 w-5 fill-gray-400 pr-0 align-middle sm:mr-2.5;
      }

      &.current {
        @apply m-0 rounded bg-gray-200 p-2 text-gray-700;

        svg {
          @apply fill-gray-700;
        }
      }
    }
  }
}
</style>