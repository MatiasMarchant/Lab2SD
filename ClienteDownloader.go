package main

import (
	//"bufio"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"servernamenode"
	//"io/ioutil"
	"log"
	//"math"
	//"os"
	"serverdatanode"
	//"strconv"
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

func juntarChunks(chunksLibro [] serverdatanode.ChunkLibro){
	// verificar si existen chunks vacios!
	// crear carpeta Descargas
	fmt.Println(chunksLibro)
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
		fmt.Print("---------------------------------------\n")
		fmt.Print("Ingrese una opción\n")
		fmt.Print("> 1. Solicitar listado de libros\n")
		fmt.Print("> 2. Descargar libro\n")
	
	
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
				fmt.Print("\nListado:\n")
				fmt.Print(respuestaNN_listado.Mensaje)
			}

			//---------------------------------------------------------------------------------------------------------------
		} else if opcion == 2{
			//---------------------------------------------------------------------------------------------------------------
			// Descargar Libro
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
					fmt.Println("Recibiendo Chunks...")
					chunksLibro := getChunksLibro(tituloLibro, chunks)
					juntarChunks(chunksLibro)

				}
			}

			//---------------------------------------------------------------------------------------------------------------

		} else {
			fmt.Print("Error al ingresar opción\n")
		}

	}








/*

	fmt.Print("Ingresar nombre de carpeta donde están libros\n")
	fmt.Print("> ")
	var carpeta_libros string
	_, err := fmt.Scanln(&carpeta_libros)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	carpeta_libros = strings.TrimRight(carpeta_libros, "\n")

	// https://stackoverflow.com/questions/14668850/list-directory-in-go
	files, err_files := ioutil.ReadDir("./" + carpeta_libros + "/")
	if err_files != nil {
		log.Fatal(err_files)
	}

	for indice_libro, f := range files {
		fmt.Print(indice_libro, ". ", f.Name(), "\n")
	}
	fmt.Print("---------------------------------------\n")
	fmt.Print("Ingresar numero de libro a subir\n")
	fmt.Print("> ")
	var indice_libro_a_subir string
	_, err2 := fmt.Scanln(&indice_libro_a_subir)
	if err2 != nil {
		log.Fatalf("Error al recibir variable indice_libro_a_subir: %s", err2)
	}

	integerdice_libro_a_subir, _ := strconv.Atoi(strings.TrimRight(indice_libro_a_subir, "\n"))
	fmt.Println("Libro escogido:", files[integerdice_libro_a_subir].Name())


	/*
		Hasta acá separa los libros en chunks de 250 kB y lo siguiente sería
		enviar chunks a un DataNode, pero voy a rearmarlo para probar que funcione
	*/

	/*
		newFileName := strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf") + "_reconstruido" + ".pdf" // Estos trims y + no deberian ser necesario pq despues se reconstruye en otra carpeta
		_, err5 := os.Create(newFileName)

		if err5 != nil {
			log.Fatalf("Error al crear archivo (err5): %s", err5)
		}

		// Set the newFileName file to APPEND MODE!!
		// Open files r and w

		file, err = os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

		if err != nil {
			log.Fatalf("Error al abrir archivo: %s", err)
		}

		// IMPORTANTE! do not defer a file.Close when opening a file for APPEND mode!
		// defer file.Close()

		// just information on which part of the new file we are appending
		var writePosition int64 = 0

		for j := uint64(0); j < totalPartsNum; j++ {

			// Read a chunk
			currentChunkFileName := strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf") + "_" + strconv.FormatUint(j, 10)

			newFileChunk, err := os.Open(currentChunkFileName)

			if err != nil {
				log.Fatalf("Error al abrir chunk: %s", err)
			}
			defer newFileChunk.Close()

			chunkInfo, err := newFileChunk.Stat()
			if err != nil {
				log.Fatalf("Error al obtener info de newFileChunk: %s", err)
			}

			// Calculate the bytes size of each chunk
			// We are not going to rely on previous data and constant

			var chunkSize int64 = chunkInfo.Size()
			chunkBufferBytes := make([]byte, chunkSize)

			fmt.Println("Appending at position : [", writePosition, "] bytes")
			writePosition = writePosition + chunkSize
			// Read into chunkBufferBytes
			reader := bufio.NewReader(newFileChunk)
			_, err = reader.Read(chunkBufferBytes)

			if err != nil {
				log.Fatalf("Error al leer de chunk: %s", err)
			}

			// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
			// Write/save buffer to disk
			//ioutilWriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

			n, err := file.Write(chunkBufferBytes)

			if err != nil {
				log.Fatalf("Error al escribir en chunkBufferBytes: %s", err)
			}

			file.Sync() // flush to disk

			// Free up the buffer for next cycle
			// Should not be a problem if the chunk size is small, but
			// Can be resource hogging if the chunk size is huge.
			// Also a good practice to clean up your own plate after eating

			chunkBufferBytes = nil // Reset or empty our buffer
			fmt.Println("Written ", n, " bytes")

			fmt.Println("Recombining part [", j, "] into: ", newFileName)
		}


		// Now, we close the newFileName
		file.Close()
	*/
}
