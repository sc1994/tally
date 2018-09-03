<template>
  <div>
    <mu-appbar :class="'top-van ' + appbarStyle" :title="title">
      <mu-button icon slot="left" v-if="lefticon" @click="leftevent" :style="letfstyle">
        <mu-icon :value="lefticon"></mu-icon>
      </mu-button>
      <mu-button icon flat slot="right">
        <i class="fa fa-github" style="font-size: 26px;margin-right: 15px;"></i>
      </mu-button>
    </mu-appbar>
    <mu-container style="position: absolute;top: 65px;">
      <slot></slot>
      <div style="height:70px;"></div>
    </mu-container>
    <mu-bottom-nav :value.sync="value" v-on:change="goto" class="bot-van">
      <mu-bottom-nav-item value="/" icon=":fa fa-circle-o-notch" class="bot-font-icon">
      </mu-bottom-nav-item>
      <mu-bottom-nav-item value="/tally" icon=":fa fa-pie-chart" class="bot-font-icon"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/me" :icon="iconThree" class="bot-font-icon" :style="colorThree">
      </mu-bottom-nav-item>
    </mu-bottom-nav>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  props: {
    lefticon: {
      type: String
    },
    leftevent: {
      type: Function
    },
    letfstyle: {
      type: String
    }
  },
  data() {
    return {
      title: "",
      value: "",
      iconThree: ":fa fa-user-circle-o",
      colorThree: ""
    };
  },
  methods: {
    goto(val) {
      this.$router.push({ path: val });
    },
    init() {
      this.$store.dispatch("initUnreadNumber");
      switch (this.$router.currentRoute.path) {
        case "/":
          this.title = "添加一笔";
          break;
        case "/tally":
          this.title = "账单";
          break;
        case "/me":
          this.title = "我";
          this.colorThree = "color:#2196f3";
          break;
      }
      this.value = this.$router.currentRoute.path;
      if (this.unreadNumber == 0) {
        this.iconThree = ":fa fa-user-circle-o";
      } else if (this.unreadNumber > 9) {
        this.iconThree = "filter_9_plus";
        if (this.$router.currentRoute.path != "/me") {
          this.colorThree = "color:#ff4081";
        } else {
          this.colorThree = "";
        }
      } else {
        this.iconThree = "filter_" + this.unreadNumber;
        if (this.$router.currentRoute.path != "/me") {
          this.colorThree = "color:#ff4081";
        } else {
          this.colorThree = "";
        }
      }
    }
  },
  computed: {
    ...mapState(["appbarStyle", "unreadNumber"])
  },
  watch: {
    $route(val) {
      this.init();
    },
    unreadNumber(val) {
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


