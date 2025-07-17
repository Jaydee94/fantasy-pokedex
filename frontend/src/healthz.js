import axios from 'axios';

export async function checkBackendHealth() {
  const backendUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080';
  try {
    const response = await axios.get(`${backendUrl}/healthz`);
    return response.status === 200;
  } catch (err) {
    return false;
  }
}
