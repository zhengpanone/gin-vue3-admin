<template>
  <el-dropdown>
      <span class="el-dropdown-link">
      {{ store.user?.account }}
      <el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item >个人中心</el-dropdown-item>
        <el-dropdown-item >退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>
<script lang="ts" setup>
import { ArrowDown } from '@element-plus/icons-vue'
import { indexStore} from '@/store/index'
const store = indexStore()
</script>
<style scoped>
</style>