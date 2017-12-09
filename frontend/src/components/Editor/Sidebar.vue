<template>
<div class="menu">
  <p class="menu-label">
    General
  </p>
  <ul class="menu-list">
    <li><a @click="change('points')" :class="{'is-active': value === 'points'}">Add action</a></li>
    <li><a @click="change('category')" :class="{'is-active': value === 'category'}">Add category</a></li>
  </ul>
</div>
</template>

<script>
export default {
  props: ['value'],
  methods: {
    change(tab) {
      this.$emit('input', tab)
    }
  }

}
</script>

<style>
.menu {
  padding: 0.5rem;
}
</style>
