<template>
  <main-layout :alert.sync="alert" :loading.sync="loading" :user.sync="user" :initUser="initUser">
    <mu-container class="button-wrapper">
      <mu-paper :z-depth="2" style="height:110px;padding:10px">
        <mu-list textline="two-line">
          <mu-list-item avatar button :ripple="false">
            <mu-list-item-action>
              <mu-avatar style="width:60px;height:60px">
                <img src="/src/static/head-default.png">
              </mu-avatar>
            </mu-list-item-action>
            <mu-list-item-content style="margin-left: 22px;">
              <mu-list-item-title>这里是昵称</mu-list-item-title>
              <mu-list-item-sub-title>账号 : suncheng</mu-list-item-sub-title>
            </mu-list-item-content>
          </mu-list-item>
        </mu-list>
      </mu-paper>
      <br/>
      <mu-list toggle-nested>
        <mu-sub-header>个人设置</mu-sub-header>
        <mu-list-item button :ripple="false" nested :open="open === 'base'" @toggle-nested="open = arguments[0] ? 'base' : ''">
          <mu-list-item-title>基本信息</mu-list-item-title>
          <mu-list-item-action>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'base'"></mu-icon>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
          </mu-list-item-action>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>头像</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>昵称</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>密码</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>
              每月预算 :
              <span class="span-money">2,876</span>
              元
            </mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
        </mu-list-item>
        <mu-list-item button :ripple="false" nested :open="open === 'money'" @toggle-nested="open = arguments[0] ? 'money' : ''">
          <mu-list-item-title>余额</mu-list-item-title>
          <mu-list-item-action>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'money'"></mu-icon>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
          </mu-list-item-action>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>
              银行卡 :
              <span class="span-money">2,876</span>
              元
            </mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>
              支付宝 :
              <span class="span-money">2,876</span>
              元
            </mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>
              微信 :
              <span class="span-money">2,876</span>
              元
            </mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>
              定期 :
              <span class="span-money">2,876</span>
              元
            </mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
        </mu-list-item>
        <mu-list-item button :ripple="false" nested :open="open === 'type'" @toggle-nested="open = arguments[0] ? 'type' : ''">
          <mu-list-item-title>账单种类维护</mu-list-item-title>
          <mu-list-item-action>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'type'"></mu-icon>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
          </mu-list-item-action>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>吃饭</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>工资</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>房租</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>添加更多</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="add" color="Red"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
        </mu-list-item>
        <mu-list-item button :ripple="false" nested :open="open === 'loan'" @toggle-nested="open = arguments[0] ? 'loan' : ''">
          <mu-list-item-title>贷款/分期</mu-list-item-title>
          <mu-list-item-action>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'loan'"></mu-icon>
            <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
          </mu-list-item-action>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>
              房贷:
              <span class="span-content">
                每月
                <span class="span-money loan">2,876</span>
                元->2019年5月结束
              </span>
            </mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>
              信用卡分期 :
              <span class="span-content">
                每月
                <span class="span-money loan">2,876</span>
                元->2019年5月结束
              </span>
            </mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="edit" color="Teal"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
          <mu-list-item button :ripple="true" slot="nested">
            <mu-list-item-title>添加更多</mu-list-item-title>
            <mu-list-item-action>
              <mu-icon value="add" color="Red"></mu-icon>
            </mu-list-item-action>
          </mu-list-item>
        </mu-list-item>
      </mu-list>
      <br/><br/>
      <mu-flex justify-content="center" align-items="center">
        <mu-button round color="success" @click="loginOut" full-width>退出登陆</mu-button>
      </mu-flex>
    </mu-container>
  </main-layout>
</template>

<script>
import MainLayout from '../layouts/Main.vue'

export default {
  components: {
    MainLayout
  },
  data() {
    return {
      open: '',
      alert: {
        msg: '',
        type: 'success',
        show: false
      },
      loading: false,
      initUser: 1,
      user: {}
    }
  },
  methods: {
    loginOut() {
      var that = this
      axios
        .get('/removetoken/' + localStorage.getItem('token'))
        .then(result => {
          localStorage.clear()
          alert('已退出登陆')
          that.$root.currentRoute = '/login'
          window.history.pushState(null, routes['/login'], '/login')
        })
    }
  }
}
</script>

<style>
.span-money {
  font-size: 16px;
  font-weight: 900;
  color: #ff5722;
  margin-left: 10px;
}

.loan {
  color: #4caf50;
  margin-left: 0px;
}

.span-content {
  color: #9e9e9e;
  font-size: 13px;
}
</style>
