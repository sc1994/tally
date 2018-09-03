<template>
  <layoutmain :lefticon="'toc'" :leftevent="()=>open=!open" :letfstyle="open?'':'transform:rotate(180deg);'">
    <mu-container ref="container">
      <mu-load-more :loading="loading" @load="load">
        <template v-for="list in group">
          <mu-list textline="two-line">
            <mu-sub-header>{{list.key}}</mu-sub-header>
            <mu-list-item avatar :ripple="false" v-for="item,index in list.value" :key="index">
              <mu-list-item-action>
                <mu-avatar>
                  <img :src="item.headImg">
                </mu-avatar>
              </mu-list-item-action>
              <mu-list-item-content>
                <mu-list-item-title>{{item.money}}元
                  <span class="mu-item-sub-title">&nbsp;-&nbsp;{{item.type}}</span>
                </mu-list-item-title>
                <mu-list-item-sub-title>
                  <span style="color:black;">{{item.channel}}</span>
                  &nbsp;-&nbsp;{{item.mode}}
                </mu-list-item-sub-title>
              </mu-list-item-content>
              <mu-list-item-action>
                <mu-list-item-after-text>{{$format(item.ctime, "hh:mm")}}</mu-list-item-after-text>
                <mu-button icon color="#2196f3" @click="currentTally=item;openSetTally=true" v-if="item.userID==currentUser.id">
                  <mu-icon value="edit_location"></mu-icon>
                </mu-button>
              </mu-list-item-action>
            </mu-list-item>
          </mu-list>
          <mu-divider />
        </template>
      </mu-load-more>
    </mu-container>
    <mu-drawer :open.sync="open" letf :docked="true" style="top:55px;border-radius:0px;" :z-depth="0">
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
    <settally :open.sync="openSetTally" :currentItem="currentTally"></settally>
  </layoutmain>
</template>

<script>
import layoutmain from "@/layout/main";
import settally from "@/components/settally";
import { mapState } from "vuex";

export default {
  components: {
    layoutmain,
    settally
  },
  data() {
    return {
      open: false,
      group: [],
      loading: false,
      pageIndex: 1,
      notNextList: false,
      openSetTally: false,
      currentTally: {}
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
          pageSize: 10,
          onlyMe: false
        })
        .then(response => {
          if (response.result) {
            if (response.body != null) {
              that.pageIndex++;
              var g = that
                .$linq(response.body)
                .groupBy("r=>(r.ctime+'').substring(0,10)")
                .select("{key:$.key(),value:$.toArray()}")
                .toArray();
              if (that.group.length > 0 && g.length > 0) {
                if (that.group[that.group.length - 1].key == g[0].key) {
                  that.group[that.group.length - 1].value.push(...g[0].value);
                  g = g.slice(1);
                }
              }
              if (g.length > 0) {
                that.group.push(...g);
              }
              if (response.body.length < 10) {
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
