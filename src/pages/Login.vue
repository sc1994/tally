<template>
  <mu-container>
    <mu-paper class="head-paper" circle :z-depth="3">
      <img src="/src/static/head-default.png" class="head-img" />
    </mu-paper>
    <mu-tabs :value.sync="active" inverse color="#ff4081" text-color="rgba(0, 0, 0, .54)" full-width>
      <mu-tab>登 陆</mu-tab>
      <mu-tab>注 册</mu-tab>
    </mu-tabs>
    <mu-form ref="form" :model="loginModel" style="height: 28px;margin-top: 36px;">
      <mu-form-item label="用户名" prop="username" :rules="usernameRules">
        <mu-text-field v-model="loginModel.username" :error-text="nameErrorText" prop="username" @blur="chenkName"></mu-text-field>
      </mu-form-item>
      <mu-form-item label="密码" prop="password1" :rules="passwordRules">
        <mu-text-field type="password" v-model="loginModel.password1" prop="password"></mu-text-field>
      </mu-form-item>
      <mu-form-item label="确认密码" prop="password2" :rules="passwordRules" v-if="active == 1">
        <mu-text-field type="password" :error-text="pwdErrorText" v-model="loginModel.password2" prop="password"></mu-text-field>
      </mu-form-item>
      <mu-form-item prop="remember" v-if="active == 0">
        <mu-checkbox label="记住密码" v-model="loginModel.remember"></mu-checkbox>
      </mu-form-item>
      <mu-form-item>
        <mu-button round full-width color="primary" style="width: 95%;" @click="submit">
          {{active == 1 ? "注 册" : "登 陆"}}
        </mu-button>
      </mu-form-item>
    </mu-form>
    <mu-snackbar :color="color.color" :open.sync="color.open">
      <mu-icon left :value="icon"></mu-icon>
      {{color.message}}
      <mu-button flat slot="action" color="#fff" @click="color.open = false">Close</mu-button>
    </mu-snackbar>
    <loading :show="loading"></loading>
  </mu-container>
</template> 

<script>
import routes from '../routes'
import loading from '../components/Loading.vue'

export default {
  components: {
    loading
  },
  data() {
    return {
      loading: false,
      active: 0,
      pwdErrorText: '',
      nameErrorText: '',
      loginModel: {
        username: '',
        password1: '',
        password2: '',
        remember: false
      },
      color: {
        color: '',
        message: '',
        open: false,
        timeout: 3000
      },
      usernameRules: [
        {
          validate: val => !!val,
          message: '必须填写用户名'
        },
        {
          validate: val => val.length >= 3,
          message: '用户名长度大于3'
        }
      ],
      passwordRules: [
        {
          validate: val => !!val,
          message: '必须填写密码'
        },
        {
          validate: val => val.length >= 3 && val.length <= 10,
          message: '密码长度大于3小于10'
        }
      ]
    }
  },
  methods: {
    submit() {
      // 注册
      var that = this
      this.$refs.form.validate().then(result => {
        if (!result) return
        if (
          that.active == 1 &&
          that.loginModel.password1 != that.loginModel.password2
        ) {
          that.pwdErrorText = '两次输入的密码不一致'
          return
        }
        if (that.nameErrorText.length > 0) {
          return
        }
        that.loading = true
        that.pwdErrorText = ''
        let url = that.active == 0 ? '/signin' : '/signup'
        // 发送请求
        axios
          .post(url, {
            name: that.loginModel.username,
            password: that.loginModel.password1,
            remember: that.loginModel.remember
          })
          .then(result => {
            var data = result.data
            if (data) {
              if (that.active == 1) {
                if (data.result) {
                  that.openMsg('success', '注册成功，请登陆')
                  that.active = 0
                } else {
                  that.openMsg('error', '注册失败，请检查用户名是否重复')
                }
                that.loading = false
                return
              }
              if (data.user && data.user.name == that.loginModel.username) {
                that.openMsg('success', '登陆成功')
                localStorage.setItem('token', data.token)
                setTimeout(() => {
                  this.$root.currentRoute = '/'
                  window.history.pushState(null, routes['/'], '/')
                }, 800)
                that.loading = false
                return
              }
            }
            that.openMsg('error', '账号或密码错误')
            that.loading = false
          })
          .catch(err => {
            that.openMsg('error', '系统异常')
            console.log(err)
            that.loading = false
          })
      })
    },
    openMsg(color, msg) {
      var that = this
      that.color.message = msg
      that.color.color = color
      that.color.open = true
      setTimeout(() => {
        that.color.open = false
      }, that.color.timeout)
    },
    chenkName() {
      var that = this
      if (that.active == 0 || that.loginModel.username.length <= 3) {
        that.nameErrorText = ''
        return
      }
      that.loading = true
      axios
        .get('/signupcheck/' + that.loginModel.username)
        .then(result => {
          var data = result.data
          if (data) {
            if (data.exist) {
              that.nameErrorText = '已存在的用户名'
            } else {
              that.nameErrorText = ''
            }
          } else {
            that.openMsg('error', '网络异常')
          }
          that.loading = false
        })
        .catch(err => {
          that.openMsg('error', '系统异常')
          console.log(err)
          that.loading = false
        })
    }
  },
  computed: {
    icon() {
      return {
        success: 'check_circle',
        info: 'info',
        warning: 'priority_high',
        error: 'warning'
      }[this.color.color]
    }
  },
  watch: {
    active(val) {
      this.pwdErrorText = ''
      this.nameErrorText = ''
      this.loginModel = {
        username: '',
        password1: '',
        password2: ''
      }
      if (val === 1) {
        this.chenkName()
        if (this.loginModel.password1 != this.loginModel.password2) {
          this.pwdErrorText = '两次输入的密码不一致'
          return
        }
      } else {
        this.nameErrorText = ''
      }
      setTimeout(() => {
        this.$refs.form.clear()
      }, 20)
    }
  }
}
</script>

<style scoped>
.head-paper {
  height: 120px;
  width: 120px;
  margin: 35px auto;
}

.head-img {
  width: 110px;
  margin-left: 5px;
  margin-top: 5px;
}

.loader {
  width: 100%;
  height: 100%;
  background-color: rgba(120, 144, 156, 0.5);
  float: left;
  top: 0px;
  left: 0px;
  display: -webkit-flex;
  justify-content: center;
  display: flex;
  align-items: center;
  z-index: 9999;
  position: fixed;
}
</style>
