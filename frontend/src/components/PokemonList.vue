<template>
  <div>
    <v-alert
      v-if="alertMsg"
      :type="alertType"
      dense
      class="mb-2"
      border="start"
      transition="fade-transition"
    >
      {{ alertMsg }}
    </v-alert>
    <v-progress-linear
      v-if="loading"
      indeterminate
      color="primary"
      class="mb-4"
    />

    <v-dialog v-model="showDeleteDialog" max-width="400">
      <v-card>
        <v-card-title>Confirm Deletion</v-card-title>
        <v-card-text>
          Do you really want to delete the Pokémon <b>{{ pokemonToDelete?.Name }}</b>?
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="red" @click="confirmDelete">Delete</v-btn>
          <v-btn text @click="showDeleteDialog = false">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-text-field
      v-model="search"
      label="Search Pokémon"
      prepend-inner-icon="mdi-magnify"
      class="mb-4"
      clearable
    />

    <div class="d-flex align-center mb-2" style="gap: 1rem">
      <v-icon color="primary">mdi-counter</v-icon>
      <span>Pokémon found: <b>{{ filteredPokemons.length }}</b></span>
      <span v-if="filteredPokemons.length !== pokemons.length" class="text-grey">(Total: {{ pokemons.length }})</span>
    </div>

    <v-row>
      <v-col
        v-for="pokemon in filteredPokemons"
        :key="pokemon.Name"
        cols="6"
        md="4"
        sm="6"
      >
        <PokemonCard :pokemon="pokemon" @delete="askDelete" @preview-image="showImagePreview" />
      </v-col>
    </v-row>
    <ImageModal v-if="previewImage" :image="previewImage" @close="previewImage = ''" />
  </div>
<!-- ...existing code... -->
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import api from '../services/api'
import PokemonCard from './PokemonCard.vue'
import ImageModal from './ImageModal.vue'

const pokemons = ref([])
const loading = ref(false)

async function fetchPokemons() {
  loading.value = true
  try {
    const res = await api.get('/pokemon')
    pokemons.value = res.data
  } catch (error) {
    console.error(error)
    showAlert('Could not load Pokémon list. Please try again later.', 'error')
  } finally {
    loading.value = false
  }
}

onMounted(fetchPokemons)

const alertMsg = ref('')
const alertType = ref('success')

function showAlert(msg, type = 'success') {
  alertMsg.value = msg
  alertType.value = type
  setTimeout(() => {
    alertMsg.value = ''
  }, 2000)
}

const showDeleteDialog = ref(false)
const pokemonToDelete = ref(null)

function askDelete(pokemon) {
  pokemonToDelete.value = pokemon
  showDeleteDialog.value = true
}

async function confirmDelete() {
  if (!pokemonToDelete.value) return
  loading.value = true
  try {
    await api.delete(`/pokemon/${encodeURIComponent(pokemonToDelete.value.Name)}`)
    pokemons.value = pokemons.value.filter(p => p.Name !== pokemonToDelete.value.Name)
    showAlert('Pokémon deleted successfully', 'success')
  } catch (error) {
    console.error(error)
    if (error.response && error.response.data && error.response.data.error) {
      showAlert('Could not delete Pokémon: ' + error.response.data.error, 'error')
    } else {
      showAlert('An unexpected error occurred while deleting the Pokémon.', 'error')
    }
  } finally {
    showDeleteDialog.value = false
    pokemonToDelete.value = null
    loading.value = false
  }
}

const search = ref('')

const filteredPokemons = computed(() => {
  if (!search.value) return pokemons.value
  const term = search.value.toLowerCase()
  return pokemons.value.filter(p =>
    p.Name.toLowerCase().includes(term) ||
    (p.Types && p.Types.some(t => t.toLowerCase().includes(term))) ||
    (p.Attacks && p.Attacks.some(a => a.toLowerCase().includes(term)))
  )
})

const previewImage = ref('')
function showImagePreview(image) {
  previewImage.value = image
}
</script>
