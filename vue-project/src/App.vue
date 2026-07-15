<script setup>
import "./assets/buttons.css"
import { ref } from "vue"

const message = ref("")
const output = ref("")

const dot = new Audio("/sounds/dot.wav")
const dash = new Audio("/sounds/dash.wav")

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

async function handleConvert() {
  if (!message.value.trim()) {
    output.value = ""
    return
  }

  try {
    const response = await fetch("http://localhost:8080/morse-it", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        input: message.value,
      }),
    })

    const data = await response.json()
    output.value = data.output
  } catch (err) {
    console.error(err)
    output.value = "Unable to connect to API."
  }
}

function handleCopy() {
  if (!output.value) return
  navigator.clipboard.writeText(output.value)
}

function handleClear() {
  message.value = ""
  output.value = ""
}

async function handlePlay() {
  for (const c of output.value) {
    if (c === ".") {
      dot.cloneNode().play()
      await sleep(150)
    } else if (c === "-") {
      dash.cloneNode().play()
      await sleep(350)
    } else if (c === " ") {
      await sleep(250)
    } else if (c === "/") {
      await sleep(700)
    }
  }
}
</script>

<template>
  <div id="app">
    <div class="container">
      <img
        src="/bibble.png"
        alt="Bibble"
        class="logo"
      />

      <h1>Aaron's Morse Code Converter</h1>

      <input
        v-model="message"
        class="input"
        placeholder="Enter text or Morse code..."
        @keyup.enter="handleConvert"
      />

      <button
        class="convert-button"
        @click="handleConvert"
      >
        Convert
      </button>

      <textarea
        v-model="output"
        class="output"
        readonly
        placeholder="Output..."
      ></textarea>

      <div class="button-container">
        <button
          class="copy-button"
          @click="handleCopy"
        >
          Copy
        </button>

        <button
          class="clear-button"
          @click="handleClear"
        >
          Clear
        </button>

        <button
          class="play-button"
          @click="handlePlay"
        >
          Play
        </button>
      </div>
    </div>
  </div>
</template>

<style>
html,
body,
#app {
  margin: 0;
  width: 100%;
  height: 100%;
}

body {
  background: #1e1e1e;
  color: white;
  font-family: Arial, Helvetica, sans-serif;
}

#app {
  display: flex;
  justify-content: center;
  align-items: center;
}

.container {
  width: 500px;
  max-width: 90%;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

.logo {
  width: 100px;
}

h1 {
  margin: 0;
}

.input,
.output {
  width: 100%;
  box-sizing: border-box;
  padding: 12px;
  border-radius: 8px;
  border: none;
  font-size: 16px;
}

.output {
  height: 120px;
  resize: none;
}

.convert-button,
.copy-button,
.clear-button,
.play-button {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  background: #5d04f7;
  color: white;
  font-size: 15px;
}

.convert-button:hover,
.copy-button:hover,
.clear-button:hover,
.play-button:hover {
  background: #5005ab;
}

.button-container {
  display: flex;
  gap: 12px;
}
</style>