import Vue from 'vue'
import Router from 'vue-router'
import Sign from '@/pages/Sign'
// import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/sign',
      name: 'Sign',
      component: Sign
    }
  ]
})
