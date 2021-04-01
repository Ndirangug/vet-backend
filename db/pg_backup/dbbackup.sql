--
-- PostgreSQL database dump
--

-- Dumped from database version 12.6 (Ubuntu 12.6-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.2 (Ubuntu 13.2-1.pgdg20.04+1)

-- Started on 2021-04-01 18:47:26 EAT

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 26896)
-- Name: postgis; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;


--
-- TOC entry 3891 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION postgis; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION postgis IS 'PostGIS geometry and geography spatial types and functions';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 204 (class 1259 OID 26791)
-- Name: farmers; Type: TABLE; Schema: public; Owner: vets_backend
--

CREATE TABLE public.farmers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text,
    last_name text,
    email text,
    phone text,
    address text,
    latitude numeric,
    longitude numeric
);


ALTER TABLE public.farmers OWNER TO vets_backend;

--
-- TOC entry 203 (class 1259 OID 26789)
-- Name: farmers_id_seq; Type: SEQUENCE; Schema: public; Owner: vets_backend
--

CREATE SEQUENCE public.farmers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.farmers_id_seq OWNER TO vets_backend;

--
-- TOC entry 3892 (class 0 OID 0)
-- Dependencies: 203
-- Name: farmers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vets_backend
--

ALTER SEQUENCE public.farmers_id_seq OWNED BY public.farmers.id;


--
-- TOC entry 208 (class 1259 OID 26823)
-- Name: services; Type: TABLE; Schema: public; Owner: vets_backend
--

CREATE TABLE public.services (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    description text
);


ALTER TABLE public.services OWNER TO vets_backend;

--
-- TOC entry 207 (class 1259 OID 26821)
-- Name: services_id_seq; Type: SEQUENCE; Schema: public; Owner: vets_backend
--

CREATE SEQUENCE public.services_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.services_id_seq OWNER TO vets_backend;

--
-- TOC entry 3893 (class 0 OID 0)
-- Dependencies: 207
-- Name: services_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vets_backend
--

ALTER SEQUENCE public.services_id_seq OWNED BY public.services.id;


--
-- TOC entry 212 (class 1259 OID 26857)
-- Name: sessions; Type: TABLE; Schema: public; Owner: vets_backend
--

CREATE TABLE public.sessions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    date timestamp with time zone,
    latitude numeric,
    longitude numeric,
    farmer_id bigint,
    veterinary_id bigint
);


ALTER TABLE public.sessions OWNER TO vets_backend;

--
-- TOC entry 211 (class 1259 OID 26855)
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: vets_backend
--

CREATE SEQUENCE public.sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sessions_id_seq OWNER TO vets_backend;

--
-- TOC entry 3894 (class 0 OID 0)
-- Dependencies: 211
-- Name: sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vets_backend
--

ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;


--
-- TOC entry 214 (class 1259 OID 26879)
-- Name: vet_service_sessions; Type: TABLE; Schema: public; Owner: vets_backend
--

CREATE TABLE public.vet_service_sessions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    vet_service_id bigint,
    session_id bigint,
    units bigint
);


ALTER TABLE public.vet_service_sessions OWNER TO vets_backend;

--
-- TOC entry 213 (class 1259 OID 26877)
-- Name: vet_service_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: vets_backend
--

CREATE SEQUENCE public.vet_service_sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vet_service_sessions_id_seq OWNER TO vets_backend;

--
-- TOC entry 3895 (class 0 OID 0)
-- Dependencies: 213
-- Name: vet_service_sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vets_backend
--

ALTER SEQUENCE public.vet_service_sessions_id_seq OWNED BY public.vet_service_sessions.id;


--
-- TOC entry 210 (class 1259 OID 26835)
-- Name: vet_services; Type: TABLE; Schema: public; Owner: vets_backend
--

CREATE TABLE public.vet_services (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    veterinary_id bigint,
    service_id bigint,
    unit text,
    charge_per_unit numeric
);


ALTER TABLE public.vet_services OWNER TO vets_backend;

--
-- TOC entry 209 (class 1259 OID 26833)
-- Name: vet_services_id_seq; Type: SEQUENCE; Schema: public; Owner: vets_backend
--

CREATE SEQUENCE public.vet_services_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vet_services_id_seq OWNER TO vets_backend;

