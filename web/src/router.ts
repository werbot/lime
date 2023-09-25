import { createRouter, createWebHistory } from "vue-router";
import { getCookie } from "@/utils";
// @ts-ignore
import * as NProgress from "nprogress";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/install",
      name: "install",
      meta: { layouts: "BlankLayouts" },
      component: () => import("@/pages/Install.vue"),
    },
    {
      path: "/signin",
      name: "signin",
      meta: { layouts: "BlankLayouts" },
      component: () => import("@/pages/Signin.vue"),
    },
    {
      path: "/:pathMatch(.*)*",
      name: "404",
      meta: { layouts: "BlankLayouts" },
      component: () => import("@/pages/404.vue"),
    },
  ],
});

router.beforeEach((to, from, next) => {
  NProgress.start();

  let isAuthenticated = false;
  let token = getCookie("token");
  if (token) {
    isAuthenticated = true;
  }

  if (to.path === "/install") next();
  else if (!isAuthenticated && to.name !== "signin") next({ name: "signin" });
  else if (isAuthenticated && to.name == "signin") next({ name: "products" });
  else next();
});

router.afterEach(() => {
  NProgress.done();
});

export default router;