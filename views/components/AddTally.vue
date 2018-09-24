<template>
  <dialoghead :title="consume+'：'+money+' 元'" :open.sync="thatOpenTally" :buttonclick="submit" :buttonicon="buttonicon">
    <mu-stepper :active-step="thatStep" orientation="vertical">
      <mu-step>
        <mu-step-label>
          收入/支出/预支
          <b>{{tallyForm.mode.length>0 ? " ：":""}}{{tallyForm.mode}}</b>
        </mu-step-label>
        <mu-step-content>
          <mu-list>
            <mu-list-item v-if="item.content.length>0 && !item.hide" :key="index" v-for="item,index in manyType.modes">
              <mu-list-item-content>
                <mu-radio :label="item.content" :value="item.content" v-model="tallyForm.mode" @change="thatStep++"></mu-radio>
              </mu-list-item-content>
            </mu-list-item>
          </mu-list>
        </mu-step-content>
      </mu-step>
      <mu-step>
        <mu-step-label>
          方式
          <b>{{tallyForm.channel.length>0 ? " ：":""}}{{tallyForm.channel}}</b>
        </mu-step-label>
        <mu-step-content>
          <mu-list>
            <mu-list-item v-if="item.content.length>0 && !item.hide" :key="index" v-for="item,index in manyType.channels">
              <mu-list-item-content>
                <mu-radio :label="item.content" :value="item.content" v-model="tallyForm.channel" @change="thatStep++"></mu-radio>
              </mu-list-item-content>
            </mu-list-item>
          </mu-list>
          <mu-button flat @click="thatStep--" color="primary">上一步</mu-button>
        </mu-step-content>
      </mu-step>
      <mu-step>
        <mu-step-label>
          其他
        </mu-step-label>
        <mu-step-content>
          <br/>
          <mu-form :model="tallyForm" label-position="top" label-width="70">
            <mu-form-item prop="date" label="日期">
              <mu-date-input v-model="tallyForm.date" type="dateTime" actions style="height: 30px;"></mu-date-input>
            </mu-form-item>
            <mu-form-item prop="input" label="备注">
              <mu-text-field v-model="tallyForm.remark" style="height: 30px;"></mu-text-field>
            </mu-form-item>
            <mu-form-item>
              <mu-button flat @click="thatStep--" color="primary">上一步</mu-button>
            </mu-form-item>
          </mu-form>
        </mu-step-content>
      </mu-step>
    </mu-stepper>
  </dialoghead>
</template>
 
 <script>
import { mapState } from "vuex";
import dialoghead from "@/layout/dialog";

export default {
  components: {
    dialoghead
  },
  props: {
    openTally: {
      type: Boolean,
      required: true
    },
    consume: {
      type: String,
      required: true
    },
    money: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      thatOpenTally: false,
      thatStep: -1,
      tallyForm: {
        mode: "",
        channel: "",
        date: "",
        remark: ""
      },
      manyType: {
        channels: [],
        modes: [{ content: "收入" }, { content: "支出" }, { content: "预支" }],
        consumes: []
      },
      buttonicon: ""
    };
  },
  computed: {
    ...mapState(["currentUser"])
  },
  methods: {
    submit() {
      var that = this;
      var loading = that.$loading({});
      that.$axios
        .post("/tally/add", {
          money: parseFloat(that.money),
          type: that.consume,
          mode: that.tallyForm.mode,
          channel: that.tallyForm.channel,
          remark: that.tallyForm.remark,
          ttime: new Date(that.tallyForm.date).toISOString()
        })
        .then(response => {
          if (response.code == 0) {
            that.$toast.success("添加成功");
            setTimeout(() => {
              that.$emit("update:openTally", false);
              that.thatOpenTally = false;
              that.$emit("update:consume", "");
              that.$emit("update:money", "");
            }, 1200);
          }
          loading.close();
        });
    }
  },
  watch: {
    currentUser(val) {
      this.manyType.channels = [
        {
          content: "",
          default: []
        },
        ...val.channels
      ];
      this.manyType.consumes = this.$linq(val.consumes)
        .select(x => x.content)
        .toArray();
    },
    thatStep(val) {
      var that = this;
      this.buttonicon = val == 2 ? "done" : ""; // 是否显示完成按钮
      if (val == 0) {
        // 当前选中的类型
        var currentConsume = that
          .$linq(that.currentUser.consumes)
          .firstOrDefault(x => x.content == that.consume);

        if (currentConsume) {
          that.manyType.modes = that
            .$linq(that.manyType.modes)
            .select(x => {
              x.hide = currentConsume.default.indexOf(x.content) < 0;
              return x;
            })
            .toArray();
        }
        var hideMode = that
          .$linq(that.manyType.modes)
          .where(x => !x.hide)
          .toArray();
        if (hideMode.length == 1) {
          that.thatStep += 1;
          that.tallyForm.mode = hideMode[0].content;
        }
      } else if (val == 1) {
        if (this.tallyForm.mode != "预支") {
          this.tallyForm.channel = "余额";
          this.thatStep += 1;
        }
        that
          .$linq(that.manyType.channels)
          .select(x => {
            x.hide = x.default.indexOf(that.tallyForm.mode) < 0;
            return x;
          })
          .toArray();
      } else if (val == 2) {
        that.tallyForm.date = new Date();
        var flag = that.tallyForm.mode == "收入" ? "收取" : "支付";
        that.tallyForm.remark = `${that.consume}${that.tallyForm.mode}了${
          that.money
        }元，通过${that.tallyForm.channel}${flag}。`;
      }
    },
    thatOpenTally(val) {
      if (!val) {
        // 初始化记录步骤里面的数据,但是保留金额和类型
        this.thatStep = -1;
        this.tallyForm = {
          mode: "",
          channel: "",
          date: "",
          remark: ""
        };
        this.$emit("update:openTally", false);
      } else {
        this.thatStep = 0;
      }
    },
    openTally(val) {
      if (val) {
        this.thatOpenTally = true;
      }
    }
  }
};
</script>
 