-- Технические sequence для таблиц, где DDL уже ссылается на nextval(...).
-- Добавлены, потому что в экспортированном ddlTables.sql они не были выгружены отдельными CREATE SEQUENCE.
CREATE SEQUENCE IF NOT EXISTS public.album_owners_id_seq AS integer;
CREATE SEQUENCE IF NOT EXISTS public.playlist_owners_id_seq AS integer;
CREATE SEQUENCE IF NOT EXISTS public.track_owners_id_seq AS integer;

-- public.albums определение

-- Drop table

-- DROP TABLE public.albums;

CREATE TABLE public.albums (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	title varchar(255) NOT NULL,
	description text NULL,
	cover_url text NULL,
	release_date date NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT albums_pkey PRIMARY KEY (id)
);


-- public.artists определение

-- Drop table

-- DROP TABLE public.artists;

CREATE TABLE public.artists (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(255) NOT NULL,
	description text NULL,
	avatar_url text NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT content_owners_pkey PRIMARY KEY (id)
);


-- public.genres определение

-- Drop table

-- DROP TABLE public.genres;

CREATE TABLE public.genres (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(100) NOT NULL,
	CONSTRAINT genres_name_key UNIQUE (name),
	CONSTRAINT genres_pkey PRIMARY KEY (id)
);



-- public.track_statuses определение

-- Drop table

-- DROP TABLE public.track_statuses;

CREATE TABLE public.track_statuses (
	id serial4 NOT NULL,
	code varchar(30) NOT NULL,
	"name" varchar(100) NOT NULL,
	CONSTRAINT track_statuses_code_key UNIQUE (code),
	CONSTRAINT track_statuses_pkey PRIMARY KEY (id)
);


-- public.users определение

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	email text NULL,
	"name" text NOT NULL,
	hash_password text NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);


-- public.album_artists определение

-- Drop table

-- DROP TABLE public.album_artists;

CREATE TABLE public.album_artists (
	id int4 DEFAULT nextval('album_owners_id_seq'::regclass) NOT NULL,
	album_id uuid NOT NULL,
	artist_id uuid NOT NULL,
	"role" varchar(50) DEFAULT 'artist'::character varying NOT NULL,
	CONSTRAINT album_owners_album_id_owner_id_role_key UNIQUE (album_id, artist_id, role),
	CONSTRAINT album_owners_pkey PRIMARY KEY (id),
	CONSTRAINT album_owners_album_id_fkey FOREIGN KEY (album_id) REFERENCES public.albums(id) ON DELETE CASCADE,
	CONSTRAINT album_owners_owner_id_fkey FOREIGN KEY (artist_id) REFERENCES public.artists(id) ON DELETE CASCADE
);


-- public.oauth_accounts определение

-- Drop table

-- DROP TABLE public.oauth_accounts;

