<template>
  <div class="form-control" :for="id" :class="class">
    <label class="label">
      <span v-if="name" class="text">{{ fullName }}</span>
      <span v-if="error" class="error">{{ error }}</span>
    </label>
    <Field :type="type" :name="id" :rules="rules" :id="id" v-model="value" :class="error ? 'error' : ''" class="input" :disabled="disabled" :placeholder="placeholder"
      :autocomplete="autocomplete" />
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { Field } from 'vee-validate';

const props = defineProps({
  name: String,
  modelValue: {
    required: true,
  },
  id: String,
  type: {
    type: String,
    default: "text",
  },
  class: String,
  disabled: Boolean,
  required: Boolean,
  placeholder: String,
  autocomplete: String,
  rules: String,
  error: String
});

const emits = defineEmits(["update:modelValue"]);

const value = computed({
  get: () => props.modelValue,
  set: (val) => emits("update:modelValue", val),
});

const fullName = computed(() => `${props.name}${props.required ? "*" : ""}`);
</script>