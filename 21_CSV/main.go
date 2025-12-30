package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type DataRecord struct {
	ID                                 string
	Name                               string
	Abstraction                        string
	Status                             string
	Description                        string
	Alternate                          string
	Terms                              string
	LikelihoodOfAttack                 string
	TypicalSeverity                    string
	RelatedAttackPatternsExecutionFlow string
	Prerequisites                      string
	SkillsRequiredResourcesRequired    string
	Indicators                         string
	Consequences                       string
	Mitigations                        string
	ExampleInstances                   string
	RelatedWeaknesses                  string
	TaxonomyMappings                   string
	Notes                              string
}

/*
-------------------------------------------------

	CSVファイルを生成する関数

-------------------------------------------------
*/
func writeCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// ヘッダ
	header := []string{
		"ID", "Name", "Abstraction", "Status", "Description",
		"Alternate", "Terms", "LikelihoodOfAttack", "TypicalSeverity",
		"RelatedAttackPatternsExecutionFlow", "Prerequisites",
		"SkillsRequiredResourcesRequired", "Indicators", "Consequences",
		"Mitigations", "ExampleInstances", "RelatedWeaknesses",
		"TaxonomyMappings", "Notes",
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	// サンプルデータ
	record := []string{
		"1000", "SampleName", "Yes", "Active", "Description sample",
		"", "", "High", "Medium",
		"", "", "", "", "",
		"", "", "", "", "",
	}
	if err := writer.Write(record); err != nil {
		return err
	}

	return nil
}

/*
-------------------------------------------------

	CSVファイルを読み込む関数

-------------------------------------------------
*/
func readCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // 列数不一致を許可

	line := 0
	for {
		record, err := reader.Read()
		line++
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("line %d read error: %v\n", line, err)
			continue
		}

		fmt.Printf("line %d\n", line)
		for i, v := range record {
			fmt.Printf("  [%02d] %q\n", i, v)
		}
		fmt.Println("--------------------------------")
	}
	return nil
}

/*
-------------------------------------------------

	main

-------------------------------------------------
*/
func main() {
	filename := "1000.csv"

	// CSV生成
	if err := writeCSV(filename); err != nil {
		fmt.Println("write error:", err)
		return
	}

	// CSV読込
	if err := readCSV(filename); err != nil {
		fmt.Println("read error:", err)
		return
	}
}
