// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    isAdmin: false,
    adminError: null,
  }),
  actions: {
    async login(password) {
      try {
        const res = await import('../services/api').then(m => m.default.post('/admin/login', { password }))
        this.isAdmin = true;
        this.adminError = null;
        return true;
      } catch (e) {
        this.adminError = e.response?.data?.error || 'Wrong password';
        this.isAdmin = false;
        return false;
      }
    },
    logout() {
      this.isAdmin = false;
      this.adminError = null;
    },
  },
})
