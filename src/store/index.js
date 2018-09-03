import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    currentUser: {},
    appbarStyle: "",
    unreadNumber: 0
  },
  mutations: {
    changeUser(state, data) {
      state.currentUser = data
    },
    changeAppbarStyle(state, data) {
      state.appbarStyle = data
    },
    changeUnreadNumber(state, data) {
      state.unreadNumber = data
    }
  },
  actions: {
    initUser({
      commit
    }, data) {
      var that = this._vm;
      var vuex = this;
      that.$axios.get("/getuser/" + localStorage.getItem("token"))
        .then(response => {
          if (response.result) {
            commit("changeUser", response.user)
            vuex.dispatch("initUnreadNumber");
          } else {
            that.$toast.warning("token失效,稍后将跳转到登陆页面")
            if (response)
              setTimeout(() => {
                data.$router.push({
                  path: "/sign"
                });
              }, 1200)
          }
        })
    },
    sendMessage({
      commit
    }, data) {
      var that = this._vm;
      var loading = that.$loading({});
      that.$axios
        .post("/sendmessage", {
          tid: data.tid,
          fid: this.state.currentUser.id,
          content: data.content,
          needTouch: data.needTouch,
          type: data.type
        })
        .then(response => {
          if (response.result) {
            that.$toast.success("已发送");
          } else {
            that.$toast.error("网络异常,请重试");
          }
          loading.close();
        })
    },
    initUnreadNumber({
      commit
    }) {
      if (!this.state.currentUser.id) return
      var that = this._vm;
      that.$axios.get("/getmessageunreadcount/" + this.state.currentUser.id).then(response => {
        commit("changeUnreadNumber", response.result)
      });
    }
  }
})
