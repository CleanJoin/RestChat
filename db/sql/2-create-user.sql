CREATE TABLE "db".users (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	username varchar NOT NULL,
	"password" varchar NOT NULL
);
CREATE UNIQUE INDEX users_id_idx ON "db".users USING btree (id);

-- Permissions

ALTER TABLE "db".users OWNER TO "restchat";
GRANT ALL ON TABLE "db".users TO "restchat";