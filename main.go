//package main es muy importante
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	
)

//ioutil is for read a file for example

// coemntario

/*
	multilinea xD
*/
func main() {
	const pi = 3.1416
	var dog string
	//para guardar punteros
	var dog2 *string
	dog2 = &dog
	dog = "perrito"
	cat := 1
	fmt.Println(dog, cat, pi)
	//imprimo el tipo y el valor de la variable dog en esos parametros luego del string y & para sacar memoria
	fmt.Printf("Tipo %T , Valor %v dirección de memoria %v\n",dog,dog,&dog)

	fmt.Printf("tipo %T valor %v desreferenciacion %s\n",dog2,dog2, *dog2)
	
	// arreglos

	var food [3] uint8
	food [0] = 1
	food [1] = 2
	food [2] = 3
	fmt.Printf("tipo %T valor %v\n ",food[2], food[2])

	//slices son apuntadores a un array pero no poseen datos 

	set := [7]string{"a","bb","ccc","ddd","eeee","fffff","gggggg"}
	// este es el slice indico de qué posición eempizo apuntar y hasta cuál acá es de 0,3 porque no toma el ultimo a lo ciclo python, a su vez puedo 
	//indicarle indice final o inicial opcional
	letras := set[:5]
	letras1 := set[3:]
	//como es apuntador cambia el dato en el arreglo general desde el slice y en todos los otros slice posible }s del mismo array 
	letras1[0] = "aaNuevo"
	//al hacer append cambia el dato de la siguiente posición final del slice por ejemplo 
	letras = append(letras,"perra")
	//cuando el append llega al limite del array y lo sobrepasa, crea un nuevo arreglo copia el slice y agrega los elementos nuevos ejemplo y 
	//pierde la referencia al array original
	letras = append(letras,"perra1")
	letras = append(letras,"perra2")
	letras = append(letras,"perra3")
	letras = append(letras,"perra4")
	fmt.Println("array:",set)
	fmt.Println("letras1:",letras1)
	fmt.Println("letras:",letras)
	//len del arreglo total o sea el set
	fmt.Println("len: ",len(letras))
	//capacidad es el numero de elementos pero a partir donde se apunta o sea si el arreglo es de 10 posiciones y mi slice empieza en el 4 tengo capacidad de 5
	fmt.Println("len: ",cap(letras))
	
	//maps are key value structure
	animals := make(map[string]int)
	animals["perro"] = 0
	animals["cat"] = 1
	fmt.Println(animals)
	fruits:=map[string]string{
		"apple" : "manzana",
		"banana" : "banano",
		"pera" : "no se",
	}
	fmt.Println(fruits)
	//borrar de un map
	delete(fruits,"apple")
	fmt.Println(fruits)

	//para crear "clases" es con struct
	type course struct{
		name string
		teacher string
		country string
	}
	db:= course{
		name:"Go",
		teacher: "Armando",
		country: "Colombia",
	}
	// o de esta forma

	db1 := course{"Caparzo", "Pepito", "Colombia"}

	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", db1)
	fmt.Printf("%+v\n", db.name)
	fmt.Printf("%+v\n", db.teacher)
	fmt.Printf("%+v\n", db.country)
	//condicional if
	if false {
		fmt.Println("priemra parte")
	}else if false {
		fmt.Println("segunda parte")
	}else{
		fmt.Println("else final")
	}
	//switch no se usa break
	perro := "perro"
	switch perro {
	case "Perro":
		fmt.Println("Es un perro")
	case "basura":
		fmt.Println("Es basura")
	default:
		fmt.Println("no es nada mka")
	}
	sise := 10
	//for
	for i := 0; i < sise; i++ {
		fmt.Println(i)
	}
	//for forever para sockets o cosas así, imprime infinitametne
	/*i:=1
	for{
		fmt.Println(i)
		i++
	}*/
	//for range para slices, se ejecuta el numero de elemntos del arreglo osea 5 lo mismo para key value structures
	nums := []uint8{2,4,5,6,8}
	for i,v := range nums{
		fmt.Println("indice ",i, " Valor ",v)
		fmt.Println("indice ",i, " Valor ", nums[i]*2)
	}
	//esta vuelta cuando recorro strings me rotarna es el valor en bytes de la letra
	hello:="hello"
	for _,v := range hello{
		fmt.Println(v)
		
	}
	for _,v := range hello{
		fmt.Println(string(v))
		
	}
	hellos("Armando", "Rios")
	hola := "perra"
	fmt.Println("\n",hola)
	change(&hola)
	fmt.Println("\n",hola)
	texto := "BiTcH PlEsEa"
	minuscula , mayuscula := convert(texto)
	fmt.Println(minuscula,mayuscula)


	//MANEJO DE ERORRES MUY IMPORTANTE
	content, err := ioutil.ReadFile("./things.txt")
	if err != nil {
		fmt.Println("Ocurrió un error %v", err)
		return
	}
	fmt.Println(string(content))

	result , err := division(10, 2)
	if err != nil{
		fmt.Println("Ocurrió un error %v", err)
		return
	}
	fmt.Println("el resultado es ",result)
	
	
	//FUNCIONES ANIDADAS O SEA RECIBVEN Y RETORNAN FUNCIONES
	nums1 := []int{1,2,34,5,6,63,12}
	result1 := filter(nums1, func (num int) bool{
		if num <= 10{
			return true
		}
		return false
	})
	fmt.Println("the result is ", result1)
	//SI LE QUITO LOS PARENTESIS FINALES NO EJECUTA LA SEGUNDA FUNCIÓN Y SOLAMENTE RECIBO EL ESPACIO EN MEMORIA DODNE ESTÁ GUARDADA LA FUNCIÓN, PERO con ellos
	// puedo ejecutar al segunda y obtrener el nombre bien
	x := HelloRaro("Armando")()
	fmt.Println(x)

	defer fmt.Println("Holaaaa")
	
	//CREAR ARCHIVOS
	file , err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo %v",err)
		return
	}
	//SE EJECUTARÁ SIEMPRE AL ACABAR LA EJECUCIÓN ASÍ QUE NO ES NECESARIO EL CLOSE DLEL IF
	defer file.Close()
	_,err = file.Write([]byte("Hello bitch"))
	if err != nil {
		//file.Close()
		fmt.Printf("Ocurrió un error al escribir el archivo %v",err)
		return
	}
	//IMPORTAR MIS PAQUETES --------------------------------------------------------------------------------------------------------------------------------------------
	saludarpkg.English()



}








