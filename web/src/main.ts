import { createApp } from "vue";

import App from "@/app.vue";
import Router from "@/router";
import Notifications from 'notiwind'

import "@/assets/app.css";

const app = createApp(App);
app.use(Router);
app.use(Notifications);
app.mount("#app");
