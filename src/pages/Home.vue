<template>
  <div>
    <main-layout :alert.sync="alert" :loading.sync="loading" :user.sync="user" :initUser="initUser">
      <mu-paper :z-depth="0" class="paper-title">
        <Residue :height="165" :residue="90" style="margin-top: -10px;float:left;"></Residue>
        <div style="margin-top: 10px;">
          <span class="span-title">当月剩余：
            <span class="span-money">50,203</span> 元</span>
          <br /><br />
          <span class="span-title">当月消费：
            <span class="span-money" style="color:#4caf50">356</span> 元</span>
          <br /><br />
          <span class="span-title">当月预支：
            <span class="span-money" style="color:#4caf50">478</span> 元</span>
        </div>
      </mu-paper>
      <mu-paper :z-depth="1" style="height:100px;padding:10px">
        <mu-text-field @blur="showMode" v-model="tallyForm.money" style="height:80px;width:50%" label="金额" prefix="￥" type="number" label-float></mu-text-field>
        <mu-auto-complete @change="showMode" v-model="tallyForm.consume" :data="manyType.consumes" style="height:80px;width:45%" label="种类" :max-search-results="5" open-on-focus label-float></mu-auto-complete>
      </mu-paper>
      <br/>
      <mu-list>
        <mu-sub-header>
          最近消费
          <mu-button flat color="primary" class="button-right">查看全部</mu-button>
        </mu-sub-header>
        <mu-divider></mu-divider>
        <mu-list-item :key="index" v-for="item,index in tallyList" avatar :ripple="false">
          <mu-list-item-action>
            <mu-avatar>
              <img src="/src/static/head-default.png">
            </mu-avatar>
          </mu-list-item-action>
          <mu-list-item-content>
            <mu-list-item-title>{{item.remark}}</mu-list-item-title>
            <mu-list-item-sub-title>{{formatDate(item.ctime)}}</mu-list-item-sub-title>
          </mu-list-item-content>
        </mu-list-item>
      </mu-list>
      <mu-list>
        <mu-sub-header>
          快速记录
          <mu-button flat color="primary" class="button-right">添加更多</mu-button>
        </mu-sub-header>
        <mu-divider></mu-divider>
        <mu-list-item avatar button :ripple="true">
          <mu-list-item-action style="width:280px">
            15 元
          </mu-list-item-action>
          <mu-list-item-title>午餐</mu-list-item-title>
          <mu-list-item-action>
            <mu-button icon color="#4caf50">
              <mu-icon right value="flight_takeoff"></mu-icon>
            </mu-button>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item avatar button :ripple="true">
          <mu-list-item-action style="width:280px">
            8 元
          </mu-list-item-action>
          <mu-list-item-title>公交</mu-list-item-title>
          <mu-list-item-action>
            <mu-button icon color="#4caf50">
              <mu-icon right value="flight_takeoff"></mu-icon>
            </mu-button>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item avatar button :ripple="true">
          <mu-list-item-action style="width:280px">
            8 元
          </mu-list-item-action>
          <mu-list-item-title>公交</mu-list-item-title>
          <mu-list-item-action>
            <mu-button icon color="#4caf50">
              <mu-icon right value="flight_takeoff"></mu-icon>
            </mu-button>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item avatar button :ripple="true">
          <mu-list-item-action style="width:280px">
            8 元
          </mu-list-item-action>
          <mu-list-item-title>公交</mu-list-item-title>
          <mu-list-item-action>
            <mu-button icon color="#4caf50">
              <mu-icon right value="flight_takeoff"></mu-icon>
            </mu-button>
          </mu-list-item-action>
        </mu-list-item>
        <!-- 留个空的撑开样式 -->
        <mu-list-item avatar button :ripple="false">
          <mu-list-item-action style="width:280px">
          </mu-list-item-action>
          <mu-list-item-title></mu-list-item-title>
          <mu-list-item-action>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list>
    </main-layout>

    <mu-dialog fullscreen :open.sync="openScroll">
      <mu-appbar color="primary" :title="tallyForm.consume+'：'+tallyForm.money+' 元'">
        <mu-button slot="left" icon @click="openScroll = false">
          <mu-icon value="close"></mu-icon>
        </mu-button>
        <mu-button slot="right" flat @click="submit" v-if="step==2">
          完成
        </mu-button>
      </mu-appbar>
      <br/>
      <mu-stepper :active-step="step" orientation="vertical">
        <mu-step :style="'height:'+stepHeight[0]+'px;'">
          <mu-step-label>
            收入/支出/预支
            <b>{{tallyForm.mode.length>0 ? " ：":""}}{{tallyForm.mode}}</b>
          </mu-step-label>
          <mu-step-content>
            <mu-list>
              <mu-list-item v-if="item.content.length>0 && !item.hide" :key="index" v-for="item,index in manyType.modes">
                <mu-list-item-content>
                  <mu-radio :label="item.content" :value="item.content" v-model="tallyForm.mode" @change="step++"></mu-radio>
                </mu-list-item-content>
              </mu-list-item>
            </mu-list>
          </mu-step-content>
        </mu-step>
        <mu-step :style="'height:'+stepHeight[1]+'px;'">
          <mu-step-label>
            方式
            <b>{{tallyForm.channel.length>0 ? " ：":""}}{{tallyForm.channel}}</b>
          </mu-step-label>
          <mu-step-content>
            <mu-list>
              <mu-list-item v-if="item.content.length>0 && !item.hide" :key="index" v-for="item,index in manyType.channels">
                <mu-list-item-content>
                  <mu-radio :label="item.content" :value="item.content" v-model="tallyForm.channel" @change="step++"></mu-radio>
                </mu-list-item-content>
              </mu-list-item>
            </mu-list>
            <mu-button flat @click="step--" color="primary">上一步</mu-button>
          </mu-step-content>
        </mu-step>
        <mu-step :style="'height:'+stepHeight[2]+'px;'">
          <mu-step-label>
            其他
          </mu-step-label>
          <mu-step-content>
            <br/>
            <mu-form :model="tallyForm" label-position="top" label-width="70">
              <mu-form-item prop="date" label="日期">
                <mu-date-input v-model="tallyForm.date" type="dateTime" actions style="height: 30px;"></mu-date-input>
              </mu-form-item>
              <mu-form-item prop="input" label="备注">
                <mu-text-field v-model="tallyForm.remark" style="height: 30px;"></mu-text-field>
              </mu-form-item>
              <mu-form-item>
                <mu-button flat @click="step--" color="primary">上一步</mu-button>
              </mu-form-item>
            </mu-form>
          </mu-step-content>
        </mu-step>
      </mu-stepper>
    </mu-dialog>
  </div>
