CREATE TABLE "restchat".users (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	username varchar NOT NULL,
	"password" varchar NOT NULL
);
CREATE UNIQUE INDEX users_id_idx ON "restchat".users USING btree (id);

-- Permissions

ALTER TABLE "restchat".users OWNER TO "restchat";
GRANT ALL ON TABLE "restchat".users TO "restchat";