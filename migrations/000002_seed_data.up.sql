-- Данные из dbExport/exportData.
-- schema_migrations намеренно не переносится: её ведёт migrate/migrate.

-- users_202606180141.sql
INSERT INTO public.users (id,email,name,hash_password,updated_at,created_at) VALUES
	 ('924ee559-5a78-4183-a940-fee9b6a8cbbf'::uuid,'DanDanDam@example.com','Dan Ti','$2a$14$YDAYPr2G2AZ.bGu3nSr1AuLuadaK7axitZxowGB8J8hu8E7DVOi/G','2026-05-22 01:31:48.830726','2026-05-22 01:31:48.830726'),
	 ('05d75892-2b0b-4abd-a944-7b2ed02a82fa'::uuid,'','','$2a$14$QctceFoFtm9dtXYEFdCGkOGvl7RzrQ9aS8zabjUJ8MUW8EXcSDERS','2026-05-22 01:32:22.350111','2026-05-22 01:32:22.350111'),
	 ('4deaca1b-d6c8-47f2-ad49-2523f6027be9'::uuid,'DantITAasdT@mail.ru','Dniiiiil23','$2a$14$mxxYEL/UJz19Cu3kcvKtoumxoOm11XV.7HB1lxJvh/8CoV7e/Q7re','2026-05-22 01:50:32.186448','2026-05-22 01:50:32.186448'),
	 ('41d98110-fb05-41ca-9fd9-627a2a81258e'::uuid,NULL,'Даниил Титовский',NULL,'2026-05-22 01:55:00.280318','2026-05-22 01:55:00.280318'),
	 ('44959a7a-564f-4915-b6ae-35d5f226614d'::uuid,'daniil.titovsk@gmail.com','asdfasdf','$2a$14$8/uEWGUT75qfB3irpCMCqO./RBMl1YcVDxv2ptd.2K8PVzuTBmzhG','2026-05-22 01:58:13.467838','2026-05-22 01:58:13.467838'),
	 ('46658189-3d01-4fa5-927e-eabb861be4c5'::uuid,NULL,'Даниил Титовский',NULL,'2026-05-22 01:58:43.351526','2026-05-22 01:58:43.351526'),
	 ('00000000-0000-0000-0000-000000000000'::uuid,'Jopa@example.com','Dan Ti','$2a$14$JdsNLt3ZPmUHgGV6hW9Xk.Mf/JZ8ZIlIFhCcS73N1HJE0J6W/D4p2','2026-06-11 00:51:00.626392','2026-06-11 00:51:00.626392'),
	 ('95577eac-c92c-4861-b632-efafd2122a32'::uuid,'dan@mail.ru','Daniil Titovsk','$2a$14$PFxVI/pGnnDF3z5HWTCLy.02n3TiL00LZCzqZlhZtw2jkCjgFJB3m','2026-06-11 01:56:03.303536','2026-06-11 01:28:40.661612');

-- artists_202606180141.sql
INSERT INTO public.artists (id,name,description,avatar_url,created_at,updated_at) VALUES
	 ('11111111-1111-1111-1111-111111111111'::uuid,'Imagine Dragons','Американская группа','/owners/imagine-dragons.jpg','2026-06-12 01:28:44.939691','2026-06-12 01:28:44.939691'),
	 ('22222222-2222-2222-2222-222222222222'::uuid,'The Weeknd','Канадский исполнитель','/owners/the-weeknd.jpg','2026-06-12 01:28:44.939691','2026-06-12 01:28:44.939691'),
	 ('33333333-3333-3333-3333-333333333333'::uuid,'Linkin Park','Американская рок-группа','/owners/linkin-park.jpg','2026-06-12 01:28:44.939691','2026-06-12 01:28:44.939691'),
	 ('44444444-4444-4444-4444-444444444444'::uuid,'Daft Punk','Французский электронный дуэт','/owners/daft-punk.jpg','2026-06-12 01:28:44.939691','2026-06-12 01:28:44.939691'),
	 ('55555555-5555-5555-5555-555555555555'::uuid,'Billie Eilish','Американская певица','/owners/billie-eilish.jpg','2026-06-12 01:28:44.939691','2026-06-12 01:28:44.939691');

-- genres_202606180141.sql
INSERT INTO public.genres (id,name) VALUES
	 ('cccccccc-cccc-cccc-cccc-ccccccccccc1'::uuid,'Rock'),
	 ('cccccccc-cccc-cccc-cccc-ccccccccccc2'::uuid,'Pop'),
	 ('cccccccc-cccc-cccc-cccc-ccccccccccc3'::uuid,'Electronic'),
	 ('cccccccc-cccc-cccc-cccc-ccccccccccc4'::uuid,'Alternative'),
	 ('cccccccc-cccc-cccc-cccc-ccccccccccc5'::uuid,'Indie');

