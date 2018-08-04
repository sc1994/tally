<template>
  <mu-snackbar :color="type" :open.sync="thatShow">
    <mu-icon left :value="icon"></mu-icon>
    {{msg}}
    <mu-button flat slot="action" color="#fff" @click="thatClose">关闭</mu-button>
  </mu-snackbar>
</template>

<script>
export default {
  props: {
    msg: {
      type: String,
      required: true
    },
    type: {
      type: String,
      required: true
    },
    show: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      thatShow: this.show
    }
  },
  methods: {
    thatClose() {
      this.$emit('update:show', false)
      this.thatShow = false
    }
  },
  computed: {
    icon() {
      return {
        success: 'check_circle',
        info: 'info',
        warning: 'priority_high',
        error: 'warning'
      }[this.thatType]
    }
  },
  watch: {
    show(val) {
      var that = this
      that.thatShow = that.show
      if (val) {
        setTimeout(() => {
          that.thatClose()
        }, 2000)
      }
    }
  }
}
</script>

