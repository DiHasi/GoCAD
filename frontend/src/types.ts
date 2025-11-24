export interface GraphLink {
    source: string
    target: string
    lineStyle?: {
        width?: number
        color?: string
        curveness?: number
    }
    label?: {
        show?: boolean
        formatter?: string
        fontSize?: number
        color?: string
    }
}

export interface ElementInfo {
    name: number
    x: number
    y: number
}

export interface ParseResult {
    elem_names: string[]
    net_names: string[]
    Q: Record<string, Record<string, number>>
    R: Record<string, Record<string, number>>
    plot: number[][]
    D: number[][]
    OptPlot: Record<string, Record<string, number>>
    X: Record<string, ElementInfo[]>
}