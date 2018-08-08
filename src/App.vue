<template>
  <div>
    <mu-appbar class="top-van" title="记录"></mu-appbar>
    <mu-container style="position: absolute;top: 55px;">
      <router-view/>
    </mu-container>
    <mu-bottom-nav :value.sync="value" v-on:change="goto" class="bot-van" v-if="showBot">
      <mu-bottom-nav-item value="/" title="添加一笔" icon="add_shopping_cart"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/bill" title="查账" icon="show_chart"></mu-bottom-nav-item>
      <mu-bottom-nav-item value="/me" title="我的" icon="account_box"></mu-bottom-nav-item>
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
      showBot: false
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
  background-color: #424242 !important;
  color: #fafafa !important;
  height: 47px !important;
}

.top-van > .mu-appbar-title {
  font-size: 15px !important;
  font-weight: 500 !important;
}
</style>
