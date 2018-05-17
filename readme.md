# learn-go-with-tests

https://quii.gitbook.io/learn-go-with-tests/

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
