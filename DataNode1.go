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
	"strings"

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

		respuestaNN, err_NN := cNameNodeNN.EnvioMensajeTest(context.Background(), &mensajetestNN)

		if err_NN != nil {
			fmt.Printf("> Sin respuesta NameNode.\n")
		} else {
			fmt.Printf("|Cliente| NameNode responde : %s", respuestaNN.Mensaje)
		}

	}
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

		respuesta_DN2, err_DN2 := cDataNode2.EnvioMensajeTest(context.Background(), &mensajetest_DN2)

		if err_DN2 != nil {
			fmt.Printf("> Sin respuesta DataNode2.\n")
			flag = false
		} else {
			fmt.Printf("|Cliente| DataNode 2 responde: %s", respuesta_DN2.Mensaje)
			flag = true
		}

	}
	return flag
}

func enviar_a_DataNode3(mensaje_cliente string) bool {
	//--------------------------------------------------------------------
	// Conexion a DataNode 3
	var conn_DN3 *grpc.ClientConn
	conn_DN3, err_DN3 := grpc.Dial("dist39:9003", grpc.WithInsecure())
	flag := true
	if err_DN3 != nil {
		fmt.Printf("¡Sin conexión DataNode 3!\n")
		flag = false
	} else {

		defer conn_DN3.Close()

		cDataNode3 := serverdatanode.NewDataNodeServiceClient(conn_DN3)
		mensajetest_3 := serverdatanode.MensajeTest{
			Mensaje: mensaje_cliente,
		}

		respuesta_3, err_DN3 := cDataNode3.EnvioMensajeTest(context.Background(), &mensajetest_3)

		if err_DN3 != nil {
			fmt.Printf("> Sin respuesta DataNode3.\n")
			flag = false
		} else {
			fmt.Printf("|Cliente| DataNode 3 responde: %s", respuesta_3.Mensaje)
			flag = true
		}

	}
	return flag
}

/*
func Enviar_Propuesta(propuesta serverdatanode.Propuesta, destinatario string) {
	switch destinatario {
	case "DataNode1":
		// Caso DataNode1

	case "DataNode2":
		// Caso DataNode2

	case "DataNode3":
		// Caso DataNode3

	case "NameNode":
		// Caso NameNode (para centralizado)

	}
}
*/

