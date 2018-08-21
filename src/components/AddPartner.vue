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
      <mu-expansion-panel :expand="panel === item.name" @change="toggle(item.name)" v-for="item in list" :key="item.id">
        <div slot="header">{{item.name}}</div>
        <mu-flex align-items="center">
          <mu-list textline="two-line">
            <mu-list-item avatar :ripple="false" button>
              <mu-list-item-action>
                <mu-avatar>
                  <img :src="item.headImg">
                </mu-avatar>
              </mu-list-item-action>
              <mu-list-item-content>
                <mu-list-item-title>{{item.nick}}</mu-list-item-title>
                <mu-list-item-sub-title>
                  {{item.intro ? item.intro : "这家伙很懒什么都没说"}}
                </mu-list-item-sub-title>
              </mu-list-item-content>
            </mu-list-item>
          </mu-list>
        </mu-flex>
        <mu-button slot="action" flat color="primary" @click="send(item.id)">
          发送请求
          <mu-icon left value="send"></mu-icon>
        </mu-button>
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
      panel: "",
      oldSearchValue: [],
      list: []
    };
  },
  methods: {
    toggle(panel) {
      this.panel = panel === this.panel ? "" : panel;
    },
    send(tid) {
      var that = this;
      var loading = this.$loading({});
      that.$axios
        .post("/sendmessage", {
          tid: tid,
          fid: that.currentUser.id,
          content: that.currentUser.nick + "向您发送了添加小伙伴的请求",
          needTouch: true
        })
        .then(response => {
          if (response.data.result) {
            that.$toast.success("已发送");
          } else {
            that.$toast.error("网络异常,请重试");
          }
          loading.close();
        })
        .catch(error => {
          that.$toast.error("网络异常,请重试");
          loading.close();
        });
    },
    search() {
      var that = this;
      that.$axios
        .get("/findusersbynick/" + that.searchValue)
        .then(response => {
          if (response.data.result != null) {
            that.list = response.data.result;
          } else {
            that.list = [];
            that.$toast.info("没有相关用户");
          }
          that.oldSearchValue = [];
        })
        .catch(error => {
          that.$toast.error("网络异常,请重试");
        });
    }
  },
  watch: {
    searchValue(val) {
      var that = this;
      that.oldSearchValue.push(val);
      setTimeout(oldSearchValue => {
        if (that.oldSearchValue.length > 0) {
          if (val == that.oldSearchValue[that.oldSearchValue.length - 1]) {
            if (val) {
              that.search();
            }
          }
        }
      }, 500);
    }
  },
  computed: {
    ...mapState(["currentUser", "appbarStyle"])
  }
};
</script>

<style>
</style>
