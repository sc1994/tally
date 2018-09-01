<template>
  <minor :title="'消息中心'" :righticon="righticon" :rightevent="readlyAll">
    <messagelist :list="todays"></messagelist>
    <messagelist :list="yesterdays"></messagelist>
    <messagelist :list="earlier"></messagelist>
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
import messagelist from "@/components/messagelist";

export default {
  components: {
    minor,
    messagelist
  },
  data() {
    return {
      openAlert: false,
      messageGroup: [],
      currentItem: {},
      righticon: ""
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
            that.righticon =
              that
                .$linq(response.result)
                .where(x => !x.needTouch && x.status == 1)
                .toArray().length > 0
                ? "done_all"
                : "";
          that.messageGroup = that
            .$linq(response.result)
            .groupBy("r=>(r.ctime+'').substring(0,10)")
            .select("{key:$.key(),value:$.toArray()}")
            .toArray();
          // todo 按今天昨天更早来分页
        });
    },
    readlyAll() {
      var that = this;
      var loading = that.$loading();
      that.$axios
        .get("/readmessageall/" + that.currentUser.id)
        .then(response => {
          if (response.result) {
            that.$toast.success("无需确认的消息, 已标记为已读");
            that.$router.back(-1);
          }
          loading.close();
        })
        .catch(error => {
          loading.close();
        });
    }
  },
  computed: {
    ...mapState(["currentUser"]),
    todays() {
      var that = this;
      var temp = this.$linq(this.messageGroup)
        .where(
          x => x.key == that.$format(new Date(), "yyyy-MM-dd").substring(0, 10)
        )
        .toArray();
      if (temp.length > 0) {
        return temp[0].value;
      }
      return [];
    },
    yesterdays() {
      var that = this;
      var now = new Date();
      now.setDate(now.getDate() - 1);
      var temp = this.$linq(this.messageGroup)
        .where(x => x.key == that.$format(now, "yyyy-MM-dd").substring(0, 10))
        .toArray();
      if (temp.length > 0) {
        return temp[0].value;
      }
      return [];
    },
    earlier() {
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
      var earline = [];
      temp.forEach(x => {
        x.forEach(i => {
          earline.push(i);
        });
      });
      return earline;
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

