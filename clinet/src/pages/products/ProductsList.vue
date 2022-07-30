<template>
  <div>
    <section>
      <product-filter></product-filter>
    </section>
    <section>
      <base-card>
        <div v-if="isLoading">
          <base-spinner></base-spinner>
        </div>

        <div v-else-if="hasProducts" class="row">
          <div v-for="product in getProducts" :key="product.id" class="column">
            <product-item
              :name="product.name"
              :img="product.imageUrl"
              :price="product.price"
            ></product-item>
          </div>
        </div>

        <h3 v-else>No products found.</h3>
      </base-card>
    </section>
  </div>
</template>

<script>
import ProductItem from "../../components/products/ProductItem.vue";
import ProductFilter from "../../components/products/ProductFilter.vue";

export default {
  components: {
    ProductItem,
    ProductFilter,
  },

  data() {
    return {
      isLoading: false,
      error: null,
    };
  },

  computed: {
    hasProducts() {
      return (
        !this.isLoading && this.$store.getters["productPaginate/hasProducts"]
      );
    },
    getProducts() {
      return this.$store.getters["productPaginate/productPaginations"];
    },
  },

  methods: {
    async loadProductPaginations(refresh = false) {
      this.isLoading = true;
      try {
        await this.$store.dispatch("productPaginate/loadProductPaginations", {
          forceRefresh: refresh,
        });
      } catch (error) {
        this.error = error.message || "Something went wrong!";
      }
      this.isLoading = false;
    },
  },
  
  created() {
    this.loadProductPaginations();
  },
};
</script>

<style scoped>
.controls {
  display: flex;
  justify-content: right;
}

.column {
  float: left;
  width: 25%;
  padding: 0 10px;
  margin-top: 20px;
}

.row {
  margin: 0 -5px;
}

.row:after {
  content: "";
  display: table;
  clear: both;
}

@media screen and (max-width: 600px) {
  .column {
    width: 100%;
    display: block;
    margin-bottom: 20px;
  }
}
</style>
