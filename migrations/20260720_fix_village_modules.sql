BEGIN;

ALTER TABLE villages
    ADD COLUMN IF NOT EXISTS region TEXT,
    ADD COLUMN IF NOT EXISTS hamlet_one TEXT,
    ADD COLUMN IF NOT EXISTS hamlet_two TEXT,
    ADD COLUMN IF NOT EXISTS north_border TEXT,
    ADD COLUMN IF NOT EXISTS east_border TEXT,
    ADD COLUMN IF NOT EXISTS south_border TEXT,
    ADD COLUMN IF NOT EXISTS west_border TEXT,
    ADD COLUMN IF NOT EXISTS area TEXT,
    ADD COLUMN IF NOT EXISTS is_active BOOLEAN NOT NULL DEFAULT FALSE;

ALTER TABLE villages
    ALTER COLUMN population TYPE TEXT USING population::TEXT;

ALTER TABLE potential_detail
    ADD COLUMN IF NOT EXISTS id BIGSERIAL;

INSERT INTO categories (name, slug, type)
SELECT 'Galeri', 'galeri', 'gallery'
WHERE NOT EXISTS (
    SELECT 1 FROM categories WHERE slug = 'galeri' AND type = 'gallery'
);

COMMIT;
