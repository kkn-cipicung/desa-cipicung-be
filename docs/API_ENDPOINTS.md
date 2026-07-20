# Dokumentasi Endpoint Cipicung Backend

Dokumen ini mengikuti implementasi route, handler, service, dan model pada source code. Semua path di bawah menggunakan base URL:

```text
Local      : http://localhost:8080/api
Production : https://cipicung.id/api
```

Semua request dan response memakai `Content-Type: application/json`, kecuali `/ping` yang tidak memiliki body. Endpoint bertanda **Protected** membutuhkan salah satu dari:

```http
Authorization: Bearer <access_token>
```

atau cookie `access_token` yang otomatis diberikan saat login.

## Format umum response

Response sukses dengan data:

```json
{
  "code": 200,
  "message": "...",
  "data": {}
}
```

Response sukses tanpa data (create/update/delete/activate):

```json
{
  "code": 200,
  "message": "..."
}
```

Response error:

```json
{
  "code": 400,
  "message": "Invalid request payload",
  "error": "detail error"
}
```

Field `error` hanya dikirim untuk status di bawah 500 dan hanya jika handler menerima detail error. Error 500 tidak membocorkan detail internal. Endpoint protected dapat mengembalikan `401 {"code":401,"message":"Unauthorized"}`.

Untuk semua endpoint list, body boleh tidak dikirim. `limit <= 0` menjadi `10`, maksimum `100`; `index < 0` menjadi `0`. `index` adalah offset data, bukan nomor halaman.

## Bentuk data response

Contoh objek yang dipakai berulang di bagian endpoint:

### Category

```json
{
  "id": 1,
  "name": "Berita",
  "slug": "berita",
  "type": "news",
  "created_at": "2026-07-20 10:00:00"
}
```

### Dashboard

```json
{
  "id": 1,
  "creator": { "id": 1, "name": "Admin Cipicung" },
  "category": { "id": 1, "name": "Banner" },
  "title": "Selamat Datang di Desa Cipicung",
  "description": "Deskripsi dashboard",
  "media_id": 10,
  "is_active": true,
  "created_at": "2026-07-20 10:00:00"
}
```

`media_id` dapat bernilai `null`.

### News

```json
{
  "id": 1,
  "category": { "id": 1, "name": "Berita Desa" },
  "uploader": { "id": 1, "name": "Admin Cipicung" },
  "title": "Kegiatan Desa",
  "description": "Isi berita",
  "media_id": 10,
  "source": "",
  "created_at": "2026-07-20 10:00:00"
}
```

Catatan: model memiliki field `source`, tetapi mapper service saat ini tidak mengisinya sehingga nilai response selalu string kosong. `media_id` dapat `null`.

### Potential

```json
{
  "id": 1,
  "category": { "id": 2, "name": "UMKM" },
  "title": "Kerajinan Bambu",
  "subtitle": "Produk unggulan desa",
  "slug": "kerajinan-bambu",
  "description": "Deskripsi potensi",
  "location": { "id": 3 },
  "owner": { "name": "Budi", "msisdn": "08123456789" },
  "media_id": 11,
  "created_at": "2026-07-20 10:00:00"
}
```

`media_id` dapat `null`; `created_at` menjadi `""` bila database mengembalikan nilai kosong.

### Business

```json
{
  "id": 1,
  "category": { "id": 2, "name": "UMKM" },
  "owner_name": "Budi",
  "business_name": "Warung Cipicung",
  "description": "Usaha makanan lokal",
  "phone": "08123456789",
  "address": "Desa Cipicung",
  "location_id": 3,
  "instagram": "@warungcipicung",
  "facebook": "Warung Cipicung",
  "created_at": "2026-07-20 10:00:00"
}
```

`location_id`, `instagram`, dan `facebook` dapat `null`.

## System

### `GET /ping`

Payload: tidak ada.

Response `200`:

```json
{
  "code": 200,
  "message": "pong"
}
```

## Auth

### `POST /auth/register`

Payload:

```json
{
  "name": "Admin Cipicung",
  "username": "admin.cipicung",
  "password": "secret123"
}
```

Ketentuan aktual: `name` 2–100 karakter; `username` 3–32 karakter dan hanya huruf kecil, angka, `_`, atau `.` setelah dinormalisasi ke lowercase; validasi panjang/kompleksitas password register saat ini dinonaktifkan.