func HacerPropuesta(metodo string, NombreLibroSubido string) {
	var Arreglo_indices_partes_libro []string // Ya no guarda indices, sino que los nombres de los chunks en el directorio
	if metodo == "distribuido" {
		// Enviar mensajes a datanodes para ver si están vivos
		err := enviar_a_DataNode2("DataNode1 pregunta estas vivo?\n")
		flagDN2vivo := true
		if err != true {
			fmt.Printf("DataNode2 no está vivo\n")
			flagDN2vivo = false
		}
		err = enviar_a_DataNode3("DataNode1 pregunta estas vivo?\n")
		flagDN3vivo := true
		if err != true {
			fmt.Printf("DataNode3 no está vivo\n")
			flagDN3vivo = false
		}

		//fmt.Println("flagDN2: %v", flagDN2vivo)
		//fmt.Println("flagDN3: %v", flagDN3vivo)

		// Contar cantidad de partes del libro
		files, err_files := ioutil.ReadDir("./")
		if err_files != nil {
			log.Printf("err_files, no puede leer directorio: %v", err_files)
		}
		for indice_parte, f := range files {
			//fmt.Printf("Nombre scan: %s\n", f.Name())
			//fmt.Printf("NombreLibroSubido: %s\n", NombreLibroSubido)
			if strings.Contains(f.Name(), NombreLibroSubido) {
				Arreglo_indices_partes_libro = append(Arreglo_indices_partes_libro, f.Name())
			}
		}

		fmt.Printf("Partes a repartir:\n")
		fmt.Printf("Arreglo_indices_partes_libro = %v\n", Arreglo_indices_partes_libro)
		// for _, ind := range Arreglo_indices_partes_libro {
		// 	fmt.Printf("ind = %v\n", ind)
		// 	//fmt.Printf("valor = %v\n", valor)
		// 	indint, _ := strconv.Atoi(ind)
		// 	fmt.Printf("%s\n", files[indint].Name())
		// }

		//cant_chunks := len(Arreglo_indices_partes_libro) // Cantidad de chunks

		// Ver cuales hay vivos y repartir con serverdatanode.Propuesta
		Arreglo_copia := Arreglo_indices_partes_libro
		fmt.Printf("%v", Arreglo_copia)
		var PartesDN1 []string
		var PartesDN2 []string
		var PartesDN3 []string
		Propuesta := serverdatanode.Propuesta{
			NombreLibroSubido: NombreLibroSubido,
			PartesDN1:         PartesDN1,
			PartesDN2:         PartesDN2,
			PartesDN3:         PartesDN3,
		}
		Propuesta.PartesDN1 = append(Propuesta.PartesDN1, Arreglo_copia[len(Arreglo_copia)-1])
		i := len(Arreglo_copia) - 1
		// Borrar elemento
		Arreglo_copia[i] = Arreglo_copia[len(Arreglo_copia)-1]
		Arreglo_copia[len(Arreglo_copia)-1] = ""
		Arreglo_copia = Arreglo_copia[:len(Arreglo_copia)-1]

		if flagDN2vivo == true {
			Propuesta.PartesDN2 = append(Propuesta.PartesDN2, Arreglo_copia[len(Arreglo_copia)-1])
			// Borrar elemento
			i = len(Arreglo_copia) - 1
			Arreglo_copia[i] = Arreglo_copia[len(Arreglo_copia)-1]
			Arreglo_copia[len(Arreglo_copia)-1] = ""
			Arreglo_copia = Arreglo_copia[:len(Arreglo_copia)-1]
		}
		if flagDN3vivo == true {
			Propuesta.PartesDN3 = append(Propuesta.PartesDN3, Arreglo_copia[len(Arreglo_copia)-1])
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
		fmt.Printf("Nombre libro: %s\n", NombreLibroSubido)
		fmt.Println("Propuesta.PartesDN1: %v", Propuesta.PartesDN1)
		fmt.Println("Propuesta.PartesDN2: %v", Propuesta.PartesDN2)
		fmt.Println("Propuesta.PartesDN3: %v", Propuesta.PartesDN3)

		//respuesta_propuesta_DN2 := Enviar_Propuesta(&Propuesta, "DataNode2")
		//respuesta_propuesta_DN3 := Enviar_Propuesta(&Propuesta, "DataNode3")

	}
	/*
		else { // Centralizado

		}
	*/

}

func main() {
	fmt.Printf("#### DataNode 1 ####\n\n")

	// Escucha en el puerto 9001
	lis, err_s := net.Listen("tcp", ":9001")
	if err_s != nil {
		log.Fatalf("Error en DataNode 1 al escuchar en puerto 9001: %v", err_s)
	}
	fmt.Println("DataNode 1 escuchando en puerto 9001.\n")
	s := serverdatanode.Server{}

	s.FlagLibroSubido = false

	// Servidor gRPC
	grpcServer := grpc.NewServer()

	serverdatanode.RegisterDataNodeServiceServer(grpcServer, &s)

	go func() {
		for {

			time.Sleep(10 * time.Second)

			mensajeaNN := "Mensaje de prueba DataNode 1 a NameNode\n"
			mensajeaD2 := "Mensaje de prueba DataNode 1 a DataNode 2\n"
			mensajeaD3 := "Mensaje de prueba DataNode 1 a DataNode 3\n"
			enviar_a_NameNode(mensajeaNN)
			enviar_a_DataNode2(mensajeaD2)
			enviar_a_DataNode3(mensajeaD3)
		}
	}()

	go func() {
		for {
			if s.FlagLibroSubido == true {
				// Llamar funcion
				NombreLibroSubido := s.NombreLibroSubido // Va sin ".pdf"
				s.FlagLibroSubido = false
				HacerPropuesta("distribuido", NombreLibroSubido) // POR AHORA DISTRIBUIDO PQ AUN NO SE IMPLMEENTA CENTRALIZADOOOOOOOOOOO
			}
		}
	}()

	if err_s := grpcServer.Serve(lis); err_s != nil {
		log.Fatalf("Error DataNode 1 en servidor gRPC en el puerto 9001: %v", err_s)
	}

}
