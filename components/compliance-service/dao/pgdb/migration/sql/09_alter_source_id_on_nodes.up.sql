ALTER TABLE IF EXISTS nodes
  ADD COLUMN IF NOT EXISTS source_id TEXT DEFAULT NULL,
  ADD COLUMN IF NOT EXISTS date_added TIMESTAMP NOT NULL DEFAULT NOW();
