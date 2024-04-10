<template>
  <Skeleton class="text-gray-200" v-if="!drawer.data" />
  <div v-else>
    <header>
      <h1>
        Clone <span class="text-green-600">{{ oldName }}</span> pattern
      </h1>
    </header>

    <Form @submit="onSubmit" v-slot="{ errors }" class="space-y-4">
      <FormInput name="Name" v-model="drawer.data.name" :error="errors.name" id="name" type="text" rules="required" title="Name" />

      <div class="pt-8">
        <div class="flex">
          <div class="flex-none">
            <button type="submit" class="btn btn-green" :disabled="loadingStatus">
              <div v-if="loadingStatus">
                <span>Loading...</span>
              </div>
              <span v-else>Clone and edit</span>
            </button>
            <button class="btn ml-5" @click="closeDrawer()">Close</button>
          </div>
          <div class="grow"></div>
        </div>
      </div>
    </Form>
  </div>
</template>

<script setup lang="ts">
import { ref, inject } from "vue";
import { Skeleton, FormInput } from "@/components";
import { randomString } from "@/utils";
import { Form } from "vee-validate";
import { apiPost } from "@/utils/api";

const loadingStatus = ref(false);
const getPatterns = inject("getPatterns") as Function;
const openDrawerEdit = inject("openDrawerEdit") as Function;
const closeDrawer = inject("closeDrawer") as Function;

const props = defineProps({
  drawer: {
    type: Object,
    required: true,
  },
});
const oldName = props.drawer.data.name;

const regex = /#[A-Za-z0-9_]{10}/g;
if (regex.test(props.drawer.data.name)) {
  props.drawer.data.name = props.drawer.data.name.replace(regex, `#${randomString(10)}`);
} else {
  props.drawer.data.name = `${oldName} #${randomString(10)} copy`;
}

const onSubmit = async () => {
  loadingStatus.value = true;
  try {
    const res = await apiPost(`/_/api/pattern/${props.drawer.data.id}`, {}, props.drawer.data,);
    if (res.code === 200) {
      getPatterns();
      openDrawerEdit(res.result.id);
    }
  } catch (error) {
    console.error("Error fetching sign data:", error);
  }
  loadingStatus.value = false;
};
</script>
