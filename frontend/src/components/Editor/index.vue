<template>
  <div class="columns">
    <div class="column is-one-quarter"><Sidebar v-model="activeTab"/></div>
    <div class="column">
      <transition name="fade" mode="out-in">
        <ActionEditor v-if="activeTab === 'points'"/>
        <AddCategory v-if="activeTab === 'category'"/>
      </transition>
    </div>
  </div>
</template>

<script>
import Sidebar from './Sidebar'
import ActionEditor from './ActionEditor'
import AddCategory from './AddCategory'
export default {
  components: {
    Sidebar,
    ActionEditor,
    AddCategory
  },
  data() {
    return {
      activeTab: "points"
    }
  }
}
</script>

<style>

</style>
