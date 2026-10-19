import { appConfig } from './appConfig'

const base = `${appConfig.apiBaseUrl}/v1`

export const apiRoutes = {
  arithmetic: {
    add: `${base}/arithmetic/add`,
    subtract: `${base}/arithmetic/subtract`,
    multiply: `${base}/arithmetic/multiply`,
    divide: `${base}/arithmetic/divide`,
  },
  math: {
    power: `${base}/math/power`,
    sqrt: `${base}/math/sqrt`,
    percentage: `${base}/math/percentage`,
  },
  expression: {
    evaluate: `${base}/expressions/evaluate`,
  },
} as const