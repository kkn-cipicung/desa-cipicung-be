--
-- PostgreSQL database dump
--

\restrict KBfONGEFv8LCUUzzCJOAHliGmZLJURbJyTPgReYTdHhtmentDybTYN8gzsKZ2Sq

-- Dumped from database version 18.3 (Ubuntu 18.3-1.pgdg24.04+1)
-- Dumped by pg_dump version 18.3 (Ubuntu 18.3-1.pgdg24.04+1)

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
-- Name: set_updated_at(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.set_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_updated_at() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: activity_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.activity_logs (
    id integer NOT NULL,
    user_id integer,
    activity text,
    ip character varying(45),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.activity_logs OWNER TO postgres;

--
-- Name: activity_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.activity_logs ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.activity_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: businesses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.businesses (
    id integer NOT NULL,
    category_id integer,
    owner_name character varying(150),
    business_name character varying(255) NOT NULL,
    description text,
    phone character varying(30),
    address text,
    instagram character varying(255),
    facebook character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    location_id integer
);


ALTER TABLE public.businesses OWNER TO postgres;

--
-- Name: businesses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.businesses ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.businesses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    slug character varying(150) NOT NULL,
    type character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.categories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: documents; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.documents (
    id integer NOT NULL,
    category_id integer,
    uploaded_by integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    media_id integer,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    source character varying(100)
);


ALTER TABLE public.documents OWNER TO postgres;

--
-- Name: documents_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.documents ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.documents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.events (
    id integer NOT NULL,
    author_id integer NOT NULL,
    title character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    description text,
    location character varying(255),
    start_date timestamp without time zone,
    end_date timestamp without time zone,
    status character varying(30),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    category_id integer NOT NULL
);


ALTER TABLE public.events OWNER TO postgres;

--
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.events ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.events_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: galleries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.galleries (
    id integer NOT NULL,
    created_by integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    category_id integer NOT NULL,
    media_id integer,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    is_active boolean DEFAULT false
);


ALTER TABLE public.galleries OWNER TO postgres;

--
-- Name: galleries_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.galleries ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.galleries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: locations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.locations (
    id integer NOT NULL,
    latitude double precision,
    longitude double precision,
    created_by_id integer,
    title character varying(100),
    description character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.locations OWNER TO postgres;

--
-- Name: locations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.locations ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.locations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: media; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.media (
    id integer NOT NULL,
    uploaded_by integer NOT NULL,
    entity_type character varying(50) NOT NULL,
    entity_id integer NOT NULL,
    role character varying(50) NOT NULL,
    original_name character varying(255),
    file_name character varying(255),
    file_path character varying(500),
    mime_type character varying(100),
    file_size bigint,
    caption character varying(255),
    order_number integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.media OWNER TO postgres;

--
-- Name: media_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.media ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.media_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: officials; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.officials (
    id integer NOT NULL,
    village_id integer NOT NULL,
    name character varying(150) NOT NULL,
    "position" character varying(100) NOT NULL,
    phone character varying(30),
    email character varying(150),
    description text,
    order_number integer DEFAULT 0,
    is_active boolean DEFAULT true,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    start_date date,
    finish_date date
);


ALTER TABLE public.officials OWNER TO postgres;

--
-- Name: officials_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.officials ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.officials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    id integer NOT NULL,
    category_id integer,
    author_id integer NOT NULL,
    type character varying(50) NOT NULL,
    title character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    excerpt text,
    content text,
    publish_start timestamp without time zone,
    publish_end timestamp without time zone,
    is_pinned boolean DEFAULT false,
    status character varying(30) DEFAULT 'draft'::character varying,
    views integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.posts ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: potential_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.potential_detail (
    title character varying(100),
    detail character varying(100),
    description character varying(255),
    id bigint NOT NULL
);


ALTER TABLE public.potential_detail OWNER TO postgres;

--
-- Name: potential_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.potential_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.potential_detail_id_seq OWNER TO postgres;

--
-- Name: potential_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.potential_detail_id_seq OWNED BY public.potential_detail.id;


--
-- Name: potentials; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.potentials (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    subtitle character varying(100),
    description character varying(255),
    owner_name character varying,
    owner_msisdn character varying(100),
    category_id integer,
    slug character varying(100),
    media_id integer,
    location_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.potentials OWNER TO postgres;

--
-- Name: potentials_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.potentials ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.potentials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: user_session_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_session_log (
    id integer NOT NULL,
    user_id integer NOT NULL,
    access_token text NOT NULL,
    refresh_token text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.user_session_log OWNER TO postgres;

--
-- Name: user_session_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_session_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_session_log_id_seq OWNER TO postgres;

--
-- Name: user_session_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_session_log_id_seq OWNED BY public.user_session_log.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    role character varying(20) NOT NULL,
    password character varying(255) NOT NULL,
    is_active boolean DEFAULT true,
    last_login timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    username character varying(100),
    name character varying(100)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: villages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.villages (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    province character varying(100),
    regency character varying(100),
    district character varying(100),
    postal_code character varying(10),
    address text,
    phone character varying(30),
    email character varying(150),
    website character varying(255),
    latitude numeric(10,7),
    longitude numeric(10,7),
    vision text,
    history text,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    village_code character varying(100),
    total_population integer,
    total_family integer,
    title character varying(100),
    wide character varying(100),
    mission character varying(100)[],
    elevation character varying,
    coordinate character varying,
    population text,
    region text,
    hamlet_one text,
    hamlet_two text,
    north_border text,
    east_border text,
    south_border text,
    west_border text,
    area text,
    is_active boolean DEFAULT false NOT NULL
);


ALTER TABLE public.villages OWNER TO postgres;

--
-- Name: villages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.villages ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.villages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: potential_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.potential_detail ALTER COLUMN id SET DEFAULT nextval('public.potential_detail_id_seq'::regclass);


--
-- Name: user_session_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_session_log ALTER COLUMN id SET DEFAULT nextval('public.user_session_log_id_seq'::regclass);


--
-- Data for Name: activity_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: businesses; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.businesses OVERRIDING SYSTEM VALUE VALUES (2, 1, 'ilham', 'aren gula', 'ilham', '0990099990', 'dskasdkdask', NULL, NULL, '2026-07-13 20:41:20.419611', '2026-07-13 20:41:20.419611', NULL) ON CONFLICT DO NOTHING;


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories OVERRIDING SYSTEM VALUE VALUES (1, 'Dashboard', 'dashboard', 'dashboard', '2026-07-04 11:40:35.863008', '2026-07-04 11:40:35.863008') ON CONFLICT DO NOTHING;
INSERT INTO public.categories OVERRIDING SYSTEM VALUE VALUES (2, 'Pendidikan', 'pendidikan', 'news', '2026-07-09 00:21:26.60145', '2026-07-09 00:21:26.60145') ON CONFLICT DO NOTHING;
INSERT INTO public.categories OVERRIDING SYSTEM VALUE VALUES (3, 'Kuliner', 'kuliner', 'business', '2026-07-10 17:34:32.745891', '2026-07-10 17:34:32.745891') ON CONFLICT DO NOTHING;
INSERT INTO public.categories OVERRIDING SYSTEM VALUE VALUES (11, 'Galeri', 'galeri', 'gallery', '2026-07-20 16:32:39.906825', '2026-07-20 16:32:39.906825') ON CONFLICT DO NOTHING;


--
-- Data for Name: documents; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.documents OVERRIDING SYSTEM VALUE VALUES (4, 2, 6, 'ilham', '12133113', '2026-07-12 21:40:20.639767', NULL, '2026-07-12 21:40:20.639767', NULL) ON CONFLICT DO NOTHING;
INSERT INTO public.documents OVERRIDING SYSTEM VALUE VALUES (5, 1, 5, 'ilham', 'dakasdksadk', '2026-07-13 22:46:45.748125', NULL, '2026-07-13 22:46:45.748125', NULL) ON CONFLICT DO NOTHING;
INSERT INTO public.documents OVERRIDING SYSTEM VALUE VALUES (7, 2, 6, 'dasdas', 'dsadsa', '2026-07-20 16:34:58.170009', 21, '2026-07-20 16:34:58.170009', NULL) ON CONFLICT DO NOTHING;


--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: galleries; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.galleries OVERRIDING SYSTEM VALUE VALUES (1, 5, 'dashboard', 'dashboard2', '2026-07-04 13:34:57.553767', 1, NULL, '2026-07-20 11:55:37.138943', false) ON CONFLICT DO NOTHING;
INSERT INTO public.galleries OVERRIDING SYSTEM VALUE VALUES (2, 5, 'dashboard', 'dashboard2', '2026-07-04 13:41:08.130782', 1, 15, '2026-07-20 15:32:28.514496', true) ON CONFLICT DO NOTHING;
INSERT INTO public.galleries OVERRIDING SYSTEM VALUE VALUES (7, 6, 'dsakasdkdsa', 'dsaksaksadkdsa', '2026-07-20 16:33:22.892896', 11, 20, '2026-07-20 16:33:22.892896', false) ON CONFLICT DO NOTHING;


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: media; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.media OVERRIDING SYSTEM VALUE VALUES (15, 6, 'dashboard', 2, 'image', NULL, NULL, 'uploads/dashboard/1784536348514020947.png', 'image/png', NULL, NULL, 0, '2026-07-20 15:32:28.514496', '2026-07-20 15:32:28.514496') ON CONFLICT DO NOTHING;
INSERT INTO public.media OVERRIDING SYSTEM VALUE VALUES (20, 6, 'gallery', 7, 'image', NULL, NULL, 'uploads/gallery/1784540002891569326.png', 'image/png', NULL, NULL, 0, '2026-07-20 16:33:22.892896', '2026-07-20 16:33:22.892896') ON CONFLICT DO NOTHING;
INSERT INTO public.media OVERRIDING SYSTEM VALUE VALUES (21, 6, 'news', 7, 'image', NULL, NULL, 'uploads/news/1784540098169274813.png', 'image/png', NULL, NULL, 0, '2026-07-20 16:34:58.170009', '2026-07-20 16:34:58.170009') ON CONFLICT DO NOTHING;


--
-- Data for Name: officials; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: potential_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: potentials; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.potentials VALUES (1, 'ilham', 'ilhamm', 'sdkaskasd', 'ilham', '121122112', 1, 'ilham', NULL, NULL, '2026-07-14 21:52:27.089467', '2026-07-14 21:52:27.089467') ON CONFLICT DO NOTHING;
INSERT INTO public.potentials VALUES (2, 'ilham', 'ilhamm', 'sdkaskasd', 'ilham', '121122112', 1, 'ilham', NULL, NULL, '2026-07-14 21:52:50.540365', '2026-07-14 21:52:50.540365') ON CONFLICT DO NOTHING;


--
-- Data for Name: user_session_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_session_log VALUES (1, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODMxNDcxOTYsImV4cCI6MTc4MzE0ODA5Nn0.vAOm8Pz6aQbGtqn8eHk6cgbfB60QE6SzAd8aXG6VRmE', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzMTQ3MTk2LCJleHAiOjE3ODM3NTE5OTZ9.tFsJT9X59Wij2cAWfwzKt8BdtGPIt7nHVV0y3ucmzlo', '2026-07-04 13:39:56.014167') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (2, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM2Nzc5MDEsImV4cCI6MTc4MzY3ODgwMX0.DQGYZ-YeYcpRsdcimiY2fdLj-zeDglH5K7WsqnaWz7M', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNjc3OTAxLCJleHAiOjE3ODQyODI3MDF9.Q_RNOXKTh1wEA0SY5VDltCCF-yAwCAd41j0Q4EGaul8', '2026-07-10 17:05:01.576576') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (3, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM2ODIzNzEsImV4cCI6MTc4MzY4MzI3MX0.fFBdKH7bS78TZlvAk1tUu_lOjWcueLAuyLm33t2qu9E', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNjgyMzcxLCJleHAiOjE3ODQyODcxNzF9.ZIsZ8sPAEw7MimQOWTz5G5shrXxtj9lo5HOTeSRZ_Lw', '2026-07-10 18:19:31.973338') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (4, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM3Njk1NzgsImV4cCI6MTc4Mzc3MDQ3OH0.BU7VjZxZZz_isco-ge62NgKYU4BzmEZKRsYcfxLE2Qc', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNzY5NTc4LCJleHAiOjE3ODQzNzQzNzh9.K9AEi9cnVor5P8HSWTGX7K9gWEVVHw1Qe-xX6-CSVS4', '2026-07-11 18:32:58.538562') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (5, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM3ODA5NDYsImV4cCI6MTc4Mzc4MTg0Nn0.V0lFihfXog_BCKSEa25JJezFGgyWNbGC-bbaitBb2Lc', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzNzgwOTQ2LCJleHAiOjE3ODQzODU3NDZ9.Jj1Y8inuBKLhADfB1x7uYyz3cIXK0_w6kUQN2WHL8iU', '2026-07-11 21:42:26.984505') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (6, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM4NjU2NjAsImV4cCI6MTc4Mzg2NjU2MH0.hCTL1ciZ1Mx83RMBoYhci17KqzjKVTYGlzvAD-Ua02A', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzODY1NjYwLCJleHAiOjE3ODQ0NzA0NjB9.yM2WyxDYr6uEztOfS9a6qJB6VzIgv_HRTD7rtJVxA8Y', '2026-07-12 21:14:20.316928') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (7, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM4NjcxOTQsImV4cCI6MTc4Mzg2ODA5NH0.l39-VN5bWYMriRMe4ml_JQoskUSWPMpSvlDeTedQLRs', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzODY3MTk0LCJleHAiOjE3ODQ0NzE5OTR9.OMDuts7iVnezSs0q-NOxMvZExIKHurOK4ncsOsnMUgQ', '2026-07-12 21:39:54.213027') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (8, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM5NDk0MzcsImV4cCI6MTc4Mzk1MDMzN30.dniYiag_LfVagi_mRmRX9TUARvnoUJJXvqMAIRghxmY', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzOTQ5NDM3LCJleHAiOjE3ODQ1NTQyMzd9.sQN3aa9OSy4zYD5OeB8G3zAFjkG8tNPN3GpB9UXWD1w', '2026-07-13 20:30:37.447884') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (9, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODM5NTc1NDIsImV4cCI6MTc4Mzk1ODQ0Mn0.Z9r7Os7dhMoM3otIvabyap4Ga1CqyoUIE8wNJrmepYc', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzgzOTU3NTQyLCJleHAiOjE3ODQ1NjIzNDJ9.nSzFbH6m4iS4IXt9ag1w_SqAvzsvns_ObdX0EPHQSyk', '2026-07-13 22:45:42.747328') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (10, 5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQwMzM5NjMsImV4cCI6MTc4NDYzODc2M30.vuhXg2eSZ4g5BJkPPsj4fYIVrKBIYkJyXz62sc4bl0E', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJ1c2VybmFtZSI6ImlsaGFtZ29kMTQiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0MDMzOTYzLCJleHAiOjE3ODQ2Mzg3NjN9.O36aKCtEdYkE88RH_G6DX1l9ca3aMr0ZJZMacuxwNfY', '2026-07-14 19:59:23.114334') ON CONFLICT DO NOTHING;
INSERT INTO public.user_session_log VALUES (11, 6, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiaXNzIjoiY2lwaWN1bmcuaWQiLCJpYXQiOjE3ODQyOTgxODgsImV4cCI6MTc4NDkwMjk4OH0.o_rgsDsq0pBr4ldCKoP2lpXm0rteEBbFS_FQ8xVM0QQ', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJ1c2VybmFtZSI6InJvYnkiLCJyb2xlIjoiYWRtaW4iLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImlzcyI6ImNpcGljdW5nLmlkIiwiaWF0IjoxNzg0Mjk4MTg4LCJleHAiOjE3ODQ5MDI5ODh9.y1em8-5Tvn7J6tbj_IiemqzTNvVKzosJG9_L8H-Bckg', '2026-07-17 21:23:08.354655') ON CONFLICT DO NOTHING;


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users OVERRIDING SYSTEM VALUE VALUES (5, 'admin', '$2a$10$vM7LYExPv2CNvn9kMilcC.0ov2UH3ZhXAQ/SPV790bg.aBMZY554a', true, '2026-07-14 19:59:23.103474', '2026-07-04 10:57:40.252264', '2026-07-04 10:57:40.252264', 'ilhamgod14', 'ilham') ON CONFLICT DO NOTHING;
INSERT INTO public.users OVERRIDING SYSTEM VALUE VALUES (6, 'admin', '$2a$10$dKemcSEvrqE4Q2k2Y7XKzeP1VqtnTW8amqfUAV.MvmN3Z.QRiEaSO', true, '2026-07-17 21:23:08.349048', '2026-07-10 17:00:38.372401', '2026-07-17 21:23:08.349048', 'roby', 'ilham') ON CONFLICT DO NOTHING;


--
-- Data for Name: villages; Type: TABLE DATA; Schema: public; Owner: postgres
--



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

SELECT pg_catalog.setval('public.categories_id_seq', 11, true);


--
-- Name: documents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.documents_id_seq', 7, true);


--
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.events_id_seq', 1, false);


--
-- Name: galleries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.galleries_id_seq', 7, true);


--
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.locations_id_seq', 3, true);


--
-- Name: media_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.media_id_seq', 21, true);


--
-- Name: officials_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.officials_id_seq', 2, true);


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

SELECT pg_catalog.setval('public.potentials_id_seq', 4, true);


--
-- Name: user_session_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_session_log_id_seq', 15, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 10, true);


--
-- Name: villages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.villages_id_seq', 8, true);


--
-- Name: activity_logs activity_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity_logs
    ADD CONSTRAINT activity_logs_pkey PRIMARY KEY (id);


--
-- Name: businesses businesses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: categories categories_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_slug_key UNIQUE (slug);


--
-- Name: documents documents_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.documents
    ADD CONSTRAINT documents_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: events events_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_slug_key UNIQUE (slug);


--
-- Name: galleries galleries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.galleries
    ADD CONSTRAINT galleries_pkey PRIMARY KEY (id);


--
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);


--
-- Name: media media_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.media
    ADD CONSTRAINT media_pkey PRIMARY KEY (id);


--
-- Name: officials officials_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.officials
    ADD CONSTRAINT officials_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: posts posts_slug_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_slug_key UNIQUE (slug);


--
-- Name: potentials potentials_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.potentials
    ADD CONSTRAINT potentials_pkey PRIMARY KEY (id);


--
-- Name: user_session_log user_session_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_session_log
    ADD CONSTRAINT user_session_log_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: villages villages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.villages
    ADD CONSTRAINT villages_pkey PRIMARY KEY (id);


--
-- Name: idx_activity_user; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_activity_user ON public.activity_logs USING btree (user_id);


--
-- Name: idx_business_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_business_name ON public.businesses USING btree (business_name);


--
-- Name: idx_document_title; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_document_title ON public.documents USING btree (title);


--
-- Name: idx_events_start_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_events_start_date ON public.events USING btree (start_date);


--
-- Name: idx_media_entity; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_media_entity ON public.media USING btree (entity_type, entity_id);


--
-- Name: idx_official_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_official_order ON public.officials USING btree (order_number);


--
-- Name: idx_posts_publish_start; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_publish_start ON public.posts USING btree (publish_start);


--
-- Name: idx_posts_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_status ON public.posts USING btree (status);


--
-- Name: idx_posts_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_posts_type ON public.posts USING btree (type);


--
-- Name: activity_logs set_activity_logs_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_activity_logs_updated_at BEFORE UPDATE ON public.activity_logs FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: businesses set_businesses_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_businesses_updated_at BEFORE UPDATE ON public.businesses FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: categories set_categories_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_categories_updated_at BEFORE UPDATE ON public.categories FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: documents set_documents_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_documents_updated_at BEFORE UPDATE ON public.documents FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: events set_events_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_events_updated_at BEFORE UPDATE ON public.events FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: galleries set_galleries_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_galleries_updated_at BEFORE UPDATE ON public.galleries FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: locations set_locations_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_locations_updated_at BEFORE UPDATE ON public.locations FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: media set_media_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_media_updated_at BEFORE UPDATE ON public.media FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: officials set_officials_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_officials_updated_at BEFORE UPDATE ON public.officials FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: posts set_posts_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_posts_updated_at BEFORE UPDATE ON public.posts FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: potentials set_potentials_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_potentials_updated_at BEFORE UPDATE ON public.potentials FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: users set_users_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: villages set_villages_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER set_villages_updated_at BEFORE UPDATE ON public.villages FOR EACH ROW EXECUTE FUNCTION public.set_updated_at();


--
-- Name: activity_logs activity_logs_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.activity_logs
    ADD CONSTRAINT activity_logs_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: businesses businesses_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL;


--
-- Name: businesses businesses_location_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_location_id_fkey FOREIGN KEY (location_id) REFERENCES public.locations(id);


--
-- Name: documents documents_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.documents
    ADD CONSTRAINT documents_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL;


--
-- Name: documents documents_uploaded_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.documents
    ADD CONSTRAINT documents_uploaded_by_fkey FOREIGN KEY (uploaded_by) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: events events_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: businesses fk_business_location; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT fk_business_location FOREIGN KEY (location_id) REFERENCES public.locations(id);


--
-- Name: documents fk_document_media; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.documents
    ADD CONSTRAINT fk_document_media FOREIGN KEY (media_id) REFERENCES public.media(id);


--
-- Name: events fk_event_categories; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT fk_event_categories FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: galleries fk_gallery_media; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.galleries
    ADD CONSTRAINT fk_gallery_media FOREIGN KEY (media_id) REFERENCES public.media(id);


--
-- Name: potentials fk_potential_location; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.potentials
    ADD CONSTRAINT fk_potential_location FOREIGN KEY (location_id) REFERENCES public.locations(id);


--
-- Name: locations fk_potential_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT fk_potential_user FOREIGN KEY (created_by_id) REFERENCES public.users(id);


--
-- Name: potentials fk_potentials_media; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.potentials
    ADD CONSTRAINT fk_potentials_media FOREIGN KEY (media_id) REFERENCES public.media(id);


--
-- Name: galleries galleries_created_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.galleries
    ADD CONSTRAINT galleries_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: media media_uploaded_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.media
    ADD CONSTRAINT media_uploaded_by_fkey FOREIGN KEY (uploaded_by) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: officials officials_village_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.officials
    ADD CONSTRAINT officials_village_id_fkey FOREIGN KEY (village_id) REFERENCES public.villages(id) ON DELETE CASCADE;


--
-- Name: posts posts_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: posts posts_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL;


--
-- Name: user_session_log user_session_log_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_session_log
    ADD CONSTRAINT user_session_log_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict KBfONGEFv8LCUUzzCJOAHliGmZLJURbJyTPgReYTdHhtmentDybTYN8gzsKZ2Sq

