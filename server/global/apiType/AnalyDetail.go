package apiType

import "github.com/EasyGolang/goTools/mOKX"

type AnalyTickerType struct {
	List        []mOKX.TypeTicker                `json:"List"`        // 列表
	AnalyWhole  []mOKX.TypeWholeTickerAnaly      `json:"AnalyWhole"`  // 大盘分析结果
	AnalySingle map[string][]mOKX.AnalySliceType `json:"AnalySingle"` // 单个币种分析结果
	Unit        string                           `json:"Unit"`
	WholeDir    int                              `json:"WholeDir"`
}
