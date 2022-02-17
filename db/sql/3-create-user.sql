CREATE TABLE users (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	username varchar NOT NULL,
	"password" varchar NOT NULL
);
CREATE UNIQUE INDEX users_id_idx ON "UserModel".users USING btree (id);

-- Permissions

ALTER TABLE "UserModel".users OWNER TO "restChat";
GRANT ALL ON TABLE "UserModel".users TO "restChat";

CREATE SEQUENCE "UserModel".users_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;

-- Permissions

ALTER SEQUENCE "UserModel".users_id_seq OWNER TO "restChat";
GRANT ALL ON SEQUENCE "UserModel".users_id_seq TO "restChat";