CREATE TABLE "restchat".messages (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	userid int4 NOT NULL,
	"text" varchar NOT NULL,
	message_time timestamp NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX messages_id_idx ON "restchat".messages USING btree (id);

-- Permissions

ALTER TABLE "restchat".messages OWNER TO "restchat";
GRANT ALL ON TABLE "restchat".messages TO "restchat";