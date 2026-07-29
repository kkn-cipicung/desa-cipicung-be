--
-- PostgreSQL database dump
--

\restrict fTkA4ED6k6tG2m5TfekAiZ0fbxcijU5lGxZS0PwxlW9xSxoPoU8AiqyLaqObgEi

-- Dumped from database version 18.4 (Ubuntu 18.4-1.pgdg22.04+1)
-- Dumped by pg_dump version 18.4 (Ubuntu 18.4-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, role, password, is_active, last_login, created_at, updated_at, username, name) OVERRIDING SYSTEM VALUE VALUES (6, 'admin', '$2a$10$dKemcSEvrqE4Q2k2Y7XKzeP1VqtnTW8amqfUAV.MvmN3Z.QRiEaSO', true, '2026-07-17 21:23:08.349048', '2026-07-10 17:00:38.372401', '2026-07-17 21:23:08.349048', 'roby', 'ilham');
INSERT INTO public.users (id, role, password, is_active, last_login, created_at, updated_at, username, name) OVERRIDING SYSTEM VALUE VALUES (5, 'admin', '$2a$10$vM7LYExPv2CNvn9kMilcC.0ov2UH3ZhXAQ/SPV790bg.aBMZY554a', true, '2026-07-21 13:41:02.923259', '2026-07-04 10:57:40.252264', '2026-07-21 13:41:02.923259', 'ilhamgod14', 'ilham');
INSERT INTO public.users (id, role, password, is_active, last_login, created_at, updated_at, username, name) OVERRIDING SYSTEM VALUE VALUES (12, 'admin', '$2a$10$5nUedfzzlzcTii7YxQOa1O8i6mHetbvwz1lD7KJDli7v90Bt/hvbO', true, '2026-07-26 18:58:39.624646', '2026-07-26 16:13:15.508189', '2026-07-26 18:58:39.624646', 'ramadit', 'Ramadit');
INSERT INTO public.users (id, role, password, is_active, last_login, created_at, updated_at, username, name) OVERRIDING SYSTEM VALUE VALUES (13, 'admin', '$2a$10$CMiRgT34Ud47bvVvOLIlW.9OsaE6h.3McoGlD8hH5yGCkXKJEPKum', true, '2026-07-27 04:45:40.686745', '2026-07-27 04:45:15.075169', '2026-07-27 04:45:40.686745', 'cipicung26', 'Hermawan Sutisna');
INSERT INTO public.users (id, role, password, is_active, last_login, created_at, updated_at, username, name) OVERRIDING SYSTEM VALUE VALUES (11, 'admin', '$2a$10$BifFVEBFCB.4Dm/VoLa9VOqQKExLJ91pZn1gYUrsVOYhRF/PIsX66', true, '2026-07-29 02:32:18.87128', '2026-07-21 14:35:53.137457', '2026-07-29 02:32:18.87128', 'jia', 'Nabiilah Nur Fauziyyah');


