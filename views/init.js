import Axios from 'axios';
import Vue from 'vue';
import Linq from 'linq';
import Toast from "muse-ui-toast";

// axios 初始化----------------------------------------------------------
if (location.href.endsWith("test")) {
  Axios.defaults.baseURL = 'http://localhost';
} else {
  Axios.defaults.baseURL = 'http://118.24.27.231:8888';
}

// 添加请求拦截器
Axios.interceptors.request.use(function (config) {
  config.headers = {
    "Content-Type": "application/json",
    "token": localStorage.getItem("token")
  }
  return config;
}, function (error) {
  // 对请求错误做些什么
  Toast.error("系统错误");
  console.log(error)
  return Promise.reject(error);
});

// 添加响应拦截器
Axios.interceptors.response.use(function (response) {
  if (response.data.code != 0) {
    Toast.warning(response.data.msg)
  } else {
    if (response.data.data && response.data.data.token) {
      localStorage.setItem("token", response.data.data.token)
    }
  }
  return response.data;
}, function (error) {
  Toast.error("系统错误");
  console.log(error)
  return Promise.reject(error);
});


// vue 初始化------------------------------------------------------------
Vue.prototype.$fileUrl = "http://118.24.27.231:81/uploadfile"
Vue.prototype.$axios = Axios;
Vue.prototype.$linq = Linq.from;
// 一些通用方法
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
  var decimals = 1;
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
