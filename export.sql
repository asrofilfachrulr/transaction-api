--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Ubuntu 14.2-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.2 (Ubuntu 14.2-1.pgdg20.04+1)

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

SET default_table_access_method = heap;

--
-- Name: customer_addresses; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.customer_addresses (
    id bigint NOT NULL,
    address text NOT NULL,
    customer_id bigint NOT NULL
);


ALTER TABLE public.customer_addresses OWNER TO anya;

--
-- Name: customer_addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.customer_addresses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customer_addresses_id_seq OWNER TO anya;

--
-- Name: customer_addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.customer_addresses_id_seq OWNED BY public.customer_addresses.id;


--
-- Name: customers; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.customers (
    id bigint NOT NULL,
    customer_name character varying(50) NOT NULL
);


ALTER TABLE public.customers OWNER TO anya;

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customers_id_seq OWNER TO anya;

--
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- Name: order_items; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.order_items (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    product_id bigint NOT NULL
);


ALTER TABLE public.order_items OWNER TO anya;

--
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_items_id_seq OWNER TO anya;

--
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- Name: order_payments; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.order_payments (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    payment_method_id bigint
);


ALTER TABLE public.order_payments OWNER TO anya;

--
-- Name: order_payments_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.order_payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_payments_id_seq OWNER TO anya;

--
-- Name: order_payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.order_payments_id_seq OWNED BY public.order_payments.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    customer_address_id bigint NOT NULL
);


ALTER TABLE public.orders OWNER TO anya;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO anya;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.payment_methods (
    id bigint NOT NULL,
    name bigint NOT NULL,
    is_active boolean NOT NULL
);


ALTER TABLE public.payment_methods OWNER TO anya;

--
-- Name: payment_methods_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.payment_methods_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payment_methods_id_seq OWNER TO anya;

--
-- Name: payment_methods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.payment_methods_id_seq OWNED BY public.payment_methods.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: anya
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    name bigint NOT NULL,
    price numeric NOT NULL
);


ALTER TABLE public.products OWNER TO anya;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: anya
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO anya;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: anya
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: customer_addresses id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.customer_addresses ALTER COLUMN id SET DEFAULT nextval('public.customer_addresses_id_seq'::regclass);


--
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- Name: order_payments id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_payments ALTER COLUMN id SET DEFAULT nextval('public.order_payments_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_methods_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Data for Name: customer_addresses; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.customer_addresses (id, address, customer_id) FROM stdin;
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.customers (id, customer_name) FROM stdin;
\.


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.order_items (id, order_id, product_id) FROM stdin;
\.


--
-- Data for Name: order_payments; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.order_payments (id, order_id, payment_method_id) FROM stdin;
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.orders (id, customer_address_id) FROM stdin;
\.


--
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.payment_methods (id, name, is_active) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: anya
--

COPY public.products (id, name, price) FROM stdin;
\.


--
-- Name: customer_addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.customer_addresses_id_seq', 1, false);


--
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.customers_id_seq', 1, false);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.order_items_id_seq', 1, false);


--
-- Name: order_payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.order_payments_id_seq', 1, false);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.orders_id_seq', 1, false);


--
-- Name: payment_methods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.payment_methods_id_seq', 1, false);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: anya
--

SELECT pg_catalog.setval('public.products_id_seq', 1, false);


--
-- Name: customer_addresses customer_addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.customer_addresses
    ADD CONSTRAINT customer_addresses_pkey PRIMARY KEY (id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- Name: order_payments order_payments_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_payments
    ADD CONSTRAINT order_payments_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: orders fk_customer_addresses_orders; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_customer_addresses_orders FOREIGN KEY (customer_address_id) REFERENCES public.customer_addresses(id) ON DELETE CASCADE;


--
-- Name: customer_addresses fk_customers_customer_addresses; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.customer_addresses
    ADD CONSTRAINT fk_customers_customer_addresses FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON DELETE CASCADE;


--
-- Name: order_items fk_orders_order_items; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_orders_order_items FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: order_payments fk_orders_order_payments; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_payments
    ADD CONSTRAINT fk_orders_order_payments FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: order_payments fk_payment_methods_order_payments; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_payments
    ADD CONSTRAINT fk_payment_methods_order_payments FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id) ON DELETE SET NULL;


--
-- Name: order_items fk_products_order_items; Type: FK CONSTRAINT; Schema: public; Owner: anya
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_products_order_items FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

