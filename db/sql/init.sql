CREATE SCHEMA "UserModel" AUTHORIZATION "restchat";

CREATE TABLE "UserModel".messages (
	id int4 NOT NULL,
	"text" varchar NOT NULL,
	user_id int4 NOT NULL,
	"time" timestamptz NOT NULL
);
CREATE UNIQUE INDEX messages_user_id_idx ON "UserModel".messages USING btree (user_id);

ALTER TABLE "UserModel".messages OWNER TO "restchat";
GRANT ALL ON TABLE "UserModel".messages TO "restchat";

CREATE TABLE "UserModel".users (
	username varchar NOT NULL,
	passwordhash varchar NOT NULL,
	id int4 NOT NULL
);
CREATE INDEX users_id_idx ON "UserModel".users USING btree (id);

ALTER TABLE "UserModel".users OWNER TO "restchat";
GRANT ALL ON TABLE "UserModel".users TO "restchat";

GRANT ALL ON SCHEMA "UserModel" TO "restchat";

INSERT INTO "UserModel".users (username,passwordhash,id) VALUES
	 ('Andrey','ertdfghjuehjuik345535',1);
INSERT INTO "UserModel".messages (id,"text",user_id,"time") VALUES
	 (1,'Привет всем',1,'2022-12-02 15:57:00+03');