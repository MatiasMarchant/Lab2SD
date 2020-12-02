package servernamenode

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"serverdatanode"
)

type Propuesta struct {
	NombreLibroSubido string
	PartesDN1         []string
	PartesDN2         []string
	PartesDN3         []string // Strings para hacer la concatenacion mas facil
}

type Server struct {
}

//-------------------------------------------------------------------------------------------------------------
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

//-----------------------------------------------------------------------------------------------------------------

func listaDeLibros() string {
	listado := ""
	nLibro := 0
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		p_linea := strings.Split(linea, " ")[0]
		if _, err := strconv.Atoi(p_linea); err != nil {
			nLibro += 1
			n := strconv.Itoa(nLibro)
			if p_linea == ""{
				fmt.Printf("Vacio\n")
				fmt.Println(p_linea)
			} //else {
			//	listado += n + " " + p_linea + "\n"
			//}
			listado += n + " " + p_linea + "\n"
			
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return listado
}

func ubicacionLibro(tituloLibro string) string {
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	libroEncontrado := false
	n_chunk := 0
	max_chunks := 0
	chunks_totales := ""

	for scanner.Scan() {
		linea := scanner.Text()
		splt_linea := strings.Split(linea, " ")
		i_linea := splt_linea[0]
		ii_linea := splt_linea[1]

		if !libroEncontrado && i_linea == tituloLibro {
			libroEncontrado = true
			max_chunks, _ = strconv.Atoi(ii_linea)
		} else if libroEncontrado {
			if n_chunk < max_chunks {
				chunks_totales += linea + "\n"
				n_chunk += 1
			} else {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return chunks_totales
}

func (s *Server) EnvioMensajeTest(ctx context.Context, message *MensajeTest) (*MensajeTest, error) {

	if message.Mensaje == "listadoLibros" {
		fmt.Printf("> Enviando listado de libros\n")
		respuestaDataNode := listaDeLibros()
		return &MensajeTest{Mensaje: respuestaDataNode}, nil
	} else if strings.Contains(message.Mensaje, "ubicacion") {
		tituloLibro := strings.Split(message.Mensaje, " ")[1]
		fmt.Printf("> Enviando ubicación de libro '" + tituloLibro + "'\n")
		respuestaDataNode := ubicacionLibro(tituloLibro)
		return &MensajeTest{Mensaje: respuestaDataNode}, nil
	}

	respuestaNameNode := "Name node recibe: " + message.Mensaje
	return &MensajeTest{Mensaje: respuestaNameNode}, nil
}

// https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *Server) EscribirEnLog(ctx context.Context, message *EscrituraLog) (*MensajeTest, error) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return &MensajeTest{Mensaje: "mal"}, nil
	}
	defer f.Close()

	// Podemos asumir que el
	// mismo libro no se sube mas de una vez

	// Asumiendo q el mismo libro no se sube mas de una vez:
	// Se escribe Nombre Libro y cant partes
	f.WriteString(message.NombreLibro + " " + strconv.Itoa(int(message.CantPartes)) + "\n")
	PartesDN1 := strings.Split(message.PartesDN1, ",")
	PartesDN2 := strings.Split(message.PartesDN2, ",")
	PartesDN3 := strings.Split(message.PartesDN3, ",")
	// Sacarle el nombre a las partes, para solo escribir el indice
	// Y sortear de menor a mayor

	for indice, palabra := range PartesDN1 {
		palabra = strings.TrimPrefix(palabra, message.NombreLibro+"_")
		PartesDN1[indice] = palabra
	}
	sort.Strings(PartesDN1)
	for indice, palabra := range PartesDN2 {
		palabra = strings.TrimPrefix(palabra, message.NombreLibro+"_")
		PartesDN2[indice] = palabra
	}
	sort.Strings(PartesDN2)
	for indice, palabra := range PartesDN3 {
		palabra = strings.TrimPrefix(palabra, message.NombreLibro+"_")
		PartesDN3[indice] = palabra
	}
	sort.Strings(PartesDN3)

	// Ahora si se puede escribir
	intmessageCantPartes := int(message.CantPartes)
	for i := 0; i < intmessageCantPartes; i++ {
		iastring := strconv.Itoa(i)
		if stringInSlice(iastring, PartesDN1) {
			f.WriteString(iastring + " dist37\n")
		}
		if stringInSlice(iastring, PartesDN2) {
			f.WriteString(iastring + " dist38\n")
		}
		if stringInSlice(iastring, PartesDN3) {
			f.WriteString(iastring + " dist39\n")
		}
	}

	return &MensajeTest{Mensaje: "bien"}, nil
}

func generar_nueva_propuesta(Propuesta *Propuestagrpc, flagDN1vivo bool, flagDN2vivo bool, flagDN3vivo bool) (*Propuestagrpc, error){
	propuesta_DN1 := strings.Split(Propuesta.PartesDN1, ",")
	propuesta_DN2 := strings.Split(Propuesta.PartesDN2, ",")
	propuesta_DN3 := strings.Split(Propuesta.PartesDN3, ",")

	// DN1 muerto. DN2 y DN3 vivos
	if !flagDN1vivo && flagDN2vivo && flagDN3vivo{
		Propuesta.PartesDN1 = " "
		for i, prop := range propuesta_DN1 { 
			if i%2 == 0{
				propuesta_DN2 = append(propuesta_DN2, prop)
			} else{
				propuesta_DN3 = append(propuesta_DN3, prop)
			}
		}
	}
	// DN2 muerto. DN1 y DN3 vivos
	if !flagDN2vivo && flagDN1vivo && flagDN3vivo{
		Propuesta.PartesDN2 = " "
		for i, prop := range propuesta_DN2 { 
			if i%2 == 0{
				propuesta_DN1 = append(propuesta_DN1, prop)
			} else{
				propuesta_DN3 = append(propuesta_DN3, prop)
			}
		}
	}
	// DN3 muerto. DN1 y DN2 vivos
	if !flagDN3vivo && flagDN1vivo && flagDN2vivo{
		Propuesta.PartesDN3 = " "
		for i, prop := range propuesta_DN3 { 
			if i%2 == 0{
				propuesta_DN1 = append(propuesta_DN1, prop)
			} else{
				propuesta_DN3 = append(propuesta_DN3, prop)
			}
		}
	}
	// DN3 vivo. DN1 y DN2 muertos
	if !flagDN1vivo && !flagDN2vivo && flagDN3vivo{
		Propuesta.PartesDN1 = " "
		Propuesta.PartesDN2 = " "
		propuesta_DN3 = append(propuesta_DN3, propuesta_DN1...)
		propuesta_DN3 = append(propuesta_DN3, propuesta_DN2...)
		Propuesta.PartesDN3 = strings.Join(propuesta_DN3, ",")		
	}
	// DN2 vivo. DN1 y DN3 muertos
	if !flagDN1vivo && !flagDN3vivo && flagDN2vivo{
		Propuesta.PartesDN1 = " "
		Propuesta.PartesDN3 = " "
		propuesta_DN2 = append(propuesta_DN2, propuesta_DN1...)
		propuesta_DN2 = append(propuesta_DN2, propuesta_DN3...)
		Propuesta.PartesDN2 = strings.Join(propuesta_DN2, ",")		
	}
	// DN1 vivo. DN2 y DN3 muertos
	if !flagDN2vivo && !flagDN3vivo && flagDN1vivo{
		Propuesta.PartesDN2 = " "
		Propuesta.PartesDN3 = " "
		propuesta_DN1 = append(propuesta_DN1, propuesta_DN2...)
		propuesta_DN1 = append(propuesta_DN1, propuesta_DN3...)
		Propuesta.PartesDN1 = strings.Join(propuesta_DN1, ",")		
	}

	return Propuesta, nil	
}

func (s *Server) Propuesta_Centralizado(ctx context.Context, Propuesta *Propuestagrpc) (*Propuestagrpc, error) {
	fmt.Printf("> Propuesta recibida\n")

	err1 := enviar_a_DataNode1("NameNode pregunta estas vivo?\n")
	flagDN1vivo := true
	if err1 != true {
		flagDN1vivo = false
	}

	err2 := enviar_a_DataNode2("NameNode pregunta estas vivo?\n")
	flagDN2vivo := true
	if err2 != true {
		flagDN2vivo = false
	}

	err3 := enviar_a_DataNode3("NameNode pregunta estas vivo?\n")
	flagDN3vivo := true
	if err3 != true {
		flagDN3vivo = false
	}

	if flagDN1vivo && flagDN2vivo && flagDN3vivo{
		fmt.Printf("> Propuesta aceptada\n")
		return Propuesta, nil
	} else{
		fmt.Printf("> Propuesta rechazada\n")
		fmt.Printf("DataNode1: %t, DataNode2: %t, DataNode3: %t\n", flagDN1vivo,flagDN2vivo,flagDN3vivo)

		nuevaPropuesta, _ := generar_nueva_propuesta(Propuesta, flagDN1vivo, flagDN2vivo, flagDN3vivo)

		fmt.Printf("> Nueva propuesta generada\n")
		return nuevaPropuesta, nil
	}




	
}