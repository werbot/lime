import { defineStore } from "pinia";

interface ErrorState {
  message: null | string;
  errors: any;
}

// Define a store named 'error'
export const useErrorStore = defineStore("error", {
  state: (): ErrorState => ({
    message: null,
    errors: {},
  }),

  getters: {
    errorCode(state) {
      return Object.keys(state.errors);
    },
  },

  actions: {
    resetStore() {
      this.$reset();
    },

    /**
     * Setter method for the `message` property.
     * @param message - The error message to be stored.
     */
    setErrorMessage(message: string) {
      this.message = message;
    },

    /**
     * Setter method for the `errors` property.
     * @param errors - The errors to be stored.
     */
    setErrors(errors: any) {
      this.errors = errors;
    },
  },
});
