package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("message")
	output := morse(input)

	fmt.Println("your input was: ", input)
	fmt.Println("your morse code: ", output)

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
				h1 {
					color: white;
				}
				form {
					margin-bottom: 20px;
				}
				input[type="text"] {
					padding: 10px;
					width: 300px;
					border: 1px solid #ccc;
					border-radius: 4px;
				}
				button {
					padding: 10px 15px;
					background-color: #4c00ff;
					color: white;
					border: none;
					border-radius: 4px;
					cursor: pointer;
				}
				button:hover {
					background-color: #0056b3;
				}
			</style>
		</head>
        <body>
		<div class="container">
		<div style="display: flex; flex-direction: column; align-items: center; justify-content: center; height: 100vh;">
            <h1>Aaron's morse code converter</h1>
            <form method="GET">
                <input type="text" name="message">
                <button type="submit">Convert</button>
            </form>
			<p>your input was: `+input+`</p>
			<p>your morse code: `+output+`</p>
		</div>
        </body>
        </html>
    `)
}

func morse(input string) string {
	morseCodestrings := map[string]string{
		"A": ".-", "B": "-...",
		"C": "-.-.", "D": "-..", "E": ".",
		"F": "..-.", "G": "--.", "H": "....", "I": "..", "J": ".---",
		"K": "-.-", "L": ".-..", "M": "--", "N": "-.", "O": "---",
		"P": ".--.", "Q": "--.-", "R": ".-.", "S": "...", "T": "-",
		"U": "..-", "V": "...-", "W": ".--", "X": "-..-", "Y": "-.--",
		"Z": "--..",
	}

	morseCodenumber := map[string]string{
		"0": "-----", "1": ".----", "2": "..---", "3": "...--",
		"4": "....-", "5": ".....", "6": "-....", "7": "--...",
		"8": "---..", "9": "----.",
	}

	input = strings.ToUpper(input)

	var morseOutput []string
	for _, char := range input {
		if code, exists := morseCodestrings[string(char)]; exists {
			morseOutput = append(morseOutput, code)
		} else if code, exists := morseCodenumber[string(char)]; exists {
			morseOutput = append(morseOutput, code)
		}
	}

	return strings.Join(morseOutput, " ")
}

func main() {

	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
