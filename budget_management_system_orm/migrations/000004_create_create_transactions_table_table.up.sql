CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY
  amount FLOAT
  type VARCHAR(256)
  budget_id BIGINT REFERENCES budgets(id)
);


//#// apply this migration and also add primary key to budgets if not added