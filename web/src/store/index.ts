import { App } from "vue";
import { createPinia } from "pinia";

import { useErrorStore } from "@/store/modules/error";

export default {
  install: (app: App) => {
    app.use(createPinia());

    const stores = {
      errorStore: useErrorStore(),
    };

    Object.entries(stores).forEach(([name, store]) => {
      app.config.globalProperties[`$${name}`] = store;
    });
  },
};

export {
  useErrorStore, 
};