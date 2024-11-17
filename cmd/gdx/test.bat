@echo off
go install .

if exist ..\test (
    rmdir /s /q ..\test
)

gdx init ../test
gdx run ../test