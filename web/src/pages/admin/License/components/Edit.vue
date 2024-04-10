<template>
  <Skeleton class="text-gray-200" v-if="!drawer" />
  <div v-else>
    <header>
      <h1>Edit license</h1>
    </header>

    {{ drawer.data }}

    <div class="pt-8">
      <div class="flex">
        <div class="flex-none">
          <button type="submit" class="btn btn-green" :disabled="loadingStatus">
            <div v-if="loadingStatus">
              <span>Loading...</span>
            </div>
            <span v-else>Save</span>
          </button>
          <button class="btn ml-5" @click="closeDrawer()">Close</button>
        </div>
        <div class="grow"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, inject } from 'vue';
import { Skeleton, FormInput } from "@/components";
import { Form } from "vee-validate";

const openDrawerEdit = inject("openDrawerEdit") as Function;
const closeDrawer = inject('closeDrawer') as Function;

const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});

const loadingStatus = ref(false);

const data = ref({
  email: null,
  password: null,
});

const onSubmit = async () => {
  loadingStatus.value = true;

  loadingStatus.value = false;
};
</script>