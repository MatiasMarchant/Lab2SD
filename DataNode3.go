package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"serverdatanode"
	"servernamenode"
	//"strconv"
	"os"
	"strings"
	"math/rand"
    "time"
)

func enviar_a_NameNode(mensaje_cliente string) {
	//--------------------------------------------------------------------
	// Conexion a NameNode
	var conn_NN *grpc.ClientConn
	conn_NN, err_NN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if err_NN != nil {
		fmt.Printf("¡Sin conexión NameNode!\n")
	} else {

		defer conn_NN.Close()

		cNameNodeNN := servernamenode.NewNameNodeServiceClient(conn_NN)

		mensajetestNN := servernamenode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		_, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajetestNN)

		if err_NN != nil {
			fmt.Printf("> Sin respuesta NameNode.\n")
		} 

	}
}

func enviar_a_DataNode1(mensaje_cliente string) bool {
	//--------------------------------------------------------------------
	// Conexion a DataNode 1
	var conn_DN1 *grpc.ClientConn
	conn_DN1, err_DN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
	flag := true
	if err_DN1 != nil {
		fmt.Printf("¡Sin conexión DataNode 1!\n")
		flag = false
	} else {

		defer conn_DN1.Close()

		cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
		mensajetest_1 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		_, err_DN1 := cDataNode1.EnvioMensajeTest(context.Background(), &mensajetest_1)

		if err_DN1 != nil {
			fmt.Printf("> Sin respuesta DataNode1.\n")
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
		fmt.Printf("¡Sin conexión DataNode 2!\n")
		flag = false
	} else {
		defer conn_DN2.Close()

		cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
		mensajetest_DN2 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		_, err_DN2 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)

		if err_DN2 != nil {
			fmt.Printf("> Sin respuesta DataNode2.\n")
			flag = false
		} else {
			flag = true
		}

	}
	return flag
}

func Enviar_Propuesta(propuesta serverdatanode.Propuesta, destinatario string) bool {
	// Transformacion propuesta para que sea enviable por grpc

	Propuesta_grpc := serverdatanode.Propuestagrpc{
		NombreLibroSubido: propuesta.NombreLibroSubido,
		PartesDN1:         strings.Join(propuesta.PartesDN1, ","),
		PartesDN2:         strings.Join(propuesta.PartesDN2, ","),
		PartesDN3:         strings.Join(propuesta.PartesDN3, ","),
	}

	switch destinatario {
	case "DataNode1":
		// Caso DataNode1
		// Conexion a DataNode 1
		var conn_DN1 *grpc.ClientConn
		conn_DN1, err_DN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
		if err_DN1 != nil {
			fmt.Printf("¡Sin conexión DataNode 1!\n")
			return false
		} else {
			defer conn_DN1.Close()

			cDataNode1 := serverdatanode.NewDataNodeServiceClient(conn_DN1)
			// Enviar propuesta por gRPC
			respuesta_DN1, err_DN1 := cDataNode1.Propuesta_Distribuido(context.Background(), &Propuesta_grpc)

			if err_DN1 != nil {
				fmt.Printf("> Error al enviar propuesta.\n")
				return false
			}

			return respuesta_DN1.Booleano
		}

	case "DataNode2":
		// Caso DataNode2
		// Conexion a DataNode 2
		var conn_DN2 *grpc.ClientConn
		conn_DN2, err_DN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
		if err_DN2 != nil {
			fmt.Printf("¡Sin conexión DataNode 2!\n")
			return false
		} else {
			defer conn_DN2.Close()

			cDataNode2 := serverdatanode.NewDataNodeServiceClient(conn_DN2)
			// Enviar propuesta por gRPC
			respuesta_DN2, err_DN2 := cDataNode2.Propuesta_Distribuido(context.Background(), &Propuesta_grpc)

			if err_DN2 != nil {
				fmt.Printf("> Error al enviar propuesta.\n")
				return false

			}

			return respuesta_DN2.Booleano
		}

	case "DataNode3":
		// Caso DataNode3
		// Conexion a DataNode 3
		var conn_DN3 *grpc.ClientConn
		conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
		if err_DN3 != nil {
			fmt.Printf("¡Sin conexión DataNode 3!\n")
			return false
		} else {
			defer conn_DN3.Close()

			cDataNode3 := serverdatanode.NewDataNodeServiceClient(conn_DN3)

			respuesta_DN3, err_DN3 := cDataNode3.Propuesta_Distribuido(context.Background(), &Propuesta_grpc)

			if err_DN3 != nil {
				fmt.Printf("> Error al enviar propuesta.\n")
				return false
			}

			return respuesta_DN3.Booleano

		}
	}
	return false
}

