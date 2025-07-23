<template>
  <v-btn class="ma-2" color="primary" @click="dialog = true">
    Admin Login
  </v-btn>
  <v-dialog v-model="dialog" max-width="400" persistent>
    <v-card>
      <v-card-title>Admin Login</v-card-title>
      <v-card-text>
        <form @submit.prevent="handleLogin">
          <v-text-field
            v-model="password"
            label="Password"
            type="password"
            @keyup.enter="handleLogin"
          />
          <v-alert v-if="error" dense type="error">{{ error }}</v-alert>
          <v-progress-linear
            v-if="loading"
            class="mb-2 mt-2"
            color="primary"
            indeterminate
          />
        </form>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" :loading="loading" @click="handleLogin">Login</v-btn>
        <v-btn text @click="dialog = false">Cancel</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
  import { computed, ref } from 'vue'
  import { useAppStore } from '../stores/app'

  const dialog = ref(false)
  const password = ref('')
  const appStore = useAppStore()
  const loading = ref(false)
  const error = computed(() => appStore.adminError)

  async function handleLogin () {
    loading.value = true
    try {
      const result = await appStore.login(password.value)
      if (result) {
        dialog.value = false
        password.value = ''
      }
    } catch (error) {
      // Log technical error for developers
      console.error(error)
      // Show user-friendly error message
      alert('An unexpected error occurred during login. Please try again.')
    } finally {
      loading.value = false
    }
  }
</script>
