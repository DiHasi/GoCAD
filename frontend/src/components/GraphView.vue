<script setup lang="ts">
import {nextTick, ref} from "vue";
import ChartControls from "./ChartControls.vue";
import GraphChart from "./GraphChart.vue";
import TableView from "./TableView.vue";

import { OpenFileDialog, ReadFile, ParseFile } from "../../wailsjs/go/main/App";
import {parser} from "../../wailsjs/go/models";
import ParseResult = parser.ParseResult;

const isOptimized = ref(false);
const parseData = ref<ParseResult | null>(null);
const loading = ref(false);

const activeTab = ref<'graph' | 'tables'>('graph');

const openFile = async () => {
  loading.value = true;
  try {
    const path = await OpenFileDialog();
    if (!path) return;

    const content = await ReadFile(path);
    if (!content) return;

    parseData.value = await ParseFile(content);

  } finally {
    isOptimized.value = false;
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

    <div class="tabs">
      <button :class="{active: activeTab === 'graph'}" @click="activeTab = 'graph'">Граф</button>
      <button :class="{active: activeTab === 'tables'}" @click="activeTab = 'tables'">Матрицы</button>
    </div>

    <div class="tab-content">
      <GraphChart
          v-if="activeTab === 'graph'"
          :data="parseData"
          :optimized="isOptimized"
      />

      <TableView
          v-if="activeTab === 'tables'"
          :data="parseData"
      />
    </div>

  </div>
</template>

<style scoped>
.chart-wrapper {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.tabs {
  display: flex;
  border-bottom: 1px solid #ccc;
  background: #f5f5f5;
}

.tabs button {
  padding: 10px 20px;
  border: none;
  background: transparent;
  cursor: pointer;
  font-size: 15px;
}

.tabs button.active {
  background: white;
  border-bottom: 2px solid #1976d2;
  font-weight: 600;
}

.tab-content {
  flex: 1;
  overflow: auto;
  background-color: #fafafa;
}
</style>