func Enviar_Propuesta_NameNode(propuesta servernamenode.Propuestagrpc) servernamenode.Propuestagrpc{

	// Caso NameNode (para centralizado)
	// Conexion a NameNode
	var conn_NN *grpc.ClientConn
	conn_NN, err_NN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if err_NN != nil {
		log.Fatalf("¡Sin conexión con NameNode!\n")	
	} 
	defer conn_NN.Close()

	cNameNode := servernamenode.NewNameNodeServiceClient(conn_NN)
	// Enviar propuesta por gRPC
	respuesta_NN, err_NN := cNameNode.Propuesta_Centralizado(context.Background(), &propuesta)

	if err_NN != nil {
		log.Fatalf("> Error al enviar propuesta a NameNode.\n")			
	}
	return *respuesta_NN
}

func EscribirEnLog(Propuesta serverdatanode.Propuesta, ID int, cant_partes int) {
	// Llamar funcion escritura sobre log namenode
	// Usar Ricart-Agrawala(Pedir acceso a lugar critico)
	// Si todos responden -> Escribir con gRPC

	// Estructura
	// -------------------------------- log.txt
	// Nombre_Libro_1 Cantidad_Partes_1
	// parte_1_1 ip_maquina
	// parte_1_2 ip_maquina
	// ..
	// Nombre_Libro_2 Cantidad_Partes_2
	// parte_2_1 ip_maquina
	// ..
	// ----------------------------------------

	// Conexion NN
	var connNN *grpc.ClientConn
	connNN, errNN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if errNN != nil {
		fmt.Println("¡Sin conexión NameNode!\n")
	} else {
		defer connNN.Close()

		cNameNodeNN := servernamenode.NewNameNodeServiceClient(connNN)
		mensaje := servernamenode.EscrituraLog{
			NombreLibro: Propuesta.NombreLibroSubido,
			CantPartes:  int32(cant_partes),
			PartesDN1:   strings.Join(Propuesta.PartesDN1, ","),
			PartesDN2:   strings.Join(Propuesta.PartesDN2, ","),
			PartesDN3:   strings.Join(Propuesta.PartesDN3, ","),
		}
		cNameNodeNN.EscribirEnLog(context.Background(), &mensaje)
	}
}

func EscribirEnLog_Centralizado(Propuesta servernamenode.Propuestagrpc, ID int, cant_partes int) {

	// Conexion NN
	var connNN *grpc.ClientConn
	connNN, errNN := grpc.Dial("dist40:9000", grpc.WithInsecure())
	if errNN != nil {
		fmt.Println("¡Sin conexión NameNode!\n")
	} else {
		defer connNN.Close()

		cNameNodeNN := servernamenode.NewNameNodeServiceClient(connNN)
		mensaje := servernamenode.EscrituraLog{
			NombreLibro: Propuesta.NombreLibroSubido,
			CantPartes:  int32(cant_partes),
			PartesDN1:  Propuesta.PartesDN1,
			PartesDN2:  Propuesta.PartesDN2,
			PartesDN3:  Propuesta.PartesDN3,
		}
		cNameNodeNN.EscribirEnLog(context.Background(), &mensaje)
	}
}

