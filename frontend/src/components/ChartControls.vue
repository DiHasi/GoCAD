<script setup lang="ts">
const props = defineProps<{
  optimized: boolean,
  FBefore: number,
  FAfter: number,
  loading: boolean
}>();
defineEmits(["open-file", "toggle-plot"]);

const improved = props.FAfter < props.FBefore;
</script>

<template>
  <div class="controls">
    <div class="bar">
      <button @click="$emit('open-file')" class="toggle-btn" :disabled="loading">
        Загрузить файл
      </button>

      <button @click="$emit('toggle-plot')" class="toggle-btn" :disabled="loading">
        {{ optimized ? "Показать обычное размещение" : "Показать оптимизированное" }}
      </button>
    </div>

    <div class="f-display" v-if="FBefore != -1 && FAfter != -1">
      F: {{ Number(FBefore).toFixed(2) }}
      <span v-if="optimized" :class="{ improved: Number(FAfter) < Number(FBefore) }">
        →
        {{ Number(FAfter).toFixed(2) }}
      </span>
    </div>
  </div>
</template>

<style scoped>
.bar {
  display: flex;
  gap: 10px;
}

.controls {
  padding: 10px;
  display: flex;
  gap: 10px;
  justify-content: space-between;
  background-color: #f5f5f5;
  border-bottom: 1px solid #ccc;
}

.toggle-btn {
  background: #1976d2;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 15px;
  transition: background 0.2s;
}

.toggle-btn:hover {
  background: #1565c0;
}

.toggle-btn:disabled {
  background: #90a4ae; /* серый фон */
  cursor: not-allowed;
  color: #eceff1;
}

.toggle-btn:disabled:hover {
  background: #90a4ae; /* убрать эффект hover */
}

.f-display {
  font-weight: 500;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  font-size: 16px;
  display: flex;
  align-items: center;
}

.f-display .improved {
  color: #2e7d32; /* зеленый цвет при улучшении */
  font-weight: 700;
}
</style>
