# Dokumentasi API Public (Tanpa Authentication) - Cipicung Backend

Dokumen ini berisi daftar endpoint API yang **tidak membutuhkan autentikasi** (tidak menggunakan `authMiddleware`) atau **khusus digunakan untuk menampilkan data**.

Semua path di bawah menggunakan base URL:
```text
Local      : http://localhost:8080/api
Production : https://cipicung.id/api
```

Semua request dan response menggunakan header `Content-Type: application/json`, kecuali `GET /ping` yang tidak membutuhkan request body.

---

## Format Umum Response

### 1. Response Sukses dengan Data
```json
{
  "code": 200,
  "message": "...",
  "data": {} // Bisa berupa object atau array
}
```

### 2. Response Error
```json
{
  "code": 400,
  "message": "...",
  "error": "detail error jika ada"
}
```

---

## Daftar Endpoint Berdasarkan Fitur

### 1. System

#### `GET /ping`
Digunakan untuk mengecek status keaktifan server.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "pong"
  }
  ```

---

### 2. Auth
Endpoint untuk pendaftaran akun baru dan login untuk mendapatkan token akses.

#### `POST /auth/register`
* **Payload:**
  ```json
  {
    "name": "Nama Lengkap",
    "username": "username_anda",
    "password": "password123"
  }
  ```
* **Response `201`:**
  ```json
  {
    "code": 201,
    "message": "User registered successfully"
  }
  ```

#### `POST /auth/login`
* **Payload:**
  ```json
  {
    "username": "username_anda",
    "password": "password123"
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "User logged in successfully",
    "data": {
      "access_token": "<jwt-access-token>"
    }
  }
  ```
  *Catatan: Selain mendapatkan `access_token` di body response, browser juga akan otomatis dipasangkan cookie `access_token` dan `refresh_token`.*

---

### 3. Category
Mendapatkan kategori yang tersedia di sistem (misal kategori berita, galeri, dsb).

#### `POST /category/list`
* **Payload (Opsional):**
  ```json
  {
    "limit": 10,
    "index": 0,
    "type": "news" // Filter tipe kategori, e.g. "news", "gallery", "business", etc. Boleh dikosongkan.
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Categories retrieved successfully",
    "data": [
      {
        "id": 1,
        "name": "Berita Desa",
        "slug": "berita-desa",
        "type": "news",
        "created_at": "2026-07-20 10:00:00"
      }
    ]
  }
  ```

#### `POST /category/detail`
* **Payload (Wajib):**
  ```json
  {
    "id": 1
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Category retrieved successfully",
    "data": {
      "id": 1,
      "name": "Berita Desa",
      "slug": "berita-desa",
      "type": "news",
      "created_at": "2026-07-20 10:00:00"
    }
  }
  ```

---

### 4. Dashboard
Menampilkan informasi dashboard/banner utama.

#### `POST /dashboard/list`
* **Payload (Opsional):**
  ```json
  {
    "limit": 10,
    "index": 0
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Dashboards retrieved successfully",
    "data": [
      {
        "id": 1,
        "creator": {
          "id": 1,
          "name": "Admin Cipicung"
        },
        "category": {
          "id": 1,
          "name": "Banner"
        },
        "title": "Selamat Datang di Desa Cipicung",
        "description": "Deskripsi dashboard banner",
        "media_id": 10, // Bisa bernilai null
        "is_active": true,
        "created_at": "2026-07-20 10:00:00"
      }
    ]
  }
  ```

#### `POST /dashboard/detail`
Mendapatkan data dashboard terbaru yang tersimpan.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Dashboard retrieved successfully",
    "data": {
      "id": 1,
      "creator": {
        "id": 1,
        "name": "Admin Cipicung"
      },
      "category": {
        "id": 1,
        "name": "Banner"
      },
      "title": "Selamat Datang di Desa Cipicung",
      "description": "Deskripsi dashboard banner",
      "media_id": 10,
      "is_active": true,
      "created_at": "2026-07-20 10:00:00"
    }
  }
  ```

#### `POST /dashboard/active`
Mendapatkan data dashboard yang berstatus aktif saat ini.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Active dashboard retrieved successfully",
    "data": {
      "id": 1,
      "creator": {
        "id": 1,
        "name": "Admin Cipicung"
      },
      "category": {
        "id": 1,
        "name": "Banner"
      },
      "title": "Selamat Datang di Desa Cipicung",
      "description": "Deskripsi dashboard banner",
      "media_id": 10,
      "is_active": true,
      "created_at": "2026-07-20 10:00:00"
    }
  }
  ```

---

### 5. Map
Menampilkan data geografis/demografis peta desa.

#### `POST /map/detail`
Mendapatkan data informasi peta desa terbaru.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Map retrieved successfully",
    "data": {
      "elevation": "120 mdpl",
      "coordinate": "-6.5561,107.4421",
      "hamlet_one": 2500, // Jumlah populasi Dusun I
      "hamlet_two": 2500, // Jumlah populasi Dusun II
      "population": 5000  // Otomatis terhitung dari hamlet_one + hamlet_two
    }
  }
  ```

