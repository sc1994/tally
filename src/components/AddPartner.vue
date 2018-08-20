<template>
  <mu-dialog transition="slide-bottom" fullscreen :open.sync="open">
    <mu-appbar color="primary" title="选择小伙伴" :class="appbarStyle">
      <mu-button slot="left" icon @click="$emit('update:open', false)">
        <mu-icon value="close"></mu-icon>
      </mu-button>
    </mu-appbar>
    <div style="padding: 24px;">
      <mu-text-field v-model="searchValue" label="搜索对方的账号" action-icon="restore" style="width: 100%;"></mu-text-field><br/>
    </div>
    <div style="padding: 0px 9px;margin-top: -21px;">
      <mu-expansion-panel :expand="panel === 'panel1'" @change="toggle('panel1')">
        <div slot="header">suncheng</div>
        <mu-flex align-items="center">
          <mu-list textline="two-line">
            <mu-list-item avatar :ripple="false" button>
              <mu-list-item-action>
                <mu-avatar>
                  <img src="/static/images/head-default.png">
                </mu-avatar>
              </mu-list-item-action>
              <mu-list-item-content>
                <mu-list-item-title>孙大诚</mu-list-item-title>
                <mu-list-item-sub-title>
                  这家伙很懒什么都没说
                </mu-list-item-sub-title>
              </mu-list-item-content>
            </mu-list-item>
          </mu-list>
        </mu-flex>
        <mu-button slot="action" flat color="primary">发送请求</mu-button>
      </mu-expansion-panel>
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
  data() {
    return {
      searchValue: "",
      panel: ""
    };
  },
  methods: {
    toggle(panel) {
      this.panel = panel === this.panel ? "" : panel;
    }
  },
  computed: {
    ...mapState(["currentUser", "appbarStyle"])
  }
};
</script>

<style>
</style>
