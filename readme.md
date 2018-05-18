# learn-go-with-tests

<https://quii.gitbook.io/learn-go-with-tests/>

## hello

Debemos intentar usar constantes siempre que sea posible para mejorar el rendimiento de la aplicación y el mantenimiento.

Las funciones pueden ser públicas y privadas, se diferencian si empiezan con mayúsculas, públicas, o minúsculas, privadas.

Podemos utilizar un named return. La variable existirá dentro de la función y si llamamos a return nos devolverá directamente la variable con el valor que contenga.

```go
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

//En Go, las funciones privadas empiezan en minúsculas, y las publicas en mayúsculas.
func greetingPrefix (language string) (prefix string){
    switch language {
    case spanish:
        prefix=spanishHelloPrefix
    case french:
        prefix=frenchHelloPrefix
    default:
        prefix=helloPrefix
    }
    return
}
```

Generamos test en el archivo `nombre_archivo_test.go`.
Debemos importar el paquete `testing` y podremos crear test para las funciones con `TestNombrefuncion`

Como argumento hay que pasar `t *testing.T`

```go
func TestHello(t *testing.T) {
	got := Hello("YOU", "")
	want := "Hello, YOU"
    if got != want {
    	t.Errorf("got '%s' want '%s'", got, want)
    }
}
```

Ejecutamos los tests con `go test`

Podemos ejecutar subtests con `t.run` y usar funciones dentro de los test para refactorizar

```go
//Dentro de TestHello
esCorrecto := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got '%s' want '%s'", got, want)
    }
}
t.Run("saying hello to one person", func(t *testing.T) {
    got := Hello("YOU", "")
    want := "Hello, YOU"
    esCorrecto(t, got, want)
})
```

`t.Helper()` permite que si hay un error, en vez de indicar la línea de la función `esCorrecto` indique la línea que ha llamado a la función `esCorrecto`.

## integers

Podemos incluir un comentario antes de la función para para que aparezca en el Godoc

```go
// Add takes two integers and returns the sum of them
func Add(x, y int) int {
    return x + y
}
```

Podemos ejecutar `go test -v` para obtener mayor feedback de `go test`

En los test podemos crear funciones con `ExampleNombrefuncion` para incluir ejemplos en godoc
Si le indicamos el comentario `// Output: X` el ejemplo se ejecutará.

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```

## iteration

En Go solo existe `for` para realizar bucles. este `for` es flexible y permite realizar bluces clásicos `for(i=0;i<2;i++)`, `while`, `continue` y `break`.

```go
//The most basic type, with a single condition.
i := 1
for i <= 3 {
	fmt.Println(i)
	i = i + 1
}
//A classic initial/condition/after for loop.
for j := 7; j <= 9; j++ {
	fmt.Println(j)
}

//for without a condition will loop repeatedly until you break out of the loop
// or return from the enclosing function.
for {
	fmt.Println("loop")
	break
}
//You can also continue to the next iteration of the loop.
for n := 0; n <= 5; n++ {
	if n%2 == 0 {
		continue
	}
	fmt.Println(n)
}
```

de: <https://gobyexample.com/for>

Podemos realizar benchmarks para obtener el rendimiento de una función en Go. En archivo `_test.go` si empezamos una función con el prefijo `Benchmark` podemos luego ejecutar el comando `go test -bench=.` para obtener resultados:

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i <; b.N; i++ {
		Repeat("a", 10)
	}
}
```

## array_slices

Los arrays tienen una longitud fija y las slices tienen una logitud dinamica. podemos inicializarlos, y obtener el valor de una posición en concreto, añadir un valor, obtener la logitud...

```go
var arr1 [5]int
arr2:= [5]int{1,2,3,4,5}

var sli1 []int
sli2:= []int{1,2,3}

arr2[0] // 1
sli2[2] // 3

sli2=append(sli2,4)
sli2[3] // 4

len(arr2) // 5
```

Podemos declarar un slice de slices

```go
sli3:= [][]string
```

En los bucles, podemos recorrerlos usando `range`. `range` devuelve dos valores. Si no vamos a usar i, podemos pasar el operador blank.

```go
for i, num := range numbers {
	i 	//El indice de la posición actual
	num	//El valor de la posición actual
}

for _, num := range numbers {
	summa += num
}
```

A las funciones podemos pasarle un número variable de argumentos usando el operador `...`.

```go
func sum(nums ...int) (total int){
    for _, num := range nums {
        total += num
    }
    return
}
```

de: <https://gobyexample.com/variadic-functions>
En el caso superior nums se trata como un array de `int`, como si fuera `[]int`

no podemos comparar directamente el contenido de slices. Solo se puede comparar si fuera `nil`. Si utilizamos la libreria `reflect`, tenemos a nuestra disposición `reflect.DeepEqual`

```go
got := []int{1, 2, 3})
expected := []int{6, 10}
if !reflect.DeepEqual(got, expected) {
	t.Errorf("expected %v got %v", expected, got)
}
```

Podemos obtener un slices de otro slice con nombreSlice[low:high]

```go
chars:=[]string{"a","b","c"}
charsSliced := chars[1:] // "b","c"
```

## Structs, methods & interfaces

Podemos crear estructuras que tengan varias variables asociadas. Podemos asignar métodos a estas variables declarando una función e indicando a que estructura se asocia.
```go
type Rectangle struct {
	Width  float64
	Height float64
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

rectangle := Rectangle{10.0, 10.0}
rectangle := Rectangle{Width: 12, Height: 6}

perimetro := rectangle.Perimeter()
```

Las interfaces nos permiten crear funciones asociadas a distintos tipos.

```go
type Shape interface {
    Perimeter() float64
}
```
¿Como se transforma `Rectangle` en `Shape`?
Creamos un nuevo tipo `Shape` pero declarándolo como `interface` en vez de `struct`. Normalmente tendrias que escribir que `Rectangle` implementa `Shape` pero en Go no.
En nuestro caso, `Rectangle` tiene un método con el nombre `Permiter` que devuelve un `float64`, por lo tanto es un `Shape` y ya quedan asociados.
