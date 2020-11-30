package main

import (
	//"bufio"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"servernamenode"
	"io/ioutil"
	"log"
	//"math"
	"os"
	"serverdatanode"
	"strconv"
	"strings"
)

//message ChunkLibro {
//    string Nombre = 1;
//    bytes Chunk = 2;
//}

func pedir_a_DataNode1(chunk string) *serverdatanode.ChunkLibro{
	//--------------------------------------------------------------------
	// Conexion a DataNode 1
	partBuffer := make([]byte, 0)
	chunk_vacio := serverdatanode.ChunkLibro{
		Nombre: "vacio",
		Chunk:  partBuffer,
	}

	var conn_DN1 *grpc.ClientConn
	conn_DN1, err_DN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
	if err_DN1 != nil {
		fmt.Printf("Sin conexión DataNode 1\n")
		return &chunk_vacio

	} else {
		defer conn_DN1.Close()
		cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
		peticion_chunk_DN1 := serverdatanode.MensajeTest{
			Mensaje: chunk,
		}
		chunk_retorno, err_DN1 := cDataNode1.DownloaderDescargaLibro(context.Background(), &peticion_chunk_DN1)
		if err_DN1 != nil {
			fmt.Printf("> Sin respuesta DataNode 1.\n")
			return &chunk_vacio
		}		
		return chunk_retorno
	}
}

func pedir_a_DataNode2(chunk string) *serverdatanode.ChunkLibro{
	//--------------------------------------------------------------------
	// Conexion a DataNode 2
	partBuffer := make([]byte, 0)
	chunk_vacio := serverdatanode.ChunkLibro{
		Nombre: "vacio",
		Chunk:  partBuffer,
	}

	var conn_DN2 *grpc.ClientConn
	conn_DN2, err_DN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
	if err_DN2 != nil {
		fmt.Printf("Sin conexión DataNode 2\n")
		return &chunk_vacio

	} else {
		defer conn_DN2.Close()
		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
		peticion_chunk_DN2 := serverdatanode.MensajeTest{
			Mensaje: chunk,
		}
		chunk_retorno, err_DN2 := cDataNode2.DownloaderDescargaLibro(context.Background(), &peticion_chunk_DN2)
		if err_DN2 != nil {
			fmt.Printf("> Sin respuesta DataNode 2.\n")
			return &chunk_vacio
		}		
		return chunk_retorno
	}
}

func pedir_a_DataNode3(chunk string) *serverdatanode.ChunkLibro{
	//--------------------------------------------------------------------
	// Conexion a DataNode 3
	partBuffer := make([]byte, 0)
	chunk_vacio := serverdatanode.ChunkLibro{
		Nombre: "vacio",
		Chunk:  partBuffer,
	}

	var conn_DN3 *grpc.ClientConn
	conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
	if err_DN3 != nil {
		fmt.Printf("Sin conexión DataNode 3\n")
		return &chunk_vacio

	} else {
		defer conn_DN3.Close()
		cDataNode3 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
		peticion_chunk_DN3 := serverdatanode.MensajeTest{
			Mensaje: chunk,
		}
		chunk_retorno, err_DN3 := cDataNode3.DownloaderDescargaLibro(context.Background(), &peticion_chunk_DN3)
		if err_DN3 != nil {
			fmt.Printf("> Sin respuesta DataNode 3.\n")
			return &chunk_vacio
		}		
		return chunk_retorno
	}
}

func getChunksLibro(tituloLibro string, chunks string) [] serverdatanode.ChunkLibro{

	str_chunks_arr := strings.Split(chunks, "\n")
	str_chunks_arr = str_chunks_arr[:len(str_chunks_arr) - 1]

	var chunks_libro []serverdatanode.ChunkLibro

	for _, i := range str_chunks_arr{
		
		i_split := strings.Split(i, " ")
		n_chunk := i_split[0]
		maquina := i_split[1]

		chunk := tituloLibro+"_"+n_chunk
		
		if maquina == "dist37"{
			chunkLibro := pedir_a_DataNode1(chunk)
			chunks_libro = append(chunks_libro, *chunkLibro)
		} else if maquina == "dist38"{
			chunkLibro := pedir_a_DataNode2(chunk)
			chunks_libro = append(chunks_libro, *chunkLibro)
		} else if maquina == "dist39" {
			chunkLibro := pedir_a_DataNode3(chunk)
			chunks_libro = append(chunks_libro, *chunkLibro)
		} else {
			log.Fatalf("Error al contactar maquina: %s", maquina)
		}		

	}
	return chunks_libro
}

