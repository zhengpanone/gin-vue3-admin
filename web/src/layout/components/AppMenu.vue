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

    <el-sub-menu index="/system">
      <template #title>
        <el-icon><Location /></el-icon>
        <span>系统管理</span>
      </template>

      <el-menu-item index="/system/user-list">
        用户管理
      </el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>
<script lang="ts" setup>
import { Setting, Location } from '@element-plus/icons-vue'
import { indexStore } from '@/store/index'
import { storeToRefs } from 'pinia'
import {getBaseMenuTree} from '@/api/menu'
import {ref} from "vue";

const menuTreeData = ref([])

const init = async()=>{
  // 获取所有菜单树
  const res = await getBaseMenuTree()
  menuTreeData.value = res.data.menus

}

init()

const store = indexStore()

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
