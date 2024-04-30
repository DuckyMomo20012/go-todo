create or replace function set_updated_at_timestamp()
returns trigger as $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ language plpgsql;

create or replace trigger set_updated_at_timestamp_trigger
before update on task
for each row
execute procedure set_updated_at_timestamp();
