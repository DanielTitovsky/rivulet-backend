-- Откат данных, загруженных миграцией 000002_seed_data.up.sql.
DELETE FROM public.track_genres;
DELETE FROM public.track_artists;
DELETE FROM public.playlist_tracks;
DELETE FROM public.playlist_artists;
DELETE FROM public.album_artists;
DELETE FROM public.tracks;
DELETE FROM public.playlists;
DELETE FROM public.refresh_tokens;
DELETE FROM public.oauth_accounts;
DELETE FROM public.track_statuses;
DELETE FROM public.albums;
DELETE FROM public.genres;
DELETE FROM public.artists;
DELETE FROM public.users;

ALTER SEQUENCE IF EXISTS public.album_owners_id_seq RESTART WITH 1;
ALTER SEQUENCE IF EXISTS public.playlist_owners_id_seq RESTART WITH 1;
ALTER SEQUENCE IF EXISTS public.track_owners_id_seq RESTART WITH 1;
ALTER SEQUENCE IF EXISTS public.track_statuses_id_seq RESTART WITH 1;
ALTER SEQUENCE IF EXISTS public.playlist_tracks_id_seq RESTART WITH 1;
ALTER SEQUENCE IF EXISTS public.track_genres_id_seq RESTART WITH 1;