Response `201`:

```json
{
  "code": 201,
  "message": "User registered successfully"
}
```

Error khusus: `400` payload/auth tidak valid, `409` username sudah dipakai.

### `POST /auth/login`

Payload:

```json
{
  "username": "admin.cipicung",
  "password": "secret123"
}
```

Response `200`:

```json
{
  "code": 200,
  "message": "User logged in successfully",
  "data": {
    "access_token": "<jwt-access-token>"
  }
}
```

Login juga memasang cookie access dan refresh token. Error khusus: `401` kredensial salah, `403` user tidak aktif.

## Category

### `POST /category/create` — Protected

Payload (semua wajib):

```json
{ "name": "Berita Desa", "type": "news" }
```

Response `201`:

```json
{ "code": 201, "message": "Category created successfully" }
```

### `POST /category/list`

Payload opsional:

```json
{ "limit": 10, "index": 0, "type": "news" }
```

`type` boleh string kosong untuk semua tipe.

Response `200`:

```json
{
  "code": 200,
  "message": "Categories retrieved successfully",
  "data": [
    { "id": 1, "name": "Berita Desa", "slug": "berita-desa", "type": "news", "created_at": "2026-07-20 10:00:00" }
  ]
}
```

### `POST /category/detail`

Payload: `{ "id": 1 }`

Response `200`: envelope dengan message `Category retrieved successfully` dan `data` berupa satu objek **Category**.

### `POST /category/update` — Protected

Payload (semua wajib secara efektif):

```json
{ "id": 1, "name": "Berita Terbaru", "type": "news" }
```

Response `200`: `{ "code": 200, "message": "Category updated successfully" }`

### `POST /category/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Category deleted successfully" }`

## Dashboard

### `POST /dashboard/create` — Protected

Payload (`category_id`, `title`, `description` wajib):

```json
{
  "category_id": 1,
  "title": "Selamat Datang di Desa Cipicung",
  "description": "Deskripsi dashboard",
  "media_id": "data:image/png;base64,iVBORw0KGgo..."
}
```

`created_by` tidak perlu dikirim karena diambil dari token. `media_id` adalah string gambar base64/data URL atau `null`, bukan ID numerik.

Response `201`: `{ "code": 201, "message": "Dashboard created successfully" }`

### `POST /dashboard/list`

Payload opsional: `{ "limit": 10, "index": 0 }`

Response `200`: envelope dengan message `Dashboards retrieved successfully` dan `data` berupa array objek **Dashboard**.

### `POST /dashboard/detail`

Payload: tidak ada. Endpoint mengambil dashboard terbaru.

Response `200`: envelope dengan message `Dashboard retrieved successfully` dan `data` berupa satu objek **Dashboard**.

### `POST /dashboard/active`

Payload: tidak ada.

Response `200`: envelope dengan message `Active dashboard retrieved successfully` dan `data` berupa satu objek **Dashboard**.

### `POST /dashboard/update` — Protected

Payload (`id`, `category_id`, `title`, `description` wajib; `is_active` opsional/default `false`):

```json
{
  "id": 1,
  "category_id": 1,
  "title": "Judul Baru",
  "description": "Deskripsi baru",
  "media_id": "data:image/png;base64,iVBORw0KGgo...",
  "is_active": true
}
```

Response `200`: `{ "code": 200, "message": "Dashboard updated successfully" }`

### `POST /dashboard/activate` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Dashboard activated successfully" }`

### `POST /dashboard/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Dashboard deleted successfully" }`

## Map

### `POST /map/create` — Protected

Payload (`elevation`, `coordinate`, `hamlet_one`, dan `hamlet_two` wajib):

```json
{ "elevation": "120 mdpl", "coordinate": "-6.5561,107.4421", "hamlet_one": 2500, "hamlet_two": 2500 }
```

`population` tidak diterima dari client. Backend menghitung `population = hamlet_one + hamlet_two`.

Response `201`: `{ "code": 201, "message": "Map created successfully" }`

### `POST /map/detail`

Payload: tidak ada. Mengambil data map terbaru.

Response `200`:

```json
{
  "code": 200,
  "message": "Map retrieved successfully",
  "data": { "elevation": "120 mdpl", "coordinate": "-6.5561,107.4421", "hamlet_one": 2500, "hamlet_two": 2500, "population": 5000 }
}
```