//por valor esto es sacar copia no es directamente en memoria, para memoría toca usar punteros
func hellos(firstName string, lastName string){
	fmt.Printf("Hello %s %s ",firstName, lastName)
}
//por referencia y le pasamos cuando la invocamos en el main con la dirección y acá con la referencia a la memoria
func change(value *string){
	*value = "hola"
}
//return func
func sum(num1,num2 int) int  {
	return num1+num2
}
//return multiples values
func convert(text string) (string, string){
	min := strings.ToLower(text)
	may := strings.ToUpper(text)
	return min,may
}
//MANEJO DE ERORRES MUY IMPORTANTE
func division(dividendo, divisor int) (int, error){
	if divisor == 0 {
		return 0,errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}
//SEGUNDA VERSION DE MANEJO DE ERRORES SINTAXIS CORTA SOLAMENTE RECOMENDADO PARA FUNCIONES CORTAS

func division1(dividendo, divisor int) (result int,err error){
	if divisor == 0 {
		//Acá como nombro los parametros de retorno entonces no requiero especificar el retorno solametne iniciarlizar el err Y dejar RETURN vacio
		err = errors.New("No se puede dividir por cero")
		return 
	}
	result = dividendo / divisor
	return 
}
//FUNCIONES ANIDADAS O SEA RECIBVEN Y RETORNAN FUNCIONES

func filter(nums []int, callback func(int) bool) []int{
	result := []int{}
	for _,v := range nums{
		if callback(v){
			result = append(result,v)
		}
	}
	return result
}
//RETORNA UNA FUNCIÓN
func HelloRaro(name string) func() string{
	return func() string{
		return "Hello " + name
	}
}
//FUNCIONES VARIATICAS RECIBIR VARIOS APRAMETROS DE UN MISMO TIPO CAMBIANDO COSNTANTE MENTE POR EJEMPLO SUMA num1,num2 pero luego puedo agregar num3,num4 sin modificar
// la declaración de la función
func sumVariatica(nums ...int) int{
	total := 0
	for _,v := range nums{
		total += v
	}
	return total
}

//DEFER SIRVE PARA APLAZAR FUNCIOENS PUES LAS EJECUCIONES AL FINAL CUANDO SE VA ACABAR YA TODO EL RPOGRAMA PERO LOS VALORES SE GUARDAN EN EL MISMO ISNTANTE O SEA SI IMPRIMO
// 5 CON DEFER y luego lo modifico la modificación no aparece


// Hay otra vuelta relaconada con DEFER QUE es RECOVER Y PANIC es como una excepción controlada. el recover es para recuperarse de esa excepción