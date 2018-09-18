<template>
  <mu-drawer :open.sync="open" right :docked="true" style="top:55px;border-radius:0px;padding:15px 15px;" :z-depth="0" width="80%">
    <mu-list>
      <mu-list-item>
        <mu-list-item-action>
          <mu-icon value="monetization_on"></mu-icon>
        </mu-list-item-action>
        <mu-list-item-title>共计:{{total}} 元</mu-list-item-title>
      </mu-list-item>
    </mu-list>
    <mu-divider style="position: fixed;left: 0px;" />
    <mu-form :model="searchForm" label-position="top" style="margin-top: 20px;">
      <mu-form-item prop="partners" label="小伙伴">
        <mu-checkbox v-model="searchForm.partners" :value="currentUser.id" label="我"></mu-checkbox>
        <template v-for="item in currentUser.partners">
          <mu-checkbox v-model="searchForm.partners" :value="item.id" :label="item.nick"></mu-checkbox>
        </template>
      </mu-form-item>
      <mu-form-item prop="modes" label="方式">
        <mu-checkbox v-model="searchForm.modes" value="收入" label="收入"></mu-checkbox>
        <mu-checkbox v-model="searchForm.modes" value="支出" label="支出"></mu-checkbox>
        <mu-checkbox v-model="searchForm.modes" value="预支" label="预支"></mu-checkbox>
      </mu-form-item>
      <mu-form-item prop="consumes" label="消费渠道">
        <template v-for="item in currentUser.channels">
          <mu-checkbox v-model="searchForm.channels" :value="item.content" :label="item.content"></mu-checkbox>
        </template>
      </mu-form-item>
      <mu-form-item prop="ttime" label="日期">
        <mu-date-input v-model="searchForm.ttime" type="month"></mu-date-input>
      </mu-form-item>
    </mu-form>
    <div style="text-align: right;">
      <mu-button color="primary" @click="search()">
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
    },
    searchForm: {
      type: Object,
      default: false
    }
  },
  data() {
    return {
      total: "0"
    };
  },
  methods: {
    getTotal() {
      var that = this;
      this.$axios
        .post("/tally/total", {
          uids: this.searchForm.partners,
          btime: new Date(
            new Date(this.searchForm.ttime).getFullYear(),
            new Date(this.searchForm.ttime).getMonth(),
            1
          ).toISOString(),
          etime: new Date(
            new Date(this.searchForm.ttime).getFullYear(),
            new Date(this.searchForm.ttime).getMonth() + 1,
            1
          ).toISOString(),
          bMoney: 0,
          eMoney: 999999,
          types: [],
          modes: this.searchForm.modes,
          channels: this.searchForm.channels
        })
        .then(response => {
          if (response.data > 0) {
            that.total = parseFloat(response.data).toFixed(1);
          }
        });
    },
    search() {
      this.$emit("update:open", false);
    }
  },
  watch: {
    open(val) {
      if (val) {
        this.getTotal();
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
