<template>
  <mu-drawer :open.sync="thatOpen" right :docked="false" style="top:55px;border-radius:0px;padding:15px 15px;" :z-depth="0" width="80%">
    <mu-list>
      <mu-list-item>
        <mu-list-item-action>
          <mu-icon value="monetization_on"></mu-icon>
        </mu-list-item-action>
        <mu-list-item-title>共计: {{$numberFormat(total)}} 元</mu-list-item-title>
      </mu-list-item>
    </mu-list>
    <mu-divider/>
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
      <mu-form-item prop="ttime" label="日期">
        <mu-date-input v-model="searchForm.ttime" type="month"></mu-date-input>
      </mu-form-item>
      <mu-form-item prop="consumes" label="类型">
        <template v-for="item in copyConsumes">
          <mu-chip :color="item.selected?item.color:'#424242'" @click="selected(item)" style="margin-top: 5px;margin-right: 5px;">
            <mu-avatar :color="item.selected?item.color:'#424242'" :size="23">{{item.count}}</mu-avatar>{{item.content}}
          </mu-chip>
        </template>

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
      total: "0",
      copyConsumes: [],
      thatOpen: false,
      colors: [
        "#43a047",
        "#689f38",
        "#558b2f",
        "#827717",
        "#33691e",
        "#0097a7",
        "#00897b",
        "#009688",
        "#00695c",
        "#0288d1",
        "#0277bd",
        "#01579b",
        "#651fff",
        "#7e57c2",
        "#5e35b1",
        "#536dfe",
        "#304ffe",
        "#2962ff",
        "#448aff",
        "#1976d2",
        "#e040fb",
        "#7b1fa2",
        "#ab47bc",
        "#ec407a",
        "#ff4081",
        "#f44336",
        "#ff1744",
        "#ec407a",
        "#ef5350",
        "#ab47bc"
      ]
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
          modes: this.searchForm.modes
        })
        .then(response => {
          if (response.data > 0) {
            that.total = parseFloat(response.data).toFixed(1);
          }
        });
    },
    search() {
      this.$emit("update:open", false);
    },
    randomColor() {
      var i = Math.floor(Math.random() * this.colors.length);
      return this.colors[i];
    },
    selected(item) {
      item.selected = !item.selected;
      this.copyConsumes = this.$linq(this.copyConsumes)
        .orderByDescending(x => x.selected)
        .thenByDescending(x => x.count)
        .thenByDescending(x => new Date(x.utime))
        .toArray();
    }
  },
  watch: {
    open(val) {
      if (val) {
        this.getTotal();
        this.thatOpen = true;
      }
    },
    thatOpen(val) {
      if (!val) {
        this.$emit("update:open", false);
      }
    },
    currentUser(val) {
      if (!val) return;
      this.copyConsumes = JSON.parse(JSON.stringify(val.consumes));
      var that = this;
      this.copyConsumes.forEach(x => {
        x.color = that.randomColor();
        x.selected = true;
      });
    }
  },
  computed: {
    ...mapState(["currentUser"])
  }
};
</script>

<style>
.mu-overlay {
  opacity: 0;
}
</style>
