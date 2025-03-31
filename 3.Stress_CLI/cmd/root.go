package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

type Result struct {
	Index        int
	StatusCode   int
	Duration     time.Duration
	ErrorMessage string
}

var (
	url         string
	total       int
	concurrency int
	output      string
	token       bool
)

var rootCmd = &cobra.Command{
	Use:   "bench",
	Short: "Simulador de requisições concorrentes.",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark(url, total, concurrency, output)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "URL para testar (obrigatório)")
	rootCmd.PersistentFlags().IntVarP(&total, "total", "t", 100, "Total de requisições")
	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", 10, "Número de requisições simultâneas")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "relatorio.txt", "Nome do arquivo de saída")
	rootCmd.PersistentFlags().BoolVarP(&token, "token", "k", false, "Tipo do teste de requisição (Token ou IP)")

	rootCmd.MarkPersistentFlagRequired("url")
}

func benchmark(url string, totalRequests, concurrency int, reportFile string) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)
	results := make([]Result, totalRequests)

	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()

			reqStart := time.Now()

			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				results[i] = Result{Index: i, ErrorMessage: err.Error()}
				return
			}

			if token {
				req.Header.Set("API_KEY", "eyJ0eXAiOiJhdCtqd3QiLCJhbGciOiJFUzI1NiIsImtpZCI6IjViOWE0ZmQyY2U3MzliNzQ4M2E1ZjZiMTFhOWVmODk4In0.eyJpc3MiOiJodHRwczovL2lkcC5sb2NhbCIsImF1ZCI6ImFwaTEiLCJzdWIiOiI1YmU4NjM1OTA3M2M0MzRiYWQyZGEzOTMyMjIyZGFiZSIsImNsaWVudF9pZCI6Im15X2NsaWVudF9hcHAiLCJleHAiOjE3NDM0MzA4NDMsImlhdCI6MTc0MzQyNzI0MywianRpIjoiZWMzZjRkYTNkOWY2NjRmODU0MjQxMTY1YzUyOGVmMWMifQ.BSNjukY76NEVbq2OH4LSGaK2pfvrVFqZHdZ3At3fD1MOMCfdwZSfDX1CGFWgIgTi-sXXJ59quHwg0X1UXTsYHA")
			}

			resp, err := http.DefaultClient.Do(req)
			duration := time.Since(reqStart)

			result := Result{
				Index:    i,
				Duration: duration,
			}

			if err != nil {
				result.ErrorMessage = err.Error()
				results[i] = result
				return
			}
			result.StatusCode = resp.StatusCode

			if resp.StatusCode == 429 {
				bodyBytes, _ := io.ReadAll(resp.Body)
				result.ErrorMessage = fmt.Sprintf("HTTP 429: %s", string(bodyBytes))
			}

			resp.Body.Close()

			results[i] = result
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(start)

	generateReport(results, totalTime, totalRequests, concurrency, reportFile)
}

func generateReport(results []Result, totalTime time.Duration, totalRequests int, concurrency int, filename string) {
	var successCount, failCount int
	for _, res := range results {
		if res.ErrorMessage == "" {
			successCount++
		} else {
			failCount++
		}
	}

	timestamp := time.Now().Format("20060102_150405")
	newFileName := fmt.Sprintf("%s_%s.txt", filename, timestamp)

	file, err := os.Create(newFileName)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo de relatório: %v\n", err)
		return
	}
	defer file.Close()

	writer := io.MultiWriter(file, os.Stdout)

	fmt.Fprintf(writer, "\n\n========================================== Relatório \n")
	fmt.Fprintf(writer, "URL: %s\n", url)
	fmt.Fprintf(writer, "Requests: %d\n", totalRequests)
	fmt.Fprintf(writer, "Concurrency: %d\n\n", concurrency)
	fmt.Fprintf(writer, "=> Quantidade total de requests realizados: %d\n", len(results))
	fmt.Fprintf(writer, "=> Quantidade de requests com status HTTP 200: %d\n", successCount)
	fmt.Fprintf(writer, "=> Quantidade de requests com falha (ou status != 200): %d\n", failCount)
	fmt.Fprintf(writer, "=> Tempo total gasto na execução: %.2f ms\n", totalTime.Seconds()*1000)

	fmt.Fprintf(writer, "\n------------------------------------ Erros \n")
	for _, res := range results {
		if res.ErrorMessage != "" {
			fmt.Fprintf(writer, "Requisição %d: %s \n", res.Index+1, res.ErrorMessage)
		}
	}

	fmt.Printf("\nRelatório gerado em '%s'\n", newFileName)
	fmt.Printf("==========================================")
}
