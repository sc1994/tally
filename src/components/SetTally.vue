<template>
  <dialoghead :title="'修改账单'" :open.sync="thatOpen" :buttonclick="submit" buttonicon="done">
    <mu-form :model="form" label-position="right" label-width="68">
      <mu-form-item prop="input" label="金额">
        <mu-text-field v-model="form.money" type="number"></mu-text-field>
      </mu-form-item>
      <mu-form-item prop="select" label="种类">
        <mu-auto-complete v-model="form.type" :data="consumes" :max-search-results="8" open-on-focus label-float></mu-auto-complete>
      </mu-form-item>
      <mu-form-item prop="date" label="模式">
        <mu-radio v-model="form.mode" value="收入" label="收入"></mu-radio>
        <mu-radio v-model="form.mode" value="支出" label="支出"></mu-radio>
        <mu-radio v-model="form.mode" value="预支" label="预支"></mu-radio>
      </mu-form-item>
      <mu-form-item prop="radio" label="方式">
        <mu-radio v-model="form.channel" :value="item.content" :label="item.content" v-for="item in currentUser.channels" :key="item.content"></mu-radio>
      </mu-form-item>
      <mu-form-item prop="checkbox" label="日期">
        <mu-date-input v-model="form.ctime" type="dateTime" actions></mu-date-input>
      </mu-form-item>
      <mu-form-item prop="switch" label="备注">
        <mu-text-field v-model="form.remark"></mu-text-field>
      </mu-form-item>
    </mu-form>
    <div style="text-align: right;margin-right: 10px;margin-top: 10px;">
      <mu-button color="red" @click="openAlert=true">
        删除当前账单
        <mu-icon right value="delete"></mu-icon>
      </mu-button>
    </div>
    <mu-dialog title="确认操作 ?" width="600" max-width="80%" :esc-press-close="false" :overlay-close="false" :open.sync="openAlert">
      是否确定删除当前订单,删除后不可恢复!
      <mu-button slot="actions" flat color="primary" @click="openAlert=false">取消</mu-button>
      <mu-button slot="actions" flat color="primary" @click="del()">删除</mu-button>
    </mu-dialog>
  </dialoghead>
</template>

<script>
import dialoghead from "@/layout/dialog";
import { mapState } from "vuex";

export default {
  props: {
    open: {
      type: Boolean
    },
    currentItem: {
      type: Object
    }
  },
  components: {
    dialoghead
  },
  data() {
    return {
      thatOpen: false,
      consumes: [],
      labelPosition: "top",
      form: {
        id: 0,
        money: 0,
        type: "",
        mode: "",
        channel: "",
        ctime: "",
        remark: ""
      },
      thatSetInterval: 0,
      openAlert: false
    };
  },
  methods: {
    submit() {
      var that = this;
      var loading = that.$loading({});
      that.form.money = parseFloat(that.form.money);
      that.form.token = localStorage.getItem("token");
      that.$axios.post("/updatetallybyid", that.form).then(response => {
        if (response.result) {
          that.$toast.success("修改完成");
          // 反过来赋值
          that.currentItem.money = that.form.money;
          that.currentItem.type = that.form.type;
          that.currentItem.mode = that.form.mode;
          that.currentItem.ctime = that.form.ctime;
          that.currentItem.remark = that.form.remark;
          // 关闭窗口
          setTimeout(() => {
            if (that.thatSetInterval != 0) {
              clearInterval(that.thatSetInterval);
            }
            that.form = {
              id: 0,
              money: 0,
              type: "",
              mode: "",
              channel: "",
              ctime: "",
              remark: ""
            };
            that.thatOpen = false;
          }, 800);
        } else {
          that.$toast.info("数据异常,请重试");
        }
        loading.close();
      });
    },
    del() {
      var that = this;
      var loading = that.$loading({});
      that.$axios
        .get(
          `/deletetallybyid/${that.form.id}/${localStorage.getItem("token")}`
        )
        .then(response => {
          if (response.result) {
            that.$toast.success("已删除");
            that.openAlert = false;
            that.currentItem.tid = "";
            setTimeout(() => {
              that.thatOpen = false;
            }, 300);
          } else {
            that.$toast.warning("发生异常,稍后重试~~~");
          }
          loading.close();
        })
        .catch(() => {
          loading.close();
        });
    }
  },
  watch: {
    currentUser(val) {
      if (!val) return;
      this.consumes = this.$linq(val.consumes)
        .select(x => x.content)
        .toArray();
    },
    open(val) {
      if (val) this.thatOpen = true;
    },
    thatOpen(val) {
      if (!val) {
        this.$emit("update:currentItem", {});
        this.$emit("update:open", false);
      }
    },
    currentItem(val) {
      if (!val.money) return;
      this.form = {
        id: val.tid,
        money: val.money,
        type: val.type,
        mode: val.mode,
        channel: val.channel,
        ctime: val.ctime,
        remark: val.remark
      };
    },
    form(val) {
      if (!val.money) return;
      var flag = val.mode == "收入" ? "收取" : "支付";
      var def = `${val.type}${val.mode}了${val.money}元，通过${
        val.channel
      }${flag}。`;
      if (def == val.remark) {
        this.thatSetInterval = setInterval(() => {
          var flag = val.mode == "收入" ? "收取" : "支付";
          val.remark = `${val.type}${val.mode}了${val.money}元，通过${
            val.channel
          }${flag}。`;
        }, 800);
      }
    }
  },
  computed: {
    ...mapState(["currentUser"])
  }
};
</script>

<style>
</style>

