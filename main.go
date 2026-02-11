package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var estadoMap = map[string]int{
	"RS": 0,
	"DF": 1, "GO": 1, "MT": 1, "MS": 1, "TO": 1,
	"AC": 2, "AM": 2, "AP": 2, "PA": 2, "RO": 2, "RR": 2,
	"CE": 3, "MA": 3, "PI": 3,
	"AL": 4, "PB": 4, "PE": 4, "RN": 4,
	"BA": 5, "SE": 5,
	"MG": 6,
	"ES": 7, "RJ": 7,
	"SP": 8,
	"PR": 9, "SC": 9,
}

func calcDigitos(cpf []int) (int, int) {
	sum1 := 0
	for i := 0; i < 9; i++ {
		sum1 += cpf[i] * (10 - i)
	}
	rest := sum1 % 11
	dv1 := 0
	if rest >= 2 {
		dv1 = 11 - rest
	}

	sum2 := 0
	for i := 0; i < 9; i++ {
		sum2 += cpf[i] * (11 - i)
	}
	sum2 += dv1 * 2
	rest2 := sum2 % 11
	dv2 := 0
	if rest2 >= 2 {
		dv2 = 11 - rest2
	}

	return dv1, dv2
}

func gerarCPF(estado string) string {
	cpf := make([]int, 9)

	for i := 0; i < 8; i++ {
		cpf[i] = rand.Intn(10)
	}

	if dig, ok := estadoMap[strings.ToUpper(estado)]; ok {
		cpf[8] = dig
	} else {
		cpf[8] = rand.Intn(10)
	}

	dv1, dv2 := calcDigitos(cpf)
	cpf = append(cpf, dv1, dv2)

	result := ""
	for _, d := range cpf {
		result += strconv.Itoa(d)
	}

	return fmt.Sprintf("%s.%s.%s-%s", result[0:3], result[3:6], result[6:9], result[9:11])
}

func main() {
	rand.Seed(time.Now().UnixNano())

	estado := flag.String("estado", "", "Sigla do estado (SP, RJ, PE...)")
	qtd := flag.Int("quantidade", 0, "Quantidade de CPFs a gerar")
	flag.Parse()


	if *qtd <= 0 {
		cpf := gerarCPF(*estado)
		fmt.Println("CPF gerado:", cpf)
		return
	}

	file, err := os.Create("cpfs.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	for i := 1; i <= *qtd; i++ {
		cpf := gerarCPF(*estado)
		file.WriteString(cpf + "\n")

		if i%10 == 0 {
			fmt.Printf("Progresso: %d/%d CPFs gerados\n", i, *qtd)
		}
	}

	fmt.Println("Concluído! Arquivo salvo como cpfs.txt")
}
