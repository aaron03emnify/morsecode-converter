package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var morseMap = map[string]string{
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".",
	"F": "..-.", "G": "--.", "H": "....", "I": "..", "J": ".---",
	"K": "-.-", "L": ".-..", "M": "--", "N": "-.", "O": "---",
	"P": ".--.", "Q": "--.-", "R": ".-.", "S": "...", "T": "-",
	"U": "..-", "V": "...-", "W": ".--", "X": "-..-", "Y": "-.--",
	"Z": "--..",
	"0": "-----", "1": ".----", "2": "..---", "3": "...--",
	"4": "....-", "5": ".....", "6": "-....", "7": "--...",
	"8": "---..", "9": "----.",
}

var reverseMap = reverseMapping(morseMap)

func home(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("message")
	output := morse(input)

	fmt.Fprint(w, `
	<html>
	<head>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #1e1e1e;
				color: white;
				margin: 0;
				padding: 20px;
			}
			input[type="text"] {
				padding: 10px;
				width: 300px;
			}
			button {
				padding: 10px;
				background-color: #1b09e2;
				color: white;
				border: none;
				cursor: pointer;
			}
			button:hover {
				background-color: #0f0a9c;
			}
		</style>
	</head>
	<body>
		<div style="display:flex;flex-direction:column;align-items:center;justify-content:center;height:100vh;">
			<h1>Aaron's morse code converter</h1>

			<form method="GET">
				<input type="text" name="message" value="`+input+`">
				<button type="submit">Convert</button>
			</form>

			<p>Your input: `+input+`</p>
			<p>Output: `+output+`</p>
		</div>
	</body>
	</html>
	`)
}

func morse(input string) string {
	if strings.ContainsAny(input, ".-") {
		var output []string

		for _, code := range strings.Split(input, " ") {
			if letter, ok := reverseMap[code]; ok {
				output = append(output, letter)
			}
		}

		return strings.Join(output, "")
	}

	input = strings.ToUpper(input)

	var output []string

	for _, char := range input {
		if code, ok := morseMap[string(char)]; ok {
			output = append(output, code)
		}
	}

	return strings.Join(output, " ")
}

func reverseMapping(m map[string]string) map[string]string {
	reversed := make(map[string]string)

	for key, value := range m {
		reversed[value] = key
	}

	return reversed
}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
