<template>
  <v-app>
    <AdminSetPassword v-if="showSetPassword" />
    <v-app-bar app color="blue-grey-darken-2" dark height="80">
      <div style="position: absolute; left: 0; right: 0; top: 0; bottom: 0; display: flex; align-items: center; justify-content: center; pointer-events: none;">
        <v-img
          src="@/assets/fantasy-pokedex-logo.png"
          max-height="100"
          style="max-width: 200px; margin: 0 auto; pointer-events: auto;"
        />
      </div>
      <div style="margin-left: auto; display: flex; align-items: center; z-index: 1;">
        <v-btn text to="/">Home</v-btn>
        <v-btn text to="/create">Create</v-btn>
        <v-btn v-if="isAdmin" text @click="logout">Logout (Admin)</v-btn>
        <AdminLogin v-else />
      </div>
    </v-app-bar>

    <v-main>
      <v-container class="py-5" fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import AdminLogin from './components/AdminLogin.vue'
import AdminSetPassword from './components/AdminSetPassword.vue'
import { useAppStore } from './stores/app'
import { computed, ref, onMounted } from 'vue'
import { isAdminPasswordSet } from './services/admin'

const appStore = useAppStore()
const isAdmin = computed(() => appStore.isAdmin)
const showSetPassword = ref(false)

onMounted(async () => {
  try {
    const set = await isAdminPasswordSet()
    showSetPassword.value = !set
  } catch (e) {
    showSetPassword.value = false
  }
})

function logout() {
  appStore.logout()
}
</script>
