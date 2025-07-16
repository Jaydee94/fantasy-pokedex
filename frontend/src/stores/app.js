// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    isAdmin: false,
    adminError: null,
  }),
  actions: {
    login(password) {
      // Simple hardcoded password for demo; replace with real auth in production
      if (password === 'admin123') {
        this.isAdmin = true;
        this.adminError = null;
        return true;
      } else {
        this.adminError = 'Wrong password';
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
