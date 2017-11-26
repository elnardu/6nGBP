<template>
  <div class="navbar is-danger">
    <div class="navbar-brand">
      <router-link to="/" class="navbar-item is-size-4">
        6N G🅱️P
      </router-link>
      <button @click="toggleNavbar" class="button navbar-burger is-danger">
        <span></span>
        <span></span>
        <span></span>
      </button>
    </div>
    <div class="navbar-menu" :class="{'is-active': navbarActive}">
      <div class="navbar-end">
        <router-link to="/dashboard" class="navbar-item">Dashboard</router-link>
        <router-link v-if="!logged" to="/login" class="navbar-item">Login</router-link>
        <router-link v-if="!logged" to="/register" class="navbar-item">Register</router-link>
        <a v-if="logged" @click="logout" class="navbar-item">Logout</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["logged"],
  data() {
    return {
      navbarActive: false
    };
  },
  methods: {
    toggleNavbar() {
      this.navbarActive = !this.navbarActive;
    },
    logout() {
      this.$emit('logout')
    }
  }
};
</script>

<style>

</style>

