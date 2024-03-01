import { createRouter, createWebHistory } from "vue-router";
import { getCookie } from "@/utils";
import * as NProgress from "nprogress";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "manager",
      meta: { layout: "Blank" },
      component: () => import("@/pages/Manager.vue"),
    },
    {
      path: "/signin",
      name: "signin",
      meta: { layout: "Signin" },
      component: () => import("@/pages/Signin.vue"),
    },
    {
      path: "/_",
      name: "adminLicense",
      meta: { layout: "Blank" },
      children: [
        {
          path: "install",
          name: "adminInstall",
          component: () => import("@/pages/admin/Install.vue"),
        },
        {
          path: "signin",
          name: "adminSignin",
          meta: { layout: "Signin" },
          component: () => import("@/pages/admin/Signin.vue"),
        },
        {
          path: "license",
          name: "adminLicense",
          component: () => import("@/pages/admin/License.vue"),
        },
      ],
    },
    {
      path: "/:pathMatch(.*)*",
      name: "404",
      meta: { layout: "404" },
      component: () => import("@/pages/404.vue"),
    },
  ],
});

router.beforeEach((to, from) => {
  NProgress.start();

  loadLayoutMiddleware(to);

  let isAuthenticated = false;
  if (getCookie("token")) {
    isAuthenticated = true;
  }

  let isAdmin = false;
  if (getCookie("admin")) {
    isAdmin = true;
  }

  if (to.path.startsWith("/_")) {
    if (!isAdmin && to.name !== "adminSignin") {
      return { name: "adminSignin" };
    }
    if (isAdmin && to.name === "adminSignin") {
      return { name: "adminLicense" };
    }
  } else {
    if (!isAuthenticated && to.name !== "signin") {
      return { name: "signin" };
    }
    if (isAuthenticated && to.name === "signin") {
      return { name: "manager" };
    }
  }
});

router.afterEach(() => {
  NProgress.done();
});

async function loadLayoutMiddleware(route: any): Promise<void> {
  let layoutComponent: any;
  try {
    layoutComponent = await import(`@/layouts/${route.meta.layout}.vue`);
  } catch (e) {
    console.error("Error occurred in processing of layout: ", e);
    console.log("Mounted default layout `Blank`");
    layoutComponent = await import(`@/layouts/Blank.vue`);
  }

  route.meta.layoutComponent = layoutComponent.default;
}

export default router;
