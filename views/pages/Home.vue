<template>
  <layoutmain>
    <mu-paper :z-depth="0" class="paper-title" :style="'height:'+titleWidth+'px;'">
      <residue :user="currentUser"></residue>
    </mu-paper>
    <mu-paper :z-depth="1" style="height:100px;padding:10px">
      <mu-text-field @blur="showMode" v-model="tallyForm.money" style="height:80px;width:50%" label="金额" prefix="￥" type="number" label-float></mu-text-field>
      <mu-auto-complete @change="showMode" v-model="tallyForm.consume" :data="consumes" style="height:80px;width:45%" label="种类" :max-search-results="5" open-on-focus label-float></mu-auto-complete>
    </mu-paper>
    <br/>
    <mu-list v-if="tallyList.length!=0">
      <mu-sub-header>
        最近消费
      </mu-sub-header>
      <mu-divider></mu-divider>
      <mu-list-item :key="index" v-for="item,index in tallyList" avatar :ripple="false">
        <mu-list-item-action>
          <mu-avatar>
            <img :src="item.headImg">
          </mu-avatar>
        </mu-list-item-action>
        <mu-list-item-content>
          <mu-list-item-title>{{item.remark}}</mu-list-item-title>
          <mu-list-item-sub-title>{{$format(item.ttime, "yyyy-MM-dd hh:mm")}}</mu-list-item-sub-title>
        </mu-list-item-content>
      </mu-list-item>
    </mu-list>
    <mu-list>
      <mu-sub-header>
        快速记录
      </mu-sub-header>
      <mu-divider></mu-divider>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="text-align: right;">
          15 元
          <br> 微信支出
        </mu-list-item-action>
        <mu-list-item-title style="text-align: center;">午餐</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon value="fast_forward"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
    </mu-list>
    <addtally :openTally.sync="openTally" :money="tallyForm.money" :consume="tallyForm.consume"></addtally>
  </layoutmain>
</template>
<script>
import layoutmain from "@/layout/main";
import residue from "@/components/residue";
import addtally from "@/components/addtally";
import { mapState } from "vuex";

export default {
  components: {
    layoutmain,
    residue,
    addtally
  },
  data() {
    return {
      tallyForm: {
        money: "",
        consume: ""
      },
      consumes: [],
      openTally: false,
      tallyList: [],
      titleWidth: document.body.clientWidth * 0.72
    };
  },
  methods: {
    showMode() {
      if (
        this.tallyForm.money.length <= 0 ||
        this.tallyForm.consume.length <= 0
      )
        return;
      let that = this;
      setTimeout(() => {
        that.openTally = true;
      }, 220);
      var isExist = false;
      this.consumes.forEach(x => {
        if (x == this.tallyForm.consume) {
          isExist = true;
        }
      });
      if (!isExist) {
        this.currentUser.consumes.push({
          content: this.tallyForm.consume,
          count: 1,
          default: ["收入", "支出", "预支"]
        });
      }
    },
    initLate() {
      var that = this;
      var types = this.$linq(this.currentUser.consumes)
        .select(x => x.content)
        .toArray();
      that.$axios
        .post("/tally/get", {
          uids: [this.currentUser.id],
          btime: new Date(
            new Date().getFullYear(),
            new Date().getMonth(),
            1
          ).toISOString(),
          etime: new Date(
            new Date().getFullYear(),
            new Date().getMonth() + 1,
            1
          ).toISOString(),
          bMoney: 0,
          eMoney: 999999,
          types: types,
          modes: [],
          channels: [],
          pageIndex: 1,
          pageSize: 3
        })
        .then(response => {
          if (response.code == 0) {
            if (response.data != null) that.tallyList = response.data;
          }
        });
    }
  },
  computed: {
    ...mapState(["currentUser"])
  },
  watch: {
    currentUser(val) {
      this.consumes = [];
      val.consumes.forEach(x => {
        this.consumes.push(x.content);
      });
      this.initLate();
    },
    openTally(val) {
      // todo 刷新时机待考虑
      if (!val) {
        this.tallyForm = {
          money: "",
          consume: ""
        };
        this.initLate();
        this.$store.dispatch("initUser", { $router: this.$router });
      }
    }
  }
};
</script>


<style scoped>
.button-right {
  margin-left: 60%;
}
.span-title {
  color: rgb(20, 20, 20);
  margin-right: 0px;
}
.span-money {
  font-size: 20px;
  font-weight: 600;
  color: #ff5722;
  margin-left: 0px;
}
.paper-title {
  margin-bottom: 10px;
  background-image: linear-gradient(60deg, #29323c 0%, #485563 100%);
  /* background-color: #efefef; */
}
</style>
