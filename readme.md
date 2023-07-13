On new project setup, run migration file migrate/migrate.go if using posgresql
In other case - do such migrations:

-- Table: public.quest_times


CREATE TABLE IF NOT EXISTS public.quest_times
(
    id bigint NOT NULL DEFAULT nextval('quest_times_id_seq'::regclass),
    "time" bigint,
    quest_id bigint,
    CONSTRAINT quest_times_pkey PRIMARY KEY (id),
    CONSTRAINT fk_quest_times_quest FOREIGN KEY (quest_id)
        REFERENCES public.quest_structures (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

----------------------------------

-- Table: public.quest_structures

CREATE TABLE IF NOT EXISTS public.quest_structures
(
    id bigint NOT NULL DEFAULT nextval('quest_structures_id_seq'::regclass),
    content text COLLATE pg_catalog."default",
    "character" text COLLATE pg_catalog."default",
    quest_reward_id bigint,
    CONSTRAINT quest_structures_pkey PRIMARY KEY (id),
    CONSTRAINT fk_quest_structures_quest_reward FOREIGN KEY (quest_reward_id)
        REFERENCES public.quest_descriptions (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

-----------------------------------
-- Table: public.quest_descriptions

CREATE TABLE IF NOT EXISTS public.quest_descriptions
(
    id bigint NOT NULL DEFAULT nextval('quest_descriptions_id_seq'::regclass),
    questgiver_name text COLLATE pg_catalog."default",
    reward_lp text COLLATE pg_catalog."default",
    reward_exp text COLLATE pg_catalog."default",
    reward_local_quality text COLLATE pg_catalog."default",
    reward_local_quality_additional text COLLATE pg_catalog."default",
    reward_by text COLLATE pg_catalog."default",
    reward_item text COLLATE pg_catalog."default",
    CONSTRAINT quest_descriptions_pkey PRIMARY KEY (id)
)

-----------------------------------

migration now automatical in DB INITIALISE.

---------
docked volume on windows : \\wsl.localhost\docker-desktop-data\data\docker\volumes\questerdeploy_db\_data
