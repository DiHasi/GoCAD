import * as echarts from "echarts";
import type { ElementInfo } from "../types";
import {CallbackDataParams} from "echarts/types/src/util/types";
import {parser} from "../../wailsjs/go/models";
import ParseResult = parser.ParseResult;

export function useGraphRenderer() {
    const renderChart = (instance: echarts.ECharts, data: ParseResult, useOptimized: boolean) => {
        const { elem_names, net_names, NetElements } = data;

        const plotKey = useOptimized ? "opt_plot" : "plot";
        const plot = (data as any)[plotKey] as ElementInfo[];
        if (!plot) return;

        const nodes = plot.map(el => ({
            id: elem_names[el.name] ?? `E${el.name}`,
            name: elem_names[el.name] ?? `E${el.name}`,
            x: el.x,
            y: el.y,
            itemStyle: {
                color: useOptimized ? "#ffb74d" : "#90caf9",
                borderColor: useOptimized ? "#f57c00" : "#1e88e5",
                borderWidth: 2
            },
            label: { show: true }
        }));

        const links: any[] = [];
        const pairMap = new Map<string, { nodes: [string, string]; nets: string[] }>();

        for (const [netId, elems] of Object.entries(NetElements)) {
            const arr = elems as ElementInfo[];
            if (arr.length < 2) continue;

            const netName = net_names[Number(netId)] ?? `Net ${netId}`;

            for (let i = 0; i < arr.length; i++) {
                for (let j = i + 1; j < arr.length; j++) {
                    const a = elem_names[arr[i].name] ?? `E${arr[i].name}`;
                    const b = elem_names[arr[j].name] ?? `E${arr[j].name}`;
                    const key = a < b ? `${a}|${b}` : `${b}|${a}`;

                    if (!pairMap.has(key)) {
                        pairMap.set(key, { nodes: [a, b], nets: [netName] });
                    } else {
                        pairMap.get(key)!.nets.push(netName);
                    }
                }
            }
        }

        for (const { nodes: [source, target], nets } of pairMap.values()) {
            for (let k = 0; k < nets.length; k++) {
                const curv = nets.length === 1 ? 0 : -0.2 + (0.4 / (nets.length - 1)) * k;

                links.push({
                    source,
                    target,
                    lineStyle: { width: 2, color: "#43a047", curveness: curv },
                    label: { show: true, formatter: nets[k], fontSize: 12, color: "#2e7d32" }
                });
            }
        }

        instance.setOption({
            animationDurationUpdate: 800,
            animationEasingUpdate: "cubicInOut",
            tooltip: {
                trigger: "item",
                formatter: (params: CallbackDataParams) => {
                    const data: any = params.data;
                    return params.dataType === "edge"
                        ? `${data.source} → ${data.target}`
                        : data.name;
                }
            },
            series: [
                {
                    type: "graph",
                    layout: "none",
                    data: nodes,
                    links,
                    roam: true,
                    scaleLimit: { min: 0.1, max: 10 },
                    lineStyle: { opacity: 0.9 },
                    emphasis: { focus: "adjacency" }
                }
            ]
        });
    };

    return { renderChart };
}