-- albums_202606180141.sql
INSERT INTO public.albums (id,title,description,cover_url,release_date,created_at,updated_at) VALUES
	 ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1'::uuid,'Evolve','Альбом Imagine Dragons','/albums/evolve.jpg','2017-06-23','2026-06-13 00:02:12.990281','2026-06-13 00:02:12.990281'),
	 ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa2'::uuid,'After Hours','Альбом The Weeknd','/albums/after-hours.jpg','2020-03-20','2026-06-13 00:02:12.990281','2026-06-13 00:02:12.990281'),
	 ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa3'::uuid,'Meteora','Альбом Linkin Park','/albums/meteora.jpg','2003-03-25','2026-06-13 00:02:12.990281','2026-06-13 00:02:12.990281'),
	 ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa4'::uuid,'Random Access Memories','Альбом Daft Punk','/albums/ram.jpg','2013-05-17','2026-06-13 00:02:12.990281','2026-06-13 00:02:12.990281'),
	 ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa5'::uuid,'Happier Than Ever','Альбом Billie Eilish','/albums/happier-than-ever.jpg','2021-07-30','2026-06-13 00:02:12.990281','2026-06-13 00:02:12.990281');

-- track_statuses_202606180141.sql
INSERT INTO public.track_statuses (code,name) VALUES
	 ('draft','Черновик'),
	 ('published','Опубликован'),
	 ('hidden','Скрыт'),
	 ('blocked','Заблокирован'),
	 ('deleted','Удалён');

-- oauth_accounts_202606180141.sql
INSERT INTO public.oauth_accounts (id,user_id,provider,provider_user_id,provider_email,created_at) VALUES
	 ('2169ac93-5031-4643-9ee5-c6f696a40eaa'::uuid,'41d98110-fb05-41ca-9fd9-627a2a81258e'::uuid,'google','106731046982290559008','daniil.titovsk@gmail.com','2026-05-22 01:55:00.282381');

