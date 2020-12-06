<template>
  <div id="app">
    <el-header class="headnav">
      <el-row>
        <el-col :span="24">
          <div class="home-link">
            <a href="/">TMPFiles</a>
          </div>
        </el-col>
      </el-row>
    </el-header>
      <router-view />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({});
</script>

<style lang="scss">

body {
  margin: 0;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
#app {
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 100vw;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  //text-align: center;
  color: #2c3e50;
}

.headnav {
  display: flex;
  align-items: center;
  z-index: 10000;

  a {
    font-weight: bold;
    color: #2c3e50;
    text-decoration: none;
    font-size: 20px;
  }
}
</style>
