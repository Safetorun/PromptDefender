package main

import (
	"encoding/csv"
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
	"strconv"
	"strings"
)

func stringToFloatArray(str string) ([]float64, error) {
	str = strings.Trim(str, "[]")
	strs := strings.Split(str, ",")
	floats := make([]float64, len(strs))
	for i, s := range strs {
		f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return nil, err
		}
		floats[i] = f
	}
	return floats, nil
}

func floatArrayToString(arr []float64) string {
	strs := make([]string, len(arr))
	for i, val := range arr {
		strs[i] = fmt.Sprintf("%v", val)
	}
	return strings.Join(strs, ", ")
}
func main() {
	f := jen.NewFilePathName("../embeddings", "embeddings")

	csvFile, err := os.Open("scanned.csv")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	badwords, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("error reading from file: %v", err)
		panic(err)
	}

	var embeddingValues [][]float64
	var embeddingValuesStrs []string

	for i, badword := range badwords {
		if i == 0 {
			continue
		}
		floatArr, err := stringToFloatArray(badword[4])
		if err != nil {
			fmt.Printf("Error converting string to float array: %v", err)
			panic(err)
		}
		embeddingValues = append(embeddingValues, floatArr)
		embeddingValuesStrs = append(embeddingValuesStrs, fmt.Sprintf("{%s}", floatArrayToString(floatArr)))
	}

	hardcodedArray := fmt.Sprintf("[][]float64{%s}", strings.Join(embeddingValuesStrs, ", "))

	f.Func().
		Id("GetBadwords").
		Params().
		Index().Index().Float64().
		Block(
			jen.Id("badwords").Op(":=").Add(jen.Op(hardcodedArray)),
			jen.Return(jen.Id("badwords")),
		)

	newFile, err := os.Create("../embeddings/badwords.gen.go")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	defer newFile.Close()

	fmt.Printf("%+v", f)

	if err = f.Render(newFile); err != nil {
		fmt.Printf("error rendering file: %v", err)
		panic(err)
	}
}
