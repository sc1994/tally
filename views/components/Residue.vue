<template>
  <canvas id="mountNode" height="246" style="width: 100%"></canvas>
</template>

<script>
import F2 from "@antv/f2";
import { mapState } from "vuex";

export default {
  props: {
    user: {
      type: Object,
      required: true
    }
  },
  data() {
    return {};
  },
  watch: {
    user(val) {
      var ticks = [];
      var mulriple = parseInt((val.budget / 50).toFixed(1));
      for (let index = 0; index <= val.budget; index++) {
        if (index % mulriple == 0) {
          ticks.push(index);
        }
      }
      ticks.push(val.budget);

      var Shape = F2.Shape;
      var data = [
        {
          pointer: `\n剩余预算 ${parseFloat(
            val.budget - val.haveBeenUsed
          ).toFixed(1)} 元\n\n预支 ${val.haveBeenAdvance} 元`,
          value: parseFloat(val.haveBeenUsed).toFixed(1),
          length: 0,
          y: 1
        }
      ];
      //自定义绘制数据的的形状
      Shape.registerShape("point", "dashBoard", {
        getPoints: function getPoints(cfg) {
          var x = cfg.x;
          if (x > 1) {
            x = 1;
          }
          if (x < 0) {
            x = 0;
          }
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
              stroke: "#d84315",
              lineWidth: 2
            }
          });

          var text = cfg.origin._origin.value.toString();
          var text1 = container.addShape("Text", {
            attrs: {
              text: text + " 元",
              x: cfg.center.x,
              y: cfg.center.y,
              fill: "#d84315",
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

      chart.source(data, {
        value: {
          type: "linear",
          min: 0,
          max: val.budget,
          ticks: ticks,
          nice: false
        },
        length: {
          type: "linear",
          min: 0,
          max: 1
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
          lineWidth: 3,
          length: -6
        },
        label: null,
        grid: null,
        line: null
      });

      chart.axis("y", false);
      var userEnd = val.haveBeenUsed - mulriple / 3;
      userEnd = userEnd > 10 ? userEnd : 10;
      userEnd = userEnd > val.budget ? val.budget - 1 : userEnd;
      //绘制仪表盘辅助元素
      if (val.haveBeenUsed > 0) {
        chart.guide().arc({
          start: [10, 1.05],
          end: [userEnd, 1.05],
          style: {
            strokeStyle: "#d84315",
            lineWidth: 4,
            lineCap: "round"
          }
        });
      }

      chart.guide().arc({
        start: [userEnd + mulriple / 3, 1.05],
        end: [val.budget - 10, 1.05],
        style: {
          strokeStyle: "#1890FF",
          lineWidth: 4,
          lineCap: "round"
        }
      });

      // 周边辅助信息
      chart.guide().text({
        position: [val.budget, 1.2],
        content: val.budget + "",
        style: {
          fillStyle: "#ccc",
          font: "18px Arial",
          textAlign: "center"
        }
      });

      chart.guide().text({
        position: [0, 1.2],
        content: "0",
        style: {
          fillStyle: "#ccc",
          font: "18px Arial",
          textAlign: "center"
        }
      });

      chart.guide().text({
        position: [val.budget / 2, 1.16],
        content: val.budget / 2 + "",
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
    }
  }
};
</script>
