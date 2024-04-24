<template>
  <label :for="id" class="toggle">
    <input type="checkbox" :id="id" class="peer sr-only" v-model="value" :checked="Boolean(value)" :disabled="disabled" />
    <div
      class="peer h-6 w-11 rounded-full after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:border after:bg-white after:transition-all after:content-[''] peer-checked:after:translate-x-full"
      :class="color()">
    </div>
    <span v-if="name" class="ml-3 text-sm font-medium">{{ name }}</span>
  </label>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { customAlphabet } from "nanoid";

const nanoid = customAlphabet('1234567890abcdef', 10);

const props = defineProps({
  modelValue: {
    required: false,
  },
  name: {
    type: String,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  id: {
    type: String,
  },
  lColor: {
    type: String,
  },
  rColor: {
    type: String,
    default: "green",
  }
});

const emits = defineEmits(["update:modelValue"]);
const value = computed({
  get: () => {
    return props.modelValue;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});

const color = () => {
  const { rColor, lColor } = props;
  const rightColor = `peer-checked:bg-${rColor}-500 peer-checked:after:border-${rColor}-600`;
  const leftColor = lColor ? `bg-${lColor}-500 after:border-${lColor}-600` : "bg-gray-200 after:border-gray-300";
  return `${rightColor} ${leftColor}`;
}


const id = props.id ? props.id : "toggle_" + nanoid();
</script>

<style lang="scss">
.toggle {
  @apply relative inline-flex cursor-pointer select-none items-center;
}
</style>