<template>
  <mu-dialog :title="title" width="600" max-width="80%" :esc-press-close="false" :overlay-close="false" :open.sync="thenAlert">
    <mu-text-field v-model="value" :type="textType" :label="oldValue" :prefix="prefix" :error-text="textError" label-float></mu-text-field>
    <mu-button slot="actions" flat color="primary" @click="thenAlert=false">取消</mu-button>
    <mu-button slot="actions" flat color="primary" @click="submitBase">保存</mu-button>
  </mu-dialog>
</template>

<script>
export default {
  props: {
    alert: {
      type: Boolean,
      required: true
    },
    user: {
      type: Object,
      required: true
    },
    type: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      thenAlert: this.alert,
      title: "",
      value: "",
      oldValue: "",
      prefix: "",
      textType: "number",
      textError: ""
    };
  },
  methods: {
    submitBase() {
      var that = this;
      var loading = that.$loading();
      var request = {
        headImg: that.user.headImg,
        nick: that.user.nick,
        budget: parseFloat(that.user.budget),
        fixDate: parseFloat(that.user.fixDate),
        wechatPay: parseFloat(that.user.wechatPay),
        aliPay: parseFloat(that.user.aliPay),
        backCard: parseFloat(that.user.backCard),
        cash: parseFloat(that.user.cash)
      };
      switch (this.type) {
        case "nick":
          request.nick = that.value;
          break;
        case "budget":
          request.budget = parseFloat(that.value);
          break;
        case "backCard":
          request.backCard = parseFloat(that.value);
          break;
        case "aliPay":
          request.aliPay = parseFloat(that.value);
          break;
        case "wechatPay":
          request.wechatPay = parseFloat(that.value);
          break;
        case "fixDate":
          request.fixDate = parseFloat(that.value);
          break;
      }
      that.$axios
        .post("/user/set", request)
        .then(response => {
          if (response.code == 0) {
            // that.$toast.success("修改成功");
            that.$store.dispatch("initUser", { $router: this.$router }); // 刷新用户信息
            that.thenAlert = false;
          }
          loading.close();
        })
        .catch(() => {
          loading.close();
        });
    },
    switchType() {
      switch (this.type) {
        case "nick":
          this.title = "修改昵称";
          this.oldValue = this.user.nick + "";
          this.prefix = "昵称: ";
          this.textType = "string";
          break;
        case "budget":
          this.title = "修改预算";
          this.oldValue = this.user.budget + "";
          this.prefix = "预算: ";
          break;
        case "backCard":
          this.title = "修改银行卡余额";
          this.oldValue = this.user.backCard + "";
          this.prefix = "银行卡余额: ";
          break;
        case "aliPay":
          this.title = "修改支付宝余额";
          this.oldValue = this.user.aliPay + "";
          this.prefix = "支付宝余额: ";
          break;
        case "wechatPay":
          this.title = "修改微信余额";
          this.oldValue = this.user.wechatPay + "";
          this.prefix = "微信余额: ";
          break;
        case "fixDate":
          this.title = "修改定期余额";
          this.oldValue = this.user.fixDate + "";
          this.prefix = "定期余额: ";
          break;
      }
    }
  },
  watch: {
    thenAlert(val) {
      if (!val) {
        this.$emit("update:alert", val);
      }
    },
    alert(val) {
      if (val) {
        this.switchType();
        this.thenAlert = val;
        this.value = "";
      }
    }
  }
};
</script>

<style>
.mu-dialog-body {
  height: 66px;
}
</style>

