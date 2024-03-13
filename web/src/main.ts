import { createApp } from "vue";

import App from "@/App.vue";
import Router from "@/router";
import Notifications from "notiwind";
import Store from "@/store";

import { defineRule } from "vee-validate";
import * as rules from "@vee-validate/rules";

import "virtual:svg-icons-register";
import "@/assets/app.scss";

// validate rules
Object.keys(rules).forEach((rule) => {
  defineRule(rule, rules[rule]);
});

const app = createApp(App);
app.use(Store);
app.use(Router);
app.use(Notifications);
app.mount("#app");
