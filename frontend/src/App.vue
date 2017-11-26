<template>
  <div id="app">
    <Navbar :logged="logged" @logout="logout"/>
    <router-view @token="setToken"/>
  </div>
</template>

<script>
import Navbar from "./components/Navbar";
export default {
  name: "app",
  components: {
    Navbar
  },
  data() {
    return {
      logged: false,
      token: ""
    };
  },
  mounted() {
    if (localStorage.token) {
      this.$http
        .post("/graphql", {
          query: `{ me { fullname } }`
        }, {
          headers: {
            "Authorization": "bearer " + localStorage.token
          }
        })
        .then(res => {
          if (res.body.data) {
            this.setToken(localStorage.token)
          }
        })
    }
  },
  methods: {
    setToken(token) {
      localStorage.token = token;
      this.logged = true;
      this.token = token;
      this.$http.interceptors.push(function(request, next) {
        request.headers.set("Authorization", "bearer " + token)
        next()
      })
    },
    logout() {
      localStorage.token = ""
      this.logged = false
      this.token = ""
      this.$http.interceptors.pop()
    }
  }
};
</script>

<style lang="scss">
@import "node_modules/bulma/bulma";
#app {
  // width: 100vw;
  // height: 100vh;
  // margin: 0;
  // padding: 0;
}
</style>
