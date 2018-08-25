<template>
  <div>
    <mu-appbar :class="'top-van ' + appbarStyle">
      <mu-button icon slot="left">
        <mu-icon value="reply"></mu-icon>
      </mu-button>
      <mu-button flat slot="right" v-if="righticon" @click="rightevent">
        <mu-icon value="menu"></mu-icon>
      </mu-button>
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
      type: Function
    }
  },
  computed: {
    ...mapState(["appbarStyle"])
  }
};
</script>

<style>
</style>

