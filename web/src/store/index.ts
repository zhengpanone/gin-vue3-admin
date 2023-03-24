import { defineStore } from 'pinia'
export interface State{
  name: string,
  isCollapse: boolean,
  aaa:string
}
export const useMainStore = defineStore({
  id: 'main',
  state: () => ({
    name: '超级管理员',
    isCollapse: false

  }),
  getters: {},
  actions: {
    setIsCollapse () {
      this.isCollapse = !this.isCollapse
    }
  }
})
