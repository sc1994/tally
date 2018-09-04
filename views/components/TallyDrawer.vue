<template>
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
</template>

<script>
export default {
  props: {
    open: {
      type: Boolean,
      default: false
    }
  }
};
</script>

<style>
</style>
