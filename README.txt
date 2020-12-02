########## Sistema Camiones ##########

Integrantes:
	Nicolás Durán 201673513-5
	Matías Marchant 201673556-9

Grupo: N&M's


La disposición de los roles de las máquinas virtuales y los pasos a seguir son los siguiente:

- Ingresar, en todas las máquinas, a la carpeta “Lab2SD” con el comando ‘cd Lab2SD’.
- Ejecutar el archivo Makefile en cada máquina. Antes de ejecutar un Cliente, debe estar corriendo
NameNode y al menos un Datanode, a continuación cada caso:

    Dist37 -> DataNode1 (make datanode1)
    Dist38 -> DataNode2 (make datanode2)
    Dist39 -> DataNode3 (make datanode3)
    Dist40 -> NameNode (make namenode)
    Dist40 -> ClienteUploader (make clienteuploader)
    Dist40 -> ClienteDownloader (make clientedownloader)


### Consideraciones extra

Se asume que los libros son todos con peso > 750 kB c/u