-- refresh_tokens_202606180141.sql
INSERT INTO public.refresh_tokens (id,user_id,token_hash,expires_at,revoked_at,created_at) VALUES
	 ('82741bee-fe76-4649-a9b9-573a10c6fcbb'::uuid,'924ee559-5a78-4183-a940-fee9b6a8cbbf'::uuid,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjgyNzQxYmVlLWZlNzYtNDY0OS1hOWI5LTU3M2ExMGM2ZmNiYiIsInVzZXJJZCI6IjkyNGVlNTU5LTVhNzgtNDE4My1hOTQwLWZlZTliNmE4Y2JiZiIsInVzZXJFbWFpbCI6IkRhbkRhbkRhbUBleGFtcGxlLmNvbSIsImV4cCI6MTc4MjAwNTUwOCwiaWF0IjoxNzc5NDEzNTA4fQ.Z8ioXlJGwD9D2ldXn9uItSqF0CsImkMV4-yfAigr5G8','2026-06-21 01:31:48.837769','0001-01-01 00:00:00','2026-05-22 01:31:48.837769'),
	 ('2a6ec09a-e669-4a63-94b0-a54a1a94e2c3'::uuid,'05d75892-2b0b-4abd-a944-7b2ed02a82fa'::uuid,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjJhNmVjMDlhLWU2NjktNGE2My05NGIwLWE1NGExYTk0ZTJjMyIsInVzZXJJZCI6IjA1ZDc1ODkyLTJiMGItNGFiZC1hOTQ0LTdiMmVkMDJhODJmYSIsInVzZXJFbWFpbCI6IiIsImV4cCI6MTc4MjAwNTU0MiwiaWF0IjoxNzc5NDEzNTQyfQ.2h7kwDmLHG1endwj-yrWdZYFUHADjSGKg5_yLAwgk_8','2026-06-21 01:32:22.360903','0001-01-01 00:00:00','2026-05-22 01:32:22.360903'),
	 ('21e11d10-3226-4a09-96d7-fa9eb2bed153'::uuid,'4deaca1b-d6c8-47f2-ad49-2523f6027be9'::uuid,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjIxZTExZDEwLTMyMjYtNGEwOS05NmQ3LWZhOWViMmJlZDE1MyIsInVzZXJJZCI6IjRkZWFjYTFiLWQ2YzgtNDdmMi1hZDQ5LTI1MjNmNjAyN2JlOSIsInVzZXJFbWFpbCI6IkRhbnRJVEFhc2RUQG1haWwucnUiLCJleHAiOjE3ODIwMDY2MzIsImlhdCI6MTc3OTQxNDYzMn0.wlB9_lksr87wmjdi8youW3DS4t0cbgKy2PdFA0M-iKw','2026-06-21 01:50:32.188618','0001-01-01 00:00:00','2026-05-22 01:50:32.188618'),
	 ('facc9f09-620b-4e5e-a142-d660e4336d42'::uuid,'44959a7a-564f-4915-b6ae-35d5f226614d'::uuid,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZhY2M5ZjA5LTYyMGItNGU1ZS1hMTQyLWQ2NjBlNDMzNmQ0MiIsInVzZXJJZCI6IjQ0OTU5YTdhLTU2NGYtNDkxNS1iNmFlLTM1ZDVmMjI2NjE0ZCIsInVzZXJFbWFpbCI6ImRhbmlpbC50aXRvdnNrQGdtYWlsLmNvbSIsImV4cCI6MTc4MjAwNzA5MywiaWF0IjoxNzc5NDE1MDkzfQ.C6kOgYq-2iSzLTWkcyOXJ4uglwIs5FHB4Vxy6cXvYMM','2026-06-21 01:58:13.476035','0001-01-01 00:00:00','2026-05-22 01:58:13.476035');

-- playlists_202606180141.sql
INSERT INTO public.playlists (id,user_id,title,description,cover_url,visibility,share_token,created_at,updated_at) VALUES
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd1'::uuid,NULL,'Лучшее за неделю','Популярные треки недели','/playlists/week-best.jpg','public','week-best-001','2026-06-12 01:28:44.96593','2026-06-12 01:28:44.96593'),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd2'::uuid,NULL,'Рок подборка','Рок-треки','/playlists/rock.jpg','public','rock-001','2026-06-12 01:28:44.96593','2026-06-12 01:28:44.96593'),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd3'::uuid,NULL,'Поп музыка','Популярная музыка','/playlists/pop.jpg','public','pop-001','2026-06-12 01:28:44.96593','2026-06-12 01:28:44.96593'),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd4'::uuid,NULL,'Электронная музыка','Электронные треки','/playlists/electronic.jpg','public','electronic-001','2026-06-12 01:28:44.96593','2026-06-12 01:28:44.96593'),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd5'::uuid,NULL,'Вечерний плейлист','Музыка для вечера','/playlists/evening.jpg','private','evening-001','2026-06-12 01:28:44.96593','2026-06-12 01:28:44.96593'),
	 ('11111111-aaaa-4aaa-8aaa-111111111111'::uuid,'05d75892-2b0b-4abd-a944-7b2ed02a82fa'::uuid,'Мой плейлист','Личная подборка пользователя','/playlists/my-playlist.jpg','public','share-my-playlist-001','2026-06-14 02:27:33.62025','2026-06-14 02:27:33.62025'),
	 ('22222222-aaaa-4aaa-8aaa-222222222222'::uuid,'924ee559-5a78-4183-a940-fee9b6a8cbbf'::uuid,'Рок на вечер','Подборка рок-треков','/playlists/evening-rock.jpg','public','share-rock-001','2026-06-14 02:27:33.62025','2026-06-14 02:27:33.62025'),
	 ('33333333-aaaa-4aaa-8aaa-333333333333'::uuid,'00000000-0000-0000-0000-000000000000'::uuid,'Популярное','Популярные треки','/playlists/popular.jpg','public','share-popular-001','2026-06-14 02:27:33.62025','2026-06-14 02:27:33.62025'),
	 ('44444444-aaaa-4aaa-8aaa-444444444444'::uuid,'44959a7a-564f-4915-b6ae-35d5f226614d'::uuid,'Тренировка','Музыка для тренировки','/playlists/workout.jpg','private','share-workout-001','2026-06-14 02:27:33.62025','2026-06-14 02:27:33.62025'),
	 ('55555555-aaaa-4aaa-8aaa-555555555555'::uuid,'41d98110-fb05-41ca-9fd9-627a2a81258e'::uuid,'Спокойная музыка','Музыка для отдыха','/playlists/chill.jpg','public','share-chill-001','2026-06-14 02:27:33.62025','2026-06-14 02:27:33.62025');
INSERT INTO public.playlists (id,user_id,title,description,cover_url,visibility,share_token,created_at,updated_at) VALUES
	 ('66666666-aaaa-4aaa-8aaa-666666666666'::uuid,'95577eac-c92c-4861-b632-efafd2122a32'::uuid,'Любимые треки','Подборка любимых треков','/playlists/favorites.jpg','private','share-favorites-001','2026-06-14 02:27:33.62025','2026-06-14 02:27:33.62025');

-- tracks_202606180141.sql
INSERT INTO public.tracks (id,album_id,title,description,cover_url,duration_seconds,release_date,is_explicit,is_streamable,is_downloadable,created_at,updated_at,audio_storage_key,status_id) VALUES
	 ('ddc8d79f-0e86-42a1-b658-f67325351823'::uuid,'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1'::uuid,'Thunder','Popular Imagine Dragons single','/covers/tracks/thunder.jpg',187,'2017-04-27',false,true,false,'2026-06-13 00:02:22.732706','2026-06-13 00:02:22.732706','tracks/imagine_dragons/thunder.mp3',2),
	 ('9e933a6f-d35d-4140-be9d-135bc95077cc'::uuid,'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa2'::uuid,'Blinding Lights','One of the most streamed songs','/covers/tracks/blinding_lights.jpg',200,'2019-11-29',false,true,false,'2026-06-13 00:02:22.732706','2026-06-13 00:02:22.732706','tracks/the_weeknd/blinding_lights.mp3',2),
	 ('70d1892c-47b0-4657-ad68-e56234ad49fc'::uuid,'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa4'::uuid,'Get Lucky','Daft Punk feat. Pharrell Williams','/covers/tracks/get_lucky.jpg',369,'2013-04-19',false,true,false,'2026-06-13 00:02:22.732706','2026-06-13 00:02:22.732706','tracks/daft_punk/get_lucky.mp3',2),
	 ('341ffba7-fadf-4ebc-ab23-a4ed5ee7dca6'::uuid,'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa5'::uuid,'Happier Than Ever','Title track','/covers/tracks/happier_than_ever.jpg',298,'2021-07-30',true,true,false,'2026-06-13 00:02:22.732706','2026-06-13 00:02:22.732706','tracks/billie_eilish/happier_than_ever.mp3',2),
	 ('33015630-09c7-4917-921b-27e2e77f257a'::uuid,'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa4'::uuid,'JOPA','Обновлённое описание трека','covers/tracks/believer.jpg',204,'2099-02-01',true,true,true,'2026-06-13 00:02:22.732706','2026-06-14 23:33:19.627915','tracks/imagine_dragons/popa.mp3',4),
	 ('0947ffc3-fe28-4ee1-ba84-35495989cf0d'::uuid,'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa3'::uuid,'Believer','Обновлённое описание трека','covers/tracks/believer.jpg',204,'2017-02-01',false,true,false,'2026-06-13 00:02:22.732706','2026-06-13 23:55:26.805588','songs/Caudal - Still Circuit.mp3',3),
	 ('00000000-0000-0000-0000-000000000000'::uuid,'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa3'::uuid,'Кишечный фонк','Ритмы привокзального шансона примеком из глубинки','covers/tracks/new-track.jpg',12,'2024-01-01',false,true,false,'2026-06-15 00:47:43.549013','2026-06-15 00:47:43.549013','tracks/new-track.mp3',2);

-- album_artists_202606180141.sql
-- empty export file

-- playlist_artists_202606180141.sql
INSERT INTO public.playlist_artists (playlist_id,artist_id) VALUES
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd1'::uuid,'11111111-1111-1111-1111-111111111111'::uuid),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd2'::uuid,'33333333-3333-3333-3333-333333333333'::uuid),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd3'::uuid,'22222222-2222-2222-2222-222222222222'::uuid),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd4'::uuid,'44444444-4444-4444-4444-444444444444'::uuid),
	 ('dddddddd-dddd-dddd-dddd-ddddddddddd5'::uuid,'55555555-5555-5555-5555-555555555555'::uuid);

