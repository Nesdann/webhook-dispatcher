# Interfaces en Go, Java y Scala

## 1. ¿Qué es una interfaz?

Una interfaz es una forma de definir un contrato: "cualquier cosa que quiera ser usada aquí debe tener estos métodos/funciones".

La idea es separar dos cosas:

- la lógica de negocio o la lógica de uso
- la implementación concreta

Esto permite que el código dependa de una abstracción y no de un detalle concreto como PostgreSQL, una API externa, una clase en particular o un archivo físico.

La clave es esta frase:

> La interfaz describe lo que necesita el consumidor, no cómo lo implementa el proveedor.

---

## 2. Interfaces en Go

En Go, una interfaz se define como una lista de métodos.

```go
package main

type EventRepository interface {
    CreateEvent(event Event) error
}

type Event struct {
    ID   string
    Name string
}
```

Ahora, cualquier tipo que tenga el método `CreateEvent(event Event) error` cumple con esa interfaz automáticamente.

```go
type PostgresRepository struct{}

func (p *PostgresRepository) CreateEvent(event Event) error {
    // INSERT INTO events ...
    return nil
}

func SaveEvent(repo EventRepository, e Event) error {
    return repo.CreateEvent(e)
}
```

Esto funciona aunque no haya ninguna palabra como `implements`.

### Característica fundamental de Go: structural typing

En Go no hace falta declarar explícitamente que una estructura implementa una interfaz. Si la estructura tiene todos los métodos necesarios, ya vale.

```go
type FakeRepo struct{}

func (f *FakeRepo) CreateEvent(event Event) error {
    return nil
}
```

`*FakeRepo` puede usarse donde se espera un `EventRepository`, aunque ni siquiera se haya escrito `implements`.

Esto es muy útil para tests, porque podés inyectar un repositorio falso sin depender de una base real.

---

## 3. ¿Por qué esto es tan importante en Go?

Porque Go favorece el diseño por interfaces y por inyección de dependencias.

Un servicio o handler no debería depender directamente de PostgreSQL, Redis, Firebase, etc. Debe depender de una abstracción.

Ejemplo:

```go
type Service struct {
    repo EventRepository
}

func (s *Service) Handle(event Event) error {
    return s.repo.CreateEvent(event)
}
```

Así `Service` no sabe ni le importa qué implementación concreta hay detrás.

Eso significa que podés cambiar de infraestructura sin tocar la lógica de negocio.

---

## 4. Interfaces en Java

En Java, una interfaz es una especie de contrato formal. Una clase debe declararlo explícitamente.

```java
public interface EventRepository {
    void createEvent(Event event) throws Exception;
}
```

Y luego:

```java
public class PostgresRepository implements EventRepository {
    @Override
    public void createEvent(Event event) throws Exception {
        // INSERT a la base
    }
}
```

### Diferencia clave

En Java, la clase debe decir `implements EventRepository`.

Eso hace que el compilador pueda verificar la relación.

La interfaz está diseñada para ser implementada por clases concretas, y la clase debe cumplir la firma exacta.

---

## 5. Interfaces en Scala

En Scala, lo más parecido a una interfaz es un `trait`.

```scala
trait EventRepository {
  def createEvent(event: Event): Either[Error, Unit]
}
```

Y luego una implementación:

```scala
class PostgresRepository extends EventRepository {
  override def createEvent(event: Event): Either[Error, Unit] = {
    // INSERT
    Right(())
  }
}
```

### ¿Qué diferencia tiene Scala?

Scala mezcla ideas de clases, traits y composición.

Un trait puede:

- definir métodos abstractos
- definir métodos concretos
- ser mezclado en clases
- ser reutilizado en varias jerarquías

Ejemplo:

```scala
trait Logger {
  def log(msg: String): Unit = println(msg)
}

class Service extends Logger with EventRepository {
  override def createEvent(event: Event): Either[Error, Unit] = Right(())
}
```

Entonces en Scala, la idea de contrato es muy poderosa y más flexible que en Java.

---

## 6. Comparación entre Go, Java y Scala

### Go

- Las interfaces se usan por estructura.
- No hace falta `implements`.
- El compilador verifica si el tipo tiene los métodos requeridos.
- Es muy simple y muy expresivo.
- Fomenta composición y dependencias abstractas.

### Java

- La interfaz es un contrato formal.
- La clase debe declarar explícitamente `implements`.
- Es más rígido y más orientado a jerarquías.
- Muy común en frameworks grandes, Spring, etc.

