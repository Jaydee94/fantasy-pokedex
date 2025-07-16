<template>
  <div>
    <RadarChart :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup>
import { Radar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, RadialLinearScale, PointElement, LineElement, Filler } from 'chart.js'
import { computed } from 'vue'

ChartJS.register(Title, Tooltip, Legend, RadialLinearScale, PointElement, LineElement, Filler)

const props = defineProps({
  stats: {
    type: Object,
    required: true,
    // Example: { HP: 60, Attack: 80, Defense: 70, Speed: 90, SpAttack: 75, SpDefense: 65 }
  }
})

const chartData = computed(() => ({
  labels: ['HP', 'Attack', 'Defense', 'Speed', 'Sp. Atk', 'Sp. Def'],
  datasets: [
    {
      label: 'Stats',
      data: [
        props.stats.HP || 0,
        props.stats.Attack || 0,
        props.stats.Defense || 0,
        props.stats.Speed || 0,
        props.stats.SpAttack || 0,
        props.stats.SpDefense || 0
      ],
      backgroundColor: 'rgba(53, 106, 188, 0.2)',
      borderColor: '#356abc',
      pointBackgroundColor: '#356abc',
      fill: true,
    }
  ]
}))

const chartOptions = {
  responsive: true,
  plugins: {
    legend: { display: false },
    title: { display: true, text: 'Pokémon Stats' }
  },
  scales: {
    r: {
      min: 0,
      max: 150,
      ticks: { stepSize: 25 }
    }
  }
}

const RadarChart = Radar
</script>
