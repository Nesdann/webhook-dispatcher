# Go: funciones, métodos y validación

Este documento resume lo que hablamos sobre Go, especialmente sobre funciones, métodos y cómo validar un `EventRequest`.

---

## 1) Función en Go

Una función se define así:

```go
func nombre(parametro Tipo) TipoDevuelto {
	return valor
}
```

Ejemplo:

```go
func validateRequest(req EventRequest) error {
	if req.EventType == "" {
		return fmt.Errorf("eventType is required")
	}
	return nil
}
```

### Qué significa:

- `func` = define una función
- `validateRequest` = nombre de la función
- `(req EventRequest)` = recibe un parámetro llamado `req` de tipo `EventRequest`
- `error` = devuelve un error

### Cuando usarla:

Cuando la lógica no pertenece directamente al tipo, sino que es una operación externa.

Ejemplo:

```go
err := validateRequest(req)
```

---

## 2) Método en Go

Un método está asociado a un tipo.

```go
func (e *EventRequest) Validate() error {
	if e.EventType == "" {
		return fmt.Errorf("eventType is required")
	}
	return nil
}
```

### Qué significa:

- `(e *EventRequest)` = este método pertenece al tipo `EventRequest`
- `e` = variable local que apunta al valor del tipo
- `Validate` = nombre del método
- `() error` = devuelve un error

### Cómo se usa:

```go
err := req.Validate()
```

Esto se lee como:

> “Valida este `EventRequest`.”

### Cuando usarlo:

Cuando la validación pertenece directamente al dato mismo.

---

## 3) Diferencia entre función y método

### Función:

```go
func validateRequest(req EventRequest) error
```

Se llama así:

```go
err := validateRequest(req)
```

### Método:

```go
func (e *EventRequest) Validate() error
```

Se llama así:

```go
err := req.Validate()
```

### Regla práctica:

- si la lógica es del objeto → método
- si la lógica es externa → función

En un `struct` tipo `EventRequest`, normalmente es mejor usar un método.

---

## 4) `error` en Go

`error` es el tipo usado para devolver errores.

### Con `errors.New`

```go
return errors.New("eventType is required")
```

Se usa cuando el mensaje es fijo.

### Con `fmt.Errorf`

```go
return fmt.Errorf("Invalid event type: %s", e.EventType)
```

Se usa cuando querés incluir variables dentro del mensaje.

### `return nil`

Cuando todo está bien, devolvés `nil`.

```go
return nil
```

Esto significa:

> “No hubo error”.

---

## 5) Método típico para validación

```go
func (e *EventRequest) ValidateMethodPost() error {
	if e.EventType != "POST" {
		return fmt.Errorf("Invalid event type: %s", e.EventType)
	}
	return nil
}
```

### Qué hace:

- revisa si `EventType` es exactamente `"POST"`
- si no lo es, devuelve un error
- si sí lo es, devuelve `nil`

### Importante:

Si una función o método tiene una cláusula `return` en ciertos casos, siempre debe devolver algo en todos los caminos posibles.

Esto está bien:

```go
func (e *EventRequest) ValidateMethodPost() error {
	if e.EventType != "POST" {
		return fmt.Errorf("Invalid event type: %s", e.EventType)
	}
	return nil
}
```

---

## 6) `EventRequest` de ejemplo

```go
type EventRequest struct {
	EventType string `json:"eventType"`
	Payload   json.RawMessage `json:"payload"`
}
```

### Qué significa:

- `EventType` = el tipo de evento
- `Payload` = la información que llega en JSON

---

## 7) Validación y HTTP

En el handler normalmente haces esto:

```go
if r.Method != http.MethodPost {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}
```

Y después validás la estructura:

```go
var req EventRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	return
}

if err := req.ValidateMethodPost(); err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
```

### Flujo:

1. verificás el método HTTP
2. lees el JSON
3. validás los campos
4. si hay error, respondés con 400
5. si todo está bien, seguís

---

## 8) Resumen corto

### `func nombre(...)`

- función normal
- se llama por nombre

### `func (e *Tipo) Nombre()`

- método del tipo
- se llama sobre una variable del tipo

### `return fmt.Errorf(...)`

- devuelve un error con mensaje dinámico

### `return errors.New(...)`

- devuelve un error con mensaje fijo

### `return nil`

- significa que no hubo error

---

## 9) Recomendación para tu proyecto

Para tu `EventRequest`, esta forma es muy limpia:

```go
func (e *EventRequest) Validate() error {
	if strings.TrimSpace(e.EventType) == "" {
		return errors.New("eventType is required")
	}

	if len(e.Payload) == 0 {
		return errors.New("payload is required")
	}

	return nil
}
```

Esto deja el código más ordenado, más claro y más fácil de mantener.

---

## 10) Importante:

Una función o método que devuelve `error` debe retornar un `error` en todos los caminos posibles.

Esto es correcto:

```go
func Demo() error {
	if true {
		return errors.New("error")
	}
	return nil
}
```

Esto no es correcto:

```go
func Demo() error {
	if true {
		return errors.New("error")
	}
}
```

Porque si la condición no se cumple, no hay `return` al final.

---

## 11) Regla mental fácil

Lee esto así:

```go
func (e *EventRequest) Validate() error
```

> “El tipo `EventRequest` tiene un método llamado `Validate` que devuelve un error.”

Y esto:

```go
func validateRequest(req EventRequest) error
```

> “Hay una función externa que recibe una `EventRequest` y devuelve un error.”

---

## 12) Ejemplo final de uso

```go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EventRequest struct {
	EventType string `json:"eventType"`
	Payload   json.RawMessage `json:"payload"`
}

func (e *EventRequest) ValidateMethodPost() error {
	if e.EventType != "POST" {
		return fmt.Errorf("Invalid event type: %s", e.EventType)
	}
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := req.ValidateMethodPost(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
```

---

# Conclusión

Go te permite hacer dos cosas muy útiles:

- funciones para lógica externa
- métodos para lógica propia del tipo

Cuando validás un struct como `EventRequest`, los métodos son una opción muy clara y elegante.

Si querés, después podemos armar una segunda nota con:

- `json.Unmarshal`
- `json.Decoder`
- validación de payload
- ejemplos de handlers HTTP completos
- cómo probarlos con tests
