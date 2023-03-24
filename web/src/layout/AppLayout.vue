<template>
  <el-container>
    <el-aside>
      <app-menu />
    </el-aside>
    <el-container>
      <el-header>
        <app-header />
      </el-header>
      <el-main>
        <!-- 子路由出口 -->
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>
<script lang="ts" setup>
import AppMenu from './components/AppMenu.vue'
import AppHeader from './AppHeader/index.vue'

</script>
<style scoped>
.el-container {
  height: 100vh;
}

.el-header {
  background-color: #fff;
  color: #333;
  display: flex;
  align-items: center;
  justify-content: space-between;

  box-sizing: border-box;
}

.el-aside {
  background-color: #304156;
  color: #333;
  width: auto;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
}
</style>
