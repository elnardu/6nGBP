<template>
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-narrow is-half">
        <div v-if="error" class="notification is-warning">{{ error }}</div>
        <form @submit.prevent="register">
          <div class="field">
            <label class="label">Full name</label>
            <div class="control">
              <input class="input" type="text" placeholder="Full name" v-model="fullname">
            </div>
          </div>
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
      error: "",
      fullname: ""
    };
  },
  methods: {
    register() {
      this.loading = true;
      this.$http.post("/api/graphql", {
        query:
          `mutation ($login: String!, $password: String!, $fullname: String!) {
            createUser(login: $login, password: $password, fullname: $fullname) {
              id
            }
          }`,
        variables: {
          login: this.login,
          fullname: this.fullname,
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
        this.$router.push('/login?fr=true')
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
