<template>
  <div class="">
    <label class="label">Search Categories</label>
    <div class="field has-addons is-marginless">
      <div class="control is-expanded">
        <input class="input" @keydown.enter.prevent="selectFirst" type="text" placeholder="Title" v-model="search"/>
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
          <a class="dropdown-item" v-for="(c, i) in categoriesShowing" :key="i" @click="select(c)">
            {{ c.title }}
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
      categories: [],
      categoriesShowing: [],
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
            categories {
              id
              title
              points
            }
          }`
      }).then(res => {
        this.categories = res.body.data.categories
      })
    },
    selectFirst() {
      if (this.categoriesShowing.length === 0) return
      this.select(this.categoriesShowing[0])
    },
    select(category) {
      this.search = ""
      this.$emit('select', category)
    }
  },
  watch: {
    search() {
      if (this.search == "") {
        this.categoriesShowing = []
        this.isMenuOpen = false
        return
      }
      let search = this.search.toLowerCase()
      this.categoriesShowing = this.categories.filter((category) => {
        return category.title.toLowerCase().includes(search)
      }).slice(0, 10)

      if (this.categoriesShowing.length == 0) this.isMenuOpen = false
      else this.isMenuOpen = true
    }
  }
}
</script>

<style scoped>
.box {

}
</style>
