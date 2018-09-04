<template>
  <div>
    <mu-appbar :class="'top-van ' + appbarStyle">
      <mu-button icon slot="left" @click="$router.back(-1)">
        <mu-icon value="reply" left></mu-icon>
      </mu-button>
      <mu-button icon flat slot="right" @click="rightevent">
        <mu-icon :value="righticon" left v-if="righticon"></mu-icon>
      </mu-button>
      {{title}}
    </mu-appbar>
    <mu-container style="position: absolute;top: 65px;">
      <slot></slot>
    </mu-container>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  props: {
    righticon: {
      type: String
    },
    rightevent: {
      type: Function,
      default: () => {
        return;
      }
    },
    title: {
      type: String
    }
  },
  computed: {
    ...mapState(["appbarStyle"])
  }
  // methods: {
  //   clickRightIcon() {
  //     debugger
  //     this.();
  //   }
  // }
};
</script>

<style>
</style>

