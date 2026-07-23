BEGIN;

-- Update latitude and longitude to DOUBLE PRECISION (FLOAT8)
ALTER TABLE villages
    ALTER COLUMN latitude TYPE DOUBLE PRECISION USING NULLIF(latitude::TEXT, '')::DOUBLE PRECISION,
    ALTER COLUMN longitude TYPE DOUBLE PRECISION USING NULLIF(longitude::TEXT, '')::DOUBLE PRECISION;

-- Update hamlet columns to BIGINT to prevent integer overflow
ALTER TABLE villages
    ALTER COLUMN hamlet_one TYPE BIGINT USING NULLIF(hamlet_one::TEXT, '')::BIGINT,
    ALTER COLUMN hamlet_two TYPE BIGINT USING NULLIF(hamlet_two::TEXT, '')::BIGINT;

COMMIT;
