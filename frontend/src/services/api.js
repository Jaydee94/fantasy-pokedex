import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080', // deine Go-API
})

export default api
