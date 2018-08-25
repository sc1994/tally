import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    currentUser: {},
    appbarStyle: ""
  },
  mutations: {
    changeUser(state, data) {
      state.currentUser = data
    },
    changeAppbarStyle(state, data) {
      state.appbarStyle = data
    }
  },
  actions: {
    initUser({
      commit
    }, data) {
      var that = this._vm;
      that.$axios.get("/getuser/" + localStorage.getItem("token"))
        .then(result => {
          if (result.data.result) {
            commit("changeUser", result.data.user)
          } else {
            that.$toast.warning("token失效,稍后将跳转到登陆页面")
            if (result.data)
              setTimeout(() => {
                data.$router.push({
                  path: "/sign"
                });
              }, 1200)
          }
        })
        .catch(error => {
          console.log(error)
          that.$toast.error("网络异常, 请重试~")
        })
    },
    sendMessage({
      commit
    }, data) {
      debugger
      var that = this._vm;
      var loading = that.$loading({});
      that.$axios
        .post("/sendmessage", {
          tid: data.tid,
          fid: this.state.currentUser.id,
          fnick: this.state.currentUser.nick,
          fimg: this.state.currentUser.headImg,
          content: data.content,
          needTouch: data.needTouch,
          type: data.type
        })
        .then(response => {
          if (response.data.result) {
            that.$toast.success("已发送");
          } else {
            that.$toast.error("网络异常,请重试");
          }
          loading.close();
        })
        .catch(error => {
          that.$toast.error("网络异常,请重试");
          loading.close();
        });
    }
  }
})
