<template>
  <canvas id="mountNode" width="356" height="246" style="width: 356px; height: 246.813px;"></canvas>
</template>

<script>
import F2 from "@antv/f2";

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
    return {};
  },
  mounted() {
    this.$nextTick(() => {
      var Shape = F2.Shape;
      var data = [
        {
          pointer: "剩余预算\n\n已消费 60 元\n\n预支 10 元",
          value: 500,
          length: 0,
          y: 1
        }
      ];
      //自定义绘制数据的的形状
      Shape.registerShape("point", "dashBoard", {
        getPoints: function getPoints(cfg) {
          var x = cfg.x;
          var y = cfg.y;
          return [
            {
              x: x,
              y: y
            },
            {
              x: x,
              y: 0.4
            }
          ];
        },
        draw: function draw(cfg, container) {
          var point1 = cfg.points[0];
          var point2 = cfg.points[1];
          point1 = this.parsePoint(point1);
          point2 = this.parsePoint(point2);

          var line = container.addShape("Polyline", {
            attrs: {
              points: [point1, point2],
              stroke: "#1890FF",
              lineWidth: 2
            }
          });

          var text = cfg.origin._origin.value.toString();
          var text1 = container.addShape("Text", {
            attrs: {
              text: text + " 元",
              x: cfg.center.x,
              y: cfg.center.y,
              fill: "#1890FF",
              fontSize: 22,
              textAlign: "center",
              textBaseline: "bottom"
            }
          });
          var text2 = container.addShape("Text", {
            attrs: {
              text: cfg.origin._origin.pointer,
              x: cfg.center.x,
              y: cfg.center.y,
              fillStyle: "#ccc",
              textAlign: "center",
              textBaseline: "top"
            }
          });

          return [line, text1, text2];
        }
      });
      var chart = new F2.Chart({
        id: "mountNode",
        animate: false,
        pixelRatio: window.devicePixelRatio
      });
      var ticks = [];
      for (var i = 0; i <= 9987; i++) {
        if (i % 5 == 0) {
          ticks.push(i);
        }
      }
      chart.source(data, {
        value: {
          type: "linear",
          min: 0,
          max: 15,
          ticks: ticks,
          nice: false
        },
        length: {
          type: "linear",
          min: 0,
          max: 10
        },
        y: {
          type: "linear",
          min: 0,
          max: 1
        }
      });

      chart.coord("polar", {
        inner: 0,
        startAngle: -1.25 * Math.PI,
        endAngle: 0.25 * Math.PI
      });

      //配置value轴刻度线
      chart.axis("value", {
        tickLine: {
          strokeStyle: "#ccc",
          lineWidth: 2,
          length: -5
        },
        label: null,
        grid: null,
        line: null
      });

      chart.axis("y", false);

      //绘制仪表盘辅助元素
      chart.guide().arc({
        start: [0, 1.05],
        end: [4.8, 1.05],
        style: {
          strokeStyle: "#1890FF",
          lineWidth: 5,
          lineCap: "round"
        }
      });
      chart.guide().arc({
        start: [5.2, 1.05],
        end: [9.8, 1.05],
        style: {
          strokeStyle: "#ccc",
          lineWidth: 5,
          lineCap: "round"
        }
      });
      chart.guide().arc({
        start: [10.2, 1.05],
        end: [15, 1.05],
        style: {
          strokeStyle: "#ccc",
          lineWidth: 5,
          lineCap: "round"
        }
      });
      chart.guide().arc({
        start: [0, 1.2],
        end: [15, 1.2],
        style: {
          strokeStyle: "#ccc",
          lineWidth: 1
        }
      });

      chart.guide().text({
        position: [-0.5, 1.3],
        content: "0.00%",
        style: {
          fillStyle: "#ccc",
          font: "18px Arial",
          textAlign: "center"
        }
      });
      chart.guide().text({
        position: [7.5, 0.7],
        content: "7.50%",
        style: {
          fillStyle: "#ccc",
          font: "18px Arial",
          textAlign: "center"
        }
      });
      chart.guide().text({
        position: [15.5, 1.3],
        content: "15.00%",
        style: {
          fillStyle: "#ccc",
          font: "18px Arial",
          textAlign: "center"
        }
      });

      chart
        .point()
        .position("value*y")
        .size("length")
        .color("#1890FF")
        .shape("dashBoard");
      chart.render();
    });
  }
};
</script>
