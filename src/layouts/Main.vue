<template>
  <div>
    <div style="padding:10px">
      <slot></slot>
    </div>
    <botton-navigation class="bot-van"></botton-navigation>
    <Self-Loading :show="loading"></Self-Loading>
    <Self-Alert :msg="alert.msg" :show.sync="alert.show" :type="alert.type"></Self-Alert>
    <Self-User :user.sync="thatUser" v-if="showUser"></Self-User>
  </div>
</template>

<script>
import BottonNavigation from '../components/BottonNavigation.vue'
import SelfLoading from '../components/Loading.vue'
import SelfAlert from '../components/Alert.vue'
import SelfUser from '../components/User.vue'

export default {
  components: {
    BottonNavigation,
    SelfLoading,
    SelfAlert,
    SelfUser
  },
  props: {
    alert: {
      type: Object,
      required: true
    },
    loading: {
      type: Boolean,
      required: true
    },
    user: {
      type: Object,
      required: true
    },
    initUser: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      thatLoading: this.loading,
      thatUser: this.user,
      showUser: true
    }
  },
  watch: {
    thatLoading(val) {
      this.$emit('update:loading', val)
    },
    thatUser(val) {
      this.$emit('update:user', val)
    },
    initUser(val) {
      var that = this
      this.showUser = false
      setTimeout(() => {
        that.showUser = true
      }, 50)
    }
  }
}
</script>

<style scoped>
.container {
  padding-right: 0px;
  padding-left: 0px;
}

.bot-van {
  position: fixed;
  bottom: 0px;
}

body,
html {
  overflow-x: hidden;
  background-color: #eee;
}
</style>
