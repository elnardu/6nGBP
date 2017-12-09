<template>
  <div class="">
    <label class="label">Search Users</label>
    <div class="field has-addons is-marginless">
      <div class="control is-expanded">
        <input class="input" @keydown.enter.prevent="selectFirst" type="text" placeholder="Name" v-model="search"/>
      </div>
      <div class="control">
        <button class="button is-info" @click="selectFirst">
          <span class="icon is-small">
            <i class="fa fa-plus"></i>
          </span>
        </button>
      </div>
    </div>
    <div class="dropdown" :class="{'is-active': isMenuOpen}">
      <div class="dropdown-menu">
        <div class="dropdown-content">
          <a class="dropdown-item" v-for="(u, i) in usersShowing" :key="i" @click="select(u)">
            {{ u.fullname }}
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isMenuOpen: false,
      users: [],
      usersShowing: [],
      search: ""
    }
  },
  mounted() {
    this.refresh()
  },
  methods: {
    refresh() {
      this.$http.post("/api/graphql", {
        query:
          `query {
            users {
              id
              fullname
            }
          }`
      }).then(res => {
        this.users = res.body.data.users
      })
    },
    selectFirst() {
      if (this.usersShowing.length === 0) return
      this.select(this.usersShowing[0])
    },
    select(user) {
      this.search = ""
      this.$emit('select', user)
    }
  },
  watch: {
    search() {
      if (this.search == "") {
        this.usersShowing = []
        this.isMenuOpen = false
        return
      }
      let search = this.search.toLowerCase()
      this.usersShowing = this.users.filter((user) => {
        return user.fullname.toLowerCase().includes(search)
      }).slice(0, 10)

      if (this.usersShowing.length == 0) this.isMenuOpen = false
      else this.isMenuOpen = true
    }
  }
}
</script>

<style scoped>
.box {

}
</style>
