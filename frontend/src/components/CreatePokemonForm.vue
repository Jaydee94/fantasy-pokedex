<template>
  <v-form class="d-flex flex-column gap-4" @submit.prevent="submit">
    <v-text-field v-model="form.Name" :label="$t('form.name')" required />

    <v-text-field
      v-model="form.TypesString"
      :label="$t('form.types')"
    />

    <v-textarea v-model="form.PokedexEntry" :label="$t('form.pokedexEntry')" />
    <v-file-input
      v-model="form.ImageFile"
      accept="image/*"
      clearable
      :label="$t('form.uploadPicture')"
      prepend-icon="mdi-upload"
      show-size
    />

    <v-textarea v-model="form.Appearance" :label="$t('form.appearance')" />

    <v-text-field
      v-model="form.AttacksString"
      :label="$t('form.attacks')"
    />

    <v-text-field v-model="form.Ability" :label="$t('form.ability')" />
    <v-text-field
      v-model="form.HeightM"
      :label="$t('form.height')"
      step="0.01"
      type="number"
    />
    <v-text-field
      v-model="form.WeightKg"
      :label="$t('form.weight')"
      step="0.01"
      type="number"
    />

    <v-text-field v-model="form.Category" :label="$t('form.category')" />

    <v-divider class="my-2" />
    <div class="d-flex flex-wrap gap-2">
      <v-text-field v-model="form.HP" label="HP" type="number" min="0" max="255" style="max-width: 120px" />
      <v-text-field v-model="form.Attack" label="Attack" type="number" min="0" max="255" style="max-width: 120px" />
      <v-text-field v-model="form.Defense" label="Defense" type="number" min="0" max="255" style="max-width: 120px" />
      <v-text-field v-model="form.Speed" label="Speed" type="number" min="0" max="255" style="max-width: 120px" />
      <v-text-field v-model="form.SpAttack" label="Sp. Atk" type="number" min="0" max="255" style="max-width: 120px" />
      <v-text-field v-model="form.SpDefense" label="Sp. Def" type="number" min="0" max="255" style="max-width: 120px" />
    </div>

    <v-btn class="mt-2" color="primary" :disabled="loading" type="submit">
      <template v-if="loading">
        <v-progress-circular indeterminate size="20" color="white" class="mr-2" />
      </template>
      {{ $t('form.createBtn') }}
    </v-btn>
    <v-progress-linear
      v-if="loading"
      indeterminate
      color="primary"
      class="mb-2"
    />

    <v-alert
      v-if="alertMsg"
      :type="alertType"
      dense
      class="mt-4"
      border="start"
      transition="fade-transition"
    >
      {{ alertMsg }}
    </v-alert>
  </v-form>
</template>

<script setup>

import { reactive, ref } from 'vue'
import api from '../services/api'
import { nextTick } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const alertMsg = ref('')
const alertType = ref('success')

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
  HP: '',
  Attack: '',
  Defense: '',
  Speed: '',
  SpAttack: '',
  SpDefense: '',
})

const loading = ref(false)

const submit = async () => {
  if (!form.ImageFile) {
    showAlert(t('form.errorImage'), 'error')
    return
  }

  const reader = new FileReader()
  reader.addEventListener('load', async () => {
    const base64 = reader.result.split(',')[1] // Only the base64 part
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
      HP: Number.parseInt(form.HP) || 0,
      Attack: Number.parseInt(form.Attack) || 0,
      Defense: Number.parseInt(form.Defense) || 0,
      Speed: Number.parseInt(form.Speed) || 0,
      SpAttack: Number.parseInt(form.SpAttack) || 0,
      SpDefense: Number.parseInt(form.SpDefense) || 0,
    }

    loading.value = true
    try {
      await api.post('/pokemon', payload)
      showAlert(t('form.success'), 'success')
      for (const key of Object.keys(form)) {
        form[key] = key === 'ImageFile' ? null : ''
      }
    } catch (error) {
      // Log technical error for developers
      console.error(error)
      // Show user-friendly error message
      if (error.response && error.response.data && error.response.data.error) {
        showAlert(t('form.errorCreate', { error: error.response.data.error }), 'error')
      } else {
        showAlert(t('form.errorUnexpected'), 'error')
      }
    } finally {
      loading.value = false
    }
  })

  reader.readAsDataURL(form.ImageFile)
}

function showAlert(msg, type = 'success') {
  alertMsg.value = msg
  alertType.value = type
  setTimeout(() => {
    alertMsg.value = ''
  }, 2000)
}

</script>