--
-- TOC entry 3896 (class 0 OID 0)
-- Dependencies: 209
-- Name: vet_services_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vets_backend
--

ALTER SEQUENCE public.vet_services_id_seq OWNED BY public.vet_services.id;


--
-- TOC entry 206 (class 1259 OID 26807)
-- Name: veterinaries; Type: TABLE; Schema: public; Owner: vets_backend
--

CREATE TABLE public.veterinaries (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    first_name text,
    last_name text,
    summary text,
    email text,
    phone text,
    address text,
    latitude numeric,
    longitude numeric
);


ALTER TABLE public.veterinaries OWNER TO vets_backend;

--
-- TOC entry 205 (class 1259 OID 26805)
-- Name: veterinaries_id_seq; Type: SEQUENCE; Schema: public; Owner: vets_backend
--

CREATE SEQUENCE public.veterinaries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.veterinaries_id_seq OWNER TO vets_backend;

--
-- TOC entry 3897 (class 0 OID 0)
-- Dependencies: 205
-- Name: veterinaries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vets_backend
--

ALTER SEQUENCE public.veterinaries_id_seq OWNED BY public.veterinaries.id;


--
-- TOC entry 3702 (class 2604 OID 26794)
-- Name: farmers id; Type: DEFAULT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.farmers ALTER COLUMN id SET DEFAULT nextval('public.farmers_id_seq'::regclass);


--
-- TOC entry 3704 (class 2604 OID 26826)
-- Name: services id; Type: DEFAULT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.services ALTER COLUMN id SET DEFAULT nextval('public.services_id_seq'::regclass);


--
-- TOC entry 3706 (class 2604 OID 26860)
-- Name: sessions id; Type: DEFAULT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);


--
-- TOC entry 3707 (class 2604 OID 26882)
-- Name: vet_service_sessions id; Type: DEFAULT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_service_sessions ALTER COLUMN id SET DEFAULT nextval('public.vet_service_sessions_id_seq'::regclass);


--
-- TOC entry 3705 (class 2604 OID 26838)
-- Name: vet_services id; Type: DEFAULT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_services ALTER COLUMN id SET DEFAULT nextval('public.vet_services_id_seq'::regclass);


--
-- TOC entry 3703 (class 2604 OID 26810)
-- Name: veterinaries id; Type: DEFAULT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.veterinaries ALTER COLUMN id SET DEFAULT nextval('public.veterinaries_id_seq'::regclass);


--
-- TOC entry 3875 (class 0 OID 26791)
-- Dependencies: 204
-- Data for Name: farmers; Type: TABLE DATA; Schema: public; Owner: vets_backend
--

COPY public.farmers (id, created_at, updated_at, deleted_at, first_name, last_name, email, phone, address, latitude, longitude) FROM stdin;
8	2021-03-31 11:56:33.036017+03	2021-04-01 17:30:12.716532+03	\N	George	Ndirangu	ndirangu.mepawa@gmail.com		nairobi	-23.87876	36.678678676
\.


--
-- TOC entry 3879 (class 0 OID 26823)
-- Dependencies: 208
-- Data for Name: services; Type: TABLE DATA; Schema: public; Owner: vets_backend
--

COPY public.services (id, created_at, updated_at, deleted_at, name, description) FROM stdin;
1	\N	\N	\N	Artificial Insemination	By Injection
2	\N	\N	\N	TB Vaccination	Injection
\.


--
-- TOC entry 3883 (class 0 OID 26857)
-- Dependencies: 212
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: vets_backend
--

COPY public.sessions (id, created_at, updated_at, deleted_at, date, latitude, longitude, farmer_id, veterinary_id) FROM stdin;
1	2021-04-01 18:43:20.616624+03	2021-04-01 18:43:20.616624+03	\N	2021-04-01 18:43:20.597809+03	-23.87876	36.678678676	8	1
2	2021-04-01 18:45:19.239025+03	2021-04-01 18:45:19.239025+03	\N	2021-04-01 18:45:19.23883+03	-23.87876	36.678678676	8	1
\.


--
-- TOC entry 3701 (class 0 OID 27203)
-- Dependencies: 216
-- Data for Name: spatial_ref_sys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
\.


--
-- TOC entry 3885 (class 0 OID 26879)
-- Dependencies: 214
-- Data for Name: vet_service_sessions; Type: TABLE DATA; Schema: public; Owner: vets_backend
--

