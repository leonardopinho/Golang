package tests

import (
	. "github.com/leonardopinho/GoLang/3.Stress_test/cmd"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestCreateReport(t *testing.T) {
	totalRequests := 2
	concurrency := 1
	filename := "test_relatorio"
	url := "http://google.com"

	Benchmark(url, totalRequests, concurrency, filename)

	files, _ := os.ReadDir(".")
	var reportFile string
	for _, f := range files {
		if strings.HasPrefix(f.Name(), filename+"_") && strings.HasSuffix(f.Name(), ".txt") {
			reportFile = f.Name()
			break
		}
	}

	if reportFile == "" {
		t.Fatal("Arquivo de relatório não foi gerado")
	}

	content, err := os.ReadFile(reportFile)
	if err != nil {
		t.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	text := string(content)

	_ = os.Remove(reportFile)

	assert.Equal(t, strings.Contains(text, "URL: "+url), true)
	assert.Equal(t, strings.Contains(text, "Quantidade total de requests realizados: 2"), true)
	assert.Equal(t, strings.Contains(text, "Quantidade de requests com falha (ou status != 200): 0"), true)
}
