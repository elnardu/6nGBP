<template>
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-narrow is-half">
        <div v-if="error" class="notification is-warning">{{ error }}</div>
        <form @submit.prevent="doLogin">
          <div class="field">
            <label class="label">Login</label>
            <div class="control">
              <input class="input" type="text" placeholder="Login" v-model="login">
            </div>
          </div>
          <div class="field">
            <label class="label">Password</label>
            <div class="control">
              <input class="input" type="password" placeholder="Password" v-model="password">
            </div>
          </div>
          <div class="control">
            <button class="button is-danger" :class="{'is-loading': loading}" type="submit">Submit</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loading: false,
      login: "",
      password: "",
      error: ""
    };
  },
  methods: {
    doLogin() {
      this.loading = true;
      this.$http.post("/graphql", {
        query:
          `query ($login: String!, $password: String!) {
            auth(login: $login, password: $password)
          }`,
          variables: {
            login: this.login,
            password: this.password
          }
      }).then(res => {
        this.loading = false
        if (res.body.errors) {
          this.$emit('error', res.body.errors)
          console.error(res.body.errors)
          this.error = res.body.errors[0].message
          return
        }
        this.$emit('token', res.body.data.auth)
        this.$router.push('/profile')
      })
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container {
  margin-top: 1em;
  margin-left: 20px;
  margin-right: 20px;
}
</style>
