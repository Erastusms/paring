# paring
Paring adalah singkatan dari Pasar Miring. Aplikasi berbasis microservice yg bertujuan untuk menciptakan transaksi jual beli

# 🚀 Sprint 1 – Foundation & Project Setup (Week 1–2)

🎯 **Goal:** Menyiapkan pondasi proyek & tooling agar tim bisa mulai coding dengan rapi, terukur, dan siap untuk sprint berikutnya.

---

## 📅 Timeline Harian (Week 1–2)

### Week 1
- **Day 1**
  - [ ] [Backend] Buat repo utama (mono-repo).
  - [ ] [DevOps] Setup task board manual (README ini).
  - [ ] [QA] Draft test plan awal (high-level).

- **Day 2**
  - [ ] [Backend] Struktur folder per service (auth-service sebagai contoh).
  - [ ] [DevOps] Setup Dockerfile & docker-compose (dummy service).
  - [ ] [QA] Draft checklist health check endpoint.

- **Day 3**
  - [ ] [Backend] Tambah dummy endpoint `/health` & `/ping`.
  - [ ] [DevOps] Setup GitHub Actions (lint, test, build).
  - [ ] [QA] Coba curl endpoint dummy service.

- **Day 4**
  - [ ] [Backend] Setup API Gateway (Kong/Nginx/SCG) → routing ke auth-service.
  - [ ] [DevOps] Konfigurasi docker-compose untuk gateway.
  - [ ] [QA] Test akses service via gateway.

- **Day 5**
  - [ ] [Backend] Implementasi producer-consumer dummy (`order.created`).
  - [ ] [DevOps] Setup Redis/Kafka container.
  - [ ] [QA] Test event publish-subscribe minimal.

### Week 2
- **Day 6**
  - [ ] [Backend] Draft ERD (User, Product, Order dummy).
  - [ ] [DevOps] Setup PostgreSQL + MongoDB + Redis di docker-compose.
  - [ ] [QA] Review schema & rencana query.

- **Day 7**
  - [ ] [Backend] Buat migration/schema dasar (Postgres).
  - [ ] [DevOps] Validasi koneksi DB dengan dummy query.
  - [ ] [QA] Test hasil migrasi (cek tabel terbentuk).

- **Day 8**
  - [ ] [Backend] Integrasi DB ke dummy service (`GET /users` ambil dari Postgres).
  - [ ] [DevOps] Pastikan service terhubung DB via docker network.
  - [ ] [QA] Jalankan API test ke endpoint `/users`.

- **Day 9**
  - [ ] [Backend] Setup Prometheus metrics di dummy service.
  - [ ] [DevOps] Deploy Grafana + Prometheus di docker-compose.
  - [ ] [QA] Validasi dummy metric tampil di dashboard.

- **Day 10**
  - [ ] [Backend] Refactor kode (struktur lebih clean).
  - [ ] [DevOps] Review CI/CD workflow (improve jika perlu).
  - [ ] [QA] Regression test endpoint + gateway + DB connection.

---

## 📋 Task List per Role

### Backend ![progress](https://img.shields.io/badge/progress-0%25-lightgrey)
- [ ] Buat repo mono-repo dengan struktur `/services`.
- [ ] Tambah service awal: `auth-service`.
- [ ] Endpoint dummy: `/health`, `/ping`.
- [ ] API Gateway routing → `auth-service`.
- [ ] Implement producer-consumer event dummy (`order.created`).
- [ ] Draft ERD awal (User, Product, Order).
- [ ] Migration/schema dasar Postgres.
- [ ] Integrasi DB → endpoint `/users`.
- [ ] Setup dummy Prometheus metric.
- [ ] Refactor kode.

### DevOps ![progress](https://img.shields.io/badge/progress-0%25-lightgrey)
- [ ] Setup docker-compose untuk service, DB, gateway, monitoring.
- [ ] Setup Dockerfile tiap service.
- [ ] Setup GitHub Actions (lint, test, build).
- [ ] Setup Kong/Nginx API Gateway di docker-compose.
- [ ] Setup Redis/Kafka container untuk pub-sub.
- [ ] Setup PostgreSQL, MongoDB, Redis di docker-compose.
- [ ] Deploy Prometheus + Grafana.
- [ ] Review CI/CD workflow.

### QA ![progress](https://img.shields.io/badge/progress-0%25-lightgrey)
- [ ] Draft test plan awal.
- [ ] Draft checklist health check endpoint.
- [ ] Test endpoint dummy (`/health`, `/ping`).
- [ ] Test akses via API Gateway.
- [ ] Test publish-subscribe event dummy.
- [ ] Validasi DB schema terbentuk.
- [ ] Test endpoint DB (`/users`).
- [ ] Validasi Prometheus metric via Grafana.
- [ ] Regression test akhir sprint.

---

## ✅ Deliverables Sprint 1
- Mono-repo dengan minimal 1 service (`auth-service`).
- Docker-compose (service + DB + gateway + monitoring).
- API Gateway routing ke dummy service.
- Producer-consumer dummy event (`order.created`).
- Database Postgres schema dasar (User, Product, Order).
- Monitoring dasar (Prometheus + Grafana).
- README Sprint Tracking (dokumen ini).
