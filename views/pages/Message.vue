<template>
  <minor :title="'消息中心'" :righticon="righticon" :rightevent="readlyAll">
    <messagelist :list="todays"></messagelist>
    <messagelist :list="yesterdays"></messagelist>
    <messagelist :list="earlier"></messagelist>
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
      messageGroup: [],
      righticon: "",
      messages: []
    };
  },
  methods: {
    submit() {
      var that = this;
      var loading = that.$loading({});
      this.$axios
        .post("/agreemessage", that.currentItem)
        .then(response => {
          if (response.code == 0) {
            that.$toast.success("你们已经是小伙伴啦");
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
        .post("/message/get", {
          type: 0,
          status: 0,
          pageIndex: 1,
          pageSize: 100
        })
        .then(response => {
          if (response.code == 0) {
            that.messages = response.data;
            that.righticon =
              that
                .$linq(response.data)
                .where(x => !x.needTouch && x.status == 1)
                .toArray().length > 0
                ? "done_all"
                : "";
            that.messageGroup = that
              .$linq(response.data)
              .groupBy("r=>(r.ctime+'').substring(0,10)")
              .select("{key:$.key(),value:$.toArray()}")
              .toArray();
          }
        });
    },
    readlyAll() {
      var that = this;
      var loading = that.$loading();
      var ids = that
        .$linq(that.messages)
        .select(x => x.id)
        .toArray();
      that.$axios
        .post("message/set", {
          ids: ids,
          status: 2
        })
        .then(response => {
          if (response.code == 0) {
            // that.$toast.success("无需确认的消息, 已标记为已读");
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

