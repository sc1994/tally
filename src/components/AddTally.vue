<template>
  <mu-dialog fullscreen :open.sync="openTally">
    <mu-appbar color="primary" :title="consume+'：'+money+' 元'">
      <mu-button slot="left" icon @click="$emit('update:openTally', false)">
        <mu-icon value="close"></mu-icon>
      </mu-button>
      <mu-button slot="right" flat @click="submit" v-if="thatStep==2">
        完成
      </mu-button>
    </mu-appbar>
    <br/>
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
  </mu-dialog>
</template>
 
 <script>
import { mapState } from "vuex";

export default {
  props: {
    openTally: {
      type: Boolean,
      required: true
    },
    step: {
      type: Number,
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
      thatStep: this.step,
      thatOpenTally: this.openTally,
      tallyForm: {
        mode: "",
        channel: "",
        date: "",
        remark: ""
      },
      manyType: {
        channels: [],
        modes: [],
        consumes: []
      }
    };
  },
  computed: {
    ...mapState(["currentUser"])
  },
  methods: {
    submit() {}
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
      this.manyType.modes = [
        {
          content: "",
          default: []
        },
        ...val.modes
      ];
      this.manyType.consumes = [];
      val.consumes.forEach(x => {
        this.manyType.consumes.push(x.content);
      });
    },
    thatStep(val) {
      if (val == 2) {
        this.tallyForm.date = new Date();
        var flag = this.tallyForm.mode == "收入" ? "收取" : "支付";
        this.tallyForm.remark = `${this.consume}${this.tallyForm.mode}了${
          this.money
        }元，通过${this.tallyForm.channel}${flag}。`;
      }
      this.$emit("update:step", val);
    }
  }
};
</script>
 