-- playlist_tracks_202606180141.sql
INSERT INTO public.playlist_tracks (playlist_id,track_id,"position",added_at) VALUES
	 ('11111111-aaaa-4aaa-8aaa-111111111111'::uuid,'ddc8d79f-0e86-42a1-b658-f67325351823'::uuid,1,'2026-06-14 02:28:13.141192'),
	 ('11111111-aaaa-4aaa-8aaa-111111111111'::uuid,'0947ffc3-fe28-4ee1-ba84-35495989cf0d'::uuid,2,'2026-06-14 02:28:13.141192'),
	 ('22222222-aaaa-4aaa-8aaa-222222222222'::uuid,'70d1892c-47b0-4657-ad68-e56234ad49fc'::uuid,1,'2026-06-14 02:28:13.141192'),
	 ('22222222-aaaa-4aaa-8aaa-222222222222'::uuid,'341ffba7-fadf-4ebc-ab23-a4ed5ee7dca6'::uuid,2,'2026-06-14 02:28:13.141192'),
	 ('33333333-aaaa-4aaa-8aaa-333333333333'::uuid,'33015630-09c7-4917-921b-27e2e77f257a'::uuid,1,'2026-06-14 02:28:13.141192'),
	 ('33333333-aaaa-4aaa-8aaa-333333333333'::uuid,'9e933a6f-d35d-4140-be9d-135bc95077cc'::uuid,2,'2026-06-14 02:28:13.141192'),
	 ('44444444-aaaa-4aaa-8aaa-444444444444'::uuid,'ddc8d79f-0e86-42a1-b658-f67325351823'::uuid,1,'2026-06-14 02:28:13.141192'),
	 ('44444444-aaaa-4aaa-8aaa-444444444444'::uuid,'70d1892c-47b0-4657-ad68-e56234ad49fc'::uuid,2,'2026-06-14 02:28:13.141192'),
	 ('55555555-aaaa-4aaa-8aaa-555555555555'::uuid,'0947ffc3-fe28-4ee1-ba84-35495989cf0d'::uuid,1,'2026-06-14 02:28:13.141192'),
	 ('55555555-aaaa-4aaa-8aaa-555555555555'::uuid,'33015630-09c7-4917-921b-27e2e77f257a'::uuid,2,'2026-06-14 02:28:13.141192');
