<template>
  <v-form class="d-flex flex-column gap-4" @submit.prevent="submit">
    <v-text-field v-model="form.Name" label="Name" required />

    <v-text-field
      v-model="form.TypesString"
      label="Types (comma separated, e.g. Fire, Ground)"
    />

    <v-textarea v-model="form.PokedexEntry" label="Pokédex-Entry" />

    <!-- Neuer Upload Input -->
    <v-file-input
      v-model="form.ImageFile"
      accept="image/*"
      clearable
      label="Upload Picture"
      prepend-icon="mdi-upload"
      show-size
    />

    <v-textarea v-model="form.Appearance" label="Appearance" />

    <v-text-field
      v-model="form.AttacksString"
      label="AAttacks (comma separated, e.g. Fire, Ground)"
    />

    <v-text-field v-model="form.Ability" label="Ability" />
    <v-text-field
      v-model="form.HeightM"
      label="Size in Meter"
      step="0.01"
      type="number"
    />
    <v-text-field
      v-model="form.WeightKg"
      label="Weight in Kilograms"
      step="0.01"
      type="number"
    />
    <v-text-field v-model="form.Category" label="Category" />

    <v-btn class="mt-2" color="primary" :disabled="loading" type="submit">
      Pokémon erstellen
    </v-btn>
  </v-form>
</template>

<script setup>
  import { reactive, ref } from 'vue'
  import api from '../services/api'

  const form = reactive({
    Name: '',
    TypesString: '',
    PokedexEntry: '',
    Appearance: '',
    AttacksString: '',
    Ability: '',
    HeightM: null,
    WeightKg: null,
    Category: '',
    ImageFile: null,
  })

  const loading = ref(false)

  const submit = async () => {
    if (!form.ImageFile) {
      alert('Please select an image file.')
      return
    }

    const reader = new FileReader()
    reader.addEventListener('load', async () => {
      const base64 = reader.result.split(',')[1] // Nur der Base64-Teil
      const payload = {
        Name: form.Name,
        Types: form.TypesString.split(',').map(t => t.trim()),
        PokedexEntry: form.PokedexEntry,
        Appearance: form.Appearance,
        Attacks: form.AttacksString.split(',').map(a => a.trim()),
        Ability: form.Ability,
        HeightM: Number.parseFloat(form.HeightM),
        WeightKg: Number.parseFloat(form.WeightKg),
        Category: form.Category,
        ImageData: base64,
      }

      loading.value = true
      try {
        await api.post('/pokemon', payload)
        alert('Successfully created pokemon')
        for (const key of Object.keys(form)) {
          form[key] = key === 'ImageFile' ? null : ''
        }
      } catch (error) {
        console.error('Error creating pokemon', error)
        alert('Error creating pokemon')
      } finally {
        loading.value = false
      }
    })

    reader.readAsDataURL(form.ImageFile)
  }

</script>
