<template>
  <mu-container>
    <mu-bottom-nav :value.sync="value" shift v-on:change="goto">
      <mu-bottom-nav-item value="/" title="添加一笔" icon="add_shopping_cart"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/bill" title="查账" icon="show_chart"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/me" title="我的" icon="account_box"></mu-bottom-nav-item>
    </mu-bottom-nav>
  </mu-container>
</template>

<script>
import routes from '../routes'

export default {
  data() {
    return {
      value: '/'
    }
  },
  methods: {
    goto(val) {
      event.preventDefault()
      this.$root.currentRoute = val
      window.history.pushState(null, routes[val], val)
      this.value = val
    }
  },
  mounted() {
    this.value = this.$root.currentRoute
  },
}
</script>

<style scoped>
.active {
  color: cornflowerblue;
}
</style>
