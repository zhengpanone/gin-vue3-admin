<template>
  <page-container>
    <app-card>
      <template #header>
        数据筛选
      </template>
      <el-form
          :inline="true"
          ref="form"
          :model="listParams"
          :disabled="listLoading"
          @submit.prevent="handleQuery"
      >
        <el-form-item label="状态">
          <el-select
              v-model="listParams.status"
              placeholder="请选择"
              clearable
          >
            <el-option
                label="全部"
                value=""
            />
            <el-option
                label="显示"
                :value="1"
            />
            <el-option
                label="不显示"
                :value="0"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="管理员名称">
          <el-input
              v-model="listParams.name"
              clearable
              placeholder="请输入身份昵称"
          />
        </el-form-item>
        <el-form-item>
          <el-button native-type="submit">
            查询
          </el-button>
        </el-form-item>
      </el-form>
    </app-card>
    <app-card>
      <template #header>
        <el-button
            type="primary"
            @click="formVisible = true"
        >
          添加管理员
        </el-button>
      </template>
      <el-table
          :data="list"
          stripe
          style="width: 100%"
          v-loading="listLoading"
      >
        <el-table-column
            prop="id"
            label="ID"
        />
        <el-table-column
            prop="real_name"
            label="姓名"
        />
        <el-table-column
            prop="account"
            label="账号"
        />
        <el-table-column
            prop="roles"
            label="身份"
            min-width="180"
        >
          <template #default="scope">
            <el-space wrap>
              <el-tag
                  class="role-tag"
                  v-for="(item) in scope.row.roles.split(',')"
                  :key="item"
              >
                {{ item }}
              </el-tag>
            </el-space>
          </template>
        </el-table-column>
        <el-table-column
            label="最后一次登录时间"
            prop="last_time"
        />
        <el-table-column
            label="最后一次登录IP"
            prop="last_ip"
        />
        <el-table-column label="状态">
          <template #default="scope">
            <el-switch
                v-model="scope.row.status"
                active-color="#13ce66"
                inactive-color="#ff4949"
                :active-value="1"
                :inactive-value="0"
                :loading="scope.row.statusLoading"
                @change="handleStatusChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column
            label="操作"
            fixed="right"
            min-width="100"
            align="center"
        >
          <template #default="scope">
            <el-button
                type="text"
                @click="handleUpdate(scope.row.id)"
            >
              编辑
            </el-button>
            <el-popconfirm
                title="确认删除吗？"
                @confirm="handleDelete(scope.row.id)"
            >
              <template #reference>
                <el-button type="text">
                  删除
                </el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <app-pagination
          v-model:page="listParams.page"
          v-model:limit="listParams.limit"
          :list-count="listCount"
          :load-list="loadList"
          :disabled="listLoading"
      />
    </app-card>
  </page-container>
</template>
<script lang="ts" setup>
import {reactive, ref,onMounted} from 'vue'
import {getAdmins} from '@/api/admin'
import type {IListParams,Admin} from '@/api/types/admin'

const list = ref<Admin[]>([]) // 列表数据
const listParams = reactive({ // 列表数据查询参数
  page: 1,
  limit: 10,
  name: '',
  roles: '',
  status: '' as IListParams['status']
})

onMounted(()=>{
  loadList()
})

const loadList = async ()=>{
  const data = await getAdmins(listParams)
  list.value = data.data.list
}

const handleQuery = async () => {
  console.log('handleQuery')
}

</script>
<style scoped>
</style>
