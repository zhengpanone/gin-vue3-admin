<template>
  <el-dropdown>
    <span class="el-dropdown-link">
      {{ store.user?.account }}
      <el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>个人中心</el-dropdown-item>
        <el-dropdown-item @click="handleLogout">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>
<script lang="ts" setup>
import { ArrowDown } from '@element-plus/icons-vue'
import { indexStore } from '@/store/index'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouter } from 'vue-router'
import { logout } from '@/api/common'
import {getItem, removeItem} from '@/utils/storage'
import {USER} from "@/utils/constants";
import {IUserInfo} from "@/api/types/common";
const store = indexStore()
const router = useRouter()
const handleLogout = () => {
  // 退出提示
  ElMessageBox.confirm('确认退出吗？', {
    confirmButtonText: '确定',
    buttonSize: 'default',
    cancelButtonText: '取消'
  }).then(async () => {
    // 确认发送退出请求
    const userInfo = getItem<{token:string} & IUserInfo>(USER)
    await logout(userInfo?.token).then(()=>{
      removeItem(USER)
    })
    // 跳转到登录页
    await router.push({
      name: 'login'
    })
    ElMessage({
      type: 'success',
      message: '退出成功'
    })
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '已取消退出'
    })
  })
}
</script>
<style scoped></style>
