#!/bin/bash

# Ensure that SQL Server is up
sleep 15s 

# Create the Database called OMOP (see setup.sql)
/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P ${SA_PASSWORD} -d master -i setup.sql

# Create the tables in OMOP database with the CommonDataModel schema from the git repository
/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P ${SA_PASSWORD} -d omop -i ./CommonDataModel/Sql\ Server/OMOP\ CDM\ ddl\ -\ SQL\ Server.sql

