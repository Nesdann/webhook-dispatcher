# Clase rápida de criptografía para entender llaves públicas, privadas y SSH

## 1. ¿Qué es la criptografía?

La criptografía es la ciencia de proteger la información para que solo pueda ser leída o validada por las personas o sistemas correctos.

Su objetivo principal es asegurar tres cosas:

- confididencialidad: que nadie pueda leer un mensaje salvo el destinatario correcto
- integridad: que el mensaje no haya sido alterado
- autenticación: que puedas comprobar quién lo envió o quién eres

La criptografía se usa en:

- HTTPS
- SSH
- GitHub/GitLab
- pagos digitales
- autenticación de apps
- mensajes cifrados
- firmas digitales

---

## 2. ¿Qué es un algoritmo criptográfico?

Un algoritmo es una receta matemática para transformar información.

Hay dos grandes familias:

- criptografía simétrica
- criptografía asimétrica

### Criptografía simétrica

La misma clave se usa para cifrar y descifrar.

Ejemplo:

```text
clave secreta = "abc123"
```

Si usas esa misma clave para cifrar y para descifrar, es simétrica.

Problema:

- cómo compartir esa clave sin que alguien la vea

Ejemplo típico:

- AES
- ChaCha20

Se usa mucho en:

- cifrado de archivos
- conexiones rápidas
- sistemas donde ambas partes ya compartieron la clave antes

---

## 3. ¿Qué es la criptografía asimétrica?

La criptografía asimétrica usa dos claves relacionadas:

- clave pública
- clave privada

La clave pública puede compartirse libremente.
La clave privada debe quedarse secreta.

### Idea clave

- con la pública puedes cifrar o verificar
- con la privada puedes descifrar o firmar

Esto da lugar a dos usos muy importantes:

1. cifrado asimétrico
2. firma digital

---

## 4. Cifrado asimétrico

Imaginemos que Alice quiere mandar un mensaje a Bob sin que nadie más lo pueda leer.

### Paso a paso

1. Bob genera un par de claves:
   - pública
   - privada
2. Bob comparte la pública con Alice
3. Alice cifra el mensaje con la pública de Bob
4. Bob recibe el mensaje y lo descifra con su privada

### Importante

Nadie, salvo Bob, puede descifrar ese mensaje porque solo él tiene la privada.

### ¿Cuándo se usa?

- intercambio seguro de información
- autenticación inicial
- intercambio de claves simétricas
- protocolos como SSH, TLS, HTTPS

---

## 5. Firma digital

La firma digital sirve para comprobar:

- que un mensaje fue realmente enviado por una persona
- que el mensaje no fue alterado

### Paso a paso

1. Alice genera una firma usando su clave privada
2. Bob verifica la firma con la clave pública de Alice
3. si la verifica, sabe que el mensaje fue firmado por Alice

### Importante

La llave privada firma, la pública verifica.

Esto es distinto del cifrado, aunque usa matemáticas similares.

---

## 6. ¿Qué es una clave pública y una privada en la práctica?

Las llaves no son cadenas aleatorias de texto raras; son valores matemáticos grandes.

Una clave pública es un valor que se puede compartir.
Una clave privada es el valor secreto que permite actuar como dueño de esa identidad.

### Ejemplo conceptual

```text
Clave pública  = 111111111111111111...
Clave privada  = 999999999999999999...
```

Pero en realidad son números enormes y se usan con funciones matemáticas complejas.

---

## 7. ¿Qué significa que estén relacionadas?

El par de llaves está matemáticamente conectado.

Esto permite:

- que la pública valide lo que la privada firmó
- que la privada descifre lo que la pública cifró

La relación es una propiedad del algoritmo, no una “conexión casual”.

---

## 8. ¿Qué es SSH?

SSH es un protocolo usado para:

- conectarse a servidores de forma segura
- autenticarte sin contraseña
- ejecutar comandos remotos
- transferir archivos de manera protegida

### ¿Por qué necesita criptografía?

Porque hay tres problemas importantes:

- alguien puede escuchar la comunicación
- alguien puede suplantar tu identidad
- alguien puede alterar lo que mandas

SSH resuelve eso usando cifrado y autenticación por llaves.

---

## 9. ¿Cómo funciona SSH con llaves?

Cuando haces algo como:

```bash
ssh git@github.com
```

el servidor quiere saber si tú eres realmente quien dices ser.

### Flujo típico

1. tú tienes una clave privada
2. la clave pública fue subida a GitHub
3. GitHub te pide una prueba que solo tu privada puede responder
4. tu cliente usa esa clave privada
5. GitHub la verifica con la pública que tiene registrada
6. si coincide, te da acceso

Esto se hace sin enviar tu clave privada al servidor.

Es una autenticación segura por prueba criptográfica.

---

## 10. ¿Qué es `ed25519`?

`ed25519` es un algoritmo de firma digital moderno y muy popular.

Es una implementación del esquema de curva elíptica llamado Ed25519.

### ¿Por qué es tan usado?

Porque tiene varias ventajas:

- seguridad fuerte
- tamaño de clave pequeño
- velocidad buena
- muy usado en SSH moderno
- más eficiente que muchos algoritmos antiguos

### Ejemplo

```bash
ssh-keygen -t ed25519 -C "tu-email@ejemplo.com"
```

Eso genera un par de llaves usando el algoritmo `ed25519`.

---

## 11. ¿Por qué se usa Ed25519 en SSH?

Porque es:

- moderno
- seguro
- rápido
- compatible con sistemas nuevos

Muchas nuevas instalaciones usan `ed25519` por defecto.

