# Laboratorio 2 INF-343

## Integrantes

Grupo N&M's
[Nicolás Durán](https://github.com/NcolasDran) 201673513-5
[Matías Marchant](https://github.com/MatiasMarchant) 201673556-9



### Disposición Máquinas Virtuales

Dist37 ->
Dist38 ->
Dist39 ->
Dist40 ->

### Consideraciones extra

Se asume que los libros son todos con peso > 750 kB c/u


### ¿Dónde debo agregar la carpeta "servernamenode"?

Para conocer dónde se deben agregar las carpetas de paquetes propios, ejecutar en la terminal

```bash
go env
```

La dirección debería ser GOROOT\src, en mi caso GOROOT es c:\go por lo tanto, la carpeta servernamenode debería ser agregada a c:\go\src

### Para compilar .proto

```bash
protoc --go_out=plugins=grpc:carpetadestino archivoacompilar.proto
```