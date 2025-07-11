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
      accept=".png,image/png"
      clearable
      label="Bild hochladen (PNG)"
      prepend-icon="mdi-upload"
      :rules="[fileRule]"
      show-size
      @change="onFileChange"
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
    ImageBase64: '', // Hier speichern wir das Base64-Bild
  })

  const loading = ref(false)

  // Validation rule, dass nur PNG akzeptiert wird
  const fileRule = file => {
    if (!file) return true
    return file.type === 'image/png' || 'Nur PNG-Dateien sind erlaubt.'
  }

  const onFileChange = async files => {
    if (!files || files.length === 0) {
      form.ImageBase64 = ''
      return
    }
    const file = files[0]
    if (file.type !== 'image/png') {
      alert('Bitte nur PNG-Dateien hochladen.')
      form.ImageBase64 = ''
      return
    }

    const reader = new FileReader()
    reader.addEventListener('load', () => {
      // Ergebnis ist ein DataURL, z.B. "data:image/png;base64,...."
      const base64String = reader.result
      // Nur den Base64-Teil extrahieren
      form.ImageBase64 = base64String.split(',')[1]
    })
    reader.readAsDataURL(file)
  }

  const submit = async () => {
    if (!form.ImageBase64) {
      alert('Bitte ein PNG-Bild hochladen.')
      return
    }

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
      ImageBase64: form.ImageBase64, // Base64 Bild senden
    }

    loading.value = true
    try {
      await api.post('/pokemon', payload)
      alert('Pokémon erfolgreich erstellt!')
      // Form reset
      for (const key of Object.keys(form)) form[key] = ''
    } catch (error) {
      console.error('Fehler beim Erstellen:', error)
      alert('Fehler beim Erstellen des Pokémon')
    } finally {
      loading.value = false
    }
  }
</script>
