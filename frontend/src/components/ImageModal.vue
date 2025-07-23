<template>
  <div class="modal-overlay" @click.self="close">
    <div class="modal-content">
      <img alt="Preview" :src="image">
      <button class="close-btn" @click="close">&times;</button>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'ImageModal',
    props: {
      image: {
        type: String,
        required: true,
      },
    },
    emits: ['close'],
    methods: {
      close () {
        this.$emit('close')
      },
    },
  }
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0,0,0,0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
.modal-content {
  position: relative;
  background: #fff;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 16px rgba(0,0,0,0.3);
  max-width: 90vw;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.modal-content img {
  max-width: 80vw;
  max-height: 70vh;
  border-radius: 8px;
}
.close-btn {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  background: transparent;
  border: none;
  font-size: 2rem;
  cursor: pointer;
}
</style>
