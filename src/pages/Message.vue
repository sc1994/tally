<template>
  <minor :title="'消息中心'">
    <mu-list textline="two-line">
      <mu-sub-header v-if="isToday">今天</mu-sub-header>
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
          <mu-button v-if="item.type==1&&item.status==1" icon color="primary" @click="openAlert=true;currentItem=item">
            <mu-icon value="sentiment_satisfied"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <div v-if="isYesterday">
        <mu-divider></mu-divider>
        <mu-sub-header>昨天</mu-sub-header>
      </div>
      <mu-list-item avatar :ripple="false" button v-for="item in yesterdays" :key="item.id">
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
          <mu-button v-if="item.type==1&&item.status==1" icon color="primary" @click="openAlert=true;currentItem=item">
            <mu-icon value="sentiment_satisfied"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <div v-if="isEarlier">
        <mu-divider></mu-divider>
        <mu-sub-header>更早</mu-sub-header>
      </div>
      <mu-list-item avatar :ripple="false" button v-for="item in earlier" :key="item.id">
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
          <mu-button v-if="item.type==1&&item.status==1" icon color="primary" @click="openAlert=true;currentItem=item">
            <mu-icon value="sentiment_satisfied"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
    </mu-list>
    <mu-dialog title="是否同意 ?" width="75%" max-width="75%" :esc-press-close="false" :overlay-close="false" :open.sync="openAlert">
      <span style="color: rgba(0,0,0,.87);">{{currentItem.fnick}}&nbsp;-&nbsp;</span>
      {{currentItem.content}}
      <mu-button slot="actions" flat color="primary" @click="openAlert=false">算了</mu-button>
      <mu-button slot="actions" flat color="primary" @click="submit">没问题</mu-button>
    </mu-dialog>
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
      openAlert: false,
      messageGroup: [],
      todays: [],
      yesterdays: [],
      earlier: [],
      currentItem: {}
    };
  },
  methods: {
    submit() {
      var that = this;
      var loading = that.$loading({});
      this.$axios
        .post("/agreemessage", that.currentItem)
        .then(response => {
          if (response.result) {
            that.$toast.success("你们已经是小伙伴啦");
          } else {
            that.$toast.error("网络异常,请重试");
          }
          loading.close();
          this.openAlert = false;
        })
        .catch(() => {
          loading.close();
          this.openAlert = false;
        });
    },
    search(index) {
      var that = this;
      that.$axios
        .get(`/getmessage/${that.currentUser.id}/${index}/100`)
        .then(response => {
          if (response.result != null)
            that.messageGroup = that
              .$linq(response.result)
              .groupBy("r=>(r.ctime+'').substring(0,10)")
              .select("{key:$.key(),value:$.toArray()}")
              .toArray();
          // todo 按今天昨天更早来分页
        })
    }
  },
  computed: {
    ...mapState(["currentUser"]),
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
      now.setDate(now.getDate() - 1);
      var temp = this.$linq(this.messageGroup)
        .where(x => x.key == that.$format(now, "yyyy-MM-dd").substring(0, 10))
        .toArray();
      if (temp.length > 0) {
        this.yesterdays = temp[0].value;
      }
      return temp.length > 0;
    },
    isEarlier() {
      var that = this;
      var now = new Date();
      now.setDate(now.getDate() - 1);
      var temp = this.$linq(this.messageGroup)
        .where(
          x =>
            new Date(x.key) <
            new Date(that.$format(now, "yyyy-MM-dd").substring(0, 10))
        )
        .select(x => x.value)
        .toArray();
      temp.forEach(x => {
        x.forEach(i => {
          that.earlier.push(i);
        });
      });
      return temp.length > 0;
    }
  },
  watch: {
    currentUser() {
      this.search(1);
    }
  },
  mounted() {}
};
</script>

