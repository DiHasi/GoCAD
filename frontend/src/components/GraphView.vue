<script setup lang="ts">
import {nextTick, ref} from "vue";
import ChartControls from "./ChartControls.vue";
import GraphChart from "./GraphChart.vue";

import { OpenFileDialog, ReadFile, ParseFile } from "../../wailsjs/go/main/App";
import {parser} from "../../wailsjs/go/models";
import ParseResult = parser.ParseResult;

const isOptimized = ref(false);
const parseData = ref<ParseResult | null>(null);
const loading = ref(false);

const openFile = async () => {
  loading.value = true;
  try {
    const path = await OpenFileDialog();
    if (!path) return;

    const content = await ReadFile(path);
    if (!content) return;

    parseData.value = null;
    parseData.value = await ParseFile(content);

  } finally {
    isOptimized.value = false
    loading.value = false;
  }
};
const togglePlot = () => {
  isOptimized.value = !isOptimized.value;
};
</script>

<template>
  <div class="chart-wrapper">
    <ChartControls
        :optimized="isOptimized"
        @open-file="openFile"
        @toggle-plot="togglePlot"
        :FBefore="parseData?.F_before ?? -1"
        :FAfter="parseData?.F_after ?? -1"
        :loading="loading"
    />

    <GraphChart
        :data="parseData"
        :optimized="isOptimized"
    />
  </div>

</template>

<style scoped>
.chart-wrapper {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
</style>
