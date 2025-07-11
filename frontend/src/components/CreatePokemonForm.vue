<template>
  <v-form class="d-flex flex-column gap-4" @submit.prevent="submit">
    <v-text-field v-model="form.Name" label="Name" required />

    <v-text-field
      v-model="form.TypesString"
      label="Typen (kommagetrennt, z.B. Fire, Flying)"
    />

    <v-textarea v-model="form.PokedexEntry" label="Pokédex-Eintrag" />

    <!-- Neuer Upload Input -->
    <v-file-input
      v-model="form.ImageFile"
      accept="image/*"
      clearable
      label="Bild hochladen"
      prepend-icon="mdi-upload"
      show-size
    />

    <v-textarea v-model="form.Appearance" label="Aussehen" />

    <v-text-field
      v-model="form.AttacksString"
      label="Attacken (kommagetrennt, z. B. Tackle, Growl)"
    />

    <v-text-field v-model="form.Ability" label="Fähigkeit" />
    <v-text-field
      v-model="form.HeightM"
      label="Größe in Metern"
      step="0.01"
      type="number"
    />
    <v-text-field
      v-model="form.WeightKg"
      label="Gewicht in Kilogramm"
      step="0.01"
      type="number"
    />
    <v-text-field v-model="form.Category" label="Kategorie" />

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
      alert('Bitte ein Bild auswählen.')
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
        alert('Pokémon erfolgreich erstellt!')
        for (const key of Object.keys(form)) {
          form[key] = key === 'ImageFile' ? null : ''
        }
      } catch (error) {
        console.error('Fehler beim Erstellen:', error)
        alert('Fehler beim Erstellen des Pokémon')
      } finally {
        loading.value = false
      }
    })

    reader.readAsDataURL(form.ImageFile)
  }

</script>
