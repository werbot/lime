import { createRouter, createWebHistory } from "vue-router";
import { getCookie } from "@/utils";
import * as NProgress from "nprogress";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "manager",
      redirect: { name: "manager-license" },
      children: [
        {
          path: "licenses",
          name: "manager-license",
          meta: { name: "Licenses", icon: "ticket", layout: "Private" },
          component: () => import("@/pages/manager/License.vue"),
        },
        {
          path: "payment",
          name: "manager-payment",
          meta: { name: "Payments", icon: "banknotes", layout: "Private" },
          component: () => import("@/pages/manager/Payment.vue"),
        },
        {
          path: "setting",
          name: "manager-setting",
          meta: { name: "Settings", icon: "tooth", layout: "Private" },
          component: () => import("@/pages/manager/Setting.vue"),
        },
      ],
    },
    {
      path: "/signin",
      name: "signin",
      meta: { layout: "Signin" },
      component: () => import("@/pages/Signin.vue"),
    },
    {
      path: "/_",
      name: "admin",
      redirect: { name: "admin-license" },
      children: [
        {
          path: "install",
          name: "admin-install",
          meta: { layout: "Blank" },
          component: () => import("@/pages/admin/Install.vue"),
        },
        {
          path: "signin",
          name: "admin-signin",
          meta: { layout: "Signin" },
          component: () => import("@/pages/admin/Signin.vue"),
        },
        {
          path: "license",
          name: "admin-license",
          meta: { name: "Licenses", icon: "ticket", layout: "Private" },
          redirect: { name: "admin-license-all" },
          children: [
            {
              path: "all",
              name: "admin-license-all",
              meta: { name: "All" },
              component: () => import("@/pages/admin/License/All.vue"),
            },
            {
              path: "all",
              name: "admin-license-template",
              meta: { name: "Template" },
              component: () => import("@/pages/admin/License/Template.vue"),
            },
          ],
        },
        {
          path: "customer",
          name: "admin-customer",
          meta: { name: "Customers", icon: "users", layout: "Private" },
          component: () => import("@/pages/admin/Customer.vue"),
        },
        {
          path: "payment",
          name: "admin-payment",
          meta: { name: "Payments", icon: "banknotes", layout: "Private" },
          component: () => import("@/pages/admin/Payment.vue"),
        },
        {
          path: "setting",
          name: "admin-setting",
          meta: { name: "Settings", icon: "tooth", layout: "Private" },
          redirect: { name: "admin-setting-site" },
          children: [
            {
              path: "site",
              name: "admin-setting-site",
              meta: { name: "Site" },
              component: () => import("@/pages/admin/Setting/Site.vue"),
            },
            {
              path: "payment",
              name: "admin-setting-payment",
              meta: { name: "Payment" },
              component: () => import("@/pages/admin/Setting/Payment.vue"),
            },
            {
              path: "mail",
              name: "admin-setting-mail",
              meta: { name: "Mail" },
              component: () => import("@/pages/admin/Setting/Mail.vue"),
            },
          ],
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
  if (getCookie("manager")) {
    isAuthenticated = true;
  }

  let isAdmin = false;
  if (getCookie("admin")) {
    isAdmin = true;
  }

  if (to.path.startsWith("/_")) {
    if (!isAdmin && to.name !== "admin-signin") {
      return { name: "admin-signin" };
    }
    if (isAdmin && to.name === "admin-signin") {
      return { name: "admin-license-all" };
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
