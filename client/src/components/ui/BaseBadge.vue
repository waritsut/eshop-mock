<template>
  <span class="badge">
    {{ text }}
  </span>
</template>

<script>
export default {
  props: ["text"],
};
</script>

<style scoped>
.badge {
  background-color: #ccc;
  color: #252525;
  border-radius: 30px;
  padding: 0.5rem 1.5rem;
  display: inline-block;
  margin-right: 0.5rem;
  background-color: #8d006e;
  color: white;
  height: 35px;
}
</style>
