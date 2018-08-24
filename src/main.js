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
import Linq from 'linq'

Vue.use(MuseUI)
Vue.use(Toast)
Vue.use(Loading)

// Axios.defaults.baseURL = 'http://192.168.1.5';
Axios.defaults.baseURL = 'http://suncheng.xyz:8888';

Vue.prototype.$axios = Axios;
Vue.prototype.$linq = Linq.from;

Vue.prototype.$format = (date, fmt) => { // 时间格式化
  var that = new Date(date)
  var o = {
    "M+": that.getMonth() + 1, //月份
    "d+": that.getDate(), //日
    "h+": that.getHours(), //小时
    "m+": that.getMinutes(), //分
    "s+": that.getSeconds(), //秒
    "q+": Math.floor((that.getMonth() + 3) / 3), //季度
    "S": that.getMilliseconds() //毫秒
  };
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (that.getFullYear() + "").substr(4 - RegExp.$1.length));
  for (var k in o)
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
  return fmt;
};

Vue.prototype.$numberFormat = (number) => {
  var dec_point = "."
  var thousands_sep = ","
  var decimals = 2;
  number = (number + '').replace(/[^0-9+-Ee.]/g, '');
  var n = !isFinite(+number) ? 0 : +number,
    prec = !isFinite(+decimals) ? 0 : Math.abs(decimals),
    sep = (typeof thousands_sep === 'undefined') ? ',' : thousands_sep,
    dec = (typeof dec_point === 'undefined') ? '.' : dec_point,
    s = '',
    toFixedFix = function (n, prec) {
      var k = Math.pow(10, prec);
      return '' + Math.floor(n * k) / k;
    };
  s = (prec ? toFixedFix(n, prec) : '' + Math.floor(n)).split('.');
  var re = /(-?\d+)(\d{3})/;
  while (re.test(s[0])) {
    s[0] = s[0].replace(re, "$1" + sep + "$2");
  }

  if ((s[1] || '').length < prec) {
    s[1] = s[1] || '';
    s[1] += new Array(prec - s[1].length + 1).join('0');
  }
  return s.join(dec);
};

Vue.config.productionTip = false

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
