<template>
  <mu-drawer :open.sync="open" right :docked="true" style="top:55px;border-radius:0px;padding:15px 15px;" :z-depth="0" width="80%">
    <mu-form :model="form" label-position="top">
      <mu-form-item prop="partners" label="小伙伴">
        <mu-checkbox v-model="form.partners" :value="currentUser.id" label="我"></mu-checkbox>
        <template v-for="item in currentUser.partners">
          <mu-checkbox v-model="form.partners" :value="item.id" :label="item.nick"></mu-checkbox>
        </template>
      </mu-form-item>
      <mu-form-item prop="modes" label="方式">
        <mu-checkbox v-model="form.modes" value="收入" label="收入"></mu-checkbox>
        <mu-checkbox v-model="form.modes" value="支出" label="支出"></mu-checkbox>
        <mu-checkbox v-model="form.modes" value="预支" label="预支"></mu-checkbox>
      </mu-form-item>
      <mu-form-item prop="consumes" label="消费渠道">
        <template v-for="item in currentUser.channels">
          <mu-checkbox v-model="form.channels" :value="item.content" :label="item.content"></mu-checkbox>
        </template>
      </mu-form-item>
      <mu-form-item prop="ctime" label="日期">
        <mu-date-input v-model="form.ctime" type="month"></mu-date-input>
      </mu-form-item>
    </mu-form>
    <div style="text-align: right;">
      <mu-button color="primary">
        查询
        <mu-icon right value="search"></mu-icon>
      </mu-button>
    </div>
  </mu-drawer>
</template>

<script>
import { mapState } from "vuex";

export default {
  props: {
    open: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      form: {
        partners: [],
        channels: [],
        modes: [],
        ctime: new Date()
      }
    };
  },
  computed: {
    ...mapState(["currentUser"])
  },
  watch: {
    currentUser(val) {
      var that = this;
      that.form.partners = [val.id];
      val.partners.forEach(x => {
        that.form.partners.push(x.id);
      });
      that.currentUser.channels.forEach(x => {
        that.form.channels.push(x.content);
      });
    }
  },
  mounted() {
    this.form.modes = ["收入", "支出", "预支"];
  }
};
</script>

<style>
</style>
