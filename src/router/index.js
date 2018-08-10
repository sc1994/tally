import Vue from 'vue'
import Router from 'vue-router'
import Sign from '@/pages/Sign'
import Tally from '@/pages/Tally'
import Me from '@/pages/Me'
// import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/sign',
      name: 'Sign',
      component: Sign
    },
    {
      path: '/',
      name: 'Tally',
      component: Tally
    },
    {
      path: "/me",
      name: 'Me',
      component: Me
    }
  ]
})
