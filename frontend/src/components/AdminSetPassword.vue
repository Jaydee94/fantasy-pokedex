<template>
  <v-dialog v-model="dialog" max-width="400" persistent>
    <v-card>
      <v-card-title>Admin Passwort setzen</v-card-title>
      <v-card-text>
        <form @submit.prevent="setPassword">
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
          <v-alert v-if="error" dense type="error">{{ error }}</v-alert>
          <v-alert v-if="success" dense type="success">{{ success }}</v-alert>
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
        <v-btn color="primary" :loading="loading" @click="setPassword">Setzen</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script setup>
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  import { ref, watch } from 'vue'
  import api from '../services/api'

  const dialog = ref(true)
  const password = ref('')
  const passwordRepeat = ref('')
  const error = ref('')
  const success = ref('')
  const loading = ref(false)

  async function setPassword () {
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
    loading.value = true
    try {
      await api.post('/admin/set-password', { password: password.value })
      success.value = 'Passwort erfolgreich gesetzt!'
      setTimeout(() => {
        dialog.value = false
      }, 1500)
    } catch (error_) {
      // Log technical error for developers
      console.error(error_)
      // Show user-friendly error message
      error.value = error_.response?.data?.error || 'An unexpected error occurred while setting the password.'
    } finally {
      loading.value = false
    }
  }

</script>
