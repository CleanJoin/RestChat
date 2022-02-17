CREATE TABLE "db".messages (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	userid int4 NOT NULL,
	"text" varchar NOT NULL,
	message_time timestamp NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX messages_id_idx ON "db".messages USING btree (id);

-- Permissions

ALTER TABLE "db".messages OWNER TO "restchat";
GRANT ALL ON TABLE "db".messages TO "restchat";