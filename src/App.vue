<template>
  <div>
    <mu-appbar class="top-van" :title="title" v-if="showBot" :style="iphoneStyle"></mu-appbar>
    <mu-container style="position: absolute;top: 65px;">
      <router-view/>
    </mu-container>
    <mu-bottom-nav :value.sync="value" v-on:change="goto" class="bot-van" v-if="showBot">
      <mu-bottom-nav-item value="/" icon="add_shopping_cart"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/tally" icon="show_chart"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/me" icon="account_box"></mu-bottom-nav-item>
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
      title: "",
      iphoneStyle: ""
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
    ...mapState(["currentUser"])
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
    if (navigator.userAgent.indexOf("iPhone")) {
      this.iphoneStyle =
        "background-color:#ffffff !important;color:#000000 !important;";
    }
  }
};
</script>

<style>
.bot-van {
  position: fixed !important;
  bottom: 0px !important;
  width: 100% !important;
}

.top-van {
  position: fixed !important;
  top: 0px !important;
  width: 100% !important;
  background-color: #000000 !important;
  color: #ffffff !important;
  height: 56px !important;
  margin-top: -1px;
}

.top-van > .mu-appbar-title {
  font-size: 18px !important;
  font-weight: 500;
}

.mu-appbar {
  background-color: #000000 !important;
}
</style>
