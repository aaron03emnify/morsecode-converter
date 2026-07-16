<script setup>
import "./assets/buttons.css";
import { ref } from "vue";

const message = ref("");
const output = ref("");

const notificationText = ref("");
const showNotification = ref(false);

let notificationTimer;

const dot = new Audio("/sounds/dot.wav");
const dash = new Audio("/sounds/dash.wav");

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

function showMessage(text) {
  notificationText.value = text;
  showNotification.value = true;

  clearTimeout(notificationTimer);

  notificationTimer = setTimeout(() => {
    showNotification.value = false;
  }, 1500);
}

async function handleConvert() {
  if (!message.value.trim()) {
    output.value = "";
    showMessage("Nothing to convert");
    return;
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
    });

    if (!response.ok) {
      throw new Error(`HTTP ${response.status}`);
    }

    const data = await response.json();

    output.value = data.output;
    showMessage("Converted!");
  } catch (err) {
    console.error(err);
    output.value = err.message;
    showMessage("Error!");
  }
}

async function handleCopy() {
  if (!output.value) {
    showMessage("Nothing to copy");
    return;
  }

  await navigator.clipboard.writeText(output.value);

  showMessage("Copied!");
}

function handleClear() {
  message.value = "";
  output.value = "";

  showMessage("Cleared!");
}

async function handlePlay() {
  if (!output.value) {
    showMessage("Nothing to play");
    return;
  }

  showMessage("Playing...");

  for (const c of output.value) {
    if (c === ".") {
      dot.cloneNode().play();
      await sleep(150);
    } else if (c === "-") {
      dash.cloneNode().play();
      await sleep(350);
    } else if (c === " ") {
      await sleep(250);
    } else if (c === "/") {
      await sleep(700);
    }
  }
}
</script>

<template>
  <div id="app">
    <Transition name="notification">
      <div v-if="showNotification" class="notification">
        {{ notificationText }}
      </div>
    </Transition>

    <div class="container">
      <img src="/bibble.png" alt="Bibble" class="logo" />

      <h1 class="animated-title">Aaron's Morse Code Converter</h1>

      <input
        v-model="message"
        class="input"
        placeholder="Enter text or Morse code..."
        @keyup.enter="handleConvert"
      />

      <button class="convert-button" @click="handleConvert">Convert</button>

      <textarea
        v-model="output"
        class="output"
        readonly
        placeholder="Output..."
      ></textarea>

      <div class="button-container">
        <button class="copy-button" @click="handleCopy">Copy</button>

        <button class="clear-button" @click="handleClear">Clear</button>

        <button class="play-button" @click="handlePlay">Play</button>
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
  border: 1px solid #444;
  background-color: #2b2b2b;
  color: white;
  font-size: 16px;
  outline: none;
  transition:
    border-color 0.3s ease,
    box-shadow 0.3s ease;
}

.input::placeholder,
.output::placeholder {
  color: #999;
}

.input:focus,
.output:focus {
  border-color: #6a49ee;
  box-shadow: 0 0 8px rgba(106, 73, 238, 0.4);
}

.output {
  height: 120px;
  resize: none;
}

.animated-title {
  background: linear-gradient(
    120deg,
    #929292 0%,
    #929292 40%,
    #ffffff 50%,
    #929292 60%,
    #929292 100%
  );

  background-size: 200% auto;
  color: transparent;
  background-clip: text;
  animation: shine 6s linear infinite;
}

@keyframes shine {
  from {
    background-position: 200% center;
  }

  to {
    background-position: -200% center;
  }
}
</style>
