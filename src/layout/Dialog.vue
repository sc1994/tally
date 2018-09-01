<template>
  <mu-dialog fullscreen :open.sync="open">
    <mu-appbar :title="title" :class="appbarStyle">
      <mu-button slot="left" icon @click="$emit('update:open', false)">
        <mu-icon value="close"></mu-icon>
      </mu-button>
      <mu-button slot="right" flat @click="buttonclick" v-if="buttonshow">
        <mu-icon :value="buttonicon"></mu-icon>
      </mu-button>
    </mu-appbar>
    <br/>
    <slot></slot>
  </mu-dialog>
</template>

<script>
import { mapState } from "vuex";
export default {
  props: {
    title: {
      type: String
    },
    open: {
      type: Boolean
    },
    buttonshow: {
      type: Boolean
    },
    buttonclick: {
      type: Function
    },
    buttonicon: {
      type: String
    }
  },
  computed: {
    ...mapState(["appbarStyle"])
  }
};
</script>