COPY public.vet_service_sessions (id, created_at, updated_at, deleted_at, vet_service_id, session_id, units) FROM stdin;
1	2021-04-01 18:43:20.669702+03	2021-04-01 18:43:20.669702+03	\N	1	1	10
2	2021-04-01 18:43:20.669702+03	2021-04-01 18:43:20.669702+03	\N	1	1	20
3	2021-04-01 18:45:19.239877+03	2021-04-01 18:45:19.239877+03	\N	1	2	10
4	2021-04-01 18:45:19.239877+03	2021-04-01 18:45:19.239877+03	\N	2	2	20
\.


--
-- TOC entry 3881 (class 0 OID 26835)
-- Dependencies: 210
-- Data for Name: vet_services; Type: TABLE DATA; Schema: public; Owner: vets_backend
--

COPY public.vet_services (id, created_at, updated_at, deleted_at, veterinary_id, service_id, unit, charge_per_unit) FROM stdin;
1	\N	\N	\N	2	1	cow	200
2	\N	\N	\N	2	2	cow	500
\.


--
-- TOC entry 3877 (class 0 OID 26807)
-- Dependencies: 206
-- Data for Name: veterinaries; Type: TABLE DATA; Schema: public; Owner: vets_backend
--

COPY public.veterinaries (id, created_at, updated_at, deleted_at, title, first_name, last_name, summary, email, phone, address, latitude, longitude) FROM stdin;
2	\N	\N	\N	Mr	George	Ndirangu	\N	\N	\N	\N	\N	\N
3	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
1	\N	\N	\N	Mr	George	\N	\N	\N	\N	\N	\N	\N
4	\N	\N	\N	\N	George	\N	\N	\N	\N	\N	-1.291804	36.822632
5	\N	\N	\N	\N	George	\N	\N	\N	\N	\N	-2.291085	36.821784
6	\N	\N	\N	\N	George	\N	\N	\N	\N	\N	-1.2939599460360305	36.799532813964184
\.


--
-- TOC entry 3898 (class 0 OID 0)
-- Dependencies: 203
-- Name: farmers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vets_backend
--

SELECT pg_catalog.setval('public.farmers_id_seq', 1, false);


--
-- TOC entry 3899 (class 0 OID 0)
-- Dependencies: 207
-- Name: services_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vets_backend
--

SELECT pg_catalog.setval('public.services_id_seq', 1, false);


--
-- TOC entry 3900 (class 0 OID 0)
-- Dependencies: 211
-- Name: sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vets_backend
--

SELECT pg_catalog.setval('public.sessions_id_seq', 2, true);


--
-- TOC entry 3901 (class 0 OID 0)
-- Dependencies: 213
-- Name: vet_service_sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vets_backend
--

SELECT pg_catalog.setval('public.vet_service_sessions_id_seq', 4, true);


--
-- TOC entry 3902 (class 0 OID 0)
-- Dependencies: 209
-- Name: vet_services_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vets_backend
--

SELECT pg_catalog.setval('public.vet_services_id_seq', 1, false);


--
-- TOC entry 3903 (class 0 OID 0)
-- Dependencies: 205
-- Name: veterinaries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: vets_backend
--

SELECT pg_catalog.setval('public.veterinaries_id_seq', 1, false);


--
-- TOC entry 3710 (class 2606 OID 26801)
-- Name: farmers farmers_email_key; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.farmers
    ADD CONSTRAINT farmers_email_key UNIQUE (email);


--
-- TOC entry 3712 (class 2606 OID 26803)
-- Name: farmers farmers_phone_key; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.farmers
    ADD CONSTRAINT farmers_phone_key UNIQUE (phone);


--
-- TOC entry 3714 (class 2606 OID 26799)
-- Name: farmers farmers_pkey; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.farmers
    ADD CONSTRAINT farmers_pkey PRIMARY KEY (id);


--
-- TOC entry 3725 (class 2606 OID 26831)
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- TOC entry 3731 (class 2606 OID 26865)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3734 (class 2606 OID 26884)
-- Name: vet_service_sessions vet_service_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_service_sessions
    ADD CONSTRAINT vet_service_sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3728 (class 2606 OID 26843)
-- Name: vet_services vet_services_pkey; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_services
    ADD CONSTRAINT vet_services_pkey PRIMARY KEY (id);


