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


-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS "public"."transaction";
CREATE TABLE "public"."transaction" (
  "id" "pg_catalog"."uuid" NOT NULL DEFAULT uuid_generate_v1(),
  "trx_id" "pg_catalog"."uuid" NOT NULL DEFAULT uuid_generate_v1(),
  "created_at" "pg_catalog"."timestamp" DEFAULT now(),
  "total_amount" "pg_catalog"."numeric"
)
;

-- ----------------------------
-- Indexes structure for table transaction
-- ----------------------------
CREATE INDEX "transaction_detail_idx" ON "public"."transaction" USING btree (
  "id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "trx_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "transaction_idx" ON "public"."transaction" USING btree (
  "id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "trx_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table transaction
-- ----------------------------
ALTER TABLE "public"."transaction" ADD CONSTRAINT "transaction_pk" PRIMARY KEY ("id");


-- ----------------------------
-- Table structure for transaction_detail
-- ----------------------------
DROP TABLE IF EXISTS "public"."transaction_detail";
CREATE TABLE "public"."transaction_detail" (
  "id" "pg_catalog"."uuid" NOT NULL DEFAULT uuid_generate_v1(),
  "trx_id" "pg_catalog"."varchar" COLLATE "pg_catalog"."default",
  "quantity" "pg_catalog"."int4"
)
;

-- ----------------------------
-- Primary Key structure for table transaction_detail
-- ----------------------------
ALTER TABLE "public"."transaction_detail" ADD CONSTRAINT "transaction_detail_pk" PRIMARY KEY ("id");