</template>

<script>
import MainLayout from '../layouts/Main.vue'
import Residue from '../components/Residue.vue'

export default {
  components: {
    MainLayout,
    Residue
  },
  data() {
    return {
      tallyForm: {
        money: '',
        consume: '',
        mode: '',
        channel: '',
        date: '',
        remark: ''
      },
      manyType: {
        channels: [],
        modes: [],
        consumes: []
      },
      step: -1,
      openScroll: false,
      alert: {
        msg: '',
        type: 'success',
        show: false
      },
      loading: false,
      initUser: 1,
      user: {},
      stepHeight: [170, 50, 50],
      tallyList: []
    }
  },
  methods: {
    submit() {
      var that = this
      that.loading = true
      axios
        .post('/inserttally', {
          token: localStorage.getItem('token'),
          money: that.tallyForm.money,
          type: that.tallyForm.consume,
          mode: that.tallyForm.mode,
          channel: that.tallyForm.channel,
          remark: that.tallyForm.remark,
          ctime: that.tallyForm.date
        })
        .then(result => {
          if (result.data.result) {
            that.initUser = that.initUser + 1
            that.initLate()
            that.alert = {
              msg: '记录完成',
              type: 'success'
            }
            that.alert.show = true
            setTimeout(() => {
              that.openScroll = false
              that.tallyForm.money = ''
              that.tallyForm.consume = ''
            }, 500)
          } else {
            that.alert = {
              msg: '发送异常',
              type: 'error'
            }
            that.alert.show = true
          }
          that.loading = false
        })
        .catch(err => {
          that.alert = {
            msg: '系统异常请重试',
            type: 'error',
            show: true
          }
          that.loading = false
        })
    },
    showMode() {
      var that = this
      if (
        that.tallyForm.money.length <= 0 ||
        that.tallyForm.consume.length <= 0
      )
        return
      setTimeout(() => {
        that.openScroll = true
      }, 200)
      that.step = 0
      var isExist = false
      that.manyType.consumes.forEach(x => {
        if (x == that.tallyForm.consume) {
          isExist = true
        }
      })
      if (!isExist) {
        that.user.consumes.push({
          content: that.tallyForm.consume,
          count: 1,
          default: ['收入', '支出', '预支']
        })
      }
    },
    initLate() {
      var that = this
      axios
        .post('/gettallybyuser', {
          token: localStorage.getItem('token'),
          pageIndex: 1,
          pageSize: 3
        })
        .then(result => {
          if (result.data.result) {
            that.tallyList = result.data.body
          } else {
            that.alert.msg = '系统异常请重试'
            that.alert.type = 'error'
            that.show = true
          }
        })
    },
    formatDate(date) {
      return format(date, 'yyyy-MM-dd hh:mm')
    }
  },
  watch: {
    user(val) {
      this.manyType.channels = [
        {
          content: '',
          default: []
        },
        ...val.channels
      ]
      this.manyType.modes = [
        {
          content: '',
          default: []
        },
        ...val.modes
      ]
      this.manyType.consumes = []
      val.consumes.forEach(x => {
        this.manyType.consumes.push(x.content)
      })
    },
    step(val) {
      var that = this
      var min = 50
      var height1 = 0
      var height2 = 45 // 上一步按钮预留位置
      var height3 = 280
      var oneHeight = 60
      var onlyMode = ''
      if (val == 0) {
        that.user.consumes.forEach(c => {
          if (c.content == that.tallyForm.consume) {
            that.manyType.modes.forEach(m => {
              if (c.default.indexOf(m.content) > -1) {
                m.hide = false
                height1 += oneHeight
                onlyMode = m.content
              } else {
                m.hide = true
              }
            })
          }
        })
        if (height1 == oneHeight) {
          // 直接到第二步
          that.step = 1
          that.tallyForm.mode = onlyMode
        }
        that.stepHeight[0] = height1
        that.stepHeight[1] = min
        that.stepHeight[2] = min
      } else if (val == 1) {
        that.manyType.channels.forEach(c => {
          if (c.default.indexOf(that.tallyForm.mode) > -1) {
            c.hide = false
            height2 += oneHeight
          } else {
            c.hide = true
          }
        })
        that.stepHeight[0] = min
        that.stepHeight[1] = height2
        that.stepHeight[2] = min
      } else if (val == 2) {
        that.stepHeight[0] = min
        that.stepHeight[1] = min
        that.stepHeight[2] = height3
        that.tallyForm.date = new Date()
        var flag = that.tallyForm.mode == '收入' ? '收取' : '支付'
        that.tallyForm.remark = `${that.tallyForm.consume}${
          that.tallyForm.mode
        }了${that.tallyForm.money}元，通过${that.tallyForm.channel}${flag}。`
      }
    },
    openScroll(val) {
      if (!val) {
        this.tallyForm.mode = ''
        this.tallyForm.channel = ''
        this.step = -1
      }
    }
  },
  mounted() {
    this.showUser = true
    this.initLate()
  }
}
</script>

<style scoped>
.button-right {
  margin-left: 55%;
}
.span-title {
  color: rgb(20, 20, 20);
  margin-right: 0px;
}
.span-money {
  font-size: 20px;
  font-weight: 600;
  color: #ff5722;
  margin-left: 0px;
}
.paper-title {
  height: 180px;
  padding: 10px;
  margin-bottom: 10px;
  background-image: linear-gradient(
    to top,
    #a7a7a7 0%,
    lightgrey 1%,
    #e0e0e0 26%,
    #efefef 48%,
    #d9d9d9 75%,
    #a7a7a7 100%
  );
}
</style>
