<template>
  <dialoghead :title="'预支设置'" :open.sync="thatOpen" :buttonclick="submit" buttonicon="done">
    <mu-container>
      <mu-text-field v-model="item.content" label="名称" full-width placeholder="如:xxx信用卡"></mu-text-field>
    </mu-container>
    <mu-sub-header>还款日期</mu-sub-header>
    <div style="div-slide">
      <mu-slide-picker :slots="dateSlots" :visible-item-count="5" :values="item.rdate" @change="changerdate"></mu-slide-picker>
    </div>
  </dialoghead>
</template>

<script>
import dialoghead from "@/layout/dialog";

export default {
  props: {
    open: {
      type: Boolean
    },
    item: {
      type: Object
    }
  },
  components: {
    dialoghead
  },
  data() {
    return {
      thatOpen: false,
      name: "",
      dateSlots: [
        {
          width: "100%",
          textAlign: "right",
          values: ["当月", "次月"]
        },
        {
          width: "100%",
          textAlign: "left",
          values: (function() {
            var result = [];
            for (var i = 1; i <= 31; i++) {
              result.push(i + "日");
            }
            result.push("月底");
            return result;
          })()
        }
      ]
    };
  },
  methods: {
    changerdate(value, index) {
      this.item.rdate[index] = value;
      if (index == 0) {
        this.item.rdate = [value, "1日"];
      } else if (index == 1) {
        this.item.rdate = [this.item.rdate[0], value];
      }
    },
    submit() {
      var that = this;
      var loading = this.$loading({});
      this.item.isAdvance = true; // 一些值需要写死
      this.item.default = ["预支"];
      var url = "";
      if (this.item.id) {
        url = "/channel/set";
      } else {
        url = "/channel/add";
      }
      this.$axios.post(url, this.item).then(response => {
        if (response.code == 0) {
          setTimeout(() => {
            that.thatOpen = false;
          }, 500);
        }
        loading.close();
      });
    }
  },
  watch: {
    open(val) {
      if (val) {
        this.thatOpen = true;
      }
    },
    thatOpen(val) {
      if (!val) {
        this.$emit("update:open", false);
      }
    },
    item(val) {
      if (!val.rdate) {
        val.rdate = ["当月", "1日"];
      }
    }
  }
};
</script>

<style>
.div-slide {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  height: 100%;
  width: 100%;
}
</style>