func EnviarChunks(Propuesta serverdatanode.Propuesta) {
	for _, indicechunk := range Propuesta.PartesDN1 {
		ChunkFileName := indicechunk
		fmt.Printf("# Enviando chunk a DN1: %s", ChunkFileName+"\n")
		newFileChunk, err := os.Open(ChunkFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer newFileChunk.Close()
		chunkInfo, err := newFileChunk.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)
		newFileChunk.Read(chunkBufferBytes)

		ChunkLibro := serverdatanode.ChunkLibro{
			Nombre: ChunkFileName,
			Chunk:  chunkBufferBytes,
		}

		// Conexion
		var connDN1 *grpc.ClientConn
		connDN1, errDN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
		if errDN1 != nil {
			log.Fatalf("Error al enviar chunk: %v", errDN1)
		}
		defer connDN1.Close()
		cDataNode1 := serverdatanode.NewDataNodeServiceClient(connDN1)

		// gRPC
		cDataNode1.UploaderSubeLibro(context.Background(), &ChunkLibro)

	}
	for _, indicechunk := range Propuesta.PartesDN2 {
		ChunkFileName := indicechunk
		fmt.Printf("# Enviando chunk a DN2: %s", ChunkFileName+"\n")
		newFileChunk, err := os.Open(ChunkFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer newFileChunk.Close()
		chunkInfo, err := newFileChunk.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)
		newFileChunk.Read(chunkBufferBytes)

		ChunkLibro := serverdatanode.ChunkLibro{
			Nombre: ChunkFileName,
			Chunk:  chunkBufferBytes,
		}

		// Conexion
		var connDN2 *grpc.ClientConn
		connDN2, errDN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
		if errDN2 != nil {
			log.Fatalf("Error al enviar chunk: %v", errDN2)
		}
		defer connDN2.Close()
		cDataNode2 := serverdatanode.NewDataNodeServiceClient(connDN2)

		// gRPC
		cDataNode2.UploaderSubeLibro(context.Background(), &ChunkLibro)
	}
	for _, indicechunk := range Propuesta.PartesDN3 {
		ChunkFileName := indicechunk
		fmt.Printf("# Enviando chunk a DN3: %s", ChunkFileName+"\n")
		newFileChunk, err := os.Open(ChunkFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer newFileChunk.Close()
		chunkInfo, err := newFileChunk.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)
		newFileChunk.Read(chunkBufferBytes)

		ChunkLibro := serverdatanode.ChunkLibro{
			Nombre: ChunkFileName,
			Chunk:  chunkBufferBytes,
		}

		// Conexion
		var connDN3 *grpc.ClientConn
		connDN3, errDN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
		if errDN3 != nil {
			log.Fatalf("Error al enviar chunk: %v", errDN3)
		}
		defer connDN3.Close()
		cDataNode3 := serverdatanode.NewDataNodeServiceClient(connDN3)

		// gRPC
		cDataNode3.UploaderSubeLibro(context.Background(), &ChunkLibro)
	}
}

func EnviarChunks_Centralizado(Propuesta servernamenode.Propuestagrpc) {
	
	PropuestasPartesDN1 := strings.Split(Propuesta.PartesDN1, ",")
	PropuestasPartesDN2 := strings.Split(Propuesta.PartesDN2, ",")
	PropuestasPartesDN3 := strings.Split(Propuesta.PartesDN3, ",")

	for _, indicechunk := range PropuestasPartesDN1 {
		ChunkFileName := indicechunk
		fmt.Printf("# Enviando chunk a DN1: %s", ChunkFileName+"\n")
		newFileChunk, err := os.Open(ChunkFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer newFileChunk.Close()
		chunkInfo, err := newFileChunk.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)
		newFileChunk.Read(chunkBufferBytes)

		ChunkLibro := serverdatanode.ChunkLibro{
			Nombre: ChunkFileName,
			Chunk:  chunkBufferBytes,
		}

		// Conexion
		var connDN1 *grpc.ClientConn
		connDN1, errDN1 := grpc.Dial("dist37:9001", grpc.WithInsecure())
		if errDN1 != nil {
			log.Fatalf("Error al enviar chunk: %v", errDN1)
		}
		defer connDN1.Close()
		cDataNode1 := serverdatanode.NewDataNodeServiceClient(connDN1)

		// gRPC
		cDataNode1.UploaderSubeLibro(context.Background(), &ChunkLibro)

	}
	for _, indicechunk := range PropuestasPartesDN2 {
		ChunkFileName := indicechunk
		fmt.Printf("# Enviando chunk a DN2: %s", ChunkFileName+"\n")
		newFileChunk, err := os.Open(ChunkFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer newFileChunk.Close()
		chunkInfo, err := newFileChunk.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)
		newFileChunk.Read(chunkBufferBytes)

		ChunkLibro := serverdatanode.ChunkLibro{
			Nombre: ChunkFileName,
			Chunk:  chunkBufferBytes,
		}

		// Conexion
		var connDN2 *grpc.ClientConn
		connDN2, errDN2 := grpc.Dial("dist38:9002", grpc.WithInsecure())
		if errDN2 != nil {
			log.Fatalf("Error al enviar chunk: %v", errDN2)
		}
		defer connDN2.Close()
		cDataNode2 := serverdatanode.NewDataNodeServiceClient(connDN2)

		// gRPC
		cDataNode2.UploaderSubeLibro(context.Background(), &ChunkLibro)
	}
	for _, indicechunk := range PropuestasPartesDN3 {
		ChunkFileName := indicechunk
		fmt.Printf("# Enviando chunk a DN3: %s", ChunkFileName+"\n")
		newFileChunk, err := os.Open(ChunkFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer newFileChunk.Close()
		chunkInfo, err := newFileChunk.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)
		newFileChunk.Read(chunkBufferBytes)

		ChunkLibro := serverdatanode.ChunkLibro{
			Nombre: ChunkFileName,
			Chunk:  chunkBufferBytes,
		}

		// Conexion
		var connDN3 *grpc.ClientConn
		connDN3, errDN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
		if errDN3 != nil {
			log.Fatalf("Error al enviar chunk: %v", errDN3)
		}
		defer connDN3.Close()
		cDataNode3 := serverdatanode.NewDataNodeServiceClient(connDN3)

		// gRPC
		cDataNode3.UploaderSubeLibro(context.Background(), &ChunkLibro)
	}
}

