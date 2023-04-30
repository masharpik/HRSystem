CREATE TABLE enterprises (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE,
  staff_units INT NOT NULL,
  system_password VARCHAR(50) NOT NULL
);

CREATE TABLE departments (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  enterprise_id INT NOT NULL REFERENCES enterprises(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE positions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  description TEXT,
  wage numeric(10,2) NOT NULL,
  department_id INT NOT NULL REFERENCES departments(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE employees (
  id SERIAL PRIMARY KEY,
  position_id INT NOT NULL REFERENCES positions(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE hr_employees (
  id SERIAL PRIMARY KEY,
  login VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(50) NOT NULL,
  avatar TEXT,
  enterprise_id INT NOT NULL REFERENCES enterprises(id) ON UPDATE CASCADE ON DELETE CASCADE,
  approval BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  text TEXT NOT NULL,
  date DATE NOT NULL,
  hr_employee_id INT NOT NULL REFERENCES hr_employees(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE editors (
  id SERIAL PRIMARY KEY,
  login VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(50) NOT NULL,
  avatar TEXT,
  enterprise_id INT NOT NULL REFERENCES enterprises(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE career (
  id SERIAL PRIMARY KEY,
  position_ids JSON NOT NULL,
  employee_id INT NOT NULL REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE UNIQUE
);

CREATE TABLE family_compositions (
    id serial PRIMARY KEY,
    name varchar(50) NOT NULL UNIQUE
);

CREATE TABLE identities (
    id serial PRIMARY KEY,
    surname varchar(50) NOT NULL,
    name varchar(50) NOT NULL,
    middlename varchar(50),
    birth date NOT NULL,
    family_composition_id int REFERENCES family_compositions(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    employee_id int REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE UNIQUE
);

CREATE TABLE contracts (
    id serial PRIMARY KEY,
    TIN int NOT NULL UNIQUE,
    representative_id int REFERENCES employees(id) ON UPDATE CASCADE ON DELETE SET NULL,
    date_hiring date NOT NULL,
    place_hiring text NOT NULL,
    workplace text NOT NULL,
    contract_expiry date NOT NULL,
    work_mode text NOT NULL,
    additional_conditions text,
    employee_id int REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE passports (
    id serial PRIMARY KEY,
    series int NOT NULL,
    number int NOT NULL,
    ussied_by text NOT NULL,
    issue_date date NOT NULL,
    registration text NOT NULL,
    employee_id serial REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL UNIQUE
);

CREATE TABLE salary_changes (
    id serial PRIMARY KEY,
    amount numeric(10, 2) NOT NULL,
    change_date date NOT NULL DEFAULT NOW(),
    employee_id serial REFERENCES employees(id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL
);
