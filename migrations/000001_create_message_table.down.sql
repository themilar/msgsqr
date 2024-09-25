CREATE TABLE IF NOT EXISTS public.message
(
    id integer NOT NULL DEFAULT nextval('message_id_seq'::regclass),
    title character varying(55) COLLATE pg_catalog."default" NOT NULL,
    content text COLLATE pg_catalog."default" NOT NULL,
    created time with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT message_pkey PRIMARY KEY (id)
)