#### `POST /map/active`
Mendapatkan data informasi peta desa yang berstatus aktif.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Active map retrieved successfully",
    "data": {
      "elevation": "120 mdpl",
      "coordinate": "-6.5561,107.4421",
      "hamlet_one": 2500,
      "hamlet_two": 2500,
      "population": 5000
    }
  }
  ```

---

### 6. Gallery
Menampilkan data dokumentasi foto/media kegiatan desa.

#### `POST /gallery/list`
* **Payload (Opsional):**
  ```json
  {
    "limit": 10,
    "index": 0
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Galleries retrieved successfully",
    "data": [
      {
        "id": 1,
        "title": "Gotong Royong",
        "image": "uploads/gallery/image.jpg"
      }
    ]
  }
  ```

#### `POST /gallery/detail`
* **Payload (Wajib):**
  ```json
  {
    "id": 1
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Gallery retrieved successfully",
    "data": {
      "title": "Gotong Royong",
      "image": "uploads/gallery/image.jpg",
      "description": "Dokumentasi kegiatan warga",
      "category": [
        {
          "id": 3,
          "name": "Kegiatan"
        }
      ]
    }
  }
  ```

---

### 7. Contact
Menampilkan informasi kontak kantor kepala desa, jam pelayanan, dll.

#### `POST /contact/detail`
* **Payload:** Tidak ada.
* **Response `200`:**
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
      "social_media": [], // Dapat berupa array kosong jika belum terisi
      "service_hour": [
        { "day": "Senin-Kamis", "time": "08.00-15.00" },
        { "day": "Jumat", "time": "08.00-11.30" },
        { "day": "Sabtu-Minggu", "time": "Tutup" }
      ]
    }
  }
  ```

---

### 8. News
Menampilkan daftar dan detail berita desa.

#### `POST /news/list`
* **Payload (Opsional):**
  ```json
  {
    "limit": 10,
    "index": 0
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "News retrieved successfully",
    "data": [
      {
        "id": 1,
        "category": {
          "id": 1,
          "name": "Berita Desa"
        },
        "uploader": {
          "id": 1,
          "name": "Admin Cipicung"
        },
        "title": "Kegiatan Desa",
        "description": "Isi berita...",
        "media_id": 10, // Dapat bernilai null
        "source": "",
        "created_at": "2026-07-20 10:00:00"
      }
    ]
  }
  ```

#### `POST /news/detail`
* **Payload (Wajib):**
  ```json
  {
    "id": 1
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "News retrieved successfully",
    "data": {
      "id": 1,
      "category": {
        "id": 1,
        "name": "Berita Desa"
      },
      "uploader": {
        "id": 1,
        "name": "Admin Cipicung"
      },
      "title": "Kegiatan Desa",
      "description": "Isi berita...",
      "media_id": 10,
      "source": "",
      "created_at": "2026-07-20 10:00:00"
    }
  }
  ```

#### `POST /news/header`
Digunakan untuk mengambil judul berita saja berdasarkan ID.

* **Payload (Wajib):**
  ```json
  {
    "id": 1
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "News header retrieved successfully",
    "data": {
      "id": 1,
      "title": "Kegiatan Desa"
    }
  }
  ```

#### `POST /news/find-by-date`
Mengambil data berita berdasarkan tanggal pembuatan tertentu.

* **Payload (Wajib):**
  ```json
  {
    "date": "2026-07-20" // Harus format YYYY-MM-DD
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "News retrieved successfully",
    "data": [
      {
        "id": 1,
        "category": {
          "id": 1,
          "name": "Berita Desa"
        },
        "uploader": {
          "id": 1,
          "name": "Admin Cipicung"
        },
        "title": "Kegiatan Desa",
        "description": "Isi berita...",
        "media_id": 10,
        "source": "",
        "created_at": "2026-07-20 10:00:00"
      }
    ]
  }
  ```

---

### 9. Potential
Menampilkan daftar potensi yang ada di wilayah desa (misal potensi UMKM, pertanian, dsb).

