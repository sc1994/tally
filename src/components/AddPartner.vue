<template>
  <mu-dialog transition="slide-bottom" fullscreen :open.sync="open" :class="appbarStyle">
    <mu-appbar color="primary" title="选择小伙伴">
      <mu-button slot="left" icon @click="$emit('update:open', false)">
        <mu-icon value="close"></mu-icon>
      </mu-button>
      <mu-button slot="right" flat @click="sumbit">
        Done
      </mu-button>
    </mu-appbar>
    <div style="padding: 24px;">
      this is a fullscreen dialog
    </div>
  </mu-dialog>
</template>

<script>
import { mapState } from "vuex";

export default {
  props: {
    open: {
      type: Boolean,
      required: true
    }
  },
  methods: {
    sumbit() {
      alter("保存 todo");
      $emit("update:open", false);
    }
  },
  computed: {
    ...mapState(["currentUser", "appbarStyle"])
  }
};
</script>

<style>
</style>
