<template>
  <div>
    <mu-appbar :class="'top-van ' + appbarStyle" :title="title">
      <mu-button icon flat slot="right">
        <i class="fa fa-github" style="font-size: 26px;margin-right: 15px;"></i>
      </mu-button>
    </mu-appbar>
    <mu-container style="position: absolute;top: 65px;">
      <slot></slot>
    </mu-container>
    <mu-bottom-nav :value.sync="value" v-on:change="goto" class="bot-van">
      <mu-bottom-nav-item value="/" icon=":fa fa-circle-o-notch" class="bot-font-icon">
      </mu-bottom-nav-item>
      <mu-bottom-nav-item value="/tally" icon=":fa fa-pie-chart" class="bot-font-icon"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/me" icon=":fa fa-user-circle-o" class="bot-font-icon"></mu-bottom-nav-item>
    </mu-bottom-nav>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  data() {
    return {
      title: "",
      value: ""
    };
  },
  methods: {
    goto(val) {
      this.$router.push({ path: val });
    },
    init() {
      switch (this.$router.currentRoute.path) {
        case "/":
          this.title = "添加一笔";
          break;
        case "/tally":
          this.title = "账单";
          break;
        case "/me":
          this.title = "我";
          break;
      }
      this.value = this.$router.currentRoute.path;
    }
  },
  computed: {
    ...mapState(["appbarStyle"])
  },
  watch: {
    $route(val) {
      this.init();
    }
  },
  mounted() {
    this.init();
  }
};
</script>

<style>
.bot-font-icon {
  font-size: 28px;
  padding-top: 11px;
}

.bot-van {
  position: fixed !important;
  bottom: 0px !important;
  width: 100% !important;
}

.top-van > .mu-appbar-title {
  font-size: 18px;
  font-weight: 500;
}
</style>


