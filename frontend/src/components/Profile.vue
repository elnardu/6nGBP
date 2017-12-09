<template>
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-narrow is-half">
        <Spinner v-if="loading"/>
        <User v-else :user="me"/>
      </div>
    </div>
  </div>
</template>

<script>
import Spinner from '@/components/Spinner'
import User from '@/components/UserView'
export default {
  components: {
    Spinner,
    User
  },
  data() {
    return {
      loading: true,
      me: {}
    };
  },
  mounted() {
    this.fetch()
  },
  methods: {
    fetch() {
      this.loading = true;
      this.$http.post("/api/graphql", {
        query:
          `query {
            me {
              actions {
                addedBy {
                  fullname
                  id
                }
                date
                id
                points
                category {
                  id
                  title
                }
                comment
              }

              admin
              fullname
              points
              pointsSpent
              login
            }
          }`
      }).then(res => {
        if (res.body.errors) {
          this.$emit('error', res.body.errors)
          return
        }
        this.$set(this, 'me', res.body.data.me)
        this.loading = false
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
