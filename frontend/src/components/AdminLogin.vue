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
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" @click="handleLogin">Login</v-btn>
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

const error = computed(() => appStore.adminError)

function handleLogin() {
  if (appStore.login(password.value)) {
    dialog.value = false
    password.value = ''
  }
}
</script>
