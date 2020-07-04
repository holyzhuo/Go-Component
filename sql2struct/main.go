package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
	"regexp"
)

const (
	SQLPATH = "/sql2struct/sql.txt"

	NewLine = "\n"
)

var finalStruct string
var hasTime = false
var snakeTableName string

func main() {
	fi, err := os.Open(getSQLFilePath())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		transLine(string(a))
	}

	handleTimeImport()
	handleTableName()

	fmt.Println(finalStruct)
}

func transLine(line string) {
	line = strings.ToLower(line)

	if strings.Contains(line, " key ") {
		return
	}

	if strings.Contains(line, "create table") {
		snakeTableName = matchBackQuoteContent(line)
		finalStruct += "type " + camelString(snakeTableName) + " struct {" + NewLine
		return
	}

	if strings.Contains(line, "engine=") {
		finalStruct += "}"
		return
	}

	if strings.Contains(line, "--") || strings.Contains(line, "//") {
		return
	}

	line = strings.TrimSpace(line)
	spliteStrs := strings.Split(line, " ")
	snakeFieldName := matchBackQuoteContent(spliteStrs[0])
	camelFieldName := "  " + camelString(snakeFieldName)
	if strings.Contains(spliteStrs[1], "int") {
		finalStruct += camelFieldName + " int" + " `json:\"" + snakeFieldName + "\"`" + NewLine
	} else if strings.Contains(spliteStrs[1], "char") || strings.Contains(spliteStrs[1], "text") || strings.Contains(spliteStrs[1], "year") {
		finalStruct += camelFieldName + " string" + " `json:\"" + snakeFieldName + "\"`" + NewLine
	} else if strings.Contains(spliteStrs[1], "float") || strings.Contains(spliteStrs[1], "double") || strings.Contains(spliteStrs[1], "decimal") {
		finalStruct += camelFieldName + " float64" + " `json:\"" + snakeFieldName + "\"`" + NewLine
	} else if strings.Contains(spliteStrs[1], "date") || strings.Contains(spliteStrs[1], "time") {
		finalStruct += camelFieldName + " *time.Time" + " `json:\"" + snakeFieldName + "\"`" + NewLine
		hasTime = true
	}
}

func getSQLFilePath() string {
	return getAppPath() + SQLPATH
}

//获取当前目录
func getAppPath() string {
	pwdPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return pwdPath
}

func matchBackQuoteContent(str string) string {
	var rgx = regexp.MustCompile("`(.*?)`")
	rs := rgx.FindStringSubmatch(str)
	return rs[1]
}

func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func handleTimeImport() {
	if hasTime {
		finalStruct = "import \"time\"" + NewLine + NewLine + finalStruct
	}
}

func handleTableName() {
	camelTableName := camelString(snakeTableName)
	finalStruct += NewLine + NewLine + `func (` + camelTableName + `) TableName() string {
		return "` + snakeTableName + `"
}
`
}
