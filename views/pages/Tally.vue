<template>
  <layoutmain :lefticon="'toc'" :leftevent="()=>openDrawer=!openDrawer" :letfstyle="openDrawer?'':'transform:rotate(180deg);'" :righticon="'settings'" :rightevent="()=>openSetDrawer=!openSetDrawer" :rightstyle="openSetDrawer?'':'transform:rotate(360deg);'">
    <mu-container ref="container">
      <mu-load-more :loading="loading" @load="load">
        <template v-for="list in group">
          <mu-list textline="two-line">
            <mu-sub-header>{{list.key}}</mu-sub-header>
            <mu-list-item avatar :ripple="false" v-for="item,index in list.value" :key="index" v-if="item.tid">
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
                <mu-list-item-after-text>{{$format(item.ttime, "hh:mm")}}</mu-list-item-after-text>
                <mu-button icon color="#2196f3" @click="currentTally=item;openSetTally=true" v-if="item.uid==currentUser.id">
                  <mu-icon value="rate_review"></mu-icon>
                </mu-button>
              </mu-list-item-action>
            </mu-list-item>
          </mu-list>
          <mu-divider />
        </template>
      </mu-load-more>
    </mu-container>
    <settally :open.sync="openSetTally" :currentItem="currentTally"></settally>
    <tallysetdrawer :open.sync="openSetDrawer" :searchForm.sync="searchForm"></tallysetdrawer>
    <tallydrawer :open.sync="openDrawer"></tallydrawer>
  </layoutmain>
</template>

<script>
import layoutmain from "@/layout/main";
import settally from "@/components/settally";
import tallydrawer from "@/components/tallydrawer";
import tallysetdrawer from "@/components/tallysetdrawer";
import { mapState } from "vuex";

export default {
  components: {
    layoutmain,
    settally,
    tallydrawer,
    tallysetdrawer
  },
  data() {
    return {
      openDrawer: false,
      openSetDrawer: false,
      group: [],
      loading: false,
      pageIndex: 1,
      notNextList: false,
      openSetTally: false,
      currentTally: {},
      searchForm: {
        partners: [],
        channels: [],
        types: [],
        modes: ["收入", "支出", "预支"],
        ttime: new Date()
      }
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
      var types = [];
      if (!this.searchForm.consumes) {
        types = this.$linq(this.currentUser.consumes)
          .select(x => x.content)
          .toArray();
      } else {
        types = this.$linq(this.searchForm.consumes)
          .where(x => x.selected)
          .select(x => x.content)
          .toArray();
      }
      var date = new Date(this.searchForm.ttime);
      that.$axios
        .post("/tally/get", {
          uids: this.searchForm.partners,
          btime: new Date(date.getFullYear(), date.getMonth(), 1).toISOString(),
          etime: new Date(
            date.getFullYear(),
            date.getMonth() + 1,
            1
          ).toISOString(),
          bMoney: 0,
          eMoney: 999999,
          types: types,
          modes: this.searchForm.modes,
          pageIndex: that.pageIndex,
          pageSize: 12
        })
        .then(response => {
          if (response.code == 0) {
            if (response.data != null) {
              that.pageIndex++;
              var g = that
                .$linq(response.data)
                .groupBy("r=>(r.ttime+'').substring(0,10)")
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
              if (response.data.length < 10) {
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
    },
    removeItem(id) {
      console.log(id);
    }
  },
  computed: {
    ...mapState(["currentUser"])
  },
  watch: {
    openDrawer(val) {
      if (val) this.openSetDrawer = !val;
    },
    openSetDrawer(val) {
      if (val) this.openDrawer = !val;
      else {
        this.group = [];
        this.pageIndex = 1;
        this.notNextList = false;
        this.getList(this.pageIndex);
      }
    },
    currentUser(val) {
      this.getList(this.pageIndex);
      var that = this;
      that.searchForm.partners.push(val.id);
      if (val.partners != null && val.partners) {
        val.partners.forEach(x => {
          that.searchForm.partners.push(x.id);
        });
      }
      // val.channels.forEach(x => {
      //   that.searchForm.channels.push(x.content);
      // });
      this.searchForm.consumes = JSON.parse(JSON.stringify(val.consumes));
      this.searchForm.consumes.forEach(x => (x.selected = true)); // 默认选中
    }
  }
};
</script>

<style>
</style>
