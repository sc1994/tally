<template>
  <div>
    <mu-paper :z-depth="0" class="paper-title">
      <residue :height="165" :residue="90" style="margin-top: -10px;float:left;"></residue>
      <div style="margin-top: 10px;">
        <span class="span-title">当月剩余：
          <span class="span-money">50,203</span> 元</span>
        <br /><br />
        <span class="span-title">当月消费：
          <span class="span-money" style="color:#4caf50">356</span> 元</span>
        <br /><br />
        <span class="span-title">当月预支：
          <span class="span-money" style="color:#4caf50">478</span> 元</span>
      </div>
    </mu-paper>
    <mu-paper :z-depth="1" style="height:100px;padding:10px">
      <mu-text-field @blur="showMode" v-model="tallyForm.money" style="height:80px;width:50%" label="金额" prefix="￥" type="number" label-float></mu-text-field>
      <mu-auto-complete @change="showMode" v-model="tallyForm.consume" :data="consumes" style="height:80px;width:45%" label="种类" :max-search-results="5" open-on-focus label-float></mu-auto-complete>
    </mu-paper>
    <br/>
    <mu-list>
      <mu-sub-header>
        最近消费
        <mu-button flat color="primary" class="button-right">查看全部</mu-button>
      </mu-sub-header>
      <mu-divider></mu-divider>
      <mu-list-item :key="index" v-for="item,index in tallyList" avatar :ripple="false">
        <mu-list-item-action>
          <mu-avatar>
            <img src="/static/head-default.png">
          </mu-avatar>
        </mu-list-item-action>
        <mu-list-item-content>
          <mu-list-item-title>{{item.remark}}</mu-list-item-title>
          <mu-list-item-sub-title>{{formatDate(item.ctime)}}</mu-list-item-sub-title>
        </mu-list-item-content>
      </mu-list-item>
    </mu-list>
    <mu-list>
      <mu-sub-header>
        快速记录
        <mu-button flat color="primary" class="button-right">添加更多</mu-button>
      </mu-sub-header>
      <mu-divider></mu-divider>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          15 元
        </mu-list-item-action>
        <mu-list-item-title>午餐</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          15 元
        </mu-list-item-action>
        <mu-list-item-title>午餐</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          15 元
        </mu-list-item-action>
        <mu-list-item-title>午餐</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          15 元
        </mu-list-item-action>
        <mu-list-item-title>午餐</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          15 元
        </mu-list-item-action>
        <mu-list-item-title>午餐</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          15 元
        </mu-list-item-action>
        <mu-list-item-title>午餐</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          8 元
        </mu-list-item-action>
        <mu-list-item-title>公交</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          8 元
        </mu-list-item-action>
        <mu-list-item-title>公交</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <mu-list-item avatar button :ripple="true">
        <mu-list-item-action style="width:280px">
          8 元
        </mu-list-item-action>
        <mu-list-item-title>公交</mu-list-item-title>
        <mu-list-item-action>
          <mu-button icon color="#4caf50">
            <mu-icon right value="flight_takeoff"></mu-icon>
          </mu-button>
        </mu-list-item-action>
      </mu-list-item>
      <!-- 留个空的撑开样式 -->
      <mu-list-item avatar button :ripple="false">
        <mu-list-item-action style="width:280px">
        </mu-list-item-action>
        <mu-list-item-title></mu-list-item-title>
        <mu-list-item-action>
        </mu-list-item-action>
      </mu-list-item>
    </mu-list>
    <addtally :openTally.sync="openTally" :step.sync="step" :money="tallyForm.money" :consume="tallyForm.consume"></addtally>
  </div>
</template>

<script>
import residue from "@/components/residue";
import addtally from "@/components/addtally";
import { mapState } from "vuex";

export default {
  components: {
    residue,
    addtally
  },
  data() {
    return {
      tallyForm: {
        money: "",
        consume: ""
      },
      step: 0,
      consumes: [],
      openTally: false,
      tallyList: []
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
      }, 180);
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
      // 最近消费
    },
    formatDate() {
      // 日期格式化
      return "1-1-1-";
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
    },
    openTally(val) {
      if (!val) {
        this.tallyForm.mode = "";
        this.tallyForm.channel = "";
      }
    }
  }
};
</script>


<style scoped>
.button-right {
  margin-left: 55%;
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
  height: 180px;
  padding: 10px;
  margin-bottom: 10px;
  background-image: linear-gradient(
    to top,
    #a7a7a7 0%,
    lightgrey 1%,
    #e0e0e0 26%,
    #efefef 48%,
    #d9d9d9 75%,
    #a7a7a7 100%
  );
}
</style>
