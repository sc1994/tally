<template>
  <div>
    <mu-paper :z-depth="2" style="height:110px;padding:10px">
      <mu-list textline="two-line">
        <mu-list-item avatar button :ripple="false">
          <mu-list-item-action>
            <mu-avatar style="width:60px;height:60px">
              <img src="/static/images/head-default.png">
            </mu-avatar>
          </mu-list-item-action>
          <mu-list-item-content style="margin-left: 22px;">
            <mu-list-item-title>{{currentUser.nick}}</mu-list-item-title>
            <mu-list-item-sub-title>账号 : {{currentUser.name}}</mu-list-item-sub-title>
          </mu-list-item-content>
          <mu-list-item-action>
            <mu-badge :content="readNumber" circle color="secondary" v-if="readNumber>0">
              <mu-button icon color="blue">
                <mu-icon value="notifications"></mu-icon>
              </mu-button>
            </mu-badge>
            <mu-button icon v-else>
              <mu-icon value="notifications"></mu-icon>
            </mu-button>
          </mu-list-item-action>
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
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('nick')">
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
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('budget')">
          <mu-list-item-title>
            每月预算 :
            <span class="span-money">{{$numberFormat(currentUser.budget)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list-item>
      <mu-list-item button :ripple="false" nested :open="open === 'partner'" @toggle-nested="open = arguments[0] ? 'partner' : ''">
        <mu-list-item-title>小伙伴</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'partner'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested" @click="addpartnerOpen=true">
          <mu-list-item-title>添加更多</mu-list-item-title>
          <mu-list-item-action>
            <mu-button icon color="#f44336">
              <mu-icon value="group_add"></mu-icon>
            </mu-button>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list-item>
      <mu-list-item button :ripple="false" nested :open="open === 'money'" @toggle-nested="open = arguments[0] ? 'money' : ''">
        <mu-list-item-title>余额</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'money'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('backCard')">
          <mu-list-item-title>
            银行卡 :
            <span class="span-money">{{$numberFormat(currentUser.backCard)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('aliPay')">
          <mu-list-item-title>
            支付宝 :
            <span class="span-money">{{$numberFormat(currentUser.aliPay)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('wechatPay')">
          <mu-list-item-title>
            微信 :
            <span class="span-money">{{$numberFormat(currentUser.wechatPay)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('fixDate')">
          <mu-list-item-title>
            定期 :
            <span class="span-money">{{$numberFormat(currentUser.fixDate)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list-item>
      <mu-list-item button :ripple="false" nested :open="open === 'advance'" @toggle-nested="open = arguments[0] ? 'advance' : ''">
        <mu-list-item-title>预支</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'advance'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested">
          <mu-list-item-title>
            花呗 :
            <span class="span-money loan">{{$numberFormat(currentUser.antCheck)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested">
          <mu-list-item-title>
            信用卡 :
            <span class="span-money loan">{{$numberFormat(currentUser.creditCard)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested">
          <mu-list-item-title>
            白条 :
            <span class="span-money loan">{{$numberFormat(currentUser.whiteBar)}}</span>
            元
          </mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="edit" color="Teal"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list-item>
      <mu-list-item button :ripple="false" nested :open="open === 'type'" @toggle-nested="open = arguments[0] ? 'type' : ''">
        <mu-list-item-title>账单类型</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'type'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item :key="i" v-for="item,i in currentUser.consumes" button :ripple="true" slot="nested">
          <mu-list-item-title>{{item.content}}</mu-list-item-title>
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
    <setuserbaseinfo :user="currentUser" :type="baseInfo.type" :alert.sync="baseInfo.alert"></setuserbaseinfo>
    <addpartner :open.sync="addpartnerOpen"></addpartner>
  </div>
</template>

<script>
import setuserbaseinfo from "@/components/setuserbaseinfo";
import addpartner from "@/components/addpartner";
import { mapState } from "vuex";

export default {
  components: {
    setuserbaseinfo,
    addpartner
  },
  data() {
    return {
      open: "",
      baseInfo: {
        alert: false,
        type: ""
      },
      readNumber: 0,
      addpartnerOpen: false
    };
  },
  computed: {
    ...mapState(["currentUser"])
  },
  methods: {
    loginOut() {
      var that = this;
      that.$axios
        .get("/removetoken/" + localStorage.getItem("token"))
        .then(result => {
          localStorage.clear();
          alert("已退出登陆");
          that.$router.push({ path: "/sign" });
        });
    },
    openBaseInfoSet(type) {
      this.baseInfo = {
        alert: true,
        type: type
      };
    }
  }
};
</script>

<style>
.span-money {
  font-size: 16px;
  font-weight: 900;
  color: #ff5722;
  margin-left: 2px;
}

.loan {
  color: #4caf50;
  margin-left: 2px;
}

.span-content {
  color: #9e9e9e;
  font-size: 13px;
}
</style>