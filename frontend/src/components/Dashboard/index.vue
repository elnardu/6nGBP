<template>
  <div class="container">
      <div class="columns">
          <div class="column">
              <h1 class="is-size-2 has-text-danger has-text-centered">Current top</h1> 
              <users-top/>
          </div>
          <div class="column">
              <h1 class="is-size-2 has-text-danger has-text-centered">Latest changes</h1> 
              <latest-actions/>
          </div>
      </div>
  </div>
</template>

<script>
import UsersTop from './UsersTop'
import LatestActions from './LatestActions'
export default {
    components: {
        UsersTop,
        LatestActions
    }
}
</script>

<style scoped>
.container {
    margin-top: 2em;
}
h1 {
    margin-bottom: 0.5em;
}
</style>
