import { createStore } from "vuex";

import productsModule from "./modules/products/index";

export default createStore({
  modules: {
    productPaginate: productsModule,
  },
});