func HacerPropuesta(metodo string, NombreLibroSubido string) {
	var Arreglo_indices_partes_libro []string // Guarda los nombres de los chunks en el directorio
	//-------------------------------------------------------------------------------------------------------------------------
	fmt.Print("\n---------------------------------")
	if metodo == "distribuido" {
		fmt.Print("\n# Algoritmo Distribuido #\n\n")
		// Enviar mensajes a datanodes para ver si están vivos
		err := enviar_a_DataNode1("DataNode3 pregunta estas vivo?\n")
		flagDN1vivo := true
		if err != true {
			fmt.Printf("DataNode1 no está vivo\n")
			flagDN1vivo = false
		}
		err = enviar_a_DataNode2("DataNode3 pregunta estas vivo?\n")
		flagDN2vivo := true
		if err != true {
			fmt.Printf("DataNode2 no está vivo\n")
			flagDN2vivo = false
		}

		//fmt.Println("flagDN2: %v", flagDN2vivo)
		//fmt.Println("flagDN3: %v", flagDN3vivo)

		// Contar cantidad de partes del libro
		files, err_files := ioutil.ReadDir("./")
		if err_files != nil {
			log.Printf("err_files, no puede leer directorio: %v", err_files)
		}
		for _, f := range files {
			//fmt.Printf("Nombre scan: %s\n", f.Name())
			//fmt.Printf("NombreLibroSubido: %s\n", NombreLibroSubido)
			if strings.Contains(f.Name(), NombreLibroSubido) {
				Arreglo_indices_partes_libro = append(Arreglo_indices_partes_libro, f.Name())
			}
		}

		fmt.Printf("> Partes a repartir:\n")
		fmt.Println(Arreglo_indices_partes_libro)
		fmt.Printf("\n")

		// Ver cuales hay vivos y repartir con serverdatanode.Propuesta
		aprobado := false
		var PartesDN1 []string
		var PartesDN2 []string
		var PartesDN3 []string
		Propuesta := serverdatanode.Propuesta{
			NombreLibroSubido: NombreLibroSubido,
			PartesDN1:         PartesDN1,
			PartesDN2:         PartesDN2,
			PartesDN3:         PartesDN3,
		}
		for true {
			Propuesta.PartesDN1 = []string{}
			Propuesta.PartesDN2 = []string{}
			Propuesta.PartesDN3 = []string{}
			respuesta_propuesta_DN1 := false
			respuesta_propuesta_DN2 := false
			Arreglo_copia := Arreglo_indices_partes_libro

			Propuesta.PartesDN3 = append(Propuesta.PartesDN3, Arreglo_copia[len(Arreglo_copia)-1])
			i := len(Arreglo_copia) - 1
			// Borrar elemento
			Arreglo_copia[i] = Arreglo_copia[len(Arreglo_copia)-1]
			Arreglo_copia[len(Arreglo_copia)-1] = ""
			Arreglo_copia = Arreglo_copia[:len(Arreglo_copia)-1]

			if flagDN1vivo == true {
				Propuesta.PartesDN1 = append(Propuesta.PartesDN1, Arreglo_copia[len(Arreglo_copia)-1])
				// Borrar elemento
				i = len(Arreglo_copia) - 1
				Arreglo_copia[i] = Arreglo_copia[len(Arreglo_copia)-1]
				Arreglo_copia[len(Arreglo_copia)-1] = ""
				Arreglo_copia = Arreglo_copia[:len(Arreglo_copia)-1]
			}
			if flagDN2vivo == true {
				Propuesta.PartesDN2 = append(Propuesta.PartesDN2, Arreglo_copia[len(Arreglo_copia)-1])
				// Borrar elemento
				i = len(Arreglo_copia) - 1
				Arreglo_copia[i] = Arreglo_copia[len(Arreglo_copia)-1]
				Arreglo_copia[len(Arreglo_copia)-1] = ""
				Arreglo_copia = Arreglo_copia[:len(Arreglo_copia)-1]
			}
	
			// El resto de chunks se quedan en este datanode
			for _, elemento := range Arreglo_copia {
				Propuesta.PartesDN1 = append(Propuesta.PartesDN1, elemento)
			}

			fmt.Printf("La \"propuesta\" quedo:\n")
			fmt.Println("- Propuesta a DN1:", Propuesta.PartesDN1)
			fmt.Println("- Propuesta a DN2:", Propuesta.PartesDN2)
			fmt.Println("- Propuesta a DN3:", Propuesta.PartesDN3)

			// Envio de propuesta por casos
			// Caso 1: DN3 y DN1 vivos
			if (flagDN1vivo == true) && (flagDN2vivo == false) {
				respuesta_propuesta_DN1 = Enviar_Propuesta(Propuesta, "DataNode1")
				respuesta_propuesta_DN2 = true
			}
			// Caso 2: DN1, DN2 y DN3 vivos
			if (flagDN1vivo == true) && (flagDN2vivo == true) {
				respuesta_propuesta_DN1 = Enviar_Propuesta(Propuesta, "DataNode1")
				respuesta_propuesta_DN2 = Enviar_Propuesta(Propuesta, "DataNode2")
			}
			// Caso 3: DN3 y DN2 vivos
			if (flagDN1vivo == false) && (flagDN2vivo == true) {
				respuesta_propuesta_DN1 = true
				respuesta_propuesta_DN2 = Enviar_Propuesta(Propuesta, "DataNode2")
			}
			// Caso 4: Solo DN3 vivo
			if (flagDN1vivo == false) && (flagDN2vivo == false) {
				respuesta_propuesta_DN1 = true
				respuesta_propuesta_DN2 = true
			}

			aprobado = respuesta_propuesta_DN1 && respuesta_propuesta_DN2
			if aprobado == true {
				break
			}
		}

		// Si llega aca, entonces aprobado == true y se documenta en el registro del NameNode

		// Llamar funcion escritura sobre log namenode
		// Si todos responden -> Escribir con gRPC
		ID := 1
		EscribirEnLog(Propuesta, ID, len(Arreglo_indices_partes_libro))

		// Enviar chunks a otros DataNode
		EnviarChunks(Propuesta)

		//
	} else if metodo == "centralizado" { // Centralizado
		//-------------------------------------------------------------------------------------------------------------------------
		fmt.Print("\n# Algoritmo Centralizado #\n\n")
		// Contar cantidad de partes del libro
		files, err_files := ioutil.ReadDir("./")
		if err_files != nil {
			log.Printf("err_files, no puede leer directorio: %v", err_files)
		}
		for _, f := range files {
			if strings.Contains(f.Name(), NombreLibroSubido) {
				Arreglo_indices_partes_libro = append(Arreglo_indices_partes_libro, f.Name())
			}
		}

		fmt.Printf("> Partes a repartir:\n")
		fmt.Println(Arreglo_indices_partes_libro)


		var PartesDN1 []string
		var PartesDN2 []string
		var PartesDN3 []string

		Propuesta := servernamenode.Propuesta{
			NombreLibroSubido: NombreLibroSubido,
			PartesDN1:         PartesDN1,
			PartesDN2:         PartesDN2,
			PartesDN3:         PartesDN3,
		}

		Propuesta.PartesDN1 = []string{}
		Propuesta.PartesDN2 = []string{}
		Propuesta.PartesDN3 = []string{}

		// Diferente para cada DataNodo
		for i, nombre_chunk := range Arreglo_indices_partes_libro{
			if i == 0 {
				Propuesta.PartesDN3 = append(Propuesta.PartesDN3, nombre_chunk)
			} else if i == 1{
				Propuesta.PartesDN1 = append(Propuesta.PartesDN1, nombre_chunk)
			} else if i == 2 {
				Propuesta.PartesDN2 = append(Propuesta.PartesDN2, nombre_chunk)
			} else {
				// asignación al azar
				s := rand.NewSource(time.Now().UnixNano())
				random := rand.New(s)
				valor_random := random.Intn(3)
				if valor_random == 0 {
					Propuesta.PartesDN1 = append(Propuesta.PartesDN1, nombre_chunk)
				} else if valor_random == 1{
					Propuesta.PartesDN2 = append(Propuesta.PartesDN2, nombre_chunk)
				} else if valor_random == 2 {
					Propuesta.PartesDN3 = append(Propuesta.PartesDN3, nombre_chunk)
				} else {
					log.Fatalf("Error en random")
				}
			}
		}

		Propuesta_grpc := servernamenode.Propuestagrpc{
			NombreLibroSubido: Propuesta.NombreLibroSubido,
			PartesDN1: strings.Join(Propuesta.PartesDN1, ","),
			PartesDN2: strings.Join(Propuesta.PartesDN2, ","),
			PartesDN3: strings.Join(Propuesta.PartesDN3, ","),
		}

		// respuesta_propuesta_NN es una propuesta
		// si la propuesta enviada se aprueba, entonces respuesta_propuesta_NN = Propuesta
		// si no se aprueba la propuesta enviada, respuesta_propuesta_NN es la propuesta de NameNode
		respuesta_propuesta_NN := Enviar_Propuesta_NameNode(Propuesta_grpc)

		fmt.Printf("La \"propuesta\" quedo:\n")
		fmt.Printf("Nombre libro: %s\n", NombreLibroSubido)
		fmt.Println("PartesDN1: %v", respuesta_propuesta_NN.PartesDN1)
		fmt.Println("PartesDN2: %v", respuesta_propuesta_NN.PartesDN2)
		fmt.Println("PartesDN3: %v", respuesta_propuesta_NN.PartesDN3)

		ID := 1
		EscribirEnLog_Centralizado(respuesta_propuesta_NN, ID, len(Arreglo_indices_partes_libro))

		// Enviar chunks a otros DataNode
		EnviarChunks_Centralizado(respuesta_propuesta_NN)

	} else {
		log.Fatalf("Error en metodo")
	}
	

}


