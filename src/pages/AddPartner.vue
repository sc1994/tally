<template>
  <minor :title="'添加小伙伴'">
    <div style="padding: 24px;">
      <mu-text-field v-model="searchValue" label="搜索对方的账号" :action-icon="searching ? 'autorenew' : 'search'" style="width: 100%;"></mu-text-field><br/>
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
        <mu-button slot="action" flat color="primary" @click="send(item.id, item.nick,item.headImg)">
          发送邀请
          <mu-icon left value="send"></mu-icon>
        </mu-button>
      </mu-expansion-panel>
    </div>
  </minor>
</template>

<script>
import { mapState } from "vuex";
import minor from "@/layout/minor";

export default {
  components: {
    minor
  },
  data() {
    return {
      searchValue: "",
      panel: "",
      oldSearchValue: [],
      list: [],
      searching: false
    };
  },
  methods: {
    toggle(panel) {
      this.panel = panel === this.panel ? "" : panel;
    },
    send(tid, tnick, timg) {
      this.$store.dispatch("sendMessage", {
        tid: tid,
        fid: this.currentUser.id,
        content: "向你发送了添加小伙伴的邀请",
        needTouch: true,
        type: 1
      });
    },
    search() {
      var that = this;
      that.searching = true;
      that.$axios
        .get("/findusersbyname/" + that.searchValue)
        .then(response => {
          if (response.result != null) {
            var ids = that
              .$linq(that.currentUser.partners)
              .select(x => x.id)
              .toArray();
            that.list = that
              .$linq(response.result)
              .where(x => ids.indexOf(x.id) < 0 && that.currentUser.id != x.id)
              .toArray();
            if (that.list.length < 1) {
              that.$toast.info("可能你们已经是小伙伴了");
            }
          } else {
            that.list = [];
            that.$toast.info("查无结果~~~");
          }
          that.oldSearchValue = [];
          that.searching = false;
        })
        .catch(() => {
          that.searching = false;
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
            } else {
              that.list = [];
            }
          }
        }
      }, 500);
    }
  },
  computed: {
    ...mapState(["currentUser"])
  }
};
</script>

<style>
</style>
