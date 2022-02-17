CREATE TABLE messages (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	userid int4 NOT NULL,
	"text" varchar NOT NULL,
	message_time timestamp NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX messages_id_idx ON "UserModel".messages USING btree (id);

-- Permissions

ALTER TABLE "UserModel".messages OWNER TO "restChat";
GRANT ALL ON TABLE "UserModel".messages TO "restChat";

CREATE SEQUENCE "UserModel".messages_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;

-- Permissions

ALTER SEQUENCE "UserModel".messages_id_seq OWNER TO "restChat";
GRANT ALL ON SEQUENCE "UserModel".messages_id_seq TO "restChat";
