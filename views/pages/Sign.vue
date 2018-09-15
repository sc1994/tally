<template>
  <mu-container>
    <mu-paper class="head-paper" circle :z-depth="3">
      <img src="/static/images/head-default.png" class="head-img" />
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
  </mu-container>
</template>

<script>
export default {
  data() {
    return {
      active: 0,
      pwdErrorText: "",
      nameErrorText: "",
      loginModel: {
        username: "",
        password1: "",
        password2: "",
        remember: true
      },
      usernameRules: [
        {
          validate: val => !!val,
          message: "必须填写用户名"
        },
        {
          validate: val => val.length >= 3,
          message: "用户名长度大于2"
        }
      ],
      passwordRules: [
        {
          validate: val => !!val,
          message: "必须填写密码"
        },
        {
          validate: val => val.length > 5,
          message: "密码长度大于5"
        }
      ]
    };
  },
  methods: {
    submit() {
      var that = this;
      that.$refs.form.validate().then(result => {
        if (!result) return;
        if (
          that.active == 1 &&
          that.loginModel.password1 != that.loginModel.password2
        ) {
          that.pwdErrorText = "两次输入的密码不一致";
          return;
        }
        if (that.nameErrorText.length > 0) {
          return;
        }
        var loading = that.$loading();
        that.pwdErrorText = "";
        let url = that.active == 0 ? "/land/signin" : "/land/signup";
        that.$axios
          .post(url, {
            name: that.loginModel.username,
            password: that.loginModel.password1,
            remember: that.loginModel.remember
          })
          .then(response => {
            if (response.code == 0) {
              if (that.active == 0) {
                localStorage.setItem("token", response.data);
                that.$toast.success("登陆成功, 请稍后...");
                setTimeout(() => {
                  that.$router.push({ path: "/" });
                }, 1200);
              } else {
                that.$toast.success("注册成功, 请登录");
                that.active = 0;
              }
            }
            loading.close();
          })
          .catch(() => {
            loading.close();
          });
      });
    },
    chenkName() {
      var that = this;
      if (that.active == 0 || that.loginModel.username.length < 3) {
        that.nameErrorText = "";
        return;
      }
      var loading = that.$loading({});
      that.$axios
        .get("/signupcheck/" + that.loginModel.username)
        .then(response => {
          if (response.exist) {
            that.nameErrorText = "已存在的用户名";
          }
          loading.close();
        })
        .catch(() => {
          loading.close();
        });
    }
  },
  watch: {
    active(val) {
      this.$refs.form.clear();
      if (this.loginModel.username.length <= 2) {
        this.loginModel.username = "";
      }
      this.loginModel.password1 = "";
      this.loginModel.password2 = "";
      this.loginModel.remember = true;
    }
  }
};
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