--
-- TOC entry 3718 (class 2606 OID 26817)
-- Name: veterinaries veterinaries_email_key; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.veterinaries
    ADD CONSTRAINT veterinaries_email_key UNIQUE (email);


--
-- TOC entry 3720 (class 2606 OID 26819)
-- Name: veterinaries veterinaries_phone_key; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.veterinaries
    ADD CONSTRAINT veterinaries_phone_key UNIQUE (phone);


--
-- TOC entry 3722 (class 2606 OID 26815)
-- Name: veterinaries veterinaries_pkey; Type: CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.veterinaries
    ADD CONSTRAINT veterinaries_pkey PRIMARY KEY (id);


--
-- TOC entry 3715 (class 1259 OID 26804)
-- Name: idx_farmers_deleted_at; Type: INDEX; Schema: public; Owner: vets_backend
--

CREATE INDEX idx_farmers_deleted_at ON public.farmers USING btree (deleted_at);


--
-- TOC entry 3723 (class 1259 OID 26832)
-- Name: idx_services_deleted_at; Type: INDEX; Schema: public; Owner: vets_backend
--

CREATE INDEX idx_services_deleted_at ON public.services USING btree (deleted_at);


--
-- TOC entry 3729 (class 1259 OID 26876)
-- Name: idx_sessions_deleted_at; Type: INDEX; Schema: public; Owner: vets_backend
--

CREATE INDEX idx_sessions_deleted_at ON public.sessions USING btree (deleted_at);


--
-- TOC entry 3732 (class 1259 OID 26895)
-- Name: idx_vet_service_sessions_deleted_at; Type: INDEX; Schema: public; Owner: vets_backend
--

CREATE INDEX idx_vet_service_sessions_deleted_at ON public.vet_service_sessions USING btree (deleted_at);


--
-- TOC entry 3726 (class 1259 OID 26854)
-- Name: idx_vet_services_deleted_at; Type: INDEX; Schema: public; Owner: vets_backend
--

CREATE INDEX idx_vet_services_deleted_at ON public.vet_services USING btree (deleted_at);


--
-- TOC entry 3716 (class 1259 OID 26820)
-- Name: idx_veterinaries_deleted_at; Type: INDEX; Schema: public; Owner: vets_backend
--

CREATE INDEX idx_veterinaries_deleted_at ON public.veterinaries USING btree (deleted_at);


--
-- TOC entry 3739 (class 2606 OID 26866)
-- Name: sessions fk_farmers_sessions; Type: FK CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT fk_farmers_sessions FOREIGN KEY (farmer_id) REFERENCES public.farmers(id);


--
-- TOC entry 3738 (class 2606 OID 26849)
-- Name: vet_services fk_services_vet_services; Type: FK CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_services
    ADD CONSTRAINT fk_services_vet_services FOREIGN KEY (service_id) REFERENCES public.services(id);


--
-- TOC entry 3741 (class 2606 OID 26885)
-- Name: vet_service_sessions fk_sessions_vet_service_sessions; Type: FK CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_service_sessions
    ADD CONSTRAINT fk_sessions_vet_service_sessions FOREIGN KEY (session_id) REFERENCES public.sessions(id);


--
-- TOC entry 3742 (class 2606 OID 26890)
-- Name: vet_service_sessions fk_vet_services_vet_service_sessions; Type: FK CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_service_sessions
    ADD CONSTRAINT fk_vet_services_vet_service_sessions FOREIGN KEY (vet_service_id) REFERENCES public.vet_services(id);


--
-- TOC entry 3740 (class 2606 OID 26871)
-- Name: sessions fk_veterinaries_sessions; Type: FK CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT fk_veterinaries_sessions FOREIGN KEY (veterinary_id) REFERENCES public.veterinaries(id);


--
-- TOC entry 3737 (class 2606 OID 26844)
-- Name: vet_services fk_veterinaries_vet_services; Type: FK CONSTRAINT; Schema: public; Owner: vets_backend
--

ALTER TABLE ONLY public.vet_services
    ADD CONSTRAINT fk_veterinaries_vet_services FOREIGN KEY (veterinary_id) REFERENCES public.veterinaries(id);


-- Completed on 2021-04-01 18:47:28 EAT

--
-- PostgreSQL database dump complete
--

