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
	//"serverdatanode"
	//"strconv"
	//"strings"
)

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
			
			mensajeNN := servernamenode.MensajeTest{
				Mensaje: "listadoLibros",
			}
	
			respuestaNN, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajeNN)
			if err_NN != nil {
				fmt.Print("Error al obtener listado: %s", err_NN)
			}
			fmt.Print("\nListado:\n")
			fmt.Print(respuestaNN.Mensaje)
	
		} else if opcion == 2{
			fmt.Print("Opción 2\n")
		} else {
			fmt.Print("Error al ingresar opción")
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
