<template>
  <layoutmain>
    <mu-paper :z-depth="2" style="height:110px;padding:10px">
      <mu-list textline="three-line">
        <mu-list-item avatar button :ripple="false">
          <mu-list-item-action>
            <mu-avatar style="width:70px;height:70px">
              <img src="/static/images/head-default.png">
            </mu-avatar>
          </mu-list-item-action>
          <mu-list-item-content style="margin-left: 22px;margin-top: -17px;">
            <mu-list-item-title>{{currentUser.nick}}</mu-list-item-title>
            <mu-list-item-sub-title>
              账号 : {{currentUser.name}}
            </mu-list-item-sub-title>
            <mu-list-item-sub-title>
              备注 : {{currentUser.intro}}
            </mu-list-item-sub-title>
          </mu-list-item-content>
          <mu-list-item-action>
            <mu-badge :content="messageUnreadCount+''" circle color="secondary" v-if="messageUnreadCount>0">
              <mu-button icon color="blue" @click="$router.push({ path: 'message' })">
                <mu-icon value="notifications"></mu-icon>
              </mu-button>
            </mu-badge>
            <mu-button icon v-else @click="$router.push({ path: 'message' })">
              <mu-icon value="notifications"></mu-icon>
            </mu-button>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list>
    </mu-paper>
    <br/>
    <mu-list toggle-nested>
      <mu-sub-header>个人设置</mu-sub-header>
      <mu-divider></mu-divider>
      <mu-list-item button :ripple="false" nested :open="open === 'base'" @toggle-nested="open = arguments[0] ? 'base' : ''">
        <mu-list-item-title>基本信息</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'base'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested" @click="openUpload()">
          <mu-list-item-title>头像</mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="backup" color="Teal"></mu-icon>
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
            每月预算
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money">{{$numberFormat(currentUser.budget)}}</span>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list-item>

      <mu-list-item button :ripple="false" nested :open="open === 'partner'" @toggle-nested="open = arguments[0] ? 'partner' : ''">
        <mu-list-item-title>小伙伴</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'partner'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested" v-for="item in currentUser.partners" :key="item.id">
          <mu-list-item-action>
            <mu-avatar>
              <img :src="item.headImg">
            </mu-avatar>
          </mu-list-item-action>
          <mu-list-item-title>{{item.nick}}</mu-list-item-title>
          <mu-list-item-action>
            <mu-icon value="chat_bubble"></mu-icon>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested" @click="$router.push({ path: 'addpartner' })">
          <mu-list-item-title>添加更多</mu-list-item-title>
          <mu-list-item-action>
            <mu-button icon color="#f44336">
              <mu-icon value="zoom_in"></mu-icon>
            </mu-button>
          </mu-list-item-action>
        </mu-list-item>
      </mu-list-item>
    </mu-list>
    <mu-list toggle-nested>
      <mu-sub-header>钱钱钱</mu-sub-header>
      <mu-divider></mu-divider>
      <mu-list-item button :ripple="false" nested :open="open === 'money'" @toggle-nested="open = arguments[0] ? 'money' : ''">
        <mu-list-item-title>余额</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'money'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('backCard')">
          <mu-list-item-title>
            银行卡
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money">{{$numberFormat(currentUser.backCard)}}</span>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('aliPay')">
          <mu-list-item-title>
            支付宝
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money">{{$numberFormat(currentUser.aliPay)}}</span>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('wechatPay')">
          <mu-list-item-title>
            微信
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money">{{$numberFormat(currentUser.wechatPay)}}</span>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested" @click="openBaseInfoSet('fixDate')">
          <mu-list-item-title>
            定期 :
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money">{{$numberFormat(currentUser.fixDate)}}</span>
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
            花呗
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money loan">{{$numberFormat(currentUser.antCheck)}}</span>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested">
          <mu-list-item-title>
            信用卡
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money loan">{{$numberFormat(currentUser.creditCard)}}</span>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested">
          <mu-list-item-title>
            白条
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money loan">{{$numberFormat(currentUser.whiteBar)}}</span>
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
    <mu-list>
      <mu-sub-header>其他</mu-sub-header>
      <mu-divider></mu-divider>
      <mu-list-item button>
        <mu-list-item-title>快速消费类型</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon value="navigate_next"></mu-icon>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item button>
        <mu-list-item-title>账单类型</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon value="navigate_next"></mu-icon>
        </mu-list-item-action>
      </mu-list-item>
    </mu-list>
    <br/><br/>
    <mu-flex justify-content="center" align-items="center">
      <mu-button round color="success" @click="loginOut" full-width>退出登陆</mu-button>
    </mu-flex>
    <div style="height:80px;"></div>
    <setuserbaseinfo :user="currentUser" :type="baseInfo.type" :alert.sync="baseInfo.alert"></setuserbaseinfo>
    <vuecoreimageupload :data="uploadData" crop="local" :url="uploadUrl" @imageuploaded="uploaded" crop-ratio="1:1" extensions="png,jpeg,jpg" compress="60">
      </vue-core-image-upload>
    </vuecoreimageupload>
  </layoutmain>
</template>

<script>
import layoutmain from "@/layout/main";
import setuserbaseinfo from "@/components/setuserbaseinfo";
import vuecoreimageupload from "vue-core-image-upload";
import { mapState } from "vuex";

export default {
  components: {
    setuserbaseinfo,
    layoutmain,
    vuecoreimageupload
  },
  data() {
    return {
      open: "",
      baseInfo: {
        alert: false,
        type: ""
      },
      addpartnerOpen: false,
      messageUnreadCount: 0,
      openMessage: false,
      uploadUrl: this.$fileUrl,
      uploadData: {
        id: "",
        fileName: "head-image"
      }
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
        .then(() => {
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
    },
    openUpload() {
      document.getElementsByName("files")[0].click();
    },
    uploaded(response) {
      if (response.result) {
        // todo
        alert(response.path);
      }
    }
  },
  watch: {
    currentUser(val) {
      var that = this;
      that.$axios.get("/getmessageunreadcount/" + val.id).then(response => {
        that.messageUnreadCount = response.result;
      });
      this.uploadData.id = this.currentUser.id;
    }
  },
  mounted() {}
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