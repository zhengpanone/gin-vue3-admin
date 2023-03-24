
<template>
  <el-config-provider :locale="locale">
    <!-- 路由渲染出口 -->
    <router-view />
  </el-config-provider>
</template>
<script setup lang="ts">
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
const locale = zhCn
</script>
<style scoped>

</style>
