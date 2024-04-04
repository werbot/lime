<template>
  <div class="form-dropdown">
    <label class="label">
      <span v-if="name" class="text">{{ name }}{{ required ? "*" : "" }}</span>
      <span v-if="error" class="error">{{ error }}</span>
    </label>
    <Field v-slot="{ value }" v-model="model" as="select" :name="id" :id="id" :rules="rules" class="form-select field peer">
      <option value="" disabled>Please select</option>
      <option v-for="(option, key) in options" :key="key" :value="key" :selected="value">
        {{ option }}
      </option>
    </Field>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { Field } from "vee-validate";

const props = defineProps({
  modelValue: {
    required: true,
  },
  id: {
    type: String,
    default: "name",
  },
  name: {
    type: String,
    default: "Name",
  },
  required: {
    type: Boolean,
  },
  options: {
    type: Object,
    required: true,
  },
  rules: {
    type: String,
    default: "",
  },
  error: String,
});

const emits = defineEmits(["update:modelValue"]);
const model = computed({
  get: () => {
    return props.modelValue;
  },
  set: (val) => {
    emits("update:modelValue", val);
  },
});
</script>

<style lang="scss" scoped>
.form-dropdown {
  .field {
    @apply w-full border-none rounded bg-gray-50 ring-1 ring-gray-300;
  }
}
</style>