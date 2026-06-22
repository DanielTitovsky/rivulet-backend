-- Seed tracks for Rivulet
-- Generated from list.txt.
-- IMPORTANT:
-- 1) audio_storage_key uses prefix: songs/<filename>.
--    If your files are stored in another MinIO folder, replace 'songs/' below.
-- 2) duration_seconds is set to 180 because only filenames were provided, not audio metadata.
-- 3) Cover paths follow requested format:
--    /cover/artist/name.jpg
--    /cover/album/name.jpg
--    /cover/track/name.jpg
-- 4) If a track belongs to an album, its cover_url is set to the album cover_url.

BEGIN;

INSERT INTO public.track_statuses (code, name)
VALUES ('published', 'Published')
ON CONFLICT (code) DO NOTHING;

INSERT INTO public.genres (name)
VALUES
    ('Alternative'),
    ('Alternative Rock'),
    ('Ambient'),
    ('Electronic'),
    ('Hip-Hop'),
    ('Indie'),
    ('Indie Rock'),
    ('Pop'),
    ('Rap')
ON CONFLICT (name) DO NOTHING;

CREATE TEMP TABLE seed_artists (
    name text NOT NULL,
    description text,
    avatar_url text
) ON COMMIT DROP;

INSERT INTO seed_artists (name, description, avatar_url)
VALUES
    ('$LOTHBOI', 'Исполнитель $LOTHBOI', '/cover/artist/SLOTHBOI.jpg'),
    ('Arctic Monkeys', 'Исполнитель Arctic Monkeys', '/cover/artist/Arctic Monkeys.jpg'),
    ('BONES', 'Исполнитель BONES', '/cover/artist/BONES.jpg'),
    ('Burgos', 'Исполнитель Burgos', '/cover/artist/Burgos.jpg'),
    ('Caudal', 'Исполнитель Caudal', '/cover/artist/Caudal.jpg'),
    ('HAARPER', 'Исполнитель HAARPER', '/cover/artist/HAARPER.jpg'),
    ('Oliver Tree', 'Исполнитель Oliver Tree', '/cover/artist/Oliver Tree.jpg'),
    ('Razegod', 'Исполнитель Razegod', '/cover/artist/Razegod.jpg'),
    ('Ricky A Go Go', 'Исполнитель Ricky A Go Go', '/cover/artist/Ricky A Go Go.jpg'),
    ('Robert DeLong', 'Исполнитель Robert DeLong', '/cover/artist/Robert DeLong.jpg'),
    ('alt-J', 'Исполнитель alt-J', '/cover/artist/alt-J.jpg'),
    ('overtonight', 'Исполнитель overtonight', '/cover/artist/overtonight.jpg'),
    ('overtonight throwaways', 'Исполнитель overtonight throwaways', '/cover/artist/overtonight throwaways.jpg'),
    ('week7', 'Исполнитель week7', '/cover/artist/week7.jpg');

INSERT INTO public.artists (name, description, avatar_url)
SELECT sa.name, sa.description, sa.avatar_url
FROM seed_artists sa
WHERE NOT EXISTS (
    SELECT 1
    FROM public.artists a
    WHERE lower(a.name) = lower(sa.name)
);

CREATE TEMP TABLE seed_albums (
    title text NOT NULL,
    description text,
    cover_url text,
    artist_name text NOT NULL,
    release_date date
) ON COMMIT DROP;

INSERT INTO seed_albums (title, description, cover_url, artist_name, release_date)
VALUES
    ('AM', 'Альбом AM', '/cover/album/AM.jpg', 'Arctic Monkeys', DATE '2013-09-09'),
    ('An Awesome Wave', 'Альбом An Awesome Wave', '/cover/album/An Awesome Wave.jpg', 'alt-J', DATE '2012-05-25'),
    ('Favourite Worst Nightmare', 'Альбом Favourite Worst Nightmare', '/cover/album/Favourite Worst Nightmare.jpg', 'Arctic Monkeys', DATE '2007-04-23'),
    ('UNRENDERED', 'Альбом UNRENDERED', '/cover/album/UNRENDERED.jpg', 'BONES', NULL),
    ('Ugly Is Beautiful', 'Альбом Ugly Is Beautiful', '/cover/album/Ugly Is Beautiful.jpg', 'Oliver Tree', DATE '2020-07-17'),
    ('Why''d You Only Call Me When You''re High? (Single)', 'Альбом Why''d You Only Call Me When You''re High? (Single)', '/cover/album/Whyd You Only Call Me When Youre High (Single).jpg', 'Arctic Monkeys', DATE '2013-09-02'),
    ('Y2K', 'Альбом Y2K', '/cover/album/Y2K.jpg', 'HAARPER', NULL);

INSERT INTO public.albums (title, description, cover_url, release_date)
SELECT sa.title, sa.description, sa.cover_url, sa.release_date
FROM seed_albums sa
WHERE NOT EXISTS (
    SELECT 1
    FROM public.albums a
    WHERE lower(a.title) = lower(sa.title)
);

