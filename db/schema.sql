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
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT property_types CHECK (((property_type = 'Single-Family'::text) OR (property_type = 'Multi-family'::text) OR (property_type = 'Condo'::text) OR (property_type = 'Co-op'::text) OR (property_type = 'Townhouse'::text) OR (property_type = 'Duplex'::text) OR (property_type = 'Triple decker'::text) OR (property_type = '4-plex'::text) OR (property_type = 'High value home'::text) OR (property_type = 'Generational'::text) OR (property_type = 'Vacation home'::text) OR (property_type = 'Shopping center'::text) OR (property_type = 'Strip mall'::text) OR (property_type = 'Medical building'::text) OR (property_type = 'Educational building'::text) OR (property_type = 'Hotel'::text) OR (property_type = 'Office building'::text) OR (property_type = 'Apartment building'::text) OR (property_type = 'Manufacturing building'::text) OR (property_type = 'Warehouse'::text) OR (property_type = 'Other'::text) OR (property_type = 'Vacant land'::text) OR (property_type = 'Working farm'::text) OR (property_type = 'Ranch'::text))),
    CONSTRAINT real_estate_types CHECK (((real_estate_type = 'Residential'::text) OR (real_estate_type = 'Commercial'::text) OR (real_estate_type = 'Industrial'::text) OR (real_estate_type = 'Land'::text)))
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
-- Name: properties properties_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_pk PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20190918020349');