### `POST /map/active`

Payload: tidak ada.

Response `200`: sama seperti `/map/detail`, dengan message `Active map retrieved successfully`.

### `POST /map/update` — Protected

Payload (`id`, `elevation`, `coordinate`, `hamlet_one`, dan `hamlet_two` wajib):

```json
{ "id": 1, "elevation": "125 mdpl", "coordinate": "-6.5561,107.4421", "hamlet_one": 2550, "hamlet_two": 2550 }
```

Response `200`: `{ "code": 200, "message": "Map updated successfully" }`

### `POST /map/activate` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Map activated successfully" }`

### `POST /map/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Map deleted successfully" }`

## Gallery

### `POST /gallery/create` — Protected

Payload (`category_id`, `title`, `description` wajib):

```json
{
  "category_id": 3,
  "title": "Gotong Royong",
  "description": "Dokumentasi kegiatan warga",
  "media_id": "data:image/jpeg;base64,/9j/4AAQSk..."
}
```

`created_by` berasal dari token. `media_id` adalah base64/data URL atau `null`.

Response `201`: `{ "code": 201, "message": "Gallery created successfully" }`

### `POST /gallery/list`

Payload opsional: `{ "limit": 10, "index": 0 }`

Response `200`:

```json
{
  "code": 200,
  "message": "Galleries retrieved successfully",
  "data": [
    { "id": 1, "title": "Gotong Royong", "image": "uploads/gallery/image.jpg" }
  ]
}
```

### `POST /gallery/detail`

Payload: `{ "id": 1 }`

Response `200`:

```json
{
  "code": 200,
  "message": "Gallery retrieved successfully",
  "data": {
    "title": "Gotong Royong",
    "image": "uploads/gallery/image.jpg",
    "description": "Dokumentasi kegiatan warga",
    "category": [{ "id": 3, "name": "Kegiatan" }]
  }
}
```

### `POST /gallery/update` — Protected

Payload (`id`, `category_id`, `title`, `description` wajib):

```json
{
  "id": 1,
  "category_id": 3,
  "title": "Gotong Royong 2026",
  "description": "Dokumentasi terbaru",
  "media_id": "data:image/jpeg;base64,/9j/4AAQSk..."
}
```

Response `200`: `{ "code": 200, "message": "Gallery updated successfully" }`

### `POST /gallery/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Gallery deleted successfully" }`

## Contact

Payload contact yang dipakai create/update:

```json
{
  "name": "Cipicung",
  "province": "Jawa Barat",
  "regency": "Purwakarta",
  "district": "Sukatani",
  "postal_code": "41167",
  "address": "Kantor Kepala Desa Cipicung",
  "phone": "08123456789",
  "email": "pemdes@cipicung.id",
  "website": "https://cipicung.id"
}
```

Field wajib: `name`, `province`, `regency`, `district`, `address`.

### `POST /contact/create` — Protected

Payload: payload contact di atas.

Response `201`: `{ "code": 201, "message": "Contact created successfully" }`

### `POST /contact/detail`

Payload: tidak ada.

Response `200`:

```json
{
  "code": 200,
  "message": "Contact retrieved successfully",
  "data": {
    "office": {
      "name": "Kantor Kepala Desa Cipicung",
      "address": "Kantor Kepala Desa Cipicung",
      "district": "Sukatani",
      "regency": "Purwakarta",
      "province": "Jawa Barat",
      "postal_code": "41167"
    },
    "contact": {
      "email": "pemdes@cipicung.id",
      "phone": "08123456789",
      "website": "https://cipicung.id"
    },
    "social_media": [],
    "service_hour": [
      { "day": "Senin-Kamis", "time": "08.00-15.00" },
      { "day": "Jumat", "time": "08.00-11.30" },
      { "day": "Sabtu-Minggu", "time": "Tutup" }
    ]
  }
}
```

`social_media` saat ini selalu array kosong dan jam layanan berasal dari nilai statis service.

### `POST /contact/update` — Protected

Payload: payload contact di atas ditambah `"id": 1`. Seluruh field wajib create tetap wajib pada update.

Response `200`: `{ "code": 200, "message": "Contact updated successfully" }`

### `POST /contact/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Contact deleted successfully" }`

## News

### `POST /news/create` — Protected

