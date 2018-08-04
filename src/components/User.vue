<template>

</template>

<script>
import routes from '../routes'

export default {
  props: {
    user: {
      type: Object,
      required: true
    }
  },
  methods: {
    init() {
      console.log('初始化用户信息')
      var that = this
      var token = localStorage.getItem('token')
      axios
        .get('/getuser/' + token)
        .then(result => {
          var data = result.data
          if (data && data.result) {
            that.$emit('update:user', data.user)
          } else {
            alert('token失效')
            that.$root.currentRoute = '/login'
            window.history.pushState(null, routes['/login'], '/login')
          }
        })
        .catch(err => {
          alert('数据异常')
          console.log(err)
        })
    }
  },
  mounted() {
    this.init()
  }
}
</script>

