<template>
  <mu-dialog transition="slide-bottom" fullscreen :open.sync="open">
    <mu-appbar color="primary" title="消息" :class="appbarStyle">
      <mu-button slot="left" icon @click="$emit('update:open', false)">
        <mu-icon value="close"></mu-icon>
      </mu-button>
    </mu-appbar>
    <mu-list textline="two-line">
      <mu-sub-header v-if="isToday()">今天</mu-sub-header>
      <mu-list-item avatar :ripple="false" button v-for="item in todays" :key="item.id">
        <mu-list-item-action>
          <mu-avatar>
            <img :src="item.fimg">
          </mu-avatar>
        </mu-list-item-action>
        <mu-list-item-content>
          <mu-list-item-title>{{item.fnick}}</mu-list-item-title>
          <mu-list-item-sub-title>
            <span style="color: rgba(0, 0, 0, .87)">
              {{item.content}}
            </span>
          </mu-list-item-sub-title>
        </mu-list-item-content>
        <mu-list-item-action>
          <mu-list-item-after-text>{{$format(item.ctime,"hh:mm")}}</mu-list-item-after-text>
          <mu-button icon color="primary" @click="openAlert=true;currentItem=item">
            <mu-icon value="sentiment_satisfied"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-sub-header v-if="isYesterday()">昨天</mu-sub-header>
      <mu-sub-header>更早</mu-sub-header>
    </mu-list>
    <mu-dialog title="是否同意 ?" width="75%" max-width="75%" :esc-press-close="false" :overlay-close="false" :open.sync="openAlert">
      {{currentItem.content}}
      <mu-button slot="actions" flat color="primary" @click="openAlert=false">算了</mu-button>
      <mu-button slot="actions" flat color="primary" @click="submit">没问题</mu-button>
    </mu-dialog>
  </mu-dialog>
</template>

<script>
import { mapState } from "vuex";

export default {
  props: {
    open: {
      type: Boolean
    },
    messages: {
      type: Array
    }
  },
  data() {
    return {
      openAlert: false,
      messageGroup: [],
      todays: [],
      yesterdays: [],
      currentItem: {}
    };
  },
  methods: {
    submit() {
      alert("好的");
      this.openAlert = false;
    },
    isToday() {
      var that = this;
      var temp = this.$linq(this.messageGroup)
        .where(
          x => x.key == that.$format(new Date(), "yyyy-MM-dd").substring(0, 10)
        )
        .toArray();
      if (temp.length > 0) {
        this.todays = temp[0].value;
      }
      return temp.length > 0;
    },
    isYesterday() {
      var that = this;
      var now = new Date();
      now.setDate(now.getDate() + 1);
      var temp = this.$linq(this.messageGroup)
        .where(x => x.key == that.$format(now, "yyyy-MM-dd").substring(0, 10))
        .toArray();
      if (temp.length > 0) {
        this.yesterdays = temp[0].value;
      }
      return temp.length > 0;
    }
  },
  computed: {
    ...mapState(["currentUser", "appbarStyle"])
  },
  watch: {
    "messages.length"(val) {
      this.messageGroup = this.$linq(this.messages)
        .groupBy("r=>(r.ctime+'').substring(0,10)")
        .select("{key:$.key(),value:$.toArray()}")
        .toArray();
    }
  },
  mounted() {}
};
</script>