func guardarChunk(chunk serverdatanode.ChunkLibro){
	fileName := "Chunks/" + chunk.Nombre
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error al crear archivo: %s", err)
	}
	ioutil.WriteFile(fileName, chunk.Chunk, os.ModeAppend)
}

func juntarChunks(tituloLibro string, chunksLibro [] serverdatanode.ChunkLibro){
	// verificar si existen chunks vacios
	for _, chunk := range chunksLibro{
		if chunk.Nombre == "vacio"{
			fmt.Println("> Error al recuperar el archivo.\n")
			return
		}		
	}

	// asegurarse el orden de los chunks
	largo := len(chunksLibro)
	var orden [largo]int

	for i, chunk := range chunksLibro{
		nombre := chunk.Nombre
		pos_split := strings.Split(nombre, "_")
		pos := pos_split[len(pos_split)-1]
		pos_int, _ := strconv.Atoi(pos)
		orden[pos_int] = i
	}
		
	newFileName := "Descargas/"+tituloLibro+".pdf"
	_, err := os.Create(newFileName)
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
	file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
	// se itera el arreglo oden, que tiene el orden correcto de los chunks 
	for _, posicion := range orden{
		// se guarda chunk en la máquina
		chunkBufferBytes := chunksLibro[posicion]

		guardarChunk(chunkBufferBytes)

		n, err := file.Write(chunkBufferBytes.Chunk)
		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}
		file.Sync()
		fmt.Println("|Escribiendo ", n, " bytes|")
	}
	file.Close()
	fmt.Println("> Libro descargado.\n")
}


func main() {
	fmt.Printf("#### ClienteDownloader ####\n\n")

	// Conexion a un Namenode 
	var conn_NN *grpc.ClientConn
	conn_NN, err_NN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if err_NN != nil {
		log.Fatalf("Error al conectar con NameNode: %s", err_NN)
	}
	defer conn_NN.Close()

	cNameNodeNN := servernamenode.NewNameNodeServiceClient(conn_NN)

	for {
		fmt.Print("---------------------------------\n")
		fmt.Print("Ingrese una opción\n")
		fmt.Print("> 1. Solicitar listado de libros\n")
		fmt.Print("> 2. Descargar libro\n")
		fmt.Print("---------------------------------\n")
	
	
		var opcion int
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			log.Fatalf("Error al ingresar opción: %s", err)
		}
	
		
		if opcion == 1 {
			//---------------------------------------------------------------------------------------------------------------
			// Pedir listado de libros disponibles
			peticionNN := servernamenode.MensajeTest{
				Mensaje: "listadoLibros",
			}
	
			respuestaNN_listado, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &peticionNN)
			if err_NN != nil {
				fmt.Print("Error al obtener listado: %s", err_NN)
			} else {
				fmt.Print("\n----- Listado -----\n")
				fmt.Print(respuestaNN_listado.Mensaje)
				fmt.Print("--------------------\n")
			}

			//---------------------------------------------------------------------------------------------------------------
		} else if opcion == 2{
			//---------------------------------------------------------------------------------------------------------------
			// Descargar Libro
			fmt.Print("------------------------------\n")
			fmt.Print("Ingrese el número del libro que desea descargar:\n")
			var nLibro int
			_, err := fmt.Scanf("%d", &nLibro)
			if err != nil {
				log.Fatalf("Error al ingresar libro: %s", err)
			}

			// Pedir listado de libros disponibles
			mensajeNN := servernamenode.MensajeTest{
				Mensaje: "listadoLibros",
			}		
			respuestaNN, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajeNN)
			if err_NN != nil {
				fmt.Print("Error al obtener listado: %s", err_NN)
			} else {

				tituloLibro := strings.Split(strings.Split(respuestaNN.Mensaje, "\n")[nLibro-1], " ")[1]
				
				mensajeNN := servernamenode.MensajeTest{
					Mensaje: "ubicacion "+tituloLibro,
				}

				respuestaNN, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajeNN)
				if err_NN != nil {
					fmt.Print("Error al obtener respuesta de NameNode: %s", err_NN)
				} else {
					chunks := respuestaNN.Mensaje
					fmt.Println("> Recibiendo Chunks...")
					chunksLibro := getChunksLibro(tituloLibro, chunks)
					juntarChunks(tituloLibro, chunksLibro)
				}
			}

			//---------------------------------------------------------------------------------------------------------------

		} else {
			fmt.Print("Error al ingresar opción\n")
		}

	}
}
