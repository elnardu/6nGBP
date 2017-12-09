<template>
  <form class="action-editor form" @submit.prevent="submit">
    <UserSearch @select="addUser"/>

    <label class="label">Selected users:</label>
    <div class="field is-grouped is-grouped-multiline users-list">
      <div class="control" v-for="(u, i) in users" :key="i">
        <div class="tags has-addons">
          <span class="tag is-danger is-medium">{{u.fullname}}</span>
          <a class="tag is-delete is-medium" @click="removeUser(u.id)"></a>
        </div>
      </div>
      <div v-if="users.length === 0" class="has-text-grey-light">None</div>
    </div>
    <hr/>
    <div class="field has-addons">
      <p class="control">
        <a class="button" :class="{'is-active': isCategoryActive}" @click="isCategoryActive = true">
          <span>Set Category</span>
        </a>
      </p>
      <p class="control">
        <a class="button" :class="{'is-active': !isCategoryActive}" @click="isCategoryActive = false">
          <span>Set Points</span>
        </a>
      </p>
    </div>

    <div v-if="isCategoryActive" class="field">
      <CategoriesSearch @select="setCategory"/>
      <p class="label" v-if="category.title">Selected category: </p>
      <div class="tags has-addons" v-if="category.title">
        <span class="tag is-danger">{{ category.title }}</span>
        <span class="tag">{{ category.points }}</span>
      </div>
    </div>
    <div v-else class="field">
      <div class="field">
        <label class="label">Points</label>
        <div class="control">
          <input class="input" type="number" v-model.number="points" placeholder="Points">
        </div>
      </div>
    </div>
    <hr/>
    <div class="field">
      <label class="label">Comment</label>
      <div class="control">
        <textarea class="textarea" placeholder="Comment" v-model="comment"></textarea>
      </div>
    </div>

    <div class="control">
      <button class="button is-danger" :class="{'is-loading': loading}" type="submit">Submit</button>
    </div>
    <br/>
    <transition name="fade" mode="out-in">
      <div v-if="error" class="notification is-warning">{{ error }}</div>
      <div v-if="info" class="notification is-info">{{ info }}</div>
    </transition>
  </form>
</template>

<script>
import UserSearch from "./UserSearch";
import CategoriesSearch from "./CategoriesSearch";
export default {
  components: {
    UserSearch,
    CategoriesSearch
  },
  data() {
    return {
      users: [],
      category: {},
      points: 0,
      comment: "",
      isCategoryActive: true,

      loading: false,
      error: "",
      info: ""
    };
  },
  methods: {
    addUser(user) {
      for (let i = 0; i < this.users.length; i++) {
        if (this.users[i].id === user.id) return;
      }
      this.users.push(user);
    },
    removeUser(userId) {
      for (let i = 0; i < this.users.length; i++) {
        if (this.users[i].id === userId) {
          this.users.splice(i, 1);
          return;
        }
      }
    },
    setCategory(category) {
      this.category = category
    },
    submit() {
      if (this.users.length === 0) {
        this.error = "No users selected"
        return
      }

      this.loading = true
      this.error = ""
      this.info = ""
      this.users.forEach(u => {
        this.createAction(u.id)
      })
    },
    createAction(userId) {
      this.$http.post("/api/graphql", {
        query:
          `mutation ($user: ID!, $category: ID, $points: Int, $comment: String) {
            createAction(user: $user, category: $category, points: $points, comment: $comment) {
              id
            }
          }`,
        variables: {
          user: userId,
          category: this.isCategoryActive ? this.category.id : null,
          points: this.isCategoryActive ? null : this.points,
          comment: this.comment
        }
      }).then(res => {
        this.loading = false
        if (res.body.errors) {
          console.error(res.body.errors)
          this.error = res.body.errors[0].message
          return
        }
        this.info = `Success`
      })
    }
  }
};
</script>

<style scoped>
.action-editor {
  padding: 1rem;
}
.users-list {
  padding-bottom: 1rem;
}
</style>
