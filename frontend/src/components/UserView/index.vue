<template>
  <div>
    <h1 class="is-size-1 has-text-danger ">{{ user.fullname }}</h1>
    <div class="level">
      <div class="level-left">
        <div class="level-item">
          <h2 class="is-size-5 has-text-grey-light">@{{ user.login }}</h2>
        </div>
        <div v-if="user.admin" class="level-item">
          <span class="tag is-danger">admin</span>
        </div>
      </div>
    </div>
    <hr/>
    <div class="level points">
      <div class="level-item has-text-centered">
        <div>
          <p class="heading">GBP Earned</p>
          <p class="title">{{ user.points }}</p>
        </div>
      </div>
      <div class="level-item has-text-centered">
        <div>
          <p class="heading">GBP Spent</p>
          <p class="title">{{ user.pointsSpent }}</p>
        </div>
      </div>
      <div class="level-item has-text-centered">
        <div>
          <p class="heading">Total GBP</p>
          <p class="title">{{ user.points }}</p>
        </div>
      </div>
    </div>
    <hr/>

  </div>
</template>

<script>
export default {
  props: ["user"]
};
</script>

<style scoped>
.level .points {

}
</style>
