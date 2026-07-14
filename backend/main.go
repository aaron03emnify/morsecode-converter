package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var morseMap = map[string]string{
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".",
	"F": "..-.", "G": "--.", "H": "....", "I": "..", "J": ".---",
	"K": "-.-", "L": ".-..", "M": "--", "N": "-.", "O": "---",
	"P": ".--.", "Q": "--.-", "R": ".-.", "S": "...", "T": "-",
	"U": "..-", "V": "...-", "W": ".--", "X": "-..-", "Y": "-.--",
	"Z": "--..", " ": "/",
	"0": "-----", "1": ".----", "2": "..---", "3": "...--",
	"4": "....-", "5": ".....", "6": "-....", "7": "--...",
	"8": "---..", "9": "----.",
}

var reverseMap = reverseMapping(morseMap)

func home(c *gin.Context) {
	input := c.Query("message")
	output := morse(input)

	c.Header("Content-Type", "text/html; charset=utf-8")

	c.String(200, `
<html>
<head>
<style>
	body {
		font-family: Arial, sans-serif;
		background-color: #1e1e1e;
		color: white;
		margin: 0;
	}

	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100vh;
		gap: 15px;
	}

	input[type="text"] {
		padding: 10px;
		width: 350px;
	}

	textarea {
		padding: 10px;
		width: 350px;
		height: 80px;
		resize: none;
	}

	button {
		padding: 10px 15px;
		background-color: #5d04f7;
		color: white;
		border: none;
		cursor: pointer;
		margin: 5px;
	}

	button:hover {
		background-color: #5005ab;
	}
</style>
</head>
<body>

<div class="container">
	<h1>aaron's morse code converter</h1>

	<form method="GET">
		<input id="message" type="text" name="message" value="`+input+`">
		<button type="submit">Convert</button>
	</form>

	<textarea id="output" readonly>`+output+`</textarea>

	<div>
		<button type="button" onclick="copyOutput()">Copy</button>
		<button type="button" onclick="clearFields()">Clear</button>
		<button type="button" onclick="playMorse()">Play</button>
	</div>

	<img
		src="/pics/bibble.png"
		alt="bibble"
		style="width:100px; margin-top:20px;"
	>

</div>

<script>
const dot = new Audio("/sounds/dot.wav");
const dash = new Audio("/sounds/dash.wav");

function sleep(ms) {
	return new Promise(resolve => setTimeout(resolve, ms));
}

async function playMorse() {
	const text = document.getElementById("output").value;

	for (const c of text) {
		if (c === ".") {
			const sound = dot.cloneNode();
			sound.play();
			await sleep(150);
		} else if (c === "-") {
			const sound = dash.cloneNode();
			sound.play();
			await sleep(350);
		} else if (c === " ") {
			await sleep(250);
		} else if (c === "/") {
			await sleep(700);
		}
	}
}

function copyOutput() {
	const output = document.getElementById("output");

	if (output.value === "")
		return;

	navigator.clipboard.writeText(output.value);
}

function clearFields() {
	document.getElementById("message").value = "";
	document.getElementById("output").value = "";
}
</script>

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
	router := gin.Default()

	startTime := time.Now()

	router.GET("/", home)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"uptime": time.Since(startTime).String(),
		})
	})

	router.POST("/morse-it", func(c *gin.Context) {
		var body struct {
			Input string `json:"input"`
		}

		c.BindJSON(&body)

		c.JSON(200, gin.H{
			"input":  body.Input,
			"output": morse(body.Input),
		})
	})

	router.Static("/sounds", "../frontend/sounds")
	router.Static("/pics", "../frontend/pics")

	fmt.Println("Server running at http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