INSERT INTO public.album_artists (album_id, artist_id, role)
SELECT al.id, ar.id, 'artist'
FROM seed_albums sa
JOIN public.albums al ON lower(al.title) = lower(sa.title)
JOIN public.artists ar ON lower(ar.name) = lower(sa.artist_name)
ON CONFLICT (album_id, artist_id, role) DO NOTHING;

CREATE TEMP TABLE seed_tracks (
    filename text NOT NULL,
    title text NOT NULL,
    description text,
    cover_url text,
    duration_seconds integer NOT NULL,
    audio_storage_key text NOT NULL,
    status_code text NOT NULL,
    album_title text,
    release_date date
) ON COMMIT DROP;

INSERT INTO seed_tracks (
    filename,
    title,
    description,
    cover_url,
    duration_seconds,
    audio_storage_key,
    status_code,
    album_title,
    release_date
)
VALUES
       ('Alt-J - Blood Flood (hitmo.net).mp3', 'Bloodflood', 'Трек Bloodflood', '/cover/track/Bloodflood.jpg', 180, 'songs/Alt-J - Blood Flood (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),
    ('Alt-J - Breezeblocks (hitmo.net).mp3', 'Breezeblocks', 'Трек Breezeblocks', '/cover/track/Breezeblocks.jpg', 180, 'songs/Alt-J - Breezeblocks (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),
    ('Alt-J - Dissolve Me (hitmo.net).mp3', 'Dissolve Me', 'Трек Dissolve Me', '/cover/track/Dissolve Me.jpg', 180, 'songs/Alt-J - Dissolve Me (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),
    ('Alt-J - Fitzpleasure (hitmo.net).mp3', 'Fitzpleasure', 'Трек Fitzpleasure', '/cover/track/Fitzpleasure.jpg', 180, 'songs/Alt-J - Fitzpleasure (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),
    ('Alt-J - Intro (hitmo.net).mp3', 'Intro', 'Трек Intro', '/cover/track/Intro.jpg', 180, 'songs/Alt-J - Intro (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),
    ('Alt-J - Matilda (hitmo.net).mp3', 'Matilda', 'Трек Matilda', '/cover/track/Matilda.jpg', 180, 'songs/Alt-J - Matilda (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),
    ('Alt-J - Ms (hitmo.net).mp3', 'Ms', 'Трек Ms', '/cover/track/Ms.jpg', 180, 'songs/Alt-J - Ms (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),
    ('Alt-J - Tessellate (hitmo.net).mp3', 'Tessellate', 'Трек Tessellate', '/cover/track/Tessellate.jpg', 180, 'songs/Alt-J - Tessellate (hitmo.net).mp3', 'published', 'An Awesome Wave', DATE '2012-05-25'),

    ('Arctic Monkeys - 505.mp3', '505', 'Трек 505', '/cover/track/505.jpg', 180, 'songs/Arctic Monkeys - 505.mp3', 'published', 'Favourite Worst Nightmare', DATE '2007-04-23'),
    ('Arctic Monkeys - Stop The World I Wanna Get Off With You.mp3', 'Stop The World I Wanna Get Off With You', 'Трек Stop The World I Wanna Get Off With You', '/cover/track/Stop The World I Wanna Get Off With You.jpg', 180, 'songs/Arctic Monkeys - Stop The World I Wanna Get Off With You.mp3', 'published', 'Why''d You Only Call Me When You''re High? (Single)', DATE '2013-09-02'),
    ('Arctic Monkeys - Why''d You Only Call Me When You''re High_.mp3', 'Why''d You Only Call Me When You''re High?', 'Трек Why''d You Only Call Me When You''re High?', '/cover/track/Whyd You Only Call Me When Youre High.jpg', 180, 'songs/Arctic Monkeys - Why''d You Only Call Me When You''re High_.mp3', 'published', 'AM', DATE '2013-09-09'),
    ('Artic Monkeys - 505.mp3', '505', 'Трек 505', '/cover/track/505.jpg', 180, 'songs/Artic Monkeys - 505.mp3', 'published', 'Favourite Worst Nightmare', DATE '2007-04-23'),

    ('Bones - ContinueWithoutSaving! (hitmo.net).mp3', 'ContinueWithoutSaving!', 'Трек ContinueWithoutSaving!', '/cover/track/ContinueWithoutSaving.jpg', 180, 'songs/Bones - ContinueWithoutSaving! (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - CtrlAltDelete (hitmo.net).mp3', 'CtrlAltDelete', 'Трек CtrlAltDelete', '/cover/track/CtrlAltDelete.jpg', 180, 'songs/Bones - CtrlAltDelete (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - Importing (hitmo.net).mp3', 'Importing', 'Трек Importing', '/cover/track/Importing.jpg', 180, 'songs/Bones - Importing (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - LifeRuiner (hitmo.net).mp3', 'LifeRuiner', 'Трек LifeRuiner', '/cover/track/LifeRuiner.jpg', 180, 'songs/Bones - LifeRuiner (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - MinorSetback (hitmo.net).mp3', 'MinorSetback', 'Трек MinorSetback', '/cover/track/MinorSetback.jpg', 180, 'songs/Bones - MinorSetback (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - MissingProjectFiles (hitmo.net).mp3', 'MissingProjectFiles', 'Трек MissingProjectFiles', '/cover/track/MissingProjectFiles.jpg', 180, 'songs/Bones - MissingProjectFiles (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - MyNephewHasAWhitePickupTruck (hitmo.net).mp3', 'MyNephewHasAWhitePickupTruck', 'Трек MyNephewHasAWhitePickupTruck', '/cover/track/MyNephewHasAWhitePickupTruck.jpg', 180, 'songs/Bones - MyNephewHasAWhitePickupTruck (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - SunnyDay (hitmo.net).mp3', 'SunnyDay', 'Трек SunnyDay', '/cover/track/SunnyDay.jpg', 180, 'songs/Bones - SunnyDay (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - SystemPreferences (hitmo.net).mp3', 'SystemPreferences', 'Трек SystemPreferences', '/cover/track/SystemPreferences.jpg', 180, 'songs/Bones - SystemPreferences (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones - WhateverHelpsYouSleep (hitmo.net).mp3', 'WhateverHelpsYouSleep', 'Трек WhateverHelpsYouSleep', '/cover/track/WhateverHelpsYouSleep.jpg', 180, 'songs/Bones - WhateverHelpsYouSleep (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),
    ('Bones feat Ricky A Go Go - YouKnowIWantYou (hitmo.net).mp3', 'YouKnowIWantYou', 'Трек YouKnowIWantYou', '/cover/track/YouKnowIWantYou.jpg', 180, 'songs/Bones feat Ricky A Go Go - YouKnowIWantYou (hitmo.net).mp3', 'published', 'UNRENDERED', DATE '2017-01-01'),

    ('Burgos - Fly My Kite.mp3', 'Fly My Kite', 'Трек Fly My Kite', '/cover/track/Fly My Kite.jpg', 180, 'songs/Burgos - Fly My Kite.mp3', 'published', NULL, DATE '2020-01-01'),
    ('Burgos - Pancakes.mp3', 'Pancakes', 'Трек Pancakes', '/cover/track/Pancakes.jpg', 180, 'songs/Burgos - Pancakes.mp3', 'published', NULL, DATE '2020-01-01'),
    ('Burgos - Ugly.mp3', 'Ugly', 'Трек Ugly', '/cover/track/Ugly.jpg', 180, 'songs/Burgos - Ugly.mp3', 'published', NULL, DATE '2020-01-01'),

    ('Caudal - Still Circuit.mp3', 'Still Circuit', 'Трек Still Circuit', '/cover/track/Still Circuit.jpg', 180, 'songs/Caudal - Still Circuit.mp3', 'published', NULL, DATE '2020-01-01'),

    ('Haarper - 8 LEGGED FREAK (hitmo.net).mp3', '8 LEGGED FREAK', 'Трек 8 LEGGED FREAK', '/cover/track/8 LEGGED FREAK.jpg', 180, 'songs/Haarper - 8 LEGGED FREAK (hitmo.net).mp3', 'published', 'Y2K', DATE '2019-01-01'),
    ('Haarper - CHAINS LIKE A CENOBITE (hitmo.net).mp3', 'CHAINS LIKE A CENOBITE', 'Трек CHAINS LIKE A CENOBITE', '/cover/track/CHAINS LIKE A CENOBITE.jpg', 180, 'songs/Haarper - CHAINS LIKE A CENOBITE (hitmo.net).mp3', 'published', 'Y2K', DATE '2019-01-01'),
    ('Haarper - DEVIL MAY KRY (hitmo.net).mp3', 'DEVIL MAY KRY', 'Трек DEVIL MAY KRY', '/cover/track/DEVIL MAY KRY.jpg', 180, 'songs/Haarper - DEVIL MAY KRY (hitmo.net).mp3', 'published', 'Y2K', DATE '2019-01-01'),
    ('Haarper - LET EM BLEED (hitmo.net).mp3', 'LET EM BLEED', 'Трек LET EM BLEED', '/cover/track/LET EM BLEED.jpg', 180, 'songs/Haarper - LET EM BLEED (hitmo.net).mp3', 'published', 'Y2K', DATE '2019-01-01'),
    ('Haarper - Y2K (hitmo.net).mp3', 'Y2K', 'Трек Y2K', '/cover/track/Y2K.jpg', 180, 'songs/Haarper - Y2K (hitmo.net).mp3', 'published', 'Y2K', DATE '2019-01-01'),

    ('Oliver Tree - Alien Boy (hitmo.net).mp3', 'Alien Boy', 'Трек Alien Boy', '/cover/track/Alien Boy.jpg', 180, 'songs/Oliver Tree - Alien Boy (hitmo.net).mp3', 'published', 'Ugly Is Beautiful', DATE '2020-07-17'),
    ('Oliver Tree - Cash Machine (hitmo.net).mp3', 'Cash Machine', 'Трек Cash Machine', '/cover/track/Cash Machine.jpg', 180, 'songs/Oliver Tree - Cash Machine (hitmo.net).mp3', 'published', 'Ugly Is Beautiful', DATE '2020-07-17'),
    ('Oliver Tree - Hurt (hitmo.net).mp3', 'Hurt', 'Трек Hurt', '/cover/track/Hurt.jpg', 180, 'songs/Oliver Tree - Hurt (hitmo.net).mp3', 'published', 'Ugly Is Beautiful', DATE '2020-07-17'),
    ('Oliver Tree - Joke''s On You! (hitmo.net).mp3', 'Joke''s On You!', 'Трек Joke''s On You!', '/cover/track/Jokes On You.jpg', 180, 'songs/Oliver Tree - Joke''s On You! (hitmo.net).mp3', 'published', 'Ugly Is Beautiful', DATE '2020-07-17'),
    ('Oliver Tree - Me, Myself and I (hitmo.net).mp3', 'Me, Myself and I', 'Трек Me, Myself and I', '/cover/track/Me, Myself and I.jpg', 180, 'songs/Oliver Tree - Me, Myself and I (hitmo.net).mp3', 'published', 'Ugly Is Beautiful', DATE '2020-07-17'),
    ('Oliver Tree - Miracle Man (hitmo.net).mp3', 'Miracle Man', 'Трек Miracle Man', '/cover/track/Miracle Man.jpg', 180, 'songs/Oliver Tree - Miracle Man (hitmo.net).mp3', 'published', 'Ugly Is Beautiful', DATE '2020-07-17'),
    ('Oliver Tree - Waste My Time (hitmo.net).mp3', 'Waste My Time', 'Трек Waste My Time', '/cover/track/Waste My Time.jpg', 180, 'songs/Oliver Tree - Waste My Time (hitmo.net).mp3', 'published', 'Ugly Is Beautiful', DATE '2020-07-17'),

    ('overtonight_throwaways_overtonight_-_mirrors_demo_(SkySound.cc).mp3', 'mirrors demo', 'Трек mirrors demo', '/cover/track/mirrors demo.jpg', 180, 'songs/overtonight_throwaways_overtonight_-_mirrors_demo_(SkySound.cc).mp3', 'published', NULL, DATE '2023-01-01'),

    ('Razegod - Don''t Run (hitmo.net).mp3', 'Don''t Run', 'Трек Don''t Run', '/cover/track/Dont Run.jpg', 180, 'songs/Razegod - Don''t Run (hitmo.net).mp3', 'published', NULL, DATE '2021-01-01'),
    ('Razegod - Nerv (hitmo.net).mp3', 'Nerv', 'Трек Nerv', '/cover/track/Nerv.jpg', 180, 'songs/Razegod - Nerv (hitmo.net).mp3', 'published', NULL, DATE '2021-01-01'),
    ('Razegod - Teleport Tab (hitmo.net).mp3', 'Teleport Tab', 'Трек Teleport Tab', '/cover/track/Teleport Tab.jpg', 180, 'songs/Razegod - Teleport Tab (hitmo.net).mp3', 'published', NULL, DATE '2021-01-01'),
    ('Razegod - Which Way_ (hitmo.net).mp3', 'Which Way?', 'Трек Which Way?', '/cover/track/Which Way.jpg', 180, 'songs/Razegod - Which Way_ (hitmo.net).mp3', 'published', NULL, DATE '2021-01-01'),
    ('Razegod, $LOTHBOI - Creation (hitmo.net).mp3', 'Creation', 'Трек Creation', '/cover/track/Creation.jpg', 180, 'songs/Razegod, $LOTHBOI - Creation (hitmo.net).mp3', 'published', NULL, DATE '2021-01-01'),

    ('Robert Delong - First Person On Earth.mp3', 'First Person On Earth', 'Трек First Person On Earth', '/cover/track/First Person On Earth.jpg', 180, 'songs/Robert Delong - First Person On Earth.mp3', 'published', NULL, DATE '2018-01-01'),
    ('week7-brrring.mp3', 'brrring', 'Трек brrring', '/cover/track/brrring.jpg', 180, 'songs/week7-brrring.mp3', 'published', NULL, DATE '2023-01-01');

INSERT INTO public.tracks (
    album_id,
    title,
    description,
    cover_url,
    duration_seconds,
    release_date,
    is_explicit,
    is_streamable,
    is_downloadable,
    audio_storage_key,
    status_id
)
SELECT
    al.id,
    st.title,
    st.description,
    CASE
        WHEN al.id IS NOT NULL THEN al.cover_url
        ELSE st.cover_url
    END AS cover_url,
    st.duration_seconds,
    st.release_date,
    false,
    true,
    false,
    st.audio_storage_key,
    ts.id
FROM seed_tracks st
JOIN public.track_statuses ts ON ts.code = st.status_code
LEFT JOIN public.albums al ON st.album_title IS NOT NULL AND lower(al.title) = lower(st.album_title)
WHERE NOT EXISTS (
    SELECT 1
    FROM public.tracks t
    WHERE t.audio_storage_key = st.audio_storage_key
);

-- If the seed was already applied earlier, update covers for existing album tracks too.
UPDATE public.tracks t
SET
    cover_url = al.cover_url,
    updated_at = now()
FROM seed_tracks st
JOIN public.albums al ON st.album_title IS NOT NULL AND lower(al.title) = lower(st.album_title)
WHERE t.audio_storage_key = st.audio_storage_key;

CREATE TEMP TABLE seed_track_artists (
    filename text NOT NULL,
    artist_name text NOT NULL,
    role text NOT NULL DEFAULT 'artist'
) ON COMMIT DROP;

INSERT INTO seed_track_artists (filename, artist_name, role)
VALUES
    ('Alt-J - Blood Flood (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Alt-J - Breezeblocks (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Alt-J - Dissolve Me (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Alt-J - Fitzpleasure (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Alt-J - Intro (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Alt-J - Matilda (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Alt-J - Ms (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Alt-J - Tessellate (hitmo.net).mp3', 'alt-J', 'artist'),
    ('Arctic Monkeys - 505.mp3', 'Arctic Monkeys', 'artist'),
    ('Arctic Monkeys - Stop The World I Wanna Get Off With You.mp3', 'Arctic Monkeys', 'artist'),
    ('Arctic Monkeys - Why''d You Only Call Me When You''re High_.mp3', 'Arctic Monkeys', 'artist'),
    ('Artic Monkeys - 505.mp3', 'Arctic Monkeys', 'artist'),
    ('Bones - ContinueWithoutSaving! (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - CtrlAltDelete (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - Importing (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - LifeRuiner (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - MinorSetback (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - MissingProjectFiles (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - MyNephewHasAWhitePickupTruck (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - SunnyDay (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - SystemPreferences (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones - WhateverHelpsYouSleep (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones feat Ricky A Go Go - YouKnowIWantYou (hitmo.net).mp3', 'BONES', 'artist'),
    ('Bones feat Ricky A Go Go - YouKnowIWantYou (hitmo.net).mp3', 'Ricky A Go Go', 'featured'),
    ('Burgos - Fly My Kite.mp3', 'Burgos', 'artist'),
    ('Burgos - Pancakes.mp3', 'Burgos', 'artist'),
    ('Burgos - Ugly.mp3', 'Burgos', 'artist'),
    ('Caudal - Still Circuit.mp3', 'Caudal', 'artist'),
    ('Haarper - 8 LEGGED FREAK (hitmo.net).mp3', 'HAARPER', 'artist'),
    ('Haarper - CHAINS LIKE A CENOBITE (hitmo.net).mp3', 'HAARPER', 'artist'),
    ('Haarper - DEVIL MAY KRY (hitmo.net).mp3', 'HAARPER', 'artist'),
    ('Haarper - LET EM BLEED (hitmo.net).mp3', 'HAARPER', 'artist'),
    ('Haarper - Y2K (hitmo.net).mp3', 'HAARPER', 'artist'),
    ('Oliver Tree - Alien Boy (hitmo.net).mp3', 'Oliver Tree', 'artist'),
    ('Oliver Tree - Cash Machine (hitmo.net).mp3', 'Oliver Tree', 'artist'),
    ('Oliver Tree - Hurt (hitmo.net).mp3', 'Oliver Tree', 'artist'),
    ('Oliver Tree - Joke''s On You! (hitmo.net).mp3', 'Oliver Tree', 'artist'),
    ('Oliver Tree - Me, Myself and I (hitmo.net).mp3', 'Oliver Tree', 'artist'),
    ('Oliver Tree - Miracle Man (hitmo.net).mp3', 'Oliver Tree', 'artist'),
    ('Oliver Tree - Waste My Time (hitmo.net).mp3', 'Oliver Tree', 'artist'),
    ('overtonight_throwaways_overtonight_-_mirrors_demo_(SkySound.cc).mp3', 'overtonight throwaways', 'artist'),
    ('overtonight_throwaways_overtonight_-_mirrors_demo_(SkySound.cc).mp3', 'overtonight', 'featured'),
    ('Razegod - Don''t Run (hitmo.net).mp3', 'Razegod', 'artist'),
    ('Razegod - Nerv (hitmo.net).mp3', 'Razegod', 'artist'),
    ('Razegod - Teleport Tab (hitmo.net).mp3', 'Razegod', 'artist'),
    ('Razegod - Which Way_ (hitmo.net).mp3', 'Razegod', 'artist'),
    ('Razegod, $LOTHBOI - Creation (hitmo.net).mp3', 'Razegod', 'artist'),
    ('Razegod, $LOTHBOI - Creation (hitmo.net).mp3', '$LOTHBOI', 'featured'),
    ('Robert Delong - First Person On Earth.mp3', 'Robert DeLong', 'artist'),
    ('week7-brrring.mp3', 'week7', 'artist');

INSERT INTO public.track_artists (track_id, artist_id, role)
SELECT t.id, a.id, sta.role
FROM seed_track_artists sta
JOIN seed_tracks st ON st.filename = sta.filename
JOIN public.tracks t ON t.audio_storage_key = st.audio_storage_key
JOIN public.artists a ON lower(a.name) = lower(sta.artist_name)
ON CONFLICT (track_id, artist_id, role) DO NOTHING;

CREATE TEMP TABLE seed_track_genres (
    filename text NOT NULL,
    genre_name text NOT NULL
) ON COMMIT DROP;

INSERT INTO seed_track_genres (filename, genre_name)
VALUES
    ('Alt-J - Blood Flood (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Blood Flood (hitmo.net).mp3', 'Alternative'),
    ('Alt-J - Breezeblocks (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Breezeblocks (hitmo.net).mp3', 'Alternative'),
    ('Alt-J - Dissolve Me (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Dissolve Me (hitmo.net).mp3', 'Alternative'),
    ('Alt-J - Fitzpleasure (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Fitzpleasure (hitmo.net).mp3', 'Alternative'),
    ('Alt-J - Intro (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Intro (hitmo.net).mp3', 'Alternative'),
    ('Alt-J - Matilda (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Matilda (hitmo.net).mp3', 'Alternative'),
    ('Alt-J - Ms (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Ms (hitmo.net).mp3', 'Alternative'),
    ('Alt-J - Tessellate (hitmo.net).mp3', 'Indie Rock'),
    ('Alt-J - Tessellate (hitmo.net).mp3', 'Alternative'),
    ('Arctic Monkeys - 505.mp3', 'Alternative Rock'),
    ('Arctic Monkeys - 505.mp3', 'Indie Rock'),
    ('Arctic Monkeys - Stop The World I Wanna Get Off With You.mp3', 'Alternative Rock'),
    ('Arctic Monkeys - Stop The World I Wanna Get Off With You.mp3', 'Indie Rock'),
    ('Arctic Monkeys - Why''d You Only Call Me When You''re High_.mp3', 'Alternative Rock'),
    ('Arctic Monkeys - Why''d You Only Call Me When You''re High_.mp3', 'Indie Rock'),
    ('Artic Monkeys - 505.mp3', 'Alternative Rock'),
    ('Artic Monkeys - 505.mp3', 'Indie Rock'),
    ('Bones - ContinueWithoutSaving! (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - ContinueWithoutSaving! (hitmo.net).mp3', 'Rap'),
    ('Bones - CtrlAltDelete (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - CtrlAltDelete (hitmo.net).mp3', 'Rap'),
    ('Bones - Importing (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - Importing (hitmo.net).mp3', 'Rap'),
    ('Bones - LifeRuiner (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - LifeRuiner (hitmo.net).mp3', 'Rap'),
    ('Bones - MinorSetback (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - MinorSetback (hitmo.net).mp3', 'Rap'),
    ('Bones - MissingProjectFiles (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - MissingProjectFiles (hitmo.net).mp3', 'Rap'),
    ('Bones - MyNephewHasAWhitePickupTruck (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - MyNephewHasAWhitePickupTruck (hitmo.net).mp3', 'Rap'),
    ('Bones - SunnyDay (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - SunnyDay (hitmo.net).mp3', 'Rap'),
    ('Bones - SystemPreferences (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - SystemPreferences (hitmo.net).mp3', 'Rap'),
    ('Bones - WhateverHelpsYouSleep (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones - WhateverHelpsYouSleep (hitmo.net).mp3', 'Rap'),
    ('Bones feat Ricky A Go Go - YouKnowIWantYou (hitmo.net).mp3', 'Hip-Hop'),
    ('Bones feat Ricky A Go Go - YouKnowIWantYou (hitmo.net).mp3', 'Rap'),
    ('Burgos - Fly My Kite.mp3', 'Hip-Hop'),
    ('Burgos - Fly My Kite.mp3', 'Rap'),
    ('Burgos - Pancakes.mp3', 'Hip-Hop'),
    ('Burgos - Pancakes.mp3', 'Rap'),
    ('Burgos - Ugly.mp3', 'Hip-Hop'),
    ('Burgos - Ugly.mp3', 'Rap'),
    ('Caudal - Still Circuit.mp3', 'Electronic'),
    ('Caudal - Still Circuit.mp3', 'Ambient'),
    ('Haarper - 8 LEGGED FREAK (hitmo.net).mp3', 'Hip-Hop'),
    ('Haarper - 8 LEGGED FREAK (hitmo.net).mp3', 'Rap'),
    ('Haarper - CHAINS LIKE A CENOBITE (hitmo.net).mp3', 'Hip-Hop'),
    ('Haarper - CHAINS LIKE A CENOBITE (hitmo.net).mp3', 'Rap'),
    ('Haarper - DEVIL MAY KRY (hitmo.net).mp3', 'Hip-Hop'),
    ('Haarper - DEVIL MAY KRY (hitmo.net).mp3', 'Rap'),
    ('Haarper - LET EM BLEED (hitmo.net).mp3', 'Hip-Hop'),
    ('Haarper - LET EM BLEED (hitmo.net).mp3', 'Rap'),
    ('Haarper - Y2K (hitmo.net).mp3', 'Hip-Hop'),
    ('Haarper - Y2K (hitmo.net).mp3', 'Rap'),
    ('Oliver Tree - Alien Boy (hitmo.net).mp3', 'Alternative'),
    ('Oliver Tree - Alien Boy (hitmo.net).mp3', 'Pop'),
    ('Oliver Tree - Cash Machine (hitmo.net).mp3', 'Alternative'),
    ('Oliver Tree - Cash Machine (hitmo.net).mp3', 'Pop'),
    ('Oliver Tree - Hurt (hitmo.net).mp3', 'Alternative'),
    ('Oliver Tree - Hurt (hitmo.net).mp3', 'Pop'),
    ('Oliver Tree - Joke''s On You! (hitmo.net).mp3', 'Alternative'),
    ('Oliver Tree - Joke''s On You! (hitmo.net).mp3', 'Pop'),
    ('Oliver Tree - Me, Myself and I (hitmo.net).mp3', 'Alternative'),
    ('Oliver Tree - Me, Myself and I (hitmo.net).mp3', 'Pop'),
    ('Oliver Tree - Miracle Man (hitmo.net).mp3', 'Alternative'),
    ('Oliver Tree - Miracle Man (hitmo.net).mp3', 'Pop'),
    ('Oliver Tree - Waste My Time (hitmo.net).mp3', 'Alternative'),
    ('Oliver Tree - Waste My Time (hitmo.net).mp3', 'Pop'),
    ('overtonight_throwaways_overtonight_-_mirrors_demo_(SkySound.cc).mp3', 'Indie'),
    ('overtonight_throwaways_overtonight_-_mirrors_demo_(SkySound.cc).mp3', 'Alternative'),
    ('Razegod - Don''t Run (hitmo.net).mp3', 'Hip-Hop'),
    ('Razegod - Don''t Run (hitmo.net).mp3', 'Rap'),
    ('Razegod - Nerv (hitmo.net).mp3', 'Hip-Hop'),
    ('Razegod - Nerv (hitmo.net).mp3', 'Rap'),
    ('Razegod - Teleport Tab (hitmo.net).mp3', 'Hip-Hop'),
    ('Razegod - Teleport Tab (hitmo.net).mp3', 'Rap'),
    ('Razegod - Which Way_ (hitmo.net).mp3', 'Hip-Hop'),
    ('Razegod - Which Way_ (hitmo.net).mp3', 'Rap'),
    ('Razegod, $LOTHBOI - Creation (hitmo.net).mp3', 'Hip-Hop'),
    ('Razegod, $LOTHBOI - Creation (hitmo.net).mp3', 'Rap'),
    ('Robert Delong - First Person On Earth.mp3', 'Electronic'),
    ('Robert Delong - First Person On Earth.mp3', 'Alternative'),
    ('week7-brrring.mp3', 'Electronic');

INSERT INTO public.track_genres (track_id, genre_id)
SELECT t.id, g.id
FROM seed_track_genres stg
JOIN seed_tracks st ON st.filename = stg.filename
JOIN public.tracks t ON t.audio_storage_key = st.audio_storage_key
JOIN public.genres g ON g.name = stg.genre_name
ON CONFLICT (track_id, genre_id) DO NOTHING;

COMMIT;

-- =========================
-- USERS SEED
-- =========================

INSERT INTO users (
    id,
    email,
    name,
    hash_password
)
VALUES
(
    '00000000-0000-0000-0000-000000000101',
    'demo-user@mail.com',
    'Demo User',
      '$2a$14$oFG2YWWttbYa0dWn/eY4J.JgL49ixJI/yXKUwllYy3M4RlSpj.JAO'
),
(
    '00000000-0000-0000-0000-000000000102',
    'rivulet-listener@mail.com',
    'Rivulet Listener',
      '$2a$14$oFG2YWWttbYa0dWn/eY4J.JgL49ixJI/yXKUwllYy3M4RlSpj.JAO'
)
ON CONFLICT (email)
DO UPDATE SET
    name = EXCLUDED.name;


-- =========================
-- PLAYLISTS SEED
-- =========================

INSERT INTO playlists (
    id,
    user_id,
    title,
    description,
    cover_url,
    visibility
)
VALUES
(
    '00000000-0000-0000-0000-000000000201',
    '00000000-0000-0000-0000-000000000101',
    'My daily mix',
    'Tracks for everyday listening',
    '/cover/playlist/my-daily-mix.jpg',
    true
),
(
    '00000000-0000-0000-0000-000000000202',
    '00000000-0000-0000-0000-000000000101',
    'Indie mood',
    'Soft indie and alternative tracks',
    '/cover/playlist/indie-mood.jpg',
    true
),
(
    '00000000-0000-0000-0000-000000000203',
    '00000000-0000-0000-0000-000000000102',
    'Evening playlist',
    'Music for calm evening',
    '/cover/playlist/evening-playlist.jpg',
    false
)
ON CONFLICT (id)
DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    cover_url = EXCLUDED.cover_url,
    visibility = EXCLUDED.visibility;


-- =========================
-- PLAYLIST TRACKS SEED
-- =========================

INSERT INTO playlists (
    id,
    user_id,
    title,
    description,
    cover_url,
    visibility
)
VALUES
(
    '00000000-0000-0000-0000-000000000201',
    '00000000-0000-0000-0000-000000000101',
    'My daily mix',
    'Tracks for everyday listening',
    '/cover/playlist/my-daily-mix.jpg',
    'true'
),
(
    '00000000-0000-0000-0000-000000000202',
    '00000000-0000-0000-0000-000000000101',
    'Indie mood',
    'Soft indie and alternative tracks',
    '/cover/playlist/indie-mood.jpg',
    'true'
),
(
    '00000000-0000-0000-0000-000000000203',
    '00000000-0000-0000-0000-000000000102',
    'Evening playlist',
    'Music for calm evening',
    '/cover/playlist/evening-playlist.jpg',
    'false'
)
ON CONFLICT (id)
DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    cover_url = EXCLUDED.cover_url,
    visibility = EXCLUDED.visibility;

WITH playlist_track_seed AS (
    SELECT
        playlist_id::uuid,
        audio_storage_key,
        position
    FROM (
        VALUES
            -- My daily mix
            (
                '00000000-0000-0000-0000-000000000201',
                'songs/Oliver Tree - Alien Boy (hitmo.net).mp3',
                1
            ),
            (
                '00000000-0000-0000-0000-000000000201',
                'songs/Oliver Tree - Cash Machine (hitmo.net).mp3',
                2
            ),
            (
                '00000000-0000-0000-0000-000000000201',
                'songs/Arctic Monkeys - 505.mp3',
                3
            ),
            (
                '00000000-0000-0000-0000-000000000201',
                'songs/Arctic Monkeys - Why''d You Only Call Me When You''re High_.mp3',
                4
            ),
            (
                '00000000-0000-0000-0000-000000000201',
                'songs/Alt-J - Breezeblocks (hitmo.net).mp3',
                5
            ),

            -- Indie mood
            (
                '00000000-0000-0000-0000-000000000202',
                'songs/Alt-J - Blood Flood (hitmo.net).mp3',
                1
            ),
            (
                '00000000-0000-0000-0000-000000000202',
                'songs/Alt-J - Dissolve Me (hitmo.net).mp3',
                2
            ),
            (
                '00000000-0000-0000-0000-000000000202',
                'songs/Alt-J - Fitzpleasure (hitmo.net).mp3',
                3
            ),
            (
                '00000000-0000-0000-0000-000000000202',
                'songs/Alt-J - Matilda (hitmo.net).mp3',
                4
            ),
            (
                '00000000-0000-0000-0000-000000000202',
                'songs/Alt-J - Tessellate (hitmo.net).mp3',
                5
            ),

            -- Evening playlist
            (
                '00000000-0000-0000-0000-000000000203',
                'songs/Caudal - Still Circuit.mp3',
                1
            ),
            (
                '00000000-0000-0000-0000-000000000203',
                'songs/Robert Delong - First Person On Earth.mp3',
                2
            ),
            (
                '00000000-0000-0000-0000-000000000203',
                'songs/overtonight_throwaways_overtonight_-_mirrors_demo_(SkySound.cc).mp3',
                3
            ),
            (
                '00000000-0000-0000-0000-000000000203',
                'songs/week7-brrring.mp3',
                4
            ),
            (
                '00000000-0000-0000-0000-000000000203',
                'songs/Oliver Tree - Hurt (hitmo.net).mp3',
                5
            )
    ) AS seed(playlist_id, audio_storage_key, position)
)
INSERT INTO public.playlist_tracks (
    playlist_id,
    track_id,
    position
)
SELECT
    playlist_track_seed.playlist_id,
    tracks.id,
    playlist_track_seed.position
FROM playlist_track_seed
JOIN public.tracks
    ON tracks.audio_storage_key = playlist_track_seed.audio_storage_key
ON CONFLICT (playlist_id, track_id)
DO UPDATE SET
    position = EXCLUDED.position;
-- =========================
-- USER FAVORITE TRACKS SEED
-- =========================

WITH track_list AS (
    SELECT
        id,
        row_number() OVER (ORDER BY id) AS rn
    FROM tracks
),
favorite_track_seed AS (
    SELECT
        user_id::uuid,
        track_number
    FROM (
        VALUES
            ('00000000-0000-0000-0000-000000000101', 1),
            ('00000000-0000-0000-0000-000000000101', 3),
            ('00000000-0000-0000-0000-000000000101', 5),
            ('00000000-0000-0000-0000-000000000101', 7),

            ('00000000-0000-0000-0000-000000000102', 2),
            ('00000000-0000-0000-0000-000000000102', 4),
            ('00000000-0000-0000-0000-000000000102', 6),
            ('00000000-0000-0000-0000-000000000102', 8)
    ) AS seed(user_id, track_number)
)
INSERT INTO user_liked_tracks (
    user_id,
    track_id
)
SELECT
    favorite_track_seed.user_id,
    track_list.id
FROM favorite_track_seed
JOIN track_list
    ON track_list.rn = favorite_track_seed.track_number
ON CONFLICT DO NOTHING;