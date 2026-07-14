<script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
import './assets/buttons.css'
</script>

<template>
  <header class="header">
    <div class="header-wrapper">
      <HelloWorld class='header-title' msg="aaron's morse code converter" />
      <img alt="Vue logo" class="logo" src="/bibble.png" width="100" height="100" />
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
</template>

<style>

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  height: 100vh;
  color: var(--text-color);
  margin-top: 60px;
}

header {
  line-height: 1.3;
  display: flex;
  height: 100vh;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .title-wrapper {
    display: flex;
    place-items: flex-start;
    align-items: center;
    flex-wrap: wrap;
  }
}
</style>