--
-- Data for Name: activity_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (1, 'Dashboard', 'dashboard', 'dashboard', '2026-07-04 11:40:35.863008', '2026-07-04 11:40:35.863008');
INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (2, 'Pendidikan', 'pendidikan', 'news', '2026-07-09 00:21:26.60145', '2026-07-09 00:21:26.60145');
INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (3, 'Kuliner', 'kuliner', 'business', '2026-07-10 17:34:32.745891', '2026-07-10 17:34:32.745891');
INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (11, 'Galeri', 'galeri', 'gallery', '2026-07-20 16:32:39.906825', '2026-07-20 16:32:39.906825');
INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (12, 'Hasil Bumi', 'hasil-bumi', 'potential', '2026-07-25 07:34:12.220592', '2026-07-25 07:34:12.220592');
INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (13, 'Kerajinan', 'kerajinan', 'potential', '2026-07-25 07:34:22.154901', '2026-07-25 07:34:22.154901');
INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (14, 'Usaha Desa', 'usaha-desa', 'potential', '2026-07-25 07:34:40.676518', '2026-07-25 07:34:40.676518');
INSERT INTO public.categories (id, name, slug, type, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (15, 'Berita', 'berita', 'news', '2026-07-28 17:57:11.045616', '2026-07-28 17:57:11.045616');


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: businesses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.businesses (id, category_id, owner_name, business_name, description, phone, address, instagram, facebook, created_at, updated_at, location_id) OVERRIDING SYSTEM VALUE VALUES (2, 1, 'ilham', 'aren gula', 'ilham', '0990099990', 'dskasdkdask', NULL, NULL, '2026-07-13 20:41:20.419611', '2026-07-13 20:41:20.419611', NULL);


--
-- Data for Name: media; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (15, 6, 'dashboard', 2, 'image', NULL, NULL, 'uploads/dashboard/1784536348514020947.png', 'image/png', NULL, NULL, 0, '2026-07-20 15:32:28.514496', '2026-07-20 15:32:28.514496');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (20, 6, 'gallery', 7, 'image', NULL, NULL, 'uploads/gallery/1784540002891569326.png', 'image/png', NULL, NULL, 0, '2026-07-20 16:33:22.892896', '2026-07-20 16:33:22.892896');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (21, 6, 'news', 7, 'image', NULL, NULL, 'uploads/news/1784540098169274813.png', 'image/png', NULL, NULL, 0, '2026-07-20 16:34:58.170009', '2026-07-20 16:34:58.170009');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (22, 5, 'news', 8, 'image', NULL, NULL, 'uploads/news/1784641284535406234.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-21 13:41:24.53574', '2026-07-21 13:41:24.53574');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (23, 11, 'news', 10, 'image', NULL, NULL, 'uploads/news/1784707522564218130.png', 'image/png', NULL, NULL, 0, '2026-07-22 08:05:22.565923', '2026-07-22 08:05:22.565923');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (31, 11, 'dashboard', 9, 'image', NULL, NULL, 'uploads/dashboard/1784879995133744841.png', 'image/png', NULL, NULL, 0, '2026-07-24 07:59:55.135114', '2026-07-24 07:59:55.135114');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (32, 11, 'news', 11, 'image', NULL, NULL, 'uploads/news/1784881270451271955.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-24 08:21:10.454108', '2026-07-24 08:21:10.454108');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (34, 11, 'dashboard', 10, 'image', NULL, NULL, 'uploads/dashboard/1784907885094727750.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-24 15:44:45.102318', '2026-07-24 15:44:45.102318');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (36, 11, 'village', 9, 'image', NULL, NULL, 'uploads/dashboard/1784951014718874154.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 03:43:34.719172', '2026-07-25 03:43:34.719172');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (37, 11, 'gallery', 8, 'image', NULL, NULL, 'uploads/gallery/1784951922356789436.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 03:58:42.360174', '2026-07-25 03:58:42.360174');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (38, 11, 'potential', 2, 'image', NULL, NULL, 'uploads/potentials/1784964960867553911.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 07:36:00.868339', '2026-07-25 07:36:00.868339');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (39, 11, 'potential', 1, 'image', NULL, NULL, 'uploads/potentials/1784965149755233036.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 07:39:09.755517', '2026-07-25 07:39:09.755517');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (40, 11, 'potential', 5, 'image', NULL, NULL, 'uploads/potentials/1784965236113994898.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 07:40:36.123517', '2026-07-25 07:40:36.123517');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (41, 11, 'potential', 6, 'image', NULL, NULL, 'uploads/potentials/1784965347536814757.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 07:42:27.537131', '2026-07-25 07:42:27.537131');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (42, 11, 'potential', 7, 'image', NULL, NULL, 'uploads/potentials/1784965484066473648.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 07:44:44.067309', '2026-07-25 07:44:44.067309');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (43, 11, 'potential', 8, 'image', NULL, NULL, 'uploads/potentials/1784965514092808205.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 07:45:14.093048', '2026-07-25 07:45:14.093048');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (44, 11, 'potential', 9, 'image', NULL, NULL, 'uploads/potentials/1784965545817117792.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-25 07:45:45.817327', '2026-07-25 07:45:45.817327');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (48, 5, 'potential', 10, 'image', NULL, NULL, 'uploads/potentials/1784979307403777690.png', 'image/png', NULL, NULL, 0, '2026-07-25 11:35:07.40584', '2026-07-25 11:35:07.40584');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (49, 5, 'gallery', 11, 'image', NULL, NULL, 'uploads/gallery/1784979321083141911.png', 'image/png', NULL, NULL, 0, '2026-07-25 11:35:21.088316', '2026-07-25 11:35:21.088316');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (51, 5, 'news', 12, 'image', NULL, NULL, 'uploads/news/1784979342212410156.png', 'image/png', NULL, NULL, 0, '2026-07-25 11:35:42.214754', '2026-07-25 11:35:42.214754');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (52, 11, 'gallery', 12, 'image', NULL, NULL, 'uploads/gallery/1785131077347336020.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-27 05:44:37.352084', '2026-07-27 05:44:37.352084');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (54, 11, 'gallery', 13, 'image', NULL, NULL, 'uploads/gallery/1785254646144965910.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 16:04:06.155034', '2026-07-28 16:04:06.155034');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (55, 11, 'gallery', 14, 'image', NULL, NULL, 'uploads/gallery/1785254732057109653.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 16:05:32.057771', '2026-07-28 16:05:32.057771');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (56, 11, 'gallery', 16, 'image', NULL, NULL, 'uploads/gallery/1785267218088020745.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 19:33:38.089358', '2026-07-28 19:33:38.089358');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (57, 11, 'gallery', 15, 'image', NULL, NULL, 'uploads/gallery/1785267247721295928.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 19:34:07.722247', '2026-07-28 19:34:07.722247');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (58, 11, 'gallery', 17, 'image', NULL, NULL, 'uploads/gallery/1785267334908715943.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 19:35:34.909576', '2026-07-28 19:35:34.909576');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (59, 11, 'gallery', 18, 'image', NULL, NULL, 'uploads/gallery/1785267445195970659.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 19:37:25.196881', '2026-07-28 19:37:25.196881');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (60, 11, 'gallery', 19, 'image', NULL, NULL, 'uploads/gallery/1785267555353481366.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 19:39:15.361742', '2026-07-28 19:39:15.361742');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (61, 11, 'news', 13, 'image', NULL, NULL, 'uploads/news/1785268623194527847.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 19:57:03.196253', '2026-07-28 19:57:03.196253');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (62, 11, 'news', 14, 'image', NULL, NULL, 'uploads/news/1785268690031812602.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 19:58:10.032708', '2026-07-28 19:58:10.032708');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (63, 11, 'news', 15, 'image', NULL, NULL, 'uploads/news/1785268988016127391.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 20:03:08.016915', '2026-07-28 20:03:08.016915');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (64, 11, 'news', 16, 'image', NULL, NULL, 'uploads/news/1785269107117288062.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-28 20:05:07.118158', '2026-07-28 20:05:07.118158');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (65, 11, 'gallery', 21, 'image', NULL, NULL, 'uploads/gallery/1785292567671521393.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-29 02:36:07.675866', '2026-07-29 02:36:07.675866');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (66, 11, 'gallery', 20, 'image', NULL, NULL, 'uploads/gallery/1785292752510843642.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-29 02:39:12.5197', '2026-07-29 02:39:12.5197');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (67, 11, 'news', 19, 'image', NULL, NULL, 'uploads/news/1785292804665845783.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-29 02:40:04.667045', '2026-07-29 02:40:04.667045');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (68, 11, 'news', 18, 'image', NULL, NULL, 'uploads/news/1785292831036796781.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-29 02:40:31.039413', '2026-07-29 02:40:31.039413');
INSERT INTO public.media (id, uploaded_by, entity_type, entity_id, role, original_name, file_name, file_path, mime_type, file_size, caption, order_number, created_at, updated_at) OVERRIDING SYSTEM VALUE VALUES (69, 11, 'news', 17, 'image', NULL, NULL, 'uploads/news/1785292956551078572.jpg', 'image/jpeg', NULL, NULL, 0, '2026-07-29 02:42:36.552501', '2026-07-29 02:42:36.552501');


--
-- Data for Name: documents; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.documents (id, category_id, uploaded_by, title, description, created_at, media_id, updated_at, source) OVERRIDING SYSTEM VALUE VALUES (13, 15, 11, 'Dorong UMKM Go Digital, KKN UNSIKA 2026 Hadirkan Program CIKOBER di Desa Cipicung', 'Purwakarta, 11 Juli 2026 – Mahasiswa KKN UNSIKA 2026 Kelompok 23 melaksanakan program CIKOBER (Cipicung Ekonomi Berkembang) sebagai upaya mendukung transformasi digital UMKM di Desa Cipicung, Kecamatan Sukatani, Kabupaten Purwakarta. Kegiatan ini dilaksanakan dengan mengunjungi langsung warung dan pelaku UMKM yang ada di Desa Cipicung untuk memberikan pendampingan secara personal.

Melalui pendekatan door to door, mahasiswa KKN berdialog dengan para pelaku usaha mengenai kondisi usaha yang dijalankan, kendala yang dihadapi, serta peluang memanfaatkan media digital sebagai sarana promosi dan pemasaran. Pendampingan dilakukan secara langsung agar materi yang diberikan dapat disesuaikan dengan kebutuhan masing-masing pelaku UMKM.

Dalam program ini, mahasiswa juga membantu memperkenalkan pentingnya membangun identitas usaha, memanfaatkan media sosial sebagai media promosi, serta mendorong pelaku UMKM untuk mulai memasarkan produknya secara lebih luas melalui platform digital. Pendekatan tersebut diharapkan mampu meningkatkan daya saing UMKM lokal di tengah perkembangan teknologi.

Program CIKOBER menjadi salah satu bentuk pengabdian KKN UNSIKA 2026 dalam mendukung pemberdayaan ekonomi masyarakat Desa Cipicung. Dengan adanya pendampingan yang dilakukan secara langsung, diharapkan para pelaku UMKM semakin percaya diri untuk mengembangkan usahanya dan mampu memanfaatkan teknologi digital sebagai sarana meningkatkan penjualan serta memperluas jangkauan pasar.', '2026-07-28 19:57:03.196253', 61, '2026-07-28 19:57:03.196253', NULL);
INSERT INTO public.documents (id, category_id, uploaded_by, title, description, created_at, media_id, updated_at, source) OVERRIDING SYSTEM VALUE VALUES (14, 15, 11, 'Tingkatkan Ketahanan Pangan, KKN UNSIKA 2026 Bersama BPP Sukatani Gelar Program TAMSUR di Desa Cipicung', 'Purwakarta, 15 Juli 2026 – Mahasiswa KKN UNSIKA 2026 Kelompok 23 bekerja sama dengan Balai Penyuluhan Pertanian (BPP) Sukatani menyelenggarakan program TAMSUR (Tanam Sayur) di Desa Cipicung, Kecamatan Sukatani, Kabupaten Purwakarta. Kegiatan ini bertujuan meningkatkan pengetahuan dan keterampilan masyarakat dalam membudidayakan tanaman sayur sebagai upaya mendukung ketahanan pangan keluarga.

Kegiatan diawali dengan penyampaian materi oleh penyuluh dari BPP Sukatani mengenai teknik budidaya tanaman sayur yang baik, mulai dari pemilihan benih, pengolahan media tanam, proses penanaman, hingga perawatan tanaman agar menghasilkan panen yang optimal. Materi tersebut diharapkan dapat menjadi bekal bagi masyarakat untuk menerapkan budidaya sayuran secara mandiri di lingkungan rumah.

Setelah sesi penyuluhan, peserta mengikuti praktik penanaman secara langsung dengan didampingi oleh tim KKN UNSIKA 2026 dan penyuluh BPP Sukatani. Dalam kegiatan ini, BPP Sukatani turut memberikan bantuan benih tanaman sayur yang digunakan sebagai sarana praktik sekaligus untuk mendorong masyarakat mulai memanfaatkan lahan pekarangan sebagai sumber pangan keluarga.

Melalui kolaborasi antara KKN UNSIKA 2026 dan BPP Sukatani, program TAMSUR diharapkan mampu meningkatkan kesadaran masyarakat akan pentingnya pemanfaatan lahan pekarangan untuk budidaya sayuran. Selain mendukung ketahanan pangan rumah tangga, kegiatan ini juga diharapkan dapat menumbuhkan kebiasaan bercocok tanam yang berkelanjutan serta memberikan manfaat ekonomi bagi masyarakat Desa Cipicung.', '2026-07-28 19:58:10.032708', 62, '2026-07-28 19:58:10.032708', NULL);
INSERT INTO public.documents (id, category_id, uploaded_by, title, description, created_at, media_id, updated_at, source) OVERRIDING SYSTEM VALUE VALUES (15, 15, 11, 'Manfaatkan Limbah Sekam Padi, KKN UNSIKA 2026 Gelar Sosialisasi dan Praktik Pembuatan BRIKUNG di Desa Cipicung', 'Purwakarta, 16 Juli 2026 – Mahasiswa KKN UNSIKA 2026 Kelompok 23 menyelenggarakan program BRIKUNG (Briket Ramah Lingkungan) di Aula Kantor Desa Cipicung, Kecamatan Sukatani, Kabupaten Purwakarta. Kegiatan ini diikuti oleh warga Desa Cipicung sebagai upaya mengedukasi masyarakat mengenai pemanfaatan limbah sekam padi menjadi briket yang ramah lingkungan dan memiliki nilai guna.

Kegiatan diawali dengan penyampaian materi mengenai potensi limbah sekam padi sebagai bahan baku pembuatan briket serta manfaatnya sebagai sumber energi alternatif. Selain mengurangi limbah pertanian, pemanfaatan sekam padi menjadi briket juga diharapkan dapat memberikan nilai tambah bagi masyarakat melalui pengolahan limbah yang lebih produktif.

Setelah sesi sosialisasi, peserta mengikuti praktik pembuatan BRIKUNG secara langsung bersama mahasiswa KKN UNSIKA 2026. Warga diperkenalkan pada setiap tahapan pembuatan briket, mulai dari proses pengolahan bahan, pencampuran, pencetakan, hingga proses pengeringan. Antusiasme peserta terlihat dari keterlibatan aktif dalam setiap proses praktik dan diskusi yang berlangsung.

Melalui program BRIKUNG, KKN UNSIKA 2026 berharap masyarakat Desa Cipicung dapat memanfaatkan limbah sekam padi yang selama ini belum dimanfaatkan secara optimal menjadi produk yang lebih bernilai. Program ini menjadi salah satu bentuk pengabdian kepada masyarakat dalam mendorong inovasi berbasis potensi lokal sekaligus meningkatkan kesadaran akan pentingnya menjaga kelestarian lingkungan melalui pengelolaan limbah yang berkelanjutan.', '2026-07-28 20:03:08.016915', 63, '2026-07-28 20:03:08.016915', NULL);
INSERT INTO public.documents (id, category_id, uploaded_by, title, description, created_at, media_id, updated_at, source) OVERRIDING SYSTEM VALUE VALUES (17, 15, 11, 'Tingkatkan Kesadaran Hidup Sehat, KKN UNSIKA 2026 Bersama Puskesmas Sukatani Gelar Cek Kesehatan Gratis', 'Purwakarta – Dalam upaya meningkatkan kesadaran masyarakat akan pentingnya menjaga kesehatan, KKN UNSIKA 2026 Kelompok 23 bekerja sama dengan Puskesmas Sukatani menyelenggarakan kegiatan Cek Kesehatan Gratis bagi warga Desa Cipicung.

Kegiatan ini meliputi pemeriksaan kesehatan dasar seperti pengecekan tekanan darah, gula darah, asam urat, kolesterol, serta konsultasi kesehatan yang didampingi oleh tenaga kesehatan dari Puskesmas Sukatani. Antusiasme masyarakat terlihat dari banyaknya warga yang memanfaatkan kesempatan untuk mengetahui kondisi kesehatannya.

Melalui kegiatan ini, masyarakat diharapkan semakin peduli terhadap kesehatan diri dan keluarga serta terdorong untuk menerapkan pola hidup sehat sebagai langkah pencegahan berbagai penyakit.', '2026-07-29 02:18:37.580107', 69, '2026-07-29 02:42:36.552501', NULL);
INSERT INTO public.documents (id, category_id, uploaded_by, title, description, created_at, media_id, updated_at, source) OVERRIDING SYSTEM VALUE VALUES (19, 15, 11, 'Dukung Digitalisasi Informasi Desa, KKN UNSIKA 2026 Kembangkan SIPADES di Desa Cipicung', 'Purwakarta – Dalam mendukung transformasi digital di tingkat desa, KKN UNSIKA 2026 Kelompok 23 mengembangkan SIPADES (Sistem Informasi Profil Desa) sebagai media informasi digital yang memuat berbagai data dan potensi Desa Cipicung. Program ini bertujuan mempermudah masyarakat dalam mengakses informasi desa secara cepat, mudah, dan transparan.

Melalui SIPADES, berbagai informasi penting mengenai profil Desa Cipicung disajikan dalam satu platform, mulai dari profil desa, struktur pemerintahan, data kependudukan, potensi desa, berita dan kegiatan, galeri, hingga informasi lainnya yang dapat diakses oleh masyarakat. Kehadiran sistem ini diharapkan dapat menjadi sarana publikasi sekaligus media penyebaran informasi resmi desa.

Dalam proses pengembangannya, mahasiswa KKN UNSIKA 2026 berkolaborasi dengan Pemerintah Desa Cipicung untuk melakukan pengumpulan data, penyusunan konten, serta pengelolaan sistem agar informasi yang disajikan sesuai dengan kondisi dan kebutuhan desa.

Melalui program SIPADES, KKN UNSIKA 2026 berharap Desa Cipicung memiliki media informasi digital yang mampu meningkatkan keterbukaan informasi kepada masyarakat sekaligus memperkenalkan potensi desa kepada masyarakat luas sebagai bagian dari upaya mewujudkan desa yang adaptif terhadap perkembangan teknologi.', '2026-07-29 02:21:25.588406', 67, '2026-07-29 02:40:04.667045', NULL);
INSERT INTO public.documents (id, category_id, uploaded_by, title, description, created_at, media_id, updated_at, source) OVERRIDING SYSTEM VALUE VALUES (18, 15, 11, 'Tingkatkan Keamanan dan Kenyamanan Warga, KKN UNSIKA 2026 Hadirkan Program CITRANG di Desa Cipicung', 'Purwakarta – Sebagai upaya meningkatkan keamanan dan kenyamanan masyarakat, KKN UNSIKA 2026 Kelompok 23 melaksanakan program CITRANG (Cipicung Terang) melalui pemasangan Penerangan Jalan Umum (PJU) berbasis lampu panel surya di Desa Cipicung, Kecamatan Sukatani, Kabupaten Purwakarta.

Program CITRANG hadir sebagai solusi penerangan di titik-titik yang masih minim cahaya, sehingga dapat membantu meningkatkan visibilitas pengguna jalan, mengurangi potensi risiko kecelakaan, serta menciptakan lingkungan yang lebih aman bagi masyarakat, terutama pada malam hari.

Pemanfaatan lampu tenaga surya dipilih karena merupakan sumber energi terbarukan yang ramah lingkungan, hemat energi, serta tidak bergantung pada jaringan listrik. Dengan teknologi ini, penerangan jalan dapat beroperasi secara mandiri melalui energi matahari yang disimpan pada baterai untuk digunakan pada malam hari.

Melalui program CITRANG, KKN UNSIKA 2026 berharap keberadaan penerangan jalan umum berbasis panel surya dapat memberikan manfaat jangka panjang bagi masyarakat Desa Cipicung. Selain meningkatkan keamanan dan kenyamanan aktivitas warga pada malam hari, program ini juga menjadi bentuk dukungan terhadap pemanfaatan energi bersih dan pembangunan desa yang berkelanjutan.', '2026-07-29 02:20:04.711734', 68, '2026-07-29 02:40:31.039413', NULL);


--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: galleries; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (13, 5, 'Cek Kesehatan Gratis pada Masyarakat Desa Cipicung', 'Pemeriksaan kesehatan dan edukasi hidup sehat bagi masyarakat Desa Cipicung. Kolaborasi antara KKN UNSIKA 2026 dengan Puskesmas Sukatani', '2026-07-27 10:36:31.102654', 11, 54, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (14, 11, 'Sosialisasi dan Praktik BRIKUNG oleh KKN UNSIKA 2026', 'BRIKUNG (Briket Ramah Lingkungan Desa Cipicung) dengan memanfaatkan limbah sekam padi, diselenggarakan di Kantor Desa Cipicung', '2026-07-28 16:05:32.057771', 11, 55, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (15, 11, 'Senam Bersama Warga Desa Cipicung dan KKN UNSIKA 2026', 'Meningkatkan kesehatan, kebugaran, dan kebersamaan masyarakat Desa Cipicung melalui kegiatan senam bersama.', '2026-07-28 18:10:35.445796', 11, 57, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (18, 11, 'Tanaman Sayur (TAMSUR) untuk Ketahanan Pangan', 'Edukasi dan praktik budidaya tanaman sayur guna mendukung ketahanan pangan serta pemanfaatan lahan pekarangan masyarakat Desa Cipicung. Kolaborasi antara KKN UNSIKA 2026 dengan BPP Sukatani', '2026-07-28 19:37:25.196881', 11, 59, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (17, 11, 'Belajar Bersama di SDN Cipicung', 'Kegiatan pembelajaran interaktif untuk mendukung peningkatan semangat belajar dan wawasan siswa SDN Cipicung oleh KKN UNSIKA 2026', '2026-07-28 19:35:34.909576', 11, 58, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (16, 11, 'Ngosrek: Tradisi Gotong Royong Desa Cipicung', 'Tradisi kerja bakti masyarakat Desa Cipicung dalam menjaga kebersihan lingkungan serta mempererat kebersamaan warga.', '2026-07-28 18:14:41.136173', 11, 56, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (19, 11, 'Turnamen Sepak Bola Tingkat SD Antar-RT', 'Ajang kompetisi sepak bola bagi siswa tingkat SD antar-RT untuk menumbuhkan semangat sportivitas, kebersamaan, dan bakat olahraga di Desa Cipicung yang diselenggarakan oleh KKN UNSIKA 2026', '2026-07-28 19:39:15.361742', 11, 60, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (21, 11, 'CITRANG (Cipicung Terang)', 'Mahasiswa KKN UNSIKA melaksanakan pemasangan Penerangan Jalan Umum (PJU) berbasis panel surya untuk meningkatkan keamanan dan kenyamanan masyarakat Desa Cipicung.', '2026-07-29 02:26:59.767479', 11, 65, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (20, 11, 'Penyerahan Website SIPADES (Sistem Informasi Profil Desa)', 'Mahasiswa KKN UNSIKA 2026 menyerahkan website SIPADES kepada Pemerintah Desa Cipicung sebagai upaya mendukung digitalisasi informasi desa.', '2026-07-29 02:26:30.462189', 11, 66, '2026-07-29 12:19:08.57282', false, 'gallery');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (8, 11, 'Sosialisasi dan Praktik BRIKUNG', 'BRIKUNG (Briket Ramah Lingkungan) dengan memanfaatkan limbah sekam padi oleh KKN UNSIKA 2026', '2026-07-23 04:35:28.189876', 1, 37, '2026-07-29 12:19:08.57282', false, 'dashboard');
INSERT INTO public.galleries (id, created_by, title, description, created_at, category_id, media_id, updated_at, is_active, type) OVERRIDING SYSTEM VALUE VALUES (10, 11, 'Cipicung', 'Terletak di kaki Gunung Kacapi, Desa Cipicung tumbuh bersama potensi alamnya. Aktivitas pertambangan batu belah, pertanian, dan usaha masyarakat menjadi denyut kehidupan yang menggerakkan desa setiap hari.', '2026-07-24 08:24:17.677424', 1, 34, '2026-07-29 12:19:08.57282', true, 'dashboard');


--
-- Data for Name: villages; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.villages (id, name, province, regency, district, postal_code, address, phone, email, website, latitude, longitude, vision, history, description, created_at, updated_at, village_code, total_population, total_family, title, wide, mission, elevation, coordinate, population, region, hamlet_one, hamlet_two, north_border, east_border, south_border, west_border, area, is_active, total_rt, total_rw, rt_hamlet_one, rt_hamlet_two, rw_hamlet_one, rw_hamlet_two, ig_usn, tiktok_usn, yt_usn, total_male, total_female, demographic_religions, demographic_religion_rt, demographic_education, demographic_occupation, demographic_ages) OVERRIDING SYSTEM VALUE VALUES (9, 'Kantor Kepala Desa Kantor Kepala Desa Kantor Kepala Desa Desa Cipicung', 'Jawa Barat', 'Purwakarta', 'Sukatani', '41167', '99QM+24C, Cipicung, Kec. Sukatani, Kabupaten Purwakarta, Jawa Barat 41167', '083124581734', 'desacipicung20@gmail.com', '', -6.6075, 107.37667, '"Terwujudnya Desa Cipicung yang lebih maju, berprestasi, berbudaya, dan kreatif melalui peningkatan sumber daya manusia, pengelolaan sumber daya alam, serta pemantapan pembangunan berlandaskan keagamaan, kultural, dan budaya daerah.”', 'Desa Cipicung pada awalnya merupakan bagian dari Desa Sukamaju. Seiring pemekaran wilayah, Cipicung berdiri sebagai desa tersendiri dengan Sukamaju sebagai desa induk.

Nama "Cipicung" dipercaya berasal dari gabungan kata "Ci" (air) dan "Picung" (nama pohon). Menurut cerita para tetua, wilayah ini dulunya kawasan pegunungan dengan banyak mata air dan pohon picung tumbuh di sekitarnya.', 'Desa Cipicung terletak di Kecamatan Sukatani, Kabupaten Purwakarta, Jawa Barat. Memadukan nilai gotong royong dengan semangat wirausaha, warga desa mengelola hasil bumi seperti gula aren, arang, dan kerajinan menjadi produk yang dijual hingga ke luar desa.

Website ini adalah etalase kecil dari desa itu: profil wilayah, peta batas desa, dan potensi unggulan yang terus diperbarui.', '2026-07-24 07:20:25.299831', '2026-07-29 02:48:29.36333', NULL, NULL, 1107, 'Desa yang tumbuh dari usaha kecil.', NULL, '{"Meningkatkan ketersediaan dan kualitas infrastruktur pemerintahan desa.","Menggali potensi desa dalam rangka peningkatan Pendapatan Asli Desa (PADes).","Meningkatkan kualitas sumber daya manusia bagi aparatur dan masyarakat desa.","Meningkatkan profesionalisme pelayanan publik.","Meningkatkan kualitas hidup masyarakat di bidang kesehatan."}', '218 - 394 mdpl', '6°36''37" LS, 107°22''39" BT', '3271', 'Desa Cipicung berada di sisi barat Kecamatan Sukatani, berjarak sekitar 5 km dari kantor kecamatan dan 15 km dari pusat Kabupaten Purwakarta. Wilayahnya berupa perbukitan pada ketinggian 218–394 mdpl, terbagi menjadi dua dusun:', 1525, 1746, 'Tajur Sindang, Sukamaju', 'Cilalawi, Sukamaju', 'Linggunung, Pamoyanan', 'Sukamulya, Sindanglaya', '420 Ha (4,2 km²)', false, 11, 4, 5, 6, 2, 2, '@cipicungpemdes', '@pemdes.cipicung', '@desacipicung-x8k', 12, 1200, '[{"label": "islam", "value": 12}]', '[{"label": "2", "value": 12, "value2": 12}]', '[{"male": 12, "label": "s1", "female": 12}]', '[]', '[]');


--
-- Data for Name: officials; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (26, 9, 'Joko Mulyono', 'sekretaris-desa', NULL, NULL, NULL, 2, true, '2026-07-25 03:33:55.086297', '2026-07-25 03:48:23.847043', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (27, 9, 'Sadari Wiharta', 'Kaur-Tata-Usaha-&-Umum', NULL, NULL, NULL, 3, true, '2026-07-25 03:49:07.790755', '2026-07-25 04:01:14.36942', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (28, 9, 'Cucu Intan Sari', 'Kaur-Keuangan', NULL, NULL, NULL, 4, true, '2026-07-25 04:02:31.835011', '2026-07-25 04:02:31.835011', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (29, 9, 'Muhamad Deni Jatnika', 'Kaur-Perencanaan', NULL, NULL, NULL, 5, true, '2026-07-25 04:02:55.934047', '2026-07-25 04:02:55.934047', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (30, 9, 'Abdul Jalal Wiharja', 'Kasi-Pemerintahan', NULL, NULL, NULL, 6, true, '2026-07-25 04:03:16.6535', '2026-07-25 04:03:16.6535', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (31, 9, 'Hermawan Sutisna', 'Kasi-Kesejahteraan', NULL, NULL, NULL, 7, true, '2026-07-25 04:03:31.605127', '2026-07-25 04:03:31.605127', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (32, 9, 'Ahmad Jaeni Tahir', 'Kasi-Pelayanan', NULL, NULL, NULL, 8, true, '2026-07-25 04:03:49.487439', '2026-07-25 04:03:49.487439', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (33, 9, 'Enjang', 'Kepala-Dusun-I', NULL, NULL, NULL, 9, true, '2026-07-25 04:04:06.19182', '2026-07-25 04:04:06.19182', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (34, 9, 'Asep Saepuloh', 'Kepala-Dusun-II', NULL, NULL, NULL, 10, true, '2026-07-25 04:04:25.721888', '2026-07-25 04:04:25.721888', NULL, NULL);
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (65, 9, 'Lili Sadili', 'kepala-desa', '', '', '', 0, true, '2026-07-27 03:26:05.244848', '2026-07-27 03:26:05.244848', '2013-01-01', '2026-07-24');
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (66, 9, 'Edi Supriadi', 'kepala-desa', '', '', '', 0, false, '2026-07-27 03:26:05.244848', '2026-07-27 03:26:05.244848', '2007-01-01', '2012-12-31');
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (67, 9, 'Muhidin', 'kepala-desa', '', '', '', 0, false, '2026-07-27 03:26:05.244848', '2026-07-27 03:26:05.244848', '2000-01-01', '2006-12-31');
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (68, 9, 'Asep Hidayat', 'kepala-desa', '', '', '', 0, false, '2026-07-27 03:26:05.244848', '2026-07-27 03:26:05.244848', '1987-01-01', '1998-12-31');
INSERT INTO public.officials (id, village_id, name, "position", phone, email, description, order_number, is_active, created_at, updated_at, start_date, finish_date) OVERRIDING SYSTEM VALUE VALUES (69, 9, 'Mahpudin', 'kepala-desa', '', '', '', 0, false, '2026-07-27 03:26:05.244848', '2026-07-27 03:26:05.244848', '1983-01-01', '1986-12-31');


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: potential_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: potentials; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.potentials (id, title, subtitle, description, owner_name, owner_msisdn, category_id, slug, media_id, location_id, created_at, updated_at) VALUES (7, 'Gula Aren', 'RT 10', 'Nira disadap setiap pagi dan sore dari pohon aren di sekitar desa, lalu dimasak perlahan di atas tungku kayu hingga mengental dan dicetak menjadi gula batok. Proses ini dikerjakan turun-temurun oleh warga sebagai salah satu sumber penghasilan utama di lua', NULL, NULL, 12, 'gula-aren', 42, NULL, '2026-07-25 07:44:44.067309', '2026-07-25 07:44:44.067309');
INSERT INTO public.potentials (id, title, subtitle, description, owner_name, owner_msisdn, category_id, slug, media_id, location_id, created_at, updated_at) VALUES (8, 'Tusuk Sate', 'RT 03 & RT 04', 'Bambu dibelah, dijemur, lalu diraut satu per satu secara manual oleh pengrajin rumahan. Hasilnya berupa tusuk sate yang dipasok ke pasar dan pedagang di luar desa, menjadi kegiatan usaha rumahan yang cukup banyak digeluti warga.', NULL, NULL, 13, 'tusuk-sate', 43, NULL, '2026-07-25 07:45:14.093048', '2026-07-25 07:45:14.093048');
INSERT INTO public.potentials (id, title, subtitle, description, owner_name, owner_msisdn, category_id, slug, media_id, location_id, created_at, updated_at) VALUES (9, 'Badan Usaha Milik Desa', 'RT 02', 'Bumi Desa menaungi beberapa unit usaha yang melayani kebutuhan harian warga Cipicung mulai dari penyediaan pangan, air bersih, kebutuhan rumah tangga, sampai akses layanan keuangan tanpa harus keluar desa.', NULL, NULL, 14, 'badan-usaha-milik-desa', 44, NULL, '2026-07-25 07:45:45.817327', '2026-07-28 15:59:30.295654');


--
-- Data for Name: user_session_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (1, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODMxNDcxOTYsImV4cCI6MTc4MzE0ODA5Nn0.vAOm8Pz6aQbGtqn8eHk6cgbfB60QE6SzAd8aXG6VRmE', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzMTQ3MTk2LCJleHAiOjE3ODM3NTE5OTZ9.tFsJT9X59Wij2cAWfwzKt8BdtGPIt7nHVV0y3ucmzlo', '2026-07-04 13:39:56.014167');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (2, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM2Nzc5MDEsImV4cCI6MTc4MzY3ODgwMX0.DQGYZ-YeYcpRsdcimiY2fdLj-zeDglH5K7WsqnaWz7M', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNjc3OTAxLCJleHAiOjE3ODQyODI3MDF9.Q_RNOXKTh1wEA0SY5VDltCCF-yAwCAd41j0Q4EGaul8', '2026-07-10 17:05:01.576576');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (3, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM2ODIzNzEsImV4cCI6MTc4MzY4MzI3MX0.fFBdKH7bS78TZlvAk1tUu_lOjWcueLAuyLm33t2qu9E', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNjgyMzcxLCJleHAiOjE3ODQyODcxNzF9.ZIsZ8sPAEw7MimQOWTz5G5shrXxtj9lo5HOTeSRZ_Lw', '2026-07-10 18:19:31.973338');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (4, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM3Njk1NzgsImV4cCI6MTc4Mzc3MDQ3OH0.BU7VjZxZZz_isco-ge62NgKYU4BzmEZKRsYcfxLE2Qc', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNzY5NTc4LCJleHAiOjE3ODQzNzQzNzh9.K9AEi9cnVor5P8HSWTGX7K9gWEVVHw1Qe-xX6-CSVS4', '2026-07-11 18:32:58.538562');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (5, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM3ODA5NDYsImV4cCI6MTc4Mzc4MTg0Nn0.V0lFihfXog_BCKSEa25JJezFGgyWNbGC-bbaitBb2Lc', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNzgwOTQ2LCJleHAiOjE3ODQzODU3NDZ9.Jj1Y8inuBKLhADfB1x7uYyz3cIXK0_w6kUQN2WHL8iU', '2026-07-11 21:42:26.984505');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (6, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM4NjU2NjAsImV4cCI6MTc4Mzg2NjU2MH0.hCTL1ciZ1Mx83RMBoYhci17KqzjKVTYGlzvAD-Ua02A', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzODY1NjYwLCJleHAiOjE3ODQ0NzA0NjB9.yM2WyxDYr6uEztOfS9a6qJB6VzIgv_HRTD7rtJVxA8Y', '2026-07-12 21:14:20.316928');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (7, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM4NjcxOTQsImV4cCI6MTc4Mzg2ODA5NH0.l39-VN5bWYMriRMe4ml_JQoskUSWPMpSvlDeTedQLRs', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzODY3MTk0LCJleHAiOjE3ODQ0NzE5OTR9.OMDuts7iVnezSs0q-NOxMvZExIKHurOK4ncsOsnMUgQ', '2026-07-12 21:39:54.213027');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (8, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM5NDk0MzcsImV4cCI6MTc4Mzk1MDMzN30.dniYiag_LfVagi_mRmRX9TUARvnoUJJXvqMAIRghxmY', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzOTQ5NDM3LCJleHAiOjE3ODQ1NTQyMzd9.sQN3aa9OSy4zYD5OeB8G3zAFjkG8tNPN3GpB9UXWD1w', '2026-07-13 20:30:37.447884');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (9, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM5NTc1NDIsImV4cCI6MTc4Mzk1ODQ0Mn0.Z9r7Os7dhMoM3otIvabyap4Ga1CqyoUIE8wNJrmepYc', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzOTU3NTQyLCJleHAiOjE3ODQ1NjIzNDJ9.nSzFbH6m4iS4IXt9ag1w_SqAvzsvns_ObdX0EPHQSyk', '2026-07-13 22:45:42.747328');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (10, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQwMzM5NjMsImV4cCI6MTc4NDYzODc2M30.vuhXg2eSZ4g5BJkPPsj4fYIVrKBIYkJyXz62sc4bl0E', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0MDMzOTYzLCJleHAiOjE3ODQ2Mzg3NjN9.O36aKCtEdYkE88RH_G6DX1l9ca3aMr0ZJZMacuxwNfY', '2026-07-14 19:59:23.114334');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (11, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQyOTgxODgsImV4cCI6MTc4NDkwMjk4OH0.o_rgsDsq0pBr4ldCKoP2lpXm0rteEBbFS_FQ8xVM0QQ', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0Mjk4MTg4LCJleHAiOjE3ODQ5MDI5ODh9.y1em8-5Tvn7J6tbj_IiemqzTNvVKzosJG9_L8H-Bckg', '2026-07-17 21:23:08.354655');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (16, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQ2NDEyNjIsImV4cCI6MTc4NTI0NjA2Mn0.R2Lmbuh1waUhPCBlLxkPfAsVxmrKbT4PLpjlqGnJpc0', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0NjQxMjYyLCJleHAiOjE3ODUyNDYwNjJ9.Se3qQDIJkB-lKf-LZSbKg2OK6dmXQXiyYbWthj23lc8', '2026-07-21 13:41:02.931302');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (17, 11, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQ2NDQ1NjAsImV4cCI6MTc4NTI0OTM2MH0.GmwdxgbZEtIgQdvpuMBl6LuAFwLAsqwhVtr4OiuIVzM', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0NjQ0NTYwLCJleHAiOjE3ODUyNDkzNjB9.5J2Xv7HQVFAQXNmF4y0okPd5jnQiK7dftRoGQRvlP4k', '2026-07-21 14:36:00.788233');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (18, 11, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQ3MDc0MDcsImV4cCI6MTc4NTMxMjIwN30.QqhNAMX5Q1AXVwykrWk4Qw3GWzyAw6dozmMTjghqNx0', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0NzA3NDA3LCJleHAiOjE3ODUzMTIyMDd9.w1vNaYHhI7K6VaKGf32FoL1YeMme2ttvL4t8gNhYdQg', '2026-07-22 08:03:27.09484');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (19, 11, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQ3NzM1ODIsImV4cCI6MTc4NTM3ODM4Mn0.XpORQ5-bLoBej4Bw1KBUPpNvbtTHBjkk_Ud8_txunSk', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0NzczNTgyLCJleHAiOjE3ODUzNzgzODJ9.iou_8LuiUhkK06Kuv9mcWLALemsIGwLrPbgMxlJlTyo', '2026-07-23 02:26:22.548419');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (20, 11, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQ3ODAzNDAsImV4cCI6MTc4NTM4NTE0MH0.LZ202cxbZO-qui5E2VOzow0hOIQjwnVqeadP0LrBirY', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0NzgwMzQwLCJleHAiOjE3ODUzODUxNDB9.nMk4V3tpwrO_7LJFDxcKOTb3Wxzo-3GRf4Y0S0UoKX0', '2026-07-23 04:19:00.89298');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (21, 11, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODUwNzM1MTYsImV4cCI6MTc4NTY3ODMxNn0.i0R66JmkWt42bmXfPIT8oVBOa3r8lWIaWbzIG2E_xiE', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg1MDczNTE2LCJleHAiOjE3ODU2NzgzMTZ9.J6-8f1Sk8Wo1M_UgU5Pg6dnjsZ9n7k_O1DY_o-zXgZw', '2026-07-26 13:45:16.091239');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (22, 12, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMiwidXNlcm5hbWUiOiJyYW1hZGl0Iiwicm9sZSI6ImFkbWluIiwidG9rZW5fdHlwZSI6ImFjY2VzcyIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg1MDgyNDA5LCJleHAiOjE3ODU2ODcyMDl9.i9xcW10dcVy69GNX0n4PuNqsmdoRBSwOpZe4UtQt1HY', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMiwidXNlcm5hbWUiOiJyYW1hZGl0Iiwicm9sZSI6ImFkbWluIiwidG9rZW5fdHlwZSI6InJlZnJlc2giLCJpc3MiOiJjaXBpY3VuZy5pZCIsImlhdCI6MTc4NTA4MjQwOSwiZXhwIjoxNzg1Njg3MjA5fQ.Rf6HTZp5xDeSdLpue6Ln4VbXodvVzAS078bf39MCqkI', '2026-07-26 16:13:29.238297');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (23, 12, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMiwidXNlcm5hbWUiOiJyYW1hZGl0Iiwicm9sZSI6ImFkbWluIiwidG9rZW5fdHlwZSI6ImFjY2VzcyIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg1MDkyMzE5LCJleHAiOjE3ODU2OTcxMTl9.nHYHlDOl-h_bu4kNMfJyNaAVxQzq-adtT-9gtI4sx4g', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMiwidXNlcm5hbWUiOiJyYW1hZGl0Iiwicm9sZSI6ImFkbWluIiwidG9rZW5fdHlwZSI6InJlZnJlc2giLCJpc3MiOiJjaXBpY3VuZy5pZCIsImlhdCI6MTc4NTA5MjMxOSwiZXhwIjoxNzg1Njk3MTE5fQ.m2TOtDoXDJE82sexVk2wsMuunEd5odmQjVVJ_hkIaus', '2026-07-26 18:58:39.633731');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (24, 13, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMywidXNlcm5hbWUiOiJjaXBpY3VuZzI2Iiwicm9sZSI6ImFkbWluIiwidG9rZW5fdHlwZSI6ImFjY2VzcyIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg1MTI3NTQwLCJleHAiOjE3ODU3MzIzNDB9.chdoJc3p9kaJVui--jkHNynOg2svuwq7iMZsZpp_ARI', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMywidXNlcm5hbWUiOiJjaXBpY3VuZzI2Iiwicm9sZSI6ImFkbWluIiwidG9rZW5fdHlwZSI6InJlZnJlc2giLCJpc3MiOiJjaXBpY3VuZy5pZCIsImlhdCI6MTc4NTEyNzU0MCwiZXhwIjoxNzg1NzMyMzQwfQ.ZGE_r2-CRPCW4ZSsMKSkRM5HNCOpDW31sZNiEcyCwQw', '2026-07-27 04:45:40.695245');
INSERT INTO public.user_session_log (id, user_id, access_token, refresh_token, created_at) VALUES (25, 11, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODUyOTIzMzgsImV4cCI6MTc4NTg5NzEzOH0.o5Ti7A1zPCchYL-lZ-lvc6orjLVZmk0JtZm2-akRL9A', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMSwidXNlcm5hbWUiOiJqaWEiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg1MjkyMzM4LCJleHAiOjE3ODU4OTcxMzh9.sp4I5GML-1ok3XyEgPBwbZjylwd7O_jf6yCma0_m4lI', '2026-07-29 02:32:18.876877');


--
-- Name: activity_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.activity_logs_id_seq', 1, false);


--
-- Name: businesses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.businesses_id_seq', 3, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 15, true);


--
-- Name: documents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.documents_id_seq', 19, true);


--
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.events_id_seq', 1, false);


--
-- Name: galleries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.galleries_id_seq', 21, true);


--
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.locations_id_seq', 3, true);


--
-- Name: media_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.media_id_seq', 69, true);


--
-- Name: officials_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.officials_id_seq', 69, true);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.posts_id_seq', 1, false);


--
-- Name: potential_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.potential_detail_id_seq', 1, false);


--
-- Name: potentials_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.potentials_id_seq', 10, true);


--
-- Name: user_session_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_session_log_id_seq', 25, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 13, true);


--
-- Name: villages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.villages_id_seq', 11, true);


--
-- PostgreSQL database dump complete
--

\unrestrict fTkA4ED6k6tG2m5TfekAiZ0fbxcijU5lGxZS0PwxlW9xSxoPoU8AiqyLaqObgEi

