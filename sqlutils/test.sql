-- database test
DROP TABLE IF EXISTS staffs;
CREATE TABLE staffs (
    id serial primary key,
    name varchar(200),
    gender char(1),
    staff_type varchar(200),
    phone varchar(200),
    birthday date,
    created_on timestamp
);
