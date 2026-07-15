<script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
import './assets/buttons.css'
</script>

<template>
  <div id="app">
    <header>
      <div class="header-wrapper">
        <div class="title-wrapper">
          <img
            alt="Logo"
            class="logo"
            src="/bibble.png"
            width="100"
            height="100"
          />

          <HelloWorld
            class="header-title"
            msg="aaron's morse code converter"
          />
        </div>

        <div class="button-container">
          <button class="copy-button" @click="handleClick">Copy</button>
          <button class="clear-button" @click="handleClear">Clear</button>
          <button class="play-button" @click="handleConvert">Play</button>
        </div>
      </div>
    </header>

    <main>
      <TheWelcome />
    </main>
  </div>
</template>

<style>
#app {
  min-height: 100vh;
  margin: 0;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: var(--text-color);
  text-align: center;

  display: flex;
  flex-direction: column;
}

header {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.header-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2rem;
}

.title-wrapper {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.header-title {
  font-size: 2rem;
}

.logo {
  width: 100px;
  height: 100px;
  object-fit: contain;
}

.button-container {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

main {
  padding: 2rem;
}
</style>