Payload (`category_id`, `title`, `description` wajib):

```json
{
  "category_id": 1,
  "title": "Kegiatan Desa",
  "description": "Isi berita",
  "media_id": "data:image/jpeg;base64,/9j/4AAQSk..."
}
```

`uploaded_by` berasal dari token. `media_id` adalah base64/data URL atau `null`.

Response `201`: `{ "code": 201, "message": "News created successfully" }`

### `POST /news/list`

Payload opsional: `{ "limit": 10, "index": 0 }`

Response `200`: envelope dengan message `News retrieved successfully` dan `data` berupa array objek **News**.

### `POST /news/detail`

Payload: `{ "id": 1 }`

Response `200`: envelope dengan message `News retrieved successfully` dan `data` berupa satu objek **News**.

### `POST /news/header`

Payload: `{ "id": 1 }`

Response `200`:

```json
{
  "code": 200,
  "message": "News header retrieved successfully",
  "data": { "id": 1, "title": "Kegiatan Desa" }
}
```

### `POST /news/update` — Protected

Payload (`id`, `category_id`, `title`, `description` wajib):

```json
{
  "id": 1,
  "category_id": 1,
  "title": "Kegiatan Desa Terbaru",
  "description": "Isi berita terbaru",
  "media_id": "data:image/jpeg;base64,/9j/4AAQSk..."
}
```

Response `200`: `{ "code": 200, "message": "News updated successfully" }`

### `POST /news/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "News deleted successfully" }`

### `POST /news/find-by-date`

Payload (format wajib `YYYY-MM-DD`):

```json
{ "date": "2026-07-20" }
```

Response `200`: envelope dengan message `News retrieved successfully` dan `data` berupa array objek **News**.

## Potential

Payload potential yang dipakai create/update:

```json
{
  "category_id": 2,
  "title": "Kerajinan Bambu",
  "subtitle": "Produk unggulan desa",
  "slug": "kerajinan-bambu",
  "description": "Deskripsi potensi",
  "location_id": 3,
  "location": {
    "id": 3,
    "latitude": -6.5561,
    "longitude": 107.4421,
    "title": "Lokasi Kerajinan",
    "description": "Dusun 1"
  },
  "owner_name": "Budi",
  "owner_msisdn": "08123456789",
  "media_id": "data:image/jpeg;base64,/9j/4AAQSk..."
}
```

Wajib: `category_id > 0`, `title`, `slug`, `description`, `owner_name`. `subtitle`, `location_id`, `location`, `owner_msisdn`, `media_id` opsional. Lokasi bisa dikosongkan, dikirim sebagai `location_id`, atau dikirim sebagai `location.id`. Jika `location` dikirim tanpa `id`, backend membuat data lokasi baru dan latitude harus -90..90 serta longitude -180..180. `media_id` adalah base64/data URL atau `null`; `uploaded_by` berasal dari token.

### `POST /potential/create` — Protected

Payload: payload potential di atas.

Response `201`: `{ "code": 201, "message": "Potential created successfully" }`

### `POST /potential/list`

Payload opsional: `{ "limit": 10, "index": 0 }`

Response `200`: envelope dengan message `Potentials retrieved successfully` dan `data` berupa array objek **Potential**.

### `POST /potential/detail`

Payload: `{ "id": 1 }`

Response `200`: envelope dengan message `Potential retrieved successfully` dan `data` berupa satu objek **Potential**.

### `POST /potential/update` — Protected

Payload: payload potential di atas ditambah `"id": 1`; field wajib create tetap wajib pada update.

Response `200`: `{ "code": 200, "message": "Potential updated successfully" }`

### `POST /potential/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Potential deleted successfully" }`

## Profile

Payload profile yang dipakai create/update:

