-- +goose Up
CREATE TABLE maker.cat_file_box
(
    id         SERIAL PRIMARY KEY,
    log_id     BIGINT  NOT NULL REFERENCES public.event_logs (id) ON DELETE CASCADE,
    address_id BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    msg_sender BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    what       TEXT,
    data       NUMERIC,
    header_id  INTEGER NOT NULL REFERENCES public.headers (id) ON DELETE CASCADE,
    UNIQUE (header_id, log_id)
);

CREATE INDEX cat_file_box_log_index
    ON maker.cat_file_box (log_id);
CREATE INDEX cat_file_box_address_index
    ON maker.cat_file_box (address_id);
CREATE INDEX cat_file_box_msg_sender
    ON maker.cat_file_box (msg_sender);
CREATE INDEX cat_file_box_header_index
    ON maker.cat_file_box (header_id);

CREATE TABLE maker.cat_file_chop_lump_dunk
(
    id         SERIAL PRIMARY KEY,
    log_id     BIGINT  NOT NULL REFERENCES public.event_logs (id) ON DELETE CASCADE,
    address_id BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    msg_sender BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    what       TEXT,
    data       NUMERIC,
    header_id  INTEGER NOT NULL REFERENCES public.headers (id) ON DELETE CASCADE,
    ilk_id     INTEGER NOT NULL REFERENCES maker.ilks (id) ON DELETE CASCADE,
    UNIQUE (header_id, log_id)
);

CREATE INDEX cat_file_chop_lump_dunk_header_index
    ON maker.cat_file_chop_lump_dunk (header_id);
CREATE INDEX cat_file_chop_lump_dunk_log_index
    ON maker.cat_file_chop_lump_dunk (log_id);
CREATE INDEX cat_file_chop_lump_dunk_address_index
    ON maker.cat_file_chop_lump_dunk (address_id);
CREATE INDEX cat_file_chop_lump_dunk_msg_sender_index
    ON maker.cat_file_chop_lump_dunk (msg_sender);
CREATE INDEX cat_file_chop_lump_dunk_ilk_index
    ON maker.cat_file_chop_lump_dunk (ilk_id);

CREATE TABLE maker.cat_file_flip
(
    id         SERIAL PRIMARY KEY,
    log_id     BIGINT  NOT NULL REFERENCES public.event_logs (id) ON DELETE CASCADE,
    address_id BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    msg_sender BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    what       TEXT,
    flip       TEXT,
    header_id  INTEGER NOT NULL REFERENCES public.headers (id) ON DELETE CASCADE,
    ilk_id     INTEGER NOT NULL REFERENCES maker.ilks (id) ON DELETE CASCADE,
    UNIQUE (header_id, log_id)
);

CREATE INDEX cat_file_flip_header_index
    ON maker.cat_file_flip (header_id);
CREATE INDEX cat_file_flip_log_index
    ON maker.cat_file_flip (log_id);
CREATE INDEX cat_file_flip_address_index
    ON maker.cat_file_flip (address_id);
CREATE INDEX cat_file_flip_ilk_index
    ON maker.cat_file_flip (ilk_id);
CREATE INDEX cat_file_flip_msg_sender_index
    ON maker.cat_file_flip (msg_sender);

CREATE TABLE maker.cat_file_vow
(
    id         SERIAL PRIMARY KEY,
    log_id     BIGINT  NOT NULL REFERENCES public.event_logs (id) ON DELETE CASCADE,
    address_id BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    msg_sender BIGINT  NOT NULL REFERENCES public.addresses (id) ON DELETE CASCADE,
    what       TEXT,
    data       TEXT,
    header_id  INTEGER NOT NULL REFERENCES public.headers (id) ON DELETE CASCADE,
    UNIQUE (header_id, log_id)
);

CREATE INDEX cat_file_vow_header_index
    ON maker.cat_file_vow (header_id);
CREATE INDEX cat_file_vow_log_index
    ON maker.cat_file_vow (log_id);
CREATE INDEX cat_file_vow_address_index
    ON maker.cat_file_vow (address_id);
CREATE INDEX cat_file_vow_msg_sender
    ON maker.cat_file_vow (msg_sender);


-- +goose Down
DROP TABLE maker.cat_file_box;
DROP TABLE maker.cat_file_chop_lump_dunk;
DROP TABLE maker.cat_file_flip;
DROP TABLE maker.cat_file_vow;
