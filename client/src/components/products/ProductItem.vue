<template>
  <div class="card">
    <img :src="img" :alt="name" width="170" height="150" />
    <h4>{{ name }}</h4>
    <base-badge :text="bahtPrice"></base-badge>
    <div></div>
  </div>
</template>

<script>
export default {
  props: ["id", "name", "img", "price"],
  computed: {
    bahtPrice() {
      return "฿" + this.price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },
  },
};
</script>

<style scoped>
h3,
h4 {
  font-size: 1rem;
}
h3,
h4 {
  margin: 0.5rem 0;
}
.card {
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  padding: 16px;
  text-align: center;
  /* background-color: #f1f1f1; */
}
</style>
