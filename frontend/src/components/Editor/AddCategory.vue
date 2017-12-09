<template>
  <div class="category">
    <transition name="fade" mode="out-in">
      <div v-if="error" class="notification is-warning">{{ error }}</div>
      <div v-if="info" class="notification is-info">{{ info }}</div>
    </transition>
    <form @submit.prevent="submit">
      <div class="field">
        <label class="label">Title</label>
        <div class="control">
          <input class="input" type="text" placeholder="Title" v-model.trim="title"/>
        </div>
        <p class="help">It is unique</p>
      </div>
      <div class="field">
        <label class="label">Description</label>
        <div class="control">
          <textarea class="textarea" type="text" placeholder="Description" v-model.trim="description"/>
        </div>
      </div>
      <div class="field">
        <label class="label">Points</label>
        <div class="control">
          <input class="input" type="text" placeholder="Title" v-model.number="points"/>
        </div>
        <p class="help">It can be negative</p>
      </div>
      <div class="control">
        <button class="button is-danger" :class="{'is-loading': loading}" type="submit">Submit</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loading: false,
      error: "",
      info: "",

      title: "",
      description: "",
      points: 0
    }
  },
  methods: {
    submit() {
      this.loading = true
      this.error = ""
      this.info = ""
      this.$http.post("/api/graphql", {
        query:
          `mutation ($title: String!, $description: String!, $points: Int!) {
            createCategory(title: $title, description: $description, points: $points) {
              id
            }
          }`,
        variables: {
          title: this.title,
          description: this.description,
          points: this.points,
        }
      }).then(res => {
        this.loading = false
        if (res.body.errors) {
          console.error(res.body.errors)
          this.error = res.body.errors[0].message
          return
        }
        this.info = `New category id: ${res.body.data.createCategory.id}`
      })
    }
  }
};
</script>

<style scoped>
.category {
  padding: 1rem;
}
</style>
