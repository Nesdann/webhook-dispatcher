# Guía rápida de llaves SSH y Git con dos usuarios

## 1. ¿Qué es una llave SSH?

SSH significa Secure Shell. Es un protocolo para conectarte a otra máquina de forma segura.

Las llaves SSH son dos claves relacionadas:

- clave privada
- clave pública

La idea es:

- La clave privada se queda en tu equipo
- La clave pública se comparte con servicios como GitHub, GitLab o Bitbucket

Cuando intentas autenticarte, el servidor verifica que tú tengas la clave privada que corresponde a esa clave pública.

---

## 2. Clave privada vs clave pública

### Clave privada

- se guarda en tu equipo
- nunca la compartes
- sirve para “firmar” o “probar” que eres tú
- se mantiene segura

Ejemplo de nombre típico:

```bash
~/.ssh/id_ed25519
```

### Clave pública

- se comparte públicamente
- se sube al servicio remoto
- se usa para verificar que la clave privada correspondiente es válida

Ejemplo:

```bash
~/.ssh/id_ed25519.pub
```

### Relación matemática

La clave pública y la privada están vinculadas. La pública puede verificar algo que solo la privada puede firmar, pero no se puede obtener la privada a partir de la pública.

Eso hace que sea seguro.

---

## 3. ¿Por qué Git usa SSH?

GitHub y GitLab suelen autenticar con HTTPS o SSH.

Con SSH, no necesitas poner tu usuario y contraseña cada vez. La clave SSH actúa como identidad.

Por eso, si tienes dos cuentas de GitHub o dos perfiles distintos, puedes usar dos claves SSH distintas.

---

## 4. ¿Qué pasa si tengo dos usuarios de Git?

Esto es muy común cuando una persona tiene:

- una cuenta personal
- una cuenta de trabajo
- una cuenta de cliente
- varias organizaciones

Cada cuenta puede tener su propia clave SSH.

La clave SSH es la que se usa para autenticarte al servicio remoto, no el nombre de tu usuario local.

Eso significa que:

- cada cuenta puede tener una clave distinta
- cada clave puede estar asociada a una cuenta distinta
- el repositorio remoto decide si acepta esa clave

---

## 5. ¿Cómo sabe Git qué clave usar?

Git no “adivina” automáticamente entre varias claves.

La elección se hace por la configuración SSH del cliente.

### Lo que importa es:

- qué host estás conectando
- qué archivo de configuración SSH estás usando
- qué clave corresponde a ese host

Cuando haces:

```bash
git clone git@github.com:usuario/repo.git
```

SSH intenta usar la configuración que tenga para el dominio `github.com`.

Si no hay una configuración especial, usa la clave por defecto, normalmente:

```bash
~/.ssh/id_ed25519
```

---

## 6. ¿Cómo se configuran varias llaves?

Se usa el archivo:

```bash
~/.ssh/config
```

Ahí puedes decir:

- para `github.com` usa esta clave
- para `gitlab.com` usa otra
- para un host custom usa otra más

### Ejemplo

```bash
Host github.com-personal
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_personal
    IdentitiesOnly yes

Host github.com-work
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_work
    IdentitiesOnly yes
```

Con eso puedes clonar repositorios usando diferentes aliases:

```bash
git clone git@github.com-personal:tuusuario/tu-repo.git
```

Y otro repo con:

```bash
git clone git@github.com-work:empresa/repo.git
```

---

## 7. ¿Qué es `IdentityFile`?

Es la línea que le dice a SSH qué clave privada usar para ese host.

Ejemplo:

```bash
IdentityFile ~/.ssh/id_ed25519_personal
```

Eso significa: “cuando te conectes a este host, usa esta clave privada”.

---

## 8. ¿Qué es `IdentitiesOnly yes`?

Es una línea importante.

Le dice a SSH:

> “No pruebes otras claves disponibles; usa solo la que te indico”.

Esto evita que SSH intente claves equivocadas y te falle la autenticación.

---

## 9. ¿Cómo se genera una llave SSH?

```bash
ssh-keygen -t ed25519 -C "tu-email@ejemplo.com"
```

Te preguntará dónde guardarla. Normalmente puedes dejar la ruta por defecto.

Ejemplo:

```bash
~/.ssh/id_ed25519
```

Luego te pedirá una frase de paso opcional.

Es recomendable poner una frase de paso para seguridad.

---

## 10. ¿Cómo se ve la llave pública?

```bash
cat ~/.ssh/id_ed25519.pub
```

Copia esa línea y súbela en GitHub o GitLab.

---

## 11. ¿Cómo revisar qué llaves tienes?

```bash
ls -la ~/.ssh
```

Puedes ver algo como:

```bash
id_ed25519
id_ed25519.pub
id_ed25519_work
id_ed25519_work.pub
```

---

## 12. ¿Cómo probar si SSH funciona?

```bash
ssh -T git@github.com
```

O si usas un host custom:

```bash
ssh -T git@github.com-personal
```

Si está bien configurado, te va a responder algo parecido a:

