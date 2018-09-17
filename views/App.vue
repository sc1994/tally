<template>
  <div>
    <router-view/>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "App",
  watch: {
    $route(val) {
      if (val.path == "/sign") return;
      this.$store.dispatch("initUser", { $router: this.$router });
    }
  },
  mounted() {
    if (navigator.userAgent.indexOf("iPhone") > 0) {
      this.$store.commit("changeAppbarStyle", "top-van-backage-iphone");
    } else {
      this.$store.commit("changeAppbarStyle", "top-van-backage-oneplus");
    }
    if (window.document.URL.indexOf("/sign") < 0) {
      this.$store.dispatch("initUser", { $router: this.$router });
    }
  }
};
</script>

<style>
.top-van-backage-iphone {
  background-color: #ffffff !important;
  color: #000000 !important;
}

.top-van-backage-oneplus {
  background-color: #2b2c30 !important;
  color: #ffffff !important;
}

.top-van {
  position: fixed !important;
  top: 0px !important;
  width: 100% !important;
  height: 56px !important;
  margin-top: -1px !important;
}

.mu-divider {
  position: fixed;
  left: 0px;
}
</style>
