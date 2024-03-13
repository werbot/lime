<template>
  <nav class="header">
    <div class="logo">
      <img src="/img/logo.svg" alt="Lime by Werbot" />
    </div>

    <div class="exit" @click="signOut">
      <SvgIcon name="exit" class="mr-2 w-5 text-gray-700" />
      Exit
    </div>
  </nav>
  <hr />
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { delCookie } from "@/utils/";
import { SvgIcon } from "@/components";

const router = useRouter();

const signOut = async () => {
  delCookie("admin");
  router.push({ name: "adminSignin" });
};
</script>

<style lang="scss">
.header {
  @apply relative flex w-full flex-wrap items-center justify-between py-4;

  .logo {
    @apply mr-5 w-32 flex-none cursor-pointer sm:mr-12 sm:w-36;
  }

  .exit {
    @apply px-2 no-underline cursor-pointer flex items-center whitespace-nowrap border rounded bg-gray-200 p-1 text-gray-700;

    svg {
      @apply fill-gray-700;
    }
  }
}
</style>