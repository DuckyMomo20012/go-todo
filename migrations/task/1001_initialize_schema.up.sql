create table if not exists task (
  task_id uuid default gen_random_uuid(),
  title text not null,
  description text,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  constraint task_it_pk primary key (task_id)
);