```json
{
  "name": "Cipicung",
  "province": "Jawa Barat",
  "regency": "Purwakarta",
  "district": "Sukatani",
  "postal_code": "41167",
  "address": "Kantor Kepala Desa Cipicung",
  "phone": "08123456789",
  "email": "pemdes@cipicung.id",
  "website": "https://cipicung.id",
  "latitude": -6.5561,
  "longitude": 107.4421,
  "vision": "Visi desa",
  "mission": ["Misi pertama", "Misi kedua"],
  "history": "Sejarah desa",
  "description": "Deskripsi desa",
  "region": "Wilayah desa",
  "hamlet_one": "Dusun 1",
  "hamlet_two": "Dusun 2",
  "north_border": "Desa Utara",
  "east_border": "Desa Timur",
  "south_border": "Desa Selatan",
  "west_border": "Desa Barat",
  "area": "10 km2",
  "population": "5.000 jiwa",
  "headmen": [
    {
      "name": "Kepala Desa Periode Lama",
      "position": "kepala-desa",
      "phone": "08123456789",
      "email": "kades@cipicung.id",
      "description": "Kepala Desa Cipicung",
      "order_number": 2,
      "is_active": false,
      "start_date": "2018-01-01",
      "finish_date": "2023-12-31"
    },
    {
      "name": "Kepala Desa Aktif",
      "position": "kepala-desa",
      "phone": "08123456789",
      "email": "kades@cipicung.id",
      "description": "Kepala Desa Cipicung",
      "order_number": 1,
      "is_active": true,
      "start_date": "2024-01-01",
      "finish_date": null
    }
  ],
  "officials": [
    {
      "name": "Sekretaris Desa",
      "position": "sekretaris-desa",
      "phone": "08123456780",
      "email": "sekdes@cipicung.id",
      "description": "Sekretaris Desa Cipicung",
      "order_number": 2,
      "is_active": true,
      "start_date": "2024-01-01",
      "finish_date": null
    }
  ],
  "resource_potential": {
    "title": "Potensi Sumber Daya",
    "detail": "Pertanian dan UMKM",
    "description": "Rincian potensi sumber daya Desa Cipicung"
  }
}
```

Wajib: `name`, `province`, `regency`, `district`, `address`. Field lain opsional/default zero value. `headmen` opsional dan menerima riwayat kepala desa. Untuk setiap item, `name` dan `start_date` wajib; `position` harus `kepala-desa` (string kosong otomatis menjadi nilai tersebut); `finish_date` boleh `null` untuk masa jabatan aktif. Tanggal memakai format `YYYY-MM-DD` dan `finish_date` tidak boleh sebelum `start_date`. Rentang tahun antar-item tidak boleh beririsan: jika satu periode berakhir pada 2023, periode lain paling cepat dimulai pada 2024. `finish_date: null` dianggap masih menjabat sehingga tidak boleh memiliki periode lain setelahnya. Payload tunggal `headman` masih diterima untuk kompatibilitas.

### `POST /profile/create` — Protected

Endpoint save tunggal untuk create dan update. Tanpa `id` akan membuat profile baru; dengan `id` akan memperbarui profile tersebut. Seluruh data `villages`, riwayat `headmen`, struktur pemerintahan `officials`, dan `resource_potential` disimpan dalam satu transaksi.

Payload create: payload profile di atas tanpa `id`.

Payload update: payload profile di atas ditambah `"id": 1`.

Response `201`: `{ "code": 201, "message": "Profile created successfully" }`

Response update `200`: `{ "code": 200, "message": "Profile updated successfully" }`

### `POST /profile/detail`

Payload: `{ "id": 1 }`

Response `200`:

```json
{
  "code": 200,
  "message": "Profile retrieved successfully",
  "data": {
    "id": 1,
    "name": "Cipicung",
    "province": "Jawa Barat",
    "regency": "Purwakarta",
    "district": "Sukatani",
    "postal_code": "41167",
    "address": "Kantor Kepala Desa Cipicung",
    "phone": "08123456789",
    "email": "pemdes@cipicung.id",
    "website": "https://cipicung.id",
    "latitude": -6.5561,
    "longitude": 107.4421,
    "vision": "Visi desa",
    "mission": ["Misi pertama", "Misi kedua"],
    "history": "Sejarah desa",
    "description": "Deskripsi desa",
    "region": "Wilayah desa",
    "hamlet_one": "Dusun 1",
    "hamlet_two": "Dusun 2",
    "north_border": "Desa Utara",
    "east_border": "Desa Timur",
    "south_border": "Desa Selatan",
    "west_border": "Desa Barat",
    "area": "10 km2",
    "population": "5.000 jiwa",
    "headman": {
      "id": 1,
      "name": "Bapak Kepala Desa",
      "position": "kepala-desa",
      "phone": "08123456789",
      "email": "kades@cipicung.id",
      "description": "Kepala Desa Cipicung",
      "order_number": 1,
      "is_active": true
    },
    "created_at": "2026-07-20 10:00:00",
    "updated_at": "2026-07-20 10:00:00"
  }
}
```

