package main

import (
    "fmt"

    "github.com/amichelins/goexpert_stress_test/stress"
    "github.com/spf13/cobra"
)

func main() {
    // rootCmd represents the base command when called without any subcommands
    var rootCmd = &cobra.Command{
        Use:   "goexpert_stress_test",
        Short: "",
        Long:  ``,
        // Uncomment the following line if your bare application
        // has an action associated with it:
        Run: func(cmd *cobra.Command, args []string) {

            sUrl, Err := cmd.PersistentFlags().GetString("url")

            if Err != nil {
                fmt.Printf("Erro ao obter valor da Url: %s", Err.Error())
                return
            } else if sUrl == "" {
                fmt.Printf("Falta a URL a ser processada. Verifique!!!")
                return
            }

            nMaxRequest, Err := cmd.PersistentFlags().GetInt64("requests")

            if Err != nil {
                fmt.Printf("Erro ao obter valor das Request: %s", Err.Error())
                return
            }

            nConcurrency, Err := cmd.PersistentFlags().GetInt64("concurrency")

            if Err != nil {
                fmt.Errorf("Erro ao obter valor da Concorrencia: %s", Err)
                return
            }

            //println("URL informada foi: " + sUrl)
            //fmt.Printf(" Request %d com concorrencia de %d", nMaxRequest, nConcurrency)

            Err = stress.StressTest(sUrl, nMaxRequest, nConcurrency)
        },
    }

    rootCmd.PersistentFlags().String("url", "", "Digite a Url que deve ser testada")
    rootCmd.PersistentFlags().Int64("requests", 100, "")
    rootCmd.PersistentFlags().Int64("concurrency", 2, "")

    rootCmd.Execute()
}
