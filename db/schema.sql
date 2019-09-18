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

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: properties; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.properties (
    id bigint NOT NULL,
    property_id text NOT NULL,
    street_number text NOT NULL,
    street text NOT NULL,
    street_suffix text NOT NULL,
    city text NOT NULL,
    state text NOT NULL,
    zip text NOT NULL,
    bedrooms real,
    bathrooms real,
    price text NOT NULL,
    real_estate_type text NOT NULL,
    property_type text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: properties_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.properties_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: properties_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.properties_id_seq OWNED BY public.properties.id;


--
-- Name: property_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.property_types (
    id bigint NOT NULL,
    property_type text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: property_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.property_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: property_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.property_types_id_seq OWNED BY public.property_types.id;


--
-- Name: real_estate_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.real_estate_types (
    id bigint NOT NULL,
    real_estate_type text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: real_estate_types_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.real_estate_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: real_estate_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.real_estate_types_id_seq OWNED BY public.real_estate_types.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: properties id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties ALTER COLUMN id SET DEFAULT nextval('public.properties_id_seq'::regclass);


--
-- Name: property_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.property_types ALTER COLUMN id SET DEFAULT nextval('public.property_types_id_seq'::regclass);


--
-- Name: real_estate_types id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.real_estate_types ALTER COLUMN id SET DEFAULT nextval('public.real_estate_types_id_seq'::regclass);


--
-- Name: properties properties_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_pk PRIMARY KEY (id);


--
-- Name: property_types property_types_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.property_types
    ADD CONSTRAINT property_types_pk PRIMARY KEY (id);


--
-- Name: real_estate_types real_estate_types_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.real_estate_types
    ADD CONSTRAINT real_estate_types_pk PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: properties_property_type_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX properties_property_type_uindex ON public.properties USING btree (property_type);


--
-- Name: properties_real_estate_type_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX properties_real_estate_type_uindex ON public.properties USING btree (real_estate_type);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20190918015245'),
    ('20190918017300'),
    ('20190918020349');
