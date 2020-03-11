drop table activity cascade;
drop table users;

CREATE TABLE users (
    user_id         integer PRIMARY KEY NOT NULL,
    user_login      text NOT NULL,
    track           BOOLEAN NOT NULL,
    created_at      timestamp not null
);

insert into users (user_id,user_login,track, created_at) values (653013,'alexellisuk',true,now());
insert into users (user_id,user_login,track, created_at) values (103022,'rgee0',true,now());

CREATE TABLE activity (
    id              INT GENERATED ALWAYS AS IDENTITY,
    user_id         integer NOT NULL references users(user_id),
    activity_date   timestamp NOT NULL,
    emoji   text NOT NULL
);

insert into activity (user_id, activity_date, emoji) values (653013, now(), '👍');

drop function get_emojis;

CREATE or REPLACE FUNCTION get_emojis()
    RETURNS TABLE(emoji text, total bigint)
  AS
$$
BEGIN
RETURN QUERY select
    a.emoji,
    count(a.emoji) as total
from activity a
group by a.emoji
order by total desc;
END
$$  LANGUAGE 'plpgsql' VOLATILE;


select * from get_emojis();

