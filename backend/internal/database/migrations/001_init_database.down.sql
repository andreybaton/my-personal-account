-- 1. таблицы с зависимостями:
DROP TABLE IF EXISTS university.curriculum CASCADE;
DROP TABLE IF EXISTS university.grades CASCADE;
DROP TABLE IF EXISTS university.attendances CASCADE;
DROP TABLE IF EXISTS university.lessons CASCADE;
DROP TABLE IF EXISTS university.teacher_disciplines CASCADE;
DROP TABLE IF EXISTS university.students CASCADE;
DROP TABLE IF EXISTS university.teachers CASCADE;
DROP TABLE IF EXISTS university.disciplines CASCADE;
DROP TABLE IF EXISTS university.groups CASCADE;
DROP TABLE IF EXISTS university.specializations CASCADE;
DROP TABLE IF EXISTS university.departments CASCADE;
DROP TABLE IF EXISTS university.faculties CASCADE;

-- 2. таблицы m:n
DROP TABLE IF EXISTS university.user_roles CASCADE;

-- 3. простые таблицы
DROP TABLE IF EXISTS university.employees CASCADE;
DROP TABLE IF EXISTS university.classrooms CASCADE;

-- 4. базовые таблицы
DROP TABLE IF EXISTS university.users CASCADE;
DROP TABLE IF EXISTS university.roles CASCADE;
DROP TABLE IF EXISTS university.buildings CASCADE;

-- DROP SCHEMA IF EXISTS university CASCADE;