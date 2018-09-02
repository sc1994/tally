<template>
  <layoutmain :lefticon="'toc'" :leftevent="()=>open=!open" :letfstyle="open?'':'transform:rotate(180deg);'">
    <mu-container ref="container">
      <mu-load-more :loading="loading" @load="load">
        <mu-list>
          <template v-for="item,index in list">
            <mu-list-item avatar :ripple="false" :key="index">
              <mu-list-item-action>
                <mu-avatar>
                  <img :src="item.headImg">
                </mu-avatar>
              </mu-list-item-action>
              <mu-list-item-content>
                <mu-list-item-title>{{item.remark}}</mu-list-item-title>
                <mu-list-item-sub-title>{{$format(item.ctime, "yyyy-MM-dd hh:mm")}}</mu-list-item-sub-title>
              </mu-list-item-content>
            </mu-list-item>
            <mu-divider />
          </template>
        </mu-list>
      </mu-load-more>
    </mu-container>
    <mu-drawer :open.sync="open" letf :docked="false" style="top:55px;border-radius:0px;" :z-depth="0">
      <mu-list>
        <mu-list-item button>
          <mu-list-item-action>
            <i class="fa fa-pie-chart" style="font-size: 20px;"></i>
          </mu-list-item-action>
          <mu-list-item-title>资产配置</mu-list-item-title>
        </mu-list-item>
        <mu-list-item button>
          <mu-list-item-action>
            <i class="fa fa-bar-chart" style="font-size: 20px;"></i>
          </mu-list-item-action>
          <mu-list-item-title>柱状图统计</mu-list-item-title>
        </mu-list-item>
        <mu-list-item button>
          <mu-list-item-action>
            <i class="fa fa-area-chart" style="font-size: 20px;"></i>
          </mu-list-item-action>
          <mu-list-item-title>消费折线图</mu-list-item-title>
        </mu-list-item>
        <mu-list-item button>
          <mu-list-item-action>
            <i class="fa fa-location-arrow" style="font-size: 20px;"></i>
          </mu-list-item-action>
          <mu-list-item-title>费用去向</mu-list-item-title>
        </mu-list-item>
      </mu-list>
    </mu-drawer>
  </layoutmain>
</template>

<script>
import layoutmain from "@/layout/main";
import { mapState } from "vuex";

export default {
  components: {
    layoutmain
  },
  data() {
    return {
      open: false,
      list: [],
      loading: false,
      pageIndex: 1,
      notNextList: false
    };
  },
  methods: {
    load() {
      if (!this.notNextList) {
        console.log("下拉刷新");
        this.getList(this.pageIndex);
      }
    },
    getList(index) {
      var that = this;
      that.loading = true;
      that.$axios
        .post("/gettallybyuser", {
          token: localStorage.getItem("token"),
          pageIndex: that.pageIndex,
          pageSize: 12
        })
        .then(response => {
          if (response.result) {
            if (response.body != null) {
              that.pageIndex++;
              that.list.push(...response.body);
              if (response.body.length < 12) {
                that.notNextList = true;
              }
            } else {
              that.notNextList = true;
            }
            that.loading = false;
          }
        });
    },
    getUserHeadImg(id) {
      if (!this.currentUser) return "";
      if (!this.currentUser.partners) return "";
      for (var i = 0; i < this.currentUser.partners.length; i++) {
        if (this.currentUser.partners[i].id == id) {
          return this.currentUser.partners[i].headImg;
        }
      }
      return this.currentUser.headImg;
    }
  },
  computed: {
    ...mapState(["currentUser"])
  },
  mounted() {
    this.getList(this.pageIndex);
  }
};
</script>

<style>
</style>
