<template>
  <v-dialog v-model="dialog" persistent max-width="400">
    <v-card>
      <v-card-title>Admin Passwort setzen</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="password"
          label="Neues Passwort"
          type="password"
          @keyup.enter="setPassword"
        />
        <v-text-field
          v-model="passwordRepeat"
          label="Passwort wiederholen"
          type="password"
          @keyup.enter="setPassword"
        />
        <v-alert v-if="error" type="error" dense>{{ error }}</v-alert>
        <v-alert v-if="success" type="success" dense>{{ success }}</v-alert>
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn color="primary" @click="setPassword">Setzen</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
import { ref, watch } from 'vue'
import api from '../services/api'

const dialog = ref(true)
const password = ref('')
const passwordRepeat = ref('')
const error = ref('')
const success = ref('')

async function setPassword() {
  error.value = ''
  success.value = ''
  if (!password.value || password.value.length < 6) {
    error.value = 'Mindestens 6 Zeichen!'
    return
  }
  if (password.value !== passwordRepeat.value) {
    error.value = 'Passwörter stimmen nicht überein!'
    return
  }
  try {
    await api.post('/admin/set-password', { password: password.value })
    success.value = 'Passwort erfolgreich gesetzt!'
    setTimeout(() => {
      dialog.value = false
    }, 1500)
  } catch (e) {
    // Log technical error for developers
    console.error(e)
    // Show user-friendly error message
    error.value = e.response?.data?.error || 'An unexpected error occurred while setting the password.'
  }
}

</script>