### Scala

- Los `trait` reemplazan o complementan a la idea de interfaz.
- Se pueden mezclar en muchas clases.
- Más flexible que Java.
- Tiene una mezcla de OO y FP.

---

## 7. Diferencia central: "qué define la interfaz" en cada lenguaje

### En Go

La interfaz la define el consumidor.

```go
type EventRepository interface {
    CreateEvent(event Event) error
}
```

El servicio dice: *"necesito algo con este método"*.

Luego, cualquier tipo que lo implemente sirve.

### En Java

La interfaz la define el proveedor o el framework.

```java
public interface EventRepository {
    void createEvent(Event event);
}
```

La clase que lo implementa debe adaptarse a ese contrato.

### En Scala

La interfaz/trait puede ser mucho más rica y mezclable.

Un trait puede dar comportamiento concreto además del contrato.

---

## 8. Cómo implementa Go una interfaz muy importante: `http.Handler`

Go tiene una interfaz muy famosa y muy importante:

```go
package net/http

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Esto significa que cualquier tipo que tenga este método puede ser un handler HTTP.

```go
type MiHandler struct{}

func (h *MiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hola desde Go"))
}
```

Y luego:

```go
http.Handle("/", &MiHandler{})
```

### ¿Qué hace esto?

Hace que Go sea extremadamente flexible.

No necesitas heredar de una clase base ni implementar una jerarquía enorme. Solo tenés que definir `ServeHTTP` y ya podés pasar tu tipo a la librería HTTP.

Esto es una demostración muy clara del estilo de Go:

- define la interfaz mínima que necesitas
- cualquier tipo que la cumpla sirve
- no hay herencia obligatoria
- no hay `implements`

---

## 9. Ejemplo práctico: el caso del repositorio

```go
package main

import "fmt"

type Event struct {
    ID   string
    Name string
}

type EventRepository interface {
    CreateEvent(event Event) error
}

type PostgresRepository struct{}

func (p *PostgresRepository) CreateEvent(event Event) error {
    fmt.Println("Guardando en Postgres:", event.Name)
    return nil
}

type MemoryRepository struct{}

func (m *MemoryRepository) CreateEvent(event Event) error {
    fmt.Println("Guardando en memoria:", event.Name)
    return nil
}

type Service struct {
    repo EventRepository
}

func (s *Service) Handle(event Event) error {
    return s.repo.CreateEvent(event)
}

func main() {
    svc := Service{repo: &PostgresRepository{}}
    _ = svc.Handle(Event{ID: "1", Name: "Evento A"})

    svc2 := Service{repo: &MemoryRepository{}}
    _ = svc2.Handle(Event{ID: "2", Name: "Evento B"})
}
```

### Qué pasa aquí

`Service` no sabe si guarda en PostgreSQL o en memoria.

Solo sabe que el repositorio tiene el método `CreateEvent`.

Eso es exactamente la esencia de la interfaz en Go.

---

## 10. Resumen final

### En Go

- La interfaz describe lo mínimo que necesita el código.
- Cualquier tipo con esos métodos cumple la interfaz.
- No hace falta declarar `implements`.
- Es más simple, flexible y muy útil para tests y arquitectura.

### En Java

- La interfaz es un contrato formal.
- La clase debe implementar esa interfaz explícitamente.
- Hay más rigor en el compilador.

### En Scala

- Los `trait` permiten una solución más rica y flexible.
- Se mezclan con clases.
- Tienen más potencia que una interfaz simple de Java.

### La idea central

Go no busca “hacer que una clase diga que cumple una interfaz”.

Go busca: “¿este tipo tiene el comportamiento que necesito?”

Y eso es lo que hace que la interfaz en Go sea tan elegante y tan utilizada.

---

## 11. Conclusión

Si querés entender interfaces en Go, pensalo así:

> En Go, una interfaz no es un “tipo heredado”. Es un conjunto de capacidades que un valor debe tener.

Y eso cambia totalmente cómo se diseña software:

- menos acoplamiento
- más flexibilidad
- más fácil de testear
- más fácil de cambiar la implementación

Eso es por lo que Go es tan bueno para diseño de APIs, servicios y sistemas grandes.

---

## 12. Regla mental rápida

Cuando veas una interfaz en Go, preguntate:

> ¿Qué operación necesita este código?

Y en base a eso defines la interfaz.

Luego cualquier cosa que pueda hacer esa operación sirve.

Y eso es exactamente la diferencia más importante con Java y Scala.
