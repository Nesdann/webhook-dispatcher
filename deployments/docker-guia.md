# Guía de Docker para principiantes

## ¿Qué es Docker?

Docker es una herramienta que permite empaquetar una aplicación y todas sus dependencias en una unidad llamada contenedor. Esto hace que la aplicación se ejecute igual en cualquier computadora, ya sea en tu laptop, un servidor o en la nube.

La idea principal es evitar el problema clásico de:

- “En mi equipo funciona”
- “En el servidor falla”

Docker ayuda a que todo se ejecute de manera consistente.

---

## ¿Para qué sirve Docker?

Docker sirve para:

- Ejecutar aplicaciones en contenedores
- Crear entornos de desarrollo reproducibles
- Evitar conflictos entre versiones de software
- Ejecutar bases de datos, APIs, colas de mensajes y otros servicios
- Desplegar aplicaciones en servidores y cloud
- Aislar servicios sin afectar el sistema operativo principal

Ejemplos comunes:

- Base de datos PostgreSQL
- MongoDB
- MySQL
- Redis
- RabbitMQ
- Aplicaciones backend en Node.js, Python, Java, .NET, etc.

---

## ¿Qué es un contenedor?

Un contenedor es una instancia ligera y aislada de una aplicación o servicio.

Piensa en un contenedor como una máquina virtual muy ligera, pero no es una VM completa. En lugar de virtualizar un sistema operativo entero, el contenedor comparte el kernel del sistema operativo host.

### Diferencia entre contenedor y máquina virtual

- Máquina virtual:
  - Tiene su propio sistema operativo completo
  - Consume más recursos
  - Es más pesada

- Contenedor:
  - Usa el mismo kernel del sistema operativo host
  - Es más rápido de iniciar
  - Consume menos memoria y CPU
  - Es muy adecuado para microservicios y aplicaciones modernas

### Analogía simple

Si una aplicación es un programa, entonces:

- La imagen es el molde o plantilla
- El contenedor es la copia en ejecución de ese molde

---

## ¿Qué es una imagen?

Una imagen es un paquete que contiene:

- El sistema operativo base
- Las dependencias necesarias
- La aplicación
- Configuraciones

Por ejemplo:

```yaml
image: postgres:15-alpine
```

Eso significa que Docker descargará la imagen oficial de PostgreSQL 15 basada en Alpine Linux y creará un contenedor a partir de esa imagen.

Las imágenes se guardan en registries, como:

- Docker Hub
- GitHub Container Registry
- Azure Container Registry
- AWS ECR

---

## ¿Qué es Docker Compose?

Docker Compose es una herramienta para definir y ejecutar aplicaciones con múltiples contenedores usando un solo archivo de configuración.

En lugar de correr varios comandos manuales para levantar servicios, puedes describir todo en un archivo `docker-compose.yml` y luego simplemente ejecutar:

```bash
docker compose up -d
```

Esto levanta todos los servicios declarados en ese archivo.

---

## ¿Qué es un archivo .yml o docker-compose.yml?

Un archivo YAML (`.yml` o `.yaml`) es un archivo de configuración con estructura legible para humanos.

Docker Compose suele usar un archivo llamado:

```bash
docker-compose.yml
```

Este archivo describe:

- Qué servicios se van a levantar
- Qué imagen usa cada servicio
- Qué puertos exponer
- Qué variables de entorno cargar
- Qué volúmenes conservar datos
- Cómo se conectan entre sí

### Ejemplo básico

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: admin
      POSTGRES_DB: webhook_dispatcher
    ports:
      - "5432:5432"
