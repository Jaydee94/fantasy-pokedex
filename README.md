
# Fantasy Pokedex
<p align="center">
  <img src="frontend/src/assets/fantasy_pokedex.png" alt="Fantasy Pokedex Logo" width="300" />
</p>

A modern fullstack project for your own fantasy Pokédex, built with Go (backend), PostgreSQL, and Vue 3 + Vuetify (frontend).

## Setup (Local)

### Requirements
- Docker/Podman & Docker/Podman Compose

### Start
```bash
cd fantasy-pokedex
docker-compose up --build
```

### Admin Password
- On first start, you will be prompted in the frontend to set an admin password.
- Alternatively, you can set a password via ENV: `ADMIN_PASSWORD=yourPassword` in `.env` or in the compose file.

## Development
- Frontend: `cd frontend && npm install && npm run dev`
- Backend: `cd fantasy_pokedex && go run main.go`

---

**Have fun with your own Pokédex!**
