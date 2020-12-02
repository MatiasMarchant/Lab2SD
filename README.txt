########## Red de Bibliotecas ##########

Integrantes:
	Nicolás Durán 201673513-5
	Matías Marchant 201673556-9

Grupo: N&M's


La disposición de los roles de las máquinas virtuales y los pasos a seguir son los siguientes:

- Ingresar, en todas las máquinas, a la carpeta “Lab2SD” con el comando ‘cd Lab2SD’.

- Ejecutar el archivo Makefile en cada máquina. Antes de ejecutar un Cliente, debe estar corriendo NameNode 
y al menos un Datanode, a continuación cada caso:

    Dist37 -> DataNode1 (ejecutar 'make datanode1')
    Dist38 -> DataNode2 (ejecutar 'make datanode2')
    Dist39 -> DataNode3 (ejecutar 'make datanode3')
    Dist40 -> NameNode (ejecutar 'make namenode')
    Dist40 -> ClienteUploader (ejecutar 'make clienteuploader')
    Dist40 -> ClienteDownloader (ejecutar 'make clientedownloader')

- Al ejecutar un DataNode se debe elegir el algoritmo que se desea correr, una vez que se ingresa la preferencia el 
DataNode podrá recibir mensajes.

- Los DataNodes almacenan los Chunks en la carpeta "Chunks", se puede ingresar con ‘cd Chunks', en la máquina correspondiente.

- El ClienteDownloader almacena los Chunks en la carpeta "Chunks", se puede ingresar con ‘cd Chunks'; los libros descargados 
por el ClienteDownloader se almacenan en la carpeta "Descargas", se puede ingresar con ‘cd Chunks', en la máquina Dist40.

- Los libros que sube el ClienteUploader se encuentran en la carpeta "Libros", se puede ingresar con ‘cd Chunks'.

- Si se desean eliminar los archivos de las carpetas "Chunks", se debe ejecutar 'make cleanchunks' en la máquina deseada.
- Si se desean eliminar los archivos de la Descarga, en la máquina Dist40, se debe ejecutar 'make cleandescargas'.
- Si se desea eliminar el archivo 'log.txt', en la máquina Dist40, se debe ejecutar 'make cleanlog'.

- Para cerrar una ejecución ingresar en la consola 'Ctrl + C'.

### Consideraciones extra ###

- Se asume que los libros son todos con peso > 750 kB c/u
- La estructura del archivo 'log.txt' es:

    Nombre_Libro_1 Cantidad_Partes_1
    Valor_de_parte dist_de_Máquina
    ...
    Nombre_Libro_2 Cantidad_Partes_2
    Valor_de_parte dist_de_Máquina

Por ejemplo:

    Dracula-Stoker_Bram 7
    0 dist37
    1 dist37
    2 dist37
    ...