```

Este archivo indica que se va a crear un contenedor de PostgreSQL con:

- usuario `admin`
- contraseña `admin`
- base de datos `webhook_dispatcher`
- puerto `5432` expuesto al host

---

## ¿Cómo se usa Docker?

### 1. Verificar que Docker está instalado

```bash
docker --version
```

### 2. Verificar Docker Compose

```bash
docker compose version
```

### 3. Levantar servicios desde un archivo Compose

```bash
docker compose -f docker-compose.yml up -d
```

- `-f` = archivo que se va a usar
- `up` = crear o iniciar contenedores
- `-d` = modo detached (en segundo plano)

### 4. Ver los contenedores activos

```bash
docker ps
```

### 5. Ver logs

```bash
docker logs <nombre_o_id_del_contenedor>
```

### 6. Detener servicios

```bash
docker compose down
```

### 7. Reiniciar un servicio

```bash
docker compose restart postgres
```

### 8. Eliminar contenedores

```bash
docker rm <contenedor>
```

### 9. Eliminar imágenes

```bash
docker rmi <imagen>
```

---

## ¿Qué es un volumen?

Un volumen es un espacio de almacenamiento persistente que se guarda fuera del contenedor.

Esto es importante porque los contenedores normalmente no guardan datos de forma permanente si se eliminan.

### Ejemplo

```yaml
volumes:
  pgdata:
```

Y luego se usa así:

```yaml
volumes:
  - pgdata:/var/lib/postgresql/data
```

Esto significa:

- `pgdata` = volumen nombrado
- `/var/lib/postgresql/data` = ruta donde PostgreSQL guarda sus datos

### ¿Por qué se usan volúmenes?

Porque permiten:

- Guardar bases de datos
- Mantener archivos de configuración
- Evitar perder información al reiniciar el contenedor
- Separar datos de la aplicación

---

## ¿Qué es `systemctl`?

`systemctl` es la herramienta que usa Linux para gestionar servicios del sistema operativo.

Por ejemplo:

```bash
sudo systemctl start postgresql
sudo systemctl status postgresql
sudo systemctl stop postgresql
```

Se usa para servicios que ya están instalados como programas del sistema, como:

- PostgreSQL
- MySQL
- Nginx
- Redis
- RabbitMQ

### ¿Qué hace exactamente?

`systemctl` inicia, detiene, reinicia o habilita servicios del sistema operativo. La idea es manejar servicios de forma nativa en la máquina donde están instalados.

No es un sistema para “empaquetar” la aplicación junto con sus dependencias. Solo administra el servicio que ya existe dentro del sistema.

---

## ¿Por qué Docker tiene ventaja sobre `systemctl`?

Docker tiene una gran ventaja: hace que el entorno sea reproducible y portable.

### Con `systemctl`

Tú instalas PostgreSQL o RabbitMQ directamente en la máquina y configuras todo manualmente.

Eso significa:

- cada máquina requiere su propia configuración
- si cambias de PC o servidor, tienes que repetir pasos
- puede haber diferencias entre entornos
- se complica el despliegue y mantenimiento

### Con Docker

Tú defines todo en un archivo `docker-compose.yml` y lo levantas con un comando:

```bash
docker compose up -d
```

Eso hace que el entorno se cree igual en cualquier máquina que tenga Docker instalado.

### Ejemplo de portabilidad

Si en una máquina tienes:

```yaml
services:
  postgres:
    image: postgres:15-alpine
