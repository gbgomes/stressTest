/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stresstest",
	Short: "Realiza o teste de stress em uma determinada URL",
	Long: `
Este app realiza o teste de carga em um serviço web.
O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas conforme abaixo:
--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.

Ao fim do processamanto será apresentado um relatório com as seguintes informações:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		numReq, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		stressTest(url, numReq, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "Url a ser testada")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().IntP("requests", "r", 50, "Número de requests a serem realizadas. Assume 50 se não for informado.")
	rootCmd.Flags().IntP("concurrency", "c", 10, "Número de chamadas simultâneas para a URL. Assume 10 se não for informado.")
}

func stressTest(url string, numReq, concurrency int) {
	var m sync.Mutex
	var contador int64

	fmt.Printf("URL: %s, Requests: %d, Concurrency: %d\n\n", url, numReq, concurrency)

	inicio := time.Now()

	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(numReq)

	hash := NewRepository()

	go func() {
		defer close(ch)
		for i := 0; i < concurrency; i++ {
			ch <- url
		}
	}()

	for i := 0; i < numReq; i++ {
		go func(c chan string, m *sync.Mutex) {
			res, err := http.Get(<-c)
			m.Lock()
			if err != nil {
				//println("erro chamando URL. Verifique o formato da URL.")
				hash.Incrementa("999")
			} else {
				//println("URL OK")
				//fmt.Printf("Status %d\n", res.StatusCode)
				hash.Incrementa(fmt.Sprint(res.StatusCode))
			}
			contador++
			m.Unlock()
			wg.Done()
		}(ch, &m)
	}
	wg.Wait()

	fim := time.Now()
	//	fmt.Printf("Inicio: %v:%v:%v:%v", inicio.Hour(), inicio.Minute(), inicio.Second(), inicio.Nanosecond())
	//	println()
	//	fmt.Printf("Fim: %v:%v:%v:%v", fim.Hour(), fim.Minute(), fim.Second(), fim.Nanosecond())
	//	println()
	//	println()
	//
	dif := fim.Sub(inicio)
	fmt.Printf("Tempo total de execução: %s\n\n", dif.String())
	fmt.Printf("Quantidade total de requests realizados: %v\n", contador)

	var detalhamento, detgen string
	var ok, erro int
	for _, item := range hash.DataHM {
		if item.Key == "200" {
			ok = item.Value
		} else if item.Key == "999" {
			erro += item.Value
			detgen += fmt.Sprintf("  - Outros erros, Quantidade: %d\n", item.Value)
		} else {
			erro += item.Value
			detalhamento += fmt.Sprintf("  - Erro %s, Quantidade: %d\n", item.Key, item.Value)
		}
	}
	detalhamento += detgen

	fmt.Printf("Quantidade total de requests com sucesso (http status 200): %d\n", ok)
	fmt.Printf("Quantidade total de requests com Erro: %d\n", erro)
	fmt.Printf("Distribuição dos erros:\n")
	fmt.Printf("%s\n", detalhamento)
}
