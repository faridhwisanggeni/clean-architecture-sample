# clean-architecture-sample

This is sample for my student learning basic of SOLID design principle with clean code.


Here's the DDL of the table product


`-- DROP TABLE public.product;

CREATE TABLE public.product (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar(100) NULL,
	price numeric NULL,
	quantity int4 NULL,
	CONSTRAINT product_pk PRIMARY KEY (id)
);
CREATE INDEX product_id_idx ON public.product USING btree (id, name);`