INSERT INTO public.playlist_tracks (playlist_id,track_id,"position",added_at) VALUES
	 ('66666666-aaaa-4aaa-8aaa-666666666666'::uuid,'341ffba7-fadf-4ebc-ab23-a4ed5ee7dca6'::uuid,1,'2026-06-14 02:28:13.141192');

-- track_artists_202606180141.sql
INSERT INTO public.track_artists (track_id,artist_id,"role") VALUES
	 ('ddc8d79f-0e86-42a1-b658-f67325351823'::uuid,'11111111-1111-1111-1111-111111111111'::uuid,'artist'),
	 ('9e933a6f-d35d-4140-be9d-135bc95077cc'::uuid,'22222222-2222-2222-2222-222222222222'::uuid,'artist'),
	 ('70d1892c-47b0-4657-ad68-e56234ad49fc'::uuid,'44444444-4444-4444-4444-444444444444'::uuid,'artist'),
	 ('341ffba7-fadf-4ebc-ab23-a4ed5ee7dca6'::uuid,'55555555-5555-5555-5555-555555555555'::uuid,'artist'),
	 ('0947ffc3-fe28-4ee1-ba84-35495989cf0d'::uuid,'11111111-1111-1111-1111-111111111111'::uuid,'artist'),
	 ('0947ffc3-fe28-4ee1-ba84-35495989cf0d'::uuid,'22222222-2222-2222-2222-222222222222'::uuid,'artist'),
	 ('33015630-09c7-4917-921b-27e2e77f257a'::uuid,'11111111-1111-1111-1111-111111111111'::uuid,'artist'),
	 ('33015630-09c7-4917-921b-27e2e77f257a'::uuid,'22222222-2222-2222-2222-222222222222'::uuid,'artist'),
	 ('00000000-0000-0000-0000-000000000000'::uuid,'33333333-3333-3333-3333-333333333333'::uuid,'artist'),
	 ('00000000-0000-0000-0000-000000000000'::uuid,'55555555-5555-5555-5555-555555555555'::uuid,'artist');

-- track_genres_202606180141.sql
INSERT INTO public.track_genres (track_id,genre_id) VALUES
	 ('9e933a6f-d35d-4140-be9d-135bc95077cc'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc2'::uuid),
	 ('70d1892c-47b0-4657-ad68-e56234ad49fc'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc4'::uuid),
	 ('ddc8d79f-0e86-42a1-b658-f67325351823'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc5'::uuid),
	 ('0947ffc3-fe28-4ee1-ba84-35495989cf0d'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc1'::uuid),
	 ('33015630-09c7-4917-921b-27e2e77f257a'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc1'::uuid),
	 ('33015630-09c7-4917-921b-27e2e77f257a'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc2'::uuid),
	 ('00000000-0000-0000-0000-000000000000'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc1'::uuid),
	 ('00000000-0000-0000-0000-000000000000'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc3'::uuid),
	 ('341ffba7-fadf-4ebc-ab23-a4ed5ee7dca6'::uuid,'cccccccc-cccc-cccc-cccc-ccccccccccc5'::uuid);

-- user_liked_tracks_202606180141.sql
-- empty export file

