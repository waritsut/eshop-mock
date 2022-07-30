export default {
  productPaginations(state) {
    return state.productPaginations;
  },
  hasProducts(state) {
    return state.productPaginations && state.productPaginations.length > 0;
  },
};
