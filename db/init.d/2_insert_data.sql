USE `sample`;

INSERT INTO
  todos(`name`, `is_done`, `created_at`, `updated_at`) VALUE(
    'hoge',
    true,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  );

INSERT INTO
  todos(`name`, `is_done`, `created_at`, `updated_at`) VALUE(
    'fuga',
    false,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  );

INSERT INTO
  todos(`name`, `is_done`, `created_at`, `updated_at`) VALUE(
    'piyo',
    true,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  );