CREATE SCHEMA "restchat" AUTHORIZATION "restchat";

CREATE TABLE "restchat".messages (
	id integer NOT NULL,
	"text" varchar NOT NULL,
	user_id integer NOT NULL,
	"time" timestamptz NOT NULL
);

CREATE UNIQUE INDEX messages_user_id_idx ON "restchat".messages USING btree (user_id);

ALTER TABLE "restchat".messages OWNER TO "restchat";

GRANT ALL ON TABLE "restchat".messages TO "restchat";

CREATE TABLE "reschat".users (
	id integer NOT NULL,
	username varchar NOT NULL,
	passwordhash varchar NOT NULL
);

CREATE INDEX users_id_idx ON "restchat".users USING btree (id);

ALTER TABLE "restchat".users OWNER TO "restchat";

GRANT ALL ON TABLE "restchat".users TO "restchat";

GRANT ALL ON SCHEMA "restchat" TO "restchat";

INSERT INTO "restchat".users (username,passwordhash,id) VALUES
	 ('Andrey','ertdfghjuehjuik345535',1);

INSERT INTO "restchat".messages (id,"text",user_id,"time") VALUES
	 (1,'Привет всем',1,'2022-12-02 15:57:00+03');