#### `POST /potential/list`
* **Payload (Opsional):**
  ```json
  {
    "limit": 10,
    "index": 0
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Potentials retrieved successfully",
    "data": [
      {
        "id": 1,
        "category": {
          "id": 2,
          "name": "UMKM"
        },
        "title": "Kerajinan Bambu",
        "subtitle": "Produk unggulan desa",
        "slug": "kerajinan-bambu",
        "description": "Deskripsi potensi kerajinan bambu",
        "location": {
          "id": 3
        },
        "owner": {
          "name": "Budi",
          "msisdn": "08123456789"
        },
        "media_id": 11, // Dapat bernilai null
        "created_at": "2026-07-20 10:00:00"
      }
    ]
  }
  ```

#### `POST /potential/detail`
* **Payload (Wajib):**
  ```json
  {
    "id": 1
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Potential retrieved successfully",
    "data": {
      "id": 1,
      "category": {
        "id": 2,
        "name": "UMKM"
      },
      "title": "Kerajinan Bambu",
      "subtitle": "Produk unggulan desa",
      "slug": "kerajinan-bambu",
      "description": "Deskripsi potensi kerajinan bambu",
      "location": {
        "id": 3
      },
      "owner": {
        "name": "Budi",
        "msisdn": "08123456789"
      },
      "media_id": 11,
      "created_at": "2026-07-20 10:00:00"
    }
  }
  ```

---

### 10. Profile
Menampilkan informasi profil desa (visi misi, sejarah, struktur organisasi, dsb).

#### `POST /profile/detail`
* **Payload (Wajib):**
  ```json
  {
    "id": 1
  }
  ```
* **Response `200`:**
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
      "headman": { // Informasi kepala desa aktif, dapat bernilai null
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

#### `POST /profile/region-boundary`
Mendapatkan informasi batas wilayah desa dan kependudukan.

* **Payload:** Tidak ada.
* **Response `200`:**
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

#### `POST /profile/vision-mission`
Mendapatkan informasi visi dan misi desa.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Profile vision mission retrieved successfully",
    "data": {
      "vision": "Visi desa",
      "mission": ["Misi pertama", "Misi kedua"]
    }
  }
  ```

#### `POST /profile/government-structure`
Mendapatkan data struktur organisasi dan daftar perangkat desa.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Government structure retrieved successfully",
    "data": [
      {
        "id": 1,
        "name": "Bapak Kepala Desa",
        "position": "kepala-desa",
        "phone": "08123456789",
        "email": "kades@cipicung.id",
        "description": "Kepala Desa Cipicung",
        "order_number": 1,
        "is_active": true,
        "start_date": "2024-01-01",
        "finish_date": null
      },
      {
        "id": 2,
        "name": "Bapak Sekretaris",
        "position": "sekretaris-desa",
        "phone": "08123456780",
        "email": "sekdes@cipicung.id",
        "description": "Sekretaris Desa Cipicung",
        "order_number": 2,
        "is_active": true,
        "start_date": "2024-01-01",
        "finish_date": null
      }
    ]
  }
  ```

#### `POST /profile/resource-potential`
Mendapatkan informasi potensi sumber daya desa.

* **Payload:** Tidak ada.
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Resource potential retrieved successfully",
    "data": {
      "title": "Potensi Sumber Daya",
      "detail": "Pertanian dan UMKM",
      "description": "Rincian potensi sumber daya Desa Cipicung"
    }
  }
  ```

---

### 11. Business
Menampilkan daftar usaha masyarakat / UMKM.

#### `POST /business/list`
* **Payload (Opsional):**
  ```json
  {
    "limit": 10,
    "index": 0,
    "type": "umkm" // Filter tipe kategori usaha, boleh dikosongkan.
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Businesses retrieved successfully",
    "data": [
      {
        "id": 1,
        "category": {
          "id": 2,
          "name": "UMKM"
        },
        "owner_name": "Budi",
        "business_name": "Warung Cipicung",
        "description": "Usaha makanan lokal",
        "phone": "08123456789",
        "address": "Desa Cipicung",
        "location_id": 3, // Dapat bernilai null
        "instagram": "@warungcipicung", // Dapat bernilai null
        "facebook": "Warung Cipicung", // Dapat bernilai null
        "created_at": "2026-07-20 10:00:00"
      }
    ]
  }
  ```

#### `POST /business/detail`
* **Payload (Wajib):**
  ```json
  {
    "id": 1
  }
  ```
* **Response `200`:**
  ```json
  {
    "code": 200,
    "message": "Business retrieved successfully",
    "data": {
      "id": 1,
      "category": {
        "id": 2,
        "name": "UMKM"
      },
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
  }
  ```