func main() {
	fmt.Printf("#### DataNode 3 ####\n\n")

	fmt.Print("---------------------------------\n")
	fmt.Print("Ingrese el algoritmo de Exclusión Mutua que desea ejecutar:\n")
	fmt.Print("> 1. Distribuido\n")
	fmt.Print("> 2. Centralizado\n")
	fmt.Print("---------------------------------\n")
	var algoritmo int
	eleccion_algoritmo := ""
	_, err := fmt.Scanf("%d", &algoritmo)
	if err != nil {
		log.Fatalf("Error al ingresar opción: %s", err)
	}
	if algoritmo == 1{
		eleccion_algoritmo = "distribuido"
	} else if algoritmo == 2{
		eleccion_algoritmo = "centralizado"
	} else {
		log.Fatalf("Error al ingresar opción: %s", algoritmo)
	}

	//##########################################################################
	// Escucha en el puerto 9003
	lis, err_s := net.Listen("tcp", ":9003")
	if err_s != nil {
		log.Fatalf("Error en DataNode 3 al escuchar en puerto 9003: %v", err_s)
	}
	fmt.Println("DataNode 3 escuchando en puerto 9003.\n")
	s := serverdatanode.Server{}

	s.FlagLibroSubido = false

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)

	go func() {
		for {
			if s.FlagLibroSubido == true {
				// Llamar funcion
				NombreLibroSubido := s.NombreLibroSubido // Va sin ".pdf"
				s.FlagLibroSubido = false
				HacerPropuesta(eleccion_algoritmo, NombreLibroSubido) 
			}
		}
	}()

	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 3 en servidor gRPC en el puerto 9003: %v", err_s)
	}

}
