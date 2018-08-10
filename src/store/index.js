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
        initUser({ commit }, data) {
            var that = this._vm;
            that.$axios.get("/getuser/" + localStorage.getItem("token"))
                .then(result => {
                    if (result.data.result) {
                        commit("changeUser", result.data.user)
                    } else {
                        that.$toast.warning("token失效,稍后将跳转到登陆页面")
                        if (result.data)
                            setTimeout(() => {
                                data.$router.push({ path: "/sign" });
                            }, 1200)
                    }
                })
                .catch(error => {
                    console.log(error)
                    that.$toast.error("网络异常, 请重试~")
                })
        }
    }
})