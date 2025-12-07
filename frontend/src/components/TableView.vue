<script setup lang="ts">
import {parser} from "../../wailsjs/go/models";

const props = defineProps<{
  data: parser.ParseResult | null
}>();
</script>

<template>
  <div v-if="!data" class="empty">
    Нет данных. Загрузите файл.
  </div>

  <div v-else class="tables">

    <!-- ================= Q ================= -->
    <h2>Матрица Q: связи элемент → сеть</h2>
    <table>
      <thead>
      <tr>
        <th>Элемент</th>
        <th v-for="(net, n) in data.net_names" :key="n">{{ net }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(el, e) in data.elem_names" :key="e">
        <td>{{ el }}</td>
        <td v-for="(_, n) in data.net_names" :key="n">
          {{ data.Q[e]?.[n] ?? 0 }}
        </td>
      </tr>
      </tbody>
    </table>

    <h2>Матрица R: связи элемент → элемент</h2>
    <table>
      <thead>
      <tr>
        <th>Элемент</th>
        <th v-for="(el, i) in data.elem_names" :key="i">{{ el }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(el, i) in data.elem_names" :key="i">
        <td>{{ el }}</td>
        <td v-for="(_, j) in data.elem_names" :key="j">
          {{ data.R[i]?.[j] ?? 0 }}
        </td>
      </tr>
      </tbody>
    </table>

    <h2>Матрица D: расстояния</h2>
    <table>
      <thead>
      <tr>
        <th>Элемент</th>
        <th v-for="(el, i) in data.elem_names" :key="i">{{ el }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(el, i) in data.elem_names" :key="i">
        <td>{{ el }}</td>
        <td v-for="(_, j) in data.elem_names" :key="j">
          {{ Number(data.D[i][j]).toFixed(2) }}
        </td>
      </tr>
      </tbody>
    </table>

  </div>
</template>

<style scoped>
.empty {
  padding: 20px;
  font-size: 16px;
}

.tables {
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 40px;
}

table {
  font-weight: 500;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  font-size: 16px;
  border-collapse: collapse;
  min-width: 600px;
}

th, td {
  border: 1px solid #ccc;
  padding: 6px 10px;
  text-align: center;
}

th {
  background: #e3f2fd;
}

h2 {
  font-weight: 500;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  font-size: 16px;
  margin-bottom: 10px;
}
</style>
