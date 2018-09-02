<template>
  <mu-list textline="two-line" v-if="list.length>0">
    <mu-sub-header>{{title}}</mu-sub-header>
    <mu-list-item avatar :ripple="false" button v-for="item in list" :key="item.id">
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
        <mu-button icon color="error" v-if="!item.needTouch&&item.status==1" :ripple="false">
          <mu-icon value="lens" :size="1"></mu-icon>
        </mu-button>
      </mu-list-item-action>
    </mu-list-item>
    <mu-divider v-if="title!='更早'"></mu-divider>
    <mu-dialog title="是否同意 ?" width="75%" max-width="75%" :esc-press-close="false" :overlay-close="false" :open.sync="openAlert">
      <span style="color: rgba(0,0,0,.87);">{{currentItem.fnick}}&nbsp;-&nbsp;</span>
      {{currentItem.content}}
      <mu-button slot="actions" flat color="primary" @click="openAlert=false">算了</mu-button>
      <mu-button slot="actions" flat color="primary" @click="submit()">没问题</mu-button>
    </mu-dialog>
  </mu-list>
</template>

<script>
export default {
  props: {
    list: {
      type: Array,
      default: []
    }
  },
  data() {
    return {
      openAlert: false,
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
    }
  },
  computed: {
    title() {
      if (this.list.length < 1) return "";
      var now = new Date();
      if (
        (this.list[0].ctime + "").substring(0, 10) ==
        this.$format(now, "yyyy-MM-dd").substring(0, 10)
      ) {
        return "今天";
      }
      now.setDate(now.getDate() - 1);
      if (
        (this.list[0].ctime + "").substring(0, 10) ==
        this.$format(now, "yyyy-MM-dd").substring(0, 10)
      ) {
        return "昨天";
      }
      now.setDate(now.getDate() - 1);
      if (
        (this.list[0].ctime + "").substring(0, 10) ==
        this.$format(now, "yyyy-MM-dd").substring(0, 10)
      ) {
        return "更早";
      }
    }
  }
};
</script>

