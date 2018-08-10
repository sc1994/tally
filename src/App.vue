<template>
  <div>
    <mu-appbar :class="'top-van ' + appbarStyle" :title="title" v-if="showBot">
      <mu-button icon flat slot="right">
        <i class="fa fa-github" style="font-size: 26px;margin-right: 15px;"></i>
      </mu-button>
    </mu-appbar>
    <mu-container style="position: absolute;top: 65px;">
      <router-view/>
    </mu-container>
    <mu-bottom-nav :value.sync="value" v-on:change="goto" class="bot-van" v-if="showBot">
      <mu-bottom-nav-item value="/" icon=":fa fa-circle-o-notch" class="bot-font-icon">
      </mu-bottom-nav-item>
      <mu-bottom-nav-item value="/tally" icon=":fa fa-pie-chart" class="bot-font-icon"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/me" icon=":fa fa-user-circle-o" class="bot-font-icon"></mu-bottom-nav-item>
    </mu-bottom-nav>
  </div>
</template>

<script>
import { mapState } from "vuex";
import theme from "muse-ui/lib/theme";

theme.use("light");

export default {
  name: "App",
  data() {
    return {
      value: "/",
      showBot: false,
      title: ""
    };
  },
  methods: {
    goto(val) {
      this.$router.push({ path: val });
    },
    init() {
      if (this.showBot) {
        this.$store.dispatch("initUser", { $router: this.$router });
      }
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
    }
  },
  computed: {
    ...mapState(["currentUser", "appbarStyle"])
  },
  watch: {
    $route(val) {
      console.log(val);
      this.value = val.path;
      this.showBot = val.path != "/sign";
      this.init();
    }
  },
  mounted() {
    this.value = this.$router.currentRoute.path;
    this.showBot = this.$router.currentRoute.path != "/sign";
    this.init();
    if (navigator.userAgent.indexOf("iPhone") > 0) {
      this.$store.commit("changeAppbarStyle", "top-van-backage-iphone");
    } else {
      this.$store.commit("changeAppbarStyle", "top-van-backage-oneplus");
    }
  }
};
</script>

<style>
.bot-van {
  position: fixed;
  bottom: 0px;
  width: 100%;
}

.top-van {
  position: fixed;
  top: 0px;
  width: 100%;
  height: 56px;
  margin-top: -1px;
}

.top-van-backage-iphone {
  background-color: #ffffff !important;
  color: #000000 !important;
}

.top-van-backage-oneplus {
  background-image: linear-gradient(0deg, #666 0, #000) !important;
  color: #ffffff !important;
}

.top-van > .mu-appbar-title {
  font-size: 18px;
  font-weight: 500;
}

.bot-font-icon {
  font-size: 28px;
  padding-top: 11px;
}
</style>
