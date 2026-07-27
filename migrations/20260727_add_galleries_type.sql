BEGIN;

ALTER TABLE galleries
    ADD COLUMN IF NOT EXISTS type VARCHAR(50) NOT NULL DEFAULT 'gallery';

UPDATE galleries g
SET type = c.type
FROM categories c
WHERE g.category_id = c.id
    AND c.type IN ('dashboard', 'gallery');

CREATE INDEX IF NOT EXISTS idx_galleries_type ON galleries(type);

COMMIT;
