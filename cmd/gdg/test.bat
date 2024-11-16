@echo off
go install .

if exist ..\test (
    rmdir /s /q ..\test
)

gdg init ../test
gdg run ../test