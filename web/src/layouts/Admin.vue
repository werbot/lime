<template>
  <div class="px-6">
    <Header />

    <div class="content" >
      <div class="left">
        <Navigation />
      </div>
      <div class="right">
        <div id="alert"></div>
        <RouterView />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Header, Navigation } from "@/components";
</script>