`headman` dapat `null`.

### `POST /profile/region-boundary`

Payload: tidak ada.

Response `200`:

```json
{
  "code": 200,
  "message": "Profile region boundary retrieved successfully",
  "data": {
    "region": "Wilayah desa",
    "hamlet_one": "Dusun 1",
    "hamlet_two": "Dusun 2",
    "north_border": "Desa Utara",
    "east_border": "Desa Timur",
    "south_border": "Desa Selatan",
    "west_border": "Desa Barat",
    "area": "10 km2",
    "population": "5.000 jiwa"
  }
}
```

### `POST /profile/vision-mission`

Payload: tidak ada.

Response `200`:

```json
{
  "code": 200,
  "message": "Profile vision mission retrieved successfully",
  "data": { "vision": "Visi desa", "mission": ["Misi pertama", "Misi kedua"] }
}
```

### `POST /profile/government-structure`

Payload: tidak ada.

Response `200`:

```json
{
  "code": 200,
  "message": "Government structure retrieved successfully",
  "data": [
    { "name": "Bapak Kepala Desa", "position": "kepala-desa" },
    { "name": "Bapak Sekretaris", "position": "sekretaris-desa" }
  ]
}
```

### `POST /profile/resource-potential`

Payload: tidak ada.

Response `200`:

```json
{
  "code": 200,
  "message": "Resource potential retrieved successfully",
  "data": { "title": "Potensi Sumber Daya", "detail": "Deskripsi potensi sumber daya" }
}
```

### `POST /profile/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Profile deleted successfully" }`

## Business

### `POST /business/create` — Protected

Payload (`category_id`, `owner_name`, `business_name`, `description`, `phone`, `address` wajib):

```json
{
  "category_id": 2,
  "owner_name": "Budi",
  "business_name": "Warung Cipicung",
  "description": "Usaha makanan lokal",
  "phone": "08123456789",
  "address": "Desa Cipicung",
  "location_id": 3,
  "instagram": "@warungcipicung",
  "facebook": "Warung Cipicung"
}
```

`location_id`, `instagram`, `facebook` opsional; `location_id: 0` dinormalisasi menjadi `null`.

Response `201`: `{ "code": 201, "message": "Business created successfully" }`

### `POST /business/list`

Payload opsional:

```json
{ "limit": 10, "index": 0, "type": "umkm" }
```

`type` boleh kosong untuk semua kategori.

Response `200`: envelope dengan message `Businesses retrieved successfully` dan `data` berupa array objek **Business**.

### `POST /business/detail`

Payload: `{ "id": 1 }`

Response `200`: envelope dengan message `Business retrieved successfully` dan `data` berupa satu objek **Business**.

### `POST /business/update` — Protected

Payload: `id` dan `category_id` wajib; field lain opsional dan hanya field yang dikirim yang diubah.

```json
{
  "id": 1,
  "category_id": 2,
  "owner_name": "Budi Santoso",
  "business_name": "Warung Cipicung Baru",
  "description": "Deskripsi baru",
  "phone": "08123456789",
  "address": "Desa Cipicung",
  "location_id": 3,
  "instagram": "@warungcipicung",
  "facebook": "Warung Cipicung"
}
```

Response `200`: `{ "code": 200, "message": "Business updated successfully" }`

### `POST /business/delete` — Protected

Payload: `{ "id": 1 }`

Response `200`: `{ "code": 200, "message": "Business deleted successfully" }`

## Status error yang berlaku

- `400 Bad Request`: JSON tidak valid, field wajib kosong/tidak ada, ID `0`, tanggal/koordinat tidak valid.
- `401 Unauthorized`: token endpoint protected tidak ada, tidak valid, atau kedaluwarsa; login gagal.
- `403 Forbidden`: akun login tidak aktif.
- `404 Not Found`: data detail/update/delete/activate terkait tidak ditemukan (sesuai hasil repository).
- `409 Conflict`: username register sudah digunakan.
- `500 Internal Server Error`: kegagalan database, media, token, atau error internal lainnya.

Nilai contoh di dokumen ini menggambarkan struktur dan tipe data response; nilai aktual mengikuti data database.
