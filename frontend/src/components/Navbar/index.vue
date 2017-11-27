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
        <router-link to="/dashboard" class="navbar-item" active-class="is-active">Dashboard</router-link>
        <router-link v-if="!logged" to="/login" class="navbar-item" active-class="is-active">Login</router-link>
        <router-link v-if="!logged" to="/register" class="navbar-item" active-class="is-active">Register</router-link>
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link">
            Other
          </a>
          <div class="navbar-dropdown is-right">
            <router-link class="navbar-item" to="/apiexplorer" active-class="is-active">
              Public API
            </router-link>
            <a href="//github.com/elnardu/6nGBP" class="navbar-item" target="_blank">Source Code</a>
          </div>
        </div>
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
      this.$emit("logout");
    }
  }
};
</script>

<style scoped>

</style>

