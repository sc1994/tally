// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import App from './App';
import router from './router';
import MuseUI from 'muse-ui';
import 'muse-ui/dist/muse-ui.css';
import Store from './store';
import 'typeface-roboto';
import Axios from 'axios'
import Toast from "muse-ui-toast";
import 'muse-ui-loading/dist/muse-ui-loading.css'; // load css
import Loading from 'muse-ui-loading';
import Loadsh from 'lodash.uniqby'

Vue.use(MuseUI)
Vue.use(Toast)
Vue.use(Loading)
Axios.defaults.baseURL = 'http://localhost';

Vue.prototype.$axios = Axios;
Vue.prototype.$_ = Loadsh;

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store: Store,
  components: { App },
  template: '<App/>'
})