CREATE TABLE public.oauth_accounts (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	user_id uuid NOT NULL,
	provider varchar(50) NOT NULL,
	provider_user_id varchar(255) NOT NULL,
	provider_email varchar(255) NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT oauth_accounts_pkey PRIMARY KEY (id),
	CONSTRAINT oauth_accounts_provider_provider_user_id_key UNIQUE (provider, provider_user_id),
	CONSTRAINT oauth_accounts_user_id_provider_key UNIQUE (user_id, provider),
	CONSTRAINT oauth_accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.playlists определение

-- Drop table

-- DROP TABLE public.playlists;

CREATE TABLE public.playlists (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	user_id uuid NULL,
	title varchar(255) NOT NULL,
	description text NULL,
	cover_url text NULL,
	visibility varchar(30) DEFAULT 'public'::character varying NOT NULL,
	share_token varchar(100) NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT playlists_pkey PRIMARY KEY (id),
	CONSTRAINT playlists_share_token_key UNIQUE (share_token),
	CONSTRAINT playlists_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.refresh_tokens определение

-- Drop table

-- DROP TABLE public.refresh_tokens;

CREATE TABLE public.refresh_tokens (
	id uuid NOT NULL,
	user_id uuid NOT NULL,
	token_hash text NOT NULL,
	expires_at timestamp NOT NULL,
	revoked_at timestamp NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id),
	CONSTRAINT refresh_tokens_token_hash_key UNIQUE (token_hash),
	CONSTRAINT refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.tracks определение

-- Drop table

-- DROP TABLE public.tracks;

CREATE TABLE public.tracks (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	album_id uuid NULL,
	title varchar(255) NOT NULL,
	description text NULL,
	cover_url text NULL,
	duration_seconds int4 NOT NULL,
	release_date date NULL,
	is_explicit bool DEFAULT false NOT NULL,
	is_streamable bool DEFAULT true NOT NULL,
	is_downloadable bool DEFAULT false NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	audio_storage_key text NULL,
	status_id int4 NOT NULL,
	CONSTRAINT tracks_duration_seconds_check CHECK ((duration_seconds > 0)),
	CONSTRAINT tracks_pkey PRIMARY KEY (id),
	CONSTRAINT fk_tracks_status FOREIGN KEY (status_id) REFERENCES public.track_statuses(id) ON DELETE RESTRICT,
	CONSTRAINT tracks_album_id_fkey FOREIGN KEY (album_id) REFERENCES public.albums(id) ON DELETE SET NULL
);


-- public.user_liked_tracks определение

-- Drop table

-- DROP TABLE public.user_liked_tracks;

CREATE TABLE public.user_liked_tracks (
	user_id uuid NOT NULL,
	track_id uuid NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT user_liked_tracks_pkey PRIMARY KEY (user_id, track_id),
	CONSTRAINT user_liked_tracks_track_id_fkey FOREIGN KEY (track_id) REFERENCES public.tracks(id) ON DELETE CASCADE,
	CONSTRAINT user_liked_tracks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);


-- public.playlist_artists определение

-- Drop table

-- DROP TABLE public.playlist_artists;

CREATE TABLE public.playlist_artists (
	id int4 DEFAULT nextval('playlist_owners_id_seq'::regclass) NOT NULL,
	playlist_id uuid NOT NULL,
	artist_id uuid NOT NULL,
	CONSTRAINT playlist_owners_pkey PRIMARY KEY (id),
	CONSTRAINT playlist_owners_owner_id_fkey FOREIGN KEY (artist_id) REFERENCES public.artists(id) ON DELETE CASCADE,
	CONSTRAINT playlist_owners_playlist_id_fkey FOREIGN KEY (playlist_id) REFERENCES public.playlists(id) ON DELETE CASCADE
);


-- public.playlist_tracks определение

-- Drop table

-- DROP TABLE public.playlist_tracks;

CREATE TABLE public.playlist_tracks (
	id serial4 NOT NULL,
	playlist_id uuid NOT NULL,
	track_id uuid NOT NULL,
	"position" int4 DEFAULT 0 NOT NULL,
	added_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT playlist_tracks_pkey PRIMARY KEY (id),
	CONSTRAINT playlist_tracks_playlist_id_track_id_key UNIQUE (playlist_id, track_id),
	CONSTRAINT playlist_tracks_playlist_id_fkey FOREIGN KEY (playlist_id) REFERENCES public.playlists(id) ON DELETE CASCADE,
	CONSTRAINT playlist_tracks_track_id_fkey FOREIGN KEY (track_id) REFERENCES public.tracks(id) ON DELETE CASCADE
);


-- public.track_artists определение

-- Drop table

-- DROP TABLE public.track_artists;

CREATE TABLE public.track_artists (
	id int4 DEFAULT nextval('track_owners_id_seq'::regclass) NOT NULL,
	track_id uuid NOT NULL,
	artist_id uuid NOT NULL,
	"role" varchar(50) DEFAULT 'artist'::character varying NOT NULL,
	CONSTRAINT track_owners_pkey PRIMARY KEY (id),
	CONSTRAINT track_owners_track_id_owner_id_role_key UNIQUE (track_id, artist_id, role),
	CONSTRAINT track_owners_owner_id_fkey FOREIGN KEY (artist_id) REFERENCES public.artists(id) ON DELETE CASCADE,
	CONSTRAINT track_owners_track_id_fkey FOREIGN KEY (track_id) REFERENCES public.tracks(id) ON DELETE CASCADE
);


-- public.track_genres определение

-- Drop table

-- DROP TABLE public.track_genres;

CREATE TABLE public.track_genres (
	id serial4 NOT NULL,
	track_id uuid NOT NULL,
	genre_id uuid NOT NULL,
	CONSTRAINT track_genres_pkey PRIMARY KEY (id),
	CONSTRAINT track_genres_track_id_genre_id_key UNIQUE (track_id, genre_id),
	CONSTRAINT track_genres_genre_id_fkey FOREIGN KEY (genre_id) REFERENCES public.genres(id) ON DELETE RESTRICT,
	CONSTRAINT track_genres_track_id_fkey FOREIGN KEY (track_id) REFERENCES public.tracks(id) ON DELETE CASCADE
);