```

Y en otra máquina haces exactamente lo mismo, ambos levantarán el mismo servicio con la misma imagen.

Con `systemctl`, eso no ocurre automáticamente, porque cada servidor puede tener diferente versión, configuración, directorios, permisos y dependencias.

---

## ¿Por qué no tiene la misma ventaja?

Porque `systemctl` no encapsula la aplicación ni sus dependencias. Solo administra procesos del sistema operativo.

Es decir:

- Docker = “empaqueto la aplicación y la ejecuto igual en cualquier lugar”
- `systemctl` = “administro un servicio que ya está instalado en esta máquina”

### Ejemplo sencillo

#### Opción 1: con Docker

```bash
docker compose up -d
```

Siempre que el repositorio tenga el archivo `docker-compose.yml`, otra persona puede levantar exactamente el mismo entorno.

#### Opción 2: con `systemctl`

```bash
sudo apt install postgresql
sudo systemctl start postgresql
```

Eso solo funciona si la otra máquina tiene PostgreSQL instalado, configurado y preparado igual. No es automático ni portable.

---

## ¿Cuándo conviene cada uno?

### Usa Docker cuando:

- quieres un entorno consistente
- trabajas con varios servicios a la vez
- necesitas reproducir el proyecto en otra PC o servidor
- quieres aislar dependencias y versiones
- tu proyecto tiene base de datos, cola, backend y frontend juntos

### Usa `systemctl` cuando:

- solo quieres administrar un servicio ya instalado en una máquina
- no necesitas portabilidad
- el entorno es simple y está configurado en esa máquina solamente

---

## ¿Qué son los puertos en Docker?

Los puertos permiten exponer servicios del contenedor hacia el host o hacia otras máquinas.

Ejemplo:

```yaml
ports:
  - "5432:5432"
```

Esto dice:

- `5432` del host
- `5432` del contenedor

Es decir, tú puedes acceder a PostgreSQL desde tu máquina local usando el puerto 5432.

### ¿Por qué Docker “escucha” en ese puerto si el contenedor también tiene su propio puerto?

Porque Docker actúa como intermedio.

El contenedor no tiene una IP real accesible directamente como si fuera un proceso normal del sistema. Docker crea una red virtual y el contenedor se comunica dentro de esa red. Cuando tú publicas un puerto con:

```yaml
ports:
  - "5432:5432"
```

lo que hace Docker es:

- abrir el puerto `5432` del host
- redirigir el tráfico que llega ahí hacia el puerto `5432` del contenedor

Es decir, la máquina local no habla directamente con PostgreSQL “desde afuera” sin pasar por Docker. Docker recibe el tráfico, lo reenvía al contenedor y devuelve la respuesta.

### Analogía simple

Es como si Docker fuera un portero o un conserje:

- tú vas a la puerta `5432` de tu máquina
- Docker recibe la llamada
- la entrega al contenedor correcto
- y responde al cliente

### Ejemplo de RabbitMQ

```yaml
ports:
  - "5672:5672"
  - "15672:15672"
```

- `5672` = puerto de conexión de mensajes
- `15672` = puerto de la UI de administración

Esto significa que Docker está escuchando esos puertos en el host y reenvía el tráfico al contenedor de RabbitMQ.

---

## ¿Qué son las variables de entorno?

Las variables de entorno se usan para configurar un contenedor sin cambiar el código.

Ejemplo:

```yaml
environment:
  POSTGRES_USER: admin
  POSTGRES_PASSWORD: admin
  POSTGRES_DB: webhook_dispatcher
```

Esto configura la base de datos automáticamente al iniciar el contenedor.

---

## ¿Qué es `docker compose up -d`?

Es el comando más usado para levantar el entorno completo definido en un archivo Compose.

### Significado

- `docker` = comando principal
- `compose` = herramienta para múltiples contenedores
- `up` = crear/iniciar
- `-d` = background, en segundo plano

Cuando lo ejecutas, Docker:

- lee el archivo YAML
- crea una red interna
- crea los volúmenes
- descarga las imágenes necesarias
- inicia cada contenedor

---

## ¿Qué cosas se pueden hacer con Docker?

Además de levantar bases de datos y servicios, Docker permite hacer muchas cosas:

### 1. Ejecutar APIs y backend

```yaml
services:
  api:
    build: .
    ports:
      - "8080:8080"
```

### 2. Ejecutar frontends

```yaml
services:
  frontend:
    image: node:20
    working_dir: /app
    command: npm run dev
