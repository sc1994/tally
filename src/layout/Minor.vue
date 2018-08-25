<template>
  <div>
    <mu-appbar :class="'top-van ' + appbarStyle">
      <mu-button icon slot="left" @click="$router.back(-1)">
        <mu-icon value="reply" left></mu-icon>
      </mu-button>
      {{title}}
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
    },
    title: {
      type: String
    }
  },
  computed: {
    ...mapState(["appbarStyle"])
  }
};
</script>

<style>
</style>

