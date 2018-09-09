OHDSI DB SQL SERVER FOR LINUX DOCKER CONTAINER
==============================================

This docker container creates a SQL SERVER for LINUX with a Database called OMOP (after 15 seconds) with the OHDSI schema.

Build with:

    docker build -t rstanleyhum/ohdsidb:sqlserver .

Start with:  

    docker run -it -e ACCEPT_EULA=Y -e SA_PASSWORD={Your Password} -p 1433:1433 -d rstanleyhum/ohdsidb:sqlserver

Enter container and connect with sql server with:

    docker exec -it <container_name|id> /bin/bash
    /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P {Your Password} -d {database}

