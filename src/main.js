import 'muse-ui-loading/dist/muse-ui-loading.css'
import Vue from 'vue'
import routes from './routes'
import MuseUI from 'muse-ui'
import Loading from 'muse-ui-loading'
Vue.use(Loading)

Vue.use(MuseUI)
Vue.use(Loading)
Vue.use(Loading)









const app = new Vue({
  el: '#app',
  data: {
    currentRoute: window.location.pathname
  },
  computed: {
    ViewComponent () {
      const matchingView = routes[this.currentRoute]
      return matchingView
        ? require('./pages/' + matchingView + '.vue')
        : require('./pages/404.vue')
    }
  },
  render (h) {
    return h(this.ViewComponent)
  }
})

window.addEventListener('popstate', () => {
  app.currentRoute = window.location.pathname
})