```

### 3. Ejecutar bases de datos

- PostgreSQL
- MySQL
- Redis
- MongoDB

### 4. Ejecutar colas de mensajes

- RabbitMQ
- Kafka
- Redis Streams

### 5. Crear microservicios

Cada servicio puede estar en un contenedor independiente y comunicarse a través de una red Docker.

### 6. Probar aplicaciones con dependencias difíciles

Por ejemplo:

- versiones específicas de Node
- Java
- Python
- bases de datos con configuraciones particulares

### 7. Desplegar en production

Aunque en producción normalmente se usan herramientas más avanzadas, Docker es la base de muchos despliegues modernos.

---

## ¿Qué es una red en Docker?

Los contenedores pueden comunicarse entre sí a través de redes Docker.

Por ejemplo, tu aplicación backend puede conectarse a la base de datos PostgreSQL usando el nombre del contenedor como host, por ejemplo:

```text
postgres
```

No necesitas usar `localhost` si están dentro de la misma red Docker.

---

## ¿Qué es el archivo `docker-compose.yml` en este proyecto?

En este proyecto, probablemente se usa para levantar servicios como:

- PostgreSQL
- RabbitMQ
- posiblemente la aplicación principal

Un ejemplo de configuración básica sería:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    container_name: webhook-db
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: admin
      POSTGRES_DB: webhook_dispatcher
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  rabbitmq:
    image: rabbitmq:3-management-alpine
    container_name: webhook-queue
    environment:
      RABBITMQ_DEFAULT_USER: admin
      RABBITMQ_DEFAULT_PASS: admin
    ports:
      - "5672:5672"
      - "15672:15672"

volumes:
  pgdata:
```

### ¿Qué hace este ejemplo?

- Levanta PostgreSQL
- Levanta RabbitMQ
- Expone puertos para acceso local
- Guarda la base de datos en un volumen persistente
- Configura usuarios y contraseñas

---

## Comandos útiles de Docker

### Ver información

```bash
docker ps
```

```bash
docker images
```

```bash
docker volume ls
```

```bash
docker network ls
```

### Levantar entorno

```bash
docker compose up -d
```

### Detener entorno

```bash
docker compose down
```

### Ver logs

```bash
docker compose logs -f
```

### Reiniciar servicios

```bash
docker compose restart
```

### Ejecutar un comando dentro del contenedor

```bash
docker exec -it <nombre_del_contenedor> bash
```

Ejemplo:

```bash
docker exec -it webhook-db bash
```

---

## Ventajas de Docker

- Portabilidad
- Reproducibilidad
- Entornos aislados
- Facilidad de despliegue
- Escalabilidad
- Acelera el desarrollo

---

## Desventajas o consideraciones

- Requiere conocer algunos conceptos básicos
- El manejo de redes y almacenamiento puede complicarse
- No sustituye una infraestructura de despliegue completa en producción
- Debe usarse con buenas prácticas de seguridad y mantenimiento

---

## Resumen rápido

Docker sirve para encapsular aplicaciones y servicios en contenedores para que sean fáciles de ejecutar, mover y desplegar.

Un contenedor es una ejecución aislada de una aplicación.

Un archivo `docker-compose.yml` define cómo se levantan varios contenedores juntos.

Los volúmenes permiten guardar datos persistentes.

Los puertos permiten acceder a los servicios desde el host.

Las variables de entorno permiten configurar la aplicación sin cambiar el código.

---

## Siguiente paso recomendado

Prueba estos comandos en tu proyecto:

```bash
cd /home/nf/Desktop/webhookDispatcherEngine/deployments
docker compose -f docker-compose.yml up -d
docker ps
docker compose logs
```

Y luego:

```bash
docker compose down
```

Esto te ayudará a entender cómo Docker levanta y baja los servicios del entorno.

---

## Bibliografía rápida

- Documentación oficial de Docker: https://docs.docker.com/
- Docker Compose: https://docs.docker.com/compose/

---

## Nota final

Docker no es un reemplazo del sistema operativo, sino una herramienta para empaquetar y ejecutar aplicaciones de manera consistente y aislada. Es especialmente útil para proyectos con bases de datos, colas de mensajes, APIs y microservicios.
