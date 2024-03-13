import { useRouter, useRoute } from "vue-router";

export const mainMenu = () => {
  const route = useRoute();
  const router = useRouter();

  const mainSection = (route.path as string).startsWith("/_")
    ? "admin"
    : "manager";
  const sections =
    router.options.routes.find((e) => e.name === mainSection)?.children || [];
  return sections.reduce((acc, route) => {
    if (route.meta?.layout === "Private") {
      acc.push({
        name: route.meta.name,
        link: { name: route.name },
        icon: route.meta.icon,
      });
    }
    return acc;
  }, []);
};

export const secondMenu = (section: string) => {
  const route = useRoute();
  const router = useRouter();

  const mainSection = (route.path as string).startsWith("/_")
    ? "admin"
    : "manager";
  const sections =
    router.options.routes.find((e) => e.name === mainSection)?.children || [];
  return sections.reduce((acc, route) => {
    if (route.path === section) {
      route.children.find((e) => {
        acc.push({
          name: e.meta.name,
          link: { name: e.name },
        });
      });
    }
    return acc;
  }, []);
};
