-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Extend ilk_state with frob_events
CREATE FUNCTION api.ilk_state_frobs(state api.ilk_state)
  RETURNS SETOF api.frob_event AS
$$
  SELECT * FROM api.all_frobs(state.ilk_name)
  WHERE block_height <= state.block_height
$$ LANGUAGE sql STABLE;


-- Extend ilk_state with file events
CREATE FUNCTION api.ilk_state_files(state api.ilk_state)
  RETURNS SETOF api.file_event AS
$$
  SELECT * FROM api.ilk_files(state.ilk_name)
  WHERE block_height <= state.block_height
$$ LANGUAGE sql STABLE;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP FUNCTION api.ilk_state_frobs(api.ilk_state);
DROP FUNCTION api.ilk_state_files(api.ilk_state);