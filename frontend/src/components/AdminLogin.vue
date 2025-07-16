<template>
  <v-btn color="primary" class="ma-2" @click="dialog = true">
    Admin Login
  </v-btn>
  <v-dialog v-model="dialog" persistent max-width="400">
    <v-card>
      <v-card-title>Admin Login</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="password"
          label="Password"
          type="password"
          @keyup.enter="handleLogin"
        />
        <v-alert v-if="error" type="error" dense>{{ error }}</v-alert>
        <v-progress-linear
          v-if="loading"
          indeterminate
          color="primary"
          class="mb-2 mt-2"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" @click="handleLogin" :loading="loading">Login</v-btn>
        <v-btn text @click="dialog = false">Cancel</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useAppStore } from '../stores/app'

const dialog = ref(false)
const password = ref('')
const appStore = useAppStore()
const loading = ref(false)

const error = computed(() => appStore.adminError)

async function handleLogin() {
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