Es mejor que algoritmos más antiguos como RSA en muchos casos.

---

## 12. ¿Qué es RSA?

RSA es otro algoritmo de clave pública muy famoso.

Tiene la idea de la clave pública y privada, pero funciona con números grandes y una matemática distinta.

Por ejemplo, antiguamente se usaba mucho en SSH.

Pero hoy muchos prefieren `ed25519` porque:

- es más eficiente
- tiene tamaño de clave menor
- es más sencillo de operar

---

## 13. ¿Qué es el “par de llaves”?

Cuando generas una llave SSH, normalmente creas algo así:

```bash
~/.ssh/id_ed25519
~/.ssh/id_ed25519.pub
```

- `id_ed25519` = clave privada
- `id_ed25519.pub` = clave pública

Se usan juntas como un conjunto inseparable.

La privada nunca se comparte.
La pública sí.

---

## 14. ¿Cómo se usa esto en aplicaciones reales?

### Ejemplo 1: acceso a GitHub

- tú subes tu clave pública a tu cuenta GitHub
- GitHub la guarda
- cuando haces `git push`, SSH usa tu clave privada local
- GitHub confirma que esa clave privada corresponde a la pública registrada

### Ejemplo 2: HTTPS seguro

Cuando abres un sitio web con HTTPS:

- el servidor tiene un par de llaves
- la pública se entrega al navegador
- el navegador usa esa clave para verificar la identidad del servidor
- la sesión se cifra con una clave de sesión

### Ejemplo 3: firma de software

Muchas herramientas firman paquetes y releases con una clave privada.
Luego, cualquiera puede verificar la firma con la pública.

---

## 15. ¿Qué es la firma digital en términos más simples?

Es como poner un sello personal en un documento digital.

- solo tú puedes poner ese sello si tienes la privada
- cualquiera puede comprobar que el sello es tuyo con la pública

Y si alguien altera el documento, la firma deja de validar.

---

## 16. ¿Y la “frase de paso”?

Al generar una clave SSH puedes poner una frase de paso.

Ejemplo:

```bash
Enter passphrase (empty for no passphrase):
```

Esto significa:

- si alguien roba tu clave privada, igual no puede usarla sin la frase
- mejora la seguridad

La frase no se comparte con el servidor. Solo desbloquea la clave localmente.

---

## 17. ¿Qué significa “cifrar”?

Cifrar significa transformar un mensaje en algo que no puede entenderse sin la clave correcta.

Es como esconder un mensaje en un formato que solo puede ser leído por la persona indicada.

### Ejemplos de cifrado

- cifrar un archivo con una contraseña
- cifrar una conexión SSH
- cifrar una API key en tránsito

---

## 18. ¿Qué significa “descifrar”?

Es el proceso inverso: recuperar el contenido original usando la clave correcta.

Por ejemplo:

- el servidor descifra lo que tú le enviaste
- tu máquina descifra el contenido recibido

---

## 19. ¿Por qué la clave privada no se comparte?

Porque si la compartes, cualquiera podría:

- autenticarse como tú
- firmar cosas en tu nombre
- acceder a repositorios
- suplantarte

Por eso la clave privada debe quedarse solo en el equipo personal o en un gestor seguro.

---

## 20. ¿Qué diferencia hay entre HTTPS y SSH?

### HTTPS

Usado para navegadores y sitios web.

- autentica servidores
- protege conexiones web
- usa certificados digitales

### SSH

Usado para acceder a servidores y repositorios.

- autentica al usuario con clave pública/privada
- protege la terminal remota
- se usa mucho para Git y servidores

---

## 21. ¿Qué hace la matemática detrás de todo esto?

Detrás hay funciones matemáticas avanzadas de:

- curvas elípticas
- aritmética modular
- producto de números grandes
- funciones hash
- problemas computacionales difíciles de revertir

La idea es que sea fácil verificar algo con la clave pública, pero extremadamente difícil deducir la privada a partir de la pública.

Eso crea la seguridad.

---

## 22. ¿Por qué no basta con usar una contraseña?

Porque una contraseña puede:

- ser robada
- ser reutilizada
- ser interceptada
- ser adivinada o forzada

Las llaves públicas/privadas permiten una autenticación más sólida.

Además, no es necesario enviar la contraseña por la red.

---

## 23. Resumen súper corto

- La criptografía protege información.
- La criptografía asimétrica usa par de claves: pública y privada.
- La pública se comparte.
- La privada se guarda en secreto.
- SSH usa este sistema para autenticarte sin contraseña.
- `ed25519` es un algoritmo moderno, rápido y muy seguro para llaves.
- GitHub/GitLab usan esto para saber si eres tú.

---

## 24. Regla mental fácil

### Para recordar:

- pública = comparto
- privada = guardo en secreto
- SSH = protocolo que usa esto para entrar seguro
- ed25519 = algoritmo moderno para generar esas llaves

---

## 25. Ejemplo final

```text
Tú generas la llave privada.
La subes la pública a GitHub.
Cuando haces push, GitHub dice: “¿Puedes demostrar que tienes la privada correspondiente?”
Tu cliente responde usando la privada.
GitHub verifica con la pública.
Si coincide, te deja entrar.
```

Eso es lo que hace que GitHub “te conozca” sin que tengas que poner tu contraseña cada vez.

---

## 26. Concepto clave para no olvidar

La seguridad no está en “ocultar todo”, sino en usar matemáticas difíciles de romper y claves separadas para identidad y validación.

Eso es lo que convierten a la criptografía asimétrica en la base de SSH, HTTPS, firmas digitales y Git.
