<template>
  <canvas id="mountNode"></canvas>
</template>

<script>
import F2 from '@antv/f2'

export default {
  props: {
    height: {
      type: Number,
      required: true
    },
    residue: {
      type: Number,
      required: true
    }
  },
  data() {
    return {}
  },
  mounted() {
    this.$nextTick(() => {
      var that = this
      var data = [
        {
          x: '1',
          y: that.residue
        }
      ]
      var chart = new F2.Chart({
        id: 'mountNode',
        pixelRatio: window.devicePixelRatio,
        height: that.height,
        width: that.height
      })
      chart.source(data, {
        y: {
          max: 100,
          min: 0
        }
      })
      chart.axis(false)
      chart.tooltip(false)
      chart.coord('polar', {
        transposed: true,
        innerRadius: 0.9,
        radius: 1
      })
      chart.guide().arc({
        start: [0, 0],
        end: [1, 99.98],
        top: false,
        style: {
          lineWidth: 8,
          stroke: '#ccc'
        }
      })
      chart.guide().text({
        position: ['50%', '50%'],
        content: that.residue + '%',
        style: {
          fontSize: 24,
          fill: '#1890FF'
        }
      })
      chart
        .interval()
        .position('x*y')
        .size(8)
        .animate({
          appear: {
            duration: 2000,
            easing: 'cubicIn'
          }
        })
      chart.render()
    })
  }
}
</script>
