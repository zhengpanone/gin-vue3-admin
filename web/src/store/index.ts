import { defineStore } from 'pinia'
import { IUserInfo } from '@/api/types/common'
import { setItem, getItem } from '@/utils/storage'
import { USER } from '@/utils/constants'

const state = {
  name: null,
  isCollapse: false,
  user: getItem<{token: string} & IUserInfo>(USER)
}

export type State = typeof state

export const indexStore = defineStore({
  id: 'index',
  state: () => { return state },
  getters: {},
  actions: {
    setIsCollapse() {
      this.isCollapse = !this.isCollapse
    },
    setUser(user: {token:string} & IUserInfo | null) {
      this.user = user
      setItem(USER, user)
    }
  }
})
