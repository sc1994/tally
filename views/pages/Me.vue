<template>
  <layoutmain>
    <mu-paper :z-depth="2" style="height:110px;padding:10px">
      <mu-list textline="three-line">
        <mu-list-item avatar button :ripple="false">
          <mu-list-item-action>
            <mu-avatar style="width:70px;height:70px">
              <img :src="currentUser.headImg">
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
            <mu-badge :content="unreadNumber+''" circle color="secondary" v-if="unreadNumber>0">
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
            <mu-icon value="cloud_upload" color="Teal"></mu-icon>
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
      <mu-list-item button :ripple="false" nested :open="open === 'advance'" @toggle-nested="open = arguments[0] ? 'advance' : ''">
        <mu-list-item-title>预支</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'advance'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested" v-for="item in advance">
          <mu-list-item-title>
            {{item._id}}
            <span class="span-content" v-if="item.rdate">
              {{item.rmonth}}{{item.rday}}
            </span>
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money loan">{{$numberFormat(item.money)}}</span>
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
        <mu-list-item-title>定时账单</mu-list-item-title>
        <mu-list-item-action>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_up" v-if="open === 'loan'"></mu-icon>
          <mu-icon class="toggle-icon" size="24" value="keyboard_arrow_down" v-else></mu-icon>
        </mu-list-item-action>
        <mu-list-item button :ripple="true" slot="nested">
          <mu-list-item-title>
            房贷
            <span class="span-content">
              每月12日
            </span>
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money loan">{{$numberFormat(2987)}}</span>
          </mu-list-item-action>
        </mu-list-item>
        <mu-list-item button :ripple="true" slot="nested">
          <mu-list-item-title>
            信用卡分期
            <span class="span-content">
              每月20日
            </span>
          </mu-list-item-title>
          <mu-list-item-action>
            <span class="span-money loan">{{$numberFormat(2987)}}</span>
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
    <br/>
    <vuecoreimageupload :data="uploadData" crop="local" :url="uploadUrl" @imageuploaded="uploaded" crop-ratio="1:1" extensions="png,jpeg,jpg" compress="60" :text="''">
    </vuecoreimageupload>
    <mu-flex justify-content="center" align-items="center">
      <mu-button round color="success" @click="loginOut" full-width>退出登陆</mu-button>
    </mu-flex>
    <setuserbaseinfo :user="currentUser" :type="baseInfo.type" :alert.sync="baseInfo.alert"></setuserbaseinfo>
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
      openMessage: false,
      uploadUrl: this.$fileUrl,
      uploadData: {
        id: "",
        fileName: "head-image-" + new Date().getMilliseconds()
      },
      advance: []
    };
  },
  computed: {
    ...mapState(["currentUser", "unreadNumber"])
  },
  methods: {
    loginOut() {
      var that = this;
      that.$axios
        .get("/land/signout/" + localStorage.getItem("token"))
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
        this.currentUser.headImg = response.path;
        var that = this;
        that.$axios
          .post("/user/set", {
            headImg: this.currentUser.headImg,
            nick: this.currentUser.nick,
            budget: parseFloat(this.currentUser.budget),
            fixDate: parseFloat(this.currentUser.fixDate),
            wechatPay: parseFloat(this.currentUser.wechatPay),
            aliPay: parseFloat(this.currentUser.aliPay),
            backCard: parseFloat(this.currentUser.backCard),
            cash: parseFloat(this.currentUser.cash),
            utime: this.currentUser.utime
          })
          .then(response => {
            if (response.code == 0) {
              that.$toast.success("设置成功");
            }
          });
      }
    },
    getAdvance() {
      var that = this;
      var date = new Date();
      this.$axios
        .post("/tally/getadvance", {
          btime: new Date(date.getFullYear(), date.getMonth(), 1).toISOString(),
          etime: new Date(
            date.getFullYear(),
            date.getMonth() + 1,
            1
          ).toISOString()
        })
        .then(response => {
          if (response.code == 0 && response.data) {
            that.advance = response.data;
          }
        });
    }
  },
  watch: {
    currentUser(val) {
      this.uploadData.id = this.currentUser.id;
    }
  },
  mounted() {
    this.getAdvance();
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