<template>
  <div style="height: 400px;">
    <BarChart :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup>
  import { BarElement, CategoryScale, Chart as ChartJS, Legend, LinearScale, Title, Tooltip } from 'chart.js'
  import { computed } from 'vue'
  import { Bar } from 'vue-chartjs'

  ChartJS.register(BarElement, CategoryScale, LinearScale, Title, Tooltip, Legend)

  const props = defineProps({
    stats: {
      type: Object,
      required: true,
      // Example: { HP: 60, Attack: 80, Defense: 70, Speed: 90, SpAttack: 75, SpDefense: 65 }
    },
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
          props.stats.SpDefense || 0,
        ],
        backgroundColor: '#356abc',
        borderRadius: 6,
        maxBarThickness: 30,
      },
    ],
  }))

  const chartOptions = {
    indexAxis: 'y', // horizontal bar chart
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: { display: false },
      title: { display: true, text: 'Pokémon Stats' },
    },
    scales: {
      x: {
        min: 0,
        max: 150,
        ticks: { stepSize: 25 },
        grid: { display: false },
      },
      y: {
        grid: { display: false },
      },
    },
  }

  const BarChart = Bar
</script>
