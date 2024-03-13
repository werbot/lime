<template>
  <div class="artboard">
    <header>
      <h1>Licenses</h1>
    </header>
    <Tabs :tabs="secondMenu('license')" />

    <div class="artboard-content p-5">section under development</div>
  </div>
</template>

<script setup lang="ts">
import { Tabs } from "@/components";
import { secondMenu } from "@/utils/menu";
</script>