```bash
Hi username! You've successfully authenticated, but GitHub does not provide shell access.
```

---

## 13. ¿Cómo verificar qué clave está usando SSH?

Puedes probar con:

```bash
ssh -vT git@github.com
```

Eso te mostrará qué clave privada está intentando usar.

Si quieres una comprobación más directa, puedes ir a:

```bash
~/.ssh/config
```

y revisar el host correspondiente.

---

## 14. ¿Qué pasa con Git local y el usuario de Git?

Git tiene dos cosas distintas:

1. usuario de Git para los commits
2. autenticación SSH para conectarse al remoto

### Usuario para commits

```bash
git config --global user.name "Tu Nombre"
git config --global user.email "tu-email@ejemplo.com"
```

Esto sirve para los commits, no para autenticación SSH.

### Autenticación SSH

Eso se resuelve con la clave SSH configurada para el host remoto.

No tienen que coincidir necesariamente.

Ejemplo:

- `user.email` = tu email personal
- clave SSH = una clave dedicada a la empresa

Ambos pueden ser diferentes, y eso está bien.

---

## 15. ¿Cómo usar dos cuentas GitHub con dos llaves?

### Paso 1: generar dos llaves

```bash
ssh-keygen -t ed25519 -C "personal@email.com" -f ~/.ssh/id_ed25519_personal
ssh-keygen -t ed25519 -C "work@email.com" -f ~/.ssh/id_ed25519_work
```

### Paso 2: subir las públicas a cada cuenta

```bash
cat ~/.ssh/id_ed25519_personal.pub
cat ~/.ssh/id_ed25519_work.pub
```

Copia cada una y súbela en GitHub en:

- Settings → SSH and GPG keys

### Paso 3: crear configuración SSH

```bash
nano ~/.ssh/config
```

Agrega:

```bash
Host github.com-personal
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_personal
    IdentitiesOnly yes

Host github.com-work
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_work
    IdentitiesOnly yes
```

### Paso 4: clonar repositorios con cada alias

```bash
git clone git@github.com-personal:tuusuario/repo-personal.git
```

```bash
git clone git@github.com-work:empresa/repo-work.git
```

---

## 16. ¿Cómo saber qué llave se está usando al hacer push o pull?

Puedes probar con:

```bash
ssh -T git@github.com-personal
ssh -T git@github.com-work
```

Eso te muestra qué cuenta está autenticando.

También puedes revisar la configuración de SSH:

```bash
cat ~/.ssh/config
```

---

## 17. ¿Qué pasa si Git usa la llave equivocada?

Entonces te puede tocar esto:

- “Permission denied (publickey)”
- no te reconoce la cuenta
- no puedes hacer push
- GitHub dice que la clave no está asociada a ese usuario

Eso normalmente significa que:

- la clave pública no está subida a la cuenta correcta
- la configuración SSH apunta a la clave equivocada
- no hay `Host` definido para ese dominio

---

## 18. Recomendación práctica

Lo mejor es separar por cuenta:

- personal → `id_ed25519_personal`
- trabajo → `id_ed25519_work`

Y dejar la configuración SSH como un mapa claro de qué host usa qué llave.

---

## 19. Comandos rápidos para revisar tu setup

### Ver las llaves existentes

```bash
ls -la ~/.ssh
```

### Ver la configuración SSH

```bash
cat ~/.ssh/config
```

### Generar una nueva llave

```bash
ssh-keygen -t ed25519 -C "tu-email@example.com"
```

### Mostrar la clave pública

```bash
cat ~/.ssh/id_ed25519.pub
```

### Probar autenticación con GitHub

```bash
ssh -T git@github.com
```

### Ver a qué clave está intentando conectarse SSH

```bash
ssh -vT git@github.com
```

---

## 20. Resumen final

La clave pública y la privada funcionan como un par:

- la pública se comparte
- la privada se guarda en secreto
- el servidor verifica que la privada corresponde a la pública

En Git, esto sirve para autenticarte sin contraseña.

Si tienes varias cuentas, puedes usar varias llaves y configurarlas por host en `~/.ssh/config`.

Eso te permite tener una identidad para trabajo y otra para personal, sin mezclar cuentas ni permisos.

---

## 21. Comandos que te puedo pedir que me tires para revisar tu caso real

Si quieres, te puedo guiar paso a paso con estos comandos:

```bash
ls -la ~/.ssh
cat ~/.ssh/config
ssh -T git@github.com
ssh -vT git@github.com
git config --global --list
```

Con esa salida te digo exactamente:

- qué llaves tienes
- cuál está configurada por defecto
- qué cuenta está usando GitHub
- qué debes cambiar para usar una u otra clave

---

## 22. Concepto clave para recordar

La llave pública se sube al servicio remoto.
La privada nunca sale de tu equipo.

Git usa SSH para autenticarte, y SSH usa la configuración del cliente para decidir qué clave privada emplear.

Eso es lo que permite tener dos usuarios y dos llaves sin conflicto.
