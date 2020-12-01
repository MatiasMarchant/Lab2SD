package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math"
	"os"
	"serverdatanode"
	"strconv"
	"strings"
	"math/rand"
    "time"
)

func enviar_a_DataNode1(mensaje_cliente string) bool {
	//--------------------------------------------------------------------
	// Conexion a DataNode 1
	var conn_DN1 *grpc.ClientConn
	conn_DN1, err_DN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
	flag := true
	if err_DN1 != nil {
		flag = false
	} else {
		defer conn_DN1.Close()
		cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
		mensajetest_1 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}
		_, err_DN1 := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_1)
		// fmt.Printf(">>> Mensaje enviado\n")
		if err_DN1 != nil {
			flag = false
		} else {
			flag = true
		}
	}
	return flag
}

func enviar_a_DataNode2(mensaje_cliente string) bool {
	//--------------------------------------------------------------------
	// Conexion a DataNode 2
	var conn_DN2 *grpc.ClientConn
	conn_DN2, err_DN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
	flag := true
	if err_DN2 != nil {
		flag = false
	} else {
		defer conn_DN2.Close()
		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
		mensajetest_DN2 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}
		_, err_DN2 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)
		// fmt.Printf(">>> Mensaje enviado\n")
		if err_DN2 != nil {
			flag = false
		} else {
			flag = true
		}
	}
	return flag
}


func enviar_a_DataNode3(mensaje_cliente string) bool{
	//--------------------------------------------------------------------
	// Conexion a DataNode 3
	var conn_DN3 *grpc.ClientConn
	conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
	flag := true
	if err_DN3 != nil {
		flag = false
	} else {
		defer conn_DN3.Close()
		cDataNode3 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
		mensajetest_3 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}
		_, err_DN3 := cDataNode3.EnvioMensajeTest(context.Background(), &mensajetest_3)
		// fmt.Printf(">>> Mensaje enviado\n")
		if err_DN3 != nil {
			flag = false
		} else {
			flag = true
		}
	}
	return flag
}


func main() {
	fmt.Printf("#### ClienteUploader ####\n\n")
	

	fmt.Print("\nIngresar nombre de carpeta donde están libros\n")
	fmt.Print("> ")
	var carpeta_libros string
	_, err_scan := fmt.Scanln(&carpeta_libros)

	for {
		fmt.Print("-------- Libros --------\n")
		if err_scan != nil {
			fmt.Fprintln(os.Stderr, err_scan)
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

		// https://www.socketloop.com/tutorials/golang-recombine-chunked-files-example
		fileToBeChunked := "./" + carpeta_libros + "/" + files[integerdice_libro_a_subir].Name()
		file, err3 := os.Open(fileToBeChunked)
		if err3 != nil {
			log.Fatalf("Error al abrir archivo (err3): %s", err3)
		}
		defer file.Close()
		fileInfo, _ := file.Stat()
		var fileSize int64 = fileInfo.Size()
		const fileChunk = 250 * (1 << 10) // (1 << 10 equivale a 1 Kilobyte), multiplicado 250 veces
		//const fileChunk = 250000 // 250000 bytes segun propiedades windows, README

		// Calculate total number of parts the file will be chunked into

		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		fmt.Printf("Dividiento en %d partes.\n", totalPartsNum)

		// Conexion a un datanote random
		var dist [3]string
		dist[0] = "dist37:9001"
		dist[1] = "dist38:9002"
		dist[2] = "dist39:9003"

		errDN1 := enviar_a_DataNode1("Cliente pregunta estas vivo?\n")
		flagDN1vivo := true
		if errDN1 != true {
			flagDN1vivo = false
		}	
		errDN2 := enviar_a_DataNode2("Cliente pregunta estas vivo?\n")
		flagDN2vivo := true
		if errDN2 != true {
			flagDN2vivo = false
		}	
		errDN3 := enviar_a_DataNode3("Cliente pregunta estas vivo?\n")
		flagDN3vivo := true
		if errDN3 != true {
			flagDN3vivo = false
		}

		valor_random := -1
		
		for{
			s := rand.NewSource(time.Now().UnixNano())
			random := rand.New(s)
			valor_random = random.Intn(3)
			if valor_random == 0 && flagDN1vivo{
				break
			} 
			if valor_random == 1 && flagDN2vivo{
				break
			} 
			if valor_random == 2 && flagDN3vivo{
				break
			} 

		}

		var conn_DN *grpc.ClientConn
		conn_DN, err_DN := grpc.Dial(dist[valor_random], grpc.WithInsecure())
		if err_DN != nil {
			log.Fatalf("Error al conectar cliente uploader al DataNode: %s", err_DN)
		}
		defer conn_DN.Close()
		cDataNode := serverdatanode.NewDataNodeServiceClient(conn_DN)

		for i := uint64(0); i < totalPartsNum; i++ {
			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			// Write to disk
			fileName := strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf") + "_" + strconv.FormatUint(i, 10)

			fmt.Println("> Dividiendo en: ", fileName)

			ChunkLibro := serverdatanode.ChunkLibro{
				Nombre: fileName,
				Chunk:  partBuffer,
			}
			respuesta, _ := cDataNode.UploaderSubeLibro(context.Background(), &ChunkLibro)
			// fmt.Printf(">>> Mensaje enviado\n")
			fmt.Println(respuesta.Mensaje)

		}

		// Una vez termina de enviar un libro, le comunica al DN que terminó
		// para que así el DN proceda a hacer la propuesta.
		// Asume que el DN esta corriendo
		_, err := cDataNode.UploaderTerminoDeSubirLibro(context.Background(), &serverdatanode.MensajeTest{Mensaje: strings.TrimRight(files[integerdice_libro_a_subir].Name(), ".pdf")})
		// fmt.Printf(">>> Mensaje enviado\n")
		if err != nil {
			fmt.Printf("Error al avisar al DN que se subio libro %s: %v\n", files[integerdice_libro_a_subir], err)
		}
	}

}
