import api from './api'

export async function isAdminPasswordSet() {
  const res = await api.get('/admin/password-set')
  return res.data.set
}
