<template>
  <div class="chart-container">
    <div ref="chartInnerRef" class="chart"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch, nextTick } from "vue";
import * as echarts from "echarts";
import { useGraphRenderer } from "./useGraphRenderer";
import {parser} from "../../wailsjs/go/models";

const props = defineProps<{
  data: parser.ParseResult | null;
  optimized: boolean;
}>();

const chartInnerRef = ref<HTMLDivElement | null>(null);
let chartInstance: echarts.ECharts | null = null;

const { renderChart } = useGraphRenderer();

onMounted(async () => {
  await nextTick();
  chartInstance = echarts.init(chartInnerRef.value!);
  window.addEventListener("resize", () => chartInstance?.resize());

  if (props.data) {
    renderChart(chartInstance, props.data, props.optimized);
  }
});

watch(
    () => [props.data, props.optimized],
    () => {
      if (!chartInstance) return;

      if (!props.data) {
        chartInstance.clear();
      } else {
        renderChart(chartInstance, props.data, props.optimized);
      }
    }
);

onBeforeUnmount(() => {
  chartInstance?.dispose();
});
</script>

<style scoped>
.chart-container {
  flex: 1;
  display: flex;
}
.chart {
  flex: 1;
  background-color: #fafafa;
  border: 1px solid #ccc;
}
</style>
