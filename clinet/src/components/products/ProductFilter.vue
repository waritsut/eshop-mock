<template>
  <base-card
    ><h2>I'm looking for ...</h2>
    <span class="filter-option">
      <input type="checkbox" id="1" checked @change="setFilter" />
      <label for="furniture">Furniture</label>
    </span>
    <span class="filter-option">
      <input type="checkbox" id="2" checked @change="setFilter" />
      <label for="electric-appliance">Electric Appliance</label>
    </span>
    <span class="filter-option">
      <input type="checkbox" id="3" checked @change="setFilter" />
      <label for="clothes">Clothes</label>
    </span>
    <span class="filter-option">
      <input type="checkbox" id="4" checked @change="setFilter" />
      <label for="sport">Sport</label>
    </span>
    <span class="filter-option">
      <input type="checkbox" id="4" checked @change="setFilter" />
      <label for="accessory">Accessory</label>
    </span>
  </base-card>
</template>

<script>
export default {};
</script>

<style scoped>
h2 {
  margin: 0.5rem 0;
}
</style>
