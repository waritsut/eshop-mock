import { createRouter, createWebHistory } from "vue-router";

import ProductsList from "../pages/products/ProductsList.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", redirect: "/catalogs" },
    { path: "/catalogs", component: ProductsList },
  ],
});

export default router;
