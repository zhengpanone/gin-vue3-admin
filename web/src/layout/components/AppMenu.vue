<template>
  <el-menu
    active-text-color="#ffd04b"
    background-color="#304156"
    class="el-menu-vertical-demo"
    default-active="2"
    text-color="#fff"
    router
    :collapse="isCollapse"
  >
    <el-menu-item index="/">
      <el-icon><Setting /></el-icon>
      <span>首页</span>
    </el-menu-item>

    <el-sub-menu index="/product">
      <template #title>
        <el-icon><Location /></el-icon>
        <span>商品</span>
      </template>

      <el-menu-item index="/product/product-list">
        商品列表
      </el-menu-item>
      <el-menu-item index="/product/product-classify">
        商品分类
      </el-menu-item>
      <el-menu-item index="/product/product-attr">
        商品规格
      </el-menu-item>
      <el-menu-item index="/product/product-reply">
        商品评论
      </el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>
<script lang="ts" setup>
import { Setting, Location } from '@element-plus/icons-vue'
import { useMainStore } from '@/store/index'
import { storeToRefs } from 'pinia'

const store = useMainStore()

const { isCollapse } = storeToRefs(store)
</script>

<style lang="scss" scoped>
.el-menu{
  border-right: none;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
