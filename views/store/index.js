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
      that.$axios.get("/user/get")
        .then(response => {
          if (response.code == 0) {
            response.data.consumes = that.$linq(response.data.consumes)
              .orderByDescending(x => x.count)
              .thenByDescending(x => new Date(x.utime))
              .thenByDescending(x => new Date(x.ctime))
              .toArray();
            commit("changeUser", response.data)
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
        .post("/message/add", {
          tid: data.tid,
          fid: this.state.currentUser.id,
          content: data.content,
          needTouch: data.needTouch,
          type: data.type,
          status: 1
        })
        .then(response => {
          if (response.code == 0) {
            that.$toast.success("已发送");
          }
          loading.close();
        })
    },
    initUnreadNumber({
      commit
    }) {
      if (!this.state.currentUser.id) return
      var that = this._vm;
      that.$axios.get("/message/getcount/1").then(response => {
        if (response.code == 0) {
          commit("changeUnreadNumber", response.data)
        }
      });
    }
  }
})
