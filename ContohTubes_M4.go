package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(input *bufio.Scanner) (text string) {
	input.Scan()
	inputan := input.Text()
	token := strings.ToLower(inputan)
	token = string(token + " #")
	return token

}

type Tokens struct {
	tipe  string
	value string
}

func getTokens(words []string, tokens *[]Tokens) {
	var tipeToken string
	for _, word := range words {
		if word == "for" || word == "break" || word == "if" {
			if word == "for" {
				tipeToken = "statement"
			} else {
				tipeToken = word
			}
		} else if word == "fmt.println('hello,tugas')" {
			tipeToken = "aksi"
		} else if word == ";" {
			tipeToken = ";"
		} else if word == "true" || word == "false" {
			tipeToken = "bool"
		} else if word >= "0" && word <= "9" {
			tipeToken = "number"
		} else if word == "+" || word == "-" || word == "/" || word == "*" || word == "<" || word == ">" || word == "<=" || word == ">=" || word == "==" || word == "!" || word == "=" || word == ":=" {
			tipeToken = "operator"
		} else if word == "{" || word == "}" {
			if word == "{" {
				tipeToken = "{"
			} else {
				tipeToken = "}"
			}
		} else if word == "#" {
			tipeToken = "EOF"
		} else if word == word+"++" {
			tipeToken = "increment"
		} else if word >= "a" && word <= "z" {
			tipeToken = "string"
		}

		token := Tokens{tipe: tipeToken, value: word}
		*tokens = append(*tokens, token)
	}
}

func main() {
	fmt.Println("Input token")
	inputan := bufio.NewScanner(os.Stdin)
	token := getInput(inputan)

	var tokens []Tokens
	words := strings.Split(token, " ")
	getTokens(words, &tokens)

	dictionary := make(map[string][][]string)
	dictionary["statement"] = [][]string{{"for"}}
	dictionary["{"] = [][]string{{"{"}}
	dictionary["loop body"] = [][]string{{"string", "operator", "variable"}, {"aksi"}}
	dictionary["kondisi"] = [][]string{{"string", "operator", "string", "operator", "variable"}}
	dictionary["post kondisi"] = [][]string{{"variable"}, {"inisasi"}, {"bool"}}
	dictionary["inisiasi"] = [][]string{{"string", "operator", "variable"}}
	dictionary["if"] = [][]string{{"if"}}
	dictionary["break"] = [][]string{{"break"}}
	dictionary["variable"] = [][]string{{"string"}, {"variable"}, {"bool"}}
	dictionary["number"] = [][]string{{"number"}}
	dictionary["string"] = [][]string{{"string"}}
	dictionary["operator"] = [][]string{{"operator"}}

	var stack []string
	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i].tipe == "EOF" {
			stack = append(stack, tokens[i].value)
		} else {
			stack = append(stack, tokens[i].tipe)
		}
	}

	fmt.Println(stack)
}
