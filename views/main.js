// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import App from './App';
import router from './router';
import MuseUI from 'muse-ui';
import 'muse-ui/dist/muse-ui.css';
import Store from './store';
import 'typeface-roboto';
import Toast from "muse-ui-toast";
import 'muse-ui-loading/dist/muse-ui-loading.css';
import Loading from 'muse-ui-loading';
import Helpers from 'muse-ui/lib/Helpers';
import Init from "./init"

Vue.use(MuseUI)
Vue.use(Toast)
Vue.use(Loading)
Vue.use(Helpers)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store: Store,
  components: {
    App
  },
  template: '<